import Vue from 'vue'
import App from './App.vue'
import router from './router.js'
import VueToasted from 'vue-toasted'
import VueMeta from 'vue-meta'
import axios from 'axios'
import VueAxios from 'vue-axios'
import { BootstrapVue, IconsPlugin } from 'bootstrap-vue'



// Import Bootstrap an BootstrapVue CSS files (order is important)
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

import '../public/assets/materialdesignicons/css/materialdesignicons.min.css'
// Make BootstrapVue available throughout your project
Vue.use(BootstrapVue)
Vue.use(VueToasted)
Vue.use(IconsPlugin)
axios.defaults.baseURL = process.env.VUE_APP_BASE_URL
Vue.use(VueAxios, axios)

Vue.use(VueMeta)
Vue.config.productionTip = false

new Vue({
  render: h => h(App),
  router,
  VueMeta
}).$mount('#app')
