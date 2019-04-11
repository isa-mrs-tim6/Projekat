<template>
    <div>
        <airline-admin-navbar></airline-admin-navbar>
        <v-form>
            <v-container fluid>
                <h1>Add a flight</h1>
                <v-layout row wrap>
                    <v-flex xs12 sm3>
                        <h3>From:</h3>
                    </v-flex>
                    <v-flex xs12 sm3>
                        <h3>To:</h3>
                    </v-flex>
                </v-layout>

                <v-layout row wrap>
                    <v-flex xs12 sm3>
                        <v-select v-model="flight.from" label="Name" :items="destinations" item-text="Name"
                                  item-value="Name" :options="this.destinations" prepend-icon="map"></v-select>
                    </v-flex>
                    <v-flex xs12 sm3>
                        <v-select v-model="flight.to" label="Name" :items="destinations" item-text="Name"
                                  item-value="Name" :options="this.destinations" prepend-icon="map"></v-select>
                    </v-flex>
                </v-layout>
                <h3>Airplane:</h3>
                <v-layout row wrap>
                    <v-flex xs12 sm3>
                        <v-select v-model="flight.airplane" label="Airplane" :items="airplanes" item-text="Name"
                                  item-value="Name" :options="this.airplanes" prepend-icon="airplanemode_active"></v-select>
                    </v-flex>
                </v-layout>
                <h3>Departure: </h3>
                <v-layout row wrap>
                    <v-flex xs12 sm3>
                        <v-menu ref="menuDeparture" v-model="menuDeparture" :close-on-content-click="false" lazy transition="scale-transition"
                                :nudge-right="40" :return-value.sync="flight.departureDate" offset-y full-width max-width="290px" min-width="290px">
                            <template v-slot:activator="{ on }">
                                <v-text-field v-model="flight.departureDate" label="Departure date" prepend-icon="event" readonly v-on="on"></v-text-field>
                            </template>
                            <v-date-picker no-title scrollable v-model = "flight.departureDate">
                                <v-spacer></v-spacer>
                                <v-btn flat color="primary" @click="menuDeparture = false">Cancel</v-btn>
                                <v-btn flat color="primary" @click="$refs.menuDeparture.save(flight.departureDate)">OK</v-btn>
                            </v-date-picker>
                        </v-menu>
                    </v-flex>
                    <v-flex xs12 sm3>
                        <v-text-field v-model="flight.departureTime" :mask="mask" label="Departure time" prepend-icon="timelapse"></v-text-field>
                    </v-flex>
                </v-layout>
                <h3>Arrival:</h3>
                <v-layout row-wrap>
                    <v-flex xs12 sm3>
                        <v-menu ref="menuArrival" v-model="menuArrival" :close-on-content-click="false" lazy transition="scale-transition"
                                :nudge-right="40" :return-value.sync="flight.arrivalDate" offset-y full-width max-width="290px" min-width="290px">
                            <template v-slot:activator="{ on }">
                                <v-text-field v-model="flight.arrivalDate" label="Arrival date" prepend-icon="event" readonly v-on="on"></v-text-field>
                            </template>
                            <v-date-picker no-title scrollable v-model = "flight.arrivalDate">
                                <v-spacer></v-spacer>
                                <v-btn flat color="primary" @click="menuArrival = false">Cancel</v-btn>
                                <v-btn flat color="primary" @click="$refs.menuArrival.save(flight.arrivalDate)">OK</v-btn>
                            </v-date-picker>
                        </v-menu>
                    </v-flex>
                    <v-flex xs12 sm3>
                        <v-text-field v-model="flight.arrivalTime" :mask='mask' label="Arrival time" prepend-icon="timelapse"></v-text-field>
                    </v-flex>
                </v-layout>
                <h3>Travel information: </h3>
                <v-layout row-wrap>
                    <v-flex xs12 sm3>
                        <v-text-field v-model="flight.travelTime" label="Travel time" prepend-icon="timelapse"></v-text-field>
                    </v-flex>
                    <v-flex xs12 sm3>
                        <v-text-field v-model="flight.travelDistance" label="Travel distance [km]" prepend-icon="map"></v-text-field>
                    </v-flex>
                </v-layout>
                <h3>Add layover:</h3>
                <div v-for="(layover, index) in layovers" :key="index">
                    <v-layout row-wrap>
                        <v-flex xs12 sm3>
                            <v-text-field label="Layover" prepend-icon="map" v-model="layover.Address"></v-text-field>
                        </v-flex>
                        <v-btn small color="error" @click="deleteRow(index)">Delete</v-btn>
                    </v-layout>
                </div>
                <v-btn small @click="addRow">Add layover</v-btn>
                <h3>Price:</h3>
                <v-layout row-wrap>
                    <v-flex xs12 sm2>
                        <v-text-field v-model="flight.firstClass" label="First class price" prepend-icon="attach_money"></v-text-field>
                    </v-flex>
                    <v-flex sm2>
                        <v-text-field v-model="flight.businessClass" label="Business class price" prepend-icon="attach_money"></v-text-field>
                    </v-flex>
                    <v-flex sm2>
                        <v-text-field v-model="flight.economyClass" label="Economy class price" prepend-icon="attach_money"></v-text-field>
                    </v-flex>
                </v-layout>

                <v-btn v-on:click.prevent="sub">Submit</v-btn>
                <v-btn>Reset</v-btn>
            </v-container>
        </v-form>
    </div>
