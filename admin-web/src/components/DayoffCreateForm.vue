<script setup>
import { ref } from 'vue';
import { defineEmits } from 'vue';

const emits = defineEmits(['dayoff-created']);

var day = ref("");

const url =`${import.meta.env.VITE_API_URL}/dayoff`;
async function createDayoffAsync() {
    const requestOptions = {
        method: 'POST',
        headers: { 'AuthToken': import.meta.env.VITE_API_AUTH_TOKEN },
        body: JSON.stringify(new Date(day.value))
    };
    await fetch(url, requestOptions);

    day.value = "";
    emits('dayoff-created');
}
</script>

<template>
<form>
    <input v-model="day" type="date" />
    <div id="button" @click="createDayoffAsync()">Create</div>
</form>
</template>

<style scoped>
form {
    margin-bottom: 1em;
    display:grid;
    grid-template-columns: 15fr 1fr;
    border: var(--border-size) dotted var(--color-text);
}

input {
    border: 0;
    padding: 0.5em 0.5em;
    font-family: inherit;
    font-size: inherit;
}

#button {
    border-left: var(--border-size) dotted var(--color-text);
    display: flex;
    cursor: pointer;
    justify-content : center;
    font-weight: bold;
    padding: 0.5em 0.5em;
}

#button:hover {
    color: var(--color-text-highlight)
}
</style>