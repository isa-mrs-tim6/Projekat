<template>
    <v-container fluid>
        <v-layout align-center justify-space-around row fill-height>
            <v-flex xs12 sm3>
                <v-menu v-model="menuFrom" :close-on-content-click="false" lazy transition="scale-transition"
                        :nudge-right="40" offset-y full-width max-width="290px" min-width="290px">
                    <template v-slot:activator="{ on }">
                        <v-text-field v-model="time.from" label="From" prepend-icon="event" readonly v-on="on"></v-text-field>
                    </template>
                    <v-date-picker no-title scrollable v-model = "time.from" @input="menuFrom = false" :min="minDate" :max="maxDate"></v-date-picker>
                </v-menu>
            </v-flex>
            <v-flex xs12 sm3>
                <v-menu v-model="menuTo" :close-on-content-click="false" lazy transition="scale-transition"
                        :nudge-right="40" offset-y full-width max-width="290px" min-width="290px">
                    <template v-slot:activator="{ on }">
                        <v-text-field v-model="time.to" label="To" prepend-icon="event" readonly v-on="on"></v-text-field>
                    </template>
                    <v-date-picker no-title scrollable v-model = "time.to" @input="menuTo = false" :min="minDate" :max="maxDate"></v-date-picker>
                </v-menu>
            </v-flex>
        </v-layout>
        <v-layout align-center justify-space-around column fill-height>
            <div id="wrapper1" style="position: relative; height: 40vh; width: 80vw">
                <FinanceGraph v-bind:item="item.visitors" ref="chart" :chart-data="dataCollectionVisitors"></FinanceGraph>
            </div>
            <div id="wrapper2" style="position: relative; height: 40vh; width: 80vw">
                <FinanceGraph v-bind:item="item.currency" ref="chart" :chart-data="dataCollectionFinances"></FinanceGraph>
            </div>
        </v-layout>
        <v-layout align-center justify-space-around row fill-height>
            <v-btn @click="daily()">Daily</v-btn>
            <v-btn @click="weekly()">Weekly</v-btn>
            <v-btn @click="monthly()">Monthly</v-btn>
        </v-layout>
    </v-container>
</template>

<script>
    export default {
        name: "AirlineFinance",
        props: ["reservations"],
        data () {
            return {
                mask: 'time',
                rules: {
                    required: value => !!value || 'Required.',
                    time: value => {
                        const pattern = /(([0-1][0-9])|(2[0-3]))[0-5][0-9]/;
                        return pattern.test(value) || "Invalid time.[required format HH:mm]"
                    }
                },
                menuFrom: false,
                menuTo: false,
                dataCollectionVisitors: null,
                dataCollectionFinances: null,
                monthlyDictVisitors: new Map(),
                monthlyDictFinances: new Map(),
                weeklyDictVisitors: new Map(),
                weeklyDictFinances: new Map(),
                dailyDictVisitors: new Map(),
                dailyDictFinances: new Map(),
                item: {
                    visitors: 'visitors',
                    currency: '$',
                },
                time: {
                    from: null,
                    to: null,
                },
            }
        },created() {
            this.loadDaily();
            this.loadWeekly();
            this.loadMonthly();

            this.monthlyVisitors();
            this.monthlyFinances();
        },
    }
</script>

<style scoped>

</style>