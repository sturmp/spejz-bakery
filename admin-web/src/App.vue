<script setup>
import { ref, watch } from 'vue';
import { RouterLink, RouterView, useRoute } from 'vue-router';

const routePath = ref("pastries");
const route = useRoute();

watch(
  () => route.path,
  (path) => {
    routePath.value = path;
  });

</script>

<template>
  <header>
    <div class="wrapper">
      <nav>
        <RouterLink to="/pastries" class="link menu-item" :class="{'link-active': routePath == '/pastries'}">Pastries</RouterLink>
        <RouterLink to="/orders" class="link menu-item" :class="{'link-active': routePath == '/orders'}">Orders</RouterLink>
        <RouterLink to="/schedules" class="link menu-item" :class="{'link-active': routePath == '/schedules'}">Schedules</RouterLink>
        <RouterLink to="/dayoffs" class="link menu-item" :class="{'link-active': routePath == '/dayoffs'}">Dayoffs</RouterLink>
      </nav>
    </div>
  </header>

  <RouterView />
</template>

<style scoped>
header {
  margin-bottom: 2em;
}

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
