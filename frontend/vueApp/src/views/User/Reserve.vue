<template>
    <div>
        <UserNavBar></UserNavBar>
        <v-layout justify-center style="margin-top: 200px">
            <v-flex xs9 style="height: 100%;">
                <div v-bind:isLogIn="isLogIn"><flight-search v-on:search="search"></flight-search></div>
                <FlightGrid v-bind:flights="flights" v-bind:searchType="searchType"></FlightGrid>
            </v-flex>
        </v-layout>
    </div>
</template>

<script>
    import FlightGrid from "../../components/User/FlightGrid";
    import UserNavBar from "../../components/User/UserNavBar";
    import FlightSearch from "../../components/User/FlightSearch";
    import axios from "axios";
    import moment from 'moment';

    export default {
        name: "UserFlights",
        components: {FlightSearch, UserNavBar, FlightGrid},
        data() {
            return {
                flights: [],
                searchType: null,
                isLogIn: false,
            }
        },
        methods: {
            search: function(query) {
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
                        });
                    localStorage.setItem('departureDate',query.date);
                }
            }
        }
    }
</script>

<style scoped>

</style>