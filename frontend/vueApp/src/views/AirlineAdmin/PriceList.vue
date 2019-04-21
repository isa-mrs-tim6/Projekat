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
            <v-dialog v-model="dialogSeat" max-width="1000px">
                <v-card>
                    <v-card-title>
                        <span class="headline">Edit seat map:</span>
                    </v-card-title>
                    <v-card-text>
                        <v-container>
                            <v-layout row-wrap justify-center>
                                <v-flex xs6>
                                    <div class = "airplane">
                                        <template v-for="(seat, index) in editedSeats.Seats">
                                            <v-layout row-wrap v-if="index % editedSeats.RowWidth  === 0">
                                            </v-layout>
                                            <div class="passage" v-if="index % editedSeats.RowWidth === editedSeats.RowWidth / 2"></div>
                                            <div v-if="editedSeats.Seats[index].ReservationID !== 0" class=seat :class="seatClassBinding(index)" :key="index">
                                                <span v-if="editedSeats.Seats[index].ReservationID !== 0" class="occupied"></span>
                                            </div>
                                            <div class="seat" :key="index" @click="chooseSeat(index)" :class="seatClassBinding(index)" v-else>
                                            </div>
                                        </template>
                                    </div>
                                </v-flex>
                                <v-flex xs5>
                                    <v-layout row-wrap>
                                        <v-flex xs12>
                                            <v-text-field label="Seat number" outline v-model="editedSeats.Number" @change="changeSeat"></v-text-field>
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
    export default {
        name: "PriceList",
        components: {AirlineAdminNavbar},
        data(){
            return{
                rules:{
                    required: value => !!value || 'Required.'
                },
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
                    this.editedSeats.Seats = item.Airplane.Seats;
                    const predicate = (a, b) => {
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
                    };
                    this.editedSeats.Seats.sort(predicate);
                }

            },
            chooseSeat(index){
                this.editedSeats.Number = this.editedSeats.Seats[index].Number;
                this.editedSeats.Class = this.editedSeats.Seats[index].Class;
                this.editedSeats.SelectedSeat = index;
            },
            seatClassBinding(index){
                let seatClass = this.editedSeats.Seats[index].Class;
                let disabled = this.editedSeats.Seats[index].Disabled;
                let retval = "";
                if(disabled){
                    retval += "disabled";
                }else{
                    switch(seatClass){
                        case "BUSINESS":
                            retval = "business";
                            break;
                        case "ECONOMIC":
                            retval = "economic";
                            break;
                        case "FIRST":
                            retval = "first";
                            break;
                        default:
                            retval = "";
                    }
                }
                if (index === this.editedSeats.SelectedSeat){
                    retval += " selected"
                }
                return retval;
            },
            changeSeat(){
                if(this.editedSeats.SelectedSeat === -1){
                    return;
                }
                this.editedSeats.Seats[this.editedSeats.SelectedSeat].Number = this.editedSeats.Number;
                this.editedSeats.Seats[this.editedSeats.SelectedSeat].Class = this.editedSeats.Class;
            },
            close(){
                this.dialogSeat = false;
                this.dialog = false;

                setTimeout(()=>{
                    this.editedItem = Object.assign({}, this.defaultItem);
                    this.editedSeats = Object.assign({}, this.defaultSeats);
                    this.editedIndex = -1;
                    }, 300);
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
        watch:{
            dialog(val){
                val || this.close()
            },
            dialogSeat(val){
                val || this.close()
            },
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