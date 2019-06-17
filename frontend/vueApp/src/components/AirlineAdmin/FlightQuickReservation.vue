<template>
  <v-card>
    <v-card-text>
      <v-container row>
        <v-flex xs8>
          <v-container row>
            <v-flex xs12>
              <seat-map :editedSeats="this.flightSeats" tab='reservation'></seat-map>
            </v-flex>
          </v-container>
          <v-container row py-0>
            <v-btn>Create</v-btn>
          </v-container>
        </v-flex>`
        <v-flex xs4></v-flex>
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
  props: ["flightID", "flightSeats"],
  data() {
    return {
      results: []
    };
  },
  methods: {
    getReservation() {
      axios
        .create({ withCredentials: true })
        .get(
          "http://localhost:8000/api/flight/" + this.flightID + "/quickReservation"
        )
        .then(res => {
          results = res.data;
        });
    },
    closeDialog() {
      this.$emit("closeReservationDialog");
    },
    add() {
      console.log("dodaj");
    }
  }
};
</script>

<style>
</style>
