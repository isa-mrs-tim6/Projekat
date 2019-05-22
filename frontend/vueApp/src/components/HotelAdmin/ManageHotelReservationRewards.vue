<template>
    <v-container grid-list-xl text-xs-center fill-height>
        <v-layout align-center justify-center row wrap fill-height>
            <v-flex xs12>
                <v-card xs12 class="flexcard">
                    <v-card-title primary-title>
                        <div class="headline font-weight-medium">Manage rewards</div>
                    </v-card-title>
                    <v-card-text>
                        <v-container>
                            <v-layout v-bind="binding" align-center justify-start row wrap>
                                <v-flex xs4>
                                    <v-btn fab large color="red" @click="addRow">
                                        <v-icon x-large color="white">add</v-icon>
                                    </v-btn>
                                </v-flex>
                                <v-flex xs4 v-for="(value, index) in new_rewards" v-bind:key="index">
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
                                <v-flex xs4 v-for="(value, index) in rewards" v-bind:key="index">
                                    <v-card>
                                        <v-card-text>
                                            <v-container grid-list-xl fill-height>
                                                <v-layout column>
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
                                                        <v-select
                                                                :items="features"
                                                                name="Features"
                                                                label="Select a category"
                                                                v-model="value.Features"
                                                                v-validate="'required'"
                                                                item-text="name"
                                                        ></v-select>
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
                                                        <v-slider
                                                                v-model="value.PriceScale"
                                                                :max=1
                                                                :min=0
                                                                step="0.01"
                                                                ticks
                                                                label="Discount"
                                                                thumb-label="always"
                                                        ></v-slider>
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
        name: "ManageReservationRewards",
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
                rewards: [],
                new_rewards: [],
                backup: [],
                features: [],
                dialogm1: '',
                dialog: false
            }
        },
        mounted() {
            axios.create({withCredentials: true}).get('http://localhost:8000/api/hotel/rewards')
                .then(res => {
                    this.rewards = res.data;
                    this.backup = JSON.parse(JSON.stringify(this.rewards));
                })
                .catch(err => alert("Could not retrieve hotel rewards"))
            axios.create({withCredentials: true}).get('http://localhost:8000/api/hotel/features')
                .then(res => {
                    this.features = res.data;
                })
                .catch(err => alert("Could not retrieve features"))
        },
        methods: {
            addFeature(index) {
                if (!Number.isInteger(parseInt(this.new_rewards[index].Price)) || parseInt(this.new_rewards[index].Price) < 0) {
                    alert("Price must be a positive integer");
                    return;
                }
                this.new_rewards[index].Price = parseInt(this.new_rewards[index].Price);
                axios.create({withCredentials: true}).post('http://localhost:8000/api/hotel/rewards', this.new_rewards[index])
                    .then(res => {
                        this.rewards.unshift(this.new_rewards[index]);
                        this.backup = JSON.parse(JSON.stringify(this.rewards));
                        this.new_rewards.splice(index, 1);
                    })
                    .catch(err => alert("Could not add reward"))
            },
            cancelAddFeature(index) {
                this.new_rewards.splice(index, 1);
            },
            updateFeature(index) {
                axios.create({withCredentials: true}).put('http://localhost:8000/api/hotel/rewards', this.rewards[index])
                    .then(res => {
                        this.backup = JSON.parse(JSON.stringify(this.rewards));
                    })
                    .catch(err => alert("Could not update reward"))
            },
            resetFeature(index) {
                this.rewards = JSON.parse(JSON.stringify(this.backup));
            },
            deleteFeature(index) {
                axios.create({withCredentials: true}).delete('http://localhost:8000/api/hotel/rewards', { data: this.rewards[index] })
                    .then(res => {
                        this.rewards.splice(index, 1);
                        this.backup = JSON.parse(JSON.stringify(this.rewards));
                    })
                    .catch(err => alert("Could not delete reward"))
            },
            addRow(e) {
                e.preventDefault();
                this.new_rewards.unshift({
                    Name: null,
                    Description: null,
                    PriceScale: null,
                    Features: null,
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