<script setup lang="ts">
import { ref } from 'vue';
import { axiosInstance } from '../api';
import { useI18n } from 'vue-i18n';

const {t} = useI18n()

const email = ref('')
const password = ref('')
const errorValue = ref('')
const isShowError = ref(false)

const login = async () => {
    isShowError.value = false
    const language = localStorage.getItem("LANGUAGE") ?? "en"
    try {
        await axiosInstance.post(`/login/${language}`, {
            email: email.value,
            password: password.value,
        })
        window.location.href = "/"
    } catch (error: any) {
        isShowError.value = true
        errorValue.value = error.response.data
    }
}
</script>

<template>
    <div class="h-screen flex flex-col text-fontColor">
        <div class="m-auto">
            <div>
                <p>{{t('auth.email')}}: </p>
                <input class="border text-black" v-model="email" type="text">
            </div>
            <div>
                <p>{{t('auth.password')}}: </p>
                <input class="border text-black" v-model="password" type="password">
            </div>
            <button class="bg-myMessage px-2 py-1 rounded-lg hover:border border-slate-400 mt-2" @click="login">
                {{t('auth.login')}}
            </button>
            <br />
            <RouterLink to="/forgot_password" class="underline">{{t('auth.forgot_password')}}</RouterLink>
            <p class="text-sm text-red-500 whitespace-pre-line" v-if="isShowError">{{ errorValue }}</p>
        </div>
    </div>
</template>



<style scoped></style>