<template>
  <div id="app">
    <MyHeader />

    <MySection />

    <div class="main-layout">
      <Sidebar />
      <div class="main-content-wrapper">
        <MainContent />
      </div>
    </div>
  </div>
</template>

<script>
import MyHeader from './components/MyHeader.vue'
import Sidebar from './components/Sidebar.vue'
import MainContent from './components/MainContent.vue'
import MySection from './components/Section.vue'

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
  display: flex;
  flex-direction: column;
  height: 100%;
}

.main-layout {
  display: flex;
  flex: 1; /* 使 main-layout 占满剩余空间 */
}

.sidebar {
  width: 250px; /* 设定 sidebar 的宽度 */
  background-color: #f4f4f4; /* 可选: 设定 sidebar 的背景色 */
}

.main-content-wrapper {
  flex: 1; /* 使 main-content-wrapper 占据剩余的空间 */
  display: flex; /* 使用 flexbox 以确保子元素布局 */
}

.main-content {
  flex: 1; /* 使 main-content 占据剩余的空间 */
  overflow: hidden; /* 防止出现滚动条 */
  background-color: #fff; /* 可选: 设定 main-content 的背景色 */
}

</style>
