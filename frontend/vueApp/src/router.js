import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'
import SystemAdmin from "./views/SystemAdmin";
import HotelAdmin from "./views/HotelAdmin";
import User from "./views/User";
import RacAdmin from "./views/RacAdmin";
import VehiclesSearch from "./views/VehiclesSearch";
import RegistrationSuccess from "./views/RegistrationSuccess";
import AirlineAdmin from "./views/AirlineAdmin";

Vue.use(Router);

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      component: Home
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
      path: '/user',
      component: User
    },
    {
      path: '/airlineAdmin/:id',
      component: AirlineAdmin
    },
    {
      path: '/racAdmin/:id',
      component: RacAdmin
    },
    {
      path: '/vehiclesSearch/:id',
      component: VehiclesSearch
    },
    {
      path: '/confirmRegistration/:type/:email/',
      component: RegistrationSuccess
    }
  ]
})
