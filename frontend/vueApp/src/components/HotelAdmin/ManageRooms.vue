<template>
    <v-container grid-list-xl text-xs-center fill-height>
        <v-layout align-center justify-center row wrap fill-height>
            <v-flex xs12>
                <v-card xs12 class="flexcard">
                    <v-card-title primary-title>
                        <div class="headline font-weight-medium">Add new rooms</div>
                    </v-card-title>
                    <v-card-text class="grow">
                        <v-layout align-center justify-center row wrap fill-height>
                            <v-flex xs4>
                                <v-form ref="form" class="align-center justify-center">
                                    <v-text-field
                                            v-model="Capacity"
                                            label="Capacity"
                                            prepend-icon="group"
                                            :rules="[rules.required, rules.numeric]"
                                            class="body-2"
                                            required
                                    ></v-text-field>
                                    <v-text-field
                                            v-model="Price"
                                            label="Price"
                                            prepend-icon="euro_symbol"
                                            :rules="[rules.required, rules.numeric]"
                                            class="body-2"
                                            required
                                    ></v-text-field>
                                    <v-text-field
                                            v-model="AddNum"
                                            label="Number of rooms"
                                            prepend-icon="meeting_room"
                                            :rules="[rules.required, rules.numeric]"
                                            class="body-2"
                                            required
                                    ></v-text-field>
                                </v-form>
                            </v-flex>
                            <v-flex xs8>
                                <Rooms @update-room="updateRoom" ref="roomReference" v-bind:rooms="Rooms"/>
                            </v-flex>
                        </v-layout>
                    </v-card-text>
                    <v-card-actions>
                        <v-btn color="primary" @click="addRooms">submit</v-btn>
                        <v-btn @click="clear">clear</v-btn>
                        <v-spacer/>
                        <v-btn color="error" @click="deleteRooms">delete selected</v-btn>
                    </v-card-actions>
                </v-card>
            </v-flex>
        </v-layout>
    </v-container>
</template>

<script>
    import Rooms from "./Rooms";
    import axios from 'axios/index';

    export default {
        name: "ManageRooms",
        components: {Rooms},
        data() {
            return {
                Rooms: [],
                Capacity: '',
                Price: '',
                AddNum: '',
                Selected: Rooms.methods.getSelected(),
                rules: {
                    required: value => !!value || 'Required.',
                    numeric: value => !isNaN(value) || 'Numeric',
                },
            }
        },
        created() {
            axios.create({withCredentials: true}).get('http://ec2-18-195-170-20.eu-central-1.compute.amazonaws.com:8000/api/hotel/getRooms')
                .then(res => this.Rooms = res.data)
                .catch(err => alert("Could not retrieve hotel rooms"));
        },
        methods: {
            addRooms(e) {
                e.preventDefault();
                let rooms = [];
                for (let i = 0; i < this.AddNum; i++) {
                    const newRoom = {
                        'Number': this.Rooms.length + i + 1,
                        'Capacity': parseInt(this.Capacity),
                        'Price': parseFloat(this.Price),
                    };
                    rooms.push(newRoom);
                }
                axios.create({withCredentials: true}).post('http://ec2-18-195-170-20.eu-central-1.compute.amazonaws.com:8000/api/hotel/addRooms', rooms)
                    .then(res =>
                        axios.create({withCredentials: true}).get('http://ec2-18-195-170-20.eu-central-1.compute.amazonaws.com:8000/api/hotel/getRooms')
                            .then(res => this.Rooms = res.data)
                            .catch(err => alert("Could not retrieve hotel rooms"))
                            .catch(err => alert("Error adding new rooms")));
                this.clear();            },
            clear() {
                this.$refs.form.reset();
            },
            deleteRooms(e) {
                e.preventDefault();
                let deletedRooms = this.$refs.roomReference.getSelected();
                axios.create({withCredentials: true}).post('http://ec2-18-195-170-20.eu-central-1.compute.amazonaws.com:8000/api/hotel/deleteRooms', deletedRooms)
                    .then(res =>
                        axios.create({withCredentials: true}).get('http://ec2-18-195-170-20.eu-central-1.compute.amazonaws.com:8000/api/hotel/getRooms')
                            .then(res => this.Rooms = res.data)
                            .catch(err => alert("Could not retrieve hotel rooms"))
                    .catch(err => {
                        alert(err)
                    }))
                    .catch(err => {
                        if (err.response.status === 406)
                            { alert("One of the selected rooms is already occupied") }
                        else
                            { alert("Could not remove selected rooms") }
                    });
            },
            updateRoom: function(room) {
                axios.create({withCredentials: true}).post('http://ec2-18-195-170-20.eu-central-1.compute.amazonaws.com:8000/api/hotel/updateRoom', room)
                    .then(res => axios.create({withCredentials: true}).get('http://ec2-18-195-170-20.eu-central-1.compute.amazonaws.com:8000/api/hotel/getRooms')
                        .then(res => this.Rooms = res.data)
                        .catch(err => alert("Could not retrieve hotel rooms"))
                        .catch(err => alert("Error updating room")));
            }
        }
    }
</script>

<style scoped>

</style>