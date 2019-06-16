<template>
    <v-container grid-list-xl text-xs-center fill-height>
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