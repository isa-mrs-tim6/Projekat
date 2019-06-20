<template>
    <v-container fill-height fluid id="main">
        <v-layout>
            <v-flex>
                <r-a-c-admin-nav-drawer/>
            </v-flex>
            <v-container>
                <v-layout>
                    <v-flex align-self-center>
                        <rac-ratings style="height: 100vh;" v-bind:vehicles="Vehicles" v-bind:name="Name" v-bind:rating="Rating"></rac-ratings>
                    </v-flex>
                </v-layout>
            </v-container>
        </v-layout>
    </v-container>
</template>

<script>
    import RACAdminNavDrawer from "../../components/RACAdmin/RACAdminNavDrawer";
    import axios from "axios";
    import RacRatings from "../../components/RACAdmin/RACRatings";

    export default {
        name: "RACRatingsReport",
        components: {RacRatings, RACAdminNavDrawer},
        data: () => ({
            Vehicles: [],
            Reservations: [],
            Name: "",
            Rating: 0,
        }),
        created() {
            axios.create({withCredentials: true}).get('http://ec2-18-195-170-20.eu-central-1.compute.amazonaws.com:8000/api/rentACarCompany/getProfile')
                .then(res => {
                        this.Name = res.data.Name;
                    }
                )
                .catch(err => alert(err));
            axios.create({withCredentials: true}).get('http://ec2-18-195-170-20.eu-central-1.compute.amazonaws.com:8000/api/rentACarCompany/getVehicleRatings')
                .then(res => {
                    this.Vehicles = res.data;
                } )
                .catch(err => alert("Could not retrieve vehicles and their ratings"));
            axios.create({withCredentials: true}).get('http://ec2-18-195-170-20.eu-central-1.compute.amazonaws.com:8000/api/rentACarCompany/getReservations')
                .then(res => {
                    this.Reservations = res.data;
                    let len = 0;
                    for (let i = 0; i < this.Reservations.length; i++) {
                        if (this.Reservations[i].CompanyRating !== 0){
                            this.Rating += this.Reservations[i].CompanyRating;
                            len += 1;
                        }
                    }
                    if(len === 0){
                        len = 1;
                    }
                    if (this.Reservations.length > 0) {
                        this.Rating /= len;
                    }
                } )
                .catch(err => alert("Could not retrieve company reservations"));
        },
        mounted(){
            this.checkFirstPass();
        },
        methods: {
            checkFirstPass(){
                axios.create({withCredentials: true}).get("http://ec2-18-195-170-20.eu-central-1.compute.amazonaws.com:8000/api/admin/checkFirstPass")
                    .then(
                        res =>{
                            if(res.data === false){
                                alert("Please update your password before accessing this feature");
                                this.$router.replace("admin_profile");
                            }
                        }
                    )
            }
        },
    }
</script>

<style scoped>
    #main {
        background-image: linear-gradient(to right bottom, #560579, #561a71, #56266a, #543062, #52395b, #594460, #604f64, #665a69, #796c7c, #8c7f8f, #9f93a3, #b3a7b7);
    }
</style>