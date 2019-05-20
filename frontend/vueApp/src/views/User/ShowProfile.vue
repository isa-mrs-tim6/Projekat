<template>
    <div>
        <UserNavBar></UserNavBar>
        <v-layout align-center justify-center style="margin-top: 100px">
            <v-card style="width: 60%;">
                <v-card-title style="padding: 0">
                    <UserProfileNavBar></UserProfileNavBar>
                </v-card-title>
                <v-card-text>
                    <UserProfile></UserProfile>
                </v-card-text>
            </v-card>
        </v-layout>
    </div>
</template>

<script>
    import UserNavBar from "../../components/User/UserNavBar";
    import UserProfile from "../../components/User/UserProfile";
    import UserProfileNavBar from "../../components/User/UserProfileNavBar";
    export default {
        name: "ShowProfile",
        components: {UserProfileNavBar, UserProfile, UserNavBar}
    }
</script>

<style scoped>

</style>