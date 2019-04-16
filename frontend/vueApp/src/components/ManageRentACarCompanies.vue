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
        <RentACarCompanies v-bind:rentACarCompanies="RentACarCompanies"/>
    </div>
</template>

<script>
    import RentACarCompanies from "../components/RentACarCompanies";
    import axios from 'axios';

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

</style>