// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import VueRouter from 'vue-router'
import App from './App.vue'
import router from './router'
import ElementUI, {Message} from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import axios from 'axios' //引入axios

Vue.prototype.$axios = axios
Vue.prototype.$message = Message

Vue.config.productionTip = false
Vue.use(ElementUI)
Vue.use(Message)
Vue.use(VueRouter)


/* eslint-disable no-new */
new Vue({
  el: '#App',
  router,
  components: { App },
  template: '<App/>'
})
