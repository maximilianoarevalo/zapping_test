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

  // Proteger ruta login y crear-cuenta -> logeados no pueden acceder
  if (token && (to.path === '/' || to.path === '/crear-cuenta')) {
    next('/player'); // Se redirige al player
  // Si no esta logeado -> rederigir a login, permitir /crear-cuenta
  } else if (to.meta.requiresAuth && !token) {
    next('/');
  } else {
    next(); // Caso extra
  }
});

export default router;
