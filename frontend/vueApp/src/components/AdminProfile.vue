<template>
    <div>
        <v-form>
            <v-container>
                <v-flex xs12 sm2>
                    <v-text-field label="First Name" v-model="Name"></v-text-field>
                    <v-text-field label="Surname" v-model="Surname"></v-text-field>
                    <v-text-field label="Email" :rules="emailRules" v-model="Email"></v-text-field>
                    <v-text-field
                            v-model="Password"
                            :append-icon="show1 ? 'visibility' : 'visibility_off'"
                            :rules="[rules.required, rules.min]"
                            :type="show1 ? 'text' : 'password'"
                            name="input-10-1"
                            label="Password"
                            hint="At least 8 characters"
                            counter
                            @click:append="show1 = !show1"
                    ></v-text-field>
                    <v-text-field
                            v-model="ConfirmedPassword"
                            :append-icon="show1 ? 'visibility' : 'visibility_off'"
                            :rules="[rules.required, rules.min]"
                            :type="show1 ? 'text' : 'password'"
                            name="input-10-1"
                            label="Confirm password"
                            hint="At least 8 characters"
                            counter
                            @click:append="show1 = !show1"></v-text-field>
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
        name: "AdminProfile",
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
                emailRules: [
                    v => !!v || 'E-mail is required',
                    v => /.+@.+/.test(v) || 'E-mail must be valid'
                ],
                rules: {
                    required: value => !!value || 'Required.',
                    min: v => v.length >= 8 || 'Min 8 characters',
                },
                show1: false,
            }
        },
        beforeMount() {
            axios.create({withCredentials: true}).get("http://localhost:8000/api/admin/getProfile")
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
                };
                if(this.Password !== "") {
                    newUserProfile.Password = this.Password;
                }
                axios.create({withCredentials: true}).post("http://localhost:8000/api/admin/updateProfile", newUserProfile)
            }
        }
    }
</script>

<style scoped>

</style>