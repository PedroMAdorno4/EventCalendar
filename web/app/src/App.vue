<template>
  <div id="app">
    <Header/>
    <router-view />
  </div>
</template>

<script>
import Header from '@/components/Header.vue'

export default {
  components: {
    Header,
  }
}
</script>

<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  display: flex;
  flex-flow: column;
  height: 100%;
}

.box {
  display: flex;
  flex-flow: column;
  height: 100%;
}

.box .header {
  flex: 0 1 auto;
}

.box .content {
  height: 1%;
  flex: 1 1 auto;
}

#nav {
  a {
    outline: 0;
  }
}
</style>
