<template>
    <v-container grid-list-xl text-xs-center>
        <v-layout align-center justify-center column wrap fill-height>
            <v-flex style="width: 40vw">
                <v-form>
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

                    <v-btn style="float: left;" @click="update">update</v-btn>
                    <v-btn style="float: right;" @click="revert">revert</v-btn>
                </v-form>
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

</style>