<template>
    <div>
        <UserNavBar></UserNavBar>
        <v-layout justify-center style="margin-top: 200px">
            <v-flex xs9 style="height: 100%;">
                <v-layout row>
                    <v-flex xs11>
                        <v-text-field v-model="query" solo label = "Enter Users First name, Last name or Email"></v-text-field>
                    </v-flex>
                    <v-flex xs1>
                        <v-btn dark @click="search">Search</v-btn>
                    </v-flex>
                </v-layout>
                <template v-for="(result, index) in results">
                    <v-layout row :key="result.Email">
                        <v-flex xs12>
                        <v-card >
                            <v-card-text>
                                <v-layout row wrap>
                                    <v-flex xs2>
                                        <img src="../../assets/user/userPlaceholder.png" height="100px" width="100px">
                                    </v-flex>
                                    <v-flex xs7>
                                        <v-layout row><span class="headline">{{result.Name + " " + result.Surname}}</span></v-layout>
                                        <v-layout row><span class="subheading">{{result.Email}}</span></v-layout>
                                    </v-flex>
                                    <v-flex xs3 justify-center>
                                        <v-layout row wrap>
                                            <v-flex xs12 v-if="result.Status=== ''">
                                                <v-btn color="primary" @click="addFriend(result.ID, index)" >Add a friend</v-btn>
                                            </v-flex>
                                            <v-flex xs12 v-else>
                                                <v-btn color="primary" disabled>Pending to be accepted</v-btn>
                                            </v-flex>
                                        </v-layout>
                                    </v-flex>
                                </v-layout>
                            </v-card-text>
                        </v-card>
                        </v-flex>
                    </v-layout>
                </template>
            </v-flex>
        </v-layout>
    </div>
</template>

<script>
    import axios from 'axios';
    import UserNavBar from "../../components/User/UserNavBar";
    export default {
        name: "UserSearch",
        components: {UserNavBar},
        data:function() {
            return {
                query: "",
                results:[],
                ind: false
            };
        },
        methods:{
            search(){
                this.ind = true;
                var QueryDto = {
                    Query: this.query
                };
                axios.create({withCredentials: true}).post("http://localhost:8000/api/user/search", QueryDto).
                then(res=>{
                    this.results =res.data
                })
            },
            addFriend(id, index){
                var friendshipDto = {
                    User2ID: id
                };
                axios.create({withCredentials: true}).post("http://localhost:8000/api/user/friendRequests", friendshipDto)
                    .then(res =>{
                        this.results[index].Status = "PENDING";
                    })
            }
        }
    }
</script>

<style scoped>

</style>