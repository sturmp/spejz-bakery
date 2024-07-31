<script setup>
defineProps({
    id: Number,
    dayoff: Date,
});

const emits = defineEmits(['dayoff-deleted']);

const days = [
    'Sunday',
    'Monday',
    'Tuesday',
    'Wednesday',
    'Thursday',
    'Friday',
    'Saturday'
];

function formatDate(date) {
    const month = date.getMonth() + 1;
    const monthString = month < 10 ? `0${month}` : month;
    const day = date.getDate();
    const dayString = day < 10 ? `0${day}` : day;

    return `${date.getFullYear()}-${monthString}-${dayString} ${days[date.getDay()]}`;
}

const url =`${import.meta.env.VITE_API_URL}/dayoff/`;
async function deleteDayoffAsync(id) {
    const requestOptions = {
        method: 'DELETE',
        headers: { 'AuthToken': import.meta.env.VITE_API_AUTH_TOKEN },
    };
    await fetch(url + id, requestOptions);
    emits('dayoff-deleted');
}
</script>

<template>
    <div class="dayoff">
        <div class="dayoff-property">{{ formatDate(dayoff) }}</div>
        <div class="dayoff-property controll" @click="deleteDayoffAsync(id)">X</div>
    </div>
</template>

<style scoped>
.dayoff {
    display:grid;
    grid-template-columns: 15fr 1fr;
    border-bottom: var(--border-size) dotted var(--color-text);
    border-left: var(--border-size) dotted var(--color-text);
    border-right: var(--border-size) dotted var(--color-text);
}

.dayoff:first-of-type {
    border-top: var(--border-size) dotted var(--color-text);
}

.dayoff-property {
    display: flex;
    padding: 0.5em 0.5em;
    border-right: var(--border-size) dotted var(--color-text);

    align-items: center;
}

.dayoff-property:last-child {
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