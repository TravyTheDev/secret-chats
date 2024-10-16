<template>
    <div @click="emit('close-modal')" class="absolute w-full bg-black bg-opacity-30 h-screen top-0 left-0 flex flex-col px-8">
        <div @click.stop class="relative flex flex-col p-4 bg-white self-start m-auto top-0 bottom-0 w-auto">
            <span @click="emit('close-modal')" class="absolute top-0 right-0 -translate-x-1 text-2xl hover:cursor-pointer">&#215;</span>
            <span>{{t('create_room.room_name')}}:</span>
            <input class="border" v-model="roomName" type="text">
            <button class="self-end border p-1 mt-2" @click="createRoom">{{t('create_room.create')}}</button>
            <span class="text-sm text-red-500">{{ errorValue }}</span>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { axiosInstance } from '../api';
import { v4 as uuidv4 } from "uuid";
import { useRouter } from 'vue-router';
import { useI18n } from 'vue-i18n';

const {t} = useI18n()

const router = useRouter()

const emit = defineEmits(['close-modal'])

const roomName = ref("")
const errorValue = ref('')

const joinRoom = (id: string, name: string) => {
    router.push(`/chat/${id}/${name}`)
}

const createRoom = async () => {
    try {
        const roomID = uuidv4()
        await axiosInstance.post('/ws/create_room', {
            id: roomID,
            name: roomName.value
        })
        joinRoom(roomID, roomName.value)
    } catch (error: any) {
        errorValue.value = error.response.data
    }
}
</script>

<style scoped>

</style>