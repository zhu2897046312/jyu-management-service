<template>
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
</template>

<script>
export default {
  data() {
    return {
      isProfileMenuVisible: false,
    };
  },
  methods: {
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
.app-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  background: linear-gradient(90deg, #007bff, #0056b3);
  padding:0px 50px;
  color: white;
  height: 50px;
}

.logo {
  display: flex;
  align-items: center;
}

.logo img {
  height: 40px;
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

</style>
