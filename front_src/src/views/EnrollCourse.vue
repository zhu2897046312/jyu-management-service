<template>
  <div class="container">
    <div>
      <input type="text">
      <button>查询</button>
      <button>重置</button>
    </div>

    <div class="option">

    </div>
    <div class="search">
      <ul>
        <li>
          年级：
          <a href="">2024</a>
          <a href="">2023</a>
          <a href="">2022</a>
          <a href="">2021</a>
          <a href="">2020</a>
          <a href="">2019</a>
          <a href="">2018</a>
          <a href="">2017</a>
          <a href="">2016</a>
          <a href="">2015</a>
        </li>
        <li>
          学院：
          <a href="">信息学院</a>
          <a href="">管理学院</a>
          <a href="">人文学院</a>
          <a href="">艺术学院</a>
          <a href="">社会学院</a>
          <a href="">经济学院</a>
          <a href="">法学院</a>
          <a href="">外国语学院</a>
          <a href="">国际 relations学院</a>
          <a href="">新闻学院</a>
        </li>
        <li>
          专业：
          <a href="">信息管理</a>
          <a href="">人工智能</a>
          <a href="">数据科学</a>
          <a href="">网络与信息安全</a>
          <a href="">数字媒体</a>
          <a href="">应用心理学</a>
          <a href="">国际商务</a>
          <a href="">金学</a>
          <a href="">外国语</a>
          <a href="">日语</a>
        </li>
        <li>
          开课学院：
          <a href="">信息学院</a>
          <a href="">管理学院</a>
          <a href="">人文学院</a>
          <a href="">艺术学院</a>
          <a href="">社会学院</a>
          <a href="">经济学院</a>
          <a href="">法学院</a>
          <a href="">外国语学院</a>
          <a href="">国际 relations学院</a>
          <a href="">新闻学院</a>
        </li>
        <li>
          课程类别：
          <a href="">必修</a>
          <a href="">选修</a>
          <a href="">任选</a>
        </li>
        <li>
          课程性质：
          <a href="">公选</a>
          <a href="">专选</a>
          <a href="">通识</a>
        </li>
        <li>
          课程归属：
          <a href="">基础课程</a>
          <a href="">专业课程</a>
          <a href="">选修课程</a>
          <a href="">任选课程</a>
        </li>
        <li>
          教学模式：
          <a href="">自主</a>
          <a href="">集中</a>
          <a href="">面授</a>
        </li>
        <li>
          上课星期：
          <a href="">星期一</a>
          <a href="">星期二</a>
          <a href="">星期三</a>
          <a href="">星期四</a>
          <a href="">星期五</a>
          <a href="">星期六</a>
          <a href="">星期日</a>
        </li>
        <li>
          上课节次：
          <a href="">1-2节</a>
          <a href="">3-4节</a>
          <a href="">5-6节</a>
          <a href="">7-8节</a>
          <a href="">9-10节</a>
          <a href="">11-12节</a>
        </li>
        <li>
          是否重修：
          <a href="">是</a>
          <a href="">否</a>
        </li>
        <li>
          是否有余量：
          <a href="">是</a>
          <a href="">否</a>
        </li>
      </ul>
    </div>
    
    <!-- 选课列表 -->
    <ul class="enroll">
      <li v-for="course in courses" :key="course.course_code" class="course-item">
        <h7 class="course-header">
          <span class="course-code">{{ course.course_code }}</span>
          <a href=""  class="course-name">{{ course.course_name }}</a>
          <span class="course-credits">{{ course.credits }} 学分</span>
          <span class="course-status">选课状态：{{ getCourseStatus(course.course_code) }}</span>
          <button @click="toggleDropdown(course.course_code)" class="toggle-btn">
            <i :class="['dropdown-icon', { 'icon-up': activeDropdown === course.course_code, 'icon-down': activeDropdown !== course.course_code }]"></i>
          </button>
        </h7>

        <!-- 展示课程详细信息的区域 -->
        <transition name="fade">
          <div v-if="activeDropdown === course.course_code" class="course-details">
            <table class="course-table">
              <thead>
                <tr>
                  <th>教学班</th>
                  <th>上课教师</th>
                  <th>上课时间</th>
                  <th>教学地点</th>
                  <th>开课学院</th>
                  <th>课程类别</th>
                  <th>课程性质</th>
                  <th>教学模式</th>
                  <th>已选/容量</th>
                  <th>操作</th>
                </tr>
              </thead>
              <tbody>
                <tr>
                  <td>{{ course.class_name }}</td>
                  <td>{{ course.teacher_name }}</td>
                  <td>{{ course.class_time }}</td>
                  <td>{{ course.class_address }}</td>
                  <td>{{ course.commencement_academy }}</td>
                  <td>{{ course.course_type }}</td>
                  <td>{{ course.course_nature }}</td>
                  <td>{{ course.teaching_mode }}</td>
                  <td>{{ course.choosed_number }} / {{ course.max_student_number }}</td>
                  <td>
                    <button class="action-btn" @click="selectCourse(course.course_code)">选课</button>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </transition>
      </li>
    </ul>
  </div>
</template>

