<template>
    <section>
      <ul class="dropdown">
        <li v-for="item in menuItems" :key="item.title" class="dropdown-menu-item">
          <button @click="toggleDropdown(item.title)">
            {{ item.title }}
            <i :class="['dropdown-icon', { 'icon-up': activeDropdown === item.title, 'icon-down': activeDropdown !== item.title }]"></i>
          </button>
          <ul v-if="activeDropdown === item.title" class="dropdown-menu">
            <li v-for="option in item.options" :key="option">
              <a href="#" @click.prevent="navigateTo(option)">{{ option }}</a>
            </li>
          </ul>
        </li>
      </ul>
    </section>
</template>

<script>
import { useRouter } from 'vue-router';

export default {
  data() {
    return {
      activeDropdown: null,
      menuItems: [
      { title: '报名申请', options: ["学籍异动申请", "学生申请交流项目",
                                    "辅修报名","重修报名",
                                    "学生证补办申请","学生转专业申请",
                                    "场地预约申请","考试项目报名",
                                    "教学项目报名","学生成绩学分认定申请",
                                    "大学英语免修免考报名","成绩作废申请",
                                    "补考确认","学科竞赛报名",
                                    "创新创业报名","学生退书申请",
                                    "开放性实验项目申请","学士学位申请",
                                    "学生实习申报"] },
      { title: '信息维护', options: ["学生监护人让人信息采集", "学生个人信息维护",
                                "主修专业确认","个人培养方案",
                                "创新创业学分确认","实习过程资料",
                                "学生自主报到注册","分流专业确认",
                                "专业方向确认"] },
      { title: '选课', options: ["个人课表查询", "自主选课",
                                  "筛选结果查询","教材预订",
                                  "实验选课","实验预约",
                                  "实习选xiang'm"] },
      { title: '信息查询', options: ['Option 1', 'Option 2'] },
      { title: '教学评价', options: ['Option 1', 'Option 2'] }
    ]
    };
  },
  props: {
    option: {
      type: String,
      required: true
    }
  },
  setup() {
    const router = useRouter();

    // 映射对象，将 option 值映射到路由名称
    const routeMap = {
      '自主选课': 'EnrollCourse',
      '其他': 'AnotherComponent'
      // 添加更多映射
    };

    const navigateTo = (option) => {
      const routeName = routeMap[option];
      if (routeName) {
        router.push({ name: routeName });
      } else {
        console.error(`无匹配的路由名称: ${option}`);
      }
    };

    return {
      navigateTo
    };
  },
  methods: {
    toggleDropdown(title) {
      this.activeDropdown = this.activeDropdown === title ? null : title;
    }

  }
};
</script>
  
<style scoped>
.dropdown {
  display: flex;
  list-style: none;
  padding: 0;
  margin: 0;
}
button {
  background-color: white;
  color: black;
  border: none;
  padding: 10px 20px;
  cursor: pointer;
  font-size: 14px;
  border-radius: 5px;
  transition: background-color 0.3s, box-shadow 0.3s;
  display: flex;
  align-items: center;
}

button:hover {
  background-color: #d3d3d3; /* 设置为浅灰色 */
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
}


.dropdown-icon {
  margin-left: 5px;
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

.dropdown-menu-item {
  position: relative;
  margin: 0;
}
.dropdown-menu {
  list-style-type: none; /* 去掉小圆点 */
  position: absolute;
  top: 100%; /* 使下拉菜单出现在按钮的正下方 */
  left: 0; /* 将下拉菜单的左边缘与按钮的左边缘对齐 */
  background-color: #fff;
  border: 1px solid #ddd;
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.2);
  padding: 10px;
  border-radius: 5px;
  z-index: 1000;
  display: flex; /* 使用 flexbox 布局 */
  flex-direction: column; /* 垂直排列子项 */
  min-width: 150px; /* 设置最小宽度 */
  /* 如果需要，可以设置固定宽度 */
}

section {
    display: flex;
    flex-wrap: wrap;
    gap: 10px;
    padding: 0px 50px;
}
</style>
  