<template>
  <div id="app">
    <SelectNumberBoard></SelectNumberBoard>
  </div>
</template>

<script>
import SelectNumberBoard from "./components/SelectNumberBoard";

export default {
  name: 'App',
  components: {
    SelectNumberBoard
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
