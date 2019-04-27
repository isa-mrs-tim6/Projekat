<template>
    <v-layout justify-center>
        <v-flex xs12>
            <v-toolbar-title>{{title}}</v-toolbar-title>
            <v-container pa-0 mt-2 grid-list-md fluid>
                <v-layout row wrap>
                    <v-flex
                            v-for="item in items"
                            xs4
                    >
                        <template>
                            <v-card height="100%">
                                <v-card-title>{{item.Name}}</v-card-title>
                                <v-card-text>
                                    <v-img src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAQMAAADCCAMAAAB6zFdcAAAAYFBMVEXa2tpVVVXd3d1MTExSUlK2trZvb29LS0tTU1Ph4eGNjY2cnJzU1NRaWlpgYGBPT0+np6fGxsavr6/AwMCGhoaioqJpaWnMzMx+fn6UlJS7u7t1dXWzs7NERERkZGSdnZ1LtC8/AAACkElEQVR4nO3b626qQBRAYeYiM2NVVPCCWvv+b3lEUUDhpIpJ42Z9/0qFOCsTYFCjCAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAwCPqN/nosr9Grdfw+4a+H8xL95bx9F/+pDRL1PjSgAQ0kNLDOuFeZxHgBDZJ8N+phNxPQwCxDn3ujMPMCGsx7vX0hDXrd54pokDzZ4G5pMLwGWqfL5aa+w+Aa6EV8upDarLZpaA30wtvzKWRfbRtcg315Z+mqfYbWYHO9ufbxUBvokbmuDqaiG/wvSHprMJPcQI+2nRV0VK4NVJLdhiyxwdqNO4ej59+X04HfCJ4HemmUybsj5C7xiVG76hXyGkTH04jMT3eENIvjfBOqXcQ10KvzWc+tuiPoEEJ9ySCuQTS9DKh4ptB4afdRpDUIeVI+GXOH+pj0rjuCtAaT6+VfeVsf9kRtO6eCsAYhqz1odmk1PbIkWU86IshqoDemSqDsdHHb7k5/HjftEYQ12DY+b7C2HHXYFwtmm6StwxTVQKdONdjZOYI+XKaH9cu2cYpqEGKr7iNMin+Ug+y4bZDUoLYwvinOhPqr2m7Gj+cESQ2itX9ooGwcTWqbvXm8RgpqEOatH8Am26yx/TwzGkcR1ECrlmlQzIS7NHa2kNogrB7PBu1sfeEcSWowmf4yQbFD4xoppoFe/HYaFFweqlOjoAZPfSXFZAKfoTzZQJn9bQklqUH7ZaGDN6clVHl7KaeBnz4rvRxFTIPTheF5l6MIavAyGohpEHo5Cmjg9+M+8vIO86MbnJZFvSgJDd6DBjSgwUc3yL/NuzjzmQ2ixaHPl/abDh/64z5+3wgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAz/kH0Z07XHC7XZ8AAAAASUVORK5CYII="></v-img>
                                    <v-label>{{item.Address}}</v-label><br/>
                                    <v-label>{{item.Promo || item.Description}}</v-label>
                                </v-card-text>
                            </v-card>
                        </template>
                    </v-flex>
                </v-layout>
            </v-container>
        </v-flex>
    </v-layout>
</template>

<script>
    export default {
        name: "ResultGrid",
        props: ["items", "title"],
        data () {
            return {

            }
        },
        methods: {

        },
    }
</script>

<style scoped>

</style>