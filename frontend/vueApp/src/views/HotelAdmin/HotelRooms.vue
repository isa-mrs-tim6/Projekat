<template>
    <v-container fill-height fluid id="main">
        <v-layout>
            <v-flex>
                <HotelAdminNavDrawer/>
            </v-flex>
            <v-container>
                <v-layout>
                    <v-flex align-self-center>
                        <ManageRooms/>
                    </v-flex>
                </v-layout>
            </v-container>
        </v-layout>
    </v-container>
</template>

<script>
    import HotelAdminNavDrawer from "../../components/HotelAdmin/HotelAdminNavDrawer";
    import ManageRooms from "../../components/HotelAdmin/ManageRooms";
    import axios from 'axios';

    export default {
        name: "HotelRooms",
        components: {HotelAdminNavDrawer, ManageRooms},
        mounted(){
            this.checkFirstPass();
        },
        methods: {
            checkFirstPass(){
                axios.create({withCredentials: true}).get("http://ec2-18-195-170-20.eu-central-1.compute.amazonaws.com:8000/api/admin/checkFirstPass")
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
        background-image: linear-gradient(to left top, #ac0c0c, #b41812, #bb2218, #c32a1d, #cb3223, #d23928, #d8402c, #df4631, #e74d37, #ee543c, #f65a42, #fe6148);
    }
</style>