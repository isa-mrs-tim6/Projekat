<template>
    <v-card dark>
        <div>
            <v-layout row-wrap pt-3 ml-3>
                <v-flex xs1><span class=".body-2" :class='this.searchClassBinder("one-way")' @click='selectedSearchType = "one-way"'>One-way</span></v-flex>
                <v-flex xs1><span class=".body-2" :class='this.searchClassBinder("round")' @click='selectedSearchType = "round"'>Round-Trip</span></v-flex>
                <v-flex xs1><span class=".body-2" :class='this.searchClassBinder("multi")' @click='selectedSearchType = "multi"'>Multi-city</span></v-flex>
            </v-layout>
            <div v-if='selectedSearchType === "one-way"'>
                <v-layout row-wrap mx-3>
                    <v-flex xs3 mr-2>
                        <v-text-field label="From" prepend-icon="map" v-model="oneWay.from"></v-text-field>
                    </v-flex>
                    <v-flex xs>
                        <v-btn flat icon color="grey" large @click="swapDestinations(oneWay)">
                            <v-icon>compare_arrows</v-icon>
                        </v-btn>
                    </v-flex>
                    <v-flex xs3 mr-2>
                        <v-text-field label="To" prepend-icon="map" v-model="oneWay.to"></v-text-field>
                    </v-flex>
                    <v-flex xs3 ml-2>
                        <v-menu v-model="oneWay.menu" :close-on-content-click="false" lazy transition="scale-transition"
                                offset-y full-width max-width="290px" min-width="290px">
                            <template v-slot:activator="{ on }">
                                <v-text-field label="Departure date" prepend-icon="event" readonly v-on="on" v-model="oneWay.date"></v-text-field>
                            </template>
                            <v-date-picker v-model="oneWay.date" no-title scrollable @input="oneWay.menu = false"></v-date-picker>
                        </v-menu>
                    </v-flex>
                    <v-flex xs3 ml-2>
                        <v-menu v-model="oneWay.classMenu" :close-on-content-click="false" lazy transition="scale-transition"
                                offset-y max-width="200px" max-height="350px">
                            <template v-slot:activator="{on}" >
                                <v-text-field label="Passangers" readonly v-on="on" v-model="classAndPassengersDisplay"></v-text-field>
                            </template>
                            <v-card>
                                <v-card-title><span class=".title">Cabin class</span>
                                    <v-layout column justify-center align-center>
                                        <v-flex xs4>
                                            <span class=".body-2 plusA" pa-1 :class='this.cabinClassBinder("first", oneWay)' @click='oneWay.seatClass = "first"'>First</span>
                                        </v-flex>
                                        <v-flex xs4>
                                            <span class=".body-2 plusA" pa-1 :class='this.cabinClassBinder("business", oneWay)' @click='oneWay.seatClass = "business"'>Business</span>
                                        </v-flex>
                                        <v-flex xs4>
                                            <span class=".body-2 plusA" pa-1 :class='this.cabinClassBinder("economic", oneWay)' @click='oneWay.seatClass = "economic"'>Economic</span>
                                        </v-flex>
                                    </v-layout>
                                </v-card-title>
                            </v-card>
                            <v-divider></v-divider>
                            <v-card>
                                <v-card-title>
                                    <v-layout column>
                                        <span class=".title">Travelers</span>
                                        <v-layout pt-2 row wrap justify-center>
                                            <v-flex xs1 ma-2 d-inline-block>
                                            <span class="passengerSign" @click="--oneWay.passengers" v-if="this.oneWay.passengers > 1">
                                                -
                                            </span>
                                                <span v-else class="passengerSignDisabled">
                                                -
                                            </span>
                                            </v-flex>
                                            <v-flex xs1 ma-2 d-inline-block justify-center align-center><div>{{this.oneWay.passengers}}</div></v-flex>
                                            <v-flex xs1 ma-2 d-inline-block>
                                            <span class="passengerSign" @click="++oneWay.passengers" v-if="this.oneWay.passengers < 99">
                                                +
                                            </span>
                                                <span v-else class="passengerSignDisabled">
                                                +
                                            </span>
                                            </v-flex>
                                        </v-layout>
                                    </v-layout>
                                </v-card-title>
                            </v-card>
                        </v-menu>
                    </v-flex>
                    <v-flex xs3><v-btn @click="oneWaySearch">Search</v-btn></v-flex>
                </v-layout>
            </div>
            <div v-else-if='selectedSearchType === "round"'>
                <v-layout row-wrap mx-3>
                    <v-flex xs3 mr-2>
                        <v-text-field label="From" v-model="round.from" prepend-icon="map"></v-text-field>
                    </v-flex>
                    <v-flex xs>
                        <v-btn flat icon color="grey" large @click="swapDestinations(round)">
                            <v-icon>compare_arrows</v-icon>
                        </v-btn>
                    </v-flex>
                    <v-flex xs3 mr-2>
                        <v-text-field label="To" v-model="round.to" prepend-icon="map"></v-text-field>
                    </v-flex>
                    <v-flex xs3 ml-2>
                        <v-menu v-model="round.departureMenu" :close-on-content-click="false" lazy transition="scale-transition"
                                offset-y full-width max-width="290px" min-width="290px">
                            <template v-slot:activator="{ on }">
                                <v-text-field label="Departure date" prepend-icon="event" readonly v-on="on" v-model="round.departureDate"></v-text-field>
                            </template>
                            <v-date-picker v-model="round.departureDate" no-title scrollable @input="round.departureMenu = false"></v-date-picker>
                        </v-menu>
                    </v-flex>
                    <v-flex xs3 ml-2>
                        <v-menu v-model="round.returnMenu" :close-on-content-click="false" lazy transition="scale-transition"
                                offset-y full-width max-width="290px" min-width="290px">
                            <template v-slot:activator="{ on }">
                                <v-text-field label="Return date" prepend-icon="event" readonly v-on="on" v-model="round.returnDate"></v-text-field>
                            </template>
                            <v-date-picker v-model="round.returnDate" no-title scrollable @input="round.returnMenu = false"></v-date-picker>
                        </v-menu>
                    </v-flex>
                    <v-flex xs3 ml-2>
                        <v-menu v-model="round.classMenu" :close-on-content-click="false" lazy transition="scale-transition"
                                offset-y max-width="200px" max-height="350px">
                            <template v-slot:activator="{on}" >
                                <v-text-field label="Passangers" readonly v-on="on" v-model="classAndPassengersDisplay"></v-text-field>
                            </template>
                            <v-card>
                                <v-card-title><span class=".title">Cabin class</span>
                                    <v-layout column justify-center align-center>
                                        <v-flex xs4>
                                            <span class=".body-2 plusA" pa-1 :class='this.cabinClassBinder("first",round)' @click='round.seatClass = "first"'>First</span>
                                        </v-flex>
                                        <v-flex xs4>
                                            <span class=".body-2 plusA" pa-1 :class='this.cabinClassBinder("business", round)' @click='round.seatClass = "business"'>Business</span>
                                        </v-flex>
                                        <v-flex xs4>
                                            <span class=".body-2 plusA" pa-1 :class='this.cabinClassBinder("economic", round)' @click='round.seatClass = "economic"'>Economic</span>
                                        </v-flex>
                                    </v-layout>
                                </v-card-title>
                            </v-card>
                            <v-divider></v-divider>
                            <v-card>
                                <v-card-title>
                                    <v-layout column>
                                        <span class=".title">Travelers</span>
                                        <v-layout pt-2 row wrap justify-center>
                                            <v-flex xs1 ma-2 d-inline-block>
                                            <span class="passengerSign" @click="--round.passengers" v-if="this.round.passengers > 1">
                                                -
                                            </span>
                                                <span v-else class="passengerSignDisabled">
                                                -
                                            </span>
                                            </v-flex>
                                            <v-flex xs1 ma-2 d-inline-block justify-center align-center><div>{{this.round.passengers}}</div></v-flex>
                                            <v-flex xs1 ma-2 d-inline-block>
                                            <span class="passengerSign" @click="++round.passengers" v-if="this.round.passengers < 99">
                                                +
                                            </span>
                                                <span v-else class="passengerSignDisabled">
                                                +
                                            </span>
                                            </v-flex>
                                        </v-layout>
                                    </v-layout>
                                </v-card-title>
                            </v-card>
                        </v-menu>
                    </v-flex>
                    <v-flex xs3><v-btn @click="roundSearch">Search</v-btn></v-flex>
                </v-layout>
            </div>
            <div v-else-if='selectedSearchType === "multi"'>
                <div v-for="(layover, index) in multi.layovers" :key="index">
                    <v-layout row-wrap justify-start align-center mx-3>
                        <v-flex xs4 mr-3>
                            <v-text-field label="From" prepend-icon="map" v-model="layover.from"></v-text-field>
                        </v-flex>
                        <v-flex xs4>
                            <v-text-field label="To" prepend-icon="map" v-model="layover.to"></v-text-field>
                        </v-flex>
                    </v-layout>
                </div>
                <v-layout row-wrap justify-start mx-3>
                    <v-flex xs2 pt-2>
                        <v-btn small @click="addRow" v-if="this.multi.layovers.length < 4">Add</v-btn>
                    </v-flex>
                    <v-flex xs2 mr-3 pt-2>
                        <v-btn small @click="deleteRow">Remove</v-btn>
                    </v-flex>
                    <v-flex xs2 mr-3>
                        <v-menu v-model="multi.departureMenu" :close-on-content-click="false" lazy transition="scale-transition"
                                offset-y full-width max-width="290px" min-width="290px">
                            <template v-slot:activator="{ on }">
                                <v-text-field label="Departure date" prepend-icon="event" readonly v-on="on" v-model="multi.departureDate" ></v-text-field>
                            </template>
                            <v-date-picker v-model="multi.departureDate" no-title scrollable @input="multi.departureMenu = false"></v-date-picker>
                        </v-menu>
                    </v-flex>
                    <v-flex xs2>
                        <v-menu v-model="multi.classMenu" :close-on-content-click="false" lazy transition="scale-transition"
                                offset-y max-width="200px" max-height="350px">
                            <template v-slot:activator="{on}" >
                                <v-text-field label="Passangers" readonly v-on="on" v-model="classAndPassengersDisplay"></v-text-field>
                            </template>
                            <v-card>
                                <v-card-title><span class=".title">Cabin class</span>
                                    <v-layout column justify-center align-center>
                                        <v-flex xs4>
                                            <span class=".body-2 plusA" pa-1 :class='this.cabinClassBinder("first",multi)' @click='multi.seatClass = "first"'>First</span>
                                        </v-flex>
                                        <v-flex xs4>
                                            <span class=".body-2 plusA" pa-1 :class='this.cabinClassBinder("business", multi)' @click='multi.seatClass = "business"'>Business</span>
                                        </v-flex>
                                        <v-flex xs4>
                                            <span class=".body-2 plusA" pa-1 :class='this.cabinClassBinder("economic", multi)' @click='multi.seatClass = "economic"'>Economic</span>
                                        </v-flex>
                                    </v-layout>
                                </v-card-title>
                            </v-card>
                            <v-divider></v-divider>
                            <v-card>
                                <v-card-title>
                                    <v-layout column>
                                        <span class=".title">Travelers</span>
                                        <v-layout pt-2 row wrap justify-center>
                                            <v-flex xs1 ma-2 d-inline-block>
                                            <span class="passengerSign" @click="--multi.passengers" v-if="this.multi.passengers > 1">
                                                -
                                            </span>
                                                <span v-else class="passengerSignDisabled">
                                                -
                                            </span>
                                            </v-flex>
                                            <v-flex xs1 ma-2 d-inline-block justify-center align-center><div>{{this.multi.passengers}}</div></v-flex>
                                            <v-flex xs1 ma-2 d-inline-block>
                                            <span class="passengerSign" @click="++multi.passengers" v-if="this.multi.passengers < 99">
                                                +
                                            </span>
                                                <span v-else class="passengerSignDisabled">
                                                +
                                            </span>
                                            </v-flex>
                                        </v-layout>
                                    </v-layout>
                                </v-card-title>
                            </v-card>
                        </v-menu>
                    </v-flex>
                    <v-flex xs><v-btn @click="multiSearch">Search</v-btn></v-flex>
                </v-layout>
            </div>
            <v-snackbar v-model="Error2Snackbar" :timeout=4000 :top="true" color="error">All fields must be filled</v-snackbar>
        </div>
    </v-card>
