<script setup>
import { ref } from 'vue';
import OrderItem from '../components/OrderItem.vue';
import { fetchFromApi } from '@/modules/fetch.mjs';

const orders = ref(null);
const showCompleted = ref(false);

const url =`${import.meta.env.VITE_API_URL}/order`;
async function fetchOrdersAsync() {
    orders.value = await (await fetchFromApi(url)).json();
    orders.value.sort((a,b) => a.ScheduledDate < b.ScheduledDate);
}

function handleOrderModified() {
    fetchOrdersAsync();
}

fetchOrdersAsync();
</script>

<template>
<div class="filters">
    <div class="filter">
        <span>Completed</span>
        <input type="checkbox" v-model="showCompleted" />
    </div>
</div>
<div>
    <OrderItem class="row" v-for="order in orders.filter(item => !item.Completed || showCompleted)" v-bind:key="order.Id"
        @order-modified="handleOrderModified()"
        :id="order.Id"
        :pastryId="order.Pastry.Id"
        :pastryName="order.Pastry.Name"
        :customer="order.Customer"
        :quantity="order.Quantity"
        :prefereddate="new Date(order.PreferedDate)"
        :scheduleddate="new Date(order.ScheduledDate)"
        :completed="order.Completed">
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