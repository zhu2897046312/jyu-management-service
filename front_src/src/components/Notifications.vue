<template>
  <div class="notifications-card">
    <h3>通知</h3>
    <ul>
      <li v-for="(notification, index) in notifications" :key="index" :class="notification.type">
        <span class="type">{{ getNotificationLabel(notification.type) }}</span>
        <span class="content">{{ notification.content }}</span>
        <span class="timestamp">{{ notification.timestamp }}</span>
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  data() {
    return {
      notifications: [
        { type: '重要', content: '超星平台用户绑定的提示', timestamp: '2024-01-09' },
        { type: '提醒', content: '请同学及时完成网上评教的提醒', timestamp: '2023-12-26' },
        { type: '通知', content: '教务系统新增“问卷”在线问答的功能，欢迎同学们试用…', timestamp: '2023-10-11' },
        { type: '提醒', content: '有关考研和考研口语的一些调整，请同学们关注', timestamp: '2023-10-10' },
        { type: '重要', content: '本学期的学位英语考试安排已经确定，请大家注意', timestamp: '2023-10-09' },
      ]
    }
  },
  methods: {
    // 根据类型返回不同的文本标签
    getNotificationLabel(type) {
      switch (type) {
        case '重要':
          return '[重要]';
        case '提醒':
          return '[提醒]';
        case '通知':
          return '[通知]';
        default:
          return '';
      }
    }
  }
}
</script>

<style scoped>
.notifications-card {
  padding: 15px;
  background-color: white;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  border-radius: 8px;
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
  display: flex;
  justify-content: space-between;
  margin-bottom: 10px;
  padding: 10px;
  border-radius: 4px;
  background-color: #f9f9f9;
  font-size: 14px;
}

li.重要 .type {
  color: red;
}

li.提醒 .type {
  color: orange;
}

li.通知 .type {
  color: green;
}

.type {
  font-weight: bold;
  margin-right: 10px;
}

.content {
  flex-grow: 1;
  text-overflow: ellipsis;
  white-space: nowrap;
  overflow: hidden;
}

.timestamp {
  font-size: 12px;
  color: #888;
  margin-left: 10px;
}
</style>
