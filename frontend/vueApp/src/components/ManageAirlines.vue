<template>
    <div>
        <v-form ref="form">
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

            <v-btn @click="addRentACarCompany">submit</v-btn>
            <v-btn @click="clear">clear</v-btn>
        </v-form>
        <Airlines v-bind:airlines="Airlines"/>
    </div>
</template>

<script>
    import Airlines from "../components/Airlines";
    import axios from 'axios';

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
            addRentACarCompany(e) {
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