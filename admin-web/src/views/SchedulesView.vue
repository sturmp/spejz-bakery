<script setup>
import { ref } from 'vue';
import ScheduleItem from "../components/ScheduleItem.vue"
import ScheduleEditItem from "../components/ScheduleEditItem.vue"
import ScheduleCreateForm from "../components/ScheduleCreateForm.vue"
import { fetchFromApi } from '@/modules/fetch.mjs';

const schedules = ref(null);
const editedSchedule = ref(null)
const showOld = ref(false);

const startOfWeek = calculateStartOfWeek();

const url =`${import.meta.env.VITE_API_URL}/schedule`;
async function fetchSchedulesAsync() {
    schedules.value = await (await fetchFromApi(url)).json();
    schedules.value.sort((a, b) => a.ReadyDate < b.ReadyDate);
}

function editSchedule(schedule) {
    editedSchedule.value = schedule;
}

function handleCreate() {
    fetchSchedulesAsync();
}

function handleSubmit() {
    fetchSchedulesAsync()
        .then(() => editedSchedule.value = null);
}

function handleCancel() {
    editedSchedule.value = null;
}

fetchSchedulesAsync();

function calculateStartOfWeek() {
    var today = new Date();
    today.getDay();
    return today.setDate(today.getDate()-today.getDay());
}
</script>

<template>
    <ScheduleCreateForm @schedule-create="handleCreate()"/>
    <div class="filters">
        <div class="filter">
            <span>Old</span>
            <input type="checkbox" v-model="showOld" />
        </div>
    </div>
    <div>
        <template v-for="(schedule, index) in schedules.filter(item => showOld || new Date(item.ReadyDate) >= startOfWeek)" v-bind:key=index>
            <ScheduleItem class="row" :class="{old: new Date(schedule.ReadyDate) < startOfWeek}"
                v-if="editedSchedule == null || schedule.Pastry != editedSchedule.Pastry || schedule.ReadyDate != editedSchedule.ReadyDate"
                @click="editSchedule(schedule)"
                :pastry="schedule.Pastry.Name"
                :quantity="schedule.Quantity"
                :reserved="schedule.Reserved"
                :readydate="new Date(schedule.ReadyDate)"/>
            <template v-if="editedSchedule != null && schedule.Pastry == editedSchedule.Pastry && schedule.ReadyDate == editedSchedule.ReadyDate">
                <ScheduleEditItem class="row-edit"
                    @schedule-submit="handleSubmit()"
                    @schedule-cancel="handleCancel()"
                    :pastryId="schedule.Pastry.Id"
                    :pastryName="schedule.Pastry.Name"
                    :quantity="schedule.Quantity"
                    :reserved="schedule.Reserved"
                    :readydate="new Date(schedule.ReadyDate)"/>
            </template>
        </template>
    </div>
</template>

<style scoped>
.row {
    position: relative;
    cursor: pointer;
}

.row:hover {
    background-color: var(--hover-background-color);
}

.row-edit {
    position: relative;
    cursor: pointer;
}

.old {
    color: var(--color-text-disabled);
}
</style>