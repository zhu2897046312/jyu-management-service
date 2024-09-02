<template>
  <div class="main-content">
    <div class="top-section">
      <div class="left">
        <el-card>
          <el-row>
            <el-col :span="12">
              <el-avatar :src="studentProfile.picture" size="large"></el-avatar>
              <div class="student-info">
                <h2>{{ studentProfile.name }}</h2>
                <p>{{ studentProfile.class }}</p>
              </div>
            </el-col>
            <el-col :span="12">
              <Notifications />
            </el-col>
          </el-row>
        </el-card>
        <Schedule />
      </div>
      <Calendar />
    </div>
    <div class="bottom-section">
      <Messages />
      <Grades />
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
    
      }
    }
  }
}
</script>

<style scoped>
.main-content {
  flex-grow: 1;
  padding: 20px;
  background-color: #fafafa;
}

.top-section {
  display: flex;
}

.left {
  flex: 2;
  margin-right: 20px;
}

.bottom-section {
  margin-top: 20px;
  display: flex;
  justify-content: space-between;
}
</style>
