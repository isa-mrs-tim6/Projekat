<template>
    <div>
        <UserNavBar></UserNavBar>
        <v-layout justify-center style="margin-top: 200px">
            <v-flex xs9 style="height: 100%;">
                <Hotels v-bind:hotels="hotels"/>
            </v-flex>
        </v-layout>
    </div>
</template>

<script>
    import axios from 'axios';
    import UserNavBar from "../../components/User/UserNavBar";
    import Hotels from  "../../components/Hotels";

    export default {
        name: "HotelList",
        components: {UserNavBar, Hotels},
        data() {
            return {
                hotels: [],
            }
        },
        created() {
            axios.create({withCredentials: true}).get('http://ec2-35-159-21-254.eu-central-1.compute.amazonaws.com:8000/api/hotel/getHotels')
                .then(res => this.hotels = res.data)
                .catch(err => console.log(err));
        },
        methods: {

        }
    }
</script>

<style scoped>

</style>