<script setup lang="ts">
import { ref } from 'vue';
import { axiosInstance } from '../api';
import { useRouter } from 'vue-router';
import { useI18n } from 'vue-i18n';

const {t} = useI18n()
const router = useRouter()


const email = ref('')
const errorValue = ref('')

const requestPasswordReset = async () => {
    try {
        await axiosInstance.post('/mailer/forgot_password', {
            email: email.value
        })
        router.push('/confirmation')
    } catch (error: any) {
        errorValue.value = error.response.data
    }
}

</script>

<template>
    <div class="h-screen flex flex-col text-fontColor">
        <div class="m-auto flex flex-col gap-1">
            <p>{{t('auth.email')}}:</p>
            <input class="text-black" v-model="email" type="text">
            <button class="self-end bg-myMessage px-2 py-1 rounded-lg hover:border border-slate-400" @click="requestPasswordReset">{{ t('auth.send_email') }}</button>
            <span class="text-sm text-red-500">{{ errorValue }}</span>
        </div>
    </div>
</template>

<style scoped>

</style>