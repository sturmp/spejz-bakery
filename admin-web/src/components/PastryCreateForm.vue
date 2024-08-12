<script setup>
import { ref } from "vue";
import { fetchFromApi } from '@/modules/fetch.mjs';

const emit = defineEmits(["pastry-create"]);

const pastry = ref({
    name: "",
    description: "",
    price: "",
    unitOfMeasureId: 0,
    quantityPerPiece: "",
});

const unitOfMeasures = ref(null);

const urlFetchUnitOfMeasure = `${import.meta.env.VITE_API_URL}/unitofmeasure`;
async function fetchUnitOfMeasuresAsync() {
    unitOfMeasures.value = await (await fetchFromApi(urlFetchUnitOfMeasure)).json();
}

const urlCreatePastry =`${import.meta.env.VITE_API_URL}/pastry`;
async function createPastryAsync() {
    const newPastry = {
        name: pastry.value.name,
        description: pastry.value.description,
        price: pastry.value.price,
        unitOfMeasure: pastry.value.unitOfMeasureId,
        quantityPerPiece: pastry.value.quantityPerPiece,
    };

    const requestOptions = {
        method: 'POST',
        body: JSON.stringify(newPastry)
    };
    await (await fetchFromApi(urlCreatePastry, requestOptions)).json();
}

function handleCreate() {
    createPastryAsync(pastry)
        .then(() => emit('pastry-create'));
}

fetchUnitOfMeasuresAsync();
</script>

<template>
    <div class="pastry-create">
        <div class="pastry-property"><span>Name:</span><input v-model.trim="pastry.name" type="text"/></div>
        <div class="pastry-property"><span>Description:</span><input v-model.trim="pastry.description" type="text"/></div>
        <div class="pastry-property"><span>Price:</span><input v-model="pastry.price" type="text"/></div>
        <div class="pastry-property">
            <select v-model="pastry.unitOfMeasureId">
                <option v-for="unitOfMeasure in unitOfMeasures" :key="unitOfMeasure.Id" :value="unitOfMeasure.Id">{{ unitOfMeasure.Name }}</option>
            </select>
        </div>
        <div class="pastry-property"><span>Quantity/Price:</span><input v-model.trim="pastry.quantityPerPiece" type="text"/></div>
        <div id="create" @click="handleCreate()">Create</div>
    </div>

</template>

<style scoped>
.pastry-create {
    display:grid;
    grid-template-columns: 3fr 6fr 2fr 1fr 3fr 2fr;
    border-bottom: var(--border-size) dotted var(--color-text);
    border-left: var(--border-size) dotted var(--color-text);
    border-right: var(--border-size) dotted var(--color-text);
    border-top: var(--border-size) dotted var(--color-text);
    margin-bottom: 1em;
}

.pastry-property {
    display: flex;
    border-right: var(--border-size) dotted var(--color-text);
    align-items: center;
    padding: 0.5em;
}

.pastry-property:first-of-type {
    padding: 0.5em;
}

.pastry-property:nth-child(6) {
    border-right: 0;
}

input {
    width: 0;
    font-family: inherit;
    font-size: inherit;
    font-weight: inherit;
    border: 0;
    padding: 0.5em;
    flex: 1 1 auto;
}

#create {
    display: flex;
    justify-content: center;
    align-items: center;
    font-weight: bold;
    cursor: pointer;
}

#create:hover {
    color: var(--color-text-highlight);
}
</style>