<template xmlns:v-slot="http://www.w3.org/1999/XSL/Transform">
    <v-container grid-list-xl text-xs-center fill-height>
        <v-layout column align-center fill-height>
            <v-data-table
                    :headers="headers"
                    :items="rooms"
                    class="elevation-1"
                    item-key="ID"
            >
                <template v-slot:items="rooms">
                    <td class="text-xs-center body-2">{{rooms.item.Room.Number}}</td>
                    <td class="text-xs-center">
                        <v-rating
                                v-model="rooms.item.Rating"
                                background-color="indigo lighten-3"
                                empty-icon="$vuetify.icons.ratingFull"
                                half-increments
                                readonly
                                color="orange"
                                size="32"
                        ></v-rating>
                    </td>
                </template>
            </v-data-table>
        </v-layout>
    </v-container>
</template>

<script>
    export default {
        name: "RoomRatings",
        props: ["rooms"],
        data() {
            return {
                headers: [
                    { text: 'Number', value: 'Number', align: 'center', width: '50%' },
                    { text: 'Rating', value: 'Rating', align: 'center', width: '50%' },
                ],
            }
        },
    }
</script>

<style scoped>

</style>