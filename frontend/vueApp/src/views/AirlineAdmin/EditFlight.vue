<template>
    <div id="main">
        <airline-admin-nav-drawer></airline-admin-nav-drawer>
        <v-container grid-list-xl text-xs-center style="height: 100vh;">
            <v-layout align-center justify-center column wrap fill-height>
                <v-flex style="width: 60vw">
                    <v-card min-height="100%" class="flexcard">
                        <v-card-title primary-title>
                            <div class="headline font-weight-medium">Edit flights:</div>
                        </v-card-title>
                        <v-card-text class="grow">
                            <h2 class=".display-3">Filters</h2>
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
                            <h2 class=".display-3">Flights</h2>
                            <v-dialog v-model="dialogSeat" max-width="1000px" scrollable persistent lazy>
                                <v-card>
                                    <v-card-title>
                                            <v-flex xs12>
                                                <v-tabs centered grow color="#eeeeee" v-model="tab" slider-color="black">
                                                    <v-tab href="#addSeat" key="addSeat">Add seat</v-tab>
                                                    <v-tab href="#editSeat" key="editSeat">Edit seat</v-tab>
                                                </v-tabs>
                                            </v-flex>
                                    </v-card-title>
                                    <v-tabs-items v-model="tab">
                                        <v-tab-item value="addSeat" key="addSeatDialog">
                                            <v-card-text>
                                                <v-container>
                                                    <v-layout row-wrap justify-center>
                                                        <v-flex xs5>
                                                            <v-layout row-wrap justify-center>
                                                                <seat-map :editedSeats="this.editedSeats" :tab="this.tab"></seat-map>
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
                                                        <v-flex xs5 ml-5>
                                                            <v-layout row-wrap>
                                                                <h5 class="headline">Add seat:</h5>
                                                            </v-layout>
                                                            <v-layout row-wrap>
                                                                <v-flex xs8>
                                                                    <v-select :items="seatClass" label="Seat class" outline v-model="newSeatClass"></v-select>
                                                                </v-flex>
                                                            </v-layout>
                                                            <v-layout row>
                                                                <v-flex xs8>
                                                                    <div class="text-xs-center">
                                                                        <v-btn large @click="addSeat">Add</v-btn>
                                                                    </div>
                                                                </v-flex>
                                                            </v-layout>
                                                        </v-flex>
                                                    </v-layout>
                                                </v-container>
                                            </v-card-text>
                                        </v-tab-item>
                                        <v-tab-item value="editSeat" key="editSeatDialog">
                                            <v-card-text>
                                                <v-container>
                                                    <v-layout row-wrap justify-center>
                                                        <v-flex xs5>
                                                            <v-layout row-wrap justify-center>
                                                                <seat-map :editedSeats="this.editedSeats" :tab="this.tab"></seat-map>
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
                                                        <v-flex xs5 ml-5>
                                                            <v-layout row-wrap>
                                                                <h5 class="headline">Edit seat:</h5>
                                                            </v-layout>
                                                            <v-layout row-wrap>
                                                                <v-flex xs8>
                                                                    <v-text-field label="Seat number" outline v-model="editedSeats.Number" @change="changeSeat" readonly></v-text-field>
                                                                </v-flex>
                                                            </v-layout>
                                                            <v-layout row-wrap>
                                                                <v-flex xs8>
                                                                    <v-select :items="seatClass" label="Seat class" outline v-model="editedSeats.Class" @change="changeSeat"></v-select>
                                                                </v-flex>
                                                            </v-layout>
                                                            <v-layout row>
                                                                <v-flex xs8>
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
            <v-dialog v-model="dialogQuickReservation" max-width="1200px" persistent>
                <FlightQuickReservation :flight=this.editedItem :seats=this.editedSeats @closeReservationDialog="close" ref="QuickResComponent"></FlightQuickReservation>
            </v-dialog>
            <v-layout row-wrap mt-2>
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
                                <v-icon class="ma-1" small @click="editItem(props.item, 2)">add_box</v-icon>
                                <v-icon class="ma-1" small @click="editItem(props.item, 1)">event_seat</v-icon>
                                <v-icon class="ma-1" small @click="editItem(props.item, 0)">attach_money</v-icon>
                            </td>
                        </template>
                    </v-data-table>
                </v-flex>
            </v-layout>
                        </v-card-text>
                    </v-card>
                </v-flex>
            </v-layout>
        </v-container>
        <v-snackbar v-model="SuccessSnackbar" :timeout=4000 :top="true" color="success">Successfully updated flight</v-snackbar>
        <v-snackbar v-model="ErrorSnackbar" :timeout=4000 :top="true" color="error">Failed to update flight</v-snackbar>
    </div>
</template>

