<script setup>
import { ref } from 'vue';
import PastryItem from "../components/PastryItem.vue"

const pastries = ref(null);

const url =`${import.meta.env.VITE_API_URL}/pastry`;
async function fetchPastriesAsync() {
    const requestOptions = {
        method: 'GET',
        headers: { 'AuthToken': import.meta.env.VITE_API_AUTH_TOKEN }
    };
    pastries.value = await (await fetch(url, requestOptions)).json();
}

fetchPastriesAsync();

</script>

<template>
    <div cass="content">
        <PastryItem class="row" v-for="(pastry, index) in pastries" v-bind:key=index
            :name="pastry.Name"
            :description="pastry.Description"
            :price="pastry.Price"
            :unitOfMeasure="pastry.UnitOfMeasure"
            :quantityPerPiece="pastry.QuantityPerPiece"/>
    </div>
</template>

<style scoped>
</style>