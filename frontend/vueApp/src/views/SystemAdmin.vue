<template>
    <div>
        <SystemAdminNavbar></SystemAdminNavbar>
        <form>
            <v-text-field
                    v-model="Name"
                    label="Name"
                    required
            ></v-text-field>
            <v-text-field
                    v-model="Address"
                    label="Address"
                    required
            ></v-text-field>
            <v-text-field
                    v-model="Latitude"
                    label="Latitude"
                    required
            ></v-text-field>
            <v-text-field
                    v-model="Longitude"
                    label="Longitude"
                    required
            ></v-text-field>
            <v-text-field
                    v-model="Description"
                    label="Description"
                    required
            ></v-text-field>

            <v-btn @click="addHotel">submit</v-btn>
            <v-btn @click="clear">clear</v-btn>
        </form>
        <Hotels v-bind:hotels="hotels"/>
    </div>
</template>

<script>
    import SystemAdminNavbar from "../components/SystemAdminNavbar";
    import Hotels from "../components/Hotels";
    import axios from 'axios';

    export default {
        name: "SystemAdmin",
        components: {Hotels, SystemAdminNavbar},
        data() {
            return {
                hotels: [],
                Name: '',
                Address: '',
                Longitude: '',
                Latitude: '',
                Description: '',
            }
        },
        created() {
            axios.get('http://localhost:8000/api/hotel/getHotels')
                .then(res => this.hotels = res.data)
                .catch(err => console.log(err));
        },
        methods: {
            addHotel(e) {
                e.preventDefault();
                const newHotel = {
                    'Name': this.Name,
                    'Address': this.Address,
                    'Latitude': parseFloat(this.Latitude),
                    'Longitude': parseFloat(this.Longitude),
                    'Description': this.Description,
                };
                console.log(newHotel);
                this.clear();
                axios.post('http://localhost:8000/api/hotel/addHotel', {
                    newHotel
                })
                    .then(res =>
                        axios.get('http://localhost:8000/api/hotel/getHotels')
                        .then(res => this.hotels = res.data)
                        .catch(err => console.log(err)))
                    .catch(err => console.log(err));
            },
            clear () {
                this.Name = '';
                this.Address = '';
                this.Longitude = '';
                this.Latitude = '';
                this.Description = '';
            }
        },
    }
</script>

<style scoped>

</style>