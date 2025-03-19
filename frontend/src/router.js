// src/router.js
import { createRouter, createWebHistory } from 'vue-router';
import CreateUser from './components/CreateUser.vue';
import LoginPage from './components/LoginPage.vue';
import HlsPlayer from './components/HlsPlayer.vue';

const routes = [
  { path: '/', component: LoginPage, name: 'login' },
  { path: '/crear-cuenta', component: CreateUser, name: 'crear-cuenta' },
  { path: '/player', component: HlsPlayer, name: 'player' }
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router;
