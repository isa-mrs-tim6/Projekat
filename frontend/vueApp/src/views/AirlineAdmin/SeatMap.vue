<template>
    <div class = "airplane">
        <template v-for="(seat, index) in this.editedSeats.Seats">
            <v-layout row-wrap v-if="index % editedSeats.RowWidth  === 0">
            </v-layout>
            <div class="passage" v-if="index % editedSeats.RowWidth === editedSeats.RowWidth / 2"></div>
            <div v-if="editedSeats.Seats[index].ReservationID !== 0" class=seat :class="seatClassBinding(index)" :key="index">
                <span v-if="editedSeats.Seats[index].ReservationID !== 0" class="occupied"></span>
            </div>
            <div class="seat" :key="index" @click="chooseSeat(index)" :class="seatClassBinding(index)" v-else>
            </div>
        </template>
    </div>
</template>

<script>
    import AirlineAdminNavbar from "../../components/AirlineAdmin/AirlineAdminNavbar";
    export default {
        name: "SeatMap",
        components: {AirlineAdminNavbar},
        props:['editedSeats'],
        data(){
            return{
                Flights:[]
            }
        },
        methods:{
            seatClassBinding(index){
                let seatClass = this.editedSeats.Seats[index].Class;
                let disabled = this.editedSeats.Seats[index].Disabled;
                let retval = "";
                if(disabled){
                    retval += "disabled";
                }else{
                    switch(seatClass){
                        case "BUSINESS":
                            retval = "business";
                            break;
                        case "ECONOMIC":
                            retval = "economic";
                            break;
                        case "FIRST":
                            retval = "first";
                            break;
                        default:
                            retval = "";
                    }
                }
                if (index === this.editedSeats.SelectedSeat){
                    retval += " selected"
                }
                return retval;
            },
            chooseSeat(index){
                this.editedSeats.Number = this.editedSeats.Seats[index].Number;
                this.editedSeats.Class = this.editedSeats.Seats[index].Class;
                this.editedSeats.SelectedSeat = index;
            },
        }
    }
</script>

<style scoped>

</style>