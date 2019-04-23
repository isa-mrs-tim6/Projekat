<template>
    <v-container>
        <v-card min-height="100%" class="flexcard" style="padding: 5px">
            <v-card-title primary-title>
                <div class="headline font-weight-medium">Manage Offices:</div>
            </v-card-title>
            <v-card-text class="grow">
                <v-layout row-wrap>
                    <v-flex xs9 mr-3>
                        <h2>Filters:</h2>
                        <v-layout row-wrap>
                            <v-flex xs3>
                                <v-text-field v-model="filter.address" label="Address"></v-text-field>
                            </v-flex>
                        </v-layout>
                        <h2>Offices:</h2>
                        <v-layout row-wrap>
                            <v-flex x12>
                                <v-data-table :headers="headers" :items="officesToShow" class="elevation-1">
                                    <template v-slot:items="props">
                                        <td>{{props.item.Address}}</td>
                                        <td>{{props.item.Longitude}}</td>
                                        <td>{{props.item.Latitude}}</td>
                                        <td>
                                            <v-icon small class="mr-2" @click="editItem(props.item)">edit</v-icon>
                                            <v-icon small class="mr-2" @click="deleteItem(props.item)">delete</v-icon>
                                        </td>
                                    </template>
                                </v-data-table>
                            </v-flex>
                        </v-layout>
                    </v-flex>
                    <v-flex xs3>
                        <h2>Add Office:</h2>
                        <v-form ref="form" class="align-center justify-center">
                            <v-text-field
                                    v-model="addedItem.Address"
                                    label="Address"
                                    :rules="[rules.required]"
                                    required
                            ></v-text-field>
                            <v-text-field
                                    v-model="addedItem.Longitude"
                                    label="Longitude"
                                    :rules="[rules.required]"
                                    required
                            ></v-text-field>
                            <v-text-field
                                    v-model="addedItem.Latitude"
                                    label="Latitude"
                                    :rules="[rules.required]"
                                    required
                            ></v-text-field>
                            <v-btn color="primary" style="float: left" @click="addLocation">submit</v-btn>
                            <v-btn style="float: right" @click="clear">clear</v-btn>
                        </v-form>
                    </v-flex>
                </v-layout>
            </v-card-text>
        </v-card>
        <v-dialog v-model="dialog" persistent max-width="500px">
            <v-card>
                <v-card-title>
                    <span class="headline">Edit office info:</span>
                </v-card-title>
                <v-card-text>
                    <v-container>
                        <v-layout row-wrap>
                            <v-flex xs6>
                                <v-text-field v-model="editedItem.Address" label="Address" :rules="[rules.required]"></v-text-field>
                            </v-flex>
                        </v-layout>
                        <v-layout row-wrap>
                            <v-flex xs6>
                                <v-text-field v-model="editedItem.Longitude" label="Longitude" :rules="[rules.required]"></v-text-field>
                            </v-flex>
                        </v-layout>
                        <v-layout row-wrap>
                            <v-flex xs6>
                                <v-text-field v-model="editedItem.Latitude" label="Latitude" :rules="[rules.required]"></v-text-field>
                            </v-flex>
                        </v-layout>
                    </v-container>
                </v-card-text>
                <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn color="blue darken-1" flat @click="close">Cancel</v-btn>
                    <v-btn color="blue darken-1" flat @click="save">Save</v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>
        <v-snackbar v-model="SuccessSnackbar" :timeout=2000 :top="true" color="success">{{this.SuccessSnackbarText}}</v-snackbar>
        <v-snackbar v-model="ErrorSnackbar" :timeout=2000 :top="true" color="error">{{this.ErrorSnackbarText}}</v-snackbar>
    </v-container>
</template>

