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
        <Rooms v-bind:rooms="Rooms"/>
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
                this.clear();
            },
            clear() {
                this.$refs.form.reset();
            }
        }
    }
</script>

<style scoped>

</style>