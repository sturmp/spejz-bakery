<script setup>
import { ref } from 'vue';
import DayoffItem from "../components/DayoffItem.vue"
import DayoffCreateForm from "../components/DayoffCreateForm.vue"

const dayoffs = ref(null);

const url =`${import.meta.env.VITE_API_URL}/dayoff`;
async function fetchDayoffsAsync() {
    const requestOptions = {
        method: 'GET',
        headers: { 'AuthToken': import.meta.env.VITE_API_AUTH_TOKEN }
    };
    dayoffs.value = await (await fetch(url, requestOptions)).json();
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
    <div>
        <DayoffCreateForm @dayoff-created="handleDayoffCreated()"/>
        <template v-for="(dayoff, index) in dayoffs" v-bind:key=index>
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