<script>
    import axios from "axios";
    export default {
        name: "ManageOffices",
        data () {
            return {
                SuccessSnackbar: false,
                SuccessSnackbarText: '',
                ErrorSnackbar: false,
                ErrorSnackbarText: '',
                rules:{
                    required: value => !!value || 'Required.'
                },
                dialog: false,
                offices:[],
                filter: {
                    address: ''
                },
                addedItem:{
                    Address: '',
                    Longitude: '',
                    Latitude: ''
                },
                editedItem:{
                    ID: '',
                    Address: '',
                    Longitude: '',
                    Latitude: ''
                },
                defaultItem:{
                    ID: '',
                    Address: '',
                    Longitude: '',
                    Latitude: ''
                },
                editedIndex: -1,
                headers:[
                    {
                        text: 'Address',
                        value: "Address",
                        width: '45%'

                    },
                    {
                        text: 'Longitude',
                        value: "Longitude",
                        width: '20%'
                    },
                    {
                        text: 'Latitude',
                        value: "Latitude",
                        width: '20%'
                    },
                    {
                        text: 'Action',
                        sortable: false,
                        width: '15%'
                    }
                ]
            }
        },
        created(){
            axios.create({withCredentials: true}).get("http://localhost:8000/api/rentACarCompany/getCompanyLocations")
                .then(res => {
                    this.offices = res.data;
                });
        },
        methods:{
            deleteItem(item){
                this.editedIndex = this.offices.indexOf(item);
                this.editedItem = Object.assign({}, item);
                let office = {
                    ID: this.editedItem.ID,
                    Address: this.editedItem.Address,
                    Longitude: parseFloat(this.editedItem.Longitude),
                    Latitude: parseFloat(this.editedItem.Latitude)
                };
                Object.assign(this.offices[this.editedIndex], this.editedItem);
                axios.create({withCredentials:true}).post("http://localhost:8000/api/rentACarCompany/deleteLocation", office)
                    .then(res=>{
                        axios.create({withCredentials: true}).get("http://localhost:8000/api/rentACarCompany/getCompanyLocations")
                            .then(res => {
                                this.offices = res.data;
                            });
                        this.SuccessSnackbar = true;
                        this.SuccessSnackbarText = 'Office successfully deleted';
                    })
                    .catch(err =>{
                        this.ErrorSnackbar = true;
                        this.ErrorSnackbarText = 'Cannot delete office';
                    });
            },
            addLocation(e){
                e.preventDefault();
                if (!this.checkAddInput()){
                    this.ErrorSnackbar = true;
                    this.ErrorSnackbarText = 'Cannot add office';
                    return;
                }
                let location = {
                    Address: this.addedItem.Address,
                    Longitude: parseFloat(this.addedItem.Longitude),
                    Latitude: parseFloat(this.addedItem.Latitude)
                };
                axios.create({withCredentials: true}).post('http://localhost:8000/api/rentACarCompany/addLocation', location)
                    .then(res =>{
                        axios.create({withCredentials: true}).get('http://localhost:8000/api/rentACarCompany/getCompanyLocations')
                            .then(res => this.offices = res.data);
                        this.SuccessSnackbar = true;
                        this.SuccessSnackbarText = 'Office successfully added';
                    })
                    .catch(err => {
                        this.ErrorSnackbar = true;
                        this.ErrorSnackbarText = 'Cannot add office';
                    });
                this.clear();
            },
            clear() {
                this.$refs.form.reset();
            },
            checkAddInput(){
                if(!this.addedItem.Address || !this.addedItem.Longitude || !this.addedItem.Latitude){
                    return false;
                }
                let ind = false;
                let number = parseFloat(this.addedItem.Longitude);
                if(isNaN(number)){
                    this.addedItem.Longitude = "";
                    ind = true;
                }
                number = parseFloat(this.addedItem.Latitude);
                if(isNaN(number)){
                    this.addedItem.Latitude = "";
                    ind = true;
                }
                return ind !== true;
            },
            checkInput(){
                if(!this.editedItem.Address || !this.editedItem.Longitude || !this.editedItem.Latitude){
                    return false;
                }
                let ind = false;
                let number = parseFloat(this.editedItem.Longitude);
                if(isNaN(number)){
                    this.editedItem.Longitude = "";
                    ind = true;
                }
                number = parseFloat(this.editedItem.Latitude);
                if(isNaN(number)){
                    this.editedItem.Latitude = "";
                    ind = true;
                }
                return ind !== true;
            },
            editItem (item) {
                this.editedIndex = this.offices.indexOf(item);
                this.editedItem = Object.assign({}, item);
                this.dialog = true
            },
            close(){
                this.dialog = false;
                setTimeout(()=>{
                    this.editedItem = Object.assign({}, this.defaultItem);
                    this.editedIndex = -1;
                }, 300);
            },
            save(){
                if (!this.checkInput()){
                    this.ErrorSnackbar = true;
                    this.ErrorSnackbarText = 'Cannot edit office';
                    return;
                }
                let office = {
                    ID: this.editedItem.ID,
                    Address: this.editedItem.Address,
                    Longitude: parseFloat(this.editedItem.Longitude),
                    Latitude: parseFloat(this.editedItem.Latitude)
                };
                Object.assign(this.offices[this.editedIndex], this.editedItem);
                axios.create({withCredentials:true}).post("http://localhost:8000/api/rentACarCompany/updateLocation", office)
                    .then(res=>{
                        axios.create({withCredentials: true}).get("http://localhost:8000/api/rentACarCompany/getCompanyLocations")
                            .then(res => {
                                this.offices = res.data;
                            });
                        this.SuccessSnackbar = true;
                        this.SuccessSnackbarText = 'Office successfully edited';
                    })
                    .catch(err => {
                        this.ErrorSnackbar = true;
                        this.ErrorSnackbarText = 'Cannot edit office';
                    });
                this.close()
            }
        },
        watch:{
            dialog(val){
                val || this.close()
            }
        },
        computed:{
            officesToShow() {
                let offices = this.offices;
                if(this.filter.address.length > 0){
                    offices = offices.filter(office => office.Address.toLowerCase().includes(this.filter.address.toLowerCase()));
                }
                return offices;
            }
        }
    }
</script>

<style scoped>

</style>