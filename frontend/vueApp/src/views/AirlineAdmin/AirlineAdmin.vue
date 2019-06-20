<template>
    <v-container fill-height fluid id="main">
        <v-layout>
            <v-flex>
                <airline-admin-nav-drawer></airline-admin-nav-drawer>
            </v-flex>
            <v-container>
                <v-layout>
                    <v-flex align-self-center>
                        <v-container grid-list-xl text-xs-center fill-height>
                            <v-layout align-center justify-center column wrap fill-height>
                                <v-flex style="width: 60vw">
                                    <v-card class="flexcard">
                                        <v-card-title primary-title>
                                            <div class="headline font-weight-medium">Edit airline profile</div>
                                        </v-card-title>
                                        <v-card-text class="grow">
                                            <v-form ref="form" class="align-center justify-center">
                                                <v-container>
                                                    <v-text-field :rules="[() => !!AirlineProfile.Name || 'This field is required']" label="Airline Name"
                                                                  v-model="AirlineProfile.Name" prepend-icon="business"></v-text-field>
                                                    <v-layout>
                                                        <v-icon medium style="margin-left: 10px; margin-right: 8px">location_on</v-icon>
                                                        <gmap-autocomplete
                                                                style="width: 95%; border-bottom: 1px solid gray"
                                                                placeholder="Address"
                                                                :value="AirlineProfile.Address"
                                                                @input="value = $event.target.value"
                                                                @place_changed="getAddressData">
                                                        </gmap-autocomplete>
                                                    </v-layout>
                                                    <v-layout row>
                                                        <v-flex xs6>
                                                            <gmap-map
                                                                    :center="{lat: AirlineProfile.Latitude, lng: AirlineProfile.Longitude}"
                                                                    :zoom="8"
                                                                    style="width:100%;  height: 400px;"
                                                            >
                                                                <gmap-marker
                                                                        :position="{lat: AirlineProfile.Latitude, lng: AirlineProfile.Longitude}"
                                                                ></gmap-marker>
                                                            </gmap-map>
                                                        </v-flex>
                                                        <v-flex xs6>
                                                            <form
                                                                    enctype="multipart/form-data"
                                                                    action="http://localhost:8080/upload"
                                                                    method="post">
                                                                <div class="col-md-4 centered align-items-center mx-auto text-center">
                                                                    <img class="centered mx-auto text-center" height="400" width="100%" v-on:click="fileUpload" style="margin-bottom:50px;"  v-if="this.PictureLink" v-bind:src="this.PictureLink">
                                                                    <input type="file" accept="image/*," @change="updatePicture" id="picture" style="display: none">
                                                                </div>
                                                            </form>
                                                        </v-flex>
                                                    </v-layout>
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
        name: "AirlineAdmin",
        components: {AirlineAdminNavDrawer},
        data() {
            return {
                BackupAirlineProfile: {},
                SuccessSnackbar: false,
                ErrorSnackbar: false,
                AirlineProfile: {},
                PictureLink: '',
            }
        },
        mounted() {
            axios.create({withCredentials:true}).get('http://localhost:8000/api/airline/getProfile')
                .then(res => {
                        this.AirlineProfile = res.data;
                        this.PictureLink = 'http://localhost:8000/'+this.AirlineProfile.Picture;
                        this.BackupAirlineProfile = JSON.parse(JSON.stringify(this.AirlineProfile))
                    }
                )
                .catch(err => console.log(err));
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
            },
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
            },
            getAddressData: function (addressData) {
                this.AirlineProfile.Address = addressData.formatted_address;
                this.AirlineProfile.Latitude = addressData.geometry.location.lat();
                this.AirlineProfile.Longitude = addressData.geometry.location.lng();
            },
            updatePicture(e){
                const file = e.target.files[0];

                let formData = new FormData();
                formData.append('file', file);

                let a = null;

                axios.create({withCredentials: true}).post( 'http://localhost:8000/api/upload/updateAirlineProfilePicture',
                    formData,
                    {
                        headers: {
                            'Content-Type': 'multipart/form-data'
                        }
                    }
                ).then(res => {
                    a = res.data.toString();
                    a = a.substr(a.lastIndexOf(":")+2, a.lastIndexOf(`"`)-a.lastIndexOf(":")-2);
                    this.PictureLink = "1" + this.PictureLink;
                    this.PictureLink = this.PictureLink.substr(1, this.PictureLink.lastIndexOf("/")) + "/" + a;
                    this.AirlineProfile.Picture = a;
                });
            },
            fileUpload(e) {
                document.getElementById("picture").click();
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