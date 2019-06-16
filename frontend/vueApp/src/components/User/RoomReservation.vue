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
                            Number: {{value.Room.Number}}
                            <v-rating
                                    v-model=value.Rating
                                    background-color="orange lighten-3"
                                    color="orange"
                                    readonly
                                    medium
                            ></v-rating>
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
                <v-dialog
                        v-model="dialog"
                        width="500"
                >
                    <template v-slot:activator="{ on }">
                        <v-btn
                                color="blue lighten-2"
                                dark
                                block
                                v-on="on"
                        >
                            Extra features
                        </v-btn>
                    </template>

                    <v-card>
                        <v-card-title
                                class="headline grey lighten-2"
                                primary-title
                        >
                            Take a look at our offerings:
                        </v-card-title>

                        <v-card-text>
                            <v-flex>
                                <v-autocomplete
                                        :items="features"
                                        v-model="selected_features"
                                        label="Features"
                                        item-text="Name"
                                        item-value="Name"
                                        multiple
                                        chips
                                        return-object
                                >
                                    <template slot="selection" slot-scope="data">
                                        <v-chip
                                                :selected="data.selected"
                                                :key="JSON.stringify(data.item)"
                                                close
                                                class="chip--select-multi"
                                                @input="data.parent.selectItem(data.item)"
                                        >
                                            <v-icon style="margin-right: 10px;">
                                                {{data.item.Icon}}
                                            </v-icon>
                                            {{ data.item.Name }}
                                        </v-chip>
                                    </template>
                                    <template slot="item" slot-scope="data">
                                        <template >
                                            <v-list-tile-avatar>
                                                <v-icon>{{data.item.Icon}}</v-icon>
                                            </v-list-tile-avatar>
                                            <v-list-tile-content>
                                                <v-list-tile-title v-html="data.item.Name"></v-list-tile-title>
                                                <v-list-tile-sub-title v-html="data.item.Description"></v-list-tile-sub-title>
                                            </v-list-tile-content>
                                            <v-list-tile-action>
                                                <v-list-tile-action-text>{{ data.item.Price }}€</v-list-tile-action-text>
                                            </v-list-tile-action>
                                        </template>
                                    </template>
                                </v-autocomplete>
                            </v-flex>
                        </v-card-text>
                    </v-card>
                </v-dialog>
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
                dialog: false,
                num_of_guests: null,
                selected_features: [],
                features: [],
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
            axios.create({withCredentials: true}).get('http://localhost:8000/api/hotel/features', {
                params: {
                    hotel: this.hotelID
                }
            })
                .then(res => {
                    this.features = res.data;
                })
                .catch(err => alert("Could not retrieve features"))
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
                        rooms.push(this.rooms[i].Room)
                    }
                }

                if (rooms.length === 0) {
                    return
                }

                const query = {
                    "Rooms": rooms,
                    "Features": this.selected_features,
                    "From": this.time.from,
                    "To": this.time.to,
                };
                axios.create({withCredentials: true}).post('http://localhost:8000/api/reservations/hotel/' + this.hotelID + '/reserve/masterRef='+ this.reservationID, query)
                    .then(res => {
                        this.$router.push({ path: '/user/reserve', query: { reservationID: this.reservationID, passengers: this.passengers }});
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