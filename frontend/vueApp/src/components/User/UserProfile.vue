<template>
    <div>
        <v-form>
            <v-container>
                <v-flex xs12 sm2>
                    <v-text-field label="First Name" v-model="Name"></v-text-field>
                    <v-text-field label="Surname" v-model="Surname"></v-text-field>
                    <v-text-field label="Email" v-model="Email"></v-text-field>
                    <v-text-field label="Password" type="password" v-model="Password"></v-text-field>
                    <v-text-field label="Confirm password" type="password" v-model="ConfirmedPassword"></v-text-field>
                    <v-text-field label="Passport" v-model="Passport"></v-text-field>
                    <v-text-field label="Address" v-model="Address"></v-text-field>
                    <v-text-field label="Phone" v-model="Phone"></v-text-field>
                    <v-flex xs12 sm1>
                        <v-btn @click="updateUser">submit</v-btn>
                    </v-flex>
                </v-flex>
            </v-container>
        </v-form>
        <v-snackbar v-model="SuccessSnackbar" :timeout=2000 :top="true" color="success">{{this.SuccessSnackbarText}}</v-snackbar>
        <v-snackbar v-model="ErrorSnackbar" :timeout=2000 :top="true" color="error">{{this.ErrorSnackbarText}}</v-snackbar>
    </div>
</template>

<script>
    import axios from 'axios';
    import UserNavBar from "./UserNavBar";
    export default {
        name: "UserProfile",
        components: {UserNavBar},
        data(){
            return{
                Name:"",
                Surname:"",
                Email:"",
                OldEmail: "",
                Password:"",
                Passport: "",
                ConfirmedPassword:"",
                OldPassword:"",
                Address:"",
                Phone:"",
                SuccessSnackbar: false,
                SuccessSnackbarText: '',
                ErrorSnackbar: false,
                ErrorSnackbarText: '',
            }
        },
        beforeCreate() {
            axios.create({withCredentials: true}).get("http://localhost:8000/api/user/getProfile")
                .then(res =>{
                    this.Name = res.data.Name;
                    this.Surname = res.data.Surname;
                    this.Email = res.data.Email;
                    this.OldEmail = res.data.Email;
                    this.Address = res.data.Address;
                    this.Phone = res.data.Phone;
                    this.Passport = res.data.Passport;
                    this.OldPassword = res.data.Password;
                })
        },
        methods:{
            updateUser(e){
                e.preventDefault();
                if(this.Password !== this.ConfirmedPassword){
                    this.ErrorSnackbar = true;
                    this.ErrorSnackbarText = 'Password and confirmed password do not match';
                    return;
                }
                let newUserProfile={
                    Name: this.Name,
                    Surname: this.Surname,
                    OldEmail: this.OldEmail,
                    Email: this.Email,
                    Address: this.Address,
                    Phone: this.Phone,
                    Passport: this.Passport,
                    Password: this.Password
                };
                axios.create({withCredentials: true}).post("http://localhost:8000/api/user/updateProfile", newUserProfile)
                    .then(res =>{
                        this.SuccessSnackbar = true;
                        this.SuccessSnackbarText = 'Profile updated';
                    }).catch(err => {
                        this.ErrorSnackbar = true;
                        this.ErrorSnackbarText = 'An error has occured';
                    })
            }
        }
    }
</script>

<style scoped>

</style>