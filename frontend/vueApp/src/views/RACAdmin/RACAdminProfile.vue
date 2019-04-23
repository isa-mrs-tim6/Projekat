<template>
    <div>
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
        background-image: linear-gradient(to right top, #ffffff, #fafafb, #f5f5f6, #eff0f2, #eaebee, #e7eaee, #e4e9ee, #e1e8ee, #dfebf1, #deeef3, #ddf1f3, #ddf4f2);
    }
</style>