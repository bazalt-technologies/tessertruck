<template>
  <div id="app">
    <HeaderComponent />
    <router-view class="routerView" />
  </div>
</template>

<script>

import HeaderComponent from "@/components/HeaderComponent.vue";

export default {
  name: 'App',
  components: {
    HeaderComponent

  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  margin: 0;
}
body {
  margin: 0;
}
.routerView {
  margin-top: 15px;
}
</style>
