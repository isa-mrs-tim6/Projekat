<template>
    <v-container fill-height fluid id="main">
        <v-layout>
            <v-flex>
                <airline-admin-nav-drawer/>
            </v-flex>
            <v-container>
                <v-layout>
                    <v-flex align-self-center>
                        <v-container grid-list-xl text-xs-center fill-height>
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
                                                    <v-layout>
                                                        <v-icon medium style="margin-left: 10px; margin-right: 8px">location_on</v-icon>
                                                        <gmap-autocomplete
                                                                style="width: 95%; border-bottom: 1px solid gray"
                                                                placeholder="Address"
                                                                :value="destination.Name"
                                                                @input="value = $event.target.value"
                                                                @place_changed="getAddressData">
                                                        </gmap-autocomplete>
                                                    </v-layout>
                                                    <v-flex>
                                                        <gmap-map
                                                                :center="{lat: destination.Latitude, lng: destination.Longitude}"
                                                                :zoom="8"
                                                                style="width:100%;  height: 400px;"
                                                        >
                                                            <gmap-marker
                                                                    :position="{lat: destination.Latitude, lng: destination.Longitude}"
                                                            ></gmap-marker>
                                                        </gmap-map>
                                                    </v-flex>
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

                    </v-flex>
                </v-layout>
            </v-container>
        </v-layout>
    </v-container>
</template>

<script>
    import axios from 'axios';
    import AirlineAdminNavDrawer from "../../components/AirlineAdmin/AirlineAdminNavDrawer";
    export default {
        name: "Destination",
        components: {AirlineAdminNavDrawer},
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
            },
            getAddressData: function (addressData) {
                this.destination.Name = addressData.formatted_address;
                this.destination.Latitude = addressData.geometry.location.lat();
                this.destination.Longitude = addressData.geometry.location.lng();
            }
        }
    }
</script>

<style scoped>
    #main {
        background-image: linear-gradient(to right bottom, #142eae, #005bca, #007ed2, #009ccd, #0bb7c7, #47c0c6, #67c8c6, #81d0c7, #6ecac4, #58c4c3, #3cbdc2, #00b7c1);
    }
    .flexcard {
        display: flex;
        flex-direction: column;
    }
</style>