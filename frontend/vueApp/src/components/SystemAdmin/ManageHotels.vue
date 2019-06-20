<template>
    <v-container grid-list-xl text-xs-center fill-height>
        <v-layout align-center justify-center row wrap fill-height>
            <v-flex xs12>
                <v-card class="flexcard">
                    <v-card-title primary-title>
                        <div class="headline font-weight-medium">Add new hotels</div>
                    </v-card-title>
                    <v-card-text class="grow">
                        <v-layout align-center justify-center row wrap fill-height>
                            <v-flex>
                                <v-form ref="form" class="align-center justify-center">
                                    <v-text-field
                                            v-model="Name"
                                            label="Name"
                                            :rules="[rules.required]"
                                            prepend-icon="business"
                                            class="body-2"
                                            required
                                    ></v-text-field>
                                    <v-layout>
                                        <v-icon medium style="margin-left: 10px; margin-right: 8px">location_on</v-icon>
                                        <gmap-autocomplete
                                                style="width: 95%; border-bottom: 1px solid gray"
                                                placeholder="Address"
                                                :value="Address"
                                                @input="value = $event.target.value"
                                                @place_changed="getAddressData">
                                        </gmap-autocomplete>
                                    </v-layout>
                                    <v-text-field
                                            v-model="Description"
                                            label="Description"
                                            prepend-icon="message"
                                            class="body-2"
                                            :rules="[rules.required]"
                                            required
                                    ></v-text-field>
                                </v-form>
                            </v-flex>
                            <v-flex xs8>
                                <Hotels v-bind:hotels="hotels"/>
                            </v-flex>
                        </v-layout>
                    </v-card-text>
                    <v-card-actions>
                        <v-btn color="primary" @click="addHotel">submit</v-btn>
                        <v-btn @click="clear">clear</v-btn>
                    </v-card-actions>
                </v-card>
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
            axios.create({withCredentials: true}).get('http://ec2-18-195-170-20.eu-central-1.compute.amazonaws.com:8000/api/hotel/getHotels')
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
                axios.create({withCredentials: true}).post('http://ec2-18-195-170-20.eu-central-1.compute.amazonaws.com:8000/api/hotel/addHotel', newHotel)
                    .then(res =>
                        axios.create({withCredentials: true}).get('http://ec2-18-195-170-20.eu-central-1.compute.amazonaws.com:8000/api/hotel/getHotels')
                            .then(res => this.hotels = res.data)
                            .catch(err => alert(err)))
                    .catch(err => alert(err));
            },
            clear() {
                this.$refs.form.reset();
            },
            getAddressData: function (addressData) {
                this.Address = addressData.formatted_address;
                this.Latitude = addressData.geometry.location.lat();
                this.Longitude = addressData.geometry.location.lng();
            }
        }
    }
</script>

<style scoped>
    .flexcard {
        display: flex;
        flex-direction: column;
    }
</style>