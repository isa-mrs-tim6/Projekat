<template>
    <div>
        <form>
            <v-text-field
                    v-model="Name"
                    label="Name"
                    required
            ></v-text-field>
            <v-text-field
                    v-model="Capacity"
                    label="Capacity"
                    type="number"
                    required
            ></v-text-field>
            <v-text-field
                    v-model="Type"
                    label="Type"
                    required
            ></v-text-field>
            <v-text-field
                    v-model="PriceLow"
                    label="PriceLow"
                    type="number"
                    required
            ></v-text-field>
            <v-text-field
                    v-model="PriceHigh"
                    label="PriceHigh"
                    type="number"
                    required
            ></v-text-field>

            <v-btn @click="findVehicles">submit</v-btn>
            <v-btn @click="clear">clear</v-btn>
        </form>
        <Vehicles v-bind:vehicles="vehicles"/>
    </div>
</template>

<script>
    import Vehicles from "../components/Vehicles";
    import axios from 'axios';

    export default {
        name: "VehiclesSearch",
        components: {Vehicles},
        data() {
            return {
                vehicles: [],
                Name: '',
                Capacity: '',
                Type: '',
                PriceLow: '',
                PriceHigh: '',
                StartDate: '',
                EndDate: ''
            }
        },
        methods: {
            findVehicles(e) {
                e.preventDefault();
                const vehicleParams = {
                    'Name': this.Name,
                    'Capacity': function(val) {
                        try {
                            val =  parseInt(this.Capacity);
                        }catch (e) {
                            val = 0;
                        }
                        return val;
                    }(),
                    'Type': this.Type,
                    'PriceLow': function(val) {
                        try {
                            val =  parseFloat(this.PriceHigh);
                        }catch (e) {
                            val = 0;
                        }
                        return val;
                    }(),
                    'PriceHigh': function(val) {
                        try {
                            val =  parseFloat(this.PriceHigh);
                        }catch (e) {
                            val = 100000;
                        }
                        return val;
                    }(),
                    'StartDate': "2000-09-22T12:42:31Z",
                    'EndDate': "2000-09-22T12:42:31Z"
                };
                this.clear();
                axios.post('http://localhost:8000/api/rentACarCompany/' + this.$route.params.id + '/findVehicles', vehicleParams)
                        .then(res =>{
                                this.vehicles = [];
                                if(res.data != null){
                                    this.vehicles = res.data;
                                }
                            }
                        )
                        .catch(err => console.log(err))
                    .catch(err => console.log(err));
            },
            clear () {
                this.Name = '';
                this.Capacity = '';
                this.Type = '';
                this.PriceLow = '';
                this.PriceHigh = '';
            }
        },
    }
</script>

<style scoped>

</style>