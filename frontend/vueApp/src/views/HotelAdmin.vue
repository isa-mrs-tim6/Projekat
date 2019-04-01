<template>
    <div>
        <HotelAdminNavbar></HotelAdminNavbar>
        <form>
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
            axios.get('http://localhost:8000/api/hotel/'+this.$route.params.id+'/getProfile')
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
                if (isNaN(Number(this.AirlineProfile.Longitude)) && isNaN(Number(this.AirlineProfile.Latitude))){
                    alert("Enter a number for Longitude and Latitude");
                }else if(isNaN(Number(this.AirlineProfile.Longitude))){
                    alert("Enter a number for Longitude");
                }else if(isNaN(Number(this.AirlineProfile.Latitude))){
                    alert("Enter a number for Latitude");
                }else {
                    this.AirlineProfile.Latitude = Number(this.AirlineProfile.Latitude);
                    this.AirlineProfile.Longitude = Number(this.AirlineProfile.Longitude);
                    axios.post('http://localhost:8000/api/hotel/' + this.$route.params.id + '/updateProfile', this.hotelProfile)
                        .then(res => this.backupHotelProfile = JSON.parse(JSON.stringify(this.hotelProfile)))
                        .catch(err => console.log(err));
                }
            },
            revert () {
                this.hotelProfile = JSON.parse(JSON.stringify(this.backupHotelProfile))
            }
        },
    }
</script>

<style scoped>

</style>