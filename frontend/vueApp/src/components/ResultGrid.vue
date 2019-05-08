<template>
    <v-layout justify-center>
        <v-flex xs12>
            <v-toolbar-title>{{title}}</v-toolbar-title>
            <v-container pa-0 mt-2 grid-list-md fluid>
                <v-layout row wrap>
                    <v-flex
                            v-for="item in items"
                            :key="item.Name"
                            xs4
                    >
                        <template>
                            <v-card height="100%">
                                <v-card-title>{{item.Name}}</v-card-title>
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
    export default {
        name: "ResultGrid",
        props: ["items", "title"],
        data () {
            return {

            }
        },
        methods: {

        },
    }
</script>

<style scoped>

</style>