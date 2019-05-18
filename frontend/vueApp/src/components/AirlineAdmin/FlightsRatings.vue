<template>
    <v-container>
        <v-data-table
                :headers="headers"
                :items="flights"
                class="elevation-1"
                item-key="ID"
        >
            <template v-slot:items="flights">
                <td class="text-xs-center body-2">{{flights.item.Flight.Origin.Name}}</td>
                <td class="text-xs-center body-2">{{flights.item.Flight.Destination.Name}}</td>
                <td class="text-xs-center body-2">{{flights.item.Flight.Departure | moment}}</td>
                <td class="text-xs-center">
                    <v-rating
                            v-model="flights.item.Rating"
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
        name: "FlightsRatings",
        props: ["flights"],
        data() {
            return {
                headers: [
                    { text: 'From', value: 'From', align: 'center', width: '20%' },
                    { text: 'To', value: 'To', align: 'center', width: '20%' },
                    { text: 'Date', value: 'Date', align: 'center', width: '20%' },
                    { text: 'Rating', value: 'Rating', align: 'center', width: '40%' },
                ],
            }
        },
    }
</script>

<style scoped>

</style>