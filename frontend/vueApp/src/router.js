import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'
import UserFlights from "./views/UserFlights";
import UserHotels from "./views/UserHotels";
import UserCars from "./views/UserCars";
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
import HotelFeatures from "./views/HotelAdmin/HotelFeatures"
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
import SystemAdminReservationRewards from "./views/SystemAdmin/SystemAdminReservationRewards";
import FlightSearch from "./components/User/FlightSearch";
import FlightSearchResults from "./views/User/FlightSearchResults";
import AirlineRatings from "./views/AirlineAdmin/AirlineRatingsReport";
import AirlineFinancialReport from "./views/AirlineAdmin/AirlineFinancialReport";
import RACRatingsReport from "./views/RACAdmin/RACRatingsReport";
import RACFinancialReport from "./views/RACAdmin/RACFinancialReport";
import RACSearch from "./views/User/RACSearch";
import ShowProfile from "./views/User/ShowProfile";
import ShowReservations from "./views/User/ShowReservations";
import ShowVehicleSearch from "./views/User/ShowVehicleSearch";
import FlightReservation from "./components/User/FlightReservation";
import Reserve from "./views/User/Reserve";
import ShowFriends from "./views/User/ShowFriends";
import ReservationChoice from "./views/User/ReservationChoice";
import RoomReservation from "./components/User/RoomReservation";
import HotelReservationRewards from "./views/HotelAdmin/HotelReservationRewards";
import QuickRoomReservation from "./components/User/QuickRoomReservation";
import HotelList from "./views/User/HotelList";
import ShowQuickVehicle from "./views/User/ShowQuickVehicle";
import UserSearch from "./views/User/UserSearch";
import AirlineList from "./views/User/AirlinesList";
import AirlineProfile from "./views/User/AirlineProfile";
import AirlineFeatures from "./views/AirlineAdmin/AirlineFeatures.vue";
import AddAirplane from "./views/AirlineAdmin/AddAirplane";
import VueCookies from 'vue-cookies'


Vue.use(Router);
Vue.use(VueCookies);

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      component: Home
    },
    {
      path: '/sr',
      component: FlightSearch
    },
    {
      path: '/systemAdmin',
      component: SystemAdminProfile
    },
    {
      path: '/systemAdmin/rewards',
      component: SystemAdminReservationRewards
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
      path: '/hotelAdmin/features',
      component: HotelFeatures
    },
    {
      path: '/hotelAdmin/rewards',
      component: HotelReservationRewards
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
      component: UserFlights
    },
    {
      path: '/user/reserve',
      component: ReservationChoice
    },
    {
      path: '/reserve_flight',
      component: Reserve
    },
    {
      path: '/user_airlines',
      component: AirlineList
    },
    {
      path: '/airline/:id',
      component: AirlineProfile
    },
    {
      path: '/user_flights',
      component: Reserve
    },
    {
      path: "/user_hotels",
      component: UserHotels
    },
    {
      path: '/user/hotels',
      component: HotelList
    },
    {
      path: "/user_cars",
      component: UserCars
    },
    {
      path: "/user_search",
      component: UserSearch
    },
    {
      path: "/userProfile",
      component: ShowProfile
    },
    {
      path: "/reserve_room/:hotel_id/:reservation/:tickets",
      component: RoomReservation
    },
    {
      path: "/quick_reserve_room/:hotel_id/:reservation/:tickets",
      component: QuickRoomReservation
    },
    {
      path: "/userReservations",
      component: ShowReservations
    },
    {
      path: "/userFriends",
      component: ShowFriends
    },
    {
      path: '/airlineAdmin/add_airplane',
      component: AddAirplane
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
      path: '/airlineAdmin/ratings',
      component: AirlineRatings
    },
    {
      path: '/airlineAdmin/features',
      component: AirlineFeatures
    },
    {
      path: '/airlineAdmin/finances',
      component: AirlineFinancialReport
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
      path: '/racAdmin/ratings',
      component: RACRatingsReport
    },
    {
      path: '/racAdmin/finances',
      component: RACFinancialReport
    },
    {
      path: '/racSearch',
      component: RACSearch
    },
    {
      path: '/vehiclesSearch/:id/:locID',
      component: ShowVehicleSearch
    },
    {
      path: '/vehiclesQuick/:id/:locID',
      component: ShowQuickVehicle
    },
    {
      path: '/confirmRegistration/:type/:email/',
      component: RegistrationSuccess
    },
    {
      path: '/flightSearch',
      component: FlightSearchResults,
      name: 'fSearch'
    },
    {
      path: '/flightReservation',
      component: FlightReservation,
      name: 'fReservation'
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
