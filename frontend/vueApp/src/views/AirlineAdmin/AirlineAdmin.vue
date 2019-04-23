<template>
    <div id="main">
        <airline-admin-nav-drawer></airline-admin-nav-drawer>

        <v-container grid-list-xl text-xs-center style="height: 100vh;">
            <v-layout align-center justify-center column wrap fill-height>
                <v-flex style="width: 60vw">
                    <v-card min-height="100%" class="flexcard">
                        <v-card-title primary-title>
                            <div class="headline font-weight-medium">Edit airline profile</div>
                        </v-card-title>
                        <v-card-text class="grow">
                            <v-form ref="form" class="align-center justify-center">
                                <v-container>
                                    <v-text-field :rules="[() => !!AirlineProfile.Name || 'This field is required']" label="Airline Name"
                                                  v-model="AirlineProfile.Name" prepend-icon="business"></v-text-field>
                                    <v-text-field :rules="[() => !!AirlineProfile.Address|| 'This field is required']"
                                                  label="Airline Address" v-model="AirlineProfile.Address" prepend-icon="location_on"></v-text-field>
                                    <v-text-field :rules="[() => !!AirlineProfile.Longitude || 'This field is required']" label="Longitude"
                                                  v-model="AirlineProfile.Longitude" prepend-icon="satellite"></v-text-field>
                                    <v-text-field :rules="[() => !!AirlineProfile.Latitude || 'This field is required']" label="Latitude"
                                                  v-model="AirlineProfile.Latitude" prepend-icon="satellite"></v-text-field>
                                    <v-textarea label="Airline description" v-model="AirlineProfile.Promo" prepend-icon="description"></v-textarea>
                                </v-container>
                            </v-form>
                        </v-card-text>
                        <v-card-actions>
                            <v-btn color="primary" @click="update">Submit</v-btn>
                            <v-btn @click="revert">Revert</v-btn>
                        </v-card-actions>
                    </v-card>
                </v-flex>
            </v-layout>
        </v-container>
        <v-snackbar v-model="SuccessSnackbar" :timeout=4000 :top="true" color="success">Successfully updated airline profile</v-snackbar>
        <v-snackbar v-model="ErrorSnackbar" :timeout=4000 :top="true" color="error">Failed to update airline profile</v-snackbar>
    </div>
</template>
<script>
    import AirlineAdminNavbar from "../../components/AirlineAdmin/AirlineAdminNavbar";
    import axios from 'axios';
    import AirlineAdminNavDrawer from "../../components/AirlineAdmin/AirlineAdminNavDrawer";

    export default {
        name: "AirlineAdmin",
        components: {AirlineAdminNavDrawer, AirlineAdminNavbar},
        data() {
            return {
                BackupAirlineProfile: {},
                SuccessSnackbar: false,
                ErrorSnackbar: false,
                AirlineProfile: {},
            }
        },
        mounted() {
            axios.create({withCredentials:true}).get('http://localhost:8000/api/airline/getProfile')
                .then(res => {
                        this.AirlineProfile = res.data;
                        this.BackupAirlineProfile = JSON.parse(JSON.stringify(this.AirlineProfile))
                    }
                )
                .catch(err => console.log(err));
        },
        methods: {
            update(e) {
                e.preventDefault();
                if (!this.AirlineProfile.Name || !this.AirlineProfile.Longitude || !this.AirlineProfile.Latitude || !this.AirlineProfile.Address){
                    return;
                }
                if (isNaN(Number(this.AirlineProfile.Longitude)) && isNaN(Number(this.AirlineProfile.Latitude))) {
                    alert("Enter a number for Longitude and Latitude");
                } else if (isNaN(Number(this.AirlineProfile.Longitude))) {
                    alert("Enter a number for Longitude");
                } else if (isNaN(Number(this.AirlineProfile.Latitude))) {
                    alert("Enter a number for Latitude");
                } else {
                    this.AirlineProfile.Latitude = Number(this.AirlineProfile.Latitude);
                    this.AirlineProfile.Longitude = Number(this.AirlineProfile.Longitude);
                    axios.create({withCredentials:true}).post('http://localhost:8000/api/airline/updateProfile', this.AirlineProfile)
                        .then(res => {this.BackupAirlineProfile = JSON.parse(JSON.stringify(this.AirlineProfile));
                                this.SuccessSnackbar = true; })
                        .catch(err => this.ErrorSnackbar = true);
                }

            },
            revert() {
                this.AirlineProfile = JSON.parse(JSON.stringify(this.BackupAirlineProfile));
            }
        },
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