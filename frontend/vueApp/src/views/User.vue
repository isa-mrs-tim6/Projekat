<template>
    <div>
        <v-toolbar>
            <v-toolbar-side-icon></v-toolbar-side-icon>
            <v-toolbar-title>User</v-toolbar-title>
            <v-spacer></v-spacer>
            <v-toolbar-items class="hidden-sm-and-down">
                <v-btn @click="mainUserPage" flat>Link One</v-btn>
                <v-btn flat>Link Two</v-btn>
                <v-btn @click="profileTabComponent" flat>Profile</v-btn>
            </v-toolbar-items>
        </v-toolbar>
        <component v-bind:is="currentTabComponent"></component>
    </div>
</template>

<script>
    import MainUserPage from "./MainUserPage";
    import UserProfile from "../components/UserProfile";

    export default {
        name: "User",
        components: {MainUserPage, UserProfile},
        data() {
            return {
                currentTabComponent: "MainUserPage",
            }
        },
        methods:{
            mainUserPage(){
              this.currentTabComponent = "MainUserPage"
            },
            profileTabComponent(){
                this.currentTabComponent = "UserProfile"
            }
        }
    }
</script>

<style scoped>

</style>