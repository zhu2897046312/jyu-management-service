<template>
  <router-link to="/login"></router-link> 
  <router-view></router-view>
  <Login />
</template>

<script>
import MyHeader from './components/MyHeader.vue'
import Sidebar from './components/Sidebar.vue'
import MainContent from './components/MainContent.vue'
import MySection from './components/Section.vue'
import Index from './views/Index.vue'
import Login from './views/Login.vue'

export default {
  components: {
    MyHeader,
    Sidebar,
    MainContent,
    MySection,
    Index,
    Login
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
