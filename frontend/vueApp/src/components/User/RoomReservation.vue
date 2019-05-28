<template>
    <v-container fluid>
        <v-layout row>
            <v-flex>
                <UserNavBar/>
            </v-flex>
        </v-layout>
        <RoomSearch v-on:search="search"  v-bind:room_capacities="capacities" style="margin-top: 100px"/>
        <v-layout align-center justify-space-around row wrap fill-height>
            <v-flex xs12 style="margin-top: 2vw" v-for="(value, index) in rooms" v-bind:key="index">
                <v-card dark class="white--text">
                    <v-card-title class="headline" primary-title>
                        <v-flex xs2 style="margin-right: 2vh">
                            Number: {{value.Number}}
                        </v-flex>
                        <v-flex xs2 style="margin-right: 2vh">
                            <v-icon> group </v-icon> Capacity: {{value.Capacity}}
                        </v-flex>
                        <v-flex xs2 style="margin-right: 2vh">
                            <v-icon> euro_symbol </v-icon> Price: {{value.Price}}
                        </v-flex>
                        <v-flex xs2 style="float:right; margin-right: 2vh">
                            <v-checkbox v-model="checkboxes[index]" label="Reserve?"></v-checkbox>
                        </v-flex>
                    </v-card-title>
                </v-card>
            </v-flex>
            <v-flex xs12 style="margin-top: 2vw" v-if="allow">
                <v-btn dark block @click="reserve">Reserve</v-btn>
            </v-flex>
        </v-layout>
    </v-container>
</template>

<script>
    import RoomSearch from "./RoomSearch";
    import UserNavBar from "./UserNavBar";
    import axios from 'axios';
    import moment from 'moment';

    export default {
        name: "RoomReservation",
        components: {UserNavBar, RoomSearch},
        data() {
            return {
                num_of_guests: null,
                capacities: null,
                rooms: null,
                selected: [],
                select: false,
                checkboxes: [],
                reservationID: this.$route.params.reservation,
                hotelID: this.$route.params.hotel_id,
                tickets: this.$route.params.tickets,
                time: {
                    from: null,
                    to: null,
                },
            }
        },
        created() {
            axios.create({withCredentials: true}).get('http://localhost:8000/api/hotel/' + this.hotelID + '/getRoomCapacities')
                .then(res => {
                    this.capacities = res.data;
                })
                .catch(err => alert(err));
        },
        methods: {
            search: function (query) {
                this.num_of_guests = query.Guests;
                this.time.from = query.From;
                this.time.to = query.To;
                axios.create({withCredentials: true}).post('http://localhost:8000/api/search/' + this.hotelID + '/rooms', query)
                    .then(res => {
                        this.rooms = res.data;
                        this.checkboxes = [];
                        for (let i = 0; i < this.rooms.length; i++) {
                            this.checkboxes.push(false);
                        }
                    })
                    .catch(err => alert(err));
            },
            reserve() {
                let rooms = [];


                for (let i = 0; i < this.checkboxes.length; i++) {
                    if (this.checkboxes[i]) {
                        rooms.push(this.rooms[i])
                    }
                }

                if (rooms.length === 0) {
                    return
                }

                const query = {
                    "Rooms": rooms,
                    "From": this.time.from,
                    "To": this.time.to,
                };
                axios.create({withCredentials: true}).post('http://localhost:8000/api/reservations/hotel/' + this.hotelID + '/reserve/masterRef='+ this.reservationID, query)
                    .then(res => {
                        this.$router.push({ path: '/user', query: { reservationID: this.reservationID, passengers: this.passengers }});
                    })
                    .catch(err => alert(err));
            }
        },
        computed:{
            allow:function(){
                let retVal = this.checkboxes;
                if(this.checkboxes.length > 0){
                    retVal = retVal.filter(ind => ind);
                }
                return retVal.length > 0;
            }
        }
    }
</script>

<style scoped>

</style>