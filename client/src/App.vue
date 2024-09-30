<script setup lang="ts">
import { RouterView } from 'vue-router';
import { onMounted, provide, ref } from 'vue';
import { axiosInstance } from './api';
import Navbar from './components/Navbar.vue';

export type User = {
  id: 0,
  username: "",
  email: "",
}

const loginUser = ref<User>()
const getUser = async () => {
  try {
    const res = await axiosInstance.get("/me")
    loginUser.value = res.data
  } catch (error) {
    console.log(error)
  }
}
provide("loginUser", loginUser)

onMounted(() => {
  getUser()
  const theme = localStorage.getItem("THEME")
    if (theme) {
      document.body.classList.add(theme)
    } else {
      document.body.classList.add("light")
    }
})

const getItemKey = (item: any) => item;

</script>

<template>
  <div class="bg-backgroundColor h-screen overflow-hidden">
    <Navbar :key="getItemKey(loginUser)" />
    <RouterView />
  </div>
</template>

<style scoped>

</style>
