<template>
    <v-layout justify-center>
        <v-flex xs12>
            <v-container pa-0 mt-2 grid-list-md fluid>
                <v-layout row wrap v-for="(item, index) in items" :key="item.ID"
                          xs4>
                    <v-flex>
                        <template>
                            <v-card height="100%">
                                <v-card-title>
                                    {{item.Name}}
                                    <v-spacer></v-spacer>
                                    <v-btn @click="searchVehicles(index)">
                                        Vehicles
                                    </v-btn>
                                </v-card-title>
                                <v-card-text>
                                    <v-label>{{item.Address}}</v-label><br/>
                                    <v-label>{{item.Promo || item.Description}}</v-label>
                                </v-card-text>
                            </v-card>
                        </template>
                    </v-flex>
                </v-layout>
            </v-container>
        </v-flex>
    </v-layout>
</template>

<script>
    import axios from 'axios'

    export default {
        name: "RACSearch",
        data() {
            return {
                items: [],
            }
        },
        methods: {
            searchVehicles(index){
                let location = 1;
                this.$router.push({ path: `/vehiclesSearch/${this.items[index].ID}-${location}` })
            }
        },
        beforeCreate() {
            axios.create({withCredentials: true}).get("http://localhost:8000/api/rentACarCompany/getRentACarCompanies")
                .then(res =>{
                    this.items = res.data;
                })
        }
    }
</script>

<style scoped>

</style>