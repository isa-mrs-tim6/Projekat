<template>
    <NavigationDrawer v-bind:title="title" v-bind:items="items"/>
</template>

<script>
    import NavigationDrawer from "../NavigationDrawer";
    export default {
        name: "RACAdminNavDrawer",
        components: {NavigationDrawer},
        data () {
            return{
                title: "Rent-a-car Admin Dashboard",
                items: [
                    { title: 'Ratings', icon: 'star', route: 'ratings' },
                    { title: 'Finances', icon: 'assessment', route: 'finances' },
                    { title: 'Manage Vehicles', icon: 'directions_car', route: 'vehicles' },
                    { title: 'Manage Offices', icon: 'store', route: 'offices' },
                    { title: 'Company Profile', icon: 'business', route: 'rac_profile' },
                    { title: 'Admin Profile', icon: 'portrait', route: 'admin_profile' },
                ],
            }
        }
    }
</script>

<style scoped>

</style>