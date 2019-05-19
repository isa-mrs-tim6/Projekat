<template>
    <div>
        <v-card style="margin: 5px" dark v-for="(item, index) in this.reservations" :key="item.Master.ID">
            <v-layout row wrap>
                <v-flex xs2 class="resCenter">
                    Created:<br/>{{item.Master.CreatedAt}}
                </v-flex>
                <v-flex xs2 class="resCenter">
                    Departure:<br/>{{item.Master.ReservationFlight.Flight.Departure}}
                </v-flex>
                <v-spacer></v-spacer>
                <v-flex xs2 class="resCenter" v-if="item.Master.MasterRef !== 0">
                    INVITED BY:
                </v-flex>
                <v-flex style="text-align: right">
                    <v-btn class="button" @click="details(index)">details</v-btn><br/>
                    <v-btn class="button" color="error" :disabled="!checkFlight(item.Master.ReservationFlight.Flight.Departure)">cancel</v-btn>
                </v-flex>
            </v-layout>
        </v-card>
        <v-dialog v-model="dialog" persistent max-width="1000px">
            <v-card>
                <v-expansion-panel v-if="this.resDetails.ReservationFlightID !== 0">
                    <v-expansion-panel-content>
                        <template v-slot:header>
                            <div>Flight Reservation</div>
                        </template>
                        <v-card class="center">
                            <v-layout row>
                                <v-flex>
                                    Origin: <br/>{{this.origin}}
                                </v-flex>
                                <v-flex>
                                    Destination: <br/>{{this.destination}}
                                </v-flex>
                                <v-flex>
                                    Departure: <br/>{{this.departure}}
                                </v-flex>
                                <v-flex>
                                    Landing: <br/>{{this.landing}}
                                </v-flex>
                                <v-flex>
                                    Class: <br/>{{this.class}}
                                </v-flex>
                                <v-flex>
                                    Number: <br/>{{this.seatNum}}
                                </v-flex>
                            </v-layout>
                            <v-layout row>
                                <v-flex>
                                    Features:
                                </v-flex>
                                <v-flex v-for="(ff, index) in this.fFeatures" :key="ff.ID">
                                    Name: {{ff.Name}}<br/>
                                    Price: {{ff.Price}}
                                </v-flex>
                            </v-layout>
                            <v-layout row>
                                <v-flex xs3>
                                    Flight Rating: <v-rating
                                        v-model="this.fRating"
                                        background-color="indigo lighten-3"
                                        empty-icon="$vuetify.icons.ratingFull"
                                        readonly
                                        half-increments
                                        color="orange"></v-rating>
                                </v-flex>
                                <v-flex xs2>
                                    <v-btn>rate flight</v-btn>
                                </v-flex>
                                <v-spacer></v-spacer>
                                <v-flex xs3>
                                    Airline Rating: <v-rating
                                        v-model="this.fCompanyRating"
                                        background-color="indigo lighten-3"
                                        empty-icon="$vuetify.icons.ratingFull"
                                        readonly
                                        half-increments
                                        color="orange"></v-rating>
                                </v-flex>
                                <v-flex xs2>
                                    <v-btn>rate airline</v-btn>
                                </v-flex>
                            </v-layout>
                            <v-layout row>
                                Price: {{this.fPrice}}
                            </v-layout>
                        </v-card>
                    </v-expansion-panel-content>
                </v-expansion-panel>
                <v-expansion-panel v-if="this.resDetails.ReservationHotelID !== 0">
                    <v-expansion-panel-content>
                        <template v-slot:header>
                            <div>Hotel Reservation</div>
                        </template>
                        <v-card class="center">
                            Price: {{this.hPrice}}
                        </v-card>
                    </v-expansion-panel-content>
                </v-expansion-panel>
                <v-expansion-panel v-if="this.resDetails.ReservationRentACarID !== 0">
                    <v-expansion-panel-content>
                        <template v-slot:header>
                            <div>Vehicle Reservation</div>
                        </template>
                        <v-card class="center">
                            <v-layout row>
                                <v-flex>
                                    Beginning:<br/>
                                    {{this.vBeginning}}
                                </v-flex>
                                <v-flex>
                                    End:<br/>
                                    {{this.vEnd}}
                                </v-flex>
                                <v-flex>
                                    Location:<br/>
                                    {{this.vLocation}}
                                </v-flex>
                            </v-layout>
                            <v-layout row>
                                <v-flex>
                                    Vehicle Name:<br/>
                                    {{this.vName}}
                                </v-flex>
                                <v-flex>
                                    Vehicle Type:<br/>
                                    {{this.vType}}
                                </v-flex>
                                <v-flex>
                                    Vehicle Capacity:<br/>
                                    {{this.vCapacity}}
                                </v-flex>
                            </v-layout>
                            <v-layout row>
                                <v-flex xs3>
                                    Vehicle Rating: <v-rating
                                        v-model="this.vRating"
                                        background-color="indigo lighten-3"
                                        empty-icon="$vuetify.icons.ratingFull"
                                        readonly
                                        half-increments
                                        color="orange"></v-rating>
                                </v-flex>
                                <v-flex xs2>
                                    <v-btn>rate vehicle</v-btn>
                                </v-flex>
                                <v-spacer></v-spacer>
                                <v-flex xs3>
                                    Company Rating: <v-rating
                                        v-model="this.vCompanyRating"
                                        background-color="indigo lighten-3"
                                        empty-icon="$vuetify.icons.ratingFull"
                                        readonly
                                        half-increments
                                        color="orange"></v-rating>
                                </v-flex>
                                <v-flex xs2>
                                    <v-btn>rate company</v-btn>
                                </v-flex>
                            </v-layout>
                            <v-layout row>
                                Price: {{this.rPrice}}
                            </v-layout>
                        </v-card>
                    </v-expansion-panel-content>
                </v-expansion-panel>
                <div class="center">
                    ADDITIONAL PASSENGERS:
                </div>
                <v-layout row v-for="(slave,index) in this.slaves" :key="slave.ID"
                          class="center" style="border-top: 1px solid black; border-bottom: 1px solid black">
                    <v-flex>
                        Name:<br/>{{slave.Name}}
                    </v-flex>
                    <v-flex>
                        Surname:<br/>{{slave.Surname}}
                    </v-flex>
                    <v-flex>
                        Passport:<br/>{{slave.Passport}}
                    </v-flex>
                    <v-flex style="text-align: right">
                        <v-btn color="error">cancel</v-btn>
                    </v-flex>
                </v-layout>
                <v-layout row>
                    <v-spacer></v-spacer>
                    <v-flex xs2 class="center" style="text-align: right">
                        TOTAL: {{this.total}}
                    </v-flex>
                    <v-flex xs2 class="center" style="text-align: right">
                        <v-btn @click="close()">close</v-btn>
                    </v-flex>
                </v-layout>

            </v-card>
        </v-dialog>
    </div>
