<script setup>
import { ref } from 'vue';
import { fetchFromApi } from '@/modules/fetch.mjs';
import { useI18n } from 'vue-i18n';
const { t } = useI18n();

const emit = defineEmits(['close-button-click', 'order-submit']);

var selectedPastry = ref();
var quantity = ref(1);
var customer = ref("");
var preferedDate = ref("");
var partOfDay = ref("morning");
var isOrderSuccess = ref(false);
var invalidCustomer = ref(false);
var invalidPreferedDate = ref(false);
var showOrderSentMessage = ref(false);
const pastries = ref(null);

function handleCloseButtonClick() {
    emit('close-button-click');
}

const postUrl = `${import.meta.env.VITE_API_URL}/order`;
async function handleSubmitClick() {
    var date = new Date(preferedDate.value);

    if(isInvalidInput(date)) {
        setTimeout( () => {
            invalidCustomer.value = false;
            invalidPreferedDate.value = false;
        }, 1000)
        return;
    }

    const order = {
        'PastryId': selectedPastry.value,
        'Customer': customer.value,
        'Quantity': quantity.value,
        'PreferedDate': partOfDay.value == "morning"?  new Date(date.setHours(8)) : new Date(date.setHours(16)),
    }
    const requestOptions = {
        method: 'POST',
        body: JSON.stringify(order)
    }

    try {
        const response = await fetchFromApi(postUrl, requestOptions);
        if (response.status == 200) {
            isOrderSuccess.value = true;
        }
    } catch(ex) {
        console.log(ex);
    }

    showOrderSentMessage.value = true;
    setTimeout( () => {
            showOrderSentMessage.value = false;
            emit('order-submit');
        }, 1800)
}

function isInvalidInput(date) {
    var isAnyInputInvalid = false;

    if(customer.value === "") {
        invalidCustomer.value = true;
        isAnyInputInvalid = true;
    }

    if (isNaN(date)) {
        invalidPreferedDate.value = true;
        isAnyInputInvalid = true;
    }

    return isAnyInputInvalid;
}

const fetchUrl = `${import.meta.env.VITE_API_URL}/pastry`;
async function fetchPastriesAsync() {
    pastries.value = await (await fetchFromApi(fetchUrl)).json();
    selectedPastry.value = pastries.value[0].Id;
}

function setInitialPreferedDate() {
    const today = new Date();
    const defaultDate = new Date(today.setDate(today.getDate() + 2));
    preferedDate.value = defaultDate.toISOString().slice(0, 10);
}

fetchPastriesAsync();
setInitialPreferedDate();
</script>

<template>
<div id="container">
    <div id="message" v-if="showOrderSentMessage"><span v-html="t('ordersentmessage')"></span></div>
    <div id="close-button" @click="handleCloseButtonClick()" v-if="!showOrderSentMessage">X</div>
    <select v-model="selectedPastry">
        <option v-for="pastry in pastries" :key="pastry.Id" :value="pastry.Id">{{ pastry.Name }}</option>
    </select>
    <input v-model.number="quantity" type="number" min="0">
    <input v-model="preferedDate" type="date" :class="{ 'validation-error': invalidPreferedDate }">
    <div id="part-of-day">
        <input type="radio" id="morning" value="morning" v-model="partOfDay" />
        <label for="morning">{{ t('morning') }}</label>
        <input type="radio" id="afternoon" value="afternoon" v-model="partOfDay" />
        <label for="afternoon">{{ t('afternoon') }}</label>
    </div>
    <input v-model.trim="customer" type="text" :placeholder="t('name')" :class="{ 'validation-error': invalidCustomer }">
    <div id="submitOrderButton" @click="handleSubmitClick()">Om Nyom Nyom</div>
</div>
</template>

<style>
#container {
    width: 80vw;
    max-width: 600px;
    display: flex;
    flex-direction: column;

    background-color: #ffffff;
    box-shadow: 0em 0.5em 1em 0 rgba(0, 0, 0, 0.35), 0em 0.5em 3em 0 rgba(0, 0, 0, 0.3);

    padding: 2em;
}

#message {
    position: absolute;
    top: 0;
    left: 0;
    z-index: 100;
    height: 100%;
    width: 100%;

    display: flex;
    text-align: center;
    align-items: center;
    justify-content: center;

    background-color: #ffffff;

    font-size: 1.25rem;
}

#close-button {
    position: absolute;
    float: left;
    right: -0.65em;
    top: -0.65em;

    font-size: 1.25rem;
    font-weight: bold;
    display: flex;
    justify-content: center;
    align-items: center;
    text-align: center;

    height: 2em;
    width: 2em;
    background-color: #ffffff;
    border: 0.18em dotted #2C3140;
    border-radius: 50%;

    cursor: pointer;
}

select, input, button, #submitOrderButton, #part-of-day {
    height: 2em;
    margin-top: 1em;
    
    font-size: 1.25rem;
    text-align: center;
    
    /* border-bottom: 0.1em dotted #2C3140; */
    border-bottom: 0em dotted #2C3140;
    border-top: 0;
    border-left: 0;
    border-right: 0;
}

#container > select {
    margin-top: 0 !important;
}

#submitOrderButton {
    margin-top: 1em;

    border: 0.18em dotted #2C3140;

    display: flex;
    justify-content: center;
    align-items: center;
    text-align: center;

    cursor: pointer;
}

#part-of-day {
    display: flex;
    align-items: center;
}

#part-of-day > input {
    flex: 1 1 auto;
    margin: 0;

    font-size: 0.75rem;
}

#part-of-day > label {
    flex: 1 1 auto;
    text-align: left;
}

#submitOrderButton:hover, #close-button:hover {
    border: 0.18em dotted #9E876D;
}

@keyframes alert {
    0% {background-color: #ffffff;}
    25% {background-color: #ff8888;}
    50% {background-color: #ffffff;}
    75% {background-color: #ff8888;}
    100% {background-color: #ffffff;}
}

.validation-error {
    animation: alert 1s linear 0s 1 normal none;
}
</style>