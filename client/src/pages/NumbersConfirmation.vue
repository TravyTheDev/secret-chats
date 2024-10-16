<template>
    <div class="h-screen flex flex-col text-fontColor">
        <div class="m-auto flex flex-col gap-1">
            <p>{{ t('auth.email_numbers') }}</p>
            <input class="text-black" v-model="numbers" type="text">
            <span class="text-sm text-red-500">{{ errorValue }}</span>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue';
import { useRouter } from 'vue-router';
import { axiosInstance } from '../api';
import { useI18n } from 'vue-i18n';

const { t } = useI18n()
const router = useRouter()

const numbers = ref()
const errorValue = ref('')

const checkNumbers = async () => {
    try {
        await axiosInstance.post('/confirm_numbers', {
            numbers: Number(numbers.value)
        })
        router.push('/password_reset')
    } catch (error: any) {
        errorValue.value = error.response.data
    }
}

watch(numbers, () => {
    if (numbers.value && String(numbers.value).length === 5) {
        checkNumbers()
    }
})
</script>

<style scoped></style>