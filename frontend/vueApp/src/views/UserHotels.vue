<template>
    <div>
        <UserNavBar></UserNavBar>
        <v-layout justify-center style="margin-top: 200px">
            <v-flex xs9 style="height: 100%;">
                <HotelSearch v-on:search="search"/>
                <HotelGrid v-bind:hotels="items" v-bind:start="start" v-bind:end="end" v-bind:reservationID="reservationID" v-bind:passengers="passengers" style="margin-top: 150px"></HotelGrid>
            </v-flex>
        </v-layout>
    </div>
</template>

<script>
    import axios from "axios";
    import HotelSearch from "../components/Search/HotelSearch";
    import HotelGrid from "../components/User/HotelGrid";
    import UserNavBar from "../components/User/UserNavBar";

    export default {
        name: "UserHotels",
        components: {HotelGrid, HotelSearch, UserNavBar},
        data() {
            return {
                items: [],
                isLogIn: false,
                reservationID: Number.parseInt(this.$route.query.reservationID),
                passengers: Number.parseInt(this.$route.query.passengers),
                start: null,
                end: null,
            }
        },
        methods: {
            search: function(query) {
                this.start = query["From"];
                this.end = query["To"];
                axios.post('http://ec2-18-195-170-20.eu-central-1.compute.amazonaws.com:8000/api/search/hotels', query)
                    .then(res => {
                        this.items = res.data;
                        localStorage.setItem("hotelStart", this.start);
                        localStorage.setItem("hotelEnd", this.end)
                    })
                    .catch(err => alert(err));
            }
        }
    }
</script>

<style scoped>

</style>