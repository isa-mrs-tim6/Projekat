<template>
    <div id="main">
        <RACAdminNavDrawer/>
        <ManageVehicles/>
    </div>
</template>

<script>
    import RACAdminNavDrawer from "../../components/RACAdmin/RACAdminNavDrawer";
    import ManageVehicles from "../../components/RACAdmin/ManageVehicles";
    import axios from "axios";
    export default {
        name: "RACVehicles",
        components: {ManageVehicles, RACAdminNavDrawer},
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
    @import '../../assets/css/RACAdmin.css';
</style>