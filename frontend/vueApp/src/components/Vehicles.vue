<template>
    <div>
        <v-data-table
                :headers="headers"
                :items="vehicles"
                class="elevation-1"
        >
            <template v-slot:items="vehicles">
                <td>{{vehicles.item.Name}}</td>
                <td>{{vehicles.item.Capacity}}</td>
                <td>{{vehicles.item.Type}}</td>
                <td>{{vehicles.item.PricePerDay}}</td>
            </template>
        </v-data-table>
    </div>
</template>

<script>
    export default {
        name: "Vehicles",
        props: ["vehicles"],
        data() {
            return {
                headers: [
                    { text: 'Name', value: 'Name' },
                    { text: 'Capacity', value: 'Capacity' },
                    { text: 'Type', value: 'Type'},
                    { text: 'PricePerDay', value: 'PricePerDay' },
                ],
            }
        },
    }
</script>

<style scoped>

</style>