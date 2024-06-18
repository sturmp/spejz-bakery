<script setup>
import { ref } from 'vue'
import SiteHeader from './components/SiteHeader.vue'
import PastryList from './components/PastryList.vue'
import OrderCalendar from './components/OrderCalendar.vue'

const showPastryList = ref(true)
const showOrderCalendar = ref(false)

function handleSwitchToPastryList() {
  showPastryList.value = true;
  showOrderCalendar.value = false;
}

function handleSwitchToOrderCalendar() {
  showPastryList.value = false;
  showOrderCalendar.value = true;
}
</script>

<template>
  <main>
    <div id="header-container">
      <SiteHeader @switch-to-pastrylist="handleSwitchToPastryList" @switch-to-ordercalendar="handleSwitchToOrderCalendar" />
    </div>
    <div id="content">
      <PastryList v-show="showPastryList" />
      <OrderCalendar v-show="showOrderCalendar" />
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

.hidden {
  display: none;
}

</style>