</template>

<script>
    import AirlineAdminNavbar from "../../components/AirlineAdmin/AirlineAdminNavbar";
    import axios from 'axios';
    import moment from 'moment';

    export default {
        name: "AddFlight",
        components: {AirlineAdminNavbar},
        data() {
            return {
                mask: 'time',
                destinations: [],
                airplanes: [],
                layovers: [],
                menuDeparture: false,
                menuArrival: false,
                flight: {
                    from: "",
                    to: "",
                    airplane: "",
                    travelTime: "",
                    travelDistance: "",
                    departureDate: "",
                    departureTime: "",
                    arrivalDate: "",
                    arrivalTime: "",
                    firstClass: "",
                    businessClass: "",
                    economyClass: ""
                }
            }
        },
        created() {
            axios.get("http://localhost:8000/api/destination/getDestinations")
                .then(res => {
                    this.destinations = res.data;
                });
            axios.get("http://localhost:8000/api/airplane/getAirplanes")
                .then(res => {
                    this.airplanes = res.data;
                })
        },
        methods: {
            addRow(e) {
                e.preventDefault();
                this.layovers.push({
                    Address:""
                })
            },
            deleteRow(index) {
                this.layovers.splice(index, 1)
            },
            sub(e){
                e.preventDefault();
                var dpTime = moment(this.flight.departureTime + " " + this.flight.departureDate,"HHmm YYYY-MM-DD").valueOf();
                var ldTime = moment(this.flight.arrivalTime + " " + this.flight.arrivalTime, "HHmm YYYY-MM-DD").valueOf();
                var flightDto = {
                    Origin: this.flight.from,
                    Destination: this.flight.to,
                    Airplane: this.flight.airplane,
                    Departure: dpTime.toString(),
                    Landing: ldTime.toString(),
                    Duration: this.flight.travelTime,
                    Distance: this.flight.travelDistance,
                    Layovers: this.layovers,
                    PriceECONOMY: this.flight.economyClass,
                    PriceBUSINESS: this.flight.businessClass,
                    PriceFIRSTCLASS: this.flight.firstClass
                };
                axios.post("http://localhost:8000/api/flight/add",flightDto)
                    .then(res=> console.log("succ"))
                    .catch(err => console.log(err));
            }
        },
    }
</script>

<style scoped>

</style>