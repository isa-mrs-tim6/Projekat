<template>
    <div>
        <h2>Airlines</h2>
        <Airlines v-bind:airlines="Airlines"/>
        <h2>Hotels</h2>
        <Hotels v-bind:hotels="Hotels"></Hotels>
        <h2>Rent a car companies</h2>
        <RentACarCompanies v-bind:rent-a-car-companies="RentACarCompanies"></RentACarCompanies>
    </div>
</template>

<script>
    import Airlines from "../components/Airlines";
    import Hotels from "../components/Hotels";
    import RentACarCompanies from "../components/RentACarCompanies";
    import axios from 'axios';
    export default {
        name: "MainUserPage",
        components: {RentACarCompanies, Hotels, Airlines},
        data() {
            return {
                Airlines: [],
                Hotels: [],
                RentACarCompanies: [],
            }
        },
        created() {
            axios.create({withCredentials: true}).get('http://localhost:8000/api/airline/getAirlines')
                .then(res => this.Airlines = res.data)
                .catch(err => alert("Could not retrieve airline companies"));
            axios.create({withCredentials: true}).get('http://localhost:8000/api/hotel/getHotels')
                .then(res => this.Hotels = res.data)
                .catch(err => alert("Could not retrieve hotels"));
            axios.create({withCredentials: true}).get('http://localhost:8000/api/rentACarCompany/getRentACarCompanies')
                .then(res => this.RentACarCompanies = res.data)
                .catch(err => alert("Could not retrieve rent-a-car companies"));
        }
    }
</script>

<style scoped>

</style>