<template>
    <NavigationDrawer v-bind:title="title" v-bind:items="items"/>
</template>

<script>
    import NavigationDrawer from "../NavigationDrawer";
    export default {
        name: "HotelAdminNavDrawer",
        components: {NavigationDrawer},
        data () {
            return {
                title: "Hotel Admin Dashboard",
                items: [
                    { title: 'Ratings', icon: 'star', route: 'ratings' },
                    { title: 'Finances', icon: 'assessment', route: 'finances' },
                    { title: 'Features', icon: 'local_offer', route: 'features' },
                    { title: 'Rewards', icon: 'shop', route: 'rewards' },
                    { title: 'Rooms', icon: 'meeting_room', route: 'rooms' },
                    { title: 'Hotel Profile', icon: 'hotel', route: 'hotel_profile' },
                    { title: 'Admin Profile', icon: 'portrait', route: 'admin_profile' },
                ],
            }
        },
    }
</script>

<style scoped>

</style>