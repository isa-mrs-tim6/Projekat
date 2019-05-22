<template>
    <v-layout row-wrap>
        <v-flex x12>
            <v-data-table :headers="headers" :items="this.friends" class="elevation-1">
                <template v-slot:items="props">
                    <td>{{props.item.Name}}</td>
                    <td>{{props.item.Surname}}</td>
                    <td>{{props.item.Email}}</td>
                </template>
            </v-data-table>
        </v-flex>
    </v-layout>
</template>

<script>
    import axios from "axios";

    export default {
        name: "UserFriends",
        data (){
            return {
                friends: [],
                headers:[
                    {
                        text: 'Name',
                        value: "Name"

                    },
                    {
                        text: 'Surname',
                        value: "Surname"
                    },
                    {
                        text: 'Email',
                        value: 'Email'
                    }
                ]
            }
        },
        beforeCreate() {
            axios.create({withCredentials: true}).get("http://localhost:8000/api/user/getFriends")
                .then(res=>{
                    this.friends = res.data;
                })
        }
    }
</script>

<style scoped>

</style>