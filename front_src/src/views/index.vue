<template>
    <div id="app">
      <MyHeader />
  
      <MySection />
  
      <MainContent />
    </div>
  </template>
  
  <script>
  import MyHeader from '../components/MyHeader.vue'
  import Sidebar from '../components/Sidebar.vue'
  import MainContent from '../components/MainContent.vue'
  import MySection from '../components/Section.vue'
  
  export default {
    components: {
      MyHeader,
      Sidebar,
      MainContent,
      MySection
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
  