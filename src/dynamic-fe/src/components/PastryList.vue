<script setup>
import { ref } from 'vue';

const pastries = ref(null);

const url = "http://localhost:5555/pastry";
async function fetchPastries() {
    pastries.value = await (await fetch(url)).json();
}

fetchPastries();

</script>

<template>
    <div id="item-list">
        <div class="item-row" v-for="pastry in pastries" :key="pastry.Name">
            <div class="item">{{ pastry.Name }}<span v-html="': ' + pastry.Description"></span></div>
            <div class="price">{{ pastry.Price }}</div>
            <div class="unit-weight">/{{ pastry.UnitOfMeasure}}<span v-if="pastry.QuantityPerPiece != ''"> ({{ pastry.QuantityPerPiece}})</span></div>
        </div>
    </div>
</template>

<style scoped>
#item-list {
    display: flex;
    flex-direction: column;
    height: 100%;
}

#item-list > :nth-child(1) {
    border-top: 0.15em dotted #2C3140;
}

.item-row {
    display: grid;
    grid-template-columns: 6fr 2em 5.5em;
    border-bottom: 0.15em dotted #2C3140;
    align-items: center;
    flex: 1 1 auto;
}

.item {
    font-weight: bold;
    line-height: 1.5;
    padding-right: 0.5em;
}

.item a {
    color: #9E876D;
    text-decoration: none;
}

.item span {
    font-weight: initial;
}

.price {
    padding-right: 0.3em;
    display: flex;
    justify-content: right;
    font-weight: bold;
}

.unit-weight span {
    font-style: italic;
    font-weight: initial;
}
</style>