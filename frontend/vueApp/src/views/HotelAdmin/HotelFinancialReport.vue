<template>
    <div>
        <HotelAdminNavDrawer/>
        <HotelFinance v-bind:reservations="Reservations"></HotelFinance>
    </div>
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
</style>