<script setup>
import { formatDate } from '@/modules/datetime.mjs';

defineProps({
    pastry: String,
    quantity: Number,
    reserved: Number,
    readydate: Date,
});

</script>

<template>
    <div class="schedule">
        <div class="schedule-property">{{ pastry }}</div>
        <div class="schedule-property">{{ quantity }}</div>
        <div class="schedule-property">{{ reserved }}</div>
        <div class="schedule-property">{{ formatDate(readydate) }}</div>
    </div>
</template>

<style scoped>
.schedule {
    display:grid;
    grid-template-columns: 1fr 1fr 1fr 1fr;
    border-bottom: var(--border-size) dotted var(--color-text);
    border-left: var(--border-size) dotted var(--color-text);
    border-right: var(--border-size) dotted var(--color-text);
}

.schedule:first-of-type {
    border-top: var(--border-size) dotted var(--color-text);
}

.schedule-property {
    display: flex;
    padding: 0.5em 0.5em;
    border-right: var(--border-size) dotted var(--color-text);

    align-items: center;
}

.schedule-property:last-child {
    border-right: 0;
}
</style>