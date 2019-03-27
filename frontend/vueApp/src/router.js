import Vue from 'vue'
import Router from 'vue-router'
//import Home from './views/Home.vue'
import SystemAdmin from "./views/SystemAdmin";
import HotelAdmin from "./views/HotelAdmin";
import User from "./views/User";

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      component: SystemAdmin
    },
    {
      path: '/systemAdmin',
      component: SystemAdmin
    },
    {
      path: '/hotelAdmin',
      component: HotelAdmin
    },
    {
      path: '/user',
      component: User
    }
  ]
})
