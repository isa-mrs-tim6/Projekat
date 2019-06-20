<template>
    <div>
        <UserNavBar></UserNavBar>
        <v-layout justify-center style="margin-top: 200px">
            <v-flex xs9 style="height: 100%;">
                <v-card>
                    <v-card-title>
                        <v-layout row align-center>
                            <v-flex xs8>
                                <span class= "headline">Airlines in our system:</span>
                            </v-flex>
                            <v-flex mt-0 pt-0>
                                <v-text-field v-model="query" label="Search" ></v-text-field>
                            </v-flex>
                        </v-layout>
                    </v-card-title>
                    <v-card-text>
                        <template v-for="airline in airlinesDisplay">
                            <v-layout row :key="airline.ID" my-4 justify-center>
                                <v-flex xs10 >
                                    <v-card dark>
                                        <v-card-text>
                                            <v-layout row align-center>
                                                <v-flex xs2>
                                                    <img :src = '"http://ec2-18-195-170-20.eu-central-1.compute.amazonaws.com:8000/" + airline.Picture' width="160px" height="100px">
                                                </v-flex>
                                                <v-flex xs5 mx-5>
                                                    <v-layout row mb-4><span class="headline">Airline Name: {{airline.Name}}</span></v-layout>
                                                    <v-layout row><span class="subheading">Airline Adress: {{airline.Address}}</span></v-layout>
                                                </v-flex>
                                                <v-flex xs4 align-center>
                                                    <v-layout row  justify-center>
                                                        <v-btn large round color="indigo" @click="goToAirline(airline.ID)">Explore</v-btn>
                                                    </v-layout>
                                                </v-flex>
                                            </v-layout>
                                        </v-card-text>
                                    </v-card>
                                </v-flex>
                            </v-layout>
                        </template>
                    </v-card-text>
                </v-card>
            </v-flex>
        </v-layout>
    </div>
    
</template>
<script>
    import UserNavBar from "../../components/User/UserNavBar";
    import axios from 'axios';
    export default {
        components: {UserNavBar},
        name: "AirlineList",
        data() {
            return {
                airlines: [],
                query: "",
            }
        },
        beforeCreate() {
            axios.get("http://ec2-18-195-170-20.eu-central-1.compute.amazonaws.com:8000/api/airline")
                .then(res =>{
                    this.airlines = res.data;
                })
        },
        methods:{
            goToAirline(id){
                this.$router.push("airline/" + id);
            }
        },
        computed:{
            airlinesDisplay: function(){
                var retVal = this.airlines;
                retVal = retVal.filter(airline => airline.Name.toLowerCase().includes(this.query.toLowerCase()))
                return retVal;
            }
        }
    }
</script>

<style>

</style>
