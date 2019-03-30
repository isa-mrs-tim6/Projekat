<template>
    <h1>
        <v-container grid-list-md text-xs-center>
            <v-layout row wrap>
                <v-flex xs12 v-if="regSuccess">
                    <h1>Registration success</h1>
                    <p>Your account is now enabled</p>
                </v-flex>
                <v-flex xs12 v-else="regSuccess">
                    <h1>Unable to confirm your registration</h1>
                    <p>Please try again</p>
                </v-flex>
            </v-layout>
        </v-container>
    </h1>
</template>

<script>
    import axios from 'axios';
    export default {
        name: "RegistrationSuccess",
        data() {
            return {
                regSuccess: false
            }
        },
        created() {
            axios.get('http://localhost:8000/api/admin/'+this.$route.params.type+'/completeRegistration/q='+this.$route.params.email)
                .then(res => this.regSuccess = true)
                .catch(err => this.regSuccess = false);
        }
    }
</script>

<style scoped>

</style>