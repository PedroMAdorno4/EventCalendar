import Vue from 'vue'
import {
  BVConfigPlugin,
  BootstrapVue
} from 'bootstrap-vue'
import VueCookies from 'vue-cookies'

import './config.scss'

import App from './App.vue'
import router from './router'
import store from './store'

Vue.use(BootstrapVue)
Vue.use(VueCookies)

Vue.config.productionTip = false

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
