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
                    <v-icon
                            small
                            @click="editQuickReservations(rooms.item)"
                    >
                        edit
                    </v-icon>
                </td>
            </template>
        </v-data-table>
        <v-dialog v-model="dialog" max-width="500px">
            <template v-slot:activator="{ on }">
                <v-btn color="primary" dark class="mb-2" v-on="on">New Item</v-btn>
            </template>
            <v-card>
                <v-card-title>
                    <span class="headline">Manage quick reservations</span>
                </v-card-title>
                <v-card-text>
                    <v-container grid-list-md>
                        <v-layout wrap>
                            <v-flex xs12 v-for="(value, index) in editedItem.QuickReserveDays" v-bind:key="index">
                                <v-layout align-center justify-center row-wrap>
                                    <v-flex xs6>
                                        <v-menu v-model="menuFrom[index]" :close-on-content-click="false" lazy transition="scale-transition"
                                                :nudge-right="40" offset-y full-width max-width="290px" min-width="290px">
                                            <template v-slot:activator="{ on }">
                                                <v-text-field v-model="value.Start" label="Start" prepend-icon="event" readonly v-on="on"></v-text-field>
                                            </template>
                                            <v-date-picker no-title scrollable v-model = "value.Start" @input="menuFrom[index] = false"></v-date-picker>
                                        </v-menu>
                                    </v-flex>
                                    <v-flex xs6>
                                        <v-menu v-model="menuTo[index]" :close-on-content-click="false" lazy transition="scale-transition"
                                                :nudge-right="40" offset-y full-width max-width="290px" min-width="290px">
                                            <template v-slot:activator="{ on }">
                                                <v-text-field v-model="value.End" label="End" prepend-icon="event" readonly v-on="on"></v-text-field>
                                            </template>
                                            <v-date-picker no-title scrollable v-model = "value.End " @input="menuTo[index] = false"></v-date-picker>
                                        </v-menu>
                                    </v-flex>
                                </v-layout>
                            </v-flex>
                        </v-layout>
                    </v-container>
                </v-card-text>
                <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn color="blue darken-1" flat @click="addNewQuickReserve">Add new dates</v-btn>
                    <v-btn color="blue darken-1" flat @click="closeQuickReserve">Cancel</v-btn>
                    <v-btn color="blue darken-1" flat @click="saveQuickReserve">Save</v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>
    </div>
</template>

<script>
    import moment from 'moment';

    export default {
        name: "Rooms",
        props: ["rooms"],
        data() {
            return {
                headers: [
                    { text: 'Number', value: 'Number' },
                    { text: 'Capacity', value: 'Capacity' },
                    { text: 'Price', value: 'Price' },
                    { text: 'Quick reservation' },
                ],
                selected: [],
                numeric: value => !isNaN(value) || 'Numeric',
                editNumber: null,
                editPrice: null,
                editCapacity: null,
                dialog: false,
                editedItem: {QuickReserveDays: []},
                menuFrom: [],
                menuTo: [],
                time: {
                    from: null,
                    to: null,
                },
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
            },
            editQuickReservations(item) {
                this.editedItem = item;
                this.menuFrom = [];
                this.menuTo = [];

                for (let i = 0; i < item.QuickReserveDays.length; i++) {
                    this.menuFrom.push(false);
                    this.menuTo.push(false);
                }

                this.dialog = true;
            },
            addNewQuickReserve() {

            },
            closeQuickReserve () {
                this.dialog = false;
            },
            saveQuickReserve() {

            },
        }
    }
</script>

<style scoped>

</style>