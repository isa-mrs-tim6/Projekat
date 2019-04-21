<template>
    <div>
        <v-layout column align-center>
            <h1>{{name}}</h1>
            <v-rating
                    v-model="rating"
                    background-color="indigo lighten-3"
                    empty-icon="$vuetify.icons.ratingFull"
                    readonly
                    half-increments
                    color="indigo"
                    size="32"
            ></v-rating>
            <RoomRatings v-bind:rooms="rooms"></RoomRatings>
        </v-layout>
    </div>
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