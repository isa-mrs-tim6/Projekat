<template>
    <div>
        <UserNavBar></UserNavBar>
        <v-layout align-center justify-center style="margin-top: 100px">
            <v-card style="width: 80%">
                <VehiclesSearch></VehiclesSearch>
            </v-card>
        </v-layout>
    </div>
</template>

<script>
    import UserNavBar from "../../components/User/UserNavBar";
    import VehiclesSearch from "../../components/Search/VehiclesSearch";
    export default {
        name: "ShowVehicleSearch",
        components: {VehiclesSearch, UserNavBar}
    }
</script>

<style scoped>

</style>