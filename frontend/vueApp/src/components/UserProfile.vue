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
                    <v-text-field label="Address" v-model="Address"></v-text-field>
                    <v-text-field label="Phone" v-model="Phone"></v-text-field>
                    <v-flex xs12 sm1>
                        <v-btn @click="updateUser">submit</v-btn>
                    </v-flex>
                </v-flex>
            </v-container>
        </v-form>
    </div>
</template>

<script>
    import axios from 'axios';
    export default {
        name: "UserProfile",
        data(){
            return{
                Name:"",
                Surname:"",
                Email:"",
                Password:"",
                ConfirmedPassword:"",
                OldPassword:"",
                Address:"",
                Phone:"",
                ID:1
            }
        },
        created(){
            this.ID = Math.floor(Math.random() * (5 - 1)) + 1;
        },
        beforeMount() {
            axios.get("http://localhost:8000/api/user/"+this.ID+"/getProfile")
                .then(res =>{
                    this.Name = res.data.Name;
                    this.Surname = res.data.Surname;
                    this.Email = res.data.Email;
                    this.Address = res.data.Address;
                    this.Phone = res.data.Phone;
                    this.OldPassword = res.data.Password;
                })
        },
        methods:{
            updateUser(e){
                e.preventDefault();
                if(this.Password !== this.ConfirmedPassword){
                    alert("Password and confirm password does not match");
                    return;
                }
                var newUserProfile={
                    Name: this.Name,
                    Surname: this.Surname,
                    Email: this.Email,
                    Address: this.Address,
                    Phone: this.Phone,
                    Password: this.OldPassword
                }
                if(this.Password !== "") {
                    newUserProfile.Password = this.Password;
                }
                axios.post("http://localhost:8000/api/user/" + this.ID + "/updateProfile", newUserProfile)
            }
        }
    }
</script>

<style scoped>

</style>