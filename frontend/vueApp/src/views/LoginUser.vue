<template>
    <v-container fluid fill-height>
        <v-layout align-center justify-center>
            <v-flex xs12 sm8 md4>
                <v-card class="elevation-12">
                    <v-toolbar dark color="primary">
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
                    <v-card-actions>
                        <v-btn color="primary" @click="loginUser">Login</v-btn>
                        <v-spacer></v-spacer>
                        <v-btn @click="admin">Log in as admin</v-btn>
                    </v-card-actions>
                </v-card>
            </v-flex>
        </v-layout>
    </v-container>
</template>

<script>
    import axios from 'axios';
    export default {
        name: "LoginUser",
        data() {
            return {
                Email:"",
                Password:""
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
                            this.$router.replace("user");
                        }
                        )
                    .catch(
                        err => alert("Invalid credentials")
                    );
            },
            admin(e){
                e.preventDefault();
                this.$router.replace("admin_login");
            }
        }
    }
</script>

<style scoped>

</style>