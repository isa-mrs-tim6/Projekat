<template>
    <div>
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

            <v-btn @click="update">update</v-btn>
            <v-btn @click="revert">revert</v-btn>
        </form>
    </div>
</template>

<script>
    import axios from 'axios';

    export default {
        name: "ManageRacProfile",
        data() {
            return {
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
                if (isNaN(Number(this.racProfile.Longitude)) && isNaN(Number(this.racProfile.Latitude))){
                    alert("Enter a number for Longitude and Latitude");
                }else if(isNaN(Number(this.racProfile.Longitude))){
                    alert("Enter a number for Longitude");
                }else if(isNaN(Number(this.racProfile.Latitude))){
                    alert("Enter a number for Latitude");
                }else{
                    this.racProfile.Latitude = Number(this.racProfile.Latitude);
                    this.racProfile.Longitude = Number(this.racProfile.Longitude);
                    axios.create({withCredentials: true}).post('http://localhost:8000/api/rentACarCompany/updateProfile', this.racProfile)
                        .then(res => this.backupRACProfile = JSON.parse(JSON.stringify(this.racProfile)))
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