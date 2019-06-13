<template>
    <v-container grid-list-xl text-xs-center fill-height>
        <v-layout align-center justify-center column wrap fill-height>
            <v-flex style="width: 60vw">
                <v-card class="flexcard">
                    <v-card-title primary-title>
                        <div class="headline font-weight-medium">Register new admin accounts</div>
                    </v-card-title>
                    <v-card-text class="grow">
                        <v-form ref="form" class="align-center justify-center">
                            <v-text-field
                                    v-model="email"
                                    :rules="emailRules"
                                    label="E-mail"
                                    class="body-2"
                                    prepend-icon="email"
                                    required
                            ></v-text-field>
                            <v-text-field
                                    v-model="password"
                                    :append-icon="show1 ? 'visibility' : 'visibility_off'"
                                    :rules="[rules.required, rules.min]"
                                    :type="show1 ? 'text' : 'password'"
                                    name="input-10-1"
                                    label="Password"
                                    prepend-icon="lock"
                                    hint="At least 8 characters"
                                    class="body-2"
                                    counter
                                    @click:append="show1 = !show1"
                            ></v-text-field>
                            <v-select
                                    v-model="select"
                                    :items="items"
                                    label="Account type"
                                    prepend-icon="person"
                                    class="body-2"
                                    required
                            ></v-select>
                            <template>
                                <v-select
                                        v-model="selectCompany"
                                        :items="airlines"
                                        label="Company"
                                        prepend-icon="work"
                                        class="body-2"
                                        required
                                        item-text="Name"
                                        item-value=item
                                        return-object
                                        v-if="select === 'Airline Admin'">
                                </v-select>
                                <v-select
                                        v-model="selectCompany"
                                        :items="hotels"
                                        label="Company"
                                        prepend-icon="work"
                                        class="body-2"
                                        required
                                        item-text="Name"
                                        item-value=item
                                        return-object
                                        v-if="select === 'HotelAdmin Admin'">
                                </v-select>
                                <v-select
                                        v-model="selectCompany"
                                        :items="rac_companies"
                                        label="Company"
                                        prepend-icon="work"
                                        class="body-2"
                                        required
                                        item-text="Name"
                                        item-value=item
                                        return-object
                                        v-if="select === 'Rent-A-Car Admin'">
                                </v-select>
                            </template>
                        </v-form>
                    </v-card-text>
                    <v-card-actions>
                        <v-btn color="primary" :loading="loading" :disabled="loading" @click="register">submit</v-btn>
                        <v-btn @click="clear">clear</v-btn>
                    </v-card-actions>
                </v-card>
            </v-flex>
        </v-layout>
        <v-snackbar
                v-model="snackbar_fields.snackbar"
                :timeout="4000"
                :top="true"
                :color="snackbar_fields.color"
        >
            {{ snackbar_fields.text }}
            <v-btn
                    flat
                    @click="snackbar_fields.snackbar = false"
            >
                Close
            </v-btn>
        </v-snackbar>
    </v-container>
</template>

<script>
    import axios from 'axios/index';
    export default {
        name: "Registration",
        data(){
            return{
                email:"",
                emailRules: [
                    v => !!v || 'E-mail is required',
                    v => /.+@.+/.test(v) || 'E-mail must be valid'
                ],
                password:"",
                show1: false,
                rules: {
                    required: value => !!value || 'Required.',
                    min: v => v.length >= 8 || 'Min 8 characters',
                },
                select: null,
                items: [
                    'System Admin',
                    'Airline Admin',
                    'HotelAdmin Admin',
                    'Rent-A-Car Admin'
                ],
                loader: null,
                loading: false,
                selectCompany: null,
                airlines: [],
                hotels: [],
                rac_companies: [],
                snackbar_fields: {
                    snackbar: false,
                    text: null,
                    color: null,
                },
            }
        },
        mounted() {
            axios.get('http://localhost:8000/api/airline/getAirlines')
                .then(res => this.airlines = res.data)
                .catch(err => {
                    this.snackbar_fields.text = "Could not retrieve an airline list";
                    this.snackbar_fields.color = "error";
                    this.snackbar_fields.snackbar = true;
                    this.loader = null;
                });
            axios.get('http://localhost:8000/api/hotel/getHotels')
                .then(res => this.hotels = res.data)
                .catch(err => {
                this.snackbar_fields.text = "Could not retrieve a list of hotels";
                this.snackbar_fields.color = "error";
                this.snackbar_fields.snackbar = true;
                this.loader = null;
                });
            axios.get('http://localhost:8000/api/rentACarCompany/getRentACarCompanies')
                .then(res => this.rac_companies = res.data)
                .catch(err => {
                    this.snackbar_fields.text = "Could not retrieve a list of rent-a-car companies";
                    this.snackbar_fields.color = "error";
                    this.snackbar_fields.snackbar = true;
                    this.loader = null;
                });
        },
        methods: {
            register (e){
                e.preventDefault();

                if (this.email === "" || this.password === "" || this.select === null) {
                    alert("Please fill in the required fields");
                    return;
                }
                const credentials = {
                    Email: this.email,
                    Password: this.password,
                    accountType: this.select,
                    CompanyID: this.selectCompany === null ? 0 : this.selectCompany.ID,
                };

                this.loader = "loading";
                axios.create({withCredentials: true}).post('http://localhost:8000/api/admin/register', credentials)
                    .then(res => {
                        if (res.status === 201) {
                            this.snackbar_fields.text = "Registration successful";
                            this.snackbar_fields.color = "success";
                            this.snackbar_fields.snackbar = true;
                            this.loader = null;
                        }
                    })
                    .catch(err => {
                        this.snackbar_fields.text = "Registration unsuccessful";
                        this.snackbar_fields.color = "error";
                        this.snackbar_fields.snackbar = true;
                        this.loader = null;
                    });
            },
            clear (){
                this.email = '';
                this.password = '';
                this.select = null;
            },
        },
        watch: {
            loader (){
                const l = this.loader;
                this[l] = !this[l];
                setTimeout(() => (this[l] = false), 2000);
                this.loader = null
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