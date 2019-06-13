<template>
    <v-container fill-height fluid id="main">
        <v-layout>
            <v-flex>
                <SystemAdminNavDrawer/>
            </v-flex>
            <v-container>
                <v-layout>
                    <v-flex align-self-center>
                        <ManageReservationRewards/>
                    </v-flex>
                </v-layout>
            </v-container>
        </v-layout>
    </v-container>
</template>

<script>
    import SystemAdminNavDrawer from "../../components/SystemAdmin/SystemAdminNavDrawerr";
    import ManageReservationRewards from "../../components/SystemAdmin/ManageReservationRewards";
    export default {
        name: "SystemAdminReservationRewards",
        components: {ManageReservationRewards, SystemAdminNavDrawer}
    }
</script>

<style scoped>
    @import '../../assets/css/SystemAdmin.css';
</style>