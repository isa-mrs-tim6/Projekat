<template>
    <div id="main">
        <HotelAdminNavDrawer/>
        <ManageRooms/>
    </div>
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
                axios.create({withCredentials: true}).get("http://localhost:8000/api/admin/checkFirstPass")
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
    @import '../../assets/css/HotelAdmin.css';
</style>