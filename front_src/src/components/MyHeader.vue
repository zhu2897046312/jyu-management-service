<template>
  <div @click="hideProfileMenu" class="app-container">
    <header class="app-header">
      <div class="logo">
        <img src="../assets/logo.png" alt="Logo" />
        <h1>嘉应学院教学综合管理信息服务平台</h1>
      </div>
      <div class="profile" @click.stop="toggleProfileMenu">
        <img src="../assets/logo.png" alt="User Avatar" class="avatar"/>
        <ul v-if="isProfileMenuVisible" class="profile-menu">
          <li>修改密码</li>
          <li @click="goToProfile">信息绑定</li>
          <li @click="logout">退出登录</li>
        </ul>
      </div>
    </header>

    <nav class="main-nav">
      <ul>
        <li v-for="item in menuItems" :key="item.title" class="menu-item">
          <button @click="toggleDropdown(item.title)">
            {{ item.title }}
            <i :class="['dropdown-icon', { 'icon-up': activeDropdown === item.title, 'icon-down': activeDropdown !== item.title }]"></i>
          </button>
          <ul v-if="activeDropdown === item.title" class="dropdown-menu">
            <li v-for="option in item.options" :key="option">{{ option }}</li>
          </ul>
        </li>
      </ul>
    </nav>
  </div>
</template>

<script>
export default {
  data() {
    return {
      activeDropdown: null,
      isProfileMenuVisible: false,
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
    toggleDropdown(title) {
      this.activeDropdown = this.activeDropdown === title ? null : title;
    },
    toggleProfileMenu() {
      this.isProfileMenuVisible = !this.isProfileMenuVisible;
    },
    hideProfileMenu() {
      this.isProfileMenuVisible = false;
    },
    goToProfile() {
      // Logic to navigate to profile page
      console.log('Navigate to profile');
    },
    logout() {
      // Clear user data (if any)
      // Perform any necessary cleanup actions

      // Redirect to login page
      this.$router.push({ path: '/login' });
    }
  },
  mounted() {
    document.addEventListener('click', this.handleClickOutside);
  },
  beforeDestroy() {
    document.removeEventListener('click', this.handleClickOutside);
  },
  methods: {
    handleClickOutside(event) {
      if (!this.$el.contains(event.target)) {
        this.hideProfileMenu();
      }
    },
    // Other existing methods
    toggleDropdown(title) {
      this.activeDropdown = this.activeDropdown === title ? null : title;
    },
    toggleProfileMenu() {
      this.isProfileMenuVisible = !this.isProfileMenuVisible;
    },
    hideProfileMenu() {
      this.isProfileMenuVisible = false;
    },
    goToProfile() {
      // Logic to navigate to profile page
      console.log('Navigate to profile');
    },
    logout() {
      // Clear user data (if any)
      // Perform any necessary cleanup actions

      // Redirect to login page
      this.$router.push({ path: '/login' });
    }
  }
};
</script>

<style scoped>
.app-container {
  position: relative;
}

.app-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  background: linear-gradient(90deg, #007bff, #0056b3);
  padding: 10px 20px;
  color: white;
}

.logo {
  display: flex;
  align-items: center;
}

.logo img {
  height: 50px;
  margin-right: 10px;
}

.profile {
  display: flex;
  align-items: center;
  position: relative;
}

.avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  cursor: pointer;
}

.profile-menu {
  list-style: none;
  padding: 0;
  margin: 10px 0 0 0;
  background-color: #00bfff;
  border: 1px solid #00bfff;
  border-radius: 4px;
  position: absolute;
  right: 0;
  top: 50px;
  width: 120px;
  box-shadow: 0 2px 5px rgba(0,0,0,0.1);
  z-index: 1000;
}

.profile-menu li {
  padding: 10px;
  text-align: left;
  cursor: pointer;
}

.profile-menu li:hover {
  background-color: #d4d4d4;
}

.main-nav {
  background-color: #f8f9fa;
  padding: 0 100px; /* 去掉内边距 */
  margin: 0; /* 去掉外边距 */
}

.main-nav ul {
  list-style: none;
  padding: 0;
  margin: 0;
  display: flex;
}

.menu-item {
  position: relative;
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

.dropdown-menu {
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

.dropdown-menu li {
  padding: 8px 12px;
  cursor: pointer;
  white-space: nowrap; /* 防止文本换行 */
}

.dropdown-menu li:hover {
  background-color: #f1f1f1;
}


</style>
