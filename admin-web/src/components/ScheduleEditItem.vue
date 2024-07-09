<script setup>
import { ref } from "vue";

const props = defineProps({
    pastry: String,
    quantity: Number,
    reserved: Number,
    readydate: Date,
});

const emit = defineEmits(["submit", "cancel"]);

const schedule = ref({
    pastry: props.pastry,
    quantity: props.quantity,
    reserved: props.reserved,
    readydate: props.readydate,
});

const url =`${import.meta.env.VITE_API_URL}/schedule`;
async function updateScheduleAsync() {
    const requestOptions = {
        method: 'PUT',
        headers: { 'AuthToken': import.meta.env.VITE_API_AUTH_TOKEN },
        body: JSON.stringify(schedule.value)
    };
    await (await fetch(url, requestOptions)).json();
}

function handleSubmit() {
    updateScheduleAsync(schedule);
    emit('submit');
}

function formatDate(date) {
    const month = date.getMonth();
    const monthString = month < 10 ? `0${month}` : month;
    const day = date.getDate();
    const dayString = day < 10 ? `0${day}` : day;

    return `${date.getFullYear()}-${monthString}-${dayString} ${date.getHours()<=12? "Morning":"Afternoon"}`;
}
</script>

<template>
    <div class="schedule-edit">
        <div class="schedule-property">{{schedule.pastry}}</div>
        <div class="schedule-property"><input v-model="schedule.quantity" type="number"/></div>
        <div class="schedule-property"><input v-model="schedule.reserved" type="number"/></div>
        <div class="schedule-property">{{formatDate(schedule.readydate)}}</div>
        <div class="controlls-group">
            <div class="controll" @click="handleSubmit()">✓</div>
            <div class="controll" @click="emit('cancel')">x</div>
        </div>
    </div>

</template>

<style scoped>
.schedule-edit {
    background-color: var(--edit-background-color);
    display:grid;
    grid-template-columns: 1fr 1fr 1fr 1fr;
    border-bottom: var(--border-size) dotted var(--color-text);
    border-left: var(--border-size) dotted var(--color-text);
    border-right: var(--border-size) dotted var(--color-text);
}

.schedule-edit:first-of-type {
    border-top: var(--border-size) dotted var(--color-text);
}

.schedule-property {
    color: var(--color-edit-text);
    display: flex;
    border-right: var(--border-size) dotted var(--color-text);
    
    align-items: center;
}

.schedule-property:first-of-type {
    padding: 0.5em;
}

.schedule-property:nth-child(5) {
    border-right: 0;
}

.controlls-group {
    z-index: 10;
    position: absolute;
    top: -1.25em;
    right: -1.25em;
}

.controll {
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

.controll:hover {
    border: var(--controll-border-size) dotted var(--border-color-hover);
}

input {
    color: var(--color-edit-text);
    background-color: var(--edit-background-color);
    width: 0;
    font-family: inherit;
    font-size: inherit;
    font-weight: inherit;
    border: 0;
    padding: 0.5em;
    flex: 1 1 auto;
}
</style>