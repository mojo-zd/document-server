<template>
  <div id="app">
    <navigation></navigation>
    <contaner></contaner>
  </div>
</template>

<script type="text/javascript">
import Navigation from './components/Navigation'
import Contaner from './components/Contaner'
export default {
  components: {
    Navigation,
    Contaner
  },
  name: 'app'
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
   margin-top: -8px; 
}
</style>

<style lang="scss">
@import "./assets/vendor/bootstrap/stylesheets/bootstrap.scss" 
</style>
