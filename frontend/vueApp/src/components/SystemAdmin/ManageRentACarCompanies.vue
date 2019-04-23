<template>
    <v-container grid-list-xl text-xs-center style="height: 100vh;">
        <v-layout align-center justify-center row wrap fill-height>
            <v-flex xs12>
                <v-card min-height="100%" class="flexcard">
                    <v-card-title primary-title>
                        <div class="headline font-weight-medium">Add new rent-a-car companies</div>
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
                                    <v-text-field
                                            v-model="Address"
                                            label="Address"
                                            prepend-icon="location_on"
                                            class="body-2"
                                            :rules="[rules.required]"
                                            required
                                    ></v-text-field>
                                    <v-flex>
                                        <v-layout row wrap fill-height>
                                            <v-text-field
                                                    v-model="Latitude"
                                                    label="Latitude"
                                                    class="body-2"
                                                    prepend-icon="satellite"
                                                    :rules="[rules.required, rules.numeric]"
                                                    required
                                            ></v-text-field>
                                            <v-text-field
                                                    v-model="Longitude"
                                                    label="Longitude"
                                                    class="body-2"
                                                    prepend-icon="satellite"
                                                    :rules="[rules.required, rules.numeric]"
                                                    required
                                            ></v-text-field>
                                        </v-layout>
                                    </v-flex>
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
                                <RentACarCompanies v-bind:rentACarCompanies="RentACarCompanies"/>
                            </v-flex>
                        </v-layout>
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
            axios.create({withCredentials: true}).get('http://localhost:8000/api/rentACarCompany/getRentACarCompanies')
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

                axios.create({withCredentials: true}).post('http://localhost:8000/api/rentACarCompany/addRentACarCompany', newRentACarCompany)
                    .then(res =>
                        axios.create({withCredentials: true}).get('http://localhost:8000/api/rentACarCompany/getRentACarCompanies')
                            .then(res => this.RentACarCompanies = res.data)
                            .catch(err => alert("Could not retrieve rent-a-car companies"))
                    .catch(err => alert("Error adding new rent-a-car company")));
                this.clear();
            },
            clear() {
                this.$refs.form.reset();
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