<template>
    <div>
        <v-flex xs12>
            <template v-if="searchType ==='oneWay'">
                <v-data-table :headers="headers" :items="flights" class="elevation-1">
                    <template v-slot:items="props">
                        <td>{{props.item.Origin.Name}}</td>
                        <td>{{props.item.Destination.Name}}</td>
                        <td>{{props.item.Departure | moment}}</td>
                        <td>{{props.item.Landing | moment}}</td>
                    </template>
                </v-data-table>
            </template>
            <template v-else-if="searchType === 'round' && flights.length !== 0 && flightsReturn.length !== 0">
                <v-data-table :headers="headers" :items="flights" class="elevation-1">
                    <template v-slot:items="props">
                        <td>{{props.item.Origin.Name}}</td>
                        <td>{{props.item.Destination.Name}}</td>
                        <td>{{props.item.Departure | moment}}</td>
                        <td>{{props.item.Landing | moment}}</td>
                    </template>
                </v-data-table>
                <v-data-table :headers="headers" :items="flightsReturn" class="elevation-1">
                    <template v-slot:items="props">
                        <td>{{props.item.Origin.Name}}</td>
                        <td>{{props.item.Destination.Name}}</td>
                        <td>{{props.item.Departure | moment}}</td>
                        <td>{{props.item.Landing | moment}}</td>
                    </template>
                </v-data-table>
            </template>
        </v-flex>
    </div>
    
</template>

<script>
    import axios from 'axios';
    import moment from 'moment';

    export default {
        name: "FlightSearchResults",
        data(){
            return{
                flights:[],
                flightsReturn:[],
                searchType: "",
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
                        text: 'Landing',
                        value: "Landing",
                    },]
            }
        },
        created() {
            if(this.$route.query.returnDate){
                this.searchType = "round";
                console.log("Round");
                var searchQuery = {
                    from : this.$route.query.from,
                    to : this.$route.query.to,
                    date: this.$route.query.departureDate,
                    seatClass: this.$route.query.seatClass,
                    passengers: this.$route.query.passengers.toString()
                };
                var dateMoment = moment(searchQuery.date,"YYYY-MM-DD");
                searchQuery.date = dateMoment.valueOf().toString();
                axios.post("http://localhost:8000/api/search/oneWay",searchQuery)
                    .then(res=>{
                        this.flights = res.data;
                    });
                var searchQuery2 = {
                    from : this.$route.query.to,
                    to : this.$route.query.from,
                    date: this.$route.query.returnDate,
                    seatClass: this.$route.query.seatClass,
                    passengers: this.$route.query.passengers.toString()
                };
                var dateMoment2 = moment(searchQuery2.date,"YYYY-MM-DD");
                searchQuery2.date = dateMoment2.valueOf().toString();
                axios.post("http://localhost:8000/api/search/oneWay", searchQuery2)
                    .then(res=>{
                        this.flightsReturn = res.data;
                    });
            }else if(this.$route.query.date){
                this.searchType = "oneWay";
                console.log("One-Way");
                var searchQuery = {
                    from : this.$route.query.from,
                    to : this.$route.query.to,
                    date: this.$route.query.date,
                    seatClass: this.$route.query.seatClass,
                    passengers: this.$route.query.passengers.toString()
                };
                var dateMoment = moment(searchQuery.date,"YYYY-MM-DD");
                searchQuery.date = dateMoment.valueOf().toString();
                axios.post("http://localhost:8000/api/search/oneWay", searchQuery)
                    .then(res=>{
                        this.flights = res.data;
                    })
            }else{
                console.log("multi");
                this.searchType = "oneWay";
                var searchQuery = {
                    from : this.$route.query.from,
                    to : this.$route.query.to,
                    layovers: this.$route.query.layovers,
                    date: this.$route.query.departureDate,
                    seatClass: this.$route.query.seatClass,
                    passengers: this.$route.query.passengers.toString()
                };
                var dateMoment = moment(searchQuery.date,"YYYY-MM-DD");
                searchQuery.date = dateMoment.valueOf().toString();
                axios.post("http://localhost:8000/api/search/multi", searchQuery)
                    .then(res=>{
                        this.flights = res.data;
                    })
            }

        }
    }
</script>

<style scoped>

</style>