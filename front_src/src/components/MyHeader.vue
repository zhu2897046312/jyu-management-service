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
        { title: '报名申请', options: ['Option 1', 'Option 2'] },
        { title: '选课', options: ['Option 1', 'Option 2'] },
        { title: '信息维护', options: ['Option 1', 'Option 2'] },
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
  padding: 10px 20px;
}

.main-nav ul {
  list-style: none;
  padding: 0;
  margin: 0;
  display: flex;
}

.menu-item {
  position: relative;
  margin: 0 10px;
}

button {
  background-color: #007bff;
  color: white;
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
  background-color: #0056b3;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
}

.dropdown-icon {
  margin-left: 5px;
  width: 10px;
  height: 10px;
  display: inline-block;
  border: solid white;
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
  top: 100%;
  right: 0;
  background-color: #fff;
  border: 1px solid #ddd;
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.2);
  padding: 10px;
  border-radius: 5px;
  z-index: 1000;
}

.dropdown-menu li {
  padding: 8px 12px;
  cursor: pointer;
}

.dropdown-menu li:hover {
  background-color: #f1f1f1;
}
</style>
