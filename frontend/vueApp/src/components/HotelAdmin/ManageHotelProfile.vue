<template>
    <v-container grid-list-xl text-xs-center>
        <v-layout align-center justify-center column wrap fill-height>
            <v-flex style="width: 60vw">
                <v-card min-height="100%" class="flexcard">
                    <v-card-title primary-title>
                        <div class="headline font-weight-medium">Hotel profile</div>
                    </v-card-title>
                    <v-card-text class="grow">
                        <v-layout row>
                            <v-flex xs12>
                                <v-form ref="form" class="align-center justify-center">
                                    <v-text-field
                                            v-model="hotelProfile.Name"
                                            label="Name"
                                            required
                                            prepend-icon="business"
                                            class="body-2"
                                    ></v-text-field>
                                    <v-layout>
                                        <v-icon medium style="margin-left: 10px; margin-right: 8px">location_on</v-icon>
                                        <gmap-autocomplete
                                                style="width: 95%; border-bottom: 1px solid gray"
                                                placeholder="Address"
                                                :value="hotelProfile.Address"
                                                @input="value = $event.target.value"
                                                @place_changed="getAddressData">
                                        </gmap-autocomplete>
                                    </v-layout>
                                    <v-layout row>
                                        <v-flex xs6>
                                            <v-flex>
                                                <gmap-map
                                                        :center="{lat: hotelProfile.Latitude, lng: hotelProfile.Longitude}"
                                                        :zoom="8"
                                                        style="width:100%;  height: 400px;"
                                                >
                                                    <gmap-marker
                                                            :position="{lat: hotelProfile.Latitude, lng: hotelProfile.Longitude}"
                                                    ></gmap-marker>
                                                </gmap-map>
                                            </v-flex>
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
                                    <v-text-field
                                            v-model="hotelProfile.Description"
                                            label="Description"
                                            prepend-icon="message"
                                            class="body-2"
                                            required
                                    ></v-text-field>
                                </v-form>
                            </v-flex>
                        </v-layout>
                    </v-card-text>
                    <v-card-actions>
                        <v-btn color="primary" @click="update">update</v-btn>
                        <v-btn @click="revert">revert</v-btn>
                    </v-card-actions>
                </v-card>
            </v-flex>
        </v-layout>
        <v-snackbar
                v-model="snackbar_fields.snackbar"
                :timeout="4000"
                :top="true"
                :color="snackbar_fields.color"
        >
            {{ snackbar_fields.text }}
            <v-btn
                    flat
                    @click="snackbar_fields.snackbar = false"
            >
                Close
            </v-btn>
        </v-snackbar>
    </v-container>
</template>

<script>
    import axios from 'axios';

    export default {
        name: "ManageHotelProfile",
        data() {
            return {
                backupHotelProfile: '',
                hotelProfile: '',
                PictureLink: '',
                rules: {
                    required: value => !!value || 'Required.',
                    numeric: value => !isNaN(value) || 'Numeric',
                },
                snackbar_fields: {
                    snackbar: false,
                    text: null,
                    color: null,
                },
            }
        },
        mounted() {
            axios.create({withCredentials: true}).get('http://ec2-35-159-21-254.eu-central-1.compute.amazonaws.com:8000/api/hotel/getProfile')
                .then(res => {
                        this.hotelProfile = res.data;
                        this.PictureLink = 'http://ec2-35-159-21-254.eu-central-1.compute.amazonaws.com:8000/'+this.hotelProfile.Picture;
                        this.backupHotelProfile = JSON.parse(JSON.stringify(this.hotelProfile))
                    }
                )
                .catch(err => alert(err));
        },
        methods: {
            update(e) {
                e.preventDefault();
                if (isNaN(Number(this.hotelProfile.Longitude)) && isNaN(Number(this.hotelProfile.Latitude))){
                    alert("Enter a number for Longitude and Latitude");
                }else if(isNaN(Number(this.hotelProfile.Longitude))){
                    alert("Enter a number for Longitude");
                }else if(isNaN(Number(this.hotelProfile.Latitude))){
                    alert("Enter a number for Latitude");
                }else {
                    this.hotelProfile.Latitude = Number(this.hotelProfile.Latitude);
                    this.hotelProfile.Longitude = Number(this.hotelProfile.Longitude);
                    axios.create({withCredentials: true}).post('http://ec2-35-159-21-254.eu-central-1.compute.amazonaws.com:8000/api/hotel/updateProfile', this.hotelProfile)
                        .then(res => {
                            this.backupHotelProfile = JSON.parse(JSON.stringify(this.hotelProfile))
                            this.snackbar_fields.text = "Update successful";
                            this.snackbar_fields.color = "success";
                            this.snackbar_fields.snackbar = true;
                            this.loader = null;
                        })
                        .catch(err => {
                            this.snackbar_fields.text = "Update failed";
                            this.snackbar_fields.color = "error";
                            this.snackbar_fields.snackbar = true;
                            this.loader = null;
                        });
                }
            },
            revert () {
                this.hotelProfile = JSON.parse(JSON.stringify(this.backupHotelProfile))
            },
            getAddressData: function (addressData) {
                this.hotelProfile.Address = addressData.formatted_address;
                this.hotelProfile.Latitude = addressData.geometry.location.lat();
                this.hotelProfile.Longitude = addressData.geometry.location.lng();
            },
            updatePicture(e){
                const file = e.target.files[0];

                let formData = new FormData();
                formData.append('file', file);

                let a = null;

                axios.create({withCredentials: true}).post( 'http://ec2-35-159-21-254.eu-central-1.compute.amazonaws.com:8000/api/upload/updateHotelProfilePicture',
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
                    this.hotelProfile.Picture = a;
                });
            },
            fileUpload(e) {
                document.getElementById("picture").click();
            }
        },
    }
</script>

<style scoped>
    .flexcard {
        display: flex;
        flex-direction: column;
    }
</style>