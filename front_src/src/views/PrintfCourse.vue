<template>
    <div class="container">
      <table>
        <thead>
          <tr>
            <th colspan="9">{{ userInformations.name }}的课表</th>
          </tr>
          <tr>
            <th>时间段</th>
            <th>节次</th>
            <th>星期一</th>
            <th>星期二</th>
            <th>星期三</th>
            <th>星期四</th>
            <th>星期五</th>
            <th>星期六</th>
            <th>星期日</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(period, index) in periods" :key="index">
            <td v-if="index % 4 === 0" :rowspan="4">{{ period.time }}</td>
            <td>{{ period.session }}</td>
            <!-- 动态填充每一天的课程 -->
            <td v-for="(day, dayIndex) in 7" :key="dayIndex">
              <!-- 根据当前节次和星期几获取对应的课程信息 -->
              <span v-if="courseMap[index + 1] && courseMap[index + 1][dayIndex + 1]">
                 课程名称  ：{{ courseMap[index + 1][dayIndex + 1].course.course_name }} <br>
                 课程类别  ：{{ courseMap[index + 1][dayIndex + 1].course.course_type }} <br>
                 课程性质  ：{{ courseMap[index + 1][dayIndex + 1].course.course_nature }} <br>
                 教学班名称  ：{{ courseMap[index + 1][dayIndex + 1].course.class_name }}<br>
                 教师名称  ：{{ courseMap[index + 1][dayIndex + 1].course.teacher_name}}<br>
                 上课时间  ：{{ courseMap[index + 1][dayIndex + 1].course.class_time }}<br>
                 上课地点  ：{{ courseMap[index + 1][dayIndex + 1].course.class_address }}<br>
                学分  ：{{ courseMap[index + 1][dayIndex + 1].course.credits }}
              </span>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
</template>
  
<script>
import axios from 'axios';

export default {
data() {
    return {
    courses: [], // 存储从后端获取的课程信息
    periods: [
        { time: '上午', session: '1' },
        { time: '上午', session: '2' },
        { time: '上午', session: '3' },
        { time: '上午', session: '4' },
        { time: '下午', session: '5' },
        { time: '下午', session: '6' },
        { time: '下午', session: '7' },
        { time: '下午', session: '8' },
        { time: '晚上', session: '9' },
        { time: '晚上', session: '10' },
        { time: '晚上', session: '11' }
    ], // 预定义的时间段
    courseMap: {}, // 用于存储按节次和星期几分类的课程
    userInformations:{}
    };
},
mounted() {
    this.fetchCourses();
    this.fetchUserInformations();
},
methods: {
    async fetchCourses() {
        const account = this.$store.getters.getLoginData.account; // 当前登录用户的学号
        try {
            const response = await axios.get('http://localhost:8081/admin/GetUserCourseInfomation', {
            params: {
                account: account
            }
            });
            this.courses = response.data;
            this.parseCourses(); // 解析课程数据，生成课程时间表
        } catch (error) {
            console.error('获取课程数据失败:', error);
        }
    },
    async fetchUserInformations(){
        const account = this.$store.getters.getLoginData.account; // 当前登录用户的学号
        try {
            const response = await axios.get('http://localhost:8081/admin/GetUserInformation', {
            params: {
                account: account
            }
            });
            this.userInformations = response.data;
            console.log(response.data)
        } catch (error) {
            console.error('获取个人信息失败:', error);
        }
    },
    
    // 解析 class_time 字段并将课程数据映射到时间表中
    parseCourses() {
        this.courseMap = {}; // 重置课程表

        this.courses.forEach(course => {
            const { class_time } = course;
            const [session, day, _] = class_time.split(':'); // 解析 "1-2:星期一:1-16周" 格式
            const [startSession, endSession] = session.split('-').map(Number); // 获取起始和结束节次
            const dayIndex = this.getDayIndex(day); // 获取星期几对应的索引

            // 遍历节次并填充到相应的时间段和星期
            for (let i = startSession; i <= endSession; i++) {
            if (!this.courseMap[i]) {
                this.courseMap[i] = {}; // 初始化节次
            }
            this.courseMap[i][dayIndex] = {
                course
            };
            }
        });
    },
    
    // 将中文星期几转换为对应的数字索引 (1-7 表示星期一到星期日)
    getDayIndex(day) {
        switch (day) {
            case '星期一': return 1;
            case '星期二': return 2;
            case '星期三': return 3;
            case '星期四': return 4;
            case '星期五': return 5;
            case '星期六': return 6;
            case '星期日': return 7;
            default: return 0;
        }
    }
}
}
</script>

<style scoped>
/* 样式可以根据需要自定义 */
.container {
  margin: 0 auto;
  padding: 5%;
  background-color: #f8f9fa;
  border-radius: 10px;
  max-height: 80vh; /* 设置容器的最大高度 */
  overflow-y: auto; /* 启用垂直滚动条 */
}
table {
width: 100%;
border-collapse: collapse;
}
th, td {
border: 1px solid #ddd;
text-align: center;
padding: 8px;
}
</style>
