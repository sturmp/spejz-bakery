<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <header>
    <div class="wrapper">
      <nav>
        <RouterLink to="/pastries" class="link menu-item">Pastries</RouterLink>
        <RouterLink to="/orders" class="link menu-item">Orders</RouterLink>
        <RouterLink to="/schedules" class="link menu-item">Schedules</RouterLink>
        <RouterLink to="/dayoffs" class="link menu-item">Dayoffs</RouterLink>
      </nav>
    </div>
  </header>

  <RouterView />
</template>

<style scoped>
nav {
  display: flex;
}

.menu-item {
  flex: 1 1 auto;

  margin: 0 1em;

  text-align: center;
  font-size: 1.5em;
}

</style>
