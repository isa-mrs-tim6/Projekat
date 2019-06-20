<template>
    <NavigationDrawer height="100%" :title="$options.title" :items="$options.items"></NavigationDrawer>
</template>

<script>
    import NavigationDrawer from "../NavigationDrawer";
    export default {
        name: "AirlineAdminNavDrawer",
        components: {NavigationDrawer},
        title: "Airline Admin Dashboard",
        items:[
            {title: 'Add flight', icon: 'add', route: 'flight_add'},
            {title: 'Edit flight', icon: 'edit', route: 'flight_edit'},
            {title: 'Airplanes', icon: 'flight', route: 'add_airplane'},
            {title: 'Add destination', icon: 'add_location', route: 'destination_add'},
            {title: 'Features', icon: 'local_offer', route: 'features' },
            { title: 'Ratings', icon: 'star', route: 'ratings' },
            { title: 'Finances', icon: 'assessment', route: 'finances' },
            {title: 'Airline profile', icon: 'business', route: 'airline_profile'},
            {title: 'Admin profile', icon: 'portrait', route: 'admin_profile'},
        ]
    }
</script>

<style scoped>

</style>