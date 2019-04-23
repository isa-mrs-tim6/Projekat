<template>
    <div>
        <SystemAdminNavDrawer/>
        <ManageRentACarCompanies/>
    </div>
</template>

<script>
    import SystemAdminNavDrawer from "../../components/SystemAdmin/SystemAdminNavDrawerr";
    import ManageRentACarCompanies from "../../components/SystemAdmin/ManageRentACarCompanies";
    import axios from 'axios';

    export default {
        name: "SystemAdminRentACars",
        components: {ManageRentACarCompanies, SystemAdminNavDrawer},
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

</style>