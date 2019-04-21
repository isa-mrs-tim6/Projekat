<template xmlns:v-slot="http://www.w3.org/1999/XSL/Transform">
    <div>
        <v-data-table
                :headers="headers"
                :items="rooms"
                class="elevation-1"
                item-key="ID"
        >
            <template v-slot:items="rooms">
                <td>{{rooms.item.Room.Number}}</td>
                <td>
                    <v-rating
                            v-model="rooms.item.Rating"
                            background-color="indigo lighten-3"
                            empty-icon="$vuetify.icons.ratingFull"
                            half-increments
                            readonly
                            color="indigo"
                            size="16"
                    ></v-rating>
                </td>
            </template>
        </v-data-table>
    </div>
</template>

<script>
    export default {
        name: "RoomRatings",
        props: ["rooms"],
        data() {
            return {
                headers: [
                    { text: 'Number', value: 'Number' },
                    { text: 'Rating', value: 'Rating' },
                ],
            }
        },
    }
</script>

<style scoped>

</style>