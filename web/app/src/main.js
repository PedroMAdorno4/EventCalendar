import Vue from 'vue'
import {
  BVConfigPlugin,
  BootstrapVue,
  BootstrapVueIcons
} from 'bootstrap-vue'
import VueCookies from 'vue-cookies'
import Vuelidate from 'vuelidate'
import ScrollView from 'vue-scrollview'

import './config.scss'

import App from './App.vue'
import router from './router'
import store from './store'

Vue.use(BootstrapVue)
Vue.use(BootstrapVueIcons)
Vue.use(VueCookies)
Vue.use(Vuelidate)
Vue.use(ScrollView)

Vue.config.productionTip = false

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
