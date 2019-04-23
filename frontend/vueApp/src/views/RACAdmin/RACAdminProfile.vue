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
    #main {
        background-image: linear-gradient(to right bottom, #02bb04, #00b824, #00b535, #00b141, #00ae4c, #00ab55, #00a75e, #00a465, #00a06d, #009b73, #009778, #00927c);
    }
</style>