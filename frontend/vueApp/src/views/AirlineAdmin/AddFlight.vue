<template>
    <div id = "main">
        <airline-admin-nav-drawer></airline-admin-nav-drawer>
    <div>
        <v-container grid-list-xl text-xs-center style="height: 100vh;">
            <v-layout align-center justify-center column wrap fill-height>
                <v-flex style="width: 60vw">
                    <v-card min-height="100%" class="flexcard" >
                        <v-card-title primary-title>
                            <div class="title font-weight-medium">Add a flight</div>
                        </v-card-title>
                        <v-card-text class="grow">
                            <v-layout row wrap>
                                <v-flex xs12 sm3>
                                    <div class=".title font-weight-bold" style="text-align: left">From:</div>
                                </v-flex>
                                <v-flex xs12 sm3 justify-start>
                                    <div class=".title font-weight-bold" style="text-align: left">To:</div>
                                </v-flex>
                            </v-layout>
                            <v-layout row wrap>
                                <v-flex xs12 sm3>
                                    <v-select v-model="flight.from" label="Name" :items="fromDestinations" item-text="Name"
                                              item-value="Name" :options="fromDestinations" prepend-icon="map"></v-select>
                                </v-flex>
                                <v-flex xs12 sm3>
                                    <v-select v-model="flight.to" label="Name" :items="toDestinations" item-text="Name"
                                              item-value="Name" :options="toDestinations" prepend-icon="map"></v-select>
                                </v-flex>
                            </v-layout>
                            <div class=".title font-weight-bold" style="text-align: left">Airplane</div>
                            <v-layout row wrap>
                                <v-flex xs12 sm3>
                                    <v-select v-model="flight.airplane" label="Airplane" :items="airplanes" item-text="Name"
                                              item-value="Name" :options="this.airplanes" prepend-icon="airplanemode_active"></v-select>
                                </v-flex>
                            </v-layout>
                            <div class=".title font-weight-bold" style="text-align: left">Departure</div>
                            <v-layout row wrap justify-start>
                                <v-flex xs12 sm3>
                                    <v-menu v-model="menuDeparture" :close-on-content-click="false" lazy transition="scale-transition"
                                            :nudge-right="40" offset-y full-width max-width="290px" min-width="290px">
                                        <template v-slot:activator="{ on }">
                                            <v-text-field v-model="flight.departureDate" label="Departure date" prepend-icon="event" readonly v-on="on"></v-text-field>
                                        </template>
                                        <v-date-picker no-title scrollable v-model = "flight.departureDate" @input="menuDeparture = false" :min="minDate"></v-date-picker>
                                    </v-menu>
                                </v-flex>
                                <v-flex xs12 sm3>
                                    <v-text-field v-model="flight.departureTime" :mask="mask" :rules="[rules.time, rules.required]" label="Departure time" prepend-icon="timelapse"></v-text-field>
                                </v-flex>
                            </v-layout>
                            <div class=".title font-weight-bold" style="text-align: left">Arrival:</div>
                            <v-layout row-wrap justify-start>
                                <v-flex xs12 sm3>
                                    <v-menu v-model="menuArrival" :close-on-content-click="false" lazy transition="scale-transition"
                                            :nudge-right="40" offset-y full-width max-width="290px" min-width="290px">
                                        <template v-slot:activator="{ on }">
                                            <v-text-field v-model="flight.arrivalDate" label="Arrival date" prepend-icon="event" readonly v-on="on"></v-text-field>
                                        </template>
                                        <v-date-picker no-title scrollable v-model = "flight.arrivalDate" @input="menuArrival = false" :min="minDateArrival"></v-date-picker>
                                    </v-menu>
                                </v-flex>
                                <v-flex xs12 sm3>
                                    <v-text-field v-model="flight.arrivalTime" :mask='mask' label="Arrival time" :rules="[rules.time, rules.required]" prepend-icon="timelapse"></v-text-field>
                                </v-flex>
                            </v-layout>
                            <div class=".title font-weight-bold" style="text-align: left">Travel information:</div>
                            <v-layout row-wrap>
                                <v-flex xs12 sm3>
                                    <v-text-field v-model="flight.travelTime" label="Travel time[min]" prepend-icon="timelapse" type="number" :rules="[rules.required]"></v-text-field>
                                </v-flex>
                                <v-flex xs12 sm3>
                                    <v-text-field v-model="flight.travelDistance" label="Travel distance [km]" prepend-icon="map" type="number" :rules="[rules.required]"></v-text-field>
                                </v-flex>
                            </v-layout>
                            <div class=".title font-weight-bold" style="text-align: left">Add layover:</div>
                            <div v-for="(layover, index) in layovers" :key="index">
                                <v-layout row-wrap justify-start align-center>
                                    <v-flex xs12 sm3>
                                        <v-text-field label="Layover" prepend-icon="map" v-model="layover.Address" :rules="[rules.required]"></v-text-field>
                                    </v-flex>
                                    <v-btn small color="error" @click="deleteRow(index)">Delete</v-btn>
                                </v-layout>
                            </div>
                            <v-layout row-wrap justify-start>
                                <v-btn small @click="addRow" style="">Add layover</v-btn>
                            </v-layout>
                            <div class=".title font-weight-bold" style="text-align: left">Price:</div>
                            <v-layout row-wrap>
                                <v-flex xs12 sm2>
                                    <v-text-field v-model="flight.firstClass" label="First class price" prepend-icon="attach_money" type="number" :rules="[rules.required]"></v-text-field>
                                </v-flex>
                                <v-flex sm2>
                                    <v-text-field v-model="flight.businessClass" label="Business class price" prepend-icon="attach_money" type="number" :rules="[rules.required]"></v-text-field>
                                </v-flex>
                                <v-flex sm2>
                                    <v-text-field v-model="flight.economyClass" label="Economy class price" prepend-icon="attach_money" type="number" :rules="[rules.required]"></v-text-field>
                                </v-flex>
                            </v-layout>
                            <v-layout row-wrap>
                                <v-flex xs12 sm2>
                                    <v-text-field v-model="flight.smallSuitcase" label="Suitcase < 21kg" prepend-icon="attach_money" type="number" :rules="[rules.required]"></v-text-field>
                                </v-flex>
                                <v-flex sm2>
                                    <v-text-field v-model="flight.bigSuitcase" label="Suitcase > 21kg" prepend-icon="attach_money" type="number" :rules="[rules.required]"></v-text-field>
                                </v-flex>
                            </v-layout>
                        </v-card-text>
                        <v-card-actions>
                            <v-btn color="primary" v-on:click.prevent="sub">Submit</v-btn>
                            <v-btn @click="reset">Reset</v-btn>
                        </v-card-actions>
                    </v-card>
                </v-flex>
            </v-layout>
        </v-container>
    </div>
        <v-snackbar v-model="SuccessSnackbar" :timeout=4000 :top="true" color="success">Successfully added flight</v-snackbar>
        <v-snackbar v-model="ErrorSnackbar" :timeout=4000 :top="true" color="error">Failed to add flight</v-snackbar>
        <v-snackbar v-model="Error2Snackbar" :timeout=4000 :top="true" color="error">All fields must be filled</v-snackbar>
    </div>
