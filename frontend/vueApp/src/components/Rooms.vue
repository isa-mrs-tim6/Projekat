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
                <td>
                    <v-edit-dialog
                            :return-value.sync="rooms.item.Capacity"
                            large
                            lazy
                            @save="save(rooms.item)"
                            @cancel="cancel"
                            @open="open(rooms.item)"
                            @close="close"
                    >
                        <div>{{ rooms.item.Capacity }}</div>
                        <template v-slot:input>
                            <div class="mt-3 title">Update capacity</div>
                        </template>
                        <template v-slot:input>
                            <v-text-field
                                    v-model="editCapacity"
                                    :rules="[numeric]"
                                    label="Edit"
                                    single-line
                                    counter
                                    autofocus
                            ></v-text-field>
                        </template>
                    </v-edit-dialog>
                </td>
                <td>
                    <v-edit-dialog
                            :return-value.sync="rooms.item.Price"
                            large
                            lazy
                            @save="save(rooms.item.ID)"
                            @cancel="cancel"
                            @open="open(rooms.item)"
                            @close="close"
                    >
                        <div>{{ rooms.item.Price }}</div>
                        <template v-slot:input>
                            <div class="mt-3 title">Update price</div>
                        </template>
                        <template v-slot:input>
                            <v-text-field
                                    v-model="editPrice"
                                    :rules="[numeric]"
                                    label="Edit"
                                    single-line
                                    counter
                                    autofocus
                            ></v-text-field>
                        </template>
                    </v-edit-dialog>
                </td>
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
                selected: [],
                numeric: value => !isNaN(value) || 'Numeric',
                editNumber: null,
                editPrice: null,
                editCapacity: null,
            }
        },
        methods: {
            getSelected() {
                if (this.selected === undefined) {
                    return [];
                } else {
                    return this.selected
                }
            },
            save (roomID) {
                if (isNaN(this.editPrice) || isNaN(this.editCapacity) || isNaN(this.editNumber) ||
                    this.editPrice < 0 || this.editCapacity < 0 || this.editNumber < 0 ||
                    Number.isInteger(this.editCapacity) || Number.isInteger(this.editNumber)) {
                    alert("Fields must be numeric");
                } else {
                    this.$emit('update-room', roomID, this.editNumber, this.editPrice, this.editCapacity);
                }
            },
            cancel () {
            },
            open (room) {
                this.editNumber = room.Number;
                this.editPrice = room.Price;
                this.editCapacity = room.Capacity;
            },
            close () {
            }
        }
    }
</script>

<style scoped>

</style>