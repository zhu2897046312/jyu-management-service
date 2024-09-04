<template>
  <Index />
</template>

<script>
import MyHeader from './components/MyHeader.vue'
import Sidebar from './components/Sidebar.vue'
import MainContent from './components/MainContent.vue'
import MySection from './components/Section.vue'
import Index from './views/index.vue'

export default {
  components: {
    MyHeader,
    Sidebar,
    MainContent,
    MySection,
    Index
  }
}
</script>

<style>
html, body {
  height: 100%;
  margin: 0;
}

#app {
  height: 100%;
  overflow: hidden;
}

</style>
