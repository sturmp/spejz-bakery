<script setup>
import { ref } from 'vue';
import PastryItem from "../components/PastryItem.vue"
import PastryEditItem from "../components/PastryEditItem.vue"
import { fetchFromApi } from '@/modules/fetch.mjs';

const pastries = ref(null);
const editedPastry = ref(null)

const url =`${import.meta.env.VITE_API_URL}/pastry`;
async function fetchPastriesAsync() {
    pastries.value = await (await fetchFromApi(url)).json();
}

function editPastry(pastry) {
    editedPastry.value = pastry;
}

function handleSubmit() {
    fetchPastriesAsync()
        .then(() =>editedPastry.value = null);
}

function handleCancel() {
    editedPastry.value = null;
}

fetchPastriesAsync();
</script>

<template>
    <div>
        <template v-for="(pastry, index) in pastries" v-bind:key=index>
            <PastryItem class="row" v-if="editedPastry == null || pastry.Name != editedPastry.Name"
                @click="editPastry(pastry)"
                :name="pastry.Name"
                :description="pastry.Description"
                :price="pastry.Price"
                :unitOfMeasure="pastry.UnitOfMeasure"
                :quantityPerPiece="pastry.QuantityPerPiece"/>
            <template v-if="editedPastry != null && pastry.Name == editedPastry.Name">
                <PastryEditItem class="row-edit"
                    @submit="handleSubmit()"
                    @cancel="handleCancel()"
                    :id="pastry.Id"
                    :name="pastry.Name"
                    :description="pastry.Description"
                    :price="pastry.Price"
                    :unitOfMeasure="pastry.UnitOfMeasure"
                    :quantityPerPiece="pastry.QuantityPerPiece"/>
            </template>
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