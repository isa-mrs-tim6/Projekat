<template>
    <v-container fluid fill-height>
        <user-nav-bar></user-nav-bar>
        <v-layout align-center justify-center>
            <v-flex xs12 sm8 md4>
                <v-card class="elevation-12">
                    <v-toolbar dark>
                        <v-toolbar-title>Login</v-toolbar-title>
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
                    <div v-if="notActivated">
                        <v-btn block @click="resend">Resend verification email</v-btn>
                    </div>
                    <v-card-actions>
                        <v-btn dark @click="loginUser">Login</v-btn>
                        <v-spacer></v-spacer>
                        <v-btn @click="admin">Log in as admin</v-btn>
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
    import UserNavBar from "../components/User/UserNavBar";
    export default {
        name: "LoginUser",
        components: {UserNavBar},
        data() {
            return {
                Email:"",
                Password:"",
                notActivated: false,
                snackbar_fields: {
                    snackbar: false,
                    text: null,
                    color: null,
                },
                isInv: this.$route.query.inv
            }
        },
        methods:{
            loginUser(e){
                e.preventDefault();
                let creds={
                    Email: this.Email,
                    Password: this.Password
                };
                axios.create({withCredentials: true}).post("http://localhost:8000/api/user/login", creds)
                    .then(
                        res => {

                            if (this.isInv){
                                this.$router.push("/userReservations")
                            }else{
                                this.$router.push("/user_flights")
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
            admin(e){
                e.preventDefault();
                this.$router.replace("admin_login");
            },
            resend(){
                this.notActivated = false;
                const credentials = {
                    Email: this.Email,
                    Password: this.Password,
                    Type: "User",
                };
                axios.post("http://localhost:8000/api/mail/resend", credentials)
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