<script setup>
import { ref } from "vue";
import { formatDate } from '@/modules/datetime.mjs';
import { fetchFromApi } from '@/modules/fetch.mjs';

const props = defineProps({
    pastryId: Number,
    pastryName: String,
    quantity: Number,
    reserved: Number,
    readydate: Date,
});

const emit = defineEmits(["schedule-submit", "schedule-cancel"]);

const schedule = ref({
    pastryId: props.pastryId,
    quantity: props.quantity,
    reserved: props.reserved,
    readydate: props.readydate,
});

const url =`${import.meta.env.VITE_API_URL}/schedule`;
async function updateScheduleAsync() {
    const requestOptions = {
        method: 'PUT',
        body: JSON.stringify(schedule.value)
    };
    await (await fetchFromApi(url, requestOptions)).json();
}

function handleSubmit() {
    updateScheduleAsync(schedule)
        .then(() => emit('schedule-submit'));
}
</script>

<template>
    <div class="schedule-edit">
        <div class="schedule-property">{{props.pastryName}}</div>
        <div class="schedule-property"><input v-model="schedule.quantity" type="number"/></div>
        <div class="schedule-property"><input v-model="schedule.reserved" type="number"/></div>
        <div class="schedule-property">{{formatDate(schedule.readydate)}}</div>
        <div class="controlls-group">
            <div class="controll" @click="handleSubmit()">✓</div>
            <div class="controll" @click="emit('schedule-cancel')">x</div>
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