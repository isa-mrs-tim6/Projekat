<template>
    <div v-if="this.isLogIn === false">
        <v-toolbar fixed app dark >
            <v-toolbar-title>Front page</v-toolbar-title>
            <v-spacer></v-spacer>
            <v-btn class="button" flat @click="goFlights">Flights</v-btn>
            <v-btn class="button" flat @click="goHotels">Hotels</v-btn>
            <v-btn class="button" flat @click="goCars">Cars</v-btn>
            <v-btn class="button" flat @click="goLogin">Login</v-btn>
            <v-btn class="button" flat @click="goRegister">Register</v-btn>
        </v-toolbar>
    </div>
    <div v-else>
        <v-toolbar  fixed app dark >
            <v-toolbar-title>Front page</v-toolbar-title>
            <v-spacer></v-spacer>
            <v-btn class="button" flat @click="goReserve">Reserve</v-btn>
            <v-btn class="button" flat @click="goFlights">Flights</v-btn>
            <v-btn class="button" flat @click="goHotels">Hotels</v-btn>
            <v-btn class="button" flat @click="goCars">Cars</v-btn>
            <v-btn class="button" flat @click="goUsers">Find a user</v-btn>
            <v-btn class="button" flat @click="goProfile">Profile</v-btn>
            <v-btn class="button" flat @click="goLogout">Logout</v-btn>
        </v-toolbar>
    </div>
</template>

<script>
    import axios from "axios";

    export default {
        name: "UserNavBar",
        beforeCreate(){
            axios.create({withCredentials: true}).get("http://localhost:8000/api/user/getProfile")
                .then(res =>{
                    this.isLogIn = true;
                })
                .catch(err =>{
                    this.isLogIn = false;
                })
        },
        data(){
            return {
                isLogIn: false,
            }
        },
        methods: {
            goLogout (e){
                e.preventDefault();
                this.isLogIn = false;
                this.$cookies.remove("token");
                this.$router.push("/user")
            },
            goProfile (e){
                e.preventDefault();
                this.$router.push("/userProfile");
            },
            goFlights (e){
                e.preventDefault();
                this.$router.push("/user_flights");
            },
            goHotels (e){
                e.preventDefault();
                this.$router.push("/user/hotels");
            },
            goCars (e){
                e.preventDefault();
                this.$router.push("/user_cars");
            },
            goUsers(e){
                e.preventDefault();
                this.$router.push("/user_search");
            },
            goLogin (e){
                e.preventDefault();
                this.$router.push("/login");
            },
            goRegister (e){
                e.preventDefault();
                this.$router.push("/register");
            },
            goReserve (e){
                e.preventDefault();
                this.$router.push("/user/reserve");
            }
        }
    }
</script>

<style scoped>
    .button{
        padding: 0;
        margin: 0;
        height:100%;
    }
</style>