import axios from "axios";

const baseUrl = import.meta.env.VITE_APP_BASE_URL
// const PORT = import.meta.env.VITE_APP_PORT

export const axiosInstance = axios.create({
  baseURL: `${baseUrl}/api/v1`,
  headers: {
    "Content-Type": "application/json",
  },
  withCredentials: true,
});

axiosInstance.interceptors.response.use(
  (response) => response, 
  async (error) => {
    const originalRequest = error.config;
    if (error.response.status === 401 && !originalRequest.retry) {
        originalRequest.retry = true;
      try {
        await fetch(`${baseUrl}/api/v1/renew_token`, {
            method: "POST",
            credentials: "include"
        });

        return axiosInstance(originalRequest); 
      } catch (refreshError) {
        console.error("Token refresh failed:", refreshError);
        
        return Promise.reject(refreshError);
      } 
    }
    return Promise.reject(error);
  }
);
