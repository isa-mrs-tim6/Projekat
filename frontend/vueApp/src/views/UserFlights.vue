<template>
    <div>
        <UserNavBar></UserNavBar>
        <v-layout justify-center style="margin-top: 200px">
            <v-flex xs9 style="height: 100%;">
                <div v-bind:isLogIn="isLogIn"><flight-search></flight-search></div>
            </v-flex>
        </v-layout>
    </div>
</template>

<script>
    import ResultGrid from "../components/ResultGrid";
    import axios from "axios";
    import UserNavBar from "../components/User/UserNavBar";
    import FlightSearch from "../components/User/FlightSearch";

    export default {
        name: "UserFlights",
        components: {FlightSearch, UserNavBar, ResultGrid},
        data() {
            return {
                items: [],
                title: "Airlines in our system",
                isLogIn: false,
            }
        },
        beforeCreate() {
            axios.create({withCredentials: true}).get("http://ec2-18-195-170-20.eu-central-1.compute.amazonaws.com:8000/api/airline/getAirlines")
                .then(res =>{
                    this.items = res.data;
                })
        }
    }
</script>

<style scoped>

</style>