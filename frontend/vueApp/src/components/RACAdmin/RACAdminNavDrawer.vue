<template>
    <NavigationDrawer height="100%" v-bind:title="title" v-bind:items="items"/>
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
                    { title: 'Manage Vehicles', icon: 'portrait', route: 'vehicles' },
                    { title: 'Manage Offices', icon: 'portrait', route: 'offices' },
                    { title: 'Company Profile', icon: 'portrait', route: 'rac_profile' },
                    { title: 'Admin Profile', icon: 'portrait', route: 'admin_profile' },
                ],
            }
        }
    }
</script>

<style scoped>

</style>