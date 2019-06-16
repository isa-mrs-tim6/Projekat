<template>
    <div>
        <UserNavBar></UserNavBar>
        <v-layout align-center justify-center style="margin-top: 100px">
            <v-card style="width: 80%">
                <QuickVehicles></QuickVehicles>
            </v-card>
        </v-layout>
    </div>
</template>

<script>
    import UserNavBar from "../../components/User/UserNavBar";
    import QuickVehicles from "../../components/User/QuickVehicles";
    export default {
        name: "ShowQuickVehicle",
        components: {QuickVehicles, UserNavBar}
    }
</script>

<style scoped>

</style>