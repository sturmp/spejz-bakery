<script setup>
import { ref } from "vue";
import { fetchFromApi } from '@/modules/fetch.mjs';

const emit = defineEmits(["schedule-create"]);

const schedule = ref({
    pastryId: 0,
    pastryName: "",
    quantity: 1,
    reserved: 1,
    readydate: "",
    partOfDay: "",
});

const pastries = ref(null);

const urlFetchPastries = `${import.meta.env.VITE_API_URL}/pastry`;
async function fetchPastriesAsync() {
    pastries.value = await (await fetchFromApi(urlFetchPastries)).json();
    schedule.value.pastryId = pastries.value[0].Id;
    schedule.value.pastryName = pastries.value[0].Name;
}

const urlCreateSchedule =`${import.meta.env.VITE_API_URL}/schedule`;
async function createScheduleAsync() {
    const date = new Date(schedule.value.readydate);

    const newSchedule = {
        pastryid: schedule.value.pastryId,
        quantity: schedule.value.quantity,
        reserved: schedule.value.reserved,
        readydate: schedule.value.partOfDay == "morning"
            ? new Date(date.setHours(8))
            : new Date(date.setHours(16)),
    };

    const requestOptions = {
        method: 'POST',
        body: JSON.stringify(newSchedule)
    };
    await (await fetchFromApi(urlCreateSchedule, requestOptions)).json();
}

function handleCreate() {
    createScheduleAsync(schedule)
        .then(() => emit('schedule-create'));
}

fetchPastriesAsync();
</script>

<template>
    <div class="schedule-create">
        <div class="schedule-property">
            <select v-model="schedule.pastryId">
                <option v-for="pastry in pastries" :key="pastry.Id" :value="pastry.Id">{{ pastry.Name }}</option>
            </select>
        </div>
        <div class="schedule-property"><span>Quantity:</span><input v-model="schedule.quantity" type="number"/></div>
        <div class="schedule-property"><span>Reserved:</span><input v-model="schedule.reserved" type="number"/></div>
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

.schedule-property span {
    padding-left: 0.5em;
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