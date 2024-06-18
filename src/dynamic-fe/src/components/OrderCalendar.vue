<script setup>
</script>

<template>
    <div id="calendar">
        <h3>06.17 - 06.23</h3>
        <div class="row">
            <div class="day">Hétfő</div>
            <div class="orders">
                <div class="order">X</div>
            </div>
        </div>
        <div class="row">
            <div class="day">Kedd</div>
            <div class="orders">
                <div class="order">X</div>
            </div>
        </div>
        <div class="row">
            <div class="day">Szerda</div>
            <div class="double-row">
                <div class="orders">
                    <div class="order part-of-day">Délelőtt</div>
                    <div class="order">Kakaós csiga<br>(2/12)</div>
                    <div class="order">Kenyér<br>(1/2)</div>
                </div>
            </div>
        </div>
        <div class="row">
            <div class="day">Csütörtök</div>
            <div class="orders">
                <div class="order"></div>
            </div>
        </div>
        <div class="row">
            <div class="day">Péntek</div>
            <div class="double-row">
                <div class="orders">
                    <div class="order part-of-day">Délelőtt</div>
                    <div class="order">Bagel<br>(0/10)</div>
                    <div class="order">Kenyér<br>(2/2)</div>
                </div>
                <div class="orders">
                    <div class="order part-of-day">Délután</div>
                    <div class="order">Biscuit<br>(1kg/1kg)</div>
                </div>
            </div>
        </div>
        <div class="row">
            <div class="day">Szombat</div>
            <div class="orders">
                <div class="order">X</div>
            </div>
        </div>
        <div class="row">
            <div class="day">Vasárnap</div>
            <div class="orders">
                <div class="order">X</div>
            </div>
        </div>
    </div>
</template>

<style scoped>
#calendar {
    display: flex;
    flex-direction: column;
    height: 100%;
}

h3 {
    flex: 0 1 auto;
    text-align: center;
    display: flex;
    justify-content: center;
    align-items: center;
    margin-bottom: 1em;
    font-size: 1.4em;
    font-weight: 500;
}

.row {
    display: flex;
    flex: 1 1 auto;
}

#calendar .row:first-of-type {
    border-top: 0.15em dotted #2C3140;
}

.day {
    flex: 0 1 5em;

    font-weight: bold;
    border-left: 0.15em dotted #2C3140;
    border-right: 0.15em dotted #2C3140;
    border-bottom: 0.15em dotted #2C3140;

    display: flex;
    justify-content: center;
    align-items: center;
    text-align: center;
}

.double-row {
    display: flex;
    flex-direction: column;
    flex: 1 1 auto;
}

.orders {
    display: flex;
    flex: 1 1 auto;
}

.order {
    border-right: 0.15em dotted #2C3140;
    border-bottom: 0.15em dotted #2C3140;

    display: flex;
    justify-content: center;
    align-items: center;
    text-align: center;

    flex: 1 1 auto;
}

.part-of-day {
    writing-mode: vertical-rl;
    flex: 0 1 auto;
}
</style>