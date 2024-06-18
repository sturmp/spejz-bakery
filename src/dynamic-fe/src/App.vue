<script setup>
import SiteHeader from './components/SiteHeader.vue'
import PastryList from './components/PastryList.vue'
import OrderCalendar from './components/OrderCalendar.vue'
</script>

<template>
  <main>
    <div id="header-container">
      <SiteHeader/>
    </div>
    <div id="content">
      <!-- <PastryList /> -->
      <OrderCalendar/>
    </div>
  </main>
</template>

<style scoped>

main {
  display: flex;
  flex-direction: column;
  height: 100%;
  max-width: 800px;
  margin: auto;
}

#header-container {
  flex: 0 1 auto;
}

#content {
  flex: 1 1 auto;
}

</style>
