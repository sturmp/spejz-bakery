<script setup>
const emit = defineEmits(['order-button-click']);

function handleOrderButtonClick() {
    emit('order-button-click');
}
</script>

<template>
        <div id="order-button" @click="handleOrderButtonClick()">Rendelek</div>
</template>

<style>
#order-button {
    height: 2.5em;
    border: 0.18em dotted #2C3140;

    font-size: 1.25rem;
    font-weight: bold;
    display: flex;
    justify-content: center;
    align-items: center;
    text-align: center;

    cursor: pointer;
}

#order-button:hover {
    border: 0.18em dotted #9E876D;
}
</style>