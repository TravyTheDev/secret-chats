<template>
  <div class="flex flex-col overflow-y-hidden mx-auto max-w-screen-md bg-chatBackground text-fontColor">
    <div ref="chatBody" class="mx-2 h-[80vh] pb-4 overflow-y-auto chat-body whitespace-pre-line">
      <div class="flex flex-col mx-2" v-for="message in messages">
        <div class="self-end flex flex-col items-end max-w-[66%]" v-if="loginUser && message.userID == loginUser.id">
          <p class="font-semibold text-sm">{{ message.name }}:</p>
          <p class="border border-myMessage rounded-l-lg rounded-br-lg bg-myMessage px-2 py-1 break-words w-full">{{
            message.body }}</p>
        </div>
        <div class="self-start max-w-[66%]" v-else>
          <p class="font-semibold text-sm">{{ message.name }}:</p>
          <p class="break-words w-full border border-theirMessage rounded-r-lg rounded-bl-lg bg-theirMessage px-2 py-1">
            {{
              message.body }}</p>
        </div>
      </div>
    </div>
    <div v-if="conn" class="w-full bg-backgroundColor fixed bottom-0">
      <span class="ml-2">{{ t('chats.current_room') }}: {{ roomName }}</span>
      <div class="bg-backgroundColor py-2 flex gap-2 max-w-screen-md mx-2 text-nowrap">
        <button @click="toggleIsShowInvite">{{ t('chats.invite') }}</button>
        <textarea @input="resize" ref="textArea" rows="1"
          class="border rounded-sm w-full bg-chatBackground my-0 py-2 overflow-y-hidden"
          v-model="sentMessage"></textarea>
        <button @click="send">{{ t('chats.send') }}>></button>
      </div>
    </div>
    <InviteModal :roomName="roomName" :user="loginUser" :roomID="roomID" v-if="isShowInvite"
      @close-modal="toggleIsShowInvite" @clear="clear" />
  </div>
</template>

<script setup lang="ts">
import { nextTick, onMounted, onUnmounted, ref, watch } from "vue";
import { useRoute } from "vue-router";
import InviteModal from "./InviteModal.vue";
import { useI18n } from 'vue-i18n';
import { v4 as uuidv4 } from "uuid";

const { t } = useI18n()

const props = defineProps({
  loginUser: Object,
})

const route = useRoute();
const roomID = route.params.roomID;
const roomName = route.params.name;
const websocketAddr = import.meta.env.VITE_APP_WS_ADDR
// const PORT = import.meta.env.VITE_APP_PORT
const chatBody = ref()
const textArea = ref()

const isShowInvite = ref(false)

type Message = {
  messageID: string;
  userID: string;
  name: string;
  body: string;
};
const messages = ref<Message[]>([]);
const sentMessage = ref("");

const conn = ref();

const setWebsocket = () => {
  if (props.loginUser) {
    conn.value = new WebSocket(
      `${websocketAddr}/api/v1/ws/join_room/${roomID}/${props.loginUser.id}/${props.loginUser.username}`
    );
  }
}

onMounted(() => {
  setWebsocket()
})

watch(conn, () => {
  if (conn.value) {
    conn.value.onmessage = (e: MessageEvent) => {
      displayMessage(e);
    };
  }
});

const send = () => {
  try {
    conn.value.send(sentMessage.value);
    sentMessage.value = "";
    nextTick(() => {
      resize()
      if (window.innerWidth > 640) {
        textArea.value.focus()
      }
    })
  } catch (error) {
    console.log(error);
  }
};

const displayMessage = (e: MessageEvent) => {
  const data = JSON.parse(e.data);
  const message: Message = {
    messageID: uuidv4(),
    userID: data.userID,
    body: data.body,
    name: data.username
  }
  messages.value.push(message);
  scrollToBottom(data.userID)
  setTimeout(() => {
    removeMessage(message.messageID)
  }, 30000)
};

const clear = () => {
  messages.value.length = 0
}

const removeMessage = (id: string) => {
  const filtered = messages.value.filter((item) => item.messageID !== id)
  messages.value = filtered
}

const toggleIsShowInvite = () => {
  isShowInvite.value = !isShowInvite.value
}

onUnmounted(() => {
  conn.value.close()
})

const scrollToBottom = (id: number) => {
  if ((Number(id) === props.loginUser?.id) || chatBody.value.scrollTop + 5 > chatBody.value.scrollHeight - chatBody.value.clientHeight) {
    nextTick(() => {
      chatBody.value.scrollTop = chatBody.value.scrollHeight
    })
  }
}

const resize = () => {
  textArea.value.style.height = 'auto'
  if (textArea.value.scrollHeight < 104) {
    textArea.value.style.height = textArea.value.scrollHeight + 'px'
  } else {
    textArea.value.style.height = '6.5rem'
  }
  if (textArea.value.scrollHeight > 40) {
    textArea.value.classList.add('expanded-text')
  } else {
    textArea.value.classList.remove('expanded-text')
  }
}
</script>

<style scoped>
.chat-body {
  scroll-behavior: smooth;
}

.expanded-text {
  overflow-y: auto;
}
</style>
