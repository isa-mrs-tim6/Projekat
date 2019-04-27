<template>
    <div>
        <NavBarPreLogin></NavBarPreLogin>
        <v-layout justify-center style="margin-top: 200px">
            <v-flex xs9 style="height: 100%;">
                <div style="height: 75px; background-color: lightgray">Ovde ide komponenta za pretragu letova/hotela/vozila</div>
                <ResultGrid v-bind:items="items" v-bind:title="title" style="margin-top: 150px"></ResultGrid>
            </v-flex>
        </v-layout>
    </div>
</template>

<script>
    import ResultGrid from "../components/ResultGrid";
    import axios from "axios";
    import NavBarPreLogin from "../components/NavBarPreLogin";

    export default {
        name: "UserFlights",
        components: {NavBarPreLogin, ResultGrid},
        data() {
            return {
                items: [],
                title: "Airlines in our system",
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