</template>

<script>
    export default {
        name: "FlightSearch",
        data(){
            return{
                Error2Snackbar: false,
                selectedSearchType: "one-way",
                selectedCabinClass: "economic",
                oneWay: {
                    menu: false,
                    classMenu: false,
                    from: "",
                    to: "",
                    date: "",
                    passengers: 1,
                    seatClass: "economic",
                },
                round: {
                    classMenu: false,
                    from: "",
                    to: "",
                    date: "",
                    departureDate: "",
                    departureMenu: false,
                    returnDate: "",
                    returnMenu: false,
                    passengers: 1,
                    seatClass: "economic",
                },
                multi: {
                    layovers:[
                        {
                            from: "",
                            to: ""
                        },
                        {
                            from: "",
                            to: ""
                        }
                    ],
                    departureMenu: false,
                    passengers: 1,
                    seatClass: "economic",
                    departureDate: ""
                }

            }
        },
        methods: {
            swapDestinations(searchType) {
                let dest = searchType.from;
                searchType.from = searchType.to;
                searchType.to = dest;
            },
            searchClassBinder(searchType) {
                var retVal = "type";
                if (searchType === this.selectedSearchType) {
                    retVal += " selectedType";
                }
                return retVal;
            },
            cabinClassBinder(cabinClass, searchType) {
                var retVal = "cabinClass";
                if (cabinClass === searchType.seatClass) {
                    retVal += " selectedCabinClass";
                }
                return retVal;
            },
            passengerSignClassBinder(sign, passengerNum) {
                var retVal = "";
                if (sign === "plus") {
                    if (passengerNum === 99) {
                        retVal = "passengerSignDisabled";
                    } else {
                        retVal = "passengerSign"
                    }
                } else {
                    if (passengerNum === 0) {
                        retVal = "passengerSignDisabled";
                    } else {
                        retVal = "passengerSign"
                    }
                }
                return retVal;
            },
            addRow() {
                if(this.multi.layovers.length > 4){
                    return;
                }
                this.multi.layovers.push({
                    from: "",
                    to: ""
                })
            },
            deleteRow() {
                if(this.multi.layovers.length === 2){
                    this.selectedSearchType = "round";
                }else{
                    this.multi.layovers.splice(-1, 1)
                }
            },
            oneWaySearch(){
                if(!this.oneWay.from || !this.oneWay.to || !this.oneWay.date){
                    this.Error2Snackbar = true;
                    return;
                }

                const query = {
                    from: this.oneWay.from,
                    to: this.oneWay.to,
                    date: this.oneWay.date,
                    seatClass: this.oneWay.seatClass,
                    passengers: this.oneWay.passengers
                };
                console.log(query);
                this.$emit('search', query);
            },
            roundSearch(){
                if(!this.round.from || !this.round.to || !this.round.departureDate || !this.round.returnDate){
                    this.Error2Snackbar = true;
                    return;
                }

                const query = {
                    from: this.round.from,
                    to: this.round.to,
                    departureDate: this.round.departureDate,
                    returnDate: this.round.returnDate,
                    seatClass: this.round.seatClass,
                    passengers: this.round.passengers
                };
                this.$emit('search', query);
            },
            multiSearch(){
                if(!this.multi.layovers[0].from || !this.multi.layovers[this.multi.layovers.length - 1].to){
                    this.Error2Snackbar = true;
                    return;
                }
                if(!this.multi.layovers[0].to || !this.multi.layovers[this.multi.layovers.length - 1].from){
                    this.Error2Snackbar = true;
                    return;
                }
                if(!this.multi.departureDate){
                    this.Error2Snackbar = true;
                    return;
                }
                var fromDestination = this.multi.layovers[0].from;
                var toDestination = this.multi.layovers[this.multi.layovers.length - 1].to;
                var layovers = this.multi.layovers[0].to + ";" + this.multi.layovers[this.multi.layovers.length - 1].from;
                for (var i = 1; i < this.multi.layovers.length - 1; i++){
                    if(!this.multi.layovers[i].from || !this.multi.layovers[i].to){
                        this.Error2Snackbar = true;
                        return;
                    }
                    layovers += ";" + this.multi.layovers[i].from + ";";
                    layovers += this.multi.layovers[i].to;
                }

                const query = {
                    from: fromDestination,
                    to: toDestination,
                    layovers: layovers,
                    departureDate: this.multi.departureDate,
                    passengers: this.multi.passengers,
                    seatClass: this.multi.seatClass
                };
                this.$emit('search', query);
            },
        },
        computed:{
            classAndPassengersDisplay: function(){
                var retVal = "";
                if(this.selectedSearchType === "one-way"){
                    if (this.oneWay.passengers === 1){
                        retVal += "1 Passenger, ";
                    }else{
                        retVal += this.oneWay.passengers.toString() + " Passengers, ";
                    }
                    retVal += this.oneWay.seatClass.toUpperCase();
                }else if(this.selectedSearchType === "round"){
                    if (this.round.passengers === 1){
                        retVal += "1 Passenger, ";
                    }else{
                        retVal += this.round.passengers.toString() + " Passengers, ";
                    }
                    retVal += this.round.seatClass.toUpperCase();
                }else if(this.selectedSearchType === "multi"){
                    if (this.multi.passengers === 1){
                        retVal += "1 Passenger, ";
                    }else{
                        retVal += this.multi.passengers.toString() + " Passengers, ";
                    }
                    retVal += this.multi.seatClass.toUpperCase();
                }

                return retVal;
            },
        },
    }
