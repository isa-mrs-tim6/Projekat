<template>
    <div>
        <v-data-table
                v-model="selected"
                :headers="headers"
                :items="rooms"
                class="elevation-1"
                item-key="ID"
                select-all
        >
            <template v-slot:items="rooms">
                <td>
                    <v-checkbox
                            v-model="rooms.selected"
                            primary
                            hide-details
                    ></v-checkbox>
                </td>
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
                selected: []
            }
        },
        methods: {
            getSelected() {
                if (this.selected === undefined) {
                    return [];
                } else {
                    return this.selected
                }
            }
        }
    }
</script>

<style scoped>

</style>