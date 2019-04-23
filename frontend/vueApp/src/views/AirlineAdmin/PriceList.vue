<template>
    <div>
        <airline-admin-navbar></airline-admin-navbar>

        <v-container>
            <v-layout row-wrap mb-3>
                <v-flex>
                    <h1>Edit price list:</h1>
                </v-flex>
            </v-layout>
            <h2>Filters:</h2>
            <v-layout row-wrap>
                <v-flex xs3>
                    <v-text-field v-model="filter.fromDestination" label="Origin destination"></v-text-field>
                </v-flex>
                <v-spacer></v-spacer>
                <v-flex xs3>
                    <v-text-field v-model="filter.toDestination" label="Destination destination"></v-text-field>
                </v-flex>
                <v-spacer></v-spacer>
                <v-flex xs3>
                    <v-text-field v-model="filter.date" label="Departure date"></v-text-field>
                </v-flex>
            </v-layout>
            <h2>Flights:</h2>
            <v-dialog v-model="dialogSeat" max-width="1000px" scrollable persistent>
                <v-card>
                    <v-card-title>
                            <v-flex xs12>
                                <v-tabs centered grow color="#eeeeee" v-model="tab">
                                    <v-tabs-slider color="black"></v-tabs-slider>
                                    <v-tab href="#addSeat">Add seat</v-tab>
                                    <v-tab href="#editSeat">Edit seat</v-tab>
                                </v-tabs>
                            </v-flex>
                    </v-card-title>
                    <v-tabs-items v-model="tab">
                        <v-tab-item value="addSeat">
                            <v-card-text>
                                <v-container>
                                    <v-layout row-wrap justify-center>
                                        <v-flex xs5>
                                            <v-layout row-wrap justify-center>
                                                <seat-map :editedSeats="this.editedSeats"></seat-map>
                                            </v-layout>
                                            <v-layout row-wrap mt-3 align-center class="legendBlock">
                                                <v-flex xs6>
                                                    <div class="seat first"></div>
                                                    <span class="body-2" style="position:relative; bottom:7px">First class seat</span>
                                                </v-flex>
                                                <v-flex xs 6>
                                                    <div class="seat business"></div>
                                                    <span class="body-2" style="position:relative; bottom:7px">Business class seat</span>
                                                </v-flex>
                                            </v-layout>
                                            <v-layout row-wrap align-center class="legendBlock">
                                                <v-flex xs6>
                                                    <div class="seat economic"></div>
                                                    <span class="body-2" style="position:relative; bottom:7px">Economic class seat</span>
                                                </v-flex>
                                                <v-flex xs6>
                                                    <div class="seat disabled"></div>
                                                    <span class="body-2" style="position:relative; bottom:7px">Disabled seat</span>
                                                </v-flex>
                                            </v-layout>
                                            <v-layout row-wrap align-center class="legendBlock">
                                                <div class="seat occupied leftOccupied"></div>
                                                <span class="body-2">Occupied seat</span>
                                            </v-layout>
                                        </v-flex>
                                        <v-flex xs6 pr-2>
                                            <v-layout row-wrap>
                                                <h5 class="headline">Add seat:</h5>
                                            </v-layout>
                                            <v-layout row-wrap>
                                                <v-flex xs8>
                                                    <v-select :items="seatClass" label="Seat class" outline v-model="newSeatClass"></v-select>
                                                </v-flex>
                                            </v-layout>
                                            <v-layout row-wrap justify-center>
                                                <v-flex xs8>
                                                    <v-btn absolute large @click="addSeat">Add</v-btn>
                                                </v-flex>
                                            </v-layout>
                                        </v-flex>
                                    </v-layout>
                                </v-container>
                            </v-card-text>
                        </v-tab-item>
                        <v-tab-item value="editSeat">
                            <v-card-text>
                                <v-container>
                                    <v-layout row-wrap justify-center>
                                        <v-flex xs5>
                                            <v-layout row-wrap justify-center>
                                                <seat-map :editedSeats="this.editedSeats"></seat-map>
                                            </v-layout>
                                            <v-layout row-wrap mt-3 align-center class="legendBlock">
                                                <v-flex xs6>
                                                    <div class="seat first"></div>
                                                    <span class="body-2" style="position:relative; bottom:7px">First class seat</span>
                                                </v-flex>
                                                <v-flex xs 6>
                                                    <div class="seat business"></div>
                                                    <span class="body-2" style="position:relative; bottom:7px">Business class seat</span>
                                                </v-flex>
                                            </v-layout>
                                            <v-layout row-wrap align-center class="legendBlock">
                                                <v-flex xs6>
                                                    <div class="seat economic"></div>
                                                    <span class="body-2" style="position:relative; bottom:7px">Economic class seat</span>
                                                </v-flex>
                                                <v-flex xs6>
                                                    <div class="seat disabled"></div>
                                                    <span class="body-2" style="position:relative; bottom:7px">Disabled seat</span>
                                                </v-flex>
                                            </v-layout>
                                            <v-layout row-wrap align-center class="legendBlock">
                                                <div class="seat occupied leftOccupied"></div>
                                                <span class="body-2">Occupied seat</span>
                                            </v-layout>
                                        </v-flex>
                                        <v-flex xs5>
                                            <v-layout row-wrap>
                                                <h5 class="headline">Edit seat:</h5>
                                            </v-layout>
                                            <v-layout row-wrap>
                                                <v-flex xs12>
                                                    <v-text-field label="Seat number" outline v-model="editedSeats.Number" @change="changeSeat" readonly></v-text-field>
                                                </v-flex>
                                            </v-layout>
                                            <v-layout row-wrap>
                                                <v-flex xs12>
                                                    <v-select :items="seatClass" label="Seat class" outline v-model="editedSeats.Class" @change="changeSeat"></v-select>
                                                </v-flex>
                                            </v-layout>
                                            <v-layout row>
                                                <v-flex xs12>
                                                    <div class="text-xs-center">
                                                        <template v-if="editedSeats.SelectedSeat !== -1">
                                                            <v-btn v-if="editedSeats.Seats[editedSeats.SelectedSeat].Disabled" @click="editSeatStatus" large>Enable</v-btn>
                                                            <v-btn v-else large @click="editSeatStatus">Disable</v-btn>
                                                        </template>
                                                    </div>
                                                </v-flex>
                                            </v-layout>
                                        </v-flex>
                                        </v-layout>
                                </v-container>
                            </v-card-text>
                        </v-tab-item>
                    </v-tabs-items>
                    <v-card-actions>
                        <v-spacer></v-spacer>
                        <v-btn color="blue darken-1" flat @click="close">Cancel</v-btn>
                        <v-btn color="blue darken-1" flat @click="saveSeats">Save</v-btn>
                    </v-card-actions>
                </v-card>
            </v-dialog>
            <v-dialog v-model="dialog" persistent max-width="500px">
                <v-card>
                    <v-card-title>
                        <span class="headline">Edit price list:</span>
                    </v-card-title>
                    <v-card-text>
                        <v-container>
                            <v-layout row-wrap>
                                <v-flex xs6>
                                    <v-text-field v-model="editedItem.PriceFIRSTCLASS" label="First class price" prepend-icon="attach_money" type="number" :rules="[rules.required]"></v-text-field>
                                </v-flex>
                            </v-layout>
                            <v-layout row-wrap>
                                <v-flex xs6>
                                    <v-text-field v-model="editedItem.PriceBUSINESS" label="Business class price" prepend-icon="attach_money" type="number" :rules="[rules.required]"></v-text-field>
                                </v-flex>
                            </v-layout>
                            <v-layout row-wrap>
                                <v-flex xs6>
                                    <v-text-field v-model="editedItem.PriceECONOMY" label="Economy class price" prepend-icon="attach_money" type="number" :rules="[rules.required]"></v-text-field>
                                </v-flex>
                            </v-layout>
                            <v-layout row-wrap>
                                <v-flex xs6>
                                    <v-text-field v-model="editedItem.SmallSuitcase" label="Suitcase < 21kg" prepend-icon="attach_money" type="number" :rules="[rules.required]"></v-text-field>
                                </v-flex>
                                <v-flex xs6>
                                    <v-text-field v-model="editedItem.BigSuitcase" label="Suitcase > 21kg" prepend-icon="attach_money" type="number" :rules="[rules.required]"></v-text-field>
                                </v-flex>
                            </v-layout>
                        </v-container>
                    </v-card-text>
                    <v-card-actions>
                        <v-spacer></v-spacer>
                        <v-btn color="blue darken-1" flat @click="close">Cancel</v-btn>
                        <v-btn color="blue darken-1" flat @click="save">Save</v-btn>
                    </v-card-actions>
                </v-card>
            </v-dialog>
            <v-layout row-wrap>
                <v-flex x12>
                    <v-data-table :headers="headers" :items="flightsToShow" class="elevation-1">
                        <template v-slot:items="props">
                            <td>{{props.item.Origin.Name}}</td>
                            <td>{{props.item.Destination.Name}}</td>
                            <td>{{props.item.Departure | moment}}</td>
                            <td>{{props.item.PriceECONOMY}}</td>
                            <td>{{props.item.PriceBUSINESS}}</td>
                            <td>{{props.item.PriceFIRSTCLASS}}</td>
                            <td>{{props.item.SmallSuitcase}}</td>
                            <td>{{props.item.BigSuitcase}}</td>
                            <td>
                                <v-icon class="ma-1" small @click="editItem(props.item, 0)">attach_money</v-icon>
                                <v-icon class="ma-1" small @click="editItem(props.item, 1)">event_seat</v-icon>
                            </td>
                        </template>
                    </v-data-table>
                </v-flex>
            </v-layout>
        </v-container>
    </div>
