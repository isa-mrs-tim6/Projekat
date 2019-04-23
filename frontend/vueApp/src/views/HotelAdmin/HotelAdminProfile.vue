<template>
    <div id="main">
        <HotelAdminNavDrawer/>
        <AdminProfile/>
    </div>
</template>

<script>
    import HotelAdminNavDrawer from "../../components/HotelAdmin/HotelAdminNavDrawer";
    import AdminProfile from "../../components/AdminProfile";
    import axios from 'axios';

    export default {
        name: "HotelAdminProfile",
        components: {HotelAdminNavDrawer, AdminProfile},
        mounted(){
            this.checkFirstPass();
        },
        methods: {
            checkFirstPass(){
                axios.create({withCredentials: true}).get("http://localhost:8000/api/admin/checkFirstPass")
                    .then(
                        res =>{
                            if(res.data === false){
                                alert("Please update your password");
                            }
                        }
                    )
            }
        },
    }
</script>

<style scoped>
    @import '../../assets/css/HotelAdmin.css';
</style>