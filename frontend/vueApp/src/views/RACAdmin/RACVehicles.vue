<template>
    <v-container id="main" style="height: 100vh; max-height: 100%;" fluid>
        <v-layout style="height: 100vh; max-height: 100%;" row>
            <v-flex shrink pa-1>
                <RACAdminNavDrawer height="100%"/>
            </v-flex>
            <v-flex grow pa-1>
                <ManageVehicles height="100%"/>
            </v-flex>
        </v-layout>
    </v-container>
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
    #main {
        background-image: linear-gradient(to right top, #ffffff, #fafafb, #f5f5f6, #eff0f2, #eaebee, #e7eaee, #e4e9ee, #e1e8ee, #dfebf1, #deeef3, #ddf1f3, #ddf4f2);
    }
</style>