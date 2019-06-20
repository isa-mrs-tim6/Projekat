<template>
    <div>
        <UserNavBar></UserNavBar>
        <v-layout justify-center style="margin-top: 200px">
            <v-flex xs9 style="height: 100%;">
                <div v-bind:isLogIn="isLogIn"><flight-search v-on:search="search"></flight-search></div>
                <v-layout row>
                    <v-flex>
                        <v-card dark v-if="flights.length !== 0">
                            <v-card-text>
                                <v-layout row>
                                    <h3>Filters:</h3>
                                </v-layout>
                                <v-layout row>
                                    <v-flex>
                                        <v-combobox label="Airlines" v-model = "filterParam.Airlines" :items="airlinesComputed" multiple chips>
                                            <template v-slot:selection="data">
                                                <v-chip
                                                        :key="JSON.stringify(data.item)"
                                                        :selected="data.selected"
                                                        :disabled="data.disabled"
                                                        class="v-chip--select-multi"
                                                        @input="data.parent.selectItem(data.item)"
                                                >
                                                    <v-avatar
                                                            class="accent white--text"
                                                            v-text="data.item.slice(0, 1).toUpperCase()"
                                                    ></v-avatar>
                                                    {{ data.item }}
                                                </v-chip>
                                            </template>
                                        </v-combobox>
                                    </v-flex>
                                </v-layout>
                                <v-layout row>
                                    <v-flex>
                                        <v-layout row>
                                            <v-flex xs1><v-subheader style="padding-top: 12px">Price: </v-subheader></v-flex>
                                            <v-flex shrink style="width: 60px">
                                                <v-text-field v-model="filterParam.Price[0]" class="mt-0" hide-details single-line type="number"></v-text-field>
                                            </v-flex>
                                            <v-flex class="px-3">
                                                <v-range-slider v-model="filterParam.Price" :max=maxPrice :min="0" :step="1" value="Price range"></v-range-slider>
                                            </v-flex>
                                            <v-flex shrink style="width: 60px">
                                                <v-text-field v-model="filterParam.Price[1]" class="mt-0" hide-details single-line type="number"></v-text-field>
                                            </v-flex>
                                        </v-layout>
                                    </v-flex>
                                </v-layout>
                                <v-layout row>
                                    <v-flex>
                                        <v-layout row>
                                            <v-flex xs1><v-subheader style="padding-top: 12px">Duration[min]: </v-subheader></v-flex>
                                            <v-flex shrink style="width: 60px">
                                                <v-text-field v-model="filterParam.Duration[0]" class="mt-0" hide-details single-line type="number"></v-text-field>
                                            </v-flex>
                                            <v-flex class="px-3">
                                                <v-range-slider v-model="filterParam.Duration" :max=maxDuration :min="0" :step="1" value="Price range"></v-range-slider>
                                            </v-flex>
                                            <v-flex shrink style="width: 60px">
                                                <v-text-field v-model="filterParam.Duration[1]" class="mt-0" hide-details single-line type="number"></v-text-field>
                                            </v-flex>
                                        </v-layout>
                                    </v-flex>
                                </v-layout>
                            </v-card-text>
                        </v-card>
                    </v-flex>
                </v-layout>
                <FlightGrid v-bind:flights="flightsToShow" v-bind:searchType="searchType"></FlightGrid>
            </v-flex>
        </v-layout>
            <v-snackbar v-model="ResultsSnackbar" :timeout=4000 :top="true" color="info">{{ResultsSnackbarMessage}}</v-snackbar>            
    </div>
</template>

