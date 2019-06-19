<template>
    <div :class = "airplaneClassBinding()">
        <template v-for="(seat, index) in this.editedSeats.Seats">
            <v-layout row v-if="index % editedSeats.RowWidth  === 0">
            </v-layout>
            <div :class="passageClassBinding()" v-if="index % editedSeats.RowWidth === Math.floor(editedSeats.RowWidth / 2)"></div>
            <v-tooltip bottom :key=index>
                <template v-slot:activator="{ on }">
                    <div v-on="on" v-if="editedSeats.Seats[index].ReservationID !== 0" :class="seatClassBinding(index)" :key="index">
                        <span v-if="editedSeats.Seats[index].ReservationID !== 0" :class="occupiedClassBinding()"></span>
                    </div>
                    <div v-on="on" :key="index" @click="chooseSeat(index)" :class="seatClassBinding(index)" v-else>
                    </div>
                </template>
                <span>{{editedSeats.Seats[index].Number}}</span>
            </v-tooltip>
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
                SelectedSeats:[],
            }
        },
        methods:{
            airplaneClassBinding(){
                var retVal = "";
                if (this.tab === "reservation" || this.tab === "quickReservation"){
                    retVal += "airplane2";
                }else{
                    retVal += "airplane";
                }
                return retVal;
            },
            passageClassBinding(){
                var retVal = "";
                if (this.tab === "reservation" || this.tab === "quickReservation"){
                    retVal += "passage2";
                }else{
                    retVal += "passage";
                }
                return retVal;
            },
            occupiedClassBinding(){
                var retVal = "";
                if (this.tab === "reservation" || this.tab === "quickReservation"){
                    retVal +=" occupied2";
                }else{
                    retVal += " occupied";
                }
                return retVal;
            },
            seatClassBinding(index){
                let seatClass = this.editedSeats.Seats[index].Class;
                let disabled = this.editedSeats.Seats[index].Disabled;
                let retVal = "";
                if (this.tab === "reservation" || this.tab === "quickReservation"){
                    retVal += "seat2";
                    if(disabled){
                    retVal += " disabled2";
                    }else{
                        switch(seatClass){
                            case "BUSINESS":
                                retVal += " business2";
                                break;
                            case "ECONOMIC":
                                retVal += " economic2";
                                break;
                            case "FIRST":
                                retVal += " first2";
                                break;
                            default:
                                retVal = "";
                        }
                    }
                    if(this.editedSeats.Seats[index].ReservationID !== 0){
                        retVal += " disabled2"
                    }
                    if (this.tab === "reservation"){   
                        if (this.SelectedSeats.includes(this.editedSeats.Seats[index])){
                            retVal += "selected2";
                        }
                    }else{
                        if (index === this.editedSeats.SelectedSeat){
                        retVal += "selected2";
                    }
                    }
                }else{
                    retVal += "seat";
                    if(disabled){
                        retVal += " disabled";
                    }else{
                        switch(seatClass){
                            case "BUSINESS":
                                retVal += " business";
                                break;
                            case "ECONOMIC":
                                retVal += " economic";
                                break;
                            case "FIRST":
                                retVal += " first";
                                break;
                            default:
                                retVal = "";
                        }
                    }
                    if(this.editedSeats.Seats[index].ReservationID !== 0){
                        retVal += " disabled"
                    }
                    if (index === this.editedSeats.SelectedSeat){
                        retVal += " selected";
                    }
                }
                return retVal;
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
                    if(this.tab === "quickReservation"){
                        if(this.editedSeats.Seats[index].Disabled){
                        return;
                        }
                    }
                    this.editedSeats.Number = this.editedSeats.Seats[index].Number;
                    this.editedSeats.Class = this.editedSeats.Seats[index].Class;
                    this.editedSeats.SelectedSeat = index;
                    this.$emit('seatSelected', this.editedSeats.Seats[index], index);
                }
            },
            resetSelect(){
                console.log("opp");
                this.editedSeats.SelectedSeat = -1;
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
    @import '../../assets/css/SeatMap.css'
</style>