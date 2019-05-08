<template>
    <v-container grid-list-xl text-xs-center style="height: 100vh;">
        <v-layout column align-center fill-height>
            <v-card>
                <v-card-title>
                    <h1 class="display-2">{{name}}</h1>
                </v-card-title>
            </v-card>
            <v-rating
                    v-model="rating"
                    background-color="indigo lighten-3"
                    empty-icon="$vuetify.icons.ratingFull"
                    readonly
                    half-increments
                    color="orange"
                    size="48"
            ></v-rating>
            <vehicle-ratings style="width: 75vw" v-bind:vehicles="vehicles"></vehicle-ratings>
        </v-layout>
    </v-container>
</template>

<script>


    import VehicleRatings from "./VehicleRatings";
    export default {
        name: "RacRatings",
        components: {VehicleRatings},
        props: ["vehicles", "name", "rating"],
    }
</script>

<style scoped>
</style>