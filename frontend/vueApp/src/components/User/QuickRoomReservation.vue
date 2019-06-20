<template>
    <v-container fluid>
        <v-layout row>
            <v-flex>
                <UserNavBar/>
            </v-flex>
        </v-layout>
        <template v-if="rooms != null">
            <v-layout align-center justify-space-around row wrap fill-height style="margin-top: 100px">
                <v-flex xs12 style="margin-top: 2vw" v-for="(value, index) in rooms" v-bind:key="index">
                    <v-card dark class="white--text">
                        <v-card-title class="headline" primary-title>
                            <v-flex xs2 style="margin-right: 2vh">
                                Number: {{value.Room.Number}}
                            </v-flex>
                            <v-flex xs2 style="margin-right: 2vh">
                                <v-icon> group </v-icon> Capacity: {{value.Room.Capacity}}
                            </v-flex>
                            <v-flex xs2 style="margin-right: 2vh">
                                <v-icon> euro_symbol </v-icon>
                                <template v-if="value.Room.Price !== value.Price">
                                    Price: <span style="text-decoration: line-through;">{{value.Room.Price}}</span> {{value.Price}}
                                </template>
                                <template v-else>
                                    Price: {{value.Room.Price}}
                                </template>
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
        </template>
        <template v-else>
            <v-layout align-center justify-space-around row wrap fill-height style="margin-top: 100px">
                <v-flex xs12 style="margin-top: 2vw" class="text-xs-center">
                    <h1 class="display-4" style="margin-bottom: 2vw">No rooms available at the moment</h1>
                    <h1 class="display-3 font-weight-regular font-italic">Perhaps try the regular <a @click="takeMeThere" :href="`localhost:8080/reserve_room/${this.hotelID}/${this.reservationID}/${this.tickets}`">reservations</a>?</h1>
                    <h1 class="display-2" style="margin-top: 20vw">
                        <a @click="takeMeBack" :href="`http://localhost:8080/user/reserve?reservationID=${this.reservationID}&passengers=${this.tickets}`">Go back</a>
                    </h1>
                </v-flex>
            </v-layout>
        </template>
    </v-container>
</template>

<script>
    import UserNavBar from "./UserNavBar";
    import axios from 'axios';

    export default {
        name: "QuickRoomReservation",
        components: {UserNavBar},
        data() {
            return {
                dialog: false,
                rooms: [],
                selected: [],
                select: false,
                checkboxes: [],
                reservationID: this.$route.params.reservation,
                hotelID: this.$route.params.hotel_id,
                tickets: this.$route.params.tickets,
            }
        },
        created() {
            const query = {
                "From": localStorage.getItem("hotelStart"),
                "To": localStorage.getItem("hotelEnd"),
            };
            axios.create({withCredentials: true}).post('http://ec2-18-195-170-20.eu-central-1.compute.amazonaws.com:8000/api/search/' + this.hotelID + '/quickRooms', query)
                .then(res => {
                    this.rooms = res.data;
                    this.checkboxes = [];
                    if (this.rooms != null) {
                        for (let i = 0; i < this.rooms.length; i++) {
                            this.checkboxes.push(false);
                        }
                    }
                })
                .catch(err => alert(err));
        },
        methods: {
            reserve() {
                let rooms = [];
                for (let i = 0; i < this.checkboxes.length; i++) {
                    if (this.checkboxes[i]) {
                        rooms.push(this.rooms[i].Room)
                    }
                }
                if (rooms.length === 0) {
                    return
                }
                const query = {
                    "Rooms": rooms,
                    "From": localStorage.getItem("hotelStart"),
                    "To": localStorage.getItem("hotelEnd"),
                    "IsQuickReserve": true,
                };
                axios.create({withCredentials: true}).post('http://ec2-18-195-170-20.eu-central-1.compute.amazonaws.com:8000/api/reservations/hotel/' + this.hotelID + '/reserve/masterRef='+ this.reservationID, query)
                    .then(res => {
                        this.$router.push({ path: '/user/reserve', query: { reservationID: this.reservationID, passengers: this.tickets }});
                    })
                    .catch(err => alert(err));
            },
            takeMeThere() {
                this.$router.push({ path: `/reserve_room/${this.hotelID}/${this.reservationID}/${this.tickets}` });
            },
            takeMeBack() {
                this.$router.push({ path: '/user/reserve', query: { reservationID: this.reservationID, passengers: this.tickets }});
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