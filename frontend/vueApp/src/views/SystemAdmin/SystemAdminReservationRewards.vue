<template>
    <div id="main">
        <SystemAdminNavDrawer/>
        <ManageReservationRewards/>
    </div>
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