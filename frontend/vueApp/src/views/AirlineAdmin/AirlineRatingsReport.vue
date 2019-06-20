<template>
    <v-container fill-height fluid id="main">
        <v-layout>
            <v-flex>
                <AirlineAdminNavDrawer/>
            </v-flex>
            <v-container>
                <v-layout>
                    <v-flex align-self-center>
                        <AirlineRatings v-bind:flights="Flights" v-bind:name="Name" v-bind:rating="Rating"></AirlineRatings>
                    </v-flex>
                </v-layout>
            </v-container>
        </v-layout>
    </v-container>
</template>

<script>
    import AirlineAdminNavDrawer from "../../components/AirlineAdmin/AirlineAdminNavDrawer";
    import axios from "axios"
    import AirlineRatings from "../../components/AirlineAdmin/AirlineRatings";

    export default {
        name: "AirlineRatingsReport.vue",
        components: {AirlineRatings, AirlineAdminNavDrawer},
        data: () => ({
            Flights: [],
            Reservations: [],
            Name: "",
            Rating: 0,
        }),
        created(){
            axios.create({withCredentials: true}).get('http://localhost:8000/api/airline/getProfile')
                .then(res => {
                        this.Name = res.data.Name;
                    }
                )
                .catch(err => alert(err));
            axios.create({withCredentials: true}).get('http://localhost:8000/api/airline/getFlightRatings')
                .then(res => {
                    this.Flights = res.data;
                } )
                .catch(err => alert("Could not retrieve flights and their ratings"));
            axios.create({withCredentials: true}).get('http://localhost:8000/api/airline/getAirlineReservations')
                .then(res => {
                    this.Reservations = res.data;
                    var k = 0;
                    for (let i = 0; i < this.Reservations.length; i++) {
                        if (this.Reservations[i].CompanyRating != 0){
                            this.Rating += this.Reservations[i].CompanyRating;
                            k++;
                        }
                    }
                    if (this.Reservations.length > 0) {
                        this.Rating /= k;
                    }
                } )
                .catch(err => alert("Could not retrieve airline reservations"));
        },
        mounted(){
            this.checkFirstPass();
        },
        methods: {
            checkFirstPass(){
                axios.create({withCredentials: true}).get("http://localhost:8000/api/admin/checkFirstPass")
                    .then(
                        res =>{
                            if(res.data === false){
                                alert("Please update your password before accessing this feature");
                                this.$router.replace("admin_profile");
                            }
                        }
                    )
            },
        },
    }
</script>

<style scoped>
    #main {
        background-image: linear-gradient(to right bottom, #142eae, #005bca, #007ed2, #009ccd, #0bb7c7, #47c0c6, #67c8c6, #81d0c7, #6ecac4, #58c4c3, #3cbdc2, #00b7c1);
    }
</style>