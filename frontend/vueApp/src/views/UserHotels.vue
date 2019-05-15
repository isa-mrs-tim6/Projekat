<template>
    <div>
        <UserNavBar></UserNavBar>
        <v-layout justify-center style="margin-top: 200px">
            <v-flex xs9 style="height: 100%;">
                <HotelSearch/>
                <ResultGrid v-bind:items="items" v-bind:title="title" style="margin-top: 150px"></ResultGrid>
            </v-flex>
        </v-layout>
    </div>
</template>

<script>
    import ResultGrid from "../components/ResultGrid";
    import axios from "axios";
    import UserNavBar from "../components/UserNavBar";
    import HotelSearch from "../components/Search/HotelSearch";

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
        beforeCreate() {
            axios.create({withCredentials: true}).get("http://localhost:8000/api/hotel/getHotels")
                .then(res =>{
                    this.items = res.data;
                })
        }
    }
</script>

<style scoped>

</style>