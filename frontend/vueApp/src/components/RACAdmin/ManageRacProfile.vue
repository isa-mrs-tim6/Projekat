<template>
    <v-container grid-list-xl text-xs-center style="height: 100vh;">
        <v-layout align-center justify-center column wrap fill-height>
            <v-card min-height="100%" class="flexcard" style="padding: 5px">
                <v-card-title primary-title>
                    <div class="headline font-weight-medium">Manage Company Profile:</div>
                </v-card-title>
                <v-card-text class="grow">
                    <v-layout align-center justify-center row wrap fill-height>
                        <v-flex style="width: 30vw">
                            <form>
                                <v-text-field
                                        v-model="racProfile.Name"
                                        label="Name"
                                        required
                                ></v-text-field>
                                <v-text-field
                                        v-model="racProfile.Address"
                                        label="Address"
                                        required
                                ></v-text-field>
                                <v-text-field
                                        v-model="racProfile.Latitude"
                                        label="Latitude"
                                        required
                                ></v-text-field>
                                <v-text-field
                                        v-model="racProfile.Longitude"
                                        label="Longitude"
                                        required
                                ></v-text-field>
                                <v-text-field
                                        v-model="racProfile.Promo"
                                        label="Promo"
                                        required
                                ></v-text-field>
                            </form>
                        </v-flex>
                    </v-layout>
                </v-card-text>
                <v-card-actions>
                    <v-btn color="primary" @click="update">update</v-btn>
                    <v-btn @click="revert">revert</v-btn>
                </v-card-actions>
            </v-card>
        </v-layout>
        <v-snackbar v-model="SuccessSnackbar" :timeout=2000 :top="true" color="success">{{this.SuccessSnackbarText}}</v-snackbar>
        <v-snackbar v-model="ErrorSnackbar" :timeout=2000 :top="true" color="error">{{this.ErrorSnackbarText}}</v-snackbar>
    </v-container>
</template>

<script>
    import axios from 'axios';

    export default {
        name: "ManageRacProfile",
        data() {
            return {
                SuccessSnackbar: false,
                SuccessSnackbarText: '',
                ErrorSnackbar: false,
                ErrorSnackbarText: '',
                backupRACProfile: '',
                racProfile: '',
            }
        },
        mounted() {
            axios.create({withCredentials: true}).get('http://localhost:8000/api/rentACarCompany/getProfile')
                .then(res => {
                        this.racProfile = res.data;
                        this.backupRACProfile = JSON.parse(JSON.stringify(this.racProfile))
                    }
                )
                .catch(err => console.log(err));
        },
        methods: {
            update(e) {
                e.preventDefault();
                if (!this.racProfile.Longitude || !this.racProfile.Latitude){
                    this.ErrorSnackbar = true;
                    this.ErrorSnackbarText = "Longitude and Latitude cannot be empty";
                }else if(isNaN(parseFloat(this.racProfile.Longitude))){
                    this.ErrorSnackbar = true;
                    this.ErrorSnackbarText = "Longitude must be a number";
                }else if(isNaN(parseFloat(this.racProfile.Latitude))){
                    this.ErrorSnackbar = true;
                    this.ErrorSnackbarText = "Latitude must be a number";
                }else{
                    this.racProfile.Latitude = Number(this.racProfile.Latitude);
                    this.racProfile.Longitude = Number(this.racProfile.Longitude);
                    axios.create({withCredentials: true}).post('http://localhost:8000/api/rentACarCompany/updateProfile', this.racProfile)
                        .then(res => {
                            this.backupRACProfile = JSON.parse(JSON.stringify(this.racProfile));
                            this.SuccessSnackbar = true;
                            this.SuccessSnackbarText = "Company profile changed successfully";
                        })
                        .catch(err => console.log(err));
                }
            },
            revert () {
                this.racProfile = JSON.parse(JSON.stringify(this.backupRACProfile))
            }
        },
    }
</script>

<style scoped>

</style>