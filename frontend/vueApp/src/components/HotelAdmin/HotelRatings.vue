<template>
    <v-container>
        <v-layout column align-center>
            <h1>{{name}}</h1>
            <v-rating
                    v-model="rating"
                    background-color="indigo lighten-3"
                    empty-icon="$vuetify.icons.ratingFull"
                    readonly
                    half-increments
                    color="indigo"
                    size="48"
            ></v-rating>
            <RoomRatings style="width: 75vw" v-bind:rooms="rooms"></RoomRatings>
        </v-layout>
    </v-container>
</template>

<script>
    import RoomRatings from "./RoomRatings";

    export default {
        name: "HotelRatings",
        components: {RoomRatings},
        props: ["rooms", "name", "rating"],
    }
</script>

<style scoped>
</style>