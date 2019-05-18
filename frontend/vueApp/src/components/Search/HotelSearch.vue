<template>
    <v-card>
        <v-card-text>
            <v-container>
                <v-layout align-center justify-center row-wrap>
                    <v-flex xs3>
                        <v-text-field label="Name" prepend-icon="hotel" v-model="query.name"></v-text-field>
                    </v-flex>
                    <v-flex xs3>
                        <v-text-field label="Address" prepend-icon="place" v-model="query.address"></v-text-field>
                    </v-flex>
                    <v-flex xs2>
                        <v-text-field label="Room capacity" prepend-icon="meeting_room" v-model="query.capacity"></v-text-field>
                    </v-flex>
                    <v-flex xs2>
                        <v-text-field label="Price per night" prepend-icon="euro_symbol" v-model="query.price"></v-text-field>
                    </v-flex>
                    <v-flex xs2>
                        <v-btn @click="search">Search</v-btn>
                    </v-flex>
                </v-layout>
            </v-container>
        </v-card-text>
    </v-card>
</template>

<script>
    export default {
        name: "HotelSearch",
        data() {
            return {
                query: {
                    name: null,
                    address: null,
                    capacity: null,
                    price: null,
                }
            }
        },
        methods: {
            search(e) {
                e.preventDefault();
                const SearchQuery = {
                    "Name": this.query.name,
                    "Address": this.query.address,
                    "RoomCapacityDownLimit": 0,
                    "RoomCapacityUpLimit": parseInt(this.query.capacity),
                    "RoomPriceDownLimit": 0,
                    "RoomPriceUpLimit": parseFloat(this.query.price),
                };
                this.$emit('search', SearchQuery)
            }
        }
    }
</script>

<style scoped>

</style>