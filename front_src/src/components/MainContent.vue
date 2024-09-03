<template>
  <div class="main-content">
    <div class="container">
      <div class="box">
        <h2>{{ studentProfile.name }}</h2>
        <p>{{ studentProfile.class }}</p>
      </div>
      <Notifications class="box"/>
      <Schedule class="box"/>
      <Calendar class="box"/>
      <Messages class="box"/>
      <Grades class="box"/>
    </div>
  </div>
</template>

<script>
import Notifications from './Notifications.vue'
import Schedule from './Schedule.vue'
import Messages from './Messages.vue'
import Calendar from './Calendar.vue'
import Grades from './Grades.vue'

export default {
  components: {
    Notifications,
    Schedule,
    Messages,
    Calendar,
    Grades
  },
  data() {
    return {
      studentProfile: {
        name: '朱佳',
        class: '计算机学院计算机2023班',
        picture: '/path/to/avatar.jpg' // 头像图片路径
      }
    }
  }
}
</script>

<style scoped>
.main-content {
  display: flex;
  flex-direction: column;
  height: 100%; /* 确保 main-content 占满其父元素的高度 */
  overflow: hidden; /* 防止出现滚动条 */
}

.container {
  display: grid;
  grid-template-columns: repeat(2, 1fr); /* 每行两个盒子 */
  grid-template-rows: repeat(3, auto); /* 三行 */
  gap: 10px; /* 控制盒子之间的间距 */
  height: 100%; 
  overflow: auto; /* 允许内容溢出时出现滚动条 */
}

.box {
  padding: 10px;
  border: 1px solid #ddd;
  box-sizing: border-box; /* 确保 padding 和 border 包含在元素的总宽度和高度中 */
}
</style>
