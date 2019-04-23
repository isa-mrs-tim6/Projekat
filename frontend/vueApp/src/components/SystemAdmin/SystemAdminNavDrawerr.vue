<template>
    <NavigationDrawer v-bind:title="title" v-bind:items="items"/>
</template>

<script>
    import NavigationDrawer from "../NavigationDrawer";
    export default {
        name: "SystemAdminNavDrawer",
        components: {NavigationDrawer},
        data () {
            return {
                title: "System Admin Dashboard",
                items: [
                    { title: 'Admins', icon: 'group_add', route: 'accounts' },
                    { title: 'Airlines', icon: 'flight', route: 'airlines' },
                    { title: 'Hotels', icon: 'hotel', route: 'hotels' },
                    { title: 'Rent-a-Cars', icon: 'directions_car', route: 'rent_a_cars' },
                    { title: 'Admin Profile', icon: 'portrait', route: 'admin_profile' },
                ],
            }
        },
    }
</script>

<style scoped>

</style>