<template>
    <v-card dark>
        <v-card-text>
            <v-container>
                <v-layout align-center justify-center row-wrap>
                    <v-flex xs3>
                        <v-text-field label="Name" prepend-icon="hotel" v-model="query.name"></v-text-field>
                    </v-flex>
                    <v-flex xs3>
                        <v-text-field label="Address" prepend-icon="place" v-model="query.address" readonly></v-text-field>
                    </v-flex>
                    <v-flex xs2>
                        <v-text-field v-model="time.from" label="From" prepend-icon="event" readonly></v-text-field>
                    </v-flex>
                    <v-flex xs2>
                        <template v-if="returnDate === null">
                            <v-menu v-model="menuTo" :close-on-content-click="false" lazy transition="scale-transition"
                                    :nudge-right="40" offset-y full-width max-width="290px" min-width="290px" >
                                <template v-slot:activator="{ on }">
                                    <v-text-field v-model="time.to" label="To" prepend-icon="event" readonly v-on="on"></v-text-field>
                                </template>
                                <v-date-picker no-title scrollable v-model = "time.to" @input="menuTo = false" :min="minDate"></v-date-picker>
                            </v-menu>
                        </template>
                        <template v-else>
                            <v-text-field v-model="time.to" label="From" prepend-icon="event" readonly></v-text-field>
                        </template>
                    </v-flex>
                    <v-flex xs2>
                        <v-btn @click="search">Search</v-btn>
                    </v-flex>
                </v-layout>
            </v-container>
        </v-card-text>
    </v-card>
</template>

<script>
    import moment from 'moment';

    export default {
        name: "HotelSearch",
        data() {
            return {
                menuTo: false,
                query: {
                    name: null,
                    address: null,
                },
                time: {
                    from: null,
                    to: null,
                },
                minDate: null,
                returnDate: null,
            }
        },
        created(){
            this.query.address = localStorage.getItem("destination");
            this.time.to = localStorage.getItem("returnDate");
            this.returnDate = localStorage.getItem("returnDate");
            this.time.from = localStorage.getItem("departureDate");
            this.minDate = this.time.from;
        },
        methods: {
            search(e) {
                e.preventDefault();
                const SearchQuery = {
                    "Name": this.query.name,
                    "Address": this.query.address,
                    "From": moment(this.time.from,"YYYY-MM-DD").valueOf().toString(),
                    "To": moment(this.time.to,"YYYY-MM-DD").valueOf().toString(),
                };
                this.$emit('search', SearchQuery)
            }
        },
        watch:{
            'time.from': function(){
                const timeFrom = moment(this.time.from,"YYYY-MM-DD");
                if(this.time.to !== ""){
                    const timeTo = moment(this.time.to,"YYYY-MM-DD");
                    if (timeFrom.isAfter(timeTo)) {
                        this.time.to = ""
                    }
                }
            },
            'time.to': function() {
                const timeTo = moment(this.time.to, "YYYY-MM-DD");
                if (this.time.from !== "") {
                    const timeFrom = moment(this.time.from, "YYYY-MM-DD");
                    if (timeTo.isBefore(timeFrom)) {
                        this.time.from = ""
                    }
                }
            }
        }

    }
</script>

<style scoped>

</style>