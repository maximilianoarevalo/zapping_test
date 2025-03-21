// src/router.js
import { createRouter, createWebHistory } from 'vue-router';
import CreateUser from './components/CreateUser.vue';
import LoginPage from './components/LoginPage.vue';
import HlsPlayer from './components/HlsPlayer.vue';

const routes = [
  { path: '/', component: LoginPage, name: 'login', meta: { requiresAuth: false } },
  { path: '/crear-cuenta', component: CreateUser, name: 'crear-cuenta', meta: { requiresAuth: false } },
  { path: '/player', component: HlsPlayer, name: 'player', meta: { requiresAuth: true } }, // Proteger ruta /player con auth y token
  { path: '/:pathMatch(.*)*', redirect: '/' } // Regex de redireccion rutas externas ingresadas por usuario
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

// Middleware proteccion urls
router.beforeEach((to, from, next) => {
  const token = localStorage.getItem("token");

  // Si no hay token de autenticidad redireccionar al login
  if (to.meta.requiresAuth && !token) {
    next("/login");
  }
  else {
    next();
  }
});

export default router;
