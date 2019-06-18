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
                                half-increments
                                readonly
                                medium
                        ></v-rating>
                    </v-card-title>
                    <v-layout row>
                        <v-flex xs12>
                            <v-card-title primary-title>
                                <v-layout column>
                                    <v-layout row>
                                        <v-flex xs3>
                                            <img src="../../assets/hotel_placeholder.png">
                                        </v-flex>
                                        <v-flex xs6 style="margin-right: -25px"></v-flex>
                                        <v-flex xs3>
                                            <gmap-map
                                                    :center="{lat: value.Hotel.Latitude, lng: value.Hotel.Longitude}"
                                                    :zoom="12"
                                                    style="width:360px;  height: 240px;"
                                            >
                                                <gmap-marker
                                                        :position="{lat: value.Hotel.Latitude, lng: value.Hotel.Longitude}"
                                                ></gmap-marker>
                                            </gmap-map>
                                        </v-flex>
                                    </v-layout>
                                    <v-layout row>
                                        <v-flex xs3>
                                            <div class="title">
                                                Description:<br> {{value.Hotel.Description}}
                                            </div>
                                        </v-flex>
                                        <v-flex xs6 style="margin-right: -25px"></v-flex>
                                        <v-flex xs3>
                                            <div class="title">Address:<br> {{value.Hotel.Address}}</div>
                                        </v-flex>
                                    </v-layout>
                                </v-layout>
                            </v-card-title>
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
                this.$router.push({ path: `/quick_reserve_room/${id}/${this.reservationID}/${this.passengers}` });
            }
        }
    }
</script>

<style scoped>

</style>