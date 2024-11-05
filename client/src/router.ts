import { createWebHistory, createRouter } from 'vue-router'

import Main from './pages/Main.vue'
import Login from './pages/Login.vue'
import Register from './pages/Register.vue'
import ChatContainer from './pages/ChatContainer.vue'
import ForgotPassword from './pages/ForgotPassword.vue'
import NumbersConfirmation from './pages/NumbersConfirmation.vue'
import PasswordReset from './pages/PasswordReset.vue'


const routes = [
  { path: '/', component: Main },
  { path: '/login', component: Login },
  { path: '/register', component: Register },
  { path: '/chat/:roomID/:name', component: ChatContainer },
  { path: '/forgot_password', component: ForgotPassword },
  { path: '/confirmation', component: NumbersConfirmation },
  { path: '/password_reset', component: PasswordReset },

]

const router = createRouter({
  history: createWebHistory('/private-chats'),
  routes,
})

export default router