import Vue from 'vue'
import VueRouter from 'vue-router'
import SignUpIn from '../views/SignUpIn.vue'
import About from '../views/About.vue'
import PublicHome from '../views/PublicHome.vue'

import store from '../store/index.js'

Vue.use(VueRouter)

const routes = [{
    path: '/',
    name: 'PublicHome',
    component: PublicHome
  },
  {
    path: '/login',
    name: 'SignUpIn',
    component: SignUpIn
  },
  {
    path: '/about',
    name: 'About',
    component: About
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

router.beforeEach((to, from, next) => {
  const publicPages = ['/', '/login', '/about'];
  const authRequired = !publicPages.includes(to.path);
  const loggedIn = store.state.authenticated

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