</template>

<script>
    import AirlineAdminNavbar from "../../components/AirlineAdmin/AirlineAdminNavbar";
    import axios from 'axios';
    import moment from 'moment';
    import SeatMap from "./SeatMap";
    export default {
        name: "PriceList",
        components: {SeatMap, AirlineAdminNavbar},
        predicate : (a, b) => {
            const map = {};
            map["FIRST"] = 1;
            map["BUSINESS"] = 2;
            map["ECONOMIC"] = 3;

            if (map[a.Class] < map[b.Class]) {
                return -1;
            }

            if (map[a.Class] > map[b.Class]) {
                return 1;
            }

            return 0;
        },
        data(){
            return{
                rules:{
                    required: value => !!value || 'Required.'
                },
                tab:"addSeat",
                newSeatClass:"",
                dialog: false,
                seatNumber: 0,
                idx: 0,
                dialogSeat: false,
                flights:[],
                filter: {
                    fromDestination: '',
                    toDestination: '',
                    date: ''
                },
                seatClass:['ECONOMIC', 'BUSINESS','FIRST'],
                editedItem:{
                    ID: '',
                    PriceFIRSTCLASS: '',
                    PriceBUSINESS: '',
                    PriceECONOMY: '',
                    SmallSuitcase: '',
                    BigSuitcase: ''
                },
                defaultItem:{
                    ID: '',
                    PriceFIRSTCLASS: '',
                    PriceBUSINESS: '',
                    PriceECONOMY: '',
                    SmallSuitcase: '',
                    BigSuitcase: ''
                },
                editedSeats:{
                    RowWidth: 0,
                    RowNum: 0,
                    LastRow: 0,
                    SelectedSeat: -1,
                    Seats: [],
                    Number: "",
                    Class: ""
                },
                defaultSeats:{
                    RowWidth: 0,
                    RowNum: 0,
                    LastRow: 0,
                    SelectedSeat: -1,
                    Seats: [],
                    Number: "",
                    Class: ""
                },
                editedIndex: -1,
                headers:[
                    {
                        text: 'Origin',
                        value: "Origin"
                    },
                    {
                        text: 'Destination',
                        value: "Destination"
                    },
                    {
                        text: 'Departure',
                        value: "Departure"
                    },
                    {
                        text: 'Economy',
                        value: "Economy"
                    },
                    {
                        text: 'Business',
                        value: "Business"
                    },
                    {
                        text: 'First',
                        value: 'First'
                    },
                    {
                        text: 'Small suitcase',
                        value: 'Small'
                    },
                    {
                        text: 'Big suitcase',
                        value: 'Big'
                    },
                    {
                        text: 'Action',
                        sortable: false
                    }
                ]
            }
        },
        created(){
            axios.create({withCredentials: true}).get("http://localhost:8000/api/flight/getCompanyFlights")
                .then(res => {
                    this.flights = res.data;
                });
        },
        methods:{
            editSeatStatus(){
                let index = this.editedSeats.SelectedSeat;
                let status = this.editedSeats.Seats[index].Disabled;
                this.editedSeats.Seats[index].Disabled = !status;
            },
            checkNumbers(){
                if(!this.editedItem.PriceFIRSTCLASS || !this.editedItem.PriceBUSINESS || !this.editedItem.PriceECONOMY){
                    return false;
                }
                if(!this.editedItem.BigSuitcase || !this.editedItem.SmallSuitcase){
                    return false;
                }
                var number = parseFloat(this.editedItem.PriceECONOMY);
                var ind = false;
                if(number < 0 ){
                    this.editedItem.PriceECONOMY = "";
                    ind = true;
                }
                number = parseFloat(this.editedItem.PriceBUSINESS);
                if(number < 0 ){
                    this.editedItem.PriceBUSINESS = "";
                    ind = true;
                }
                number = parseFloat(this.editedItem.PriceFIRSTCLASS);
                if(number < 0 ){
                    this.editedItem.PriceFIRSTCLASS = "";
                    ind = true;
                }
                number = parseFloat(this.editedItem.BigSuitcase);
                if(number < 0 ){
                    this.editedItem.BigSuitcase = "";
                    ind = true;
                }
                number = parseFloat(this.editedItem.SmallSuitcase);
                if(number < 0 ){
                    this.editedItem.SmallSuitcase = "";
                    ind = true;
                }
                return ind !== true;
            },
            editItem (item, idx) {
                this.idx = idx;
                this.editedIndex = this.flights.indexOf(item);
                this.editedItem = Object.assign({}, item);
                if(this.idx === 0){
                    this.dialog = true;
                }else if(this.idx === 1){
                    this.dialogSeat = true;
                    this.editedSeats.RowWidth = item.Airplane.RowWidth;
                    this.editedSeats.Seats = Object.assign(this.editedSeats.Seats, {});
                    this.editedSeats.Seats = JSON.parse(JSON.stringify(item.Airplane.Seats));
                }

            },
            changeSeat(){
                if(this.editedSeats.SelectedSeat === -1){
                    return;
                }
                this.editedSeats.Seats[this.editedSeats.SelectedSeat].Number = this.editedSeats.Number;
                this.editedSeats.Seats[this.editedSeats.SelectedSeat].Class = this.editedSeats.Class;
                this.editedSeats.Seats.sort(predicate);
            },
            close() {
                this.dialogSeat = false;
                this.dialog = false;
                this.editedItem = Object.assign({}, this.defaultItem);
                this.editedSeats = JSON.parse(JSON.stringify(this.defaultSeats));
                this.editedIndex = -1;
            },
            save(){
                if (idx === 1) {
                    if (!this.checkNumbers()) {
                        return;
                    }
                    let flight = {
                        FlightID: this.editedItem.ID.toString(),
                        PriceFIRSTCLASS: this.editedItem.PriceFIRSTCLASS,
                        PriceBUSINESS: this.editedItem.PriceBUSINESS,
                        PriceECONOMY: this.editedItem.PriceECONOMY,
                        SmallSuitcase: this.editedItem.SmallSuitcase,
                        BigSuitcase: this.editedItem.BigSuitcase
                    };
                    Object.assign(this.flights[this.editedIndex], this.editedItem);
                    axios.create({withCredentials: true}).post("http://localhost:8000/api/priceList/update", flight)
                        .then(res => {
                            axios.create({withCredentials: true}).get("http://localhost:8000/api/flight/getCompanyFlights")
                                .then(res => {
                                    this.flights = res.data;
                                });
                        });
                }else{

                }
                this.close()
            },
            addSeat(){
                if(!this.newSeatClass){
                    return;
                }
                let Seat = {
                    Number: this.editedSeats.Seats.length,
                    Class: this.newSeatClass,
                    Disabled: false,
                    QuickReserve: false,
                    ReservationID: 0
                };
                this.editedSeats.Seats.push(Seat);
                this.editedSeats.Seats.sort(this.$options.predicate);
            },
            saveSeats(){
                Object.assign(this.flights[this.editedIndex].Airplane.Seats, this.editedSeats.Seats);
                let flight = {
                    FlightID: this.flights[this.editedIndex].ID.toString(),
                    AirplaneObject: this.flights[this.editedIndex].Airplane
                };
                axios.create({withCredentials: true}).post("http://localhost:8000/api/flight/updateSeats", flight)
                    .then(res => {
                        axios.create({withCredentials: true}).get("http://localhost:8000/api/flight/getCompanyFlights")
                            .then(res => {
                                this.flights = res.data;
                            });
                    });
                this.close();
            }
        },
        computed:{
            flightsToShow() {
                let flights = this.flights;
                if(this.filter.fromDestination.length > 0){
                    flights = flights.filter(flight => flight.Origin.Name.toLowerCase().includes(this.filter.fromDestination.toLowerCase()));
                }
                if(this.filter.toDestination.length > 0 ){
                    flights = flights.filter(flight => flight.Destination.Name.toLowerCase().includes(this.filter.toDestination.toLowerCase()));
                }
                if(this.filter.date.length > 0 ){
                    let date;
                    flights = flights.filter(flight =>{date = moment(String(flight.Departure)).format('HH:mm DD/MM/YYYY');
                                                        return date.includes(this.filter.date);
                    });
                }
                return flights;
            }
        }
    }


</script>

<style scoped>
    @import '../../assets/css/SeatMap.css';
</style>