<template>
    <v-container id="main" style="height: 100vh; max-height: 100%;" fluid>
        <v-layout style="height: 100vh; max-height: 100%;" row>
            <v-flex
                    shrink
                    pa-1
            >
                <SystemAdminNavDrawer height="100%"/>
            </v-flex>
            <v-flex
                    grow
                    pa-1
            >
                <ManageHotels/>
            </v-flex>
        </v-layout>
    </v-container>
</template>

<script>
    import SystemAdminNavDrawer from "../../components/SystemAdmin/SystemAdminNavDrawerr";
    import ManageHotels from "../../components/SystemAdmin/ManageHotels";
    import axios from 'axios';

    export default {
        name: "SystemAdminHotels",
        components: {ManageHotels, SystemAdminNavDrawer},
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