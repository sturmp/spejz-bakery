<script setup>
import { ref } from "vue";
import { fetchFromApi } from '@/modules/fetch.mjs';

const props = defineProps({
    id: Number,
    name: String,
    description: String,
    price: String,
    unitOfMeasure: String,
    quantityPerPiece: String,
    enabled: Boolean,
});

const emit = defineEmits(["submit", "cancel"]);

const pastry = ref({
    id: props.id,
    name: props.name,
    description: props.description,
    price: props.price,
    unitOfMeasure: props.unitOfMeasure,
    quantityPerPiece: props.quantityPerPiece,
    enabled: props.enabled,
});

const url =`${import.meta.env.VITE_API_URL}/pastry`;
async function updatePastryAsync() {
    const requestOptions = {
        method: 'PUT',
        body: JSON.stringify(pastry.value)
    };
    await (await fetchFromApi(url, requestOptions)).json();
}

function handleSubmit() {
    updatePastryAsync(pastry)
        .then(() => emit('submit'));
}

</script>

<template>
    <div class="pastry-edit">
        <div class="pastry-property"><input class="left" v-model.trim="pastry.name" type="text"/></div>
        <div class="pastry-property"><input class="left" v-model.trim="pastry.description" type="text"/></div>
        <div class="pastry-property"><input class="right" v-model.trim="pastry.price" type="text"/></div>
        <div class="pastry-property"><input class="left" v-model.trim="pastry.unitOfMeasure" type="text"/></div>
        <div class="pastry-property"><input class="middle" v-model.trim="pastry.quantityPerPiece" type="text"/></div>
        <div class="pastry-property"><input class="middle" v-model.trim="pastry.enabled" type="checkbox"/></div>
        <div class="controlls-group">
            <div class="controll" @click="handleSubmit()">✓</div>
            <div class="controll" @click="emit('cancel')">x</div>
        </div>
    </div>

</template>

<style scoped>
.pastry-edit {
    background-color: var(--edit-background-color);
    display:grid;
    grid-template-columns: 2fr 9fr 1fr 1fr 1fr 1fr;
    border-bottom: var(--border-size) dotted var(--color-text);
    border-left: var(--border-size) dotted var(--color-text);
    border-right: var(--border-size) dotted var(--color-text);
}

.pastry-edit:first-of-type {
    border-top: var(--border-size) dotted var(--color-text);
}

.pastry-property {
    color: var(--color-edit-text);
    display: flex;
    padding: 0.5em 0.5em;
    border-right: var(--border-size) dotted var(--color-text);
    
    align-items: center;
}

.pastry-property:first-of-type {
    padding: 0.5em;
}

.pastry-property:last-child {
    border-right: 0;
}

.controlls-group {
    z-index: 10;
    position: absolute;
    top: -1.25em;
    right: -1.25em;
}

.controll {
    margin-bottom: 0.5em;
    
    font-size: 1.5rem;
    font-weight: bold;
    display: flex;
    justify-content: center;
    align-items: center;
    text-align: center;
    
    height: 1.5em;
    width: 1.5em;
    background-color: #ffffff;
    border: var(--controll-border-size) dotted var(--border-color);
    border-radius: 50%;

    cursor: pointer;
}

.controll:hover {
    border: var(--controll-border-size) dotted var(--border-color-hover);
}

input {
    color: var(--color-edit-text);
    background-color: var(--edit-background-color);
    width: 0;
    font-family: inherit;
    font-size: inherit;
    font-weight: inherit;
    border: 0;
    flex: 1 1 auto;
}
</style>