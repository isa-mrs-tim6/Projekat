import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'
import User from "./views/User";
import VehiclesSearch from "./views/VehiclesSearch";
import RegistrationSuccess from "./views/RegistrationSuccess";
import AirlineAdmin from "./views/AirlineAdmin/AirlineAdmin";
import RegisterUser from "./views/RegisterUser";
import LoginUser from "./views/LoginUser";
import Destination from "./views/AirlineAdmin/Destination";
import UpdateDestination from "./views/AirlineAdmin/UpdateDestination";
import LoginAdmin from "./views/LoginAdmin";
import AddFlight from "./views/AirlineAdmin/AddFlight";
import EditFlight from "./views/AirlineAdmin/EditFlight";
import HotelFinancialReport from "./views/HotelAdmin/HotelFinancialReport";
import HotelAdminProfile from "./views/HotelAdmin/HotelAdminProfile";
import HotelRooms from "./views/HotelAdmin/HotelRooms";
import HotelProfile from "./views/HotelAdmin/HotelProfile";
import HotelRatingsReport from "./views/HotelAdmin/HotelRatingsReport";
import SystemAdminAccounts from "./views/SystemAdmin/SystemAdminAccounts";
import SystemAdminAirlines from "./views/SystemAdmin/SystemAdminAirlines";
import SystemAdminHotels from "./views/SystemAdmin/SystemAdminHotels";
import SystemAdminRentACars from "./views/SystemAdmin/SystemAdminRentACars";
import SystemAdminProfile from "./views/SystemAdmin/SystemAdminProfile";
import RACAdminProfile from "./views/RACAdmin/RACAdminProfile";
import RACProfile from "./views/RACAdmin/RACProfile";
import RACOffices from "./views/RACAdmin/RACOffices";
import RACVehicles from "./views/RACAdmin/RACVehicles";
import AirlineAdminProfile from "./views/AirlineAdmin/AirlineAdminProfile";

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
      component: SystemAdminProfile
    },
    {
      path: '/systemAdmin/accounts',
      component: SystemAdminAccounts
    },
    {
      path: '/systemAdmin/airlines',
      component: SystemAdminAirlines
    },
    {
      path: '/systemAdmin/hotels',
      component: SystemAdminHotels
    },
    {
      path: '/systemAdmin/rent_a_cars',
      component: SystemAdminRentACars
    },
    {
      path: '/systemAdmin/admin_profile',
      component: SystemAdminProfile
    },
    {
      path: '/hotelAdmin/ratings',
      component: HotelRatingsReport
    },
    {
      path: '/hotelAdmin/finances',
      component: HotelFinancialReport
    },
    {
      path: '/hotelAdmin/rooms',
      component: HotelRooms
    },
    {
      path: '/hotelAdmin/hotel_profile',
      component: HotelProfile
    },
    {
      path: '/hotelAdmin/admin_profile',
      component: HotelAdminProfile
    },
    {
      path: '/user',
      component: User
    },
    {
      path: '/airlineAdmin/airline_profile',
      component: AirlineAdmin
    },
    {
      path:'/airlineAdmin/flight_edit',
      component: EditFlight
    },
    {
      path: '/airlineAdmin/flight_add',
      component: AddFlight
    },
    {
      path: '/airlineAdmin/destination_add',
      component: Destination
    },
    {
      path: '/airlineAdmin/destination/:id',
      component: UpdateDestination
    },
    {
      path: '/airlineAdmin/admin_profile',
      component: AirlineAdminProfile
    },
    {
      path: '/racAdmin/admin_profile',
      component: RACAdminProfile
    },
    {
      path: '/racAdmin/rac_profile',
      component: RACProfile
    },
    {
      path: '/racAdmin/offices',
      component: RACOffices
    },
    {
      path: '/racAdmin/vehicles',
      component: RACVehicles
    },
    {
      path: '/vehiclesSearch/:id',
      component: VehiclesSearch
    },
    {
      path: '/confirmRegistration/:type/:email/',
      component: RegistrationSuccess
    },
    {
      path: '/register',
      component: RegisterUser
    },
    {
      path: '/login',
      component: LoginUser
    },
    {
      path: '/admin_login',
      component: LoginAdmin
    }
  ]
})
