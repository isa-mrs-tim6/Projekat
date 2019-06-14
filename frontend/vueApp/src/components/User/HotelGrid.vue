<template>
    <v-container fluid fill-height>
        <v-layout align-center justify-space-around row wrap fill-height>
            <v-flex xs12 style="margin-bottom: 2vw" v-for="(value, index) in hotels" v-bind:key="index">
                <v-card dark class="white--text">
                    <v-card-title class="headline" primary-title>
                        {{value.Hotel.Name}}
                        <v-rating
                                v-model=value.Rating
                                background-color="orange lighten-3"
                                color="orange"
                                readonly
                                medium
                        ></v-rating>
                    </v-card-title>
                    <v-layout row>
                        <v-flex xs10>
                            <v-card-title primary-title>
                                <v-layout row>
                                    <v-flex xs8>
                                        <div>
                                            <div class="title"> <v-icon>place</v-icon> Address: {{value.Hotel.Address}}</div>
                                        </div>
                                    </v-flex>
                                </v-layout>
                            </v-card-title>
                        </v-flex>
                        <v-flex xs2>
                            <div class="subheading">
                                Description: {{value.Hotel.Description}}
                            </div>
                        </v-flex>
                    </v-layout>
                    <v-divider light/>
                    <v-card-actions v-if="reservationID != null" class="pa-3">
                        <v-btn block @click="reserve(value.Hotel.ID)">Reserve</v-btn>
                        <v-btn block @click="quickReserve(value.Hotel.ID)">Quick Reserve</v-btn>
                    </v-card-actions>
                </v-card>
            </v-flex>
        </v-layout>
    </v-container>
</template>

<script>
    export default {
        name: "HotelGrid",
        props: ["hotels", "passengers", "reservationID", "start", "end"],
        methods: {
            reserve(id) {
                this.$router.push({ path: `/reserve_room/${id}/${this.reservationID}/${this.passengers}` });
            },
            quickReserve(id) {
                this.$router.push({ path: `/quick_reserve_room/${id}/${this.reservationID}/${this.start}/${this.end}` });
            }
        }
    }
</script>

<style scoped>

</style>