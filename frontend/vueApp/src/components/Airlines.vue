<template xmlns:v-slot="http://www.w3.org/1999/XSL/Transform">
    <div>
        <v-data-table
                :headers="headers"
                :items="airlines"
                class="elevation-1"
        >
            <template v-slot:items="airlines">
                <td>{{airlines.item.Name}}</td>
                <td>{{airlines.item.Address}}</td>
                <td>{{airlines.item.Longitude}}</td>
                <td>{{airlines.item.Latitude}}</td>
                <td>{{airlines.item.Promo}}</td>
            </template>
        </v-data-table>
    </div>
</template>

<script>
    export default {
        name: "Airlines",
        props: ["airlines"],
        data() {
            return {
                headers: [
                    { text: 'Name', value: 'Name' },
                    { text: 'Address', value: 'Address' },
                    { text: 'Longitude', value: 'Longitude' },
                    { text: 'Latitude', value: 'Latitude' },
                    { text: 'Promo', value: 'Promo' },
                ],
            }
        },
    }
</script>

<style scoped>

</style>