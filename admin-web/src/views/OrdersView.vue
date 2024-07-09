<script setup>
import { ref } from 'vue';
import OrderItem from '../components/OrderItem.vue';

const orders = ref(null);

const url =`${import.meta.env.VITE_API_URL}/order`;
async function fetchOrdersAsync() {
    const requestOptions = {
        method: 'GET',
        headers: { 'AuthToken': import.meta.env.VITE_API_AUTH_TOKEN }
    };
    orders.value = await (await fetch(url, requestOptions)).json();
}

function handleOrderModified() {
    fetchOrdersAsync();
}

fetchOrdersAsync();
</script>

<template>
<div>
    <OrderItem class="row" v-for="order in orders" v-bind:key="order.Id"
        @order-modified="handleOrderModified()"
        :id="order.Id"
        :pastry="order.Pastry"
        :customer="order.Customer"
        :quantity="order.Quantity"
        :prefereddate="new Date(order.PreferedDate)"
        :scheduleddate="new Date(order.ScheduledDate)">
    </OrderItem>
</div>
</template>

<style scoped>
.row {
    position: relative;
}

.row:hover {
    background-color: var(--hover-background-color);
}
</style>