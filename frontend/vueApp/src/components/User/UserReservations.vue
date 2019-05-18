<template>
    <div>

    </div>
</template>

<script>
    import axios from "axios";

    export default {
        name: "UserReservations",
        beforeCreate() {
            axios.create({withCredentials: true}).get("http://localhost:8000/api/reservations/getReservations").
            then(res => {
                console.log(res.data);
            })
        }
    }
</script>

<style scoped>

</style>