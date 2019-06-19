<template>
    <v-container grid-list-xl text-xs-center fill-height>
        <v-layout align-center justify-center column wrap fill-height>
            <v-flex style="width: 60vw">
                <v-card class="flexcard">
                    <v-card-title primary-title>
                        <div class="headline font-weight-medium">Admin account</div>
                    </v-card-title>
                    <v-card-text class="grow">
                        <form
                                enctype="multipart/form-data"
                                action="http://localhost:8080/upload"
                                method="post">
                            <div class="col-md-4 centered align-items-center mx-auto text-center">
                                <img class="centered mx-auto text-center" height="184" width="184" v-on:click="fileUpload" style="margin-bottom:50px;"  v-if="this.PictureLink" v-bind:src="this.PictureLink">
                                <input type="file" accept="image/*," @change="updatePicture" id="picture" style="display: none">
                            </div>
                        </form>
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
        <v-snackbar v-model="SuccessSnackbar" :timeout=4000 :top="true" color="success">Successfully updated admin profile</v-snackbar>
    </v-container>
</template>

<script>
    import axios from 'axios';
    export default {
        name: "AdminProfile",
        data(){
            return{
                SuccessSnackbar: false,
                Name:"",
                Surname:"",
                Email:"",
                Password:"",
                ConfirmedPassword:"",
                OldPassword:"",
                Address:"",
                Phone:"",
                Picture:"",
                PictureLink:"",
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
                    this.Picture = res.data.Picture;
                    this.PictureLink = 'http://localhost:8000/'+this.Picture;
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
                    Picture: this.Picture,
                };
                if(this.Password !== "") {
                    newUserProfile.Password = this.Password;
                }
                axios.create({withCredentials: true}).post("http://localhost:8000/api/admin/updateProfile", newUserProfile)
                    .then(res => this.SuccessSnackbar = true);
            },
            updatePicture(e){
                const file = e.target.files[0];

                let formData = new FormData();
                formData.append('file', file);

                let a = null;

                axios.create({withCredentials: true}).post( 'http://localhost:8000/api/upload/updateProfilePicture',
                    formData,
                    {
                        headers: {
                            'Content-Type': 'multipart/form-data'
                        }
                    }
                ).then(res => {
                    a = res.data.toString();
                    a = a.substr(a.lastIndexOf(":")+2, a.lastIndexOf(`"`)-a.lastIndexOf(":")-2);
                    this.PictureLink = "1" + this.PictureLink;
                    this.PictureLink = this.PictureLink.substr(1, this.PictureLink.lastIndexOf("/")) + "/" + a;
                    this.Picture = a;
                });
            },
            fileUpload(e) {
                document.getElementById("picture").click();
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