<template xmlns:v-slot="http://www.w3.org/1999/XSL/Transform">
    <div>
        <v-data-table
                :headers="headers"
                :items="rentACarCompanies"
                class="elevation-1"
        >
            <template v-slot:items="rentACarCompanies">
                <td>{{rentACarCompanies.item.Name}}</td>
                <td>{{rentACarCompanies.item.Address}}</td>
                <td>{{rentACarCompanies.item.Longitude}}</td>
                <td>{{rentACarCompanies.item.Latitude}}</td>
                <td>{{rentACarCompanies.item.Promo}}</td>
            </template>
        </v-data-table>
    </div>
</template>

<script>
    export default {
        name: "RentACarCompanies",
        props: ["rentACarCompanies"],
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