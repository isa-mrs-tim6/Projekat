<template>
    <v-container fill-height fluid id="main">
        <v-layout>
            <v-flex>
                <HotelAdminNavDrawer/>
            </v-flex>
            <v-container>
                <v-layout>
                    <v-flex align-self-center>
                        <HotelRatings v-bind:rooms="Rooms" v-bind:name="Name" v-bind:rating="Rating"></HotelRatings>
                    </v-flex>
                </v-layout>
            </v-container>
        </v-layout>
    </v-container>
</template>

<script>
    import HotelAdminNavDrawer from "../../components/HotelAdmin/HotelAdminNavDrawer";
    import HotelRatings from "../../components/HotelAdmin/HotelRatings";
    import axios from 'axios';

    export default {
        name: "HotelRatingsReport",
        components: {HotelAdminNavDrawer, HotelRatings},
        data: () => ({
            Rooms: [],
            Reservations: [],
            Name: "",
            Rating: 0,
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

<style scoped>
    #main {
        background-image: linear-gradient(to left top, #ac0c0c, #b41812, #bb2218, #c32a1d, #cb3223, #d23928, #d8402c, #df4631, #e74d37, #ee543c, #f65a42, #fe6148);
    }
</style>