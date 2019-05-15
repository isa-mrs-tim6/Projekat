<template>
    <div id="main">
        <r-a-c-admin-nav-drawer/>
        <rac-ratings style="height: 100vh;" v-bind:vehicles="Vehicles" v-bind:name="Name" v-bind:rating="Rating"></rac-ratings>
    </div>
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
            axios.create({withCredentials: true}).get('http://localhost:8000/api/rentACarCompany/getProfile')
                .then(res => {
                        this.Name = res.data.Name;
                    }
                )
                .catch(err => alert(err));
            axios.create({withCredentials: true}).get('http://localhost:8000/api/rentACarCompany/getVehicleRatings')
                .then(res => {
                    this.Vehicles = res.data;
                } )
                .catch(err => alert("Could not retrieve vehicles and their ratings"));
            axios.create({withCredentials: true}).get('http://localhost:8000/api/rentACarCompany/getReservations')
                .then(res => {
                    this.Reservations = res.data;
                    for (let i = 0; i < this.Reservations.length; i++) {
                        this.Rating += this.Reservations[i].CompanyRating;
                    }
                    if (this.Reservations.length > 0) {
                        this.Rating /= this.Reservations.length;
                    }
                } )
                .catch(err => alert("Could not retrieve company reservations"));
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
            }
        },
    }
</script>

<style scoped>
    #main {
        background-image: linear-gradient(to right bottom, #02bb04, #00b824, #00b535, #00b141, #00ae4c, #00ab55, #00a75e, #00a465, #00a06d, #009b73, #009778, #00927c);
    }
</style>