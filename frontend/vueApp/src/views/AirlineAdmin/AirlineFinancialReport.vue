<template>
    <div id="main">
        <AirlineAdminNavDrawer></AirlineAdminNavDrawer>
        <AirlineFinance :graphData="graphData"></AirlineFinance>
    </div>
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
            }
        },

    }
</script>

<style scoped>
    #main {
        background-image: linear-gradient(to right bottom, #142eae, #005bca, #007ed2, #009ccd, #0bb7c7, #47c0c6, #67c8c6, #81d0c7, #6ecac4, #58c4c3, #3cbdc2, #00b7c1);
    }
</style>