<template>
    <v-container fill-height fluid id="main">
        <v-layout>
            <v-flex>
                <AirlineAdminNavDrawer/>
            </v-flex>
            <v-container>
                <v-layout>
                    <v-flex align-self-center>
                        <AirlineFeaturesComponent/>
                    </v-flex>
                </v-layout>
            </v-container>
        </v-layout>
    </v-container>

</template>

<script>
    import AirlineAdminNavDrawer from "../../components/AirlineAdmin/AirlineAdminNavDrawer";
    import AirlineFeaturesComponent from "../../components/AirlineAdmin/AirlineFeaturesComponent";
    import axios from 'axios';

    export default {
        name: "AirlineFeatures",
        components: {AirlineAdminNavDrawer, AirlineFeaturesComponent},
        mounted(){
            this.checkFirstPass();
        },
        methods: {
            checkFirstPass(){
                axios.create({withCredentials: true}).get("http://ec2-18-195-170-20.eu-central-1.compute.amazonaws.com:8000/api/admin/checkFirstPass")
                    .then(
                        res =>{
                            if(res.data === false){
                                alert("Please update your password");
                                this.$router.push("/admin_profile");
                            }
                        }
                    )
            }
        },
    }
</script>

<style scoped>
    #main {
        background-image: linear-gradient(to right bottom, #142eae, #005bca, #007ed2, #009ccd, #0bb7c7, #47c0c6, #67c8c6, #81d0c7, #6ecac4, #58c4c3, #3cbdc2, #00b7c1);
    }
</style>