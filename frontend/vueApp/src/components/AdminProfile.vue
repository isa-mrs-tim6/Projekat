<template>
    <v-container grid-list-xl text-xs-center style="height: 100vh;">
        <v-layout align-center justify-center column wrap fill-height>
            <v-flex style="width: 60vw">
                <v-card min-height="100%" class="flexcard">
                    <v-card-title primary-title>
                        <div class="headline font-weight-medium">Admin account</div>
                    </v-card-title>
                    <v-card-text class="grow">
                        <v-form>
                            <v-flex>
                                <v-layout align-center justify-space-around row fill-height>
                                    <v-text-field class="body-2" prepend-icon="person" style="margin-right: 10px" label="Name" v-model="Name"></v-text-field>
                                    <v-text-field class="body-2" style="margin-left: 10px" label="Surname" v-model="Surname"></v-text-field>
                                </v-layout>
                            </v-flex>
                            <v-text-field class="body-2" prepend-icon="email" label="Email" :rules="emailRules" v-model="Email"></v-text-field>
                            <v-text-field
                                    v-model="Password"
                                    :append-icon="show1 ? 'visibility' : 'visibility_off'"
                                    prepend-icon="lock"
                                    :rules="[rules.required, rules.min]"
                                    :type="show1 ? 'text' : 'password'"
                                    name="input-10-1"
                                    label="Password"
                                    hint="At least 8 characters"
                                    class="body-2"
                                    counter
                                    @click:append="show1 = !show1"
                            ></v-text-field>
                            <v-text-field
                                    v-model="ConfirmedPassword"
                                    :append-icon="show1 ? 'visibility' : 'visibility_off'"
                                    prepend-icon="lock"
                                    :rules="[rules.required, rules.min]"
                                    :type="show1 ? 'text' : 'password'"
                                    name="input-10-1"
                                    label="Confirm password"
                                    hint="At least 8 characters"
                                    class="body-2"
                                    counter
                                    @click:append="show1 = !show1"></v-text-field>
                            <v-text-field class="body-2" prepend-icon="place" label="Address" v-model="Address"></v-text-field>
                            <v-text-field class="body-2" prepend-icon="phone" label="Phone" v-model="Phone"></v-text-field>
                        </v-form>
                    </v-card-text>
                    <v-card-actions>
                        <v-btn color="primary" @click="updateUser">submit</v-btn>
                    </v-card-actions>
                </v-card>
            </v-flex>
        </v-layout>
    </v-container>
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
                    Password: this.OldPassword,
                    FirstPassChanged: true,
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
    .flexcard {
        display: flex;
        flex-direction: column;
    }
</style>