<template>
    <div>
        <v-toolbar>
            <v-toolbar-side-icon></v-toolbar-side-icon>
            <v-toolbar-title>Hotel Admin Panel</v-toolbar-title>
            <v-spacer></v-spacer>
            <v-toolbar-items class="hidden-sm-and-down">
                <v-btn @click="hotelReport" flat>Hotel Report</v-btn>
                <v-btn @click="manageRooms" flat>Rooms</v-btn>
                <v-btn @click="manageHotelProfile" flat>Hotel Profile</v-btn>
                <v-btn @click="manageProfile" flat>Admin Profile</v-btn>
            </v-toolbar-items>
        </v-toolbar>
        <component v-bind:is="currentTabComponent"></component>
    </div>
</template>

<script>
    import ManageRooms from "../components/ManageRooms";
    import ManageHotelProfile from "../components/ManageHotelProfile";
    import AdminProfile from "../components/AdminProfile";
    import HotelReport from "./HotelAdmin/HotelReport";
    import axios from "axios";

    export default {
        name: "HotelAdmin",
        components: {ManageRooms, ManageHotelProfile, AdminProfile, HotelReport},
        data() {
            return {
                currentTabComponent: "ManageRooms",
            }
        },
        mounted(){
            this.checkFirstPass()
        },
        methods: {
            manageRooms(){
                this.currentTabComponent = "ManageRooms";
                this.checkFirstPass()
            },
            manageHotelProfile(){
                this.currentTabComponent = "ManageHotelProfile";
                this.checkFirstPass()
            },
            manageProfile(){
                this.currentTabComponent = "AdminProfile";
            },
            hotelReport(){
                this.currentTabComponent = "HotelReport";
                this.checkFirstPass()
            },
            checkFirstPass(){
                axios.create({withCredentials: true}).get("http://localhost:8000/api/admin/checkFirstPass")
                    .then(
                        res =>{
                            if(res.data === false){
                                alert("Please update your password");
                                this.manageProfile();
                            }
                        }
                    )
            }
        },
    }
</script>

<style scoped>

</style>