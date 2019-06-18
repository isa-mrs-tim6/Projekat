<template>
    <div>
        <v-card v-if="Reservations.length === 0" style="margin: 5px" dark>
            <v-card-text>
                <v-layout row wrap>
                    <v-flex>
                        There are no quick reservations for this date range.
                    </v-flex>
                </v-layout>
            </v-card-text>
        </v-card>
        <v-card v-if="Reservations.length !== 0" style="margin: 5px" dark v-for="(item, index) in this.Reservations" :key="item.ID">
            <v-card-text>
                <v-layout row wrap>
                    <v-flex xs2>
                        <v-img src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAQMAAADCCAMAAAB6zFdcAAAAYFBMVEXa2tpVVVXd3d1MTExSUlK2trZvb29LS0tTU1Ph4eGNjY2cnJzU1NRaWlpgYGBPT0+np6fGxsavr6/AwMCGhoaioqJpaWnMzMx+fn6UlJS7u7t1dXWzs7NERERkZGSdnZ1LtC8/AAACkElEQVR4nO3b626qQBRAYeYiM2NVVPCCWvv+b3lEUUDhpIpJ42Z9/0qFOCsTYFCjCAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAwCPqN/nosr9Grdfw+4a+H8xL95bx9F/+pDRL1PjSgAQ0kNLDOuFeZxHgBDZJ8N+phNxPQwCxDn3ujMPMCGsx7vX0hDXrd54pokDzZ4G5pMLwGWqfL5aa+w+Aa6EV8upDarLZpaA30wtvzKWRfbRtcg315Z+mqfYbWYHO9ufbxUBvokbmuDqaiG/wvSHprMJPcQI+2nRV0VK4NVJLdhiyxwdqNO4ej59+X04HfCJ4HemmUybsj5C7xiVG76hXyGkTH04jMT3eENIvjfBOqXcQ10KvzWc+tuiPoEEJ9ySCuQTS9DKh4ptB4afdRpDUIeVI+GXOH+pj0rjuCtAaT6+VfeVsf9kRtO6eCsAYhqz1odmk1PbIkWU86IshqoDemSqDsdHHb7k5/HjftEYQ12DY+b7C2HHXYFwtmm6StwxTVQKdONdjZOYI+XKaH9cu2cYpqEGKr7iNMin+Ug+y4bZDUoLYwvinOhPqr2m7Gj+cESQ2itX9ooGwcTWqbvXm8RgpqEOatH8Am26yx/TwzGkcR1ECrlmlQzIS7NHa2kNogrB7PBu1sfeEcSWowmf4yQbFD4xoppoFe/HYaFFweqlOjoAZPfSXFZAKfoTzZQJn9bQklqUH7ZaGDN6clVHl7KaeBnz4rvRxFTIPTheF5l6MIavAyGohpEHo5Cmjg9+M+8vIO86MbnJZFvSgJDd6DBjSgwUc3yL/NuzjzmQ2ixaHPl/abDh/64z5+3wgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAz/kH0Z07XHC7XZ8AAAAASUVORK5CYII="></v-img>
                    </v-flex>
                    <v-spacer></v-spacer>
                    <v-flex xs9>
                        <v-layout row>
                            <v-flex xs8 class="center">
                                {{item.Vehicle.Name}}
                            </v-flex>
                            <v-spacer></v-spacer>
                            <v-flex xs3 style="text-align: right;">
                                <v-btn @click="quickReserveVehicle(index)">choose</v-btn>
                            </v-flex>
                        </v-layout>
                        <v-layout row>
                            <v-flex xs3 class="center">
                                Type: {{item.Vehicle.Type}}<br/>
                                Capacity: {{item.Vehicle.Capacity}}<v-icon large>person</v-icon><br/>
                                Start: {{computedDateFormattedMomentjs(item.Beginning)}} <br/>
                                End: {{computedDateFormattedMomentjs(item.End)}}
                            </v-flex>
                            <v-spacer></v-spacer>
                            <v-flex xs3 class="center price">
                                Total:<br/>
                                <v-icon large>attach_money</v-icon>
                                {{item.Price / item.QuickReservationPriceScale}}
                            </v-flex>
                            <v-flex xs3 class="center price">
                                Your price:<br/>
                                <v-icon large>attach_money</v-icon>
                                {{item.Price}}
                            </v-flex>
                        </v-layout>
                    </v-flex>
                </v-layout>
            </v-card-text>
        </v-card>
        <v-snackbar v-model="SuccessSnackbar" :timeout=2000 :top="true" color="success">{{this.SuccessSnackbarText}}</v-snackbar>
        <v-snackbar v-model="ErrorSnackbar" :timeout=2000 :top="true" color="error">{{this.ErrorSnackbarText}}</v-snackbar>
    </div>
</template>

<script>
    import axios from "axios";
    import moment from "moment";

    export default {
        name: "QuickVehicles",
        data (){
            return {
                SuccessSnackbar: false,
                SuccessSnackbarText: '',
                ErrorSnackbar: false,
                ErrorSnackbarText: '',
                LocationID: parseInt(this.$route.params.locID),
                CompanyID: parseInt(this.$route.params.id),
                Start: this.$route.query.start,
                End: this.$route.query.end,
                Reservations: []
            }
        },
        beforeMount() {
            let params = {
                CompanyID: this.CompanyID,
                StartDate: moment(this.Start,"YYYY-MM-DD").valueOf().toString(),
                EndDate: moment(this.End,"YYYY-MM-DD").valueOf().toString(),
            };
            console.log(params);
            axios.create({withCredentials: true}).post("http://localhost:8000/api/reservations/rac/quickReservationsCompany", params)
                .then(res => {
                    this.Reservations = res.data;
                })
        },
        methods: {
            computedDateFormattedMomentjs(date) {
                return date ? moment(date).format('DD-MM-YYYY') : ''
            },
            quickReserveVehicle(index){
                let res = {
                    ReservationID: this.Reservations[index].ID,
                    MasterID: Number.parseInt(this.$route.query.reservationID),
                    LocationId: this.LocationID
                };
                axios.create({withCredentials: true}).post('http://localhost:8000/api/reservation/rac/completeQuickRes', res)
                    .then(res =>{
                        this.SuccessSnackbar = true;
                        this.SuccessSnackbarText = 'Reservation complete';
                        this.$router.push({ path: '/user/reserve', query: { reservationID: this.reservationID, passengers: this.passengers }});
                    })
                    .catch(err => {
                        this.ErrorSnackbar = true;
                        this.ErrorSnackbarText = 'Please log in to use this feature.';
                    })
            },
        }
    }
</script>

<style scoped>
    .center {
        margin: auto;
        width: 50%;
        padding: 10px;
        font-size: 24px;
        word-wrap: break-word;
    }
    .price{
        font-size: 36px;
    }
</style>