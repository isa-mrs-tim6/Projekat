<template>
    <v-layout row-wrap>
        <v-flex x12>
            <template v-if="this.requests.length !== 0">
                <v-layout row mb-2>
                    <span class="display-1">Your requests:</span>
                </v-layout>
            </template>
            <template v-for="request in requests">
                <v-card :key="request.Email">
                    <v-card-text>
                        <v-layout row wrap>
                            <v-flex xs2>
                                <img src="../../assets/user/userPlaceholder.png" height="100px" width="100px">
                            </v-flex>
                            <v-flex xs7>
                                <v-layout row><span class="headline">{{request.Name + " " + request.Surname}}</span></v-layout>
                                <v-layout row><span class="subheading">{{request.Email}}</span></v-layout>
                            </v-flex>
                            <v-flex xs3 justify-center>
                                <v-layout row wrap>
                                    <v-flex xs6>
                                        <v-btn color="primary" @click="updateRequest(request.ID, true)">Accept</v-btn>
                                    </v-flex>
                                    <v-flex xs6>
                                        <v-btn color="error" @click="updateRequest(request.ID, false)">Reject</v-btn>
                                    </v-flex>
                                </v-layout>
                            </v-flex>
                        </v-layout>
                    </v-card-text>
                </v-card>
            </template>
            <template v-if="this.friends.length === 0">
                <v-layout row justify-center mt-2>
                    <span class="display-1">You have no friends</span>
                </v-layout>
            </template>
            <template v-else>
                <v-layout row my-2>
                    <span class="display-1">Your friends:</span>
                </v-layout>
            </template>
            <template v-for="friend in friends">
                <v-card :key="friend.Email">
                    <v-card-text>
                        <v-layout row>
                            <v-flex xs2>
                                <img src="../../assets/user/userPlaceholder.png" height="100px" width="100px">
                            </v-flex>
                            <v-flex xs7>
                                <v-layout row><span class="headline">{{friend.Name + " " + friend.Surname}}</span></v-layout>
                                <v-layout row><span class="subheading">{{friend.Email}}</span></v-layout>
                            </v-flex>
                            <v-flex xs3 justify-center>
                                <v-btn color="error" @click="removeAFriends(friend.ID)">Remove a friend</v-btn>
                            </v-flex>
                        </v-layout>
                    </v-card-text>
                </v-card>
            </template>
        </v-flex>
    </v-layout>
</template>

<script>
    import axios from "axios";
    export default {
        name: "UserFriends",
        data (){
            return {
                friends: [],
                requests: [],
            }
        },
        created() {
            this.updateFriends();
            this.updateRequests();
        },
        methods:{
            updateFriends(){
                axios.create({withCredentials: true}).get("http://localhost:8000/api/user/getFriends")
                    .then(res=>{
                        this.friends = res.data;
                    })
            },
            updateRequests(){
                axios.create({withCredentials: true}).get("http://localhost:8000/api/user/friendRequests")
                    .then(res=>{
                        this.requests = res.data;
                    })
            },
            updateRequest(id, status){
                var friendshipDto = {
                    User2ID: id,
                    Status: status
                };
                axios.create({withCredentials: true}).put("http://localhost:8000/api/user/friendRequests",friendshipDto)
                    .then(res => {
                        this.updateRequests();
                        this.updateFriends();
                    })
            },
            removeAFriends(id){
                var friendshipDto = {
                    User2ID:id
                };
                axios.create({withCredentials: true}).delete("http://localhost:8000/api/user/friends",{data: friendshipDto})
                    .then(res => {
                        this.updateFriends()
                    })
            }
        },
    }
</script>

<style scoped>

</style>