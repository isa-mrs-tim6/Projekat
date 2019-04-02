<template>
    <v-toolbar>
        <v-toolbar-side-icon></v-toolbar-side-icon>
        <v-toolbar-title>Airline Admin</v-toolbar-title>
        <v-spacer></v-spacer>
        <v-toolbar-items class="hidden-sm-and-down">
            <v-btn flat>Airline Company</v-btn>
            <v-btn flat>Profile</v-btn>
        </v-toolbar-items>
    </v-toolbar>
</template>

<script>
    export default {
        name: "AirlineAdminNavbar"
    }
</script>

<style scoped>

</style>