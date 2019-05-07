<template>
    <div>
        <UserNavBar></UserNavBar>
        <v-layout justify-center style="margin-top: 150px">
            <v-flex xs9 style="height: 100%;">
                <VehiclesSearch v-bind:isLogIn="isLogIn"></VehiclesSearch>
                <ResultGrid v-bind:items="items" v-bind:title="title"></ResultGrid>
            </v-flex>
        </v-layout>
    </div>
</template>

<script>
    import ResultGrid from "../components/ResultGrid";
    import axios from "axios";
    import UserNavBar from "../components/UserNavBar";
    import VehiclesSearch from "./VehiclesSearch";

    export default {
        name: "UserCars",
        components: {VehiclesSearch, UserNavBar, ResultGrid},
        data() {
            return {
                items: [],
                title: "Rent a car companies in our system",
                isLogIn: false,
            }
        },
        beforeCreate() {
            axios.create({withCredentials: true}).get("http://localhost:8000/api/rentACarCompany/getRentACarCompanies")
                .then(res =>{
                    this.items = res.data;
                })
        }
    }
</script>

<style scoped>

</style>