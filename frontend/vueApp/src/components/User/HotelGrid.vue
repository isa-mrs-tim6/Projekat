<template>
    <v-container fluid fill-height>
        <v-layout align-center justify-space-around row wrap fill-height>
            <v-flex xs12 style="margin-bottom: 2vw" v-for="(value, index) in hotels" v-bind:key="index">
                <v-card dark class="white--text">
                    <v-card-title class="headline" primary-title>
                        {{value.Name}}
                    </v-card-title>
                    <v-layout row>
                        <v-flex xs10>
                            <v-card-title primary-title>
                                <v-layout row>
                                    <v-flex xs8>
                                        <div>
                                            <div class="title"> <v-icon>place</v-icon> Address: {{value.Address}}</div>
                                        </div>
                                    </v-flex>
                                </v-layout>
                            </v-card-title>
                        </v-flex>
                        <v-flex xs2>
                            <div class="subheading">
                                Description: {{value.Description}}
                            </div>
                        </v-flex>
                    </v-layout>
                    <v-divider light></v-divider>
                    <v-card-actions class="pa-3">
                        <v-btn block @click="reserve(value.ID)">Reserve</v-btn>
                    </v-card-actions>
                </v-card>
            </v-flex>
        </v-layout>
    </v-container>
</template>

<script>
    export default {
        name: "HotelGrid",
        props: ["hotels", "passengers", "reservationID"],
        methods: {
            reserve(id) {
                console.log(id, this.passengers, this.reservationID);
            }
        }
    }
</script>

<style scoped>

</style>