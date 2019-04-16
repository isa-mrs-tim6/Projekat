<template>
    <div>
        <v-toolbar>
            <v-toolbar-side-icon></v-toolbar-side-icon>
            <v-toolbar-title>System Admin Panel</v-toolbar-title>
            <v-spacer></v-spacer>
            <v-toolbar-items class="hidden-sm-and-down">
                <v-btn @click="manageHotelProfile" flat>Hotel Profile</v-btn>
                <v-btn @click="manageProfile" flat>Admin Profile</v-btn>
            </v-toolbar-items>
        </v-toolbar>
        <component v-bind:is="currentTabComponent"></component>
    </div>
</template>

<script>
    import ManageHotelProfile from "../components/ManageHotelProfile";
    import AdminProfile from "../components/AdminProfile";

    export default {
        name: "HotelAdmin",
        components: {ManageHotelProfile, AdminProfile},
        data() {
            return {
                currentTabComponent: "ManageHotelProfile",
            }
        },
        methods: {
            manageHotelProfile(){
                this.currentTabComponent = "ManageHotelProfile"
            },
            manageProfile(){
                this.currentTabComponent = "AdminProfile"
            },
        },
    }
</script>

<style scoped>

</style>