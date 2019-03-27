<template>
    <div>
        <HotelAdminNavbar></HotelAdminNavbar>
        <form>
            <h1> Hotel <span id="routeID">{{ $route.params.id }}</span> </h1>
            <v-text-field
                    v-model="hotelProfile.Name"
                    label="Name"
                    required
            ></v-text-field>
            <v-text-field
                    v-model="hotelProfile.Address"
                    label="Address"
                    required
            ></v-text-field>
            <v-text-field
                    v-model="hotelProfile.Latitude"
                    label="Latitude"
                    required
            ></v-text-field>
            <v-text-field
                    v-model="hotelProfile.Longitude"
                    label="Longitude"
                    required
            ></v-text-field>
            <v-text-field
                    v-model="hotelProfile.Description"
                    label="Description"
                    required
            ></v-text-field>

            <v-btn @click="update">update</v-btn>
            <v-btn @click="revert">revert</v-btn>
        </form>
    </div>
</template>

<script>
    import HotelAdminNavbar from "../components/HotelAdminNavbar";
    import axios from 'axios';

    export default {
        name: "SystemAdmin",
        components: {HotelAdminNavbar},
        data() {
            return {
                backupHotelProfile: '',
                hotelProfile: '',
            }
        },
        mounted() {
            var hotelID = document.getElementById("routeID").innerHTML;
            axios.get('http://localhost:8000/api/hotel/'+hotelID+'/getProfile')
                .then(res => {
                    this.hotelProfile = res.data;
                    this.backupHotelProfile = JSON.parse(JSON.stringify(this.hotelProfile))
                    }
                )
                .catch(err => console.log(err));
        },
        methods: {
            update(e) {
                e.preventDefault();
                var hotelID = document.getElementById("routeID").innerHTML;
                axios.post('http://localhost:8000/api/hotel/'+hotelID+'/updateProfile', this.hotelProfile)
                    .then(res => this.backupHotelProfile = JSON.parse(JSON.stringify(this.hotelProfile)))
                    .catch(err => console.log(err));
            },
            revert () {
                this.hotelProfile = JSON.parse(JSON.stringify(this.backupHotelProfile))
            }
        },
    }
</script>

<style scoped>

</style>