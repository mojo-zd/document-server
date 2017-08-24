import Vue from 'vue'
import Router from 'vue-router'
import ElementUI from 'element-ui'
import VueResource from 'vue-resource'
import 'element-ui/lib/theme-default/index.css'
import 'bootstrap/dist/css/bootstrap.css'
import '../assets/main.css'

Vue.use(Router)
Vue.use(ElementUI)
Vue.use(VueResource)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'dashboard',
      component: require('../components/documents/Main')
    },
    {
      path: '/document',
      name: 'documents',
      component: require('../components/documents/Document')
    }
  ]
})
