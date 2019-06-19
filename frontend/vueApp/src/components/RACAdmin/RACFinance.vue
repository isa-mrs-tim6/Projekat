<template xmlns:v-slot="http://www.w3.org/1999/XSL/Transform">
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
            <div style="height: 20vh;"></div>
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
    import FinanceGraph from './FinanceGraph'
    import moment from 'moment';

    export default {
        name: "RACFinance",
        components: {
            FinanceGraph
        },
        props: ["reservations"],
        data () {
            return {
                mask: 'time',
                rules:{
                    required: value => !!value || 'Required.',
                    time: value=>{
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
        },
        created() {
            this.loadDaily();
            this.loadWeekly();
            this.loadMonthly();

            this.monthlyVisitors();
            this.monthlyFinances();
        },
        methods: {
            loadDaily () {
                this.dailyDictVisitors = new Map();
                this.dailyDictFinances = new Map();
                if (!moment(this.time.from).isValid() || !moment(this.time.to).isValid()) {
                    return;
                }
                let dateFrom = moment(this.time.from);
                const dateTo = moment(this.time.to);

                while (dateFrom.isSameOrBefore(dateTo)) {
                    this.dailyDictVisitors.set(dateFrom.format('DD-MM-YYYY'), 0);
                    this.dailyDictFinances.set(dateFrom.format('DD-MM-YYYY'), 0);
                    dateFrom.add(1, 'days');
                }

                for (let i = 0; i < this.reservations.length; i++) {
                    const key = moment(this.reservations[i].Beginning).format('DD-MM-YYYY');
                    if (this.dailyDictVisitors.has(key)) {
                        this.dailyDictVisitors.set(key, this.dailyDictVisitors.get(key) + 1);
                        this.dailyDictFinances.set(key, this.dailyDictFinances.get(key) + this.reservations[i].Price);
                    }
                }
            },
            daily() {
                this.loadDaily();
                this.dailyVisitors();
                this.dailyFinances();
            },
            dailyVisitors () {
                this.dataCollectionVisitors = {
                    labels: Array.from(this.dailyDictVisitors.keys()),
                    datasets: [
                        {
                            label: 'Daily vehicles',
                            backgroundColor: '#f87979',
                            data: Array.from(this.dailyDictVisitors.values())
                        },
                    ],
                    field: "vehicles",
                }
            },
            dailyFinances () {
                this.dataCollectionFinances = {
                    labels: Array.from(this.dailyDictVisitors.keys()),
                    datasets: [
                        {
                            label: 'Daily finance',
                            backgroundColor: '#76c2f8',
                            data: Array.from(this.dailyDictFinances.values())
                        },
                    ],
                    field: "€",
                }
            },
            loadWeekly () {
                this.weeklyDictVisitors = new Map();
                this.weeklyDictFinances = new Map();
                if (!moment(this.time.from).isValid() || !moment(this.time.to).isValid()) {
                    return;
                }
                let dateFrom = moment(this.time.from);
                const dateTo = moment(this.time.to);

                while (dateFrom.isSameOrBefore(dateTo)) {
                    this.weeklyDictVisitors.set("Week " + dateFrom.isoWeek() + ", " + dateFrom.year(), 0);
                    this.weeklyDictFinances.set("Week " + dateFrom.isoWeek() + ", " + dateFrom.year(), 0);
                    dateFrom.add(1, 'weeks');
                }

                for (let i = 0; i < this.reservations.length; i++) {
                    const key = "Week " + moment(this.reservations[i].Beginning).isoWeek() + ", " +  moment(this.reservations[i].Beginning).year();
                    if (this.weeklyDictVisitors.has(key)) {
                        this.weeklyDictVisitors.set(key, this.weeklyDictVisitors.get(key) + 1);
                        this.weeklyDictFinances.set(key, this.weeklyDictFinances.get(key) + this.reservations[i].Price);
                    }
                }
            },
            weekly() {
                this.loadWeekly();
                this.weeklyVisitors();
                this.weeklyFinances();
            },
            weeklyVisitors () {
                this.dataCollectionVisitors = {
                    labels: Array.from(this.weeklyDictVisitors.keys()),
                    datasets: [
                        {
                            label: 'Weekly vehicles',
                            backgroundColor: '#f87979',
                            data: Array.from(this.weeklyDictVisitors.values())
                        },
                    ],
                    field: "vehicles",
                }
            },
            weeklyFinances () {
                this.dataCollectionFinances = {
                    labels: Array.from(this.weeklyDictVisitors.keys()),
                    datasets: [
                        {
                            label: 'Weekly finance',
                            backgroundColor: '#76c2f8',
                            data: Array.from(this.weeklyDictFinances.values())
                        },
                    ],
                    field: "€",
                }
            },
            loadMonthly () {
                const monthNames = ["January", "February", "March", "April", "May", "June",
                    "July", "August", "September", "October", "November", "December"
                ];
                this.monthlyDictVisitors = new Map();
                this.monthlyDictFinances = new Map();
                if (!moment(this.time.from).isValid() || !moment(this.time.to).isValid()) {
                    return;
                }
                let dateFrom = moment(this.time.from).startOf("month");
                const dateTo = moment(this.time.to).endOf("month");

                while (dateFrom.isSameOrBefore(dateTo)) {
                    this.monthlyDictVisitors.set(monthNames[dateFrom.month()] + ", " + dateFrom.year(), 0);
                    this.monthlyDictFinances.set(monthNames[dateFrom.month()] + ", " + dateFrom.year(), 0);
                    dateFrom.add(1, 'months');
                }

                for (let i = 0; i < this.reservations.length; i++) {
                    const key = monthNames[moment(this.reservations[i].Beginning).month()] + ", " + moment(this.reservations[i].Beginning).year();
                    if (this.monthlyDictVisitors.has(key)) {
                        this.monthlyDictVisitors.set(key, this.monthlyDictVisitors.get(key) + 1);
                        this.monthlyDictFinances.set(key, this.monthlyDictFinances.get(key) + this.reservations[i].Price);
                    }
                }
            },
            monthly() {
                this.loadMonthly();
                this.monthlyVisitors();
                this.monthlyFinances();
            },
            monthlyVisitors () {
                this.dataCollectionVisitors = {
                    labels: Array.from(this.monthlyDictVisitors.keys()),
                    datasets: [
                        {
                            label: 'Monthly vehicles',
                            backgroundColor: '#f87979',
                            data: Array.from(this.monthlyDictVisitors.values())
                        },
                    ],
                    field: "vehicles",
                }
            },
            monthlyFinances () {
                this.dataCollectionFinances = {
                    labels: Array.from(this.monthlyDictVisitors.keys()),
                    datasets: [
                        {
                            label: 'Monthly finance',
                            backgroundColor: '#76c2f8',
                            data: Array.from(this.monthlyDictFinances.values())
                        },
                    ],
                    field: "€",
                }
            },
            formatDate(date) {
                let month = `${date.getMonth() + 1}`;
                let day = `${date.getDate()}`;
                const year = date.getFullYear();

                if (month.length < 2) month = `0${month}`;
                if (day.length < 2) day = `0${day}`;

                return [year, month, day].join('-');
            },
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
        },
        computed: {
            minDate() {
                let minDate = moment(new Date(2100,0,1),"YYYY-MM-DD");
                for (let i = 0; i < this.reservations.length; i++) {
                    const date = moment(this.reservations[i].Beginning,"YYYY-MM-DD");
                    if (date.isBefore(minDate)) {
                        minDate = date;
                    }
                }
                return this.formatDate(minDate.toDate());
            },
            maxDate() {
                let maxDate = moment(new Date(1970,0,1),"YYYY-MM-DD");
                for (let i = 0; i < this.reservations.length; i++) {
                    const date = moment(this.reservations[i].End,"YYYY-MM-DD");
                    if (date.isAfter(maxDate)) {
                        maxDate = date;
                    }
                }
                return this.formatDate(maxDate.toDate());
            },
        }
    }
</script>

<style scoped>

</style>