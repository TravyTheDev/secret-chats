<template>
    <div class="relative p-2 border border-chatBackground max-w-screen-md mx-auto text-sm">
        <div class="flex justify-between text-fontColor">
            <div>
                <RouterLink to="/">
                    << {{ t('navbar.home') }}</RouterLink>
            </div>
            <div>
                <label v-if="colorTheme == 'light' || !colorTheme" for="dark_mode">{{ t('navbar.dark_mode') }}</label>
                <input class="hidden" value="dark" v-model="colorTheme" name="theme_choice" id="dark_mode" type="radio">
                <label v-if="colorTheme == 'dark'" for="light_mode">{{t('navbar.light_mode')}}</label>
                <input class="hidden" value="light" v-model="colorTheme" name="theme_choice" id="light_mode"
                    type="radio">
            </div>
            <div>
                <select class="bg-backgroundColor text-fontColor" v-model="selectedLanguage">
                    <option v-for="lang in languages" :value="lang.value">{{ lang.label }}</option>
                </select>
            </div>
            <span v-if="loginUser" @click="logout"
                class="hover:cursor-pointer border px-1 rounded-lg border-red-300">{{t('navbar.logout')}}</span>
            <div v-else>
                <RouterLink class="border px-1 rounded-lg border-red-300 mr-1" to="/login">{{t('navbar.login')}}</RouterLink>
                <RouterLink class="border px-1 rounded-lg border-red-300" to="/register">{{t('navbar.register')}}</RouterLink>
            </div>
        </div>
        <InvitationModal v-if="messageData" :message="messageData" @close-modal="messageData = ''" />
    </div>
</template>

<script setup lang="ts">
import { inject, onMounted, ref, Ref, watch } from 'vue';
import { User } from '../App.vue';
import { axiosInstance } from '../api';
import InvitationModal from './InvitationModal.vue';
import { useI18n } from 'vue-i18n';
import { languages } from '../consts';

const {t} = useI18n()
const {locale} = useI18n()
const selectedLanguage = ref(localStorage.getItem("LANGUAGE") ?? "en")
locale.value = selectedLanguage.value

const loginUser = inject<Ref<User | undefined>>("loginUser")
const eventSourceAddr = import.meta.env.VITE_APP_BASE_URL

const logout = async () => {
    await axiosInstance.post('/logout')
    window.location.href = "/"
}

const colorTheme = ref("")
const timeout = ref()
const evtSource = ref()

const setEventSource = () => {
    if (loginUser && loginUser.value) {
        evtSource.value = new EventSource(
            `${eventSourceAddr}:8000/api/v1/sse/stream/${loginUser.value.id}`
        )
    }
}

onMounted(() => {
    setEventSource()
    const theme = localStorage.getItem("THEME")
    if (theme) {
        colorTheme.value = theme
    }
})

const messageData = ref()

watch(evtSource, () => {
    if (evtSource.value) {
        evtSource.value.onmessage = (e: MessageEvent) => {
            clearTimeout(timeout.value)
            const data = JSON.parse(e.data)
            messageData.value = data
            timeout.value = setTimeout(() => {
                messageData.value = undefined
            }, 10000)
        }
        evtSource.value.onerror = (err: MessageEvent) => {
            console.error("event source failed", err)
        }
    }
})

watch(colorTheme, (newVal, oldVal) => {
    if (colorTheme.value) {
        localStorage.setItem("THEME", colorTheme.value)
        document.body.classList.add(newVal)
        if (document.body.classList.contains(oldVal)) {
            document.body.classList.remove(oldVal)
        }
    }
})

watch(selectedLanguage, () => {
    if(selectedLanguage.value){
        locale.value = selectedLanguage.value
        localStorage.setItem("LANGUAGE", selectedLanguage.value)
    }
})

</script>

<style scoped></style>