</template>

<script>
    import AirlineAdminNavbar from "../../components/AirlineAdmin/AirlineAdminNavbar";
    import axios from 'axios';
    import moment from 'moment';
    import AirlineAdminNavDrawer from "../../components/AirlineAdmin/AirlineAdminNavDrawer";

    export default {
        name: "AddFlight",
        components: {AirlineAdminNavDrawer, AirlineAdminNavbar},
        data() {
            return {
                mask: 'time',
                rules:{
                    required: value => !!value || 'Required.',
                    time: value=>{
                        const pattern = /(([0-1][0-9])|(2[0-3]))[0-5][0-9]/;
                        return pattern.test(value) || "Invalid time.[required format HH:mm]"
                    }
                },
                SuccessSnackbar: false,
                ErrorSnackbar: false,
                Error2Snackbar: false,
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
                    economyClass: "",
                    smallSuitcase: "",
                    bigSuitcase: ""
                },
                defaultFlight: {
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
                    economyClass: "",
                    smallSuitcase: "",
                    bigSuitcase: ""
                }
            }
        },
        created() {
            axios.create({withCredentials: true}).get("http://localhost:8000/api/destination/getCompanyDestinations")
                .then(res => {
                    this.destinations = res.data;
                });
            axios.create({withCredentials: true}).get("http://localhost:8000/api/airplane/getCompanyAirplanes")
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
            checkNumbers(){
                var number = parseFloat(this.flight.economyClass);
                var ind = false;
                if(number < 0 ){
                    this.flight.economyClass = "";
                    ind = true;
                }
                number = parseFloat(this.flight.businessClass);
                if(number < 0 ){
                    this.flight.businessClass = "";
                    ind = true;
                }
                number = parseFloat(this.flight.firstClass);
                if(number < 0 ){
                    this.flight.firstClass = "";
                    ind = true;
                }
                number = parseFloat(this.flight.travelDistance);
                if(number < 0 ){
                    this.flight.travelDistance = "";
                    ind = true;
                }
                number = parseFloat(this.flight.travelTime);
                if(number < 0 ){
                    this.flight.travelTime = "";
                    ind = true;
                }
                number = parseFloat(this.flight.bigSuitcase);
                if(number < 0 ){
                    this.flight.bigSuitcase = "";
                    ind = true;
                }
                number = parseFloat(this.flight.smallSuitcase);
                if(number < 0 ){
                    this.flight.smallSuitcase = "";
                    ind = true;
                }
                return ind !== true;
            },
            formatDate(date) {
                let month = `${date.getMonth() + 1}`;
                let day = `${date.getDate()}`;
                const year = date.getFullYear();

                if (month.length < 2) month = `0${month}`;
                if (day.length < 2) day = `0${day}`;

                return [year, month, day].join('-');
            },
            deleteRow(index) {
                this.layovers.splice(index, 1)
            },
            sub(e){
                e.preventDefault();
                if(!this.flight.from || !this.flight.to || !this.flight.airplane || !this.flight.departureDate || !this.flight.departureTime){
                    this.Error2Snackbar = true;
                    return;
                }
                if(!this.flight.arrivalDate || !this.flight.arrivalTime || !this.flight.travelTime || !this.flight.travelDistance){
                    this.Error2Snackbar = true;
                    return;
                }
                for(let value of this.layovers){
                    if(!value.Address){
                        this.Error2Snackbar = true;
                        return;
                    }
                }
                if(!this.flight.economyClass || !this.flight.businessClass || !this.flight.firstClass){
                    this.Error2Snackbar = true;
                    return;
                }
                if(!this.flight.smallSuitcase || !this.flight.bigSuitcase){
                    this.Error2Snackbar = true;
                    return;
                }
                const pattern = /(([0-1][0-9])|(2[0-3]))[0-5][0-9]/;
                if(!pattern.test(this.flight.arrivalTime) || !pattern.test(this.flight.departureTime)){
                    return;
                }

                if(!this.checkNumbers()){
                    return;
                }

                var dpTime = moment(this.flight.departureTime + " " + this.flight.departureDate,"HHmm YYYY-MM-DD");
                var ldTime = moment(this.flight.arrivalTime + " " + this.flight.arrivalDate, "HHmm YYYY-MM-DD");
                if (dpTime.isAfter(ldTime)){
                    alert("Arrival time is before departure time");
                    return;
                }
                var flightDto = {
                    Origin: this.flight.from,
                    Destination: this.flight.to,
                    Airplane: this.flight.airplane,
                    Departure: dpTime.toString(),
                    Landing: ldTime.toString(),
                    Duration: this.flight.travelTime,
                    Distance: this.flight.travelDistance,
                    Layovers: this.layovers,
                    BigSuitcase: this.flight.bigSuitcase,
                    SmallSuitcase: this.flight.smallSuitcase,
                    PriceECONOMY: this.flight.economyClass,
                    PriceBUSINESS: this.flight.businessClass,
                    PriceFIRSTCLASS: this.flight.firstClass
                };
                axios.create({withCredentials: true}).post("http://localhost:8000/api/flight/add",flightDto)
                    .then(res => {
                        this.SuccessSnackbar = true;
                    })
                    .catch(err => {
                        this.ErrorSnackbar = true;
                        console.log(err)});
            },
            reset(){
              this.flight = Object.assign({}, this.defaultFlight);
              this.layovers = [];
            }
        },
        watch:{
            'flight.departureDate': function(){
                var departureDate = moment(this.flight.departureDate,"YYYY-MM-DD");
                if(this.flight.arrivalDate !== ""){
                    let arrivalDate = moment(this.flight.arrivalDate,"YYYY-MM-DD");
                    if (departureDate.isAfter(arrivalDate)){
                        this.flight.arrivalDate = ""
                    }
                }
            },
        },
        computed: {
            minDateArrival(){
                if (this.flight.departureDate === ""){
                    const today = new Date();
                    return this.formatDate(today);
                }else{
                    return this.flight.departureDate;
                }
            },
            minDate() {
                const today = new Date();
                return this.formatDate(today);
            },
            fromDestinations() {
                return this.destinations.filter((destination) => destination.Name !== this.flight.to);
            },

            toDestinations() {
                return this.destinations.filter((destination) => destination.Name !== this.flight.from);
            }
        }
    }
</script>

<style scoped>
    #main {
        background-image: linear-gradient(to right bottom, #142eae, #005bca, #007ed2, #009ccd, #0bb7c7, #47c0c6, #67c8c6, #81d0c7, #6ecac4, #58c4c3, #3cbdc2, #00b7c1);
    }
    .flexcard {
        display: flex;
        flex-direction: column;
    }
</style>