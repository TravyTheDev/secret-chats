package ws

import (
	"encoding/json"
	"fmt"
	"net/http"
	"secret-chats/service/jwt"
	"secret-chats/types"
	"secret-chats/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type WSHandler struct {
	hub      *Hub
	jwtMaker *jwt.JWTMaker
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func NewWSHandler(hub *Hub, secretKey string) *WSHandler {
	return &WSHandler{
		hub:      hub,
		jwtMaker: jwt.NewJWTMaker(secretKey),
	}
}

func (h *WSHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/ws/create_room", h.createRoom).Methods("POST")
	router.HandleFunc("/ws/get_rooms", h.getRooms).Methods("GET")
	router.HandleFunc("/ws/{room_id}/get_clients", h.getClients).Methods("GET")
	router.HandleFunc("/ws/join_room/{room_id}/{user_id}/{username}", h.joinRoom).Methods("GET")
}

func (h *WSHandler) createRoom(w http.ResponseWriter, r *http.Request) {
	var req *types.CreateRoomReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "error creating room", http.StatusBadRequest)
		return
	}

	if err := utils.Validate.Struct(req); err != nil {
		errors := err.(validator.ValidationErrors)
		if errors != nil {
			http.Error(w, "room must have a name", http.StatusBadRequest)
		}
		return
	}

	h.hub.Rooms[req.ID] = &Room{
		ID:      req.ID,
		Name:    req.Name,
		Clients: make(map[string]*Client),
	}

	json.NewEncoder(w).Encode(req)
}

func (h *WSHandler) joinRoom(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
	}

	roomID := mux.Vars(r)["room_id"]
	userID := mux.Vars(r)["user_id"]
	username := mux.Vars(r)["username"]

	cl := &Client{
		Conn:     conn,
		Message:  make(chan *Message, 10),
		ID:       string(userID),
		RoomID:   roomID,
		Username: username,
	}

	m := &Message{
		Body:     fmt.Sprintf("%s has joined the room", username),
		RoomID:   roomID,
		Username: username,
		UserID:   cl.ID,
	}

	h.hub.Register <- cl
	if len(h.hub.Rooms[roomID].Clients) != 0 {
		h.hub.Broadcast <- m
	}

	go cl.writeMessage()
	cl.readMessage(h.hub)
}

func (h *WSHandler) getRooms(w http.ResponseWriter, r *http.Request) {
	rooms := make([]types.RoomRes, 0)

	for _, room := range h.hub.Rooms {
		rooms = append(rooms, types.RoomRes{
			ID:   room.ID,
			Name: room.Name,
		})
	}

	if err := json.NewEncoder(w).Encode(rooms); err != nil {
		fmt.Println(err)
		return
	}
}

func (h *WSHandler) getClients(w http.ResponseWriter, r *http.Request) {
	var clients []types.ClientRes

	roomID := mux.Vars(r)["room_id"]

	if _, ok := h.hub.Rooms[roomID]; ok {
		clients = make([]types.ClientRes, 0)
	}
	for _, client := range h.hub.Rooms[roomID].Clients {
		clients = append(clients, types.ClientRes{
			ID:       client.ID,
			Username: client.Username,
		})
	}
	if err := json.NewEncoder(w).Encode(clients); err != nil {
		fmt.Println(err)
		return
	}
}
