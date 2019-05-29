<template>
    <v-container fluid>
        <v-layout row>
            <v-flex>
                <UserNavBar/>
            </v-flex>
        </v-layout>
        <v-layout align-center justify-space-around row wrap fill-height style="margin-top: 100px">
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
            <v-flex xs12 style="margin-top: 2vw">
                <v-btn dark block @click="reserve">Reserve</v-btn>
            </v-flex>
        </v-layout>
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
                rooms: [],
                selected: [],
                select: false,
                checkboxes: [],
                reservationID: this.$route.params.reservation,
                hotelID: this.$route.params.hotel_id,
                start: this.$route.params.start,
                end:  this.$route.params.end,
            }
        },
        created() {
            const query = {
                "From": this.start,
                "To": this.end,
            };
            axios.create({withCredentials: true}).post('http://localhost:8000/api/search/' + this.hotelID + '/quickRooms', query)
                .then(res => {
                    this.rooms = res.data;
                    this.checkboxes = [];
                    for (let i = 0; i < this.rooms.length; i++) {
                        this.checkboxes.push(false);
                    }
                })
                .catch(err => alert(err));
        },
        methods: {
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
                    "From": this.start,
                    "To": this.end,
                    "IsQuickReserve": true,
                };
                axios.create({withCredentials: true}).post('http://localhost:8000/api/reservations/hotel/' + this.hotelID + '/reserve/masterRef='+ this.reservationID, query)
                    .then(res => {
                        this.$router.push({ path: '/user', query: { reservationID: this.reservationID, passengers: this.passengers }});
                    })
                    .catch(err => alert(err));
            }
        },
    }
</script>

<style scoped>

</style>