<template>
  <div class="card">
    <h3>消息</h3>
    <ul>
      <li v-for="(message, index) in messages" :key="index">
        <a href="#">
          <span class="message-content">{{ message.content }}</span>
          <span class="message-time">{{ message.time }}</span>
        </a>
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  data() {
    return {
      messages: [
        { time: '2024-05-29 09:14:49', content: '老师发布了新的作业通知' },
        { time: '2024-04-06 13:27:36', content: '作业截止时间即将到来' },
        { time: '2024-05-27 16:38:32', content: '新的学习资料已上传' },
        { time: '2024-05-05 10:23:14', content: '期中考试安排已公布' },
        { time: '2024-04-22 09:21:43', content: '教学计划已更新' }
      ]
    }
  }
}
</script>

<style scoped>
.card {
  padding: 15px;
  background-color: white;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  border-radius: 8px;
  margin-bottom: 20px;
}

h3 {
  font-size: 18px;
  margin-bottom: 15px;
}

ul {
  list-style: none;
  padding: 0;
  margin: 0;
}

li {
  margin-bottom: 10px;
  font-size: 14px;
  display: flex;
  justify-content: space-between;
  border-bottom: 1px solid #eee;
  padding-bottom: 5px;
}

.message-content {
  flex: 1;
}

.message-time {
  color: #999;
  font-size: 12px;
  white-space: nowrap;
  margin-left: 10px;
}
</style>
