<template>
    <div>
        <v-data-table
            :headers="headers"
            :items="hotels"
            class="elevation-1"
        >
            <template v-slot:items="hotels">
                <td>{{hotels.item.Name}}</td>
                <td>{{hotels.item.Address}}</td>
                <td>{{hotels.item.Longitude}}</td>
                <td>{{hotels.item.Latitude}}</td>
                <td>{{hotels.item.Description}}</td>
            </template>
        </v-data-table>
    </div>
</template>

<script>
    export default {
        name: "Hotels",
        props: ["hotels"],
        data() {
            return {
                headers: [
                    { text: 'Name', value: 'Name' },
                    { text: 'Address', value: 'Address' },
                    { text: 'Longitude', value: 'Longitude' },
                    { text: 'Latitude', value: 'Latitude' },
                    { text: 'Description', value: 'Description' },
                ],
            }
        },
    }
</script>

<style scoped>

</style>