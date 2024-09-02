<template>
  <div id="app">

    <MyHeader />

    <div class="main-layout">
      <Sidebar />
      <MainContent />
    </div>
  </div>
</template>

<script>
import MyHeader from './components/MyHeader.vue'
import Sidebar from './components/Sidebar.vue'
import MainContent from './components/MainContent.vue'

export default {
  components: {
    MyHeader,
    Sidebar,
    MainContent
  }
}
</script>

<style>
.main-layout {
  display: flex;
}
</style>
