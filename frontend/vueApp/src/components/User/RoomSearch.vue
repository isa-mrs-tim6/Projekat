<template>
    <v-card dark>
        <v-card-text>
            <v-container>
                <v-layout align-center justify-center row>
                    <v-flex xs10 style="margin-right: 10px; padding-top: 2px">
                        <v-select
                                v-model="selected_capacity"
                                :items="room_capacities"
                                :menu-props="{ maxHeight: '400' }"
                                label="Select room capacities"
                                prepend-icon="hotel"
                                multiple
                                hint="What type of rooms are you looking for?"
                                persistent-hint
                        ></v-select>
                    </v-flex>
                    <v-flex xs2>
                        <v-btn block @click="search">Search</v-btn>
                    </v-flex>
                </v-layout>
            </v-container>
        </v-card-text>
    </v-card>
</template>

<script>

    export default {
        name: "RoomSearch",
        props: ["room_capacities"],
        data() {
            return {
                menuFrom: false,
                menuTo: false,
                guests: null,
                selected_capacity: null,
            }
        },
        methods: {
            search(e) {
                e.preventDefault();
                const SearchQuery = {
                    "Guests": this.guests,
                    "Capacities": this.selected_capacity,
                    "From": localStorage.getItem("hotelStart"),
                    "To": localStorage.getItem("hotelEnd"),
                };
                this.$emit('search', SearchQuery)
            }
        },
    }
</script>

<style scoped>

</style>