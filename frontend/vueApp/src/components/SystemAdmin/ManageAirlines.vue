<template>
    <v-container grid-list-xl text-xs-center ill-height>
        <v-layout align-center justify-center row wrap fill-height>
                <v-flex xs12>
                    <v-card class="flexcard">
                        <v-card-title primary-title>
                            <div class="headline font-weight-medium">Add new airlines</div>
                        </v-card-title>
                        <v-card-text class="grow">
                            <v-layout align-center justify-center row wrap fill-height>
                                <v-flex xs4>
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
                                                v-model="Promo"
                                                label="Promo"
                                                prepend-icon="message"
                                                class="body-2"
                                                :rules="[rules.required]"
                                                required
                                        ></v-text-field>
                                    </v-form>
                                </v-flex>
                                <v-flex xs8>
                                    <Airlines v-bind:airlines="Airlines"/>
                                </v-flex>
                            </v-layout>
                        </v-card-text>
                        <v-card-actions>
                            <v-btn color="primary" @click="addAirline">submit</v-btn>
                            <v-btn @click="clear">clear</v-btn>
                        </v-card-actions>
                    </v-card>
                </v-flex>
        </v-layout>
    </v-container>
</template>

<script>
    import Airlines from "../Airlines";
    import axios from 'axios/index';
    import VueGoogleAutocomplete from 'vue-google-autocomplete'

    export default {
        name: "ManageAirlines",
        components: {Airlines},
        data() {
            return {
                Airlines: [],
                Name: '',
                Address: '',
                Longitude: '',
                Latitude: '',
                Promo: '',
                rules: {
                    required: value => !!value || 'Required.',
                    numeric: value => !isNaN(value) || 'Numeric',
                },
            }
        },
        created() {
            axios.create({withCredentials: true}).get('http://ec2-18-195-170-20.eu-central-1.compute.amazonaws.com:8000/api/airline/getAirlines')
                .then(res => this.Airlines = res.data)
                .catch(err => alert("Could not retrieve airline companies"));
        },
        methods: {
            addAirline(e) {
                e.preventDefault();

                const newAirline = {
                    'Name': this.Name,
                    'Address': this.Address,
                    'Latitude': parseFloat(this.Latitude),
                    'Longitude': parseFloat(this.Longitude),
                    'Promo': this.Promo,
                };

                if (!Object.keys(newAirline).length || isNaN(newAirline.Latitude) || isNaN(newAirline.Longitude)) {
                    alert("Please fill in all the fields properly");
                    return;
                }

                axios.create({withCredentials: true}).post('http://ec2-18-195-170-20.eu-central-1.compute.amazonaws.com:8000/api/airline/addAirline', newAirline)
                    .then(res =>
                        axios.create({withCredentials: true}).get('http://ec2-18-195-170-20.eu-central-1.compute.amazonaws.com:8000/api/airline/getAirlines')
                            .then(res => this.Airlines = res.data)
                            .catch(err => alert("Could not retrieve airline companies"))
                            .catch(err => alert("Error adding new airline company")));
                this.clear();
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