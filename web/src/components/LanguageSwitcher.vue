<script setup>
import { useI18n } from 'vue-i18n';

const { locale } = useI18n();

function handleSwitch() {
    switch(locale.value) {
        case "hu": locale.value = "en"; break;
        case "en": locale.value = "hu"; break;
    }
    localStorage.setItem('lang', locale.value);

    location.reload();
}

function initLanguage() {
    if(localStorage.getItem('lang')){
        locale.value = localStorage.getItem('lang');
        return
    }

    locale.value = 'en'
    localStorage.setItem('lang', locale.value)
}

initLanguage();
</script>

<template>
    <button @click="handleSwitch()">{{ locale }}</button>
</template>

<style scoped>
button {
    font-size: 0.95rem;
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

button:hover {
    border: 0.18em dotted #9E876D;
}
</style>