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
                <FinanceGraph v-bind:item="item.tickets" ref="chart" :chart-data="dataCollectionTickets"></FinanceGraph>
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
    import moment from 'moment';
    import FinanceGraph from '../HotelAdmin/FinanceGraph'
    export default {
        name: "AirlineFinance",
        components: {FinanceGraph},
        props: ["graphData"],
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
                dataCollectionTickets: null,
                dataCollectionFinances: null,
                monthlyDictTickets: new Map(),
                monthlyDictFinances: new Map(),
                weeklyDictTickets: new Map(),
                weeklyDictFinances: new Map(),
                dailyDictTickets: new Map(),
                dailyDictFinances: new Map(),
                item: {
                    tickets: 'tickets',
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
                this.dailyDictTickets = new Map();
                this.dailyDictFinances = new Map();
                if (!moment(this.time.from).isValid() || !moment(this.time.to).isValid()) {
                    return;
                }
                let dateFrom = moment(this.time.from);
                const dateTo = moment(this.time.to);

                while (dateFrom.isSameOrBefore(dateTo)) {
                    this.dailyDictTickets.set(dateFrom.format('DD-MM-YYYY'), 0);
                    this.dailyDictFinances.set(dateFrom.format('DD-MM-YYYY'), 0);
                    dateFrom.add(1, 'days');
                }

                for (let i = 0; i < this.graphData.length; i++) {
                    const key = moment(this.graphData[i].Departure).format('DD-MM-YYYY');
                    if (this.dailyDictTickets.has(key)) {
                        this.dailyDictTickets.set(key, this.dailyDictTickets.get(key) + 1);
                        this.dailyDictFinances.set(key, this.dailyDictFinances.get(key) + this.graphData[i].Price);
                    }
                }
            },
            daily() {
                this.loadDaily();
                this.dailyVisitors();
                this.dailyFinances();
            },
            dailyVisitors () {
                this.dataCollectionTickets = {
                    labels: Array.from(this.dailyDictTickets.keys()),
                    datasets: [
                        {
                            label: 'Daily tickets',
                            backgroundColor: '#f87979',
                            data: Array.from(this.dailyDictTickets.values())
                        },
                    ],
                    field: "tickets",
                }
            },
            dailyFinances () {
                this.dataCollectionFinances = {
                    labels: Array.from(this.dailyDictTickets.keys()),
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
                this.weeklyDictTickets = new Map();
                this.weeklyDictFinances = new Map();
                if (!moment(this.time.from).isValid() || !moment(this.time.to).isValid()) {
                    return;
                }
                let dateFrom = moment(this.time.from);
                const dateTo = moment(this.time.to);

                while (dateFrom.isoWeek() <=dateTo.isoWeek()) {
                    this.weeklyDictTickets.set("Week " + dateFrom.isoWeek() + ", " + dateFrom.year(), 0);
                    this.weeklyDictFinances.set("Week " + dateFrom.isoWeek() + ", " + dateFrom.year(), 0);
                    dateFrom.add(1, 'weeks');
                }
                for (let i = 0; i < this.graphData.length; i++) {
                    const key = "Week " + moment(this.graphData[i].Departure).isoWeek() + ", " +  moment(this.graphData[i].Departure).year();
                    if (this.weeklyDictTickets.has(key)) {
                        this.weeklyDictTickets.set(key, this.weeklyDictTickets.get(key) + 1);
                        this.weeklyDictFinances.set(key, this.weeklyDictFinances.get(key) + this.graphData[i].Price);
                    }
                }
            },
            weekly() {
                this.loadWeekly();
                this.weeklyVisitors();
                this.weeklyFinances();
            },
            weeklyVisitors () {
                this.dataCollectionTickets = {
                    labels: Array.from(this.weeklyDictTickets.keys()),
                    datasets: [
                        {
                            label: 'Weekly visitors',
                            backgroundColor: '#f87979',
                            data: Array.from(this.weeklyDictTickets.values())
                        },
                    ],
                    field: "tickets",
                }
            },
            weeklyFinances () {
                this.dataCollectionFinances = {
                    labels: Array.from(this.weeklyDictTickets.keys()),
                    datasets: [
                        {
                            label: 'Weekly tickets',
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
                this.monthlyDictTickets = new Map();
                this.monthlyDictFinances = new Map();
                if (!moment(this.time.from).isValid() || !moment(this.time.to).isValid()) {
                    return;
                }
                let dateFrom = moment(this.time.from);
                let dateFromIncrement = moment(this.time.from).startOf("month");
                const dateTo = moment(this.time.to).endOf("day");

                while (dateFromIncrement.isSameOrBefore(dateTo)) {
                    this.monthlyDictTickets.set(monthNames[dateFromIncrement.month()] + ", " + dateFromIncrement.year(), 0);
                    this.monthlyDictFinances.set(monthNames[dateFromIncrement.month()] + ", " + dateFromIncrement.year(), 0);
                    dateFromIncrement.add(1, 'months');
                }

                for (let i = 0; i < this.graphData.length; i++) {
                    if(moment(this.graphData[i].Departure).isBefore(dateFrom)){
                        continue
                    }
                    if(moment(this.graphData[i].Departure).isAfter(dateTo)){
                        continue
                    }
                    const key = monthNames[moment(this.graphData[i].Departure).month()] + ", " + moment(this.graphData[i].Departure).year();
                    if (this.monthlyDictTickets.has(key)) {
                        this.monthlyDictTickets.set(key, this.monthlyDictTickets.get(key) + 1);
                        this.monthlyDictFinances.set(key, this.monthlyDictFinances.get(key) + this.graphData[i].Price);
                    }
                }
            },
            monthly() {
                this.loadMonthly();
                this.monthlyVisitors();
                this.monthlyFinances();
            },
            monthlyVisitors () {
                this.dataCollectionTickets = {
                    labels: Array.from(this.monthlyDictTickets.keys()),
                    datasets: [
                        {
                            label: 'Monthly tickets',
                            backgroundColor: '#f87979',
                            data: Array.from(this.monthlyDictTickets.values())
                        },
                    ],
                    field: "tickets",
                }
            },
            monthlyFinances () {
                this.dataCollectionFinances = {
                    labels: Array.from(this.monthlyDictTickets.keys()),
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
                for (let i = 0; i < this.graphData.length; i++) {
                    const date = moment(this.graphData[i].Departure,"YYYY-MM-DD");
                    if (date.isBefore(minDate)) {
                        minDate = date;
                    }
                }
                return this.formatDate(minDate.toDate());
            },
            maxDate() {
                let maxDate = moment(new Date(1970,0,1),"YYYY-MM-DD");
                for (let i = 0; i < this.graphData.length; i++) {
                    const date = moment(this.graphData[i].Departure,"YYYY-MM-DD");
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