<template xmlns:v-slot="http://www.w3.org/1999/XSL/Transform">
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
                <td>
                    <v-edit-dialog
                            :return-value.sync="rooms.item.Number"
                            large
                            lazy
                            @save="save(rooms.item)"
                            @cancel="cancel"
                            @open="open(rooms.item)"
                            @close="close"
                    >
                        <div>{{ rooms.item.Number }}</div>
                        <template v-slot:input>
                            <div class="mt-3 title">Update price</div>
                        </template>
                        <template v-slot:input>
                            <v-text-field
                                    v-model="editNumber"
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
                            @save="save(rooms.item)"
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
                <td class="text-xs-center">
                    <v-checkbox style="padding-top: 17px" class="text-xs-center" @change="save(rooms.item)" v-model="rooms.item.QuickReserve"></v-checkbox>
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
                    { text: 'Quick reservation', value: 'QuickReserve' },
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
            save(room) {
                // Quick reservation field is changed
                if (this.editNumber === null && this.editPrice === null && this.editCapacity === null) {
                    this.$emit('update-room', room);
                    return;
                }

                // Price, Capacity or Number fields are changed
                if (isNaN(this.editPrice) || isNaN(this.editCapacity) || isNaN(this.editNumber)) {
                    alert("Fields must be numeric");
                } else {
                    this.editNumber = Number.parseInt(this.editNumber);
                    this.editPrice = Number.parseInt(this.editPrice);
                    this.editCapacity = Number.parseInt(this.editCapacity);

                    if (this.editPrice < 0 || this.editCapacity < 0 || this.editNumber < 0 ||
                        !Number.isInteger(this.editPrice) || !Number.isInteger(this.editCapacity) || !Number.isInteger(this.editNumber)) {
                        alert("Fields must be numeric");
                    } else {
                        room.Number = this.editNumber;
                        room.Price = this.editPrice;
                        room.Capacity = this.editCapacity;
                        this.$emit('update-room', room);
                    }
                }
            },
            cancel() {
                this.editNumber = null;
                this.editPrice = null;
                this.editCapacity = null;
            },
            open(room) {
                this.editNumber = Number.parseInt(room.Number);
                this.editPrice = Number.parseInt(room.Price);
                this.editCapacity = Number.parseInt(room.Capacity);
            },
            close () {
                this.editNumber = null;
                this.editPrice = null;
                this.editCapacity = null;
            }
        }
    }
</script>

<style scoped>

</style>