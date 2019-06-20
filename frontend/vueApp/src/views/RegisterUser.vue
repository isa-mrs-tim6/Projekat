<template>
    <div>
        <UserNavBar></UserNavBar>
        <v-layout align-center justify-center style="margin-top: 75px">
            <v-flex xs4>
                <v-card>
                    <v-card-title>
                        <h2>Registration</h2>
                    </v-card-title>
                    <v-card-text v-if="this.RegComplete === false">
                        <form>
                            <v-text-field
                                    v-model="Name"
                                    label="Name"
                                    required
                            ></v-text-field>
                            <v-text-field
                                    v-model="Surname"
                                    label="Surname"
                                    required
                            ></v-text-field>
                            <v-text-field
                                    v-model="Email"
                                    label="Email"
                                    required
                            ></v-text-field>
                            <v-text-field
                                    v-model="Password"
                                    label="Password"
                                    type="password"
                                    required
                            ></v-text-field>
                            <v-text-field
                                    v-model="ConfirmedPassword"
                                    label="Confirm password"
                                    type="password"
                                    required
                            ></v-text-field>
                            <v-text-field
                                    v-model="Passport"
                                    label="Passport"
                                    required
                            ></v-text-field>
                            <v-text-field
                                    v-model="Address"
                                    label="Address"
                                    required
                            ></v-text-field>
                            <v-text-field
                                    v-model="Phone"
                                    label="Phone"
                                    required
                            ></v-text-field>
                            <v-btn @click="registerUser">Register</v-btn>
                        </form>
                    </v-card-text>
                    <v-card-text v-else style="font-size: 14pt">
                        Registration successful! An email with an activation link has been sent to <b>{{this.Email}}</b>
                    </v-card-text>
                </v-card>
            </v-flex>
        </v-layout>
        <v-snackbar v-model="ErrorSnackbar" :timeout=2000 :top="true" color="error">{{this.ErrorSnackbarText}}</v-snackbar>
    </div>
</template>

<script>
    import UserNavBar from "../components/User/UserNavBar";
    import axios from 'axios';
    export default {
        name: "RegisterUser",
        components: {UserNavBar},
        data() {
            return {
                SuccessSnackbar: false,
                SuccessSnackbarText: '',
                ErrorSnackbar: false,
                ErrorSnackbarText: '',
                RegComplete: false,
                Name:"",
                Surname:"",
                Email:"",
                Password:"",
                ConfirmedPassword:"",
                Passport: "",
                Address:"",
                Phone:"",
            }
        },
        methods:{
            registerUser(e){
                e.preventDefault();
                if(this.Name === "" || this.Surname === "" || this.Email === "" || this.Password === ""
                    || this.ConfirmedPassword === "" || this.Address === "" || this.Phone === "" || this.Passport === ""){
                    this.ErrorSnackbar = true;
                    this.ErrorSnackbarText = 'Every field must be filled';
                    return;
                }
                if(this.Password !== this.ConfirmedPassword){
                    this.ErrorSnackbar = true;
                    this.ErrorSnackbarText = 'Password and confirm password do not match';
                    return;
                }
                let userProfile={
                    Name: this.Name,
                    Surname: this.Surname,
                    Email: this.Email,
                    Address: this.Address,
                    Phone: this.Phone,
                    Password: this.Password,
                    Passport: this.Passport
                };
                axios.post("http://ec2-35-159-21-254.eu-central-1.compute.amazonaws.com:8000/api/user/register", userProfile)
                    .then(res => this.RegComplete = true)
                    .catch(err => {
                        if (err.response.status === 406){
                            this.ErrorSnackbar = true;
                            this.ErrorSnackbarText = 'Invalid email format';
                        }else{
                            this.ErrorSnackbar = true;
                            this.ErrorSnackbarText = 'Email already taken';
                        }
                    });
            }
        }
    }
</script>

<style scoped>

</style>