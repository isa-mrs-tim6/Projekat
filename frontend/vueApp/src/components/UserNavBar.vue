<template>
    <v-toolbar fixed app dark v-if="this.isLogIn === false">
        <v-toolbar-title>Title</v-toolbar-title>
        <v-spacer></v-spacer>
        <v-btn class="button" flat @click="goFlights">Flights</v-btn>
        <v-btn class="button" flat @click="goHotels">Hotels</v-btn>
        <v-btn class="button" flat @click="goCars">Cars</v-btn>
        <v-btn class="button" flat @click="goLogin">Login</v-btn>
        <v-btn class="button" flat @click="goRegister">Register</v-btn>
    </v-toolbar>
    <v-toolbar fixed app dark v-else>
        <v-toolbar-title>Title</v-toolbar-title>
        <v-spacer></v-spacer>
        <v-btn class="button" flat @click="goFlights">Flights</v-btn>
        <v-btn class="button" flat @click="goHotels">Hotels</v-btn>
        <v-btn class="button" flat @click="goCars">Cars</v-btn>
        <v-btn class="button" flat @click="goProfile">Profile</v-btn>
        <v-btn class="button" flat @click="goLogout">Log out</v-btn>
    </v-toolbar>
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
            },
            goProfile (e){
                e.preventDefault();
                this.$router.replace("userProfile");
            },
            goFlights (e){
                e.preventDefault();
                this.$router.replace("user");
            },
            goHotels (e){
                e.preventDefault();
                this.$router.replace("user_hotels");
            },
            goCars (e){
                e.preventDefault();
                this.$router.replace("user_cars");
            },
            goLogin (e){
                e.preventDefault();
                this.$router.replace("login");
            },
            goRegister (e){
                e.preventDefault();
                this.$router.replace("register");
            },
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