import Vue from 'vue'
import VueRouter from 'vue-router'
import SignUpIn from '../views/SignUpIn.vue'
import Events from '../views/Events.vue';

import store from '../store/index.js'

Vue.use(VueRouter)

const routes = [
  {
    path: '/login',
    name: 'Entrar',
    component: SignUpIn
  },
  {
    path: '/events',
    name: 'Eventos',
    component: Events
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

router.beforeEach((to, from, next) => {
  const publicPages = ['/login'];
  const authRequired = !publicPages.includes(to.path);
  const loggedIn = store.state.authenticated || $cookies.isKey("token");

  Vue.nextTick(() => {
    document.title = to.name + ' - Calendário' || 'Calendário';
  });

  if (authRequired && !loggedIn) {
    next('/login');
  } else {
    next();
  }
});


export default router