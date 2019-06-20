<template>
    <v-container fill-height fluid id="main">
        <v-layout>
            <v-flex>
                <airline-admin-nav-drawer/>
            </v-flex>
            <v-container grid-list-xl text-xs-center fill-height>
                <v-layout align-center justify-center column wrap fill-height>
                    <v-flex style="width: 90vw">
                        <v-card min-height="100%" class="flexcard" style="padding: 5px">
                            <v-card-title primary-title>
                                <div class="headline font-weight-medium">Airplanes:</div>
                            </v-card-title>
                            <v-card-text class="grow">
                                <v-layout row-wrap>
                                    <v-flex xs9 mr-3>
                                        <v-layout row-wrap>
                                            <v-flex x12>
                                                <v-data-table :headers="headers" :items="items" class="elevation-1">
                                                    <template v-slot:items="props">
                                                        <td>{{props.item.Name}}</td>
                                                        <td>{{props.item.RowWidth}}</td>
                                                        <td>{{props.item.First}}</td>
                                                        <td>{{props.item.Business}}</td>
                                                        <td>{{props.item.Economy}}</td>
                                                    </template>
                                                </v-data-table>
                                            </v-flex>
                                        </v-layout>
                                    </v-flex>
                                    <v-flex xs3>
                                        <h2 style="margin-bottom: 10px">Add An Airplane:</h2>
                                        <v-form ref="form">
                                            <v-text-field label="Name" v-model="Name"></v-text-field>
                                            <v-text-field label="Row Width" v-model="RowWidth"></v-text-field>
                                            <v-text-field label="First Class Num" v-model="First"></v-text-field>
                                            <v-text-field label="Business Class Num" v-model="Business"></v-text-field>
                                            <v-text-field label="Economy num" v-model="Economy"></v-text-field>
                                        </v-form>
                                        <v-btn color="primary" style="float: left" @click="addAirplane()" >submit</v-btn>
                                        <v-btn style="float: right" @click="clear()">clear</v-btn>
                                    </v-flex>
                                </v-layout>
                            </v-card-text>
                        </v-card>
                    </v-flex>
                </v-layout>
            </v-container>
        </v-layout>
        <v-snackbar v-model="SuccessSnackbar" :timeout=2000 :top="true" color="success">{{SuccessSnackbarMessage}}</v-snackbar>
        <v-snackbar v-model="ErrorSnackbar" :timeout=2000 :top="true" color="error">{{ErrorSnackbarMessage}}</v-snackbar>
    </v-container>
</template>

<script>
    import AirlineAdminNavDrawer from "../../components/AirlineAdmin/AirlineAdminNavDrawer";
    import axios from "axios";

    export default {
        name: "AddAirplane",
        components: {AirlineAdminNavDrawer},
        data (){
            return {
                SuccessSnackbar: false,
                SuccessSnackbarMessage: "",
                ErrorSnackbar: false,
                ErrorSnackbarMessage: "",
                Name: '',
                RowWidth: '',
                First: '',
                Business: '',
                Economy: '',
                items: [],
                headers:[
                    {
                        text: 'Name',
                        value: "Name",
                        width: '20%'

                    },
                    {
                        text: 'Row Width',
                        value: "RowWidth",
                        width: '20%'
                    },
                    {
                        text: 'First',
                        value: "First",
                        width: '20%'
                    },
                    {
                        text: 'Business',
                        value: "Business",
                        width: '20%'
                    },
                    {
                        text: 'Economy',
                        value: "Economy",
                        width: '20%'
                    }
                ]
            }
        },
        beforeCreate() {
            axios.create({withCredentials: true}).get("http://ec2-35-159-21-254.eu-central-1.compute.amazonaws.com:8000/api/airplane/getAirplanesAndSeats")
                .then(res => {
                    this.items = res.data;
                })
        },
        mounter(){
            this.checkFirstPass();
        },
        methods: {
            checkFirstPass(){
                axios.create({withCredentials: true}).get("http://ec2-35-159-21-254.eu-central-1.compute.amazonaws.com:8000/api/admin/checkFirstPass")
                    .then(
                        res =>{
                            if(res.data === false){
                                alert("Please update your password before accessing this feature");
                                this.$router.replace("admin_profile");
                            }
                        }
                    )
            },
            clear() {
                this.$refs.form.reset();
            },
            addAirplane(){
                if(!this.Name || !this.RowWidth || !this.First || !this.Business || !this.Economy){
                    this.ErrorSnackbarMessage = "All fields must be entered!";
                    this.ErrorSnackbar = true;
                    return;
                }
                if(isNaN(parseInt(this.RowWidth))){
                    this.ErrorSnackbarMessage = "Row Width must be integer!";
                    this.ErrorSnackbar = true;
                    return;
                }else{
                    if(parseInt(this.RowWidth) < 1){
                        this.ErrorSnackbarMessage = "Row Width must be positive number!";
                        this.ErrorSnackbar = true;
                        return;
                    }
                }
                if(isNaN(parseInt(this.First))){
                    this.ErrorSnackbarMessage = "First class seat num must be integer!";
                    this.ErrorSnackbar = true;
                    return;
                }else{
                    if(parseInt(this.First) < 1){
                        this.ErrorSnackbarMessage = "First class seat num must be positive number!";
                        this.ErrorSnackbar = true;
                        return;
                    }
                }
                if(isNaN(parseInt(this.Business))){
                    this.ErrorSnackbarMessage = "Business class seat num must be integer!";
                    this.ErrorSnackbar = true;
                    return;
                }else{
                    if(parseInt(this.Business) < 1){
                        this.ErrorSnackbarMessage = "Business class seat num must be positive number!";
                        this.ErrorSnackbar = true;
                        return;
                    }
                }
                if(isNaN(parseInt(this.Economy))){
                    this.ErrorSnackbarMessage = "Economy class seat num must be integer!";
                    this.ErrorSnackbar = true;
                    return;
                }else{
                    if(parseInt(this.Economy) < 1){
                        this.ErrorSnackbarMessage = "Economy class seat num must be positive number!";
                        this.ErrorSnackbar = true;
                        return;
                    }
                }
                let params = {
                    Name: this.Name,
                    RowWidth: parseInt(this.RowWidth),
                    First: parseInt(this.First),
                    Business: parseInt(this.Business),
                    Economy: parseInt(this.Economy)
                };
                axios.create({withCredentials: true}).post("http://ec2-35-159-21-254.eu-central-1.compute.amazonaws.com:8000/api/airplane/addAirplane", params)
                    .then(res => {
                        axios.create({withCredentials: true}).get("http://ec2-35-159-21-254.eu-central-1.compute.amazonaws.com:8000/api/airplane/getAirplanesAndSeats")
                            .then(res => {
                                this.clear();
                                this.items = res.data;
                                this.SuccessSnackbarMessage = "Successfully added airplane!";
                                this.SuccessSnackbar = true;
                            })
                    })
            }
        },
    }
</script>

<style scoped>
    #main {
        background-image: linear-gradient(to right bottom, #142eae, #005bca, #007ed2, #009ccd, #0bb7c7, #47c0c6, #67c8c6, #81d0c7, #6ecac4, #58c4c3, #3cbdc2, #00b7c1);
    }
    .flexcard {
        display: flex;
        flex-direction: column;
    }
</style>