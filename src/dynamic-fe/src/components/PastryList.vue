<script setup>
</script>

<template>
    <div id="item-list">
        <div class="item-row">
            <div class="item">Biscuit<span>: nem az az édes, ez sós... Vajas pogácsa jó?! Vajas pogácsa!</span></div>
            <div class="price">3500</div>
            <div class="unit-weight">/kg</div>
        </div>
        <div class="item-row">
            <div class="item">Foccacia<span>: Olasz olajos kenyér lángos... Feltét nélkül.</span></div>
            <div class="price">~440</div>
            <div class="unit-weight">/db <span>(~125g)</span></div>
        </div>
        <div class="item-row">
            <div class="item">Kenyér<span>: Sima kenyér. Semmi extra.</span></div>
            <div class="price">600</div>
            <div class="unit-weight">/db <span>(750g)</span></div>
        </div>
        <div class="item-row">
            <div class="item">English muffin<span>: Nem, ez nem az édesség. <a href="https://www.google.com/search?client=firefox-b-d&q=english+muffin">Nézz utána!</a></span></div>
            <div class="price">150</div>
            <div class="unit-weight">/db <span>(65g)</span></div>
        </div>
        <div class="item-row">
            <div class="item">Kakaós csiga<span>: Kakaós és fel van tekerve.</span></div>
            <div class="price">500</div>
            <div class="unit-weight">/db <span>(100g)</span></div>
        </div>
        <div class="item-row">
            <div class="item">Tortilla<span>: Mexikói lapos lángos. Kaja origami.</span></div>
            <div class="price">100</div>
            <div class="unit-weight">/db <span>(~20cm)</span></div>
        </div>
        <div class="item-row">
            <div class="item">Heti különlegesség<span>: Bagel. Te a kutyára gondolsz, de az beagle. Ez olyan mint a szalagos fánk csak sós.</span></div>
            <div class="price">350</div>
            <div class="unit-weight">/db <span>(~100g)</span></div>
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