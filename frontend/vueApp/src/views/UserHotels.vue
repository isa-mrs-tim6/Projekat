<template>
    <div>
        <UserNavBar></UserNavBar>
        <v-layout justify-center style="margin-top: 200px">
            <v-flex xs9 style="height: 100%;">
                <HotelSearch v-on:search="search"/>
                <HotelGrid v-bind:hotels="items" v-bind:reservationID="reservationID" v-bind:passengers="passengers" style="margin-top: 150px"></HotelGrid>
            </v-flex>
        </v-layout>
    </div>
</template>

<script>
    import axios from "axios";
    import UserNavBar from "../components/UserNavBar";
    import HotelSearch from "../components/Search/HotelSearch";
    import HotelGrid from "../components/User/HotelGrid";

    export default {
        name: "UserHotels",
        components: {HotelGrid, HotelSearch, UserNavBar},
        data() {
            return {
                items: [],
                title: "Hotels in our system",
                isLogIn: false,
                reservationID: Number.parseInt(this.$route.query.reservationID),
                passengers: Number.parseInt(this.$route.query.passengers),
            }
        },
        methods: {
            search: function(query) {
                axios.post('http://localhost:8000/api/search/hotels', query)
                    .then(res => this.items = res.data)
                    .catch(err => alert(err));
            }
        }
    }
</script>

<style scoped>

</style>