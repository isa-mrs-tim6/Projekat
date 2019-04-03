<template>
    <div>
        <AirlineAdminNavbar></AirlineAdminNavbar>
        <v-form>
            <v-container>
                <v-text-field :rules="[() => !!AirlineProfile.Name || 'This field is required']" label="Airline Name"
                              v-model="AirlineProfile.Name"></v-text-field>
                <v-text-field :rules="[() => !!AirlineProfile.Address|| 'This field is required']"
                              label="Airline Address" v-model="AirlineProfile.Address"></v-text-field>
                <v-text-field :rules="[() => !!AirlineProfile.Longitude || 'This field is required']" label="Longitude"
                              v-model="AirlineProfile.Longitude"></v-text-field>
                <v-text-field :rules="[() => !!AirlineProfile.Latitude || 'This field is required']" label="Latitude"
                              v-model="AirlineProfile.Latitude"></v-text-field>
                <v-textarea label="Airline description" v-model="AirlineProfile.Promo"></v-textarea>
                <v-btn @click="update">Submit</v-btn>
                <v-btn @click="revert">Revert</v-btn>
            </v-container>
        </v-form>
    </div>
</template>
<script>
    import AirlineAdminNavbar from "../../components/AirlineAdmin/AirlineAdminNavbar";
    import axios from 'axios';

    export default {
        name: "AirlineAdmin",
        components: {AirlineAdminNavbar},
        data() {
            return {
                BackupAirlienProfile: {},
                AirlineProfile: {},
            }
        },
        mounted() {
            axios.get('http://localhost:8000/api/airline/' + this.$route.params.id + '/getProfile')
                .then(res => {
                        this.AirlineProfile = res.data;
                        this.BackupAirlienProfile = JSON.parse(JSON.stringify(this.AirlineProfile))
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
                    axios.post('http://localhost:8000/api/airline/' + this.$route.params.id + '/updateProfile', this.AirlineProfile)
                        .then(res => this.BackupAirlienProfile = JSON.parse(JSON.stringify(this.AirlineProfile)))
                        .catch(err => console.log(err));
                }

            },
            revert() {
                this.AirlineProfile = JSON.parse(JSON.stringify(this.BackupAirlienProfile));
            }
        },
    }
</script>
