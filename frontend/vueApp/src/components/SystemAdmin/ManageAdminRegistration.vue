<template>
    <div>
        <v-form>
            <v-container fluid>
                <v-layout align-center justify-center column fill-height/>
                    <v-flex offset-xs1 sm10 sm6>
                        <v-text-field
                                v-model="email"
                                :rules="emailRules"
                                label="E-mail"
                                required
                        ></v-text-field>
                    </v-flex>
                    <v-flex offset-xs1 sm10 sm6>
                        <v-text-field
                                v-model="password"
                                :append-icon="show1 ? 'visibility' : 'visibility_off'"
                                :rules="[rules.required, rules.min]"
                                :type="show1 ? 'text' : 'password'"
                                name="input-10-1"
                                label="Password"
                                hint="At least 8 characters"
                                counter
                                @click:append="show1 = !show1"
                        ></v-text-field>
                    </v-flex>
                    <v-flex offset-xs1 sm10 sm6>
                        <v-select
                                v-model="select"
                                :items="items"
                                label="Account type"
                                required
                        ></v-select>
                    </v-flex>
                <v-flex offset-xs1 sm10 sm6>
                    <v-btn @click="register">submit</v-btn>
                    <v-btn @click="clear">clear</v-btn>
                </v-flex>
            </v-container>

        </v-form>
    </div>
</template>

<script>
    import axios from 'axios/index';
    export default {
        name: "Registration",
        data(){
            return{
                email:"",
                emailRules: [
                    v => !!v || 'E-mail is required',
                    v => /.+@.+/.test(v) || 'E-mail must be valid'
                ],
                password:"",
                show1: false,
                rules: {
                    required: value => !!value || 'Required.',
                    min: v => v.length >= 8 || 'Min 8 characters',
                },
                select: null,
                items: [
                    'System Admin',
                    'Airline Admin',
                    'HotelAdmin Admin',
                    'Rent-A-Car Admin'
                ],
            }
        },
        methods: {
            register (e) {
                e.preventDefault();
                if (this.email === "" || this.password === "" || this.select === null) {
                    alert("Please fill in the required fields");
                    return;
                }
                const credentials = {
                    Email: this.email,
                    Password: this.password,
                    accountType: this.select
                };
                axios.create({withCredentials: true}).post('http://localhost:8000/api/admin/register', credentials)
                    .then(res => {if (res.status === 201) alert("Registration successful")})
                    .catch(err => alert("Registration unsuccessful"));
            },
            clear () {
                this.email = '';
                this.password = '';
                this.select = null;
            },
        }
    }
</script>

<style scoped>

</style>