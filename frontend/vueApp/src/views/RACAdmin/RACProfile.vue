<template>
    <div id="main">
        <RACAdminNavDrawer/>
        <ManageRacProfile/>
    </div>
</template>

<script>
    import RACAdminNavDrawer from "../../components/RACAdmin/RACAdminNavDrawer";
    import ManageRacProfile from "../../components/RACAdmin/ManageRacProfile";
    import axios from "axios";
    export default {
        name: "RACProfile",
        components: {ManageRacProfile, RACAdminNavDrawer},
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
    #main {
        background-image: linear-gradient(to right bottom, #02bb04, #00b824, #00b535, #00b141, #00ae4c, #00ab55, #00a75e, #00a465, #00a06d, #009b73, #009778, #00927c);
    }
</style>