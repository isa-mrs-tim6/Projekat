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
                            <td><v-icon small class="mr-2" @click="editItem(props.item)">edit</v-icon></td>
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
                flights:[],
                filter: {
                    fromDestination: '',
                    toDestination: '',
                    date: ''
                },
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
            editItem (item) {
                this.editedIndex = this.flights.indexOf(item);
                this.editedItem = Object.assign({}, item);
                this.dialog = true
            },
            close(){
                this.dialog = false;
                setTimeout(()=>{
                    this.editedItem = Object.assign({}, this.defaultItem);
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
            }
        },
        watch:{
            dialog(val){
                val || this.close()
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

</style>