<template>
    <div>
        <UserNavBar></UserNavBar>
        <v-layout justify-center style="margin-top: 200px">
            <v-flex xs9 style="height: 100%;">
                <HotelSearch v-on:search="search"/>
                <ResultGrid v-bind:items="items" v-bind:title="title" style="margin-top: 150px"></ResultGrid>
            </v-flex>
        </v-layout>
    </div>
</template>

<script>
    import ResultGrid from "../components/ResultGrid";
    import axios from "axios";
    import HotelSearch from "../components/Search/HotelSearch";
    import UserNavBar from "../components/User/UserNavBar";

    export default {
        name: "UserHotels",
        components: {HotelSearch, UserNavBar, ResultGrid},
        data() {
            return {
                items: [],
                title: "Hotels in our system",
                isLogIn: false,
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