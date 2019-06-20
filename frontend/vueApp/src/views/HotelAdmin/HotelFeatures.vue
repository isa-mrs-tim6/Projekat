<template>
    <v-container fill-height fluid id="main">
        <v-layout>
            <v-flex>
                <HotelAdminNavDrawer/>
            </v-flex>
            <v-container>
                <v-layout>
                    <v-flex align-self-center>
                        <HotelFeaturesComponent/>
                    </v-flex>
                </v-layout>
            </v-container>
        </v-layout>
    </v-container>

</template>

<script>
    import HotelAdminNavDrawer from "../../components/HotelAdmin/HotelAdminNavDrawer";
    import HotelFeaturesComponent from "../../components/HotelAdmin/HotelFeaturesComponent";
    import axios from 'axios';

    export default {
        name: "HotelFeatures",
        components: {HotelAdminNavDrawer, HotelFeaturesComponent},
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
                                this.$router.push("/admin_login");
                            }
                        }
                    )
            }
        },
    }
</script>

<style scoped>
    #main {
        background-image: linear-gradient(to left top, #ac0c0c, #b41812, #bb2218, #c32a1d, #cb3223, #d23928, #d8402c, #df4631, #e74d37, #ee543c, #f65a42, #fe6148);
    }
</style>