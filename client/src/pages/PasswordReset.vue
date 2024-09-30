<script setup lang="ts">
import { ref } from 'vue';
import { axiosInstance } from '../api';
import { useRouter } from 'vue-router';
import { useI18n } from 'vue-i18n';

const {t} = useI18n()
const router = useRouter()

const password = ref('')
const passwordConfirm = ref('')

const changePassword = async () => {
    try {
        await axiosInstance.post('/change_password', {
        password: password.value,
        passwordConfirm: passwordConfirm.value
    })
    router.push('/login')
    } catch (error) {
        console.log(error)
    }
    
}
</script>

<template>
    <div class="h-screen flex flex-col text-fontColor">
        <div class="m-auto flex flex-col gap-1">
            <p>{{t('auth.password')}}:</p>
            <input class="text-black" v-model="password" type="password">
            <p>{{t('auth.confirm_password')}}:</p>
            <input class="text-black" v-model="passwordConfirm" type="password">
            <button class="self-end bg-myMessage px-2 py-1 rounded-lg hover:border border-slate-400" @click="changePassword">send</button>
        </div>
    </div>
</template>

<style scoped>

</style>