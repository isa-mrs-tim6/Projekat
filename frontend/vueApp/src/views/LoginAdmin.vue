<template>
    <v-container fluid fill-height>
        <v-layout align-center justify-center>
            <v-flex xs12 sm8 md4>
                <v-card class="elevation-12">
                    <v-toolbar dark>
                        <v-toolbar-title>Admin login</v-toolbar-title>
                    </v-toolbar>
                    <v-card-text>
                        <v-form>
                            <v-select
                                    :items="items"
                                    v-model="Type"
                                    prepend-icon="account_circle"
                                    label="Account type"
                            ></v-select>
                            <v-text-field
                                    prepend-icon="person"
                                    v-model="Email"
                                    label="Email"
                                    required
                            ></v-text-field>
                            <v-text-field
                                    prepend-icon="lock"
                                    v-model="Password"
                                    label="Password"
                                    required
                                    type="password"
                            ></v-text-field>
                        </v-form>
                        <div v-if="notActivated">
                            <v-btn block @click="resend">Resend verification email</v-btn>
                        </div>
                    </v-card-text>
                    <v-card-actions>
                        <v-btn dark @click="loginAdmin">Login</v-btn>
                        <v-spacer></v-spacer>
                        <v-btn @click="goBack">Back</v-btn>
                    </v-card-actions>
                </v-card>
            </v-flex>
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
        </v-layout>
    </v-container>
</template>

<script>
    import axios from 'axios';
    export default {
        name: "LoginAdmin",
        data() {
            return {
                Email:"",
                Password:"",
                Type:"",
                items: [
                    "System Admin",
                    "Airline Admin",
                    "Hotel Admin",
                    "Rent-A-Car Admin"
                ],
                notActivated: false,
                snackbar_fields: {
                    snackbar: false,
                    text: null,
                    color: null,
                },
            }
        },
        methods:{
            loginAdmin(e){
                e.preventDefault();
                const credentials = {
                    Email: this.Email,
                    Password: this.Password,
                    Type: this.Type.split(' ').join(''),
                };
                axios.create({withCredentials: true}).post("http://ec2-35-159-21-254.eu-central-1.compute.amazonaws.com:8000/api/admin/login", credentials)
                    .then(
                        res => {
                            switch (res.data) {
                                case "AirlineAdmin":
                                    this.$router.replace("../../airlineAdmin/admin_profile");
                                    break;
                                case "HotelAdmin":
                                    this.$router.replace("../../hotelAdmin/admin_profile");
                                    break;
                                case "Rent-A-CarAdmin":
                                    this.$router.replace("../../racAdmin/admin_profile");
                                    break;
                                case "SystemAdmin":
                                    this.$router.replace("../../systemAdmin/admin_profile");
                                    break;
                                default:
                                    alert("Invalid login info");
                                    break;
                            }

                        }
                    )
                    .catch( err => {
                        if (err.response.status === 400) {
                            this.notActivated = true;
                            this.snackbar_fields.text = "Account not verified";
                            this.snackbar_fields.color = "error";
                            this.snackbar_fields.snackbar = true;
                        } else {
                            this.snackbar_fields.text = "Wrong login info";
                            this.snackbar_fields.color = "error";
                            this.snackbar_fields.snackbar = true;
                        }
                    })
            },
            goBack(e){
                e.preventDefault();
                this.$router.replace("login");
            },
            resend(){
                this.notActivated = false;
                const credentials = {
                    Email: this.Email,
                    Password: this.Password,
                    Type: this.Type.split(' ').join(''),
                };
                axios.post("http://ec2-35-159-21-254.eu-central-1.compute.amazonaws.com:8000/api/mail/resend", credentials)
                    .then(res => {
                        this.snackbar_fields.text = "Email resent";
                        this.snackbar_fields.color = "success";
                        this.snackbar_fields.snackbar = true;
                    })
                    .catch(err => {
                        this.snackbar_fields.text = "Account not found";
                        this.snackbar_fields.color = "error";
                        this.snackbar_fields.snackbar = true;
                    })
            }
        }
    }
</script>

<style scoped>

</style>