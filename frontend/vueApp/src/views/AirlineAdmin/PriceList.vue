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
                            <v-layout row-wrap>
                                <v-flex xs3 >
                                    <div class="airplane">
                                        <div v-for="n in editedSeats.RowNum">
                                            <v-layout row-wrap class="row">
                                                <div v-for="m in editedSeats.RowWidth/2">
                                                   <div class="seat" @click="chooseSeat(n*editedSeats.RowWidth + m - editedSeats.RowWidth - 1)"
                                                                                    v-bind:class="{first: editedSeats.Seats[n*editedSeats.RowWidth + m - editedSeats.RowWidth - 1].Class === 'FIRST',
                                                                                    business: editedSeats.Seats[n*editedSeats.RowWidth + m - editedSeats.RowWidth - 1].Class === 'BUSINESS',
                                                                                    selected: editedSeats.SelectedSeat === n*editedSeats.RowWidth + m - editedSeats.RowWidth - 1,
                                                                                    economic: editedSeats.Seats[n*editedSeats.RowWidth + m - editedSeats.RowWidth - 1].Class === 'ECONOMIC'
                                                                                    }"></div>
                                                    {{n*editedSeats.RowWidth + m - editedSeats.RowWidth - 1}}
                                                </div>
                                                <div class="passage"></div>
                                                <div v-for="m in editedSeats.RowWidth/2">
                                                    <div class="seat" @click="chooseSeat(n*editedSeats.RowWidth + editedSeats.RowWidth / 2 + m - editedSeats.RowWidth - 1)"
                                                                                    v-bind:class="{first: editedSeats.Seats[n*editedSeats.RowWidth + editedSeats.RowWidth / 2 + m - editedSeats.RowWidth - 1].Class === 'FIRST',
                                                                                    selected: editedSeats.SelectedSeat === n*editedSeats.RowWidth + editedSeats.RowWidth / 2 + m - editedSeats.RowWidth - 1,
                                                                                    business: editedSeats.Seats[n*editedSeats.RowWidth + editedSeats.RowWidth / 2 + m - editedSeats.RowWidth - 1].Class === 'BUSINESS',
                                                                                    economic: editedSeats.Seats[n*editedSeats.RowWidth + editedSeats.RowWidth / 2 + m - editedSeats.RowWidth - 1].Class === 'ECONOMIC'}"></div>
                                                    {{n*editedSeats.RowWidth + editedSeats.RowWidth / 2 + m - editedSeats.RowWidth - 1}}
                                                </div>
                                            </v-layout>
                                        </div>
                                        <div v-if="editedSeats.LastRow <= (editedSeats.RowWidth / 2)">
                                            <v-layout row-wrap class="row">
                                                <div v-for="a in editedSeats.LastRow">
                                                    <div class="seat" @click="chooseSeat(editedSeats.RowNum * editedSeats.RowWidth + a - 1)"
                                                                                    v-bind:class="{first: editedSeats.Seats[editedSeats.RowNum * editedSeats.RowWidth + a - 1].Class === 'FIRST',
                                                                                    selected: editedSeats.SelectedSeat === editedSeats.RowNum * editedSeats.RowWidth + a - 1,
                                                                                    business: editedSeats.Seats[editedSeats.RowNum * editedSeats.RowWidth + a - 1].Class === 'BUSINESS',
                                                                                    economic: editedSeats.Seats[editedSeats.RowNum * editedSeats.RowWidth + a - 1].Class === 'ECONOMIC'}"></div>
                                                    {{editedSeats.RowNum * editedSeats.RowWidth + a}}
                                                </div>
                                                <div class="passage"></div>
                                                <div v-for="k in (editedSeats.RowWidth - editedSeats.LastRow)">
                                                    <div class="passage"></div>
                                                </div>
                                            </v-layout>
                                        </div>
                                        <div v-else>
                                            <v-layout row-wrap class="row">
                                                <div v-for="m in editedSeats.RowWidth/2">
                                                    <div class="seat" @click="chooseSeat(editedSeats.RowNum * editedSeats.RowWidth + m - 1)"
                                                                                    v-bind:class="{first: editedSeats.Seats[editedSeats.RowNum * editedSeats.RowWidth + m - 1].Class === 'FIRST',
                                                                                    selected: editedSeats.SelectedSeat === editedSeats.RowNum * editedSeats.RowWidth + m - 1,
                                                                                    business: editedSeats.Seats[editedSeats.RowNum * editedSeats.RowWidth + m - 1].Class === 'BUSINESS',
                                                                                    economic: editedSeats.Seats[editedSeats.RowNum * editedSeats.RowWidth + m - 1].Class === 'ECONOMIC'}"></div>
                                                    {{editedSeats.RowNum * editedSeats.RowWidth + m}}
                                                </div>
                                                <div class="passage"></div>
                                                <div v-for="m in (editedSeats.LastRow -(editedSeats.RowWidth/2))">
                                                    <div class="seat" @click="chooseSeat(editedSeats.RowNum * editedSeats.RowWidth + editedSeats.RowWidth/2 + m - 1)"
                                                                                    v-bind:class="{first: editedSeats.Seats[editedSeats.RowNum * editedSeats.RowWidth + editedSeats.RowWidth/2 + m - 1].Class === 'FIRST',
                                                                                    selected: editedSeats.SelectedSeat === editedSeats.RowNum * editedSeats.RowWidth + editedSeats.RowWidth/2 + m - 1,
                                                                                    business: editedSeats.Seats[editedSeats.RowNum * editedSeats.RowWidth + editedSeats.RowWidth/2 + m - 1].Class === 'BUSINESS',
                                                                                    economic: editedSeats.Seats[editedSeats.RowNum * editedSeats.RowWidth + editedSeats.RowWidth/2 + m - 1].Class === 'ECONOMIC'}"></div>
                                                    {{editedSeats.RowNum * editedSeats.RowWidth + editedSeats.RowWidth/2 + m}}
                                                </div>
                                                <div v-for="k in (editedSeats.RowWidth - editedSeats.LastRow)">
                                                    <div class="passage"></div>
                                                </div>
                                            </v-layout>
                                        </div>
                                    </div>
                                </v-flex>
                                <v-flex xs4>

                                </v-flex>
                                <v-flex xs4>
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
                    this.editedSeats.RowNum = Math.floor(item.Airplane.Seats.length / item.Airplane.RowWidth);
                    this.editedSeats.LastRow = item.Airplane.Seats.length % item.Airplane.RowWidth;
                    this.editedSeats.Seats = item.Airplane.Seats;
                }

            },
            chooseSeat(index){
                this.editedSeats.Number = this.editedSeats.Seats[index].Number;
                this.editedSeats.Class = this.editedSeats.Seats[index].Class;
                this.editedSeats.SelectedSeat = index;
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
                if (!this.checkNumbers()){
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
                axios.create({withCredentials:true}).post("http://localhost:8000/api/priceList/update", flight)
                    .then(res=>{
                        axios.create({withCredentials: true}).get("http://localhost:8000/api/flight/getCompanyFlights")
                            .then(res => {
                                this.flights = res.data;
                            });
                    });
                this.close()
            },
            saveSeats(){
                Object.assign(this.flights[this.editedIndex].Airplane.Seats, this.editedSeats.Seats);
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
            proba(){
              this.seatNumber++;
              return this.seatNumber
            },
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