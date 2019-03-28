import Vue from 'vue'
import Router from 'vue-router'
//import Home from './views/Home.vue'
import SystemAdmin from "./views/SystemAdmin";
import HotelAdmin from "./views/HotelAdmin";
import RacAdmin from "./views/RacAdmin";
import VehiclesSearch from "./views/VehiclesSearch";

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
      path: '/hotelAdmin/:id',
      component: HotelAdmin
    },
    {
      path: '/racAdmin/:id',
      component: RacAdmin
    },
    {
      path: '/vehiclesSearch/:id',
      component: VehiclesSearch
    }
  ]
})
