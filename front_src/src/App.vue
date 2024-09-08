<template>
  <router-link to="/login"></router-link> 
  <router-view></router-view>
  <EnrollCourse />
</template>

<script>
import MyHeader from './components/MyHeader.vue'
import MainContent from './components/MainContent.vue'
import MySection from './components/Section.vue'
import Index from './views/Index.vue'
import Login from './views/Login.vue'
import EnrollCourse from './views/EnrollCourse.vue'

export default {
  components: {
    MyHeader,
    MainContent,
    MySection,
    Index,
    Login,
    EnrollCourse
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
