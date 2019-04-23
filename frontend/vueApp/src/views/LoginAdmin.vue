<template>
    <v-container fluid fill-height>
        <v-layout align-center justify-center>
            <v-flex xs12 sm8 md4>
                <v-card class="elevation-12">
                    <v-toolbar dark color="primary">
                        <v-toolbar-title>Admin login</v-toolbar-title>
                    </v-toolbar>
                    <v-card-text>
                        <v-form>
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
                    </v-card-text>
                    <v-card-actions>
                        <v-btn color="primary" @click="loginAdmin">Login</v-btn>
                        <v-spacer></v-spacer>
                        <v-btn @click="goBack">Back</v-btn>
                    </v-card-actions>
                </v-card>
            </v-flex>
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
                Password:""
            }
        },
        methods:{
            loginAdmin(e){
                e.preventDefault();
                let creds={
                    Email: this.Email,
                    Password: this.Password
                };
                axios.create({withCredentials: true}).post("http://localhost:8000/api/admin/login", creds)
                    .then(
                        res => {
                            switch (res.data) {
                                case "AirlineAdmin":
                                    this.$router.replace("../../airlineAdmin");
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
                    .catch(
                        err => alert("Invalid login info")
                    );
            },
            goBack(e){
                e.preventDefault();
                this.$router.replace("login");
            }
        }
    }
</script>

<style scoped>

</style>