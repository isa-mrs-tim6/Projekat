<template>
    <v-container fluid pa-0>
        <v-layout row>
            <v-card dark style="width: 100%;">
                <v-card-text>
                    <form>
                        <v-layout>
                            <v-flex xs3>
                                <v-text-field
                                        v-model="Name"
                                        label="Name"
                                        required
                                ></v-text-field>
                            </v-flex>
                            <v-spacer></v-spacer>
                            <v-flex xs1>
                                <v-text-field
                                        v-model="Capacity"
                                        label="Capacity"
                                        type="number"
                                        :min="0"
                                        required
                                >0</v-text-field>
                            </v-flex>
                            <v-spacer></v-spacer>
                            <v-flex xs2>
                                <v-text-field
                                        v-model="Type"
                                        label="Type"
                                        required
                                ></v-text-field>
                            </v-flex>
                            <v-spacer></v-spacer>
                            <v-flex xs2>
                                <v-menu v-model="menuStart" :close-on-content-click="false" lazy transition="scale-transition"
                                        offset-y full-width max-width="290px" min-width="290px">
                                    <template v-slot:activator="{ on }">
                                        <v-text-field v-model="StartDate" label="Start date" prepend-icon="event" readonly v-on="on" :rules="[rules.required]"></v-text-field>
                                    </template>
                                    <v-date-picker v-model="StartDate" no-title scrollable @input="menuStart = false"></v-date-picker>
                                </v-menu>
                            </v-flex>
                            <v-spacer></v-spacer>
                            <v-flex xs2>
                                <v-menu v-model="menuEnd" :close-on-content-click="false" lazy transition="scale-transition"
                                        offset-y full-width max-width="290px" min-width="290px">
                                    <template v-slot:activator="{ on }">
                                        <v-text-field v-model="EndDate" label="End date" prepend-icon="event" readonly v-on="on" :rules="[rules.required]"></v-text-field>
                                    </template>
                                    <v-date-picker v-model="EndDate" no-title scrollable @input="menuEnd = false"></v-date-picker>
                                </v-menu>
                            </v-flex>
                            <v-flex xs1>
                                <v-btn @click="findVehicles">submit</v-btn>
                            </v-flex>
                        </v-layout>
                    </form>
                    <v-layout row>
                        <v-subheader style="padding-top: 12px">Price Range: </v-subheader>
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
                </v-card-text>
            </v-card>
        </v-layout>
        <Vehicles v-bind:vehicles="vehicles"/>
    </v-container>
</template>

<script>
    import Vehicles from "../components/Vehicles";
    import axios from 'axios';
    import moment from 'moment';

    export default {
        name: "VehiclesSearch",
        components: {Vehicles},
        data() {
            return {
                vehicles: [],
                Name: '',
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
            }
        },
        methods: {
            findVehicles(e) {
                e.preventDefault();
                let start;
                let end;
                if(this.StartDate === ''){
                    start = "0";
                }else{
                    start = moment(this.StartDate).valueOf().toString();
                }
                if(this.StartDate === ''){
                    end = "0";
                }else{
                    end = moment(this.EndDate).valueOf().toString();
                }
                const vehicleParams = {
                    'Name': this.Name,
                    'Capacity': parseInt(this.Capacity),
                    'Type': this.Type,
                    'PriceLow': this.Price[0],
                    'PriceHigh': this.Price[1],
                    'StartDate': start,
                    'EndDate': end
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