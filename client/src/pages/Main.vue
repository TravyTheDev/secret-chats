<template>
  <div class="h-full overflow-hidden">
    <div class="text-fontColor flex flex-col items-center h-full">
      <div class="mt-20 text-lg mx-2">
        <h1 v-if="loginUser">{{ t('main.hello') }} {{ loginUser.username }}!</h1>
        <h2 v-if="loginUser">{{ t('main.create') }}</h2>
        <h2 v-else>{{ t('main.login_or_register') }}</h2>
        <p>{{ t('main.rooms_display_info') }}</p>
        <p>{{ t('main.invite_info') }}</p>
        <p>{{ t('main.display_invite_info') }}</p>
        <p>{{ t('main.message_info') }}</p>
        <p>{{ t('main.room_info') }}</p>
        <div v-if="loginUser" class="w-full flex justify-center mt-4">
          <button class="bg-myMessage px-2 py-1 rounded-lg hover:border border-slate-400" @click="toggleCreate">{{
            t('main.create_button') }}</button>
        </div>
      </div>
    </div>
    <CreateRoomModal v-if="isCreate" @close-modal="toggleCreate" />
  </div>
</template>

<script setup lang="ts">
import { inject, ref } from 'vue';
import { User } from '../App.vue';
import type { Ref } from 'vue'
import CreateRoomModal from '../components/CreateRoomModal.vue';
import { useI18n } from 'vue-i18n';

const { t } = useI18n()

const isCreate = ref(false)

const loginUser = inject<Ref<User | undefined>>("loginUser")
const toggleCreate = () => {
  isCreate.value = !isCreate.value
}

</script>

<style scoped></style>
