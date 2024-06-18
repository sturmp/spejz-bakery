<script setup>
</script>

<template>
    <div id="header">
        <h1 id="title">Spejz Pékség</h1>
        <div id="menu">
            <div class="menu-item">
                <a href="#item-list">
                    <img src="../assets/imgs/sourdough.png">
                    <img src="../assets/imgs/cinnamon-roll.png">
                    <img src="../assets/imgs/akara.png">
                </a>
            </div>
            <div class="menu-item">
                <a href="#calendar">
                    <img src="../assets/imgs/calendar_02.png">
                </a>
            </div>
        </div>
    </div>
</template>

<style scoped>
#header {
    display: grid;
    grid-template-columns: 1fr;
    grid-template-rows: 1fr;
    margin-bottom: 1.5em;
}

#title {
    text-align: center;
    padding-top: 0.5em;
    padding-bottom: 0.75em;
    font-weight: 500;
}

#menu {
    display: grid;
    grid-template-columns: 1fr 1fr;
    grid-template-rows: auto;
    justify-items: stretch;
    align-items: stretch;
    column-gap: 0.75em;
    margin-left: 0.5em;
    margin-right: 0.5em;
}

.menu-item {
    display: flex;
    border: 0.15em dotted #9E876D;
    padding-top: 0.5em;
    padding-bottom: 0.5em;
    justify-content: center;
}

img{
    height: 2em;
    margin-left: 0.2em;
    margin-right: 0.2em;
    position: relative;
    top: 50%;
    transform: translateY(-50%);
}
</style>