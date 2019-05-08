<template>
    <div>
        <RACAdminNavDrawer height="100%"/>
        <RACFinance v-bind:reservations="Reservations"></RACFinance>
    </div>
</template>

<script>
    import axios from 'axios';
    import RACAdminNavDrawer from "../../components/RACAdmin/RACAdminNavDrawer";
    import RACFinance from "../../components/RACAdmin/RACFinance";

    export default {
        name: "RACFinancialReport",
        components: {RACFinance, RACAdminNavDrawer},
        data: () => ({
            Reservations: [],
        }),
        created() {
            axios.create({withCredentials: true}).get('http://localhost:8000/api/rentACarCompany/getProfile')
                .then(res => {
                        this.Name = res.data.Name;
                    }
                )
                .catch(err => alert(err));
            axios.create({withCredentials: true}).get('http://localhost:8000/api/rentACarCompany/getReservations')
                .then(res => {
                    this.Reservations = res.data;
                    for (let i = 0; i < this.Reservations.length; i++) {
                        this.Rating += this.Reservations[i].Rating;
                    }
                    if (this.Reservations.length > 0) {
                        this.Rating /= this.Reservations.length;
                    }
                } )
                .catch(err => alert("Could not retrieve vehicle reservations"));
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
</style>