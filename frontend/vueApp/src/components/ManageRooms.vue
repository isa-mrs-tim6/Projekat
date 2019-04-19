<template>
    <div>
        <v-form ref="form">
            <v-text-field
                    v-model="Capacity"
                    label="Capacity"
                    :rules="[rules.required, rules.numeric]"
                    required
            ></v-text-field>
            <v-text-field
                    v-model="Price"
                    label="Price"
                    :rules="[rules.required, rules.numeric]"
                    required
            ></v-text-field>
            <v-text-field
                    v-model="AddNum"
                    label="Number of rooms"
                    :rules="[rules.required, rules.numeric]"
                    required
            ></v-text-field>
            <v-btn @click="addRooms">submit</v-btn>
            <v-btn @click="clear">clear</v-btn>
        </v-form>
        <Rooms @update-room="updateRoom" ref="roomReference" v-bind:rooms="Rooms"/>
        <v-btn @click="deleteRooms">delete</v-btn>
    </div>
</template>

<script>
    import Rooms from "./Rooms";
    import axios from 'axios';

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
            axios.create({withCredentials: true}).get('http://localhost:8000/api/hotel/getRooms')
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
                axios.create({withCredentials: true}).post('http://localhost:8000/api/hotel/addRooms', rooms)
                    .then(res =>
                        axios.create({withCredentials: true}).get('http://localhost:8000/api/hotel/getRooms')
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
                axios.create({withCredentials: true}).post('http://localhost:8000/api/hotel/deleteRooms', deletedRooms)
                    .then(res =>
                        axios.create({withCredentials: true}).get('http://localhost:8000/api/hotel/getRooms')
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
                axios.create({withCredentials: true}).post('http://localhost:8000/api/hotel/updateRoom', room)
                    .then(res => axios.create({withCredentials: true}).get('http://localhost:8000/api/hotel/getRooms')
                        .then(res => this.Rooms = res.data)
                        .catch(err => alert("Could not retrieve hotel rooms"))
                        .catch(err => alert("Error updating room")));
            }
        }
    }
</script>

<style scoped>

</style>