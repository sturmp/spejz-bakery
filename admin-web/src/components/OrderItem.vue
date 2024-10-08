<script setup>
import { ref } from 'vue';
import { formatDate } from '@/modules/datetime.mjs';
import { fetchFromApi } from '@/modules/fetch.mjs';

const props = defineProps({
    id: Number,
    pastryId: Number,
    pastryName: String,
    customer: String,
    quantity: Number,
    prefereddate: Date,
    scheduleddate: Date,
    completed: Boolean,
})

const emits = defineEmits(['order-modified']);

const schedules = ref(null);
const showSchedules = ref(false);

function isOrderScheduled(date) {
    return date.getFullYear() != 1;
}

const completeUrl =`${import.meta.env.VITE_API_URL}/order/complete/`;
async function completeOrderAsync(orderId) {
    const requestOptions = {
        method: 'PUT',
    };
    await fetchFromApi(completeUrl + orderId, requestOptions)
        .then(() => emits('order-modified'));
}

const deleteUrl =`${import.meta.env.VITE_API_URL}/order/`;
async function deleteOrderAsync(orderId) {
    const requestOptions = {
        method: 'DELETE',
    };
    await fetchFromApi(deleteUrl + orderId, requestOptions)
        .then(() => emits('order-modified'));
}

const getScheduleUrl = `${import.meta.env.VITE_API_URL}/schedule`;
async function fetchSchedulesAsync() {
    schedules.value = await (await fetchFromApi(getScheduleUrl)).json();
    schedules.value = schedules.value.filter(schedule => {
        var scheduleReadyDate = new Date(schedule.ReadyDate)
        return scheduleReadyDate >= new Date()
            && schedule.Pastry.Id == props.pastryId
            && schedule.Quantity != schedule.Reserved
    })
}

function openSelectScheduleDialog() {
    fetchSchedulesAsync();
    showSchedules.value = true;
}

const scheduleOrderUrl =`${import.meta.env.VITE_API_URL}/order/schedule`;
function scheduleOrder(orderId, orderScheduledDate) {
    const orderSchedule = {
        Id: orderId,
        ScheduledDate: orderScheduledDate
    };

    const requestOptions = {
        method: 'POST',
        body: JSON.stringify(orderSchedule)
    };
    fetchFromApi(scheduleOrderUrl, requestOptions)
        .then(() => {
            emits('order-modified');
            showSchedules.value = false;
        });
}

function cancelOrderSchedule() {
    showSchedules.value = false;
}
</script>

<template>
    <div class="order" :class="{completed: completed}">
        <div class="order-property">{{ pastryName }}</div>
        <div class="order-property">{{ customer }}</div>
        <div class="order-property">{{ quantity }}</div>
        <div class="order-property">{{ formatDate(prefereddate) }}</div>
        <div class="order-property" v-if="isOrderScheduled(scheduleddate)">{{ formatDate(scheduleddate) }}</div>
        <div class="order-property controll" v-else @click="openSelectScheduleDialog()">Schedule</div>
        <div class="order-property controll" @click="completeOrderAsync(id)">✔</div>
        <div class="order-property controll" @click="deleteOrderAsync(id)">X</div>
    </div>
    <div id="obscure" v-if="showSchedules"></div>
    <div class="schedules" v-if="showSchedules">
        <div class="schedule" v-for="schedule, index in schedules"
            v-bind:key="index"
            @click="scheduleOrder(id, schedule.ReadyDate)">
            <div>{{ formatDate(new Date(schedule.ReadyDate)) }}</div>
            <div>{{ "[" + schedule.Reserved + "/" + schedule.Quantity +"]" }}</div>
        </div>
        <div id="cancel" @click="cancelOrderSchedule()">X</div>
    </div>
</template>

<style scoped>
.order {
    display:grid;
    grid-template-columns: 4fr 4fr 4fr 4fr 4fr 1fr 1fr;
    border-bottom: var(--border-size) dotted var(--color-text);
    border-left: var(--border-size) dotted var(--color-text);
    border-right: var(--border-size) dotted var(--color-text);
}

.completed {
    color: var(--color-text-disabled);
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

#obscure {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  backdrop-filter: blur(0.1em);
  z-index: 2;
}

.schedules {
    z-index: 3;
    width: 300px;
    position: absolute;
    top: 50%;
    left: 50%;
    -webkit-transform: translate(-50%, -50%);
    transform: translate(-50%, -50%);

    display: flex;
    flex-direction: column;

    background-color: var(--color-background);

    box-shadow: 0em 0.5em 1em 0 rgba(0, 0, 0, 0.35), 0em 0.5em 3em 0 rgba(0, 0, 0, 0.3);
}

.schedule {
    display: grid;
    grid-template-columns: 3fr 1fr;
    padding: 0.5em;
    cursor: pointer;
}

.schedule:hover {
    background-color: var(--hover-background-color);
}

.schedule div {
    padding: 0 0.5em;
}

#cancel {
    position: absolute;
    top: -0.75em;
    right: -0.75em;

    margin-bottom: 0.5em;
    
    font-size: 1.5rem;
    font-weight: bold;
    display: flex;
    justify-content: center;
    align-items: center;
    text-align: center;
    
    height: 1.5em;
    width: 1.5em;
    background-color: #ffffff;
    border: var(--controll-border-size) dotted var(--border-color);
    border-radius: 50%;

    cursor: pointer;
}

#cancel:hover {
    border: var(--controll-border-size) dotted var(--border-color-hover);
}
</style>