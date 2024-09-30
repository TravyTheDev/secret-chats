<script setup lang="ts">
import { inject } from 'vue';
import ChatRoom from '../components/ChatRoom.vue';
import { User } from '../App.vue';
import type { Ref } from 'vue'

const loginUser = inject<Ref<User | undefined>>("loginUser")

</script>

<template>
    <div>
        <ChatRoom v-if="loginUser" :loginUser="loginUser" />
    </div>
</template>

<style scoped></style>