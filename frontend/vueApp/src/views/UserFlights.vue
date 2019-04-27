<template>
    <div>
        <UserNavBar></UserNavBar>
        <v-layout justify-center style="margin-top: 200px">
            <v-flex xs9 style="height: 100%;">
                <div v-bind:isLogIn="isLogIn" style="height: 75px; background-color: lightgray">Ovde ide komponenta za pretragu letova/hotela/vozila</div>
                <ResultGrid v-bind:items="items" v-bind:title="title" style="margin-top: 150px"></ResultGrid>
            </v-flex>
        </v-layout>
    </div>
</template>

<script>
    import ResultGrid from "../components/ResultGrid";
    import axios from "axios";
    import UserNavBar from "../components/UserNavBar";

    export default {
        name: "UserFlights",
        components: {UserNavBar, ResultGrid},
        data() {
            return {
                items: [],
                title: "Airlines in our system",
                isLogIn: false,
            }
        },
        beforeCreate() {
            axios.create({withCredentials: true}).get("http://localhost:8000/api/airline/getAirlines")
                .then(res =>{
                    this.items = res.data;
                })
        }
    }
</script>

<style scoped>

</style>