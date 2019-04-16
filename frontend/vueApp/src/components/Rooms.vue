<template>
    <div>
        <v-data-table
                :headers="headers"
                :items="rooms"
                class="elevation-1"
        >
            <template v-slot:items="rooms">
                <td>{{rooms.item.Number}}</td>
                <td>{{rooms.item.Capacity}}</td>
                <td>{{rooms.item.Price}}</td>
            </template>
        </v-data-table>
    </div>
</template>

<script>
    export default {
        name: "Rooms",
        props: ["rooms"],
        data() {
            return {
                headers: [
                    { text: 'Number', value: 'Number' },
                    { text: 'Capacity', value: 'Capacity' },
                    { text: 'Price', value: 'Price' },
                ],
            }
        },
    }
</script>

<style scoped>

</style>