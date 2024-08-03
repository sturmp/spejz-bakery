<script setup>
import { ref } from 'vue';
import PastryItem from "../components/PastryItem.vue"
import PastryEditItem from "../components/PastryEditItem.vue"
import { fetchFromApi } from '@/modules/fetch.mjs';
import router from '@/router';

const pastries = ref(null);
const editedPastry = ref(null);
var language = "";

function initLanguage() {
    if (localStorage.getItem('lang') == null)
    {
        localStorage.setItem('lang', 'hu');
    }

    language = localStorage.getItem('lang');
}

function handleLanguageSwitch() {
    switch(localStorage.getItem('lang')) {
        case "en": localStorage.setItem('lang', 'hu'); break;
        case "hu": localStorage.setItem('lang', 'en'); break;
    }

    router.go(0);
}

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
initLanguage();
</script>

<template>
    <div id="container">
        <button @click="handleLanguageSwitch()">{{ language }}</button>
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
#container {
    position: relative;
}

button {
    position: absolute;
    right: -1em;
    top: -1em;
    display: flex;
    z-index: 2;

    font-size: 0.95rem;
    font-weight: bold;
    justify-content: center;
    align-items: center;
    text-align: center;

    height: 2em;
    width: 2em;
    background-color: var(--color-background);
    border: 0.2em dotted var(--border-color);
    border-radius: 50%;

    cursor: pointer;
}

button:hover {
    border: 0.2em dotted var(--border-color-hover);
}

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