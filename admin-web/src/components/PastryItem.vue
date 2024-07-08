<script setup>
defineProps({
    name: String,
    description: String,
    price: String,
    unitOfMeasure: String,
    quantityPerPiece: String,
})
</script>

<template>
    <div class="pastry">
        <div class="pastry-property">{{ name }}</div>
        <div class="pastry-property">{{ description }}</div>
        <div class="pastry-property">{{ price }}</div>
        <div class="pastry-property">{{ unitOfMeasure }}</div>
        <div class="pastry-property">{{ quantityPerPiece }}</div>
    </div>
</template>

<style scoped>
.pastry {
    display:grid;
    grid-template-columns: 2fr 9fr 1fr 1fr 1fr;
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