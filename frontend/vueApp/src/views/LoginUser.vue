<template>
    <form>
        <v-text-field
                v-model="Email"
                label="Email"
                required
        ></v-text-field>
        <v-text-field
                v-model="Password"
                label="Password"
                required
                type="password"
        ></v-text-field>
        <v-btn @click="loginUser">Login</v-btn>
    </form>
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
                axios.post("http://localhost:8000/api/user/login", creds)
                    .then(
                        res => alert("Login successful")
                        )
                    .catch(
                        err => alert("Invalid credentials")
                    );
            }
        }
    }
</script>

<style scoped>

</style>