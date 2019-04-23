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
                axios.create({withCredentials: true}).post("http://localhost:8000/api/admin/login", creds)
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
                    .catch(
                        err => alert("Invalid login info")
                    );
            }
        }
    }
</script>

<style scoped>

</style>