<script>
    import axios from 'axios';
    import moment from 'moment';
    import SeatMap from "../../components/AirlineAdmin/SeatMap";
    import AirlineAdminNavDrawer from "../../components/AirlineAdmin/AirlineAdminNavDrawer";
    import FlightQuickReservation from "../../components/AirlineAdmin/FlightQuickReservation";
    export default {
        name: "EditFlight",
        components: {FlightQuickReservation, AirlineAdminNavDrawer, SeatMap},
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
                SuccessSnackbar: false,
                ErrorSnackbar: false,
                tab:"",
                newSeatClass:"",
                dialog: false,
                seatNumber: 0,
                idx: 0,
                dialogSeat: false,
                dialogQuickReservation: false,
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
                    BigSuitcase: '',
                    Airplane:{
                        Seats:[]
                    }
                },
                defaultItem:{
                    ID: '',
                    PriceFIRSTCLASS: '',
                    PriceBUSINESS: '',
                    PriceECONOMY: '',
                    SmallSuitcase: '',
                    BigSuitcase: '',
                    Airplane:{
                        Seats:[]
                    }
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
                        value: "Departure",
                    },
                    {
                        text: 'Economy',
                        value: "Economy",
                        sortable: false
                    },
                    {
                        text: 'Business',
                        value: "Business",
                        sortable: false
                    },
                    {
                        text: 'First',
                        value: 'First',
                        sortable: false
                    },
                    {
                        text: 'Small suitcase',
                        value: 'Small',
                        sortable: false
                    },
                    {
                        text: 'Big suitcase',
                        value: 'Big',
                        sortable: false
                    },
                    {
                        text: 'Action',
                        sortable: false
                    }
                ]
            }
        },
        created(){
            axios.create({withCredentials: true}).get("http://ec2-35-159-21-254.eu-central-1.compute.amazonaws.com:8000/api/flight/getCompanyFlights")
                .then(res => {
                    this.flights = res.data;
                });
        },
        mounted(){
            this.checkFirstPass();
        },
        methods:{
            checkFirstPass(){
                axios.create({withCredentials: true}).get("http://ec2-35-159-21-254.eu-central-1.compute.amazonaws.com:8000/api/admin/checkFirstPass")
                    .then(
                        res =>{
                            if(res.data === false){
                                alert("Please update your password before accessing this feature");
                                this.$router.replace("admin_profile");
                            }
                        }
                    )
            },
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
                this.editedSeats.RowWidth = item.Airplane.RowWidth;
                this.editedSeats.Seats = Object.assign(this.editedSeats.Seats, {});
                this.editedSeats.Seats = JSON.parse(JSON.stringify(item.Airplane.Seats));
                this.editedSeats.Seats.sort(this.$options.predicate);
                if(this.idx === 0){
                    this.dialog = true;
                    return;
                }else if(this.idx === 1){
                    this.tab = "addSeat";
                    this.dialogSeat = true;
                }else if( this.idx === 2){
                    this.dialogQuickReservation = true;
                    this.$refs.QuickResComponent.getReservation(this.editedItem.ID);
                }
            },
            changeSeat(){
                if(this.editedSeats.SelectedSeat === -1){
                    return;
                }
                this.editedSeats.Seats[this.editedSeats.SelectedSeat].Number = this.editedSeats.Number;
                this.editedSeats.Seats[this.editedSeats.SelectedSeat].Class = this.editedSeats.Class;
                this.editedSeats.SelectedSeat = -1;
                this.editedSeats.Number = "";
                this.editedSeats.Class = "";

            },
            close() {
                this.dialogSeat = false;
                this.dialog = false;
                this.dialogQuickReservation = false;
                this.editedItem = Object.assign({}, this.defaultItem);
                this.editedSeats = JSON.parse(JSON.stringify(this.defaultSeats));
                this.editedIndex = -1;
                axios.create({withCredentials: true}).get("http://ec2-35-159-21-254.eu-central-1.compute.amazonaws.com:8000/api/flight/getCompanyFlights")
                .then(res => {
                    this.flights = res.data;
                });
            },
            save(){
                if (!this.checkNumbers()) {
                    return;
                }
                let flight = {
                    FlightID: this.editedItem.ID.toString(),
                    PriceFIRSTCLASS: this.editedItem.PriceFIRSTCLASS.toString(),
                    PriceBUSINESS: this.editedItem.PriceBUSINESS.toString(),
                    PriceECONOMY: this.editedItem.PriceECONOMY.toString(),
                    SmallSuitcase: this.editedItem.SmallSuitcase.toString(),
                    BigSuitcase: this.editedItem.BigSuitcase.toString()
                };
                Object.assign(this.flights[this.editedIndex], this.editedItem);
                axios.create({withCredentials: true}).post("http://ec2-35-159-21-254.eu-central-1.compute.amazonaws.com:8000/api/priceList/update", flight)
                    .then(res => {
                        axios.create({withCredentials: true}).get("http://ec2-35-159-21-254.eu-central-1.compute.amazonaws.com:8000/api/flight/getCompanyFlights")
                            .then(res => {
                                this.flights = res.data;
                                this.SuccessSnackbar = true;
                            });
                    })
                    .catch(res=>{
                        this.ErrorSnackbar = true;
                    });
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
                axios.create({withCredentials: true}).post("http://ec2-35-159-21-254.eu-central-1.compute.amazonaws.com:8000/api/flight/updateSeats", flight)
                    .then(res => {
                        axios.create({withCredentials: true}).get("http://ec2-35-159-21-254.eu-central-1.compute.amazonaws.com:8000/api/flight/getCompanyFlights")
                            .then(res => {
                                this.flights = res.data;
                                this.SuccessSnackbar = true;
                            });
                    })
                    .catch(res=>{
                        this.ErrorSnackbar = true;
                    });
                this.close();
            },
        },
        watch:{
          tab: function(val){
              if(val === "addSeat"){
                  this.editedSeats.SelectedSeat = -1;
                  this.editedSeats.Number = "";
                  this.editedSeats.Class = "";
              }
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
    #main {
        background-image: linear-gradient(to right bottom, #142eae, #005bca, #007ed2, #009ccd, #0bb7c7, #47c0c6, #67c8c6, #81d0c7, #6ecac4, #58c4c3, #3cbdc2, #00b7c1);
    }
    @import '../../assets/css/SeatMap.css';    
    .flexcard {
        display: flex;
        flex-direction: column;
    }
</style>