<template>
    <div>
        <UserNavBar></UserNavBar>
        <v-container fluid fill-height style="margin-top: 200px">
            <v-layout column>
                <v-layout align-center justify-space-around row fill-height>
                    <v-flex xs3 fill-height>
                        <v-card min-height="60vh" dark tile fill-height class="flexcard">
                            <v-card-text class="grow">
                                <h1 class="display-3 text-xs-center">FLIGHTS</h1>
                                <h2 class="subheading text-xs-center">Travel all over the world</h2>
                            </v-card-text>
                            <v-card-actions>
                                <v-btn v-if="reservation == null" block @click="reserveFlight">Reserve</v-btn>
                                <v-btn v-else block disabled>Already reserved</v-btn>
                            </v-card-actions>
                        </v-card>
                    </v-flex>
                    <v-flex xs3 fill-height>
                        <v-card min-height="60vh" dark tile fill-height class="flexcard">
                            <v-card-text class="grow">
                                <h1 class="display-3 text-xs-center">HOTELS</h1>
                                <h2 class="subheading text-xs-center">Enjoy your stay in one of our many hotels</h2>
                            </v-card-text>
                            <v-card-actions>
                                <v-btn v-if="reservation != null && reservation.ReservationFlightID !== 0 && reservation.ReservationHotelID === 0" block @click="reserveHotel">Reserve</v-btn>
                                <v-btn v-else-if="reservation != null && reservation.ReservationFlightID !== 0 && reservation.ReservationHotelID !== 0" block disabled>Already reserved</v-btn>
                                <v-btn v-else block disabled>Access this feature by reserving a flight first</v-btn>
                            </v-card-actions>
                        </v-card>
                    </v-flex>
                    <v-flex xs3 fill-height>
                        <v-card min-height="60vh" dark tile fill-height class="flexcard">
                            <v-card-text class="grow">
                                <h1 class="display-3 text-xs-center">VEHICLES</h1>
                                <h2 class="subheading text-xs-center">Ride wherever, whenever</h2>
                            </v-card-text>
                            <v-card-actions>
                                <v-btn v-if="reservation != null && reservation.ReservationFlightID !== 0 && reservation.ReservationRentACarID === 0" block @click="reserveRentACar">Reserve</v-btn>
                                <v-btn v-else-if="reservation != null && reservation.ReservationFlightID !== 0 && reservation.ReservationRentACarID !== 0" block disabled>Already reserved</v-btn>
                                <v-btn v-else block disabled>Access this feature by reserving a flight first</v-btn>
                            </v-card-actions>
                        </v-card>
                    </v-flex>
                </v-layout>
            </v-layout>
        </v-container>
    </div>
</template>

<script>
    import UserNavBar from "../../components/User/UserNavBar";
    import axios from 'axios/index';

    export default {
        name: "ReservationChoice",
        components: {UserNavBar},
        data() {
            return {
                reservationID: this.$route.query.reservationID,
                passengers: this.$route.query.passengers,
                reservation: null,
            };
        },
        created(){
            if (this.reservationID != null) {
            axios.create({withCredentials:true}).get("http://ec2-18-195-170-20.eu-central-1.compute.amazonaws.com:8000/api/reservations/" + this.reservationID + "/reservation")
                .then(res =>{
                    this.reservation = res.data;
                });
            }
        },
        methods: {
            reserveFlight() {
                this.$router.push({ path: '/reserve_flight' });
            },
            reserveHotel() {
                this.$router.push({ path: '/user_hotels', query: { reservationID: this.reservationID, passengers: this.passengers}});
            },
            reserveRentACar() {
                this.$router.push({ path: '/user_cars', query: { reservationID: this.reservationID, passengers: this.passengers}});
            }
        }
    }
</script>

<style scoped>
    .flexcard {
        display: flex;
        flex-direction: column;
    }
</style>