<template>
    <div>
        <v-toolbar>
            <v-toolbar-side-icon></v-toolbar-side-icon>
            <v-toolbar-title>System Admin Panel</v-toolbar-title>
            <v-spacer></v-spacer>
            <v-toolbar-items class="hidden-sm-and-down">
                <v-btn @click="manageAccounts" flat>Manage admin accounts</v-btn>
                <v-btn @click="manageAirlines" flat>Airlines</v-btn>
                <v-btn @click="manageHotels" flat>Hotels</v-btn>
                <v-btn @click="manageRentACarCompanies" flat>Rent-a-car companies</v-btn>
            </v-toolbar-items>
        </v-toolbar>
        <component v-bind:is="currentTabComponent"></component>
    </div>
</template>

<script>
    import ManageHotels from "../components/ManageHotels";
    import ManageAdminRegistration from "../components/ManageAdminRegistration";
    import ManageAirlines from "../components/ManageAirlines";
    import ManageRentACarCompanies from "../components/ManageRentACarCompanies";
    
    export default {
        name: "SystemAdmin",
        components: {ManageHotels, ManageAdminRegistration, ManageAirlines, ManageRentACarCompanies},
        data() {
            return {
                currentTabComponent: "ManageHotels",
            }
        },
        methods: {
            manageAccounts(){
                this.currentTabComponent = "ManageAdminRegistration"
            },
            manageHotels(){
                this.currentTabComponent = "ManageHotels"
            },
            manageRentACarCompanies(){
                this.currentTabComponent = "ManageRentACarCompanies"
            },
            manageAirlines() {
                this.currentTabComponent = "ManageAirlines"
            }
        },
    }
</script>

<style scoped>

</style>