<template>
    <v-container style="height: 100vh;">
        <v-card min-height="100%" class="flexcard" style="padding: 5px">
            <v-card-title primary-title>
                    <div class="headline font-weight-medium">Manage Vehicles:</div>
            </v-card-title>
            <v-card-text class="grow">
                <v-layout row-wrap>
                    <v-flex xs9 mr-5>
                        <h2>Filters:</h2>
                        <v-layout row-wrap>
                            <v-flex ma-2>
                                <v-text-field v-model="filter.Name" label="Name"></v-text-field>
                            </v-flex>
                            <v-flex ma-2>
                                <v-text-field v-model="filter.Capacity" label="Capacity"></v-text-field>
                            </v-flex>
                            <v-flex ma-2>
                                <v-text-field v-model="filter.Type" label="Type"></v-text-field>
                            </v-flex>
                            <v-flex ma-2>
                                <v-text-field v-model="filter.PricePerDay" label="Price Per Day"></v-text-field>
                            </v-flex>
                        </v-layout>
                        <h2>Vehicles:</h2>
                        <v-layout row-wrap>
                            <v-flex xs12>
                                <v-data-table :headers="headers" :items="vehiclesToShow" class="elevation-1">
                                    <template v-slot:items="props">
                                        <td>{{props.item.Name}}</td>
                                        <td>{{props.item.Capacity}}</td>
                                        <td>{{props.item.Type}}</td>
                                        <td>{{props.item.PricePerDay}}</td>
                                        <td>{{props.item.Discount | valueConversion}}</td>
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
                        <v-layout row-wrap>
                            <v-flex xs12>
                                <h2>Add Vehicle:</h2>
                                <v-form ref="form" class="align-center justify-center">
                                    <v-text-field
                                            v-model="addedItem.Name"
                                            label="Name"
                                            :rules="[rules.required]"
                                            required
                                    ></v-text-field>
                                    <v-text-field
                                            v-model="addedItem.Capacity"
                                            label="Capacity"
                                            :rules="[rules.required]"
                                            required
                                    ></v-text-field>
                                    <v-text-field
                                            v-model="addedItem.Type"
                                            label="Type"
                                            :rules="[rules.required]"
                                            required
                                    ></v-text-field>
                                    <v-text-field
                                            v-model="addedItem.PricePerDay"
                                            label="Price Per Day"
                                            :rules="[rules.required]"
                                            required
                                    ></v-text-field>
                                    <v-checkbox
                                            v-model="addedItem.Discount"
                                            label="Discount"
                                    ></v-checkbox>
                                    <v-btn color="primary" style="float: left" @click="addVehicle">submit</v-btn>
                                    <v-btn style="float: right" @click="clear">clear</v-btn>
                                </v-form>
                            </v-flex>
                        </v-layout>
                    </v-flex>
                </v-layout>
            </v-card-text>
        </v-card>
        <v-dialog v-model="dialog" persistent max-width="500px">
            <v-card>
                <v-card-title>
                    <span class="headline">Edit Vehicle Details:</span>
                </v-card-title>
                <v-card-text>
                    <v-container>
                        <v-layout row-wrap>
                            <v-flex xs6>
                                <v-text-field v-model="editedItem.Name" label="Name" :rules="[rules.required]"></v-text-field>
                            </v-flex>
                        </v-layout>
                        <v-layout row-wrap>
                            <v-flex xs6>
                                <v-text-field v-model="editedItem.Capacity" label="Capacity" :rules="[rules.required]"></v-text-field>
                            </v-flex>
                        </v-layout>
                        <v-layout row-wrap>
                            <v-flex xs6>
                                <v-text-field v-model="editedItem.Type" label="Type" :rules="[rules.required]"></v-text-field>
                            </v-flex>
                        </v-layout>
                        <v-layout row-wrap>
                            <v-flex xs6>
                                <v-text-field
                                        v-model="editedItem.PricePerDay"
                                        label="Price Per Day"
                                        :rules="[rules.required]"
                                        required
                                ></v-text-field>
                            </v-flex>
                        </v-layout>
                        <v-layout row-wrap>
                            <v-flex xs6>
                                <v-checkbox
                                        v-model="editedItem.Discount"
                                        label="Discount"
                                ></v-checkbox>
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
        name: "ManageVehicles",
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
                vehicles:[],
                filter: {
                    Name: '',
                    Capacity: '',
                    Type: '',
                    PricePerDay: '',
                },
                addedItem:{
                    Name: '',
                    Capacity: '',
                    Type: '',
                    PricePerDay: '',
                    Discount: false
                },
                editedItem:{
                    ID: '',
                    Name: '',
                    Capacity: '',
                    Type: '',
                    PricePerDay: '',
                    Discount: false
                },
                defaultItem:{
                    ID: '',
                    Name: '',
                    Capacity: '',
                    Type: '',
                    PricePerDay: '',
                    Discount: false
                },
                editedIndex: -1,
                headers:[
                    {
                        text: 'Name',
                        value: "Name",
                        width: '40%'
                    },
                    {
                        text: 'Capacity',
                        value: "Capacity",
                        width: '5%'
                    },
                    {
                        text: 'Type',
                        value: "Type",
                        width: '25%'
                    },
                    {
                        text: 'Price/Day',
                        value: "PricePerDay",
                        width: '5%'
                    },
                    {
                        text: 'Discount',
                        value: "Discount",
                        width: '5%'
                    },
                    {
                        text: 'Action',
                        sortable: false,
                        width: '20%'
                    }
                ]
            }
        },
        created(){
            axios.create({withCredentials: true}).get("http://localhost:8000/api/rentACarCompany/getCompanyVehicles")
                .then(res => {
                    this.vehicles = res.data;
                });
        },
        methods:{
            deleteItem(item){
                this.editedIndex = this.vehicles.indexOf(item);
                this.editedItem = Object.assign({}, item);
                let vehicle = {
                    ID: this.editedItem.ID,
                    Name: this.editedItem.Name,
                    Capacity: parseInt(this.editedItem.Capacity),
                    Type: this.editedItem.Type,
                    PricePerDay: parseFloat(this.editedItem.PricePerDay),
                    Discount: this.editedItem.Discount
                };
                Object.assign(this.vehicles[this.editedIndex], this.editedItem);
                axios.create({withCredentials:true}).post("http://localhost:8000/api/rentACarCompany/deleteVehicle", vehicle)
                    .then(res=>{
                        axios.create({withCredentials: true}).get("http://localhost:8000/api/rentACarCompany/getCompanyVehicles")
                            .then(res => {
                                this.vehicles = res.data;
                            });
                        this.SuccessSnackbar = true;
                        this.SuccessSnackbarText = 'Vehicle successfully deleted'
                    })
                    .catch(err => {
                        this.ErrorSnackbar = true;
                        this.ErrorSnackbarText = 'Vehicle has active reservations and cannot be deleted';
                    });
            },
            addVehicle(e){
                e.preventDefault();
                if (!this.checkAddInput()){
                    this.ErrorSnackbar = true;
                    this.ErrorSnackbarText = 'Cannot add vehicle';
                    return;
                }
                let vehicle = {
                    Name: this.addedItem.Name,
                    Capacity: parseInt(this.addedItem.Capacity),
                    Type: this.addedItem.Type,
                    PricePerDay: parseFloat(this.addedItem.Capacity),
                    Discount: this.addedItem.Discount
                };
                axios.create({withCredentials: true}).post('http://localhost:8000/api/rentACarCompany/addVehicle', vehicle)
                    .then(res =>{
                        axios.create({withCredentials: true}).get('http://localhost:8000/api/rentACarCompany/getCompanyVehicles')
                            .then(res => this.vehicles = res.data);
                        this.SuccessSnackbar = true;
                        this.SuccessSnackbarText = 'Vehicle successfully added';
                    })
                    .catch(err => {
                        this.ErrorSnackbar = true;
                        this.ErrorSnackbarText = 'Cannot add vehicle'
                    });

                this.clear();
            },
            clear() {
                this.$refs.form.reset();
            },
            checkAddInput(){
                if(!this.addedItem.Name || !this.addedItem.Capacity || !this.addedItem.Type ||
                    !this.addedItem.PricePerDay){
                    return false;
                }
                let ind = false;
                let number = parseInt(this.addedItem.Capacity);
                if(isNaN(number)){
                    this.addedItem.Capacity = "";
                    ind = true;
                }
                number = parseFloat(this.addedItem.PricePerDay);
                if(isNaN(number)){
                    this.addedItem.PricePerDay = "";
                    ind = true;
                }
                return ind !== true;
            },
            checkInput(){
                if(!this.editedItem.Name || !this.editedItem.Capacity || !this.editedItem.Type ||
                    !this.editedItem.PricePerDay){
                    return false;
                }
                let ind = false;
                let number = parseInt(this.editedItem.Capacity);
                if(isNaN(number) || number < 0){
                    this.editedItem.Capacity = "";
                    ind = true;
                }
                number = parseFloat(this.editedItem.PricePerDay);
                if(isNaN(number || number < 0)){
                    this.editedItem.PricePerDay = "";
                    ind = true;
                }
                return ind !== true;
            },
            editItem (item) {
                this.editedIndex = this.vehicles.indexOf(item);
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
                    this.ErrorSnackbarText = 'Cannot edit vehicle';
                    return;
                }
                let vehicle = {
                    ID: this.editedItem.ID,
                    Name: this.editedItem.Name,
                    Capacity: parseInt(this.editedItem.Capacity),
                    Type: this.editedItem.Type,
                    PricePerDay: parseFloat(this.editedItem.PricePerDay),
                    Discount: this.editedItem.Discount
                };
                axios.create({withCredentials:true}).post("http://localhost:8000/api/rentACarCompany/updateVehicle", vehicle)
                    .then(res=>{
                        axios.create({withCredentials: true}).get("http://localhost:8000/api/rentACarCompany/getCompanyVehicles")
                            .then(res => {
                                Object.assign(this.vehicles[this.editedIndex], this.editedItem);
                                this.vehicles = res.data;
                            });
                        this.SuccessSnackbar = true;
                        this.SuccessSnackbarText = 'Vehicle successfully edited'
                    })
                    .catch(err => {
                        this.ErrorSnackbar = true;
                        this.ErrorSnackbarText = 'Vehicle has active reservations and cannot be edited'
                    });
                this.close()
            }
        },
        watch:{
            dialog(val){
                val || this.close()
            }
        },
        filters:{
          valueConversion: function(value){
              if(value){
                  return "Yes";
              }else{
                  return "No";
              }
          }
        },
        computed:{
            vehiclesToShow() {
                let vehicles = this.vehicles;
                if(this.filter.Name.length > 0){
                    vehicles = vehicles.filter(vehicle => vehicle.Name.toLowerCase().includes(this.filter.Name.toLowerCase()));
                }
                if(this.filter.Capacity.length > 0){
                    vehicles = vehicles.filter(vehicle => vehicle.Capacity.toString().includes(this.filter.Capacity.toString()));
                }
                if(this.filter.Type.length > 0){
                    vehicles = vehicles.filter(vehicle => vehicle.Type.toLowerCase().includes(this.filter.Type.toLowerCase()));
                }
                if(this.filter.PricePerDay.length > 0){
                    vehicles = vehicles.filter(vehicle => vehicle.PricePerDay.toString().includes(this.filter.PricePerDay.toString()));
                }
                return vehicles;
            }
        }
    }
</script>

<style scoped>

</style>