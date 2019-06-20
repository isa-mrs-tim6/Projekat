<template>
    <v-container grid-list-xl text-xs-center fill-height>
        <v-layout align-center justify-center row wrap fill-height>
            <v-flex xs12>
                <v-card class="flexcard">
                    <v-card-title primary-title>
                        <div class="headline font-weight-medium">Manage reservation rewards</div>
                    </v-card-title>
                    <v-card-text>
                        <v-layout  column>
                            <v-layout align-space-between row v-for="(value, index) in rewards" v-bind:key="index" style="padding: 2vw">
                                <v-flex xs2>
                                    <v-text-field
                                            v-model="value.RequiredNumber"
                                            label="Reservations needed"
                                            class="body-2"
                                            prepend-inner-icon="input"
                                            required
                                    ></v-text-field>
                                </v-flex>
                                <v-flex xs8>
                                    <v-slider
                                            v-model="value.PriceScale"
                                            :max=1
                                            :min=0
                                            step="0.01"
                                            ticks
                                            thumb-label="always"
                                    ></v-slider>
                                </v-flex>
                                <v-flex xs2>
                                    <v-btn small color="error" @click="deleteRow(index)">Delete</v-btn>
                                </v-flex>
                            </v-layout>
                        </v-layout>
                        <v-layout row-wrap justify-start>
                            <v-btn small @click="addRow" style="">Add reward</v-btn>
                        </v-layout>
                    </v-card-text>
                    <v-card-actions>
                        <v-btn color="primary" @click="changeRewards">submit</v-btn>
                        <v-btn @click="reset">reset</v-btn>
                    </v-card-actions>
                </v-card>
            </v-flex>
        </v-layout>
    </v-container>
</template>

<script>
    import axios from 'axios/index';

    export default {
        name: "ManageReservationRewards",
        data() {
            return {
                rewards: [],
                backup: [],
            }
        },
        mounted() {
            axios.create({withCredentials: true}).get('http://ec2-35-159-21-254.eu-central-1.compute.amazonaws.com:8000/api/reservations/rewards')
                .then(res => {
                    this.rewards = res.data;
                    this.backup = JSON.parse(JSON.stringify(this.rewards));
                })
                .catch(err => alert("Could not retrieve rewards"))
        },
        methods: {
            reset(e) {
                e.preventDefault();
                this.rewards = JSON.parse(JSON.stringify(this.backup));
            },
            changeRewards(e) {
                e.preventDefault();
                for (let i = 0; i < this.rewards.length; i++) {
                    if (!Number.isInteger(parseInt(this.rewards[i].RequiredNumber)) || parseInt(this.rewards[i].RequiredNumber) < 0) {
                        alert("Required reservations must be a positive integer");
                        return;
                    }
                    this.rewards[i].RequiredNumber = parseInt(this.rewards[i].RequiredNumber);
                    if (this.rewards[i].PriceScale > 1 || this.rewards[i].PriceScale < 0) {
                        alert("Reservation price scale must be in the range of 0 to 1");
                        return;
                    }
                }
                axios.create({withCredentials: true}).put('http://ec2-35-159-21-254.eu-central-1.compute.amazonaws.com:8000/api/reservations/rewards', this.rewards)
                    .then(res => {
                        this.backup = JSON.parse(JSON.stringify(this.rewards));
                    })
                    .catch(err => alert("Could not update rewards"))
            },
            deleteRow(index) {
                this.rewards.splice(index, 1)
            },
            addRow(e) {
                e.preventDefault();
                this.rewards.push({
                    RequiredNumber: null,
                    PriceScale:1,
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