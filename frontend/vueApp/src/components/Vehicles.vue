<template>
    <v-container>
        <v-card style="margin: 5px" dark v-for="item in vehicles">
            <v-card-text>
                <v-layout row wrap>
                    <template>
                        <v-flex xs2>
                            <v-img src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAQMAAADCCAMAAAB6zFdcAAAAYFBMVEXa2tpVVVXd3d1MTExSUlK2trZvb29LS0tTU1Ph4eGNjY2cnJzU1NRaWlpgYGBPT0+np6fGxsavr6/AwMCGhoaioqJpaWnMzMx+fn6UlJS7u7t1dXWzs7NERERkZGSdnZ1LtC8/AAACkElEQVR4nO3b626qQBRAYeYiM2NVVPCCWvv+b3lEUUDhpIpJ42Z9/0qFOCsTYFCjCAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAwCPqN/nosr9Grdfw+4a+H8xL95bx9F/+pDRL1PjSgAQ0kNLDOuFeZxHgBDZJ8N+phNxPQwCxDn3ujMPMCGsx7vX0hDXrd54pokDzZ4G5pMLwGWqfL5aa+w+Aa6EV8upDarLZpaA30wtvzKWRfbRtcg315Z+mqfYbWYHO9ufbxUBvokbmuDqaiG/wvSHprMJPcQI+2nRV0VK4NVJLdhiyxwdqNO4ej59+X04HfCJ4HemmUybsj5C7xiVG76hXyGkTH04jMT3eENIvjfBOqXcQ10KvzWc+tuiPoEEJ9ySCuQTS9DKh4ptB4afdRpDUIeVI+GXOH+pj0rjuCtAaT6+VfeVsf9kRtO6eCsAYhqz1odmk1PbIkWU86IshqoDemSqDsdHHb7k5/HjftEYQ12DY+b7C2HHXYFwtmm6StwxTVQKdONdjZOYI+XKaH9cu2cYpqEGKr7iNMin+Ug+y4bZDUoLYwvinOhPqr2m7Gj+cESQ2itX9ooGwcTWqbvXm8RgpqEOatH8Am26yx/TwzGkcR1ECrlmlQzIS7NHa2kNogrB7PBu1sfeEcSWowmf4yQbFD4xoppoFe/HYaFFweqlOjoAZPfSXFZAKfoTzZQJn9bQklqUH7ZaGDN6clVHl7KaeBnz4rvRxFTIPTheF5l6MIavAyGohpEHo5Cmjg9+M+8vIO86MbnJZFvSgJDd6DBjSgwUc3yL/NuzjzmQ2ixaHPl/abDh/64z5+3wgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAz/kH0Z07XHC7XZ8AAAAASUVORK5CYII=">   </v-img>
                        </v-flex>
                        <v-spacer></v-spacer>
                        <v-flex xs7>
                            <v-layout row>
                                <v-flex class="center">{{item.Name}}</v-flex>
                            </v-layout>
                            <v-layout row>
                                <v-flex class="center">{{item.Type}}</v-flex>
                                <v-flex class="center"><v-icon large>person</v-icon>{{item.Capacity}}</v-flex>
                            </v-layout>
                        </v-flex>
                        <v-flex xs2 class="center price"><v-icon large>attach_money</v-icon>{{item.PricePerDay}}</v-flex>
                    </template>
                </v-layout>
            </v-card-text>
        </v-card>
    </v-container>
</template>

<script>
    export default {
        name: "Vehicles",
        props: ["vehicles"],
        data() {
            return {
                headers: [
                    { text: 'Name', value: 'Name' },
                    { text: 'Capacity', value: 'Capacity' },
                    { text: 'Type', value: 'Type'},
                    { text: 'PricePerDay', value: 'PricePerDay' },
                ],
            }
        },
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