<template>
    <v-card>
        <v-card-title>
            Hotels
            <v-spacer></v-spacer>
            <v-text-field
                    v-model="search"
                    append-icon="search"
                    label="Search"
                    single-line
                    hide-details
            ></v-text-field>
        </v-card-title>
        <v-data-table
            :headers="headers"
            :items="hotels"
            :search="search"
            class="elevation-1"
        >
            <template v-slot:items="hotels">
                <td>{{hotels.item.Name}}</td>
                <td>{{hotels.item.Address}}</td>
                <td>{{hotels.item.Longitude}}</td>
                <td>{{hotels.item.Latitude}}</td>
                <td>{{hotels.item.Description}}</td>
            </template>
            <template v-slot:no-results>
                <v-alert :value="true" color="error" icon="warning">
                    Your search for "{{ search }}" found no results.
                </v-alert>
            </template>
        </v-data-table>
    </v-card>
</template>

<script>
    export default {
        name: "Hotels",
        props: ["hotels"],
        data() {
            return {
                search: '',
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