</script>

<style scoped>
    .passengerSign{
        position:relative;
        border: 1px solid #0ca7c6;
        height: 20px;
        width: 20px;
        font-size: 16px;
        line-height: 20px;
        text-align: center;
        color: #0ca7c6;
        display: inline-block;
        cursor: pointer;

    }
    .passengerSignDisabled{
        height: 20px;
        width: 20px;
        font-size: 16px;
        line-height: 20px;
        text-align: center;
        background-color: grey;
        display: inline-block;
        color: white;
        cursor: not-allowed;

    }
    .passengerSign:hover {
        color: white;
        background-color: #0ca7c6;
    }
    .passengerNumber{
        line-height: 20px;
        font-size: 16px;
        text-align: center;
        display: inline-block;
        margin-right: 8px;
        margin-left: 8px;
    }
    .cabinClass{
        color: #6d8494;
        cursor: pointer;
        font-size: 14px;
        width : 100px;
        height: 40px;
        margin: 2px;
        display: block;
        line-height: 40px;
        text-align: center;
    }
    .cabinClass:hover{
        color: #0ca7c6;
        border: 3px #00baf7 solid;
    }
    .selectedCabinClass{
        color: #0ca7c6;
        border: 2px #00baf7 solid;
        width : 100px;
        height: 40px;
    }
    .type{
        color: #6d8494;
        font-size: 14px;
        font-weight: bold;
        cursor: pointer;
    }
    .selectedType{
        color: aliceblue;
    }
</style>