<script setup>
defineProps({
    id: Number,
    pastry: String,
    customer: String,
    quantity: Number,
    prefereddate: Date,
    scheduleddate: Date,
})

const emits = defineEmits(['order-modified']);

function formatDate(date) {
    const month = date.getMonth();
    const monthString = month < 10 ? `0${month}` : month;
    const day = date.getDate();
    const dayString = day < 10 ? `0${day}` : day;

    return `${date.getFullYear()}-${monthString}-${dayString} ${date.getHours()<=12? "Morning":"Afternoon"}`;
}

function isDefaultDate(date) {
    return date.getFullYear() == 1;
}

const scheduleUrl =`${import.meta.env.VITE_API_URL}/order/schedule`;
async function scheduleOrderAsync(orderId, orderScheduledDate) {
    const orderSchedule = {
        Id: orderId,
        ScheduledDate: orderScheduledDate
    };

    const requestOptions = {
        method: 'POST',
        headers: { 'AuthToken': import.meta.env.VITE_API_AUTH_TOKEN },
        body: JSON.stringify(orderSchedule)
    };
    await fetch(scheduleUrl, requestOptions);
    emits('order-modified');
}

const deleteUrl =`${import.meta.env.VITE_API_URL}/order/`;
async function deleteOrderAsync(orderId) {
    const requestOptions = {
        method: 'DELETE',
        headers: { 'AuthToken': import.meta.env.VITE_API_AUTH_TOKEN },
    };
    await fetch(deleteUrl + orderId, requestOptions);
    emits('order-modified');
}
</script>

<template>
    <div class="order">
        <div class="order-property">{{ pastry }}</div>
        <div class="order-property">{{ customer }}</div>
        <div class="order-property">{{ quantity }}</div>
        <div class="order-property">{{ formatDate(prefereddate) }}</div>
        <div class="order-property" v-if="isDefaultDate(scheduleddate)">{{ formatDate(scheduleddate) }}</div>
        <div class="order-property controll" v-else @click="scheduleOrderAsync(id, scheduleddate)">Schedule</div>
        <div class="order-property controll" @click="deleteOrderAsync(id)">X</div>
    </div>
</template>

<style scoped>
.order {
    display:grid;
    grid-template-columns: 4fr 4fr 4fr 4fr 4fr 1fr;
    border-bottom: var(--border-size) dotted var(--color-text);
    border-left: var(--border-size) dotted var(--color-text);
    border-right: var(--border-size) dotted var(--color-text);
}

.order:first-of-type {
    border-top: var(--border-size) dotted var(--color-text);
}

.order-property {
    display: flex;
    padding: 0.5em 0.5em;
    border-right: var(--border-size) dotted var(--color-text);

    align-items: center;
}

.order-property:last-child {
    border-right: 0;
}

.controll {
    display: flex;
    cursor: pointer;
    justify-content : center;
    font-weight: bold;
}

.controll:hover {
    color: var(--color-text-highlight)
}
</style>