<template>
    <v-container grid-list-xl text-xs-center fill-height>
        <v-layout align-center justify-center row wrap fill-height>
            <v-flex xs12>
                <v-card class="flexcard">
                    <v-card-title primary-title>
                        <div class="headline font-weight-medium">Add new rent-a-car companies</div>
                    </v-card-title>
                    <v-card-text>
                        <v-container>
                            <v-layout align-center justify-start row wrap>
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
                                    <RentACarCompanies v-bind:rentACarCompanies="RentACarCompanies"/>
                                </v-flex>
                            </v-layout>
                        </v-container>
                    </v-card-text>
                    <v-card-actions>
                        <v-btn color="primary" @click="addRentACarCompany">submit</v-btn>
                        <v-btn @click="clear">clear</v-btn>
                    </v-card-actions>
                </v-card>
            </v-flex>
        </v-layout>
    </v-container>
</template>

<script>
    import RentACarCompanies from "../RentACarCompanies";
    import axios from 'axios/index';

    export default {
        name: "ManageRentACarCompanies",
        components: {RentACarCompanies},
        data() {
            return {
                RentACarCompanies: [],
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
            axios.create({withCredentials: true}).get('http://ec2-35-159-21-254.eu-central-1.compute.amazonaws.com:8000/api/rentACarCompany/getRentACarCompanies')
                .then(res => this.RentACarCompanies = res.data)
                .catch(err => alert("Could not retrieve rent-a-car companies"));
        },
        methods: {
            addRentACarCompany(e) {
                e.preventDefault();

                const newRentACarCompany = {
                    'Name': this.Name,
                    'Address': this.Address,
                    'Latitude': parseFloat(this.Latitude),
                    'Longitude': parseFloat(this.Longitude),
                    'Promo': this.Promo,
                };

                if (!Object.keys(newRentACarCompany).length || isNaN(newRentACarCompany.Latitude) || isNaN(newRentACarCompany.Longitude)) {
                    alert("Please fill in all the fields properly");
                    return;
                }

                axios.create({withCredentials: true}).post('http://ec2-35-159-21-254.eu-central-1.compute.amazonaws.com:8000/api/rentACarCompany/addRentACarCompany', newRentACarCompany)
                    .then(res =>
                        axios.create({withCredentials: true}).get('http://ec2-35-159-21-254.eu-central-1.compute.amazonaws.com:8000/api/rentACarCompany/getRentACarCompanies')
                            .then(res => this.RentACarCompanies = res.data)
                            .catch(err => alert("Could not retrieve rent-a-car companies"))
                    .catch(err => alert("Error adding new rent-a-car company")));
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