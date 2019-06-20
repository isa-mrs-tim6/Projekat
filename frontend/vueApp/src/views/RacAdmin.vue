<template>
    <div>
        <v-toolbar>
            <v-toolbar-side-icon></v-toolbar-side-icon>
            <v-toolbar-title>Rent-a-car Admin Panel</v-toolbar-title>
            <v-spacer></v-spacer>
            <v-toolbar-items class="hidden-sm-and-down">
                <v-btn @click="manageRacProfile" flat>Company Profile</v-btn>
                <v-btn @click="manageProfile" flat>Admin Profile</v-btn>
            </v-toolbar-items>
        </v-toolbar>
        <component v-bind:is="currentTabComponent"></component>
    </div>
</template>

<script>
    import ManageRacProfile from "../components/RACAdmin/ManageRacProfile";
    import AdminProfile from "../components/AdminProfile";
    import axios from "axios";

    export default {
        name: "RacAdmin",
        components: {ManageRacProfile, AdminProfile},
        data() {
            return {
                currentTabComponent: "ManageRacProfile",
            }
        },
        mounted(){
            this.checkFirstPass();
        },
        methods: {
            manageRacProfile(){
                this.currentTabComponent = "ManageRacProfile";
                this.checkFirstPass();
            },
            manageProfile(){
                this.currentTabComponent = "AdminProfile";
            },
            checkFirstPass(){
                axios.create({withCredentials: true}).get("http://ec2-35-159-21-254.eu-central-1.compute.amazonaws.com:8000/api/admin/checkFirstPass")
                    .then(
                        res =>{
                            if(res.data === false){
                                alert("Please update your password");
                                this.manageProfile();
                            }
                        }
                    )
            }
        },
    }
</script>

<style scoped>

</style>