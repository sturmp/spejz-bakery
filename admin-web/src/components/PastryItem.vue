<script setup>
defineProps({
    name: String,
    description: String,
    price: String,
    unitOfMeasure: String,
    quantityPerPiece: String,
    enabled: Boolean,
})
</script>

<template>
    <div class="pastry">
        <div class="pastry-property left">{{ name }}</div>
        <div class="pastry-property left">{{ description }}</div>
        <div class="pastry-property right">{{ price }}</div>
        <div class="pastry-property left">{{ unitOfMeasure }}</div>
        <div class="pastry-property middle">{{ quantityPerPiece }}</div>
        <div class="pastry-property middle">{{ enabled? "✔": "X"}}</div>
    </div>
</template>

<style scoped>
.pastry {
    display:grid;
    grid-template-columns: 2fr 9fr 1fr 1fr 1fr 1fr;
    border-bottom: var(--border-size) dotted var(--color-text);
    border-left: var(--border-size) dotted var(--color-text);
    border-right: var(--border-size) dotted var(--color-text);
}

.pastry:first-of-type {
    border-top: var(--border-size) dotted var(--color-text);
}

.pastry-property {
    display: flex;
    padding: 0.5em 0.5em;
    border-right: var(--border-size) dotted var(--color-text);

    align-items: center;
}

.pastry-property:last-child {
    border-right: 0;
}
</style>