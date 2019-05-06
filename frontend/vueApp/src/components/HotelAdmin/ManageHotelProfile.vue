<template>
    <v-container grid-list-xl text-xs-center style="height: 100vh;">
        <v-layout align-center justify-center column wrap fill-height>
            <v-flex style="width: 60vw">
                <v-card min-height="100%" class="flexcard">
                    <v-card-title primary-title>
                        <div class="headline font-weight-medium">Hotel profile</div>
                    </v-card-title>
                    <v-card-text class="grow">
                        <v-form ref="form" class="align-center justify-center">
                            <v-text-field
                                    v-model="hotelProfile.Name"
                                    label="Name"
                                    required
                                    prepend-icon="business"
                                    class="body-2"
                            ></v-text-field>
                            <v-text-field
                                    v-model="hotelProfile.Address"
                                    label="Address"
                                    required
                                    prepend-icon="location_on"
                                    class="body-2"
                            ></v-text-field>
                            <v-flex>
                                <v-layout row wrap fill-height>
                                    <v-text-field
                                            v-model="hotelProfile.Latitude"
                                            label="Latitude"
                                            required
                                            class="body-2"
                                            prepend-icon="satellite"
                                            :rules="[rules.required, rules.numeric]"
                                    ></v-text-field>
                                    <v-text-field
                                            v-model="hotelProfile.Longitude"
                                            label="Longitude"
                                            required
                                            class="body-2"
                                            prepend-icon="satellite"
                                            :rules="[rules.required, rules.numeric]"
                                    ></v-text-field>
                                </v-layout>
                            </v-flex>
                            <v-text-field
                                    v-model="hotelProfile.Description"
                                    label="Description"
                                    prepend-icon="message"
                                    class="body-2"
                                    required
                            ></v-text-field>
                        </v-form>
                    </v-card-text>
                    <v-card-actions>
                        <v-btn color="primary" @click="update">update</v-btn>
                        <v-btn @click="revert">revert</v-btn>
                    </v-card-actions>
                </v-card>
            </v-flex>
        </v-layout>
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
                rules: {
                    required: value => !!value || 'Required.',
                    numeric: value => !isNaN(value) || 'Numeric',
                },
            }
        },
        mounted() {
            axios.create({withCredentials: true}).get('http://localhost:8000/api/hotel/getProfile')
                .then(res => {
                        this.hotelProfile = res.data;
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
                    axios.create({withCredentials: true}).post('http://localhost:8000/api/hotel/updateProfile', this.hotelProfile)
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
    .flexcard {
        display: flex;
        flex-direction: column;
    }
</style>