<template>
    <section>
      <div class="dropdown" @click="toggleDropdown(1)">
        <button class="dropbtn">报名申请</button> 
        <div v-if="isDropdownVisible[1]" class="dropdown-content">
          <a href="#">Link 1</a>
          <a href="#">Link 2</a>
          <a href="#">Link 3</a>
        </div>
      </div>
      <div class="dropdown" @click="toggleDropdown(2)">
        <button class="dropbtn">信息维护</button>
        <div v-if="isDropdownVisible[2]" class="dropdown-content">
          <a href="#">Link 1</a>
          <a href="#">Link 2</a>
          <a href="#">Link 3</a>
        </div>
      </div>
      <div class="dropdown" @click="toggleDropdown(3)">
        <button class="dropbtn">选课</button>
        <div v-if="isDropdownVisible[3]" class="dropdown-content">
          <a href="#">Link 1</a>
          <a href="#">Link 2</a>
          <a href="#">Link 3</a>
        </div>
      </div>
      <div class="dropdown" @click="toggleDropdown(4)">
        <button class="dropbtn">信息查询</button>
        <div v-if="isDropdownVisible[4]" class="dropdown-content">
          <a href="#">Link 1</a>
          <a href="#">Link 2</a>
          <a href="#">Link 3</a>
        </div>
      </div>
      <div class="dropdown" @click="toggleDropdown(5)">
        <button class="dropbtn">教学评价</button>
        <div v-if="isDropdownVisible[5]" class="dropdown-content">
          <a href="#">Link 1</a>
          <a href="#">Link 2</a>
          <a href="#">Link 3</a>
        </div>
      </div>

      <ul>
        <li v-for="item in menuItems" :key="item.title" class="menu-item">
          <button @click="toggleDropdown(item.title)">
            {{ item.title }}
          </button>
          <ul v-if="activeDropdown === item.title" class="dropdown-menu">
            <li v-for="option in item.options" :key="option">{{ option }}</li>
          </ul>
        </li>
      </ul>
    </section>
  </template>
  
  <script>
  export default {
    data() {
      return {
        activeDropdown: null,
        isDropdownVisible: {
          1: false,
          2: false,
          3: false,
          4: false,
          5: false
        },
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
    methods: {
      toggleDropdown(id) {
        // 切换下拉菜单的显示状态
        this.isDropdownVisible[id] = !this.isDropdownVisible[id];
        // 关闭其他下拉菜单
        for (let key in this.isDropdownVisible) {
          if (parseInt(key) !== id) {
            this.isDropdownVisible[key] = false;
          }
        }
      },
      toggleDropdown(title) {
        this.activeDropdown = this.activeDropdown === title ? null : title;
      }

    }
  };
  </script>
  
<style scoped>
/* 基本样式 */
.dropdown {
    position: relative;
    display: inline-block;
    margin: 10px;
}

.dropdown-content {
    display: none;
    position: absolute;
    background-color: #f9f9f9;
    min-width: 160px;
    box-shadow: 0px 8px 16px 0px rgba(0,0,0,0.2);
    z-index: 1;
}

.dropdown-content a {
    color: black;
    padding: 12px 16px;
    text-decoration: none;
    display: block;
}

.dropdown-content a:hover {
    background-color: #f1f1f1;
}

.dropdown:hover .dropdown-content {
    display: block;
}

section {
    display: flex;
    flex-wrap: wrap;
    gap: 10px;
}
</style>
  