<template>
    <div>
        <UserNavBar></UserNavBar>
        <v-layout justify-center style="margin-top: 150px">
            <v-flex xs9 style="height: 100%;">
                <RACSearch></RACSearch>
            </v-flex>
        </v-layout>
    </div>
</template>

<script>
    import axios from "axios";
    import UserNavBar from "../components/UserNavBar";
    import VehiclesSearch from "./VehiclesSearch";
    import RACSearch from "./User/RACSearch";

    export default {
        name: "UserCars",
        components: {RACSearch, VehiclesSearch, UserNavBar},
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