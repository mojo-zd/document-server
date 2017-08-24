<template>
  <div class="application row">
    <app-menu></app-menu>
    <app-main></app-main>
  </div>
</template>

<script>
import AppMenu from './AppMenu'
import AppMain from './AppMain'
export default {
  components: {
    AppMenu,
    AppMain
  },
  name: 'hello',
  data () {
    return {
      msg: 'Welcome to Your Vue.js App'
    }
  }
}
</script>