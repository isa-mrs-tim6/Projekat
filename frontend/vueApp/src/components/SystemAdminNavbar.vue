<template>
    <v-toolbar>
        <v-toolbar-side-icon></v-toolbar-side-icon>
        <v-toolbar-title>Title</v-toolbar-title>
        <v-spacer></v-spacer>
        <v-toolbar-items class="hidden-sm-and-down">
            <v-btn flat>Manage admin accounts</v-btn>
            <v-btn flat>Airlines</v-btn>
            <v-btn flat>Hotels</v-btn>
            <v-btn flat>Rent-a-car companies</v-btn>
        </v-toolbar-items>
    </v-toolbar>
</template>

<script>
    export default {
        name: "SystemAdminNavbar"
    }
</script>

<style scoped>

</style>