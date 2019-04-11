<template>
    <form>
        <v-label>ADMIN LOGIN</v-label>
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
        <v-btn @click="loginAdmin">Login</v-btn>
    </form>
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
                axios.post("http://localhost:8000/api/admin/login", creds)
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