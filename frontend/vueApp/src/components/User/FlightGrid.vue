<template>
    <v-container>
        <v-layout column>
            <template v-if="searchType ==='oneWay'">
                <v-flex xs12 v-for="(value, index) in flights" v-bind:key="index">
                    <v-card dark class="white--text">
                        <v-card-title class="headline" primary-title>
                            {{value.Airline.Name}}
                        </v-card-title>
                        <v-layout row>
                            <v-flex xs10>
                                <v-card-title primary-title>
                                    <v-layout row>
                                        <v-flex xs5>
                                            <div>
                                                <div class="title"> <v-icon>flight_takeoff</v-icon> Origin: {{value.Origin.Name}}</div>
                                                <div class="title"> <v-icon>query_builder</v-icon> Departure: {{value.Departure | moment}}</div>
                                            </div>
                                        </v-flex>
                                        <v-flex xs5>
                                            <div>
                                                <div class="title"> <v-icon>flight_takeoff</v-icon> Destination: {{value.Destination.Name}}</div>
                                                <div class="title"> <v-icon>query_builder</v-icon> Departure: {{value.Landing | moment}}</div>
                                            </div>
                                        </v-flex>
                                    </v-layout>

                                </v-card-title>
                            </v-flex>
                            <v-flex xs2>
                                <div class="subheading">
                                    <v-icon>euro_symbol</v-icon> First class: {{value.PriceFIRSTCLASS}}
                                </div>
                                <div class="subheading">
                                    <v-icon>euro_symbol</v-icon> Business:   {{value.PriceBUSINESS}}
                                </div>
                                <div class="subheading">
                                    <v-icon>euro_symbol</v-icon> Economy:   {{value.PriceECONOMY}}
                                </div>
                            </v-flex>
                        </v-layout>
                        <v-divider light></v-divider>
                        <v-card-actions class="pa-3">
                            <v-btn block @click="reserve(value.ID, 'OneWay')">Reserve</v-btn>
                        </v-card-actions>
                    </v-card>
                </v-flex>
            </template>

            <template v-else-if="searchType === 'round' && flights.length !== 0 && flightsReturn.length !== 0">
                <h1>MULTI</h1>
                <v-data-table :items="flights" class="elevation-1">
                    <template v-slot:flights="props">
                        <td>{{props.item.Origin.Name}}</td>
                        <td>{{props.item.Destination.Name}}</td>
                        <td>{{props.item.Departure | moment}}</td>
                        <td>{{props.item.Landing | moment}}</td>
                    </template>
                </v-data-table>
                <v-data-table :items="flightsReturn" class="elevation-1">
                    <template v-slot:flights="props">
                        <td>{{props.item.Origin.Name}}</td>
                        <td>{{props.item.Destination.Name}}</td>
                        <td>{{props.item.Departure | moment}}</td>
                        <td>{{props.item.Landing | moment}}</td>
                    </template>
                </v-data-table>
            </template>

            <template v-else>
            </template>
        </v-layout>
    </v-container>
</template>

<script>
    export default {
        name: "FlightGrid",
        props: ["flights", "searchType"],
        methods: {
            reserve(id, type) {
                if (type === "OneWay") {
                    this.$router.push("/flightReservation?flightId=" + id);
                } else {
                    // TODO
                }
            }
        }
    }
</script>

<style scoped>

</style>