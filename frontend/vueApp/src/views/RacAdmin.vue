<template>
    <div>
        <RacAdminNavbar></RacAdminNavbar>
        <form>
            <h1> Rent-a-car company <span id="routeID">{{ $route.params.id }}</span> </h1>
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
    import RacAdminNavbar from "../components/RacAdminNavbar";
    import axios from 'axios';

    export default {
        name: "RacAdmin",
        components: {RacAdminNavbar},
        data() {
            return {
                backupRACProfile: '',
                racProfile: '',
            }
        },
        mounted() {
            var hotelID = document.getElementById("routeID").innerHTML;
            axios.get('http://localhost:8000/api/rentACarCompany/'+hotelID+'/getProfile')
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
                var hotelID = document.getElementById("routeID").innerHTML;
                axios.post('http://localhost:8000/api/rentACarCompany/'+hotelID+'/updateProfile', this.racProfile)
                    .then(res => this.backupRACProfile = JSON.parse(JSON.stringify(this.racProfile)))
                    .catch(err => console.log(err));
            },
            revert () {
                this.racProfile = JSON.parse(JSON.stringify(this.backupRACProfile))
            }
        },
    }
</script>

<style scoped>

</style>