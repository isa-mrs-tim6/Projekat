<template>
    <div id="main">
        <AirlineAdminNavbar></AirlineAdminNavbar>
        <v-form>
            <v-container>
                <v-text-field :rules="[() => !!Destination.Name || 'This field is required']" label="Destination Name"
                              v-model="Destination.Name"></v-text-field>
                <v-text-field :rules="[() => !!Destination.Longitude|| 'This field is required']"
                              label="Longitude" v-model="Destination.Longitude"></v-text-field>
                <v-text-field :rules="[() => !!Destination.Latitude || 'This field is required']" label="Latitude"
                              v-model="Destination.Latitude"></v-text-field>
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
        name: "UpdateDestination",
        components: {AirlineAdminNavbar},
        data() {
            return {
                BackupDestination: {},
                Destination: {},
            }
        },
        mounted() {
            axios.get('http://localhost:8000/api/destination/' + this.$route.params.id + '/getDestination')
                .then(res => {
                        this.Destination = res.data;
                        this.BackupDestination = JSON.parse(JSON.stringify(this.Destination))
                    }
                )
                .catch(err => console.log(err));
        },
        methods: {
            update(e) {
                e.preventDefault();
                if (!this.Destination.Name || !this.Destination.Longitude || !this.Destination.Latitude){
                    return;
                }
                if (isNaN(Number(this.Destination.Longitude)) && isNaN(Number(this.Destination.Latitude))) {
                    alert("Enter a number for Longitude and Latitude");
                } else if (isNaN(Number(this.Destination.Longitude))) {
                    alert("Enter a number for Longitude");
                } else if (isNaN(Number(this.Destination.Latitude))) {
                    alert("Enter a number for Latitude");
                } else {
                    this.Destination.Latitude = Number(this.Destination.Latitude);
                    this.Destination.Longitude = Number(this.Destination.Longitude);
                    axios.post('http://localhost:8000/api/destination/' + this.$route.params.id + '/updateDestination', this.Destination)
                        .then(res => this.BackupDestination = JSON.parse(JSON.stringify(this.Destination)))
                        .catch(err => console.log(err));
                }

            },
            revert() {
                this.Destination = JSON.parse(JSON.stringify(this.BackupDestination));
            }
        },
    }
</script>

<style scoped>
    #main {
        background-image: linear-gradient(to right bottom, #142eae, #005bca, #007ed2, #009ccd, #0bb7c7, #47c0c6, #67c8c6, #81d0c7, #6ecac4, #58c4c3, #3cbdc2, #00b7c1);
    }
</style>