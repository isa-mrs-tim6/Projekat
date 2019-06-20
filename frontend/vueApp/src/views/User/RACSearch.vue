<template>
    <div>
        <v-card dark>
            <v-card-title>
                <h2>Search Rent-a-car Companies:</h2>
            </v-card-title>
            <v-card-text>
                <v-container>
                    <v-layout align-center justify-center row-wrap>
                        <v-flex xs3>
                            <v-text-field label="Name" prepend-icon="directions_car" v-model="query.name"></v-text-field>
                        </v-flex>
                        <v-flex xs3>
                            <v-text-field label="Address" prepend-icon="place" v-model="query.address" :readonly="checkReadonly()"></v-text-field>
                        </v-flex>
                        <v-flex xs2>
                            <v-menu v-model="menuFrom" :close-on-content-click="false" lazy transition="scale-transition"
                                    :nudge-right="40" offset-y full-width max-width="290px" min-width="290px">
                                <template v-slot:activator="{ on }">
                                    <v-text-field v-model="time.from" label="From" prepend-icon="event" readonly v-on="on"></v-text-field>
                                </template>
                                <v-date-picker no-title scrollable v-model = "time.from" @input="menuFrom = false" :readonly="checkReadonly()"></v-date-picker>
                            </v-menu>
                        </v-flex>
                        <v-flex xs2>
                            <template v-if = "returnDate === null">
                                <v-menu v-model="menuTo" :close-on-content-click="false" lazy transition="scale-transition"
                                        :nudge-right="40" offset-y full-width max-width="290px" min-width="290px">
                                    <template v-slot:activator="{ on }">
                                        <v-text-field v-model="time.to" label="To" prepend-icon="event" readonly v-on="on"></v-text-field>
                                    </template>
                                    <v-date-picker no-title scrollable v-model = "time.to" @input="menuTo = false" :min="minDate"></v-date-picker>
                                </v-menu>
                            </template>
                            <template v-else>
                                <v-text-field v-model="time.to" label="To" prepend-icon="event" readonly></v-text-field>
                            </template>
                        </v-flex>
                        <v-flex xs2>
                            <v-btn @click="racSearch">Search</v-btn>
                        </v-flex>
                    </v-layout>
                </v-container>
            </v-card-text>
        </v-card>
        <v-layout justify-center>
            <v-flex xs12>
                <v-container pa-0 mt-2 grid-list-md fluid>
                    <v-layout row wrap v-for="(item, index) in items" :key="item.ID"
                              xs4>
                        <v-flex>
                            <v-card height="100%" dark>
                                <v-card-text>
                                    <v-layout row>
                                        <v-flex xs2>
                                            <v-img src="http://pngriver.com/wp-content/uploads/2017/11/city-buildings-png-transparent-images-clipart-icons-pngriver-download-free-skyscraper-1-featured-600x450.png"></v-img>
                                        </v-flex>
                                        <v-flex xs8>
                                            <v-layout row >
                                                <div class="headline">
                                                    {{item.Name}}
                                                </div>
                                            </v-layout>
                                            <v-layout row>
                                                <v-icon>place</v-icon>
                                                {{item.Address}}
                                            </v-layout>
                                            <v-layout row class="center">
                                                {{item.Promo}}
                                            </v-layout>
                                            <v-layout row>
                                                <v-select
                                                        v-model = "selectedLoc[index]"
                                                        :items=" item.Locations"
                                                        item-text="Address"
                                                        item-value="ID"
                                                        label="Locations"
                                                ></v-select>
                                            </v-layout>
                                        </v-flex>
                                        <v-flex xs2 class="center" style="text-align: right">
                                            <v-btn @click="searchVehicles(index)">
                                                vehicles
                                            </v-btn>
                                            <br/>
                                            <v-btn @click="quickRes(index)">
                                                quickres
                                            </v-btn>
                                        </v-flex>
                                    </v-layout>
                                </v-card-text>
                            </v-card>
                        </v-flex>
                    </v-layout>
                </v-container>
            </v-flex>
        </v-layout>
        <v-snackbar v-model="ErrorSnackbar" :timeout=2000 :top="true" color="error">{{this.ErrorSnackbarText}}</v-snackbar>
    </div>

</template>

<script>
    import axios from 'axios';
    import moment from 'moment';

    export default {
        name: "RACSearch",
        props: ["reservationID", "passengers"],
        data() {
            return {
                selectedLoc: [],
                ErrorSnackbar: false,
                ErrorSnackbarText: '',
                items: [],
                menuTo: false,
                menuFrom: false,
                query: {
                    name: null,
                    address: null,
                },
                time: {
                    from: null,
                    to: null,
                },
                returnDate: null,
                minDate:"",
            }
        },
        created(){
            this.query.address = localStorage.getItem("destination");
            this.time.to = localStorage.getItem("returnDate");
            this.returnDate = localStorage.getItem("returnDate");
            this.time.from = localStorage.getItem("departureDate");
            this.minDate = this.time.from;
        },
        methods: {
            checkReadonly(){
                return !isNaN(parseInt(this.reservationID))
            },
            quickRes(index){
                if(isNaN(parseInt(this.selectedLoc[index]))){
                    this.ErrorSnackbar = true;
                    this.ErrorSnackbarText = 'Please select a location.';
                }else{
                    this.$router.push({ path: `/vehiclesQuick/${this.items[index].ID}/${this.selectedLoc[index]}`,
                        query: { reservationID: this.reservationID, passengers: this.passengers, start: this.time.from, end: this.time.to}})
                }
            },
            searchVehicles(index){
                if(isNaN(parseInt(this.selectedLoc[index]))){
                    this.ErrorSnackbar = true;
                    this.ErrorSnackbarText = 'Please select a location.';
                }else{
                    this.$router.push({ path: `/vehiclesSearch/${this.items[index].ID}/${this.selectedLoc[index]}`, query: { reservationID: this.reservationID, passengers: this.passengers, start: this.time.from, end: this.time.to}})
                }
            },
            racSearch(e) {
                e.preventDefault();
                const SearchQuery = {
                    "Name": this.query.name,
                    "Address": this.query.address,
                    "From": moment(this.time.from,"YYYY-MM-DD").valueOf().toString(),
                    "To": moment(this.time.to,"YYYY-MM-DD").valueOf().toString(),
                };
                axios.create({withCredentials: true}).post("http://localhost:8000/api/search/rac", SearchQuery).
                    then(res => {
                        this.items = res.data;

                        for(let i=0; i<this.items.length; i++){
                            this.selectedLoc.push('');
                        }
                })
            }
        }
    }
</script>

<style scoped>
    .center {
        margin: auto 0;
        padding: 10px;
        font-size: 16px;
        word-wrap: break-word;
    }
</style>