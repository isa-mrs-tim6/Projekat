<template>
    <v-container fluid>
        <v-layout row>
            <v-flex>
                <UserNavBar></UserNavBar>
            </v-flex>
        </v-layout>
        <v-layout row>
            <v-flex>
                <v-card style="margin-top: 100px">
                    <v-card-title><h3 class=".display-2">Flight reservation</h3></v-card-title>
                    <v-card-text>
                        <v-layout row>
                            <v-flex xs6 mt-3 justify-center>
                                <seat-map :editedSeats="flight.Airplane" tab="reservation" @seatsChanged="updateSeats"></seat-map>
                            </v-flex>
                            <v-flex xs6>
                                <v-layout column v-if="this.Seats.length !== 0">
                                    <v-flex xs2><h3>Passengers:</h3></v-flex>
                                    <v-flex>
                                        <div v-for="index in this.Seats.length" :key="index">
                                            <v-layout row>
                                                <v-flex xs4>
                                                    <v-select v-model="userInfo[index - 1].Type" :items="$options.passengerTypeValues" label="Type"></v-select>
                                                </v-flex>
                                                <template v-if='userInfo[index - 1].Type === "Unregistered"'>
                                                    <v-flex xs2 mx-4>
                                                        <v-btn @click="enterUserInfoDialog(index - 1)">Passenger Info</v-btn>
                                                    </v-flex>
                                                    <v-flex xs pt-2>
                                                        <v-icon v-if='checkInfo(index - 1, "unregistered")' large>check_circle_outline</v-icon>
                                                        <v-icon v-else large>highlight_off</v-icon>
                                                    </v-flex>
                                                </template>
                                                <template v-else-if='userInfo[index - 1].Type === "Friend"'>
                                                    <v-flex xs6 ml-4>
                                                        <v-select v-model="userInfo[index - 1].Email" :items="friendsSelect" label="Friends"
                                                                  no-data-text="You have no friends to call"
                                                                  item-text="Text" item-value="Value"></v-select>
                                                    </v-flex>
                                                    <v-flex xs pt-2>
                                                        <v-icon v-if='checkInfo(index - 1, "friend")' large>check_circle_outline</v-icon>
                                                        <v-icon v-else large>highlight_off</v-icon>
                                                    </v-flex>
                                                </template>
                                            </v-layout>
                                        </div>
                                    </v-flex>
                                </v-layout>
                            </v-flex>
                        </v-layout>
                    </v-card-text>
                </v-card>
                <v-dialog v-model="dialogUserInfo" persistent max-width="500px">
                    <v-card>
                        <v-card-title>
                            <span class="headline">Enter passenger info:</span>
                        </v-card-title>
                        <v-card-text>
                            <v-container>
                                <v-layout row-wrap justify-center>
                                    <v-flex xs11>
                                        <v-text-field v-model="editedItem.Name" label="First name" prepend-icon = "perm_identity"></v-text-field>
                                    </v-flex>
                                </v-layout>
                                <v-layout row-wrap justify-center>
                                    <v-flex xs11>
                                        <v-text-field v-model="editedItem.Surname" label="Surname" prepend-icon = "perm_identity"></v-text-field>
                                    </v-flex>
                                </v-layout>
                                <v-layout row-wrap justify-center>
                                    <v-flex xs11>
                                        <v-text-field v-model="editedItem.Passport" label="Passport" prepend-icon = "perm_identity"></v-text-field>
                                    </v-flex>
                                </v-layout>
                            </v-container>
                        </v-card-text>
                        <v-card-actions>
                            <v-spacer></v-spacer>
                            <v-btn color="blue darken-1" flat @click="close">Cancel</v-btn>
                            <v-btn color="blue darken-1" flat @click="save">Save</v-btn>
                        </v-card-actions>
                    </v-card>
                </v-dialog>
            </v-flex>
        </v-layout>
    </v-container>
</template>

<script>
    import UserNavBar from "../UserNavBar";
    import SeatMap from "../AirlineAdmin/SeatMap";
    import axios from 'axios/index';

    export default {
        name: "FlightReservation",
        components: {UserNavBar, SeatMap},
        passengerTypeValues:["Friend", "Unregistered"],
        defaultItem:{
            Name: '',
            Surname: '',
            Passport: '',
        },
        data(){
            return{
                flightId: this.$route.query.flightId,
                con: 1000,
                flight:{
                    Airplane:{
                        Seats:[]
                    }
                },
                editedItem:{
                    Name: '',
                    Surname: '',
                    Passport: '',
                },
                editedItemIndex : -1,
                Seats:[],
                userInfo:[],
                friends:[],
                dialogUserInfo: false,
            }
        },
        created(){
            axios.get("http://localhost:8000/api/flight/"+ this.flightId +"/getFlight")
                 .then(res =>{
                        this.flight = res.data;
                 });
            axios.create({withCredentials:true}).get("http://localhost:8000/api/user/getFriends")
                .then(res =>{
                        this.friends = res.data;
                })
        },
        methods:{
            updateSeats(value){
                if(this.userInfo.length > value.length){
                    this.userInfo.pop();
                }else{
                    this.userInfo.push({
                        Type: "",
                        Name: "",
                        Surname: "",
                        Passport: "",
                        Email: "",
                    });
                }
                this.Seats = value;
            },
            enterUserInfoDialog(index){
                this.editedItemIndex = index;
                this.editedItem.Name = this.userInfo[index].Name;
                this.editedItem.Surname = this.userInfo[index].Surname;
                this.editedItem.Passport = this.userInfo[index].Passport;
                this.dialogUserInfo = true;
            },
            save(){
                if(!this.editedItem.Name || !this.editedItem.Surname || !this.editedItem.Passport){
                    return;
                }
                Object.assign(this.userInfo[this.editedItemIndex], this.editedItem);
                this.close();
            },
            close(){
                this.editedItem = Object.assign({}, this.$options.defaultItem);
                this.editedItemIndex = -1;
                this.dialogUserInfo = false;
            },
            checkInfo(index, type){
                if (type === "unregistered") {
                    return !(!this.userInfo[index].Name || !this.userInfo[index].Surname || !this.userInfo[index].Passport);
                }else{
                    return !(!this.userInfo[index].Email)
                }
            },
        },
        computed:{
          friendsSelect:function(){
              var list = JSON.parse(JSON.stringify(this.friends));
              var retVal = [];
              list.forEach(friend =>{retVal.push( {
                                        Text: friend.Name + " " + friend.Surname + "(" + friend.Email+")",
                                        Value: friend.Email
                                    });});
              return retVal
          }
        }
    }
</script>

<style scoped>
    @import '../../assets/css/SeatMap2.css';
   .d {
       display: flex;
       flex-wrap: nowrap;
   }
</style>