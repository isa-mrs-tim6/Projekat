<template>
    <div class = "airplane">
        <template style="display:inline-flex; flex-wrap: nowrap" v-for="(seat, index) in this.editedSeats.Seats">
            <v-layout row v-if="index % editedSeats.RowWidth  === 0">
            </v-layout>
            <div class="passage" v-if="index % editedSeats.RowWidth === Math.floor(editedSeats.RowWidth / 2)"></div>
            <div v-if="editedSeats.Seats[index].ReservationID !== 0" class=seat :class="seatClassBinding(index)" :key="index">
                <span v-if="editedSeats.Seats[index].ReservationID !== 0" class="occupied"></span>
            </div>
            <div class="seat" :key="index" @click="chooseSeat(index)" :class="seatClassBinding(index)" v-else>
            </div>
        </template>
    </div>
</template>

<script>
    export default {
        name: "SeatMap",
        components: {},
        props:['editedSeats','tab'],
        data(){
            return{
                Flights:[],
                SelectedSeats:[]
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
                if(this.editedSeats.Seats[index].ReservationID !== 0){
                    retval += " disabled"
                }
                if(this.tab === "reservation"){
                    if (this.SelectedSeats.includes(this.editedSeats.Seats[index])){
                        retval += "selected";
                    }
                }else{
                    if (index === this.editedSeats.SelectedSeat){
                        retval += " selected";
                    }
                }

                return retval;
            },
            chooseSeat(index){
                if(this.tab === "addSeat"){
                    return;
                }else if(this.tab === "reservation"){
                    if(this.editedSeats.Seats[index].Disabled){
                        return;
                    }
                    if(this.SelectedSeats.includes(this.editedSeats.Seats[index])){
                        this.removeA(this.SelectedSeats, this.editedSeats.Seats[index])
                    }else{
                        this.SelectedSeats.push(this.editedSeats.Seats[index])
                    }
                    this.$emit('seatsChanged', this.SelectedSeats)
                }else{
                    this.editedSeats.Number = this.editedSeats.Seats[index].Number;
                    this.editedSeats.Class = this.editedSeats.Seats[index].Class;
                    this.editedSeats.SelectedSeat = index;
                }
            },
            removeA(arr) {
                var what, a = arguments, L = a.length, ax;
                while (L > 1 && arr.length) {
                    what = a[--L];
                    while ((ax= arr.indexOf(what)) !== -1) {
                        arr.splice(ax, 1);
                    }
                }
                return arr;
            },
        },
    }
</script>

<style scoped>

</style>