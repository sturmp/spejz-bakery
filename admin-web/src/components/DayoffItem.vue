<script setup>
import { formatDateWithDayName } from '@/modules/datetime.mjs';
import { fetchFromApi } from '@/modules/fetch.mjs';

defineProps({
    id: Number,
    dayoff: Date,
});

const emits = defineEmits(['dayoff-deleted']);

const url =`${import.meta.env.VITE_API_URL}/dayoff/`;
async function deleteDayoffAsync(id) {
    const requestOptions = {
        method: 'DELETE',
    };
    fetchFromApi(url + id, requestOptions)
        .then(() => emits('dayoff-deleted'));
}
</script>

<template>
    <div class="dayoff">
        <div class="dayoff-property">{{ formatDateWithDayName(dayoff) }}</div>
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