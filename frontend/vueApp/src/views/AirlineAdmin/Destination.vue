<template>
    <div id="main">
        <airline-admin-nav-drawer></airline-admin-nav-drawer>
        <v-container grid-list-xl text-xs-center style="height: 100vh;">
            <v-layout align-center justify-center column wrap fill-height>
                <v-flex style="width: 60vw">
                    <v-card  min-height="100%" class="flexcard">
                        <v-card-title primary-title>
                            <div class="headline font-weight-medium">Add new destination</div>
                        </v-card-title>
                        <v-card-text class="grow">
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
                                </v-container>
                            </v-form>
                        </v-card-text>
                        <v-card-actions>
                            <v-btn color="primary" @click="addDestination">Submit</v-btn>
                        </v-card-actions>
                    </v-card>
                </v-flex>
            </v-layout>
        </v-container>
        <v-snackbar v-model="SuccessSnackbar" :timeout=4000 :top="true" color="success">Successfully added new destination</v-snackbar>
        <v-snackbar v-model="ErrorSnackbar" :timeout=4000 :top="true" color="error">Failed to add new destination</v-snackbar>
    </div>
</template>

<script>
    import axios from 'axios';
    import AirlineAdminNavbar from "../../components/AirlineAdmin/AirlineAdminNavbar";
    import AirlineAdminNavDrawer from "../../components/AirlineAdmin/AirlineAdminNavDrawer";
    export default {
        name: "Destination",
        components: {AirlineAdminNavDrawer, AirlineAdminNavbar},
        data() {
            return {
                SuccessSnackbar: false,
                ErrorSnackbar: false,
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
            axios.create({withCredentials:true}).get('http://localhost:8000/api/destination/getCompanyDestinations')
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
                    axios.create({withCredentials: true}).post('http://localhost:8000/api/destination/add', this.destination)
                        .then(res =>
                            axios.create({withCredentials:true}).get('http://localhost:8000/api/destination/getCompanyDestinations')
                                .then(res => {
                                    this.destinations = res.data;
                                    this.SuccessSnackbar = true;
                                })
                                .catch(err => {
                                    this.ErrorSnackbar = true;
                                    alert(err)
                                }))
                        .catch(err => console.log(err));
                }
            }
        }
    }
</script>

<style scoped>
    @import '../../assets/css/AirlineAdmin.css';
    .flexcard {
        display: flex;
        flex-direction: column;
    }
</style>