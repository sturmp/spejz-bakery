<script setup>
import { ref } from 'vue';
import DayoffItem from "../components/DayoffItem.vue"
import DayoffCreateForm from "../components/DayoffCreateForm.vue"
import { fetchFromApi } from '@/modules/fetch.mjs';
import { calculateStartOfWeek } from '@/modules/datetime.mjs';

const dayoffs = ref(null);
const showOld = ref(false);

const startOfWeek = calculateStartOfWeek();

const url =`${import.meta.env.VITE_API_URL}/dayoff`;
async function fetchDayoffsAsync() {
    dayoffs.value = await (await fetchFromApi(url)).json();
    dayoffs.value.sort((a, b) => a.Day > b.Day)
}

function handleDayoffCreated() {
    fetchDayoffsAsync();
}

function handleDayoffDeleted() {
    fetchDayoffsAsync();
}

fetchDayoffsAsync();
</script>

<template>
    <div class="filters">
        <div class="filter">
            <span>Old</span>
            <input type="checkbox" v-model="showOld" />
        </div>
    </div>
    <div>
        <DayoffCreateForm @dayoff-created="handleDayoffCreated()"/>
        <template v-for="(dayoff, index) in dayoffs.filter(item => showOld || new Date(item.Day) >= startOfWeek)" v-bind:key=index>
            <DayoffItem class="row"
                :id="dayoff.Id"
                :dayoff="new Date(dayoff.Day)"
                @dayoff-deleted="handleDayoffDeleted()"/>
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
</style>