<script setup>
import { ref } from 'vue'
import SiteHeader from './components/SiteHeader.vue'
import PastryList from './components/PastryList.vue'
import OrderCalendar from './components/OrderCalendar.vue'
import OrderButton from './components/OrderButton.vue'
import OrderForm from './components/OrderForm.vue'

const showPastryList = ref(true)
const showOrderCalendar = ref(false)
const showOrderForm = ref(false)

function handleSwitchToPastryList() {
  showPastryList.value = true;
  showOrderCalendar.value = false;
}

function handleSwitchToOrderCalendar() {
  showPastryList.value = false;
  showOrderCalendar.value = true;
}

function handleOrderButtonClick() {
  showOrderForm.value = true;
}

function handleOrderFormCloseButtonClick() {
  showOrderForm.value = false;
}

function handleOrderSubmit() {
  showOrderForm.value = false;
}
</script>

<template>
  <main>
    <div id="obscure" v-show="showOrderForm"></div>
    <div id="header-container">
      <SiteHeader @switch-to-pastrylist="handleSwitchToPastryList" @switch-to-ordercalendar="handleSwitchToOrderCalendar" />
    </div>
    <div id="content">
      <PastryList v-show="showPastryList" />
      <OrderCalendar v-show="showOrderCalendar" />
    </div>
    <OrderButton class="order" @order-button-click="handleOrderButtonClick"/>
    <OrderForm class="order-form" v-show="showOrderForm" @close-button-click="handleOrderFormCloseButtonClick" @order-submit="handleOrderSubmit"/>
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

#obscure {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  backdrop-filter: blur(0.1em);
  z-index: 2;
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

.order {
  flex: 0 1 auto;
  margin-top: 0.2em;
}

.order-form {
  position: absolute;
  left: 50%;
  top: 50%;
  -webkit-transform: translate(-50%, -50%);
  transform: translate(-50%, -50%);
  z-index: 3;
}

</style>
