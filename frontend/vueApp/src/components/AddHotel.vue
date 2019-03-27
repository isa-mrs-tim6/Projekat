<template>
    <form @submit="submit">
        <v-text-field
                v-model="Name"
                label="Name"
                required
                @input="$v.Name.$touch()"
                @blur="$v.Name.$touch()"
        ></v-text-field>
        <v-text-field
                v-model="Address"
                label="Address"
                required
                @input="$v.Address.$touch()"
                @blur="$v.Address.$touch()"
        ></v-text-field>
        <v-text-field
                v-model="Latitude"
                label="Latitude"
                required
                @input="$v.Latitude.$touch()"
                @blur="$v.Latitude.$touch()"
        ></v-text-field>
        <v-text-field
                v-model="Longitude"
                label="Longitude"
                required
                @input="$v.Longitude.$touch()"
                @blur="$v.Longitude.$touch()"
        ></v-text-field>
        <v-text-field
                v-model="Description"
                label="Description"
                required
                @input="$v.Description.$touch()"
                @blur="$v.Description.$touch()"
        ></v-text-field>

        <v-btn @click="submit">submit</v-btn>
        <v-btn @click="clear">clear</v-btn>
    </form>
</template>

<script>
    export default {
        name: "AddHotel",
        data() {
            return {
                Name: '',
                Address: '',
                Longitude: '',
                Latitude: '',
                Promo: '',
            }
        },
        methods: {
            submit (e) {
                e.preventDefault();
                const newHotel = {
                    Name: this.Name,
                    Address: this.Address,
                    Longitude: this.Longitude,
                    Latitude: this.Latitude,
                    Promo: this.Promo,
                };
                this.clear();
                this.$emit('add-hotel', newHotel);
            },
            clear () {
                this.Name = '';
                this.Address = '';
                this.Longitude = '';
                this.Latitude = '';
                this.Promo = '';
            },
        }
    }
</script>

<style scoped>

</style>