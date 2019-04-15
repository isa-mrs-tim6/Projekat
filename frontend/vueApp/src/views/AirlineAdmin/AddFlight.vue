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
                        <v-select v-model="flight.from" label="Name" :items="fromDestinations" item-text="Name"
                                  item-value="Name" :options="fromDestinations" prepend-icon="map"></v-select>
                    </v-flex>
                    <v-flex xs12 sm3>
                        <v-select v-model="flight.to" label="Name" :items="toDestinations" item-text="Name"
                                  item-value="Name" :options="toDestinations" prepend-icon="map"></v-select>
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
                <h3>Arrival:</h3>
                <v-layout row-wrap>
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
                <h3>Travel information: </h3>
                <v-layout row-wrap>
                    <v-flex xs12 sm3>
                        <v-text-field v-model="flight.travelTime" label="Travel time[min]" prepend-icon="timelapse" type="number" :rules="[rules.required]"></v-text-field>
                    </v-flex>
                    <v-flex xs12 sm3>
                        <v-text-field v-model="flight.travelDistance" label="Travel distance [km]" prepend-icon="map" type="number" :rules="[rules.required]"></v-text-field>
                    </v-flex>
                </v-layout>
                <h3>Add layover:</h3>
                <div v-for="(layover, index) in layovers" :key="index">
                    <v-layout row-wrap>
                        <v-flex xs12 sm3>
                            <v-text-field label="Layover" prepend-icon="map" v-model="layover.Address" :rules="[rules.required]"></v-text-field>
                        </v-flex>
                        <v-btn small color="error" @click="deleteRow(index)">Delete</v-btn>
                    </v-layout>
                </div>
                <v-btn small @click="addRow">Add layover</v-btn>
                <h3>Price:</h3>
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
                rules:{
                    required: value => !!value || 'Required.',
                    time: value=>{
                        const pattern = /(([0-1][0-9])|(2[0-3]))[0-5][0-9]/;
                        return pattern.test(value) || "Invalid time.[required format HH:mm]"
                    }
                },
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
                }
            }
        },
        created() {
            axios.create({withCredentials: true}).get("http://localhost:8000/api/destination/getCompanysDestinations")
                .then(res => {
                    this.destinations = res.data;
                });
            axios.create({withCredentials: true}).get("http://localhost:8000/api/airplane/getCompanysAirplanes")
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
                    return;
                }
                if(!this.flight.arrivalDate || !this.flight.arrivalTime || !this.flight.travelTime || !this.flight.travelDistance){
                    return;
                }
                for(let value of this.layovers){
                    if(!value.Address){
                        return;
                    }
                }
                if(!this.flight.economyClass || !this.flight.businessClass || !this.flight.firstClass){
                    return;
                }
                if(!this.flight.smallSuitcase || !this.flight.bigSuitcase){
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
                    .then(res => alert("successfully added flight"))
                    .catch(err => console.log(err));
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

</style>