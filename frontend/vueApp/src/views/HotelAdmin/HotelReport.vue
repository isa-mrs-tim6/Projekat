<template>
    <v-container fluid>
        <HotelRatings v-bind:rooms="Rooms" v-bind:name="Name" v-bind:rating="Rating"></HotelRatings>
        <HotelFinance v-bind:reservations="Reservations"></HotelFinance>
    </v-container>
</template>

<script>
    import HotelRatings from "../../components/HotelAdmin/HotelRatings";
    import HotelFinance from "../../components/HotelAdmin/HotelFinance";
    import axios from 'axios';

    export default {
        name: "HotelReport",
        components: {HotelFinance, HotelRatings},
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
        }
    }
</script>

<style scoped>

</style>