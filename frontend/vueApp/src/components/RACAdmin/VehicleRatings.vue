<template xmlns:v-slot="http://www.w3.org/1999/XSL/Transform">
    <v-container>
        <v-data-table
                :headers="headers"
                :items="vehicles"
                class="elevation-1"
                item-key="ID"
        >
            <template v-slot:items="vehicles">
                <td class="text-xs-center body-2">{{vehicles.item.Vehicle.Name}}</td>
                <td class="text-xs-center">
                    <v-rating
                            v-model="vehicles.item.Rating"
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
    </v-container>
</template>

<script>
    export default {
        name: "VehicleRatings",
        props: ["vehicles"],
        data() {
            return {
                headers: [
                    { text: 'Name', value: 'Name', align: 'center', width: '50%' },
                    { text: 'Rating', value: 'Rating', align: 'center', width: '50%' },
                ],
            }
        },
    }
</script>

<style scoped>

</style>