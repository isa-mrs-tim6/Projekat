<template>
    <v-container grid-list-xl text-xs-center fill-height>
        <v-layout align-center justify-center row wrap fill-height>
            <v-flex xs12>
                <v-card xs12 class="flexcard">
                    <v-card-title primary-title>
                        <div class="headline font-weight-medium">Manage features</div>
                    </v-card-title>
                    <v-card-text>
                        <v-container>
                            <v-layout v-bind="binding" align-center justify-start row wrap>
                                <v-flex xs4>
                                    <v-btn fab large color="red" @click="addRow">
                                        <v-icon x-large color="white">add</v-icon>
                                    </v-btn>
                                </v-flex>
                                <v-flex xs4 v-for="(value, index) in features" v-bind:key="index">
                                    <v-card>
                                        <v-card-text>
                                            <v-container grid-list-xl fill-height>
                                                <v-layout column>
                                                    <v-flex>
                                                        <v-icon x-large>{{value.Icon}}</v-icon>
                                                    </v-flex>
                                                    <v-flex >
                                                        <v-text-field
                                                                v-model="value.Name"
                                                                label="Name"
                                                                class="body-2"
                                                                prepend-inner-icon="input"
                                                                required
                                                        ></v-text-field>
                                                    </v-flex>
                                                    <v-flex>
                                                        <v-textarea
                                                                auto-grow
                                                                name="input-2-2"
                                                                label="Description"
                                                                v-model="value.Description"
                                                        ></v-textarea>
                                                    </v-flex>
                                                    <v-flex>
                                                        <v-text-field
                                                                v-model="value.Price"
                                                                label="Price"
                                                                class="body-2"
                                                                prepend-inner-icon="euro_symbol"
                                                                required
                                                        ></v-text-field>
                                                    </v-flex>
                                                    <v-flex>
                                                        <v-tooltip bottom>
                                                            <template v-slot:activator="{ on }">
                                                                <v-btn fab small color="primary" @click="updateFeature(index)" v-on="on">
                                                                    <v-icon>edit</v-icon>
                                                                </v-btn>
                                                            </template>
                                                            <span>Update feature</span>
                                                        </v-tooltip>
                                                        <v-tooltip bottom>
                                                            <template v-slot:activator="{ on }">
                                                                <v-btn fab small color="warning" @click="resetFeature(index)" v-on="on">
                                                                    <v-icon>restore</v-icon>
                                                                </v-btn>
                                                            </template>
                                                            <span>Undo changes</span>
                                                        </v-tooltip>
                                                        <v-tooltip bottom>
                                                            <template v-slot:activator="{ on }">
                                                                <v-btn fab small color="error" @click="deleteFeature(index)" v-on="on">
                                                                    <v-icon>delete</v-icon>
                                                                </v-btn>
                                                            </template>
                                                            <span>Delete feature</span>
                                                        </v-tooltip>
                                                    </v-flex>
                                                </v-layout>
                                            </v-container>
                                        </v-card-text>
                                    </v-card>
                                </v-flex>
                            </v-layout>
                        </v-container>
                    </v-card-text>
                </v-card>
            </v-flex>
        </v-layout>
    </v-container>
</template>

<script>
    import axios from 'axios/index';

    export default {
        name: "HotelFeaturesComponent",
        computed: {
            binding () {
                const binding = {};
                if (this.$vuetify.breakpoint.mdAndUp)
                    binding.row = true;
                else
                    binding.column = true;
                return binding;
            }
        },
        data() {
            return {
                features: [],
                backup: [],
            }
        },
        mounted() {
            axios.create({withCredentials: true}).get('http://localhost:8000/api/hotel/features')
                .then(res => {
                    this.features = res.data;
                    this.backup = JSON.parse(JSON.stringify(this.features));
                })
                .catch(err => alert("Could not retrieve features"))
        },
        methods: {
            updateFeature(index) {

            },
            resetFeature(index) {

            },
            deleteFeature(index) {
                this.features.splice(index, 1)
            },
            addRow(e) {
                e.preventDefault();
                this.features.unshift({
                    Icon: "add_circle_outline",
                    Name: null,
                    Description: null,
                    Price: null,
                })
            },
        }
    }
</script>

<style scoped>
    .flexcard {
        display: flex;
        flex-direction: column;
    }
</style>