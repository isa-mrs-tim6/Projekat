<template>
    <v-container grid-list-xl text-xs-center>
        <v-layout align-center justify-center column wrap fill-height>
            <v-flex style="width: 75vw">
                <Hotels style="margin-bottom: 2vh" v-bind:hotels="hotels"/>
                <v-flex style="width: 35vw; margin-left: 20vw">
                    <v-form ref="form" class="align-center justify-center">
                        <v-text-field
                                v-model="Name"
                                label="Name"
                                :rules="[rules.required]"
                                required
                        ></v-text-field>
                        <v-text-field
                                v-model="Address"
                                label="Address"
                                :rules="[rules.required]"
                                required
                        ></v-text-field>
                        <v-text-field
                                v-model="Latitude"
                                label="Latitude"
                                :rules="[rules.required, rules.numeric]"
                                required
                        ></v-text-field>
                        <v-text-field
                                v-model="Longitude"
                                label="Longitude"
                                :rules="[rules.required, rules.numeric]"
                                required
                        ></v-text-field>
                        <v-text-field
                                v-model="Description"
                                label="Description"
                                :rules="[rules.required]"
                                required
                        ></v-text-field>

                        <v-btn style="float: left" @click="addHotel">submit</v-btn>
                        <v-btn style="float: right" @click="clear">clear</v-btn>
                    </v-form>
                </v-flex>
            </v-flex>
        </v-layout>
    </v-container>
</template>

<script>
    import Hotels from "../Hotels";
    import axios from 'axios/index';

    export default {
        name: "ManageHotels",
        components: {Hotels},
        data() {
            return {
                hotels: [],
                Name: '',
                Address: '',
                Longitude: '',
                Latitude: '',
                Description: '',
                rules: {
                    required: value => !!value || 'Required.',
                    numeric: value => !isNaN(value) || 'Numeric',
                },
            }
        },
        created() {
            axios.create({withCredentials: true}).get('http://localhost:8000/api/hotel/getHotels')
                .then(res => this.hotels = res.data)
                .catch(err => console.log(err));
        },
        methods: {
            addHotel(e) {
                e.preventDefault();
                const newHotel = {
                    'Name': this.Name,
                    'Address': this.Address,
                    'Latitude': parseFloat(this.Latitude),
                    'Longitude': parseFloat(this.Longitude),
                    'Description': this.Description,
                };
                this.clear();
                axios.create({withCredentials: true}).post('http://localhost:8000/api/hotel/addHotel', newHotel)
                    .then(res =>
                        axios.create({withCredentials: true}).get('http://localhost:8000/api/hotel/getHotels')
                            .then(res => this.hotels = res.data)
                            .catch(err => alert(err)))
                    .catch(err => alert(err));
            },
            clear() {
                this.$refs.form.reset();
            }
        }
    }
</script>

<style scoped>

</style>