<script>
    import FlightGrid from "../../components/User/FlightGrid";
    import UserNavBar from "../../components/User/UserNavBar";
    import FlightSearch from "../../components/User/FlightSearch";
    import axios from "axios";
    import moment from 'moment';

    export default {
        name: "Reserve",
        components: {FlightSearch, UserNavBar, FlightGrid},
        data() {
            return {
                flights: [],
                ResultsSnackbarMessage:"",
                ResultsSnackbar: false,
                searchType: null,
                isLogIn: false,
                filterParam:{
                    Price: [0, 0],
                    Duration: [0, 0],
                    seatClass: "",
                    Airlines: ""
                }
            }
        },
        methods: {
            search: function(query) {
                this.filterParam.seatClass = query.seatClass;
                if(query.returnDate){
                    this.searchType = "round";
                    var searchQuery = {
                        from : query.from,
                        to : query.to,
                        date: query.departureDate,
                        seatClass: query.seatClass,
                        passengers: query.passengers.toString()
                    };
                    var dateMoment = moment(searchQuery.date,"YYYY-MM-DD");
                    searchQuery.date = dateMoment.valueOf().toString();
                    axios.post("http://localhost:8000/api/search/oneWay",searchQuery)
                        .then(res=>{
                            this.flights = res.data;
                            if(this.flights.length === 0){
                                this.ResultsSnackbarMessage = "No flights found to "+ query.to + " !"
                                this.ResultsSnackbar = true;
                            }
                        });
                    var searchQuery2 = {
                        from : query.to,
                        to : query.from,
                        date: query.returnDate,
                        seatClass: query.seatClass,
                        passengers: query.passengers.toString()
                    };
                    localStorage.setItem('departureDate',query.departureDate);
                    localStorage.setItem('returnDate',query.returnDate);

                    var dateMoment2 = moment(searchQuery2.date,"YYYY-MM-DD");
                    searchQuery2.date = dateMoment2.valueOf().toString();
                    axios.post("http://localhost:8000/api/search/oneWay", searchQuery2)
                        .then(res=>{
                            this.flightsReturn = res.data;
                            if(this.flightsReturn.length === 0){
                                this.ResultsSnackbarMessage = "No return flights found!"
                                this.ResultsSnackbar = true;
                            }
                        });
                }else if(query.date){
                    this.searchType = "oneWay";
                    var searchQuery = {
                        from : query.from,
                        to : query.to,
                        date: query.date,
                        seatClass: query.seatClass,
                        passengers: query.passengers.toString()
                    };
                    var dateMoment = moment(searchQuery.date,"YYYY-MM-DD");
                    searchQuery.date = dateMoment.valueOf().toString();
                    axios.post("http://localhost:8000/api/search/oneWay", searchQuery)
                        .then(res=>{
                            this.flights = res.data;
                            if(this.flights.length === 0){
                                this.ResultsSnackbarMessage = "No flights found!"
                                this.ResultsSnackbar = true;
                            }
                        })
                    localStorage.setItem('departureDate',query.date);
                }else{
                    this.searchType = "oneWay";
                    var searchQuery = {
                        from : query.from,
                        to : query.to,
                        layovers: query.layovers,
                        date: query.departureDate,
                        seatClass: query.seatClass,
                        passengers: query.passengers.toString()
                    };
                    var dateMoment = moment(searchQuery.date,"YYYY-MM-DD");
                    searchQuery.date = dateMoment.valueOf().toString();
                    axios.post("http://localhost:8000/api/search/multi", searchQuery)
                        .then(res=>{
                            this.flights = res.data;
                            if(this.flights.length === 0){
                                this.ResultsSnackbarMessage = "No flights found!"
                                this.ResultsSnackbar = true;
                            }
                        });
                    localStorage.setItem('departureDate',query.date);
                }
            }
        },
        watch:{
            maxPrice: function(val){
                this.filterParam.Price[1] = val;
            },
            maxDuration: function(val){
                this.filterParam.Duration[1] = val;
            }
        },
        computed:{
            airlinesComputed: function(){
                var retVal = [];
                this.flights.forEach(flight =>{retVal.push(flight.Airline.Name)});
                return retVal;
            },
            maxPrice: function(){
                let max = 0 ;
                if(this.flights.length !== 0) {
                    if (this.filterParam.seatClass === "economic") {
                        max = this.flights[0].PriceECONOMY;
                    } else if (this.filterParam.seatClass === "business") {
                        max = this.flights[0].PriceBUSINESS;
                    } else {
                        max = this.flights[0].PriceFIRSTCLASS;
                    }
                    for (let i = 0, len = this.flights.length; i < len; i++) {
                        let v;
                        if (this.filterParam.seatClass === "economic") {
                            v = this.flights[i].PriceECONOMY;
                        } else if (this.filterParam.seatClass === "business") {
                            v = this.flights[i].PriceBUSINESS;
                        } else {
                            v = this.flights[i].PriceFIRSTCLASS;
                        }
                        max = (v > max) ? v : max;
                    }
                }
                return max;
            },
            maxDuration: function(){
                let max = 0;
                if(this.flights.length !== 0) {
                    max = this.flights[0].Duration / 60000000000;
                    for (let i = 0, len = this.flights.length; i < len; i++) {
                        let v = this.flights[i].Duration / 60000000000;
                        max = (v > max) ? v : max;
                    }
                }
                return max;
            },
            flightsToShow: function(){
                let flights = this.flights;
                if(this.filterParam.Price[1] !== this.maxPrice){
                    if(this.filterParam.seatClass === "economic"){
                        flights = flights.filter(flight => flight.PriceECONOMY < this.filterParam.Price[1]);
                    }else if(this.filterParam.seatClass === "business"){
                        flights = flights.filter(flight => flight.PriceBUSINESS < this.filterParam.Price[1]);
                    }else{
                        flights = flights.filter(flight => flight.PriceFIRSTCLASS < this.filterParam.Price[1]);
                    }
                }
                if(this.filterParam.Duration[1] !== this.maxDuration){
                    flights = flights.filter(flight => (flight.Duration / 60000000000) < this.filterParam.Duration[1]);
                }
                if(this.filterParam.Airlines.length > 0){
                    flights = flights.filter(flight => (this.filterParam.Airlines.indexOf(flight.Airline.Name)) > -1);
                }
                return flights;
            }
        }

    }
</script>

<style scoped>

</style>