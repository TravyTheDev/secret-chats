<template>
    <div @click="emit('close-modal')" class="absolute w-full bg-black bg-opacity-30 h-screen top-0 left-0 flex flex-col px-8">
        <div @click.stop class="relative p-4 bg-white text-black self-start m-auto top-0 bottom-0 w-auto">
            <span @click="emit('close-modal')" class="absolute top-0 right-0 -translate-y-1 -translate-x-1 text-2xl hover:cursor-pointer">&#215;</span>
            <div class="py-2">
                <div>
                    <p>{{t('invites.email')}}: </p>
                    <input class="border" v-model="searchText" type="text">
                </div>
                <div>
                    <p>{{t('invites.message')}}: </p>
                    <textarea maxlength="40" class="border w-full" v-model="message" type="text"></textarea>
                </div>
            </div>
            <p class="hover:cursor-pointer" v-if="user" @click="handleInvite(user.id)">{{ user.username }}</p>
        </div>
    </div>
</template>

<script setup lang="ts">
import { inject, ref, watch } from 'vue';
import { axiosInstance } from '../api';
import { User } from '../App.vue';
import type { Ref } from 'vue'
import { useI18n } from 'vue-i18n';

const {t} = useI18n()

interface props {
    roomID: string | string[];
    roomName: string | string[];
}

const loginUser = inject<Ref<User | undefined>>("loginUser")

const props = defineProps<props>()

const emit = defineEmits(['close-modal'])
const searchText = ref('')
const message = ref('')
const user = ref<User>()

const searchUser = async (val: string) => {
    const res = await axiosInstance.get(`/search_user/${val}`)
    user.value = res.data
}

const handleSearch = (value: string) => {
    const timeoutID: number = window.setTimeout(() => {}, 0)

    for (let id:number = timeoutID; id >= 0; id -= 1) {
        window.clearTimeout(id)
    }

    setTimeout(() => {
        searchUser(value)
    }, 300)
}

watch(searchText, () => {
    if(searchText.value){
        handleSearch(searchText.value)
    }
})

const handleInvite = async (id: number) => {
    await axiosInstance.post(`/sse/send/${id}`, {
        name: loginUser?.value?.username,
        message: message.value,
        roomID: props.roomID,
        roomName: props.roomName
    })
    emit('close-modal')
}

</script>

<style scoped>

</style>