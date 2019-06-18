<template>
    <div>
        <v-container fluid>
            <v-card style="margin: 5px" dark v-for="(item, index) in vehicles" :key="item.Vehicle.ID">
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
                                    <v-btn @click="reserveVehicle(index)">choose</v-btn>
                                </v-flex>
                            </v-layout>
                            <v-layout row>
                                <v-flex xs6 class="center">
                                    Type: {{item.Vehicle.Type}}<br/>
                                    Capacity: {{item.Vehicle.Capacity}}<v-icon large>person</v-icon><br/>
                                    Rating: <v-rating
                                        v-model="item.Rating"
                                        background-color="indigo lighten-3"
                                        empty-icon="$vuetify.icons.ratingFull"
                                        readonly
                                        half-increments
                                        color="orange"
                                        style="display:inline;"
                                ></v-rating>
                                </v-flex>
                                <v-spacer></v-spacer>
                                <v-flex xs3 class="center price">
                                    Total:<br/>
                                    <v-icon large>attach_money</v-icon>
                                    {{item.Vehicle.PricePerDay}}
                                </v-flex>
                            </v-layout>
                        </v-flex>
                    </v-layout>
                </v-card-text>
            </v-card>
        </v-container>
        <v-snackbar v-model="SuccessSnackbar" :timeout=2000 :top="true" color="success">{{this.SuccessSnackbarText}}</v-snackbar>
        <v-snackbar v-model="ErrorSnackbar" :timeout=2000 :top="true" color="error">{{this.ErrorSnackbarText}}</v-snackbar>
    </div>
</template>

<script>
    import axios from "axios";

    export default {
        name: "Vehicles",
        props: ["vehicles", "startDate", "endDate", "reservationID", "passengers"],
        data (){
          return {
              SuccessSnackbar: false,
              SuccessSnackbarText: '',
              ErrorSnackbar: false,
              ErrorSnackbarText: '',
          };
        },
        methods:{
            reserveVehicle(index){
                let res = {
                    VehicleID: this.$props.vehicles[index].Vehicle.ID,
                    LocationID: parseInt(this.$route.params.locID),
                    CompanyID: parseInt(this.$route.params.id),
                    Price: this.$props.vehicles[index].Vehicle.PricePerDay,
                    StartDate: this.$props.startDate,
                    EndDate: this.$props.endDate
                };
                axios.create({withCredentials: true}).post('http://localhost:8000/api/rentACarCompany/' + this.reservationID + '/reserveVehicle', res)
                    .then(res =>{
                        this.SuccessSnackbar = true;
                        this.SuccessSnackbarText = 'Reservation complete';
                        this.$router.push({ path: '/user/reserve', query: { reservationID: this.reservationID, passengers: this.passengers }});
                    })
                    .catch(err => {
                        let vm = this
                        if(err.response.status === 401){
                            vm.ErrorSnackbar = true;
                            vm.ErrorSnackbarText = 'Please log in to use this feature.';
                        }else if(err.response.status === 400){
                            vm.ErrorSnackbar = true;
                            vm.ErrorSnackbarText = 'Someone else has already taken that vehicle.';
                            setTimeout(function () {
                                vm.$router.go()
                            }, 4000);
                        }else if(err.response.status === 500){
                            vm.ErrorSnackbar = true;
                            vm.ErrorSnackbarText = 'Please reserve a flight first';
                        }

                    })
            },
            goLogin (e){
                e.preventDefault();
                this.$router.push("/login");
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