</template>

<script>
    import axios from "axios";
    import moment from 'moment';

    export default {
        name: "UserReservations",
        data (){
            return {
                reservations: [],
                dialog: false,
                slaves: [],
                resDetails: '',
                total: '',
                fPrice: '',
                hPrice: '',
                rPrice: '',
                origin: '',
                destination: '',
                departure: '',
                landing: '',
                fRating: 0,
                fCompanyRating: 0,
                fFeatures: [],
                class: '',
                seatNum: '',
                vName: '',
                vType: '',
                vCapacity: '',
                vBeginning: '',
                vEnd: '',
                vRating: '',
                vCompanyRating: '',
                vLocation: '',
            }
        },
        beforeCreate() {
            axios.create({withCredentials: true}).get("http://localhost:8000/api/reservations/getReservations").
            then(res => {
                console.log(res.data);
                this.reservations = res.data;
                for(let i = 0; i < this.reservations.length; i++) {
                    var date = moment(this.reservations[i].Master.CreatedAt).format('DD.MM.YYYY HH:mm');
                    this.reservations[i].Master.CreatedAt = date.toString();

                    date = moment(this.reservations[i].Master.ReservationFlight.Flight.Departure).format('DD.MM.YYYY HH:mm');
                    this.reservations[i].Master.ReservationFlight.Flight.Departure = date.toString();

                    date = moment(this.reservations[i].Master.ReservationFlight.Flight.Landing).format('DD.MM.YYYY HH:mm');
                    this.reservations[i].Master.ReservationFlight.Flight.Landing = date.toString();

                    date = moment(this.reservations[i].Master.ReservationFlight.Flight.Landing).format('DD.MM.YYYY');
                    this.reservations[i].Master.ReservationRentACar.Beginning = date.toString();

                    date = moment(this.reservations[i].Master.ReservationFlight.Flight.Landing).format('DD.MM.YYYY');
                    this.reservations[i].Master.ReservationRentACar.End = date.toString();
                }
            })
        },
        methods : {
            checkFlight(d){
                let date = moment(d);
                return date.isAfter(moment());
            },
            details(index){
                this.dialog = true;
                this.slaves = this.reservations[index].Slaves;
                this.resDetails = this.reservations[index].Master;
                this.fPrice = this.reservations[index].Master.ReservationFlight.Price;
                this.hPrice = this.reservations[index].Master.ReservationHotel.Price;
                this.rPrice = this.reservations[index].Master.ReservationRentACar.Price;
                this.total = this.fPrice + this.hPrice + this.rPrice;
                this.origin = this.reservations[index].Master.ReservationFlight.Flight.Origin.Name;
                this.destination = this.reservations[index].Master.ReservationFlight.Flight.Destination.Name;
                this.departure = this.reservations[index].Master.ReservationFlight.Flight.Departure;
                this.landing = this.reservations[index].Master.ReservationFlight.Flight.Landing;
                this.fRating = this.reservations[index].Master.ReservationFlight.FlightRating;
                this.fCompanyRating = this.reservations[index].Master.ReservationFlight.CompanyRating;
                this.class = this.reservations[index].Master.ReservationFlight.Seat.Class;
                this.seatNum = this.reservations[index].Master.ReservationFlight.Seat.Number;
                this.fFeatures = this.reservations[index].Master.ReservationFlight.Features;
                for(let i=0; i < this.reservations[index].Master.ReservationFlight.Features.length; i++){
                    this.total += this.reservations[index].Master.ReservationFlight.Features[i].Price;
                }

                this.vName =  this.reservations[index].Master.ReservationRentACar.Vehicle.Name;
                this.vType = this.reservations[index].Master.ReservationRentACar.Vehicle.Type;
                this.vCapacity = this.reservations[index].Master.ReservationRentACar.Vehicle.Capacity;
                this.vBeginning = this.reservations[index].Master.ReservationRentACar.Beginning;
                this.vEnd = this.reservations[index].Master.ReservationRentACar.End;
                this.vLocation = this.reservations[index].Master.ReservationRentACar.Location;
                this.vRating = this.reservations[index].Master.ReservationRentACar.VehicleRating;
                this.vCompanyRating = this.reservations[index].Master.ReservationRentACar.CompanyRating;

            },
            close(){
                this.dialog = false;
                this.slaves = [];
                this.resDetails = '';
                this.total = '';
                this.fPrice = '';
                this.hPrice = '';
                this.rPrice = '';
                this.origin = '';
                this.destination = '';
                this.departure = '';
                this.destination = '';
                this.class = '';
                this.seatNum = '';
                this.fFeatures = [];
                this.vName =  '';
                this.vType = '';
                this.vCapacity = '';
                this.vBeginning = '';
                this.vEnd = '';
                this.vLocation = '';
                this.vRating = '';
                this.vCompanyRating = '';
            },
        }
    }
</script>

<style scoped>
    .button{
        padding: 0;
        margin: 2px 0px;
    }
    .center {
        margin: auto;
        font-size: 16px;
        padding: 20px;
        word-wrap: break-word;
    }
    .resCenter{
        margin: auto;
        font-size: 16px;
        padding: 5px;
        word-wrap: break-word;
    }
</style>