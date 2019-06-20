<template>
    <v-container fill-height fluid>
        <v-layout>
            <v-flex>
                <AirlineAdminNavDrawer/>
            </v-flex>
            <v-container>
                <v-layout>
                    <v-flex align-self-center>
                        <AirlineFinance :graphData="graphData"></AirlineFinance>
                    </v-flex>
                </v-layout>
            </v-container>
        </v-layout>
    </v-container>
</template>

<script>
    import axios from "axios"
    import AirlineFinance from "../../components/AirlineAdmin/AirlineFinance";
    import AirlineAdminNavDrawer from "../../components/AirlineAdmin/AirlineAdminNavDrawer";

    export default {
        name: "AirlineFinancialReport",
        components: {AirlineAdminNavDrawer, AirlineFinance},
        data: () => ({
            graphData: [],
        }),
        created(){
            axios.create({withCredentials: true}).get('http://localhost:8000/api/airline/getGraphData')
                .then(res => {
                    this.graphData = res.data;
                });
        },
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
            },
        },

    }
</script>

<style scoped>
</style>