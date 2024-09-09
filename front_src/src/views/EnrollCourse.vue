<template>
  <div class="container">
    <div>
      <input type="text">
      <button @click="performSearch">查询</button>
      <button @click="resetFilters">重置</button>
    </div>

    <div class="option">
      <span v-for="(tag, index) in selectedTags" :key="index" class="selected-tag">
        {{ tag }}
        <button @click="removeTag(index)">x</button> <!-- 删除标签按钮 -->
      </span>
    </div>
    <div class="search">
      <ul>
        <li v-for="(item, index) in categories" :key="index">
          {{ item.title }}：
          <div class="options-container">
            <a v-for="(option, idx) in item.options" :key="idx" href="#" @click.prevent="selectTag(item.title,option)" class="option">
              {{ option }}
            </a>
        </div>
        </li>
      </ul>
    </div>
    
    <!-- 选课列表 -->
    <ul class="enroll">
      <li v-for="course in courses" :key="course.course_code" class="course-item">
        <div class="course-header">
          <span class="course-code">{{ course.course_code }}</span>
          <a href="#" @click.prevent="selectTag('2024')"  class="course-name">{{ course.course_name }}</a>
          <span class="course-credits">{{ course.credits }} 学分</span>
          <span class="course-status">选课状态：{{ getCourseStatus(course.course_code) }}</span>
          <button @click="toggleDropdown(course.course_code)" class="toggle-btn">
            <i :class="['dropdown-icon', { 'icon-up': activeDropdown === course.course_code, 'icon-down': activeDropdown !== course.course_code }]"></i>
          </button>
        </div>

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
      selectedTags: [], // 保存已选标签
      courses: [],
      courseStatuses: [
        {course_code:"cs101",status: 1},
        {course_code:"CS103",status: 1},
        {course_code:"CS104",status: 1},
        {course_code:"CS105",status: 1},
        {course_code:"CS106",status: 1},
      ], // 课程状态
      categories: [
        {
          title: '年级',
          options: [
            '2024', '2023', 
            '2022', '2021', 
            '2020', '2019', 
            '2018', '2017', 
            '2016', '2015'
          ]
        },
        {
          title: '学院',
          options: [
            '信息学院', '管理学院', 
            '人文学院', '艺术学院', 
            '社会学院', '经济学院', 
            '法学院', '外国语学院', 
            '国际学院', '新闻学院'
          ]
        },
        {
          title: '专业',
          options: [
            '信息管理', '人工智能', 
            '数据科学', '网络与信息安全', 
            '数字媒体', '应用心理学', 
            '国际商务', '金学', 
            '外国语', '日语'
          ]
        },
        {
          title: '开课学院',
          options: [
            '信息学院', '管理学院', 
            '人文学院', '艺术学院', 
            '社会学院', '经济学院', 
            '法学院', '外国语学院', 
            '国际学院', '新闻学院'
          ]
        },
        {
          title: '课程类别',
          options: [
            '实践课', '通识任选课', 
            '通识限选课', '国防安全教育课', 
            '思想政治理论课', '语言与技能课',
            '计算机应用技术课', '健康与运动课',
            '学科基础课', '专业基础课',
          ]
        },
        {
          title: '课程性质',
          options: ['公选', '专选', '通识']
        },
        {
          title: '课程归属',
          options: [
            '人文社会科学', '自然科学与技术', 
            '艺术与审美', '教师教育',
            '客家文化','创新创业'
          ]
        },
        {
          title: '教学模式',
          options: ['线上', '线下',]
        },
        {
          title: '上课星期',
          options: [
            '星期一', '星期二', 
            '星期三', '星期四', 
            '星期五', '星期六', 
            '星期日'
          ]
        },
        // {
        //   title: '上课节次',
        //   options: [
        //     '1-2节', '3-4节', 
        //     '5-6节', '7-8节', 
        //     '9-10节', '11-12节'
        //   ]
        // },
        {
          title: '是否重修',
          options: ['是', '否']
        },
        {
          title: '是否有余量',
          options: ['是', '否']
        }
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
    // 选择标签并添加到 selectedTags 中
    selectTag(key, value) {
      const tag = `${key}:${value}`; // 使用下划线作为分隔符
      if (!this.selectedTags.includes(tag)) {
        this.selectedTags.push(tag);
      }
    },
    // 删除标签
    removeTag(index) {
      this.selectedTags.splice(index, 1);
    },
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
    },
    performSearch(){
      const conditions = {};
      
      if (!this.selectedTags){
        fetchCourses();
      }else{
        this.selectedTags.forEach(tag => {
          const [key, value] = tag.split(':');
          switch (key) {
            case '年级':
              conditions['academic_year'] = value;
              break;
            case '开课学院':
              conditions['commencement_academy'] = value;
              break;
            case '课程类别':
              conditions['course_type'] = this.convertCourseType(value);
              break;
            case '课程性质':
              conditions['course_nature'] = this.convertCourseNature(value);
              break;
            case '课程归属':
              conditions['course_affiliation'] = this.convertCourseNature(value);
              break;
            case '教学模式':
              conditions['teaching_mode'] = this.convertTeachingMode(value);
              break;
            default:
              break;
          }
        });
      console.log(conditions);
      axios({
        method: 'POST',
        url: 'http://localhost:8081/admin/courses',  
        data: conditions,
        headers: {
          'Content-Type': 'application/json'
        }
      })
      .then(res => {
        this.courses = res.data;
        console.log('查询结果:', res.data);
      })
      .catch(err => {
        console.error('查询出错:', err);
      });
      }
    },
    resetFilters() {
      this.searchQuery = '';
      this.selectedTags = [];
    },
    convertCourseType(value) {
      const types = { 
        '实践课': 1, '通识任选课': 2, 
        '通识限选课': 3, '国防安全教育课': 4, 
        '思想政治理论课': 5, '语言与技能课': 6,
        '计算机应用技术课': 7, '健康与运动课': 8,
        '学科基础课': 9, '专业基础课': 10,
      };
      return types[value] || null;
    },
    convertCourseNature(value) {
      const natures = { '通识必修': 1, '通识选修': 2, 
        '专业必修': 3, '专业选修': 4, 
        '职业必修': 5, '职业选修': 6
      };
      return natures[value] || null;
    },
    convertTeachingMode(value) {
      const modes = { '线上': 1, '线下': 2,};
      return modes[value] || null;
    },
    convertCourseAffiliation(value) {
      const affiliations = { 
        '人文社会科学': 1, '自然科学与技术': 2, 
        '艺术与审美': 3, '教师教育': 4,
        '客家文化': 5,'创新创业': 6
      };
      return affiliations[value] || null;
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

.search {
  background-color: #f9f9f9;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.search ul li {
  list-style-type: none;
  margin-bottom: 15px;
  font-size: 16px;
}

.options-container {
  flex: 1;
  display: flex;
  flex-wrap: wrap;
}

.option {
  display: inline-block;
  margin-right: 10px;
  margin-bottom: 5px;
  padding: 5px 10px;
  background-color: #e0e0e0;
  border-radius: 4px;
  text-decoration: none;
  color: #333;
  white-space: nowrap; /* 防止单个选项内换行 */
}

.options-container a:nth-child(n+2) {
  text-indent: 10px; /* 溢出部分与第一个选项对齐 */
}

.option:hover {
  background-color: #007bff;
  color: #fff;
}

.option:active {
  background-color: #0056b3;
  color: #fff;
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