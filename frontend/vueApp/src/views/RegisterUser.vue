<template>
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
</template>

<script>
    import axios from 'axios';
    export default {
        name: "RegisterUser",
        data() {
            return {
                Name:"",
                Surname:"",
                Email:"",
                Password:"",
                ConfirmedPassword:"",
                Address:"",
                Phone:"",
            }
        },
        methods:{
            registerUser(e){
                e.preventDefault();
                if(this.Name === "" || this.Surname === "" || this.Email === "" || this.Password === ""
                    || this.ConfirmedPassword === "" || this.Address === "" || this.Phone === ""){
                    alert("Every field must be filled");
                    return;
                }
                if(this.Password !== this.ConfirmedPassword){
                    alert("Password and confirm password does not match");
                    return;
                }
                let userProfile={
                    Name: this.Name,
                    Surname: this.Surname,
                    Email: this.Email,
                    Address: this.Address,
                    Phone: this.Phone,
                    Password: this.Password
                };
                axios.post("http://localhost:8000/api/user/register", userProfile)
                    .catch(err => alert("Email already taken"));
            }
        }
    }
</script>

<style scoped>

</style>