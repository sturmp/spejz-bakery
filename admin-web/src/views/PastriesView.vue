<script setup>
import { ref } from 'vue';
import PastryItem from "../components/PastryItem.vue"
import PastryEditItem from "../components/PastryEditItem.vue"

const pastries = ref(null);
const editedPastry = ref(null)

const url =`${import.meta.env.VITE_API_URL}/pastry`;
async function fetchPastriesAsync() {
    const requestOptions = {
        method: 'GET',
        headers: { 'AuthToken': import.meta.env.VITE_API_AUTH_TOKEN }
    };
    pastries.value = await (await fetch(url, requestOptions)).json();
}

function editPastry(pastry) {
    editedPastry.value = pastry;
    console.log(pastry.value);
}

function handleSubmit() {
    fetchPastriesAsync();
    editedPastry.value = null;
    fetchPastriesAsync();
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