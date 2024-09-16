<template>
    <div class="common-layout">
      <el-container>
        <!-- 顶部导航栏 -->
        <el-header class="header">
          <div class="header-content">
            <div class="logo">管理系统</div>
            <div class="user-info">
              <el-avatar src="https://via.placeholder.com/50"></el-avatar>
              <span class="username">管理员</span>
              <el-button type="primary" size="mini" @click="logout">退出</el-button>
            </div>
          </div>
        </el-header>
  
        <!-- 主体部分 -->
        <el-container>
          <!-- 左侧导航栏 -->
          <el-aside width="200px" class="aside">
            <el-menu :default-active="activeMenu" class="el-menu-vertical" @select="handleMenuSelect">
              <el-menu-item index="1">
                <i class="el-icon-menu"></i>
                <span slot="title">课程管理</span>
              </el-menu-item>
              <el-sub-menu index="2">
                <template #title>
                    <el-icon><location /></el-icon>
                    <span>账户管理</span>
                </template>
                <el-menu-item index="2-1"> 基本信息管理</el-menu-item>
                <el-menu-item index="2-2">学籍信息管理</el-menu-item>
                <el-menu-item index="2-3">联系方式管理</el-menu-item>
              </el-sub-menu>
            </el-menu>
          </el-aside>
  
          <!-- 主体内容区 -->
          <el-main class="main">
            <router-view></router-view> <!-- 路由渲染内容 -->
          </el-main>
        </el-container>
      </el-container>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        activeMenu: '1', // 默认激活的菜单
      };
    },
    methods: {
      handleMenuSelect(index) {
        this.activeMenu = index;
        // 根据菜单项跳转到不同页面
        if (index === '1') {
          this.$router.push({ name: 'CourseManagement' });
        } else if (index === '2-1') {
          this.$router.push({ name: 'BasicInfoManagement' });
        } else if (index === '2-2') {
          this.$router.push({ name: 'StudentStatusManagement' });
        } else if (index === '2-3') {
          this.$router.push({ name: 'ContactManagement' });
        }
      },
      logout() {
        // 处理退出逻辑
        this.$router.push({ name: 'Login' });
      }
    }
  };
  </script>
  
  <style scoped>
  .common-layout {
    height: 100vh;
  }
  
  .header {
    background-color: #409EFF;
    color: white;
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 0 20px;
  }
  
  .header-content {
    display: flex;
    justify-content: space-between;
    width: 100%;
  }
  
  .logo {
    font-size: 20px;
    font-weight: bold;
  }
  
  .user-info {
    display: flex;
    align-items: center;
  }
  
  .username {
    margin-left: 10px;
    margin-right: 20px;
  }
  
  .aside {
    background-color: #2d3a4b;
    color: blue;
  }
  
  .el-menu-vertical {
    height: 100%;
  }
  
  .el-menu-item, .el-submenu__title {
    color: green;
  }
  
  .el-menu-item:hover, .el-submenu:hover {
    background-color: #1f2a36;
  }
  
  .main {
    padding: 20px;
    background-color: #f5f5f5;
    height: calc(100vh - 64px); /* 减去 header 的高度 */
    overflow-y: auto;
  }
  </style>
  