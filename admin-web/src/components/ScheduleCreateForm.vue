<script setup>
import { ref } from "vue";

const emit = defineEmits(["schedule-create"]);

const schedule = ref({
    pastry: "",
    quantity: 1,
    reserved: 1,
    readydate: "",
    partOfDay: "",
});

const pastries = ref(null);

const urlFetchPastries = `${import.meta.env.VITE_API_URL}/pastry`;
async function fetchPastriesAsync() {
    const requestOptions = {
        method: 'GET',
        headers: { 'AuthToken': import.meta.env.VITE_API_AUTH_TOKEN }
    };
    pastries.value = await (await fetch(urlFetchPastries, requestOptions)).json();
    schedule.value.pastry = pastries.value[0].Name;
}

const urlCreateSchedule =`${import.meta.env.VITE_API_URL}/schedule`;
async function createScheduleAsync() {
    const date = new Date(schedule.value.readydate);

    const newSchedule = {
        pastry: schedule.value.pastry,
        quantity: schedule.value.quantity,
        reserved: schedule.value.reserved,
        readydate: schedule.value.partOfDay == "morning"
            ? new Date(date.setHours(8))
            : new Date(date.setHours(16)),
    };

    const requestOptions = {
        method: 'POST',
        headers: { 'AuthToken': import.meta.env.VITE_API_AUTH_TOKEN },
        body: JSON.stringify(newSchedule)
    };
    await (await fetch(urlCreateSchedule, requestOptions)).json();
}

function handleCreate() {
    createScheduleAsync(schedule);
    emit('schedule-create');
}

fetchPastriesAsync();
</script>

<template>
    <div class="schedule-create">
        <div class="schedule-property">
            <select v-model="schedule.pastry">
                <option v-for="pastry in pastries" :key="pastry.Name">{{ pastry.Name }}</option>
            </select>
        </div>
        <div class="schedule-property"><input v-model="schedule.quantity" type="number"/></div>
        <div class="schedule-property"><input v-model="schedule.reserved" type="number"/></div>
        <div class="schedule-property"><input v-model="schedule.readydate" type="date"/></div>
        <div class="schedule-property">
            <select v-model="schedule.partOfDay">
                <option value="morning">Délelőtt</option>
                <option value="afternoon">Délután</option>
            </select>
        </div>
        <div id="create" @click="handleCreate()">Create</div>
    </div>

</template>

<style scoped>
.schedule-create {
    display:grid;
    grid-template-columns: 1fr 1fr 1fr 1fr 1fr 1fr;
    border-bottom: var(--border-size) dotted var(--color-text);
    border-left: var(--border-size) dotted var(--color-text);
    border-right: var(--border-size) dotted var(--color-text);
    border-top: var(--border-size) dotted var(--color-text);
    margin-bottom: 1em;
}

.schedule-property {
    display: flex;
    border-right: var(--border-size) dotted var(--color-text);
    align-items: center;
}

.schedule-property:first-of-type {
    padding: 0.5em;
}

.schedule-property:nth-child(6) {
    border-right: 0;
}

input {
    width: 0;
    font-family: inherit;
    font-size: inherit;
    font-weight: inherit;
    border: 0;
    padding: 0.5em;
    flex: 1 1 auto;
}

#create {
    display: flex;
    justify-content: center;
    align-items: center;
    font-weight: bold;
    cursor: pointer;
}

#create:hover {
    color: var(--color-text-highlight);
}
</style>