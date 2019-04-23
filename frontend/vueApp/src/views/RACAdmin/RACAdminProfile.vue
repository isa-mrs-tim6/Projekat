<template>
    <div id="main">
        <RACAdminNavDrawer/>
        <AdminProfile/>
    </div>
</template>

<script>
    import RACAdminNavDrawer from "../../components/RACAdmin/RACAdminNavDrawer";
    import AdminProfile from "../../components/AdminProfile";
    import axios from 'axios';

    export default {
        name: "RACAdminProfile",
        components: {RACAdminNavDrawer, AdminProfile},
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
    @import '../../assets/css/RACAdmin.css';
</style>