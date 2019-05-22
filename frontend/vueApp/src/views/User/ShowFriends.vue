<template>
    <div>
        <UserNavBar></UserNavBar>
        <v-layout align-center justify-center style="margin-top: 100px">
            <v-card style="width: 60%;">
                <v-card-title style="padding: 0">
                    <UserProfileNavBar></UserProfileNavBar>
                </v-card-title>
                <v-card-text>
                    <UserFriends></UserFriends>
                </v-card-text>
            </v-card>
        </v-layout>
    </div>
</template>

<script>
    import UserNavBar from "../../components/User/UserNavBar";
    import UserProfileNavBar from "../../components/User/UserProfileNavBar";
    import UserReservations from "../../components/User/UserReservations";
    import UserFriends from "../../components/User/UserFriends";
    export default {
        name: "ShowFriends",
        components: {UserFriends, UserProfileNavBar, UserNavBar}
    }
</script>

<style scoped>

</style>