<template>
    <div class="h-screen flex flex-col text-fontColor">
        <div class="m-auto">
            <div>
                <p>{{ t('auth.username') }}: </p>
                <input class="border text-black" v-model="username" type="text">
            </div>
            <div>
                <p>{{ t('auth.email') }}: </p>
                <input class="border text-black" v-model="email" type="text">
            </div>
            <div class="last-name">
                <p>last name: </p>
                <input class="border text-black" v-model="lastName" type="text">
            </div>
            <div>
                <p>{{ t('auth.password') }}: </p>
                <input class="border text-black" v-model="password" type="password">
            </div>
            <div>
                <p>{{ t('auth.confirm_password') }}: </p>
                <input class="border text-black" v-model="passwordConfirm" type="password">
            </div>
            <button class="bg-myMessage px-2 py-1 rounded-lg hover:border border-slate-400 mt-2" @click="register">
                {{ t('auth.register') }}
            </button>
            <p class="text-sm text-red-500 whitespace-pre-line" v-if="isShowError">{{ errorValue }}</p>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { axiosInstance } from '../api';
import { useI18n } from 'vue-i18n';

const { t } = useI18n()

const email = ref('')
const username = ref('')
const password = ref('')
const passwordConfirm = ref('')
const lastName = ref('')
const isShowError = ref(false)
const errorValue = ref('')

const register = async () => {
    isShowError.value = false
    const language = localStorage.getItem("LANGUAGE") ?? "en"
    if (lastName.value) {
        return
    }
    try {
        await axiosInstance.post(`/register/${language}`, {
            email: email.value,
            username: username.value,
            password: password.value,
            passwordConfirm: passwordConfirm.value,
        })
        window.location.href = "/private-chats/"
    } catch (error: any) {
        isShowError.value = true
        errorValue.value = error.response.data
    }
}

</script>

<style scoped>
.last-name {
    display: none;
}
</style>