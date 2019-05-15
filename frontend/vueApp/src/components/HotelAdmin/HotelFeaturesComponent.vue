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
                                <v-flex xs4 v-for="(value, index) in new_features" v-bind:key="index">
                                    <v-card>
                                        <v-card-text>
                                            <v-container grid-list-xl fill-height>
                                                <v-layout column>
                                                    <v-flex>
                                                        <v-icon x-large>{{value.Icon}}</v-icon>
                                                    </v-flex>
                                                    <v-flex>
                                                        <v-autocomplete
                                                                v-model="value.Icon"
                                                                :items="icons_model"
                                                                :label="`Pick an icon`"
                                                                prepend-inner-icon="wallpaper"
                                                                persistent-hint
                                                        >
                                                            <template v-slot:append-outer>
                                                                <v-slide-x-reverse-transition
                                                                        mode="out-in"
                                                                >
                                                                </v-slide-x-reverse-transition>
                                                            </template>
                                                        </v-autocomplete>
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
                                                                prepend-inner-icon="note"
                                                                class="body-2"
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
                                                                <v-btn fab small color="green" @click="addFeature(index)" v-on="on">
                                                                    <v-icon color="white">done</v-icon>
                                                                </v-btn>
                                                            </template>
                                                            <span>Finish adding new feature</span>
                                                        </v-tooltip>
                                                        <v-tooltip bottom>
                                                            <template v-slot:activator="{ on }">
                                                                <v-btn fab small color="error" @click="cancelAddFeature(index)" v-on="on">
                                                                    <v-icon>cancel</v-icon>
                                                                </v-btn>
                                                            </template>
                                                            <span>Cancel adding new feature</span>
                                                        </v-tooltip>
                                                    </v-flex>
                                                </v-layout>
                                            </v-container>
                                        </v-card-text>
                                    </v-card>
                                </v-flex>
                                <v-flex xs4 v-for="(value, index) in features" v-bind:key="index">
                                    <v-card>
                                        <v-card-text>
                                            <v-container grid-list-xl fill-height>
                                                <v-layout column>
                                                    <v-flex>
                                                        <v-icon x-large>{{value.Icon}}</v-icon>
                                                    </v-flex>
                                                    <v-flex>
                                                        <v-autocomplete
                                                                v-model="value.Icon"
                                                                :items="icons_model"
                                                                :label="`Pick an icon`"
                                                                prepend-inner-icon="wallpaper"
                                                                persistent-hint
                                                        >
                                                            <template v-slot:append-outer>
                                                                <v-slide-x-reverse-transition
                                                                        mode="out-in"
                                                                >
                                                                </v-slide-x-reverse-transition>
                                                            </template>
                                                        </v-autocomplete>
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
                                                                prepend-inner-icon="note"
                                                                class="body-2"
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
                                                                <v-btn fab small color="warning" @click="resetFeature" v-on="on">
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
    import { icons } from '../../iconList';

    export default {
        name: "HotelFeaturesComponent",
        icons_model: null,
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
                new_features: [],
                backup: [],
            }
        },
        created() {
          this.icons_model = icons;
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
            addFeature(index) {
                if (!Number.isInteger(parseInt(this.new_features[index].Price)) || parseInt(this.new_features[index].Price) < 0) {
                    alert("Price must be a positive integer");
                    return;
                }
                this.new_features[index].Price = parseInt(this.new_features[index].Price);
                axios.create({withCredentials: true}).post('http://localhost:8000/api/hotel/features', this.new_features[index])
                    .then(res => {
                        this.features.unshift(this.new_features[index]);
                        this.backup = JSON.parse(JSON.stringify(this.features));
                        this.new_features.splice(index, 1);
                    })
                    .catch(err => alert("Could not add feature"))
            },
            cancelAddFeature(index) {
                this.new_features.splice(index, 1);
            },
            updateFeature(index) {
                if (!Number.isInteger(parseInt(this.features[index].Price)) || parseInt(this.features[index].Price) < 0) {
                    alert("Price must be a positive integer");
                    return;
                }
                this.features[index].Price = parseInt(this.features[index].Price);

                axios.create({withCredentials: true}).put('http://localhost:8000/api/hotel/features', this.features[index])
                    .then(res => {
                        this.backup = JSON.parse(JSON.stringify(this.features));
                    })
                    .catch(err => alert("Could not update feature"))
            },
            resetFeature(index) {
                this.features = JSON.parse(JSON.stringify(this.backup));
            },
            deleteFeature(index) {
                axios.create({withCredentials: true}).delete('http://localhost:8000/api/hotel/features', { data: this.features[index] })
                    .then(res => {
                        this.features.splice(index, 1);
                        this.backup = JSON.parse(JSON.stringify(this.features));
                    })
                    .catch(err => alert("Could not delete feature"))
            },
            addRow(e) {
                e.preventDefault();
                this.new_features.unshift({
                    Icon: "add_circle_outline",
                    Name: null,
                    Description: null,
                    Price: null,
                });
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