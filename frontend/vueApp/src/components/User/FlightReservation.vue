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
                            <v-flex xs6 mt-1 justify-center>
                                <v-layout row mb-2>
                                    <span class="title">Seat map:</span>
                                </v-layout>
                                <v-layout row>
                                    <seat-map :editedSeats="flight.Airplane" :tab="mode" @seatsChanged="updateSeats"></seat-map>
                                </v-layout>
                            </v-flex>
                            <v-flex xs6>
                                <v-layout column v-if="this.Seats.length !== 0">
                                    <v-flex xs2><h3>Passengers:</h3></v-flex>
                                    <v-flex>
                                        <div v-for="index in this.Seats.length" :key="index">
                                            <v-layout row>
                                                <v-flex grow v-if='userInfo[index - 1].Type === "Self"'>
                                                        <v-layout row>
                                                            <v-flex class="self_field">
                                                                <v-text-field label="Email" readonly v-model="self.Email"></v-text-field>
                                                            </v-flex>
                                                            <v-flex class="self_field">
                                                                <v-text-field label="Name" readonly v-model="self.Name"></v-text-field>
                                                            </v-flex>
                                                            <v-flex class="self_field">
                                                                <v-text-field label="Surname" readonly v-model="self.Surname"></v-text-field>
                                                            </v-flex>
                                                            <v-flex class="self_field">
                                                                <v-text-field label="Passport" readonly v-model="self.Passport"></v-text-field>
                                                            </v-flex>
                                                            <v-flex>
                                                                <v-icon large>check_circle_outline</v-icon>
                                                            </v-flex>
                                                        </v-layout>
                                                </v-flex>
                                                <v-flex xs4 v-else>
                                                    <v-select v-model="userInfo[index - 1].Type" :items="$options.passengerTypeValues" label="Type"></v-select>
                                                </v-flex>
                                                <template v-if='userInfo[index - 1].Type === "Unregistered"'>
                                                    <v-flex xs2 mx-4>
                                                        <v-btn @click="enterUserInfoDialog(index - 1)">Passenger Info</v-btn>
                                                    </v-flex>
                                                    <v-flex xs pt-2>
                                                        <v-icon v-if='checkInfo(index - 1, "Unregistered")' large>check_circle_outline</v-icon>
                                                        <v-icon v-else large>highlight_off</v-icon>
                                                    </v-flex>
                                                </template>
                                                <template v-else-if='userInfo[index - 1].Type === "Friend"'>
                                                    <v-flex xs6 ml-4>
                                                        <v-select v-model="userInfo[index - 1].Email" cache-items :items="friendsSelect" label="Friends"
                                                                  no-data-text="You have no friends to call"
                                                                  item-text="Text" item-value="Value"></v-select>
                                                    </v-flex>
                                                    <v-flex xs pt-2>
                                                        <v-icon v-if='checkInfo(index - 1, "Friend")' large>check_circle_outline</v-icon>
                                                        <v-icon v-else large>highlight_off</v-icon>
                                                    </v-flex>
                                                </template>
                                            </v-layout>
                                        </div>
                                        <v-flex>
                                            <v-btn @click="reserve" block color="info">Reserve</v-btn>
                                        </v-flex>
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
            <v-snackbar
                    v-model="snackbar_fields.snackbar"
                    :timeout="4000"
                    :top="true"
                    :color="snackbar_fields.color"
            >
                {{ snackbar_fields.text }}
                <v-btn
                        flat
                        @click="snackbar_fields.snackbar = false"
                >
                    Close
                </v-btn>
            </v-snackbar>
        </v-layout>
    </v-container>
</template>

<script>
    import SeatMap from "../AirlineAdmin/SeatMap";
    import axios from 'axios/index';
    import UserNavBar from "./UserNavBar";

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
                self: null,
                isLogIn: false,
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
                snackbar_fields: {
                    snackbar: false,
                    text: null,
                    color: null,
                },
            }
        },
        beforeCreate(){
            axios.create({withCredentials: true}).get("http://localhost:8000/api/user/getProfile")
                .then(res =>{
                    this.isLogIn = true;
                })
                .catch(err =>{
                    this.isLogIn = false;
                })
        },
        created(){
            axios.get("http://localhost:8000/api/flight/"+ this.flightId +"/getFlight")
                 .then(res =>{
                        this.flight = res.data;
                 });
            axios.create({withCredentials:true}).get("http://localhost:8000/api/user/getFriends")
                .then(res =>{
                        this.friends = res.data;
                });
            axios.create({withCredentials: true}).get("http://localhost:8000/api/user/getProfile")
                .then(res => this.self = res.data);
        },
        methods:{
            updateSeats(value){
                if(this.userInfo.length > value.length){
                    this.userInfo.pop();
                }else{
                    if (this.userInfo.length === 0) {
                        this.userInfo.push({
                            Type: "Self",
                            Name: this.self.Name,
                            Surname: this.self.Surname,
                            Passport: this.self.Passport,
                            Email: this.self.Email,
                        });
                    } else {
                        this.userInfo.push({
                            Type: "",
                            Name: "",
                            Surname: "",
                            Passport: "",
                            Email: "",
                        });
                    }
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
                if (type === "Unregistered") {
                    return !(!this.userInfo[index].Name || !this.userInfo[index].Surname || !this.userInfo[index].Passport);
                }else{
                    return !(!this.userInfo[index].Email)
                }
            },
            reserve(e) {
                e.preventDefault();
                localStorage.setItem('destination',this.flight.Destination.Name);
                for (let i = 0; i < this.userInfo.length; i++) {
                    if (!this.checkInfo(i, this.userInfo[i].Type)) {
                        this.snackbar_fields.text = "Please fill in the required fields";
                        this.snackbar_fields.color = "warning";
                        this.snackbar_fields.snackbar = true;
                        return;
                    }
                }

                const payload = {
                    "Seats": this.Seats,
                    "Users": this.userInfo,
                };

                axios.create({withCredentials: true}).post('http://localhost:8000/api/reservations/airline/reserve/q='+this.flightId, payload)
                    .then(res => {
                        if (res.status === 200) {
                            this.snackbar_fields.text = "Reservation successful";
                            this.snackbar_fields.color = "success";
                            this.snackbar_fields.snackbar = true;
                            this.$router.push({ path: '/user/reserve', query: { reservationID: res.data, passengers: payload.Users.length.toString() }});
                        }
                    })
                    .catch(err => {
                        this.snackbar_fields.text = "Reservation unsuccessful";
                        this.snackbar_fields.color = "error";
                        this.snackbar_fields.snackbar = true;
                        this.$router.go();
                    });
            }
        },
        computed:{
          mode: function(){
              if(this.isLogIn){
                  return "reservation";
              }else{
                  return "explore";
              }
          },
          friendsSelect:function(){
              const friends = JSON.parse(JSON.stringify(this.friends));
              let retVal = [];
              let selected = [];

              this.userInfo.forEach(user => {
                  if (user.Type === "Friend") {
                      selected.push(user);
                  }
              });

              for (let i = 0; i < selected.length; i++) {
                  for (let j = 0; j < friends.length; j++) {
                      if (friends[j].Email === selected[i].Email) {
                          friends.splice(j, 1);
                      }
                  }
              }

              friends.forEach(friend =>
                {retVal.push( {
                    Text: friend.Name + " " + friend.Surname + "(" + friend.Email+")", Value: friend.Email
                });
              });
              return retVal
          }
        }
    }
</script>

<style scoped>
    @import '../../assets/css/SeatMap.css';
   .self_field {
       margin-right: 10px
   }
</style>