<script>
import axios from 'axios';
export default {
  data() {
    return {
      activeDropdown: null,
      status: '未选',  // 初始状态
      courses: [],
      courseStatuses: [
        {course_code:"cs101",status: 1},
        {course_code:"CS103",status: 1},
        {course_code:"CS104",status: 1},
        {course_code:"CS105",status: 1},
        {course_code:"CS106",status: 1},
      ], // 课程状态
      menuItems:[
        {title:"年级：" ,option:[]},
        {title:"学院：" ,option:[]},
        {title:"专业：" ,option:[]},
        {title:"开课学院：" ,option:[]},
        {title:"课程类别：" ,option:[]},
        {title:"课程性质：" ,option:[]},
        {title:"课程归属：" ,option:[]},
        {title:"教学模式：" ,option:[]},
        {title:"上课星期：" ,option:[]},
        {title:"上课节次：" ,option:[]},
        {title:"是否重修：" ,option:[]},
        {title:"是否有余量：" ,option:[]},
      ]
    };
  },
  mounted() {
    this.fetchCourses();
    console.log(this.$store); // 打印 Vuex store 对象
    console.log(this.$store.getters); // 打印所有的 getters
    console.log(this.$store.getters.getLoginData); // 打印 loginData getter
    console.log(this.$store.getters.getLoginData.account); // 打印 loginData getter
  },
  methods: {
    async fetchCourses() {
      try {
        const response = await axios.get('http://localhost:8081/admin/GetAll');
        this.courses = response.data;
      } catch (error) {
        console.error('获取课程数据失败:', error);
      }
    },
    toggleDropdown(courseCode) {
      this.activeDropdown = this.activeDropdown === courseCode ? null : courseCode;
    },
    fetchStatus() {
      const userCourseData = {
          account: this.$store.getters.getLoginData.account,  // 当前登录用户的学号
        };

        axios({
          method: 'POST',
          url: 'http://localhost:8081/admin/GetAllByAccount',  
          data: userCourseData,
          headers: {
            'Content-Type': 'application/json'
          }
        })
        .then(res => {
          console.log('成功', res.data);

        })
        .catch(err => {
          console.error('失败', err);
        });
    },
    selectCourse(courseCode) {
      const userCourseData = {
        account: this.$store.getters.getLoginData.account,  // 当前登录用户的学号
        course_code: courseCode,  // 选课的课程代码
        status: 1
      };

      axios({
        method: 'POST',
        url: 'http://localhost:8081/admin/EnrollCourse',  
        data: userCourseData,
        headers: {
          'Content-Type': 'application/json'
        }
      })
      .then(res => {
        console.log('选课成功', res.data);

      })
      .catch(err => {
        console.error('选课失败', err);
      });
    },
    getCourseStatus(courseCode) {
      const courseStatus = this.courseStatuses.find(item => item.course_code === courseCode);
      
      if (courseStatus) {
        // 如果找到对应的 course_code，判断 status
        return courseStatus.status === 1 ? '已选' : '未选';
      } else {
        // 如果未找到对应的 course_code，返回默认值 '未选'
        return '未选';
      }
    }
  },
};
</script>

<style scoped>
/* 整体布局 */
.container {
  margin: 0 auto;
  padding: 5%;
  background-color: #f8f9fa;
  border-radius: 10px;
  max-height: 80vh; /* 设置容器的最大高度 */
  overflow-y: auto; /* 启用垂直滚动条 */
}

/* 选课列表样式 */
.enroll {
  list-style: none;
  padding: 0;
}

.course-item {
  margin-bottom: 20px;
  padding: 15px;
  background-color: #fff;
  border-radius: 8px;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
}

/* 课程标题样式 */
.course-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-size: 18px;
  font-weight: bold;
}

.course-header span,
.course-header a {
  margin-right: 10px;
}

/* 展开按钮样式 */
.toggle-btn {
  background-color: #007bff; /* 确保按钮有背景颜色 */
  color: #fff;
  border: none;
  padding: 8px 10px; /* 增加按钮的填充 */
  cursor: pointer;
  border-radius: 4px;
  outline: none;
}

.dropdown-icon {
  margin: 0;
  width: 10px;
  height: 10px;
  display: inline-block;
  border: solid black;
  border-width: 0 2px 2px 0;
  content: "";
  transform: rotate(45deg);
  transition: transform 0.3s;
}

.icon-down {
  transform: rotate(45deg);
}

.icon-up {
  transform: rotate(-135deg);
}

/* 课程详情样式 */
.course-details {
  margin-top: 10px;
  overflow: hidden;
  transition: max-height 0.1s ease-in-out;
}

/* 表格样式 */
.course-table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 10px;
}

.course-table th, .course-table td {
  padding: 12px 15px;
  border: 1px solid #ddd;
  text-align: left;
}

.course-table th {
  background-color: #f2f2f2;
  font-weight: bold;
}

.course-table tr:nth-child(even) {
  background-color: #f9f9f9;
}

/* 动画效果 */
.fade-enter-active, .fade-leave-active {
  transition: opacity 0.5s ease;
}
.fade-enter, .fade-leave-to {
  opacity: 0;
}

/* 操作按钮 */
.action-btn {
  padding: 8px 15px;
  background-color: #007bff;
  color: #fff;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.action-btn:hover {
  background-color: #0056b3;
}
</style>