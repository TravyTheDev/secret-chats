<template>
    <div class="fixed bottom-0 right-0 flex bg-white flex-col border px-4 z-10">
        <div @click.stop class="relative p-4 m-auto w-auto">
            <div class="absolute top-0 right-0 translate-x-2 text-2xl font-semibold hover:cursor-pointer" @click="emit('close-modal')">
                <span>&#215;</span>
            </div>
            <div>
                <p>{{t('invites.invite_from')}}: {{ message.name }}</p>
                <p>{{ message.message }}</p>
                <button class="border" @click="joinRoom(message.roomID, message.roomName)">{{t('invites.join')}}</button>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';
import { useI18n } from 'vue-i18n';

const {t} = useI18n()

const router = useRouter()

type Message = {
    name: string;
    message: string;
    roomID: string;
    roomName: string
}

interface props {
    message: Message
}

const props = defineProps<props>()

const emit = defineEmits(['close-modal'])

const joinRoom = (id: string, name:string) => {
    router.push(`/chat/${id}/${name}`)
    emit('close-modal')
}

</script>

<style scoped>

</style>