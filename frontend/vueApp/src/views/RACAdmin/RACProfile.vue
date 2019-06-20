<template>
    <v-container fill-height fluid id="main">
        <v-layout>
            <v-flex>
                <RACAdminNavDrawer/>
            </v-flex>
            <v-container>
                <v-layout>
                    <v-flex align-self-center>
                        <ManageRacProfile/>
                    </v-flex>
                </v-layout>
            </v-container>
        </v-layout>
    </v-container>
</template>

<script>
    import RACAdminNavDrawer from "../../components/RACAdmin/RACAdminNavDrawer";
    import ManageRacProfile from "../../components/RACAdmin/ManageRacProfile";
    import axios from "axios";
    export default {
        name: "RACProfile",
        components: {ManageRacProfile, RACAdminNavDrawer},
        mounted(){
            this.checkFirstPass();
        },
        methods: {
            checkFirstPass(){
                axios.create({withCredentials: true}).get("http://ec2-35-159-21-254.eu-central-1.compute.amazonaws.com:8000/api/admin/checkFirstPass")
                    .then(
                        res =>{
                            if(res.data === false){
                                alert("Please update your password before accessing this feature");
                                this.$router.replace("admin_profile");
                            }
                        }
                    )
            }
        },
    }
</script>

<style scoped>
    #main {
        background-image: linear-gradient(to right bottom, #560579, #561a71, #56266a, #543062, #52395b, #594460, #604f64, #665a69, #796c7c, #8c7f8f, #9f93a3, #b3a7b7);
    }
</style>