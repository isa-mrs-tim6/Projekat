<template>
    <div>
        <AirlineAdminNavbar></AirlineAdminNavbar>
        <v-data-table :headers="headers" :items="destinations" class="elevation-1">
            <template v-slot:items="destinations">
                <td>{{destinations.item.Name}}</td>
                <td>{{destinations.item.Longitude}}</td>
                <td>{{destinations.item.Latitude}}</td>
            </template>
        </v-data-table>
        <v-form>
            <v-container>
                <v-text-field :rules="[() => !!destination.Name || 'This field is required']" label="Destination Name"
                              v-model="destination.Name"></v-text-field>
                <v-text-field :rules="[() => !!destination.Longitude || 'This field is required']" label="Longitude"
                              v-model="destination.Longitude"></v-text-field>
                <v-text-field :rules="[() => !!destination.Latitude || 'This field is required']" label="Latitude"
                              v-model="destination.Latitude"></v-text-field>
                <v-btn @click="addDestination">Submit</v-btn>
            </v-container>
        </v-form>
    </div>
</template>

<script>
    import axios from 'axios';
    import AirlineAdminNavbar from "../../components/AirlineAdmin/AirlineAdminNavbar";
    export default {
        name: "Destination",
        components: {AirlineAdminNavbar},
        data() {
            return {
                headers: [
                    { text: 'Name', value: 'Name' },
                    { text: 'Longitude', value: 'Longitude' },
                    { text: 'Latitude', value: 'Latitude' },
                ],
                destinations: [],
                destination:{
                    Name:"",
                    Longitude: 0,
                    Latitude: 0
                }
            }
        },
        mounted() {
            axios.get('http://localhost:8000/api/destination/getDestinations')
                .then(res => {
                        this.destinations = res.data
                    }
                )
                .catch(err => console.log(err));
        },
        methods: {
            addDestination(e) {
                e.preventDefault();
                if (!this.destination.Name || !this.destination.Longitude || !this.destination.Latitude){
                    return;
                }
                if (isNaN(Number(this.destination.Longitude)) && isNaN(Number(this.destination.Latitude))) {
                    alert("Enter a number for Longitude and Latitude");
                } else if (isNaN(Number(this.destination.Longitude))) {
                    alert("Enter a number for Longitude");
                } else if (isNaN(Number(this.destination.Latitude))) {
                    alert("Enter a number for Latitude");
                } else {
                    this.destination.Latitude = Number(this.destination.Latitude);
                    this.destination.Longitude = Number(this.destination.Longitude);
                    axios.post('http://localhost:8000/api/destination/add', this.destination)
                        .then(res =>
                            axios.get('http://localhost:8000/api/destination/getDestinations')
                                .then(res => this.destinations = res.data)
                                .catch(err => console.log(err)))
                        .catch(err => console.log(err));
                }
            }
        }
    }
</script>

<style scoped>

</style>