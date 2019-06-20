<template>
    <v-card dark>
        <v-card-title>
            <v-text-field
                    v-model="search"
                    append-icon="search"
                    label="Search"
                    single-line
                    hide-details
            ></v-text-field>
        </v-card-title>
        <v-data-table
            :headers="headers"
            :items="hotels"
            :search="search"
            dark
            class="elevation-1"
        >
            <template v-slot:items="hotels">
                <td>
                    <v-layout row>
                        <v-flex xs4>
                            <img width="300px" height="300px" v-bind:src = '"http://localhost:8000/"+hotels.item.Picture'>
                        </v-flex>
                        <v-flex xs8>
                            <v-layout style="margin-left: 5px;" column>
                                <div class="display-1">
                                    {{hotels.item.Name}}
                                </div>
                                <div class="headline">
                                    {{hotels.item.Description}}
                                </div>
                            </v-layout>
                        </v-flex>
                    </v-layout>
                </td>
                <td>
                    <v-layout column text-xs-center>
                        <div>
                            <gmap-map
                                    :center="{lat: hotels.item.Latitude, lng: hotels.item.Longitude}"
                                    :zoom="8"
                                    style="width:100%;  height: 300px;"
                            >
                                <gmap-marker
                                        :position="{lat: hotels.item.Latitude, lng: hotels.item.Longitude}"
                                ></gmap-marker>
                            </gmap-map>
                        </div>
                        <div class="headline">
                            {{hotels.item.Address}}
                        </div>
                    </v-layout>
                </td>
            </template>
            <template v-slot:no-results>
                <v-alert :value="true" color="error" icon="warning">
                    Your search for "{{ search }}" found no results.
                </v-alert>
            </template>
        </v-data-table>
    </v-card>
</template>

<script>
    export default {
        name: "Hotels",
        props: ["hotels"],
        data() {
            return {
                search: '',
                headers: [
                    { text: 'Name', value: 'Name', width:'80%'},
                    { text: 'Address', value: 'Address', width:'20%'},
                ],
            }
        },
    }
</script>

<style scoped>

</style>