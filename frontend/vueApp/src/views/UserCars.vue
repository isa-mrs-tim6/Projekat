<template>
    <div>
        <UserNavBar></UserNavBar>
        <v-layout justify-center style="margin-top: 200px">
            <v-flex xs9 style="height: 100%;">
                <RACSearch v-bind:passengers="passengers" v-bind:reservationID="reservationID"></RACSearch>
            </v-flex>
        </v-layout>
    </div>
</template>

<script>
    import axios from "axios";
    import UserNavBar from "../components/User/UserNavBar";
    import VehiclesSearch from "../components/Search/VehiclesSearch";
    import RACSearch from "./User/RACSearch";

    export default {
        name: "UserCars",
        components: {RACSearch, VehiclesSearch, UserNavBar},
        data() {
            return {
                items: [],
                title: "Rent a car companies in our system",
                isLogIn: false,
                reservationID: Number.parseInt(this.$route.query.reservationID),
                passengers: Number.parseInt(this.$route.query.passengers),
            }
        },
        beforeCreate() {
            axios.create({withCredentials: true}).get("http://ec2-18-195-170-20.eu-central-1.compute.amazonaws.com:8000/api/rentACarCompany/getRentACarCompanies")
                .then(res =>{
                    this.items = res.data;
                })
        }
    }
</script>

<style scoped>

</style>