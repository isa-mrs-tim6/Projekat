<template>
    <v-container fluid pa-0>
        <v-layout row>
            <v-card dark style="width: 100%;">
                <v-card-title>
                    <h2>Search Vehicles:</h2>
                </v-card-title>
                <v-card-text>
                    <form>
                        <v-layout>
                            <v-flex xs2 ma-1>
                                <v-text-field
                                        v-model="Capacity"
                                        label="Capacity"
                                        type="number"
                                        :min="0"
                                        required
                                ></v-text-field>
                                <v-text-field
                                        v-model="Type"
                                        label="Type"
                                        required
                                ></v-text-field>
                            </v-flex>
                            <v-spacer></v-spacer>
                            <v-flex xs2 ma-1>
                                <v-menu v-model="menuStart" :close-on-content-click="false" lazy transition="scale-transition"
                                        offset-y full-width max-width="290px" min-width="290px">
                                    <template v-slot:activator="{ on }">
                                        <v-text-field v-model="StartDate" label="Start date" prepend-icon="event" readonly v-on="on" :rules="[rules.required]"></v-text-field>
                                    </template>
                                    <v-date-picker v-model="StartDate" no-title scrollable @input="menuStart = false"></v-date-picker>
                                </v-menu>
                                <v-menu v-model="menuEnd" :close-on-content-click="false" lazy transition="scale-transition"
                                        offset-y full-width max-width="290px" min-width="290px">
                                    <template v-slot:activator="{ on }">
                                        <v-text-field v-model="EndDate" label="End date" prepend-icon="event" readonly v-on="on" :rules="[rules.required]"></v-text-field>
                                    </template>
                                    <v-date-picker v-model="EndDate" no-title scrollable @input="menuEnd = false"></v-date-picker>
                                </v-menu>
                            </v-flex>
                            <v-spacer></v-spacer>
                            <v-flex xs7 ma-1>
                                <v-layout row>
                                    <v-subheader style="padding-top: 12px">Price Per Day: </v-subheader>
                                    <v-flex shrink style="width: 60px">
                                        <v-text-field
                                                v-model="Price[0]"
                                                class="mt-0"
                                                hide-details
                                                single-line
                                                type="number"
                                        ></v-text-field>
                                    </v-flex>
                                    <v-flex class="px-3">
                                        <v-range-slider
                                                v-model="Price"
                                                :max="200"
                                                :min="0"
                                                :step="1"
                                                value="Price range"
                                        ></v-range-slider>
                                    </v-flex>
                                    <v-flex shrink style="width: 60px">
                                        <v-text-field
                                                v-model="Price[1]"
                                                class="mt-0"
                                                hide-details
                                                single-line
                                                type="number"
                                        ></v-text-field>
                                    </v-flex>
                                </v-layout>
                            </v-flex>
                            <v-flex xs1 ma-1>
                                <v-btn @click="findVehicles">search</v-btn>
                            </v-flex>
                        </v-layout>
                    </form>
                </v-card-text>
            </v-card>
        </v-layout>
        <v-snackbar v-model="ErrorSnackbar" :timeout=2000 :top="true" color="error">{{this.ErrorSnackbarText}}</v-snackbar>
        <Vehicles v-bind:passengers="passengers" v-bind:reservationID="reservationID" v-bind:vehicles="vehicles" v-bind:startDate="StartUnix" v-bind:endDate="EndUnix"/>
    </v-container>
</template>

<script>
    import Vehicles from "../User/Vehicles";
    import axios from 'axios';
    import moment from 'moment';

    export default {
        name: "VehiclesSearch",
        props: ["reservationID", "passengers"],
        components: {Vehicles},
        data() {
            return {
                ErrorSnackbar: false,
                ErrorSnackbarText: '',
                vehicles: [],
                Capacity: 0,
                Type: '',
                Price: [0, 200],
                PriceLow: '',
                PriceHigh: '',
                StartDate: '',
                EndDate: '',
                menuStart: false,
                menuEnd: false,
                rules: {
                    required: value => !!value || 'Required.'
                },
                StartUnix: '',
                EndUnix: ''
            }
        },
        methods: {
            findVehicles(e) {
                e.preventDefault();
                if(this.StartDate === '' || this.EndDate === ''){
                    this.ErrorSnackbar = true;
                    this.ErrorSnackbarText = 'Please fill start and end date';
                    return;
                }
                this.StartUnix = moment(this.StartDate).valueOf().toString();
                this.EndUnix = moment(this.EndDate).valueOf().toString();

                const vehicleParams = {
                    'ID': parseInt(this.$route.params.id),
                    'Capacity': parseInt(this.Capacity),
                    'Type': this.Type,
                    'PriceLow': this.Price[0],
                    'PriceHigh': this.Price[1],
                    'StartDate': this.StartUnix,
                    'EndDate': this.EndUnix
                };
                this.clear();
                axios.post('http://localhost:8000/api/rentACarCompany/findVehicles', vehicleParams)
                        .then(res =>{
                            console.log(res.data);
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
            },
        },
    }
</script>

<style scoped>

</style>