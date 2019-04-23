<template>
    <v-container grid-list-xl text-xs-center>
        <v-layout align-center justify-center column wrap fill-height>
            <v-flex style="width: 75vw">
                <Airlines style="margin-top: 2vh" v-bind:airlines="Airlines"/>
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
                                v-model="Promo"
                                label="Promo"
                                :rules="[rules.required]"
                                required
                        ></v-text-field>

                        <v-btn style="float: left" @click="addAirline">submit</v-btn>
                        <v-btn style="float: right" @click="clear">clear</v-btn>
                    </v-form>
                </v-flex>
            </v-flex>
        </v-layout>
    </v-container>
</template>

<script>
    import Airlines from "../Airlines";
    import axios from 'axios/index';

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
            axios.create({withCredentials: true}).get('http://localhost:8000/api/airline/getAirlines')
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

                axios.create({withCredentials: true}).post('http://localhost:8000/api/airline/addAirline', newAirline)
                    .then(res =>
                        axios.create({withCredentials: true}).get('http://localhost:8000/api/airline/getAirlines')
                            .then(res => this.Airlines = res.data)
                            .catch(err => alert("Could not retrieve airline companies"))
                            .catch(err => alert("Error adding new airline company")));
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