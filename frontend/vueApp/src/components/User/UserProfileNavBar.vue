<template>
    <v-toolbar dark>
        <v-toolbar-title>My Profile</v-toolbar-title>
        <v-spacer></v-spacer>
        <v-btn class="button" flat @click="goFriends">Friends</v-btn>
        <v-btn class="button" flat @click="goReservations">Reservations</v-btn>
        <v-btn class="button" flat @click="goProfile">Edit profile</v-btn>
    </v-toolbar>
</template>

<script>
    export default {
        name: "UserProfileNavBar",
        methods: {
            goProfile (e){
                e.preventDefault();
                this.$router.replace("userProfile");
            },
            goReservations (e){
                e.preventDefault();
                this.$router.replace("userReservations");
            },
            goFriends (e){
                e.preventDefault();
                this.$router.replace("userFriends");
            }
        }
    }
</script>

<style scoped>
    .button{
        padding: 0px 10px;
        margin: 0;
        height:100%;
    }
</style>