<template>
    <div>
        <UserNavBar></UserNavBar>
        <v-layout justify-center style="margin-top: 200px">
            <v-flex xs9 style="height: 100%;">
                <v-card>
                    <v-card-text>
                        <v-container>
                            <v-layout row>
                                <v-flex xs3>
                                    <img :src = '"http://ec2-18-195-170-20.eu-central-1.compute.amazonaws.com:8000/" + profile.Picture' width="250px" height="150px">
                                </v-flex>
                                <v-flex xs6>
                                    <v-layout row mb-3 align-center>
                                        <v-flex><span class = "display-1">{{profile.Name}}</span></v-flex>
                                        <v-flex>
                                            <v-rating v-model="rating" background-color="indigo lighten-3" empty-icon="$vuetify.icons.ratingFull"
                                            readonly half-increments color="orange" size="48"></v-rating>
                                        </v-flex>
                                    </v-layout>
                                    <v-layout row mt-4>
                                        <v-flex><span class = "title">Description:</span></v-flex>
                                    </v-layout>
                                    <v-layout row my-2>
                                        <v-flex><span class = "subheading">{{profile.Promo}}</span></v-flex>
                                    </v-layout> 
                                </v-flex> 
                            </v-layout>
                            <v-layout row>
                                <v-flex xs3>
                                </v-flex>
                                <v-flex xs6>
                                    <v-layout row>
                                        <v-flex mt-3>
                                            <span class = "title">Address: </span><span class = "subheading">{{profile.Address}}</span>
                                        </v-flex>
                                    </v-layout>
                                    <v-layout row mt-4>
                                        <v-flex><span class = "title">Location:</span></v-flex>
                                    </v-layout>
                                    <v-layout row>
                                        <v-flex v-if="updateMap">
                                            <gmap-map :center="{lat: profile.Latitude, lng: profile.Longitude}"
                                                      :zoom="8"
                                                      style="width:500px;  height: 300px;">
                                                            <gmap-marker :position="{lat: profile.Latitude, lng: profile.Longitude}"></gmap-marker>
                                            </gmap-map>
                                        </v-flex>
                                    </v-layout>
                                </v-flex>
                            </v-layout>
                            <v-layout>
                                <v-flex>
                                    <v-layout row>
                                        <v-flex mt-3>
                                            <span class = "title">Quick reservations: </span>
                                        </v-flex>
                                    </v-layout>
                                    <div v-if="update">
                                        <div v-if="reservations.length !== 0 ">
                                            <carousel :nav = true :items = 2 ref = proba>
                                                <template v-for="(reservation,index) in reservations">
                                                    <div :key="index">
                                                        <v-card width="80%" dark>
                                                            <v-card-text>
                                                                <v-layout row align-center>
                                                                    <v-flex xs6>
                                                                        <v-layout row>
                                                                            <v-flex><span class = "subheading">From: {{reservation.OriginName}}</span></v-flex>
                                                                        </v-layout>
                                                                        <v-layout row>
                                                                            <v-flex><span class = "subheading">To: {{reservation.DestName}}</span></v-flex>
                                                                        </v-layout>
                                                                        <v-layout row>
                                                                            <v-flex><span class = "subheading">Departure: {{reservation.Departure | moment}}</span></v-flex>
                                                                        </v-layout>
                                                                    </v-flex>
                                                                    <v-flex xs5>
                                                                        <v-layout row>
                                                                            <v-flex><span class = "subheading">Seat class: {{reservation.Class}}</span></v-flex>
                                                                        </v-layout>
                                                                        <v-layout row>
                                                                            <v-flex><span class = "subheading">Normal Price: {{getRealPrice(reservation)}}</span></v-flex>
                                                                        </v-layout>
                                                                        <v-layout row>
                                                                            <v-flex><span class = "subheading">Quick Res Price: {{(reservation.Price).toFixed(2)}}</span></v-flex>
                                                                        </v-layout>
                                                                    </v-flex>
                                                                    <v-flex>
                                                                        <v-btn small round color="indigo" @click="reserve(reservation)">Reserve</v-btn>
                                                                    </v-flex>
                                                                </v-layout>
                                                            </v-card-text>
                                                        </v-card>
                                                    </div>
                                                </template>
                                            </carousel>
                                        </div>
                                        <span v-else class="subheading"> No quick reservations to display please come back later!
                                        </span>
                                    </div>
                                </v-flex>
                            </v-layout>
                        </v-container>
                    </v-card-text>
                </v-card>
            </v-flex>
        </v-layout>
        <v-snackbar v-model="SuccessSnackbar" :timeout=4000 :top="true" color="success">{{succMessage}}</v-snackbar>
        <v-snackbar v-model="ErrorSnackbar" :timeout=4000 :top="true" color="error">{{errorMessage}}</v-snackbar>
    </div>
</template>

<script>
    import UserNavBar from "../../components/User/UserNavBar";
    import carousel from 'vue-owl-carousel'
    import axios from 'axios';
export default {
    components: {UserNavBar, carousel },
    name: "AirlineProfile",
    data() {
            return {
                profile: {
                    Picture:"",
                },
                update:false,
                updateMap: false,
                reservations: [],
                succMessage: "",
                errorMessage: "",
                SuccessSnackbar: false,
                ErrorSnackbar: false,
                rating: 0,
        }
    },
    beforeCreate() {
        axios.get("http://ec2-18-195-170-20.eu-central-1.compute.amazonaws.com:8000/api/airline/"+ this.$route.params.id + "/profile")
            .then(res =>{
                this.profile = res.data;
                this.updateMap = true;
        });

        axios.get("http://ec2-18-195-170-20.eu-central-1.compute.amazonaws.com:8000/api/airline/"+ this.$route.params.id + "/rating")
            .then(res =>{
                this.rating = res.data;
        });
        
    },
    created(){
        this.getReservations();
    },
    methods:{
        getReservations(){
            axios.get("http://ec2-18-195-170-20.eu-central-1.compute.amazonaws.com:8000/api/airline/"+ this.$route.params.id + "/quickReservations")
            .then(res =>{
                this.reservations = res.data;
                this.update = true;
            });
        },
        getRealPrice(item){
            if(item.Class === "FIRST"){
                return parseFloat(item.PriceFIRSTCLASS).toFixed(2);
            }else if(item.Class === "ECONOMIC"){
                return parseFloat(item.PriceECONOMY).toFixed(2);
            }else{
                return (item.PriceBUSINESS).toFixed(2);
            }
        },
        reserve(item){
            axios.create({withCredentials: true}).get("http://ec2-18-195-170-20.eu-central-1.compute.amazonaws.com:8000/api/user/getProfile")
                .then(res=>{
                    axios.create({withCredentials: true}).post("http://ec2-18-195-170-20.eu-central-1.compute.amazonaws.com:8000/api/user/FlightQuickReservation", item)
                        .then(res=>{
                            let vm = this;
                            this.getReservations();
                            this.succMessage = "Successfully created reservation";
                            this.SuccessSnackbar = true;
                            setTimeout(function () {
                                vm.$router.go()
                            }, 3000);
                        })
                        .catch(err =>{
                            let vm = this;
                            if(err.response.status === 406){
                                this.errorMessage = "Quick reservation already taken!";
                                this.ErrorSnackbar = true;
                            }else{
                                this.errorMessage = "Failed to reserve please try again!";
                                this.ErrorSnackbar = true;
                            }
                            setTimeout(function () {
                                vm.$router.go()
                            }, 3000);
                        })
                })
                .catch(err =>{
                    this.errorMessage = "You need to be logged in"
                    this.ErrorSnackbar = true
                })
        }
    }
}
</script>

<style>

</style>
