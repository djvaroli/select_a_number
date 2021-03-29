<template>
  <div id="app">
    <Header></Header>
    <PageBanner></PageBanner>
    <SelectNumberBoard></SelectNumberBoard>
    <Footer></Footer>
  </div>
</template>

<script>
import SelectNumberBoard from "./components/SelectNumberBoard";
import Footer from "./components/navigation/Footer";
import Header from "./components/navigation/Header";
import PageBanner from "./components/utilities/PageBanner";

export default {
  name: 'App',
  components: {
    SelectNumberBoard,
    Header,
    Footer,
    PageBanner
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
}
</style>
