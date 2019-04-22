<template>
    <v-container id="main" style="height: 100vh; max-height: 100%;" fluid>
        <v-layout style="height: 100vh; max-height: 100%;" row>
            <v-flex
                    shrink
                    pa-1
            >
                <HotelAdminNavDrawer height="100%"/>
            </v-flex>
            <v-flex
                    grow
                    pa-1
            >
                <HotelFinance v-bind:reservations="Reservations"></HotelFinance>
            </v-flex>
        </v-layout>
    </v-container>
</template>

<script>
    import HotelAdminNavDrawer from "../../components/HotelAdmin/HotelAdminNavDrawer";
    import HotelFinance from "../../components/HotelAdmin/HotelFinance";
    import axios from 'axios';

    export default {
        name: "HotelReport",
        components: {HotelAdminNavDrawer, HotelFinance},
        data: () => ({
            Reservations: [],
        }),
        created() {
            axios.create({withCredentials: true}).get('http://localhost:8000/api/hotel/getProfile')
                .then(res => {
                        this.Name = res.data.Name;
                    }
                )
                .catch(err => alert(err));
            axios.create({withCredentials: true}).get('http://localhost:8000/api/hotel/getRoomRatings')
                .then(res => {
                    this.Rooms = res.data;
                } )
                .catch(err => alert("Could not retrieve hotel rooms and their ratings"));
            axios.create({withCredentials: true}).get('http://localhost:8000/api/hotel/getHotelReservations')
                .then(res => {
                    this.Reservations = res.data;
                    for (let i = 0; i < this.Reservations.length; i++) {
                        this.Rating += this.Reservations[i].HotelRating;
                    }
                    if (this.Reservations.length > 0) {
                        this.Rating /= this.Reservations.length;
                    }
                } )
                .catch(err => alert("Could not retrieve hotel reservations"));
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

<style>
    #main {
        background-image: linear-gradient(to right top, #ffffff, #fafafb, #f5f5f6, #eff0f2, #eaebee, #e7eaee, #e4e9ee, #e1e8ee, #dfebf1, #deeef3, #ddf1f3, #ddf4f2);
    }
</style>