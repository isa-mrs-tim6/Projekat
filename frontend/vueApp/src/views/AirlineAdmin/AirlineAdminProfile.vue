<template>
    <div>
        <AirlineAdminNavDrawer></AirlineAdminNavDrawer>
        <AdminProfile></AdminProfile>
    </div>
</template>

<script>
    import AirlineAdminNavDrawer from "../../components/AirlineAdmin/AirlineAdminNavDrawer";
    import AdminProfile from "../../components/AdminProfile";
    import axios from 'axios';
    export default {
        name: "AirlineAdminProfile",
        components: {AdminProfile, AirlineAdminNavDrawer},
        mounter(){
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

</style>