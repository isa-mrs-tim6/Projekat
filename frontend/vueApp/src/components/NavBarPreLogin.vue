<template>
    <v-toolbar fixed app dark>
        <v-toolbar-title>Title</v-toolbar-title>

        <v-spacer></v-spacer>

        <v-btn class="button" flat @click="goFlights">Flights</v-btn>
        <v-btn class="button" flat @click="goHotels">Hotels</v-btn>
        <v-btn class="button" flat @click="goCars">Cars</v-btn>
        <v-btn class="button" flat @click="goLogin">Login</v-btn>
        <v-btn class="button" flat @click="goRegister">Register</v-btn>

    </v-toolbar>
</template>

<script>
    export default {
        name: "NavBarPreLogin",
        methods: {
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