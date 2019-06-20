<template>
  <v-card>
    <v-container>
      <v-card-title><span class = "headline">Edit quick reservations</span></v-card-title>
    </v-container>
    <v-card-text>
      <v-container>
        <v-layout row>
        <v-flex xs7>
          <v-layout row>
            <v-flex xs12><seat-map :editedSeats="this.seats" tab='quickReservation' @seatSelected="seatSelectedEvent"></seat-map></v-flex>
          </v-layout>
          <v-layout row py-2>
            <v-flex xs3><span class = "body-2">Seat class: {{selectedSeat.Class}}</span></v-flex>
            <v-flex xs3><span class = "body-2">Seat price:{{price}}</span></v-flex>
            <v-flex xs3><span class = "body-2">Seat number:{{selectedSeat.Number}}</span></v-flex>
          </v-layout>
          <v-layout row py-3>
            <v-flex xs3><v-text-field v-model="discount" label="Discount[%]" type="number" min="1" max = "100" ></v-text-field></v-flex>
            <v-flex xs3><v-btn @click="createReservation">Create</v-btn></v-flex>
          </v-layout>
        </v-flex>
          <v-flex xs5 justify-end>
            <v-data-table :headers="headers" :items=this.results class="elevation-1">
              <template v-slot:items="props">
                <td>{{ props.item.Number }}</td>
                <td>{{ (props.item.Price).toFixed(2) }}</td>
                <td><v-icon v-if="props.item.ID == 0" @click = "remove(props.item.ReservationID, props.item.Number)">remove</v-icon></td>
              </template>
            </v-data-table> 
          </v-flex>
      </v-layout>
      </v-container>
    </v-card-text>
    <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn @click="closeDialog">Close</v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
import axios from "axios";
import SeatMap from "../../components/AirlineAdmin/SeatMap";
export default {
  name: "FlightQuickReservation",
  components: { SeatMap },
  props: ["flight","seats"],
  data() {
    return {
      results: [],
      discount: 1,
      selectedIndex: -1,
      selectedSeat:{
        Class: "",
        Number: ""
      },
      headers:[
        {text:"Seat Number", value:"number"},
        {text:"Price", value: "price"},
        {text:"Action", value: "action"},
      ]
    };
  },
  methods: {
    getReservation(idx) {
      axios
        .create({ withCredentials: true })
        .get(
          "http://ec2-18-195-170-20.eu-central-1.compute.amazonaws.com:8000/api/flight/" + idx + "/quickReservation"
        )
        .then(res => {
          this.results = res.data;
        });
    },
    remove(id, seatNumber){
      var quickReservationDto = {
        ReservationID: id,
        Number: seatNumber,
      }
      axios.create({withCredentials: true}).delete("http://ec2-18-195-170-20.eu-central-1.compute.amazonaws.com:8000/api/flight/quickReservation",{data: quickReservationDto})
                    .then(res => {
                        for (let o of this.seats.Seats) {
                          if (o.Number === seatNumber){
                            if (o.ReservationID === 0){
                              o.Disabled = false;
                            }else{
                              o.ReservationID = 0;
                            }
                            break;
                          }
                        }
                        this.getReservation(this.flight.ID)
                    })
    },
    closeDialog() {
      this.$emit("closeReservationDialog");
    },
    seatSelectedEvent(seat, index){
      this.selectedSeat = seat;
      this.selectedIndex = index;
    },
    createReservation(){
      if(this.discount === "" && this.discount < 1 || this.discount > 100 || !Number.isInteger(parseInt(this.discount))){
        alert("Bad discount");
        return;
      }
      if(this.selectedSeat.Number === ""){
        alert("Seat not choosen");
        return;
      }
      var quickReservationDto = {
        SeatId: this.selectedSeat.ID,
        Discount: parseInt(this.discount),
        FlightId: this.flight.ID
      };
      axios.create({withCredentials: true})
        .post("http://ec2-18-195-170-20.eu-central-1.compute.amazonaws.com:8000/api/flight/quickReservation",quickReservationDto)
        .then(res=>{
          this.seats.Seats[this.selectedIndex].Disabled = true;
          this.seats.SelectedSeat = -1;
          this.selectedSeat = {
            Class: "",
            Number: "",
          };
          this.getReservation(this.flight.ID);
        })
    }
  },
  computed:{
    price: function(){
      var retVal;
      if(this.selectedSeat.Class === ""){
        return "";
      }
      if(this.selectedSeat.Class === "ECONOMIC"){
        retVal = this.flight.PriceECONOMY;
      }else if(this.selectedSeat.Class === "BUSINESS"){
        retVal = this.flight.PriceBUSINESS;
      }else{
        retVal = this.flight.PriceFIRSTCLASS;
      }
      return retVal;
    }
  }
};
</script>

<style>
</style>
