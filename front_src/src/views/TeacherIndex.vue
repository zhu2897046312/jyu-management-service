<template>
    <div class="common-layout">
        <el-container>
            <el-header class="header">
                <h2>教师课程管理</h2>
                <div class="profile" @click.stop="toggleProfileMenu">
                    <img src="../assets/logo.png" alt="User Avatar" class="avatar"/>
                    <ul v-if="isProfileMenuVisible" class="profile-menu">
                        <li>修改密码</li>
                        <li @click="goToProfile">信息绑定</li>
                        <li @click="logout">退出登录</li>
                    </ul>
                </div>
            </el-header>
            <el-main class="main">
                <el-upload
                    class="upload-demo"
                    drag
                    action="http://localhost:8081/teacher/uploadExcel"
                    :on-success="handleUploadSuccess"
                    :show-file-list="false"
                    >
                    <i class="el-icon-upload"></i>
                    <div class="el-upload__text">拖拽文件或点击上传</div>
                    <div class="el-upload__text">填写下载的文件内容即可上传解析</div>
                    <div class="el-upload__tip" slot="tip">上传Excel文件</div>
                </el-upload>

                <el-table :data="courseInfo" @row-click="editStudentScores">
                    <el-table-column prop="course_name" label="课程名称"></el-table-column>
                    <el-table-column prop="class_name" label="教学班级"></el-table-column>
                    <el-table-column label="操作">
                        <template #default="scope">
                            <el-button @click="downloadStudentInfo(scope.row.course_code)" size="small">下载学生信息</el-button>
                        </template>
                    </el-table-column>
                </el-table>
            </el-main>
        </el-container>
    </div>
</template>

<script>
import axios from 'axios';

export default {
    data() {
        return {
            isProfileMenuVisible: false,
            courseInfo: [],
            isUploadDialogVisible: false,
            selectedRow: null,
            file: null,
            fileName: '',
        };
    },
    mounted() {
        this.fetchCourses();
    },
    methods: {
        toggleProfileMenu() {
            this.isProfileMenuVisible = !this.isProfileMenuVisible;
        },
        goToProfile() {
            console.log('Navigate to profile');
        },
        logout() {
            this.$router.push({ path: '/login' });
        },
        async fetchCourses() {
            const account = this.$store.getters.getLoginData.account;
            try {
                const response = await axios.get('http://localhost:8081/teacher/GetTeacherCourses', {
                    params: { account }
                });
                this.courseInfo = response.data.Courses || [];
                console.log('获取课程数据成功:', this.courseInfo);
            } catch (error) {
                console.error('获取课程数据失败:', error);
            }
        },
        editStudentScores(row) {
            console.log('编辑学生成绩:', row);
        },
        downloadStudentInfo(courseCode) {
            console.log('下载学生信息，课程代码:', courseCode);
            axios.get(`http://localhost:8081/teacher/grades_execl`, {
                params: { course_code: courseCode },
                responseType: 'blob'
            }).then(response => {
                const url = window.URL.createObjectURL(new Blob([response.data]));
                const link = document.createElement('a');
                link.href = url;
                link.setAttribute('download', `${courseCode}_学生信息.xlsx`);
                document.body.appendChild(link);
                link.click();
            }).catch(error => {
                console.error('下载失败:', error);
            });
        },
        handleUploadSuccess(response) {
            // 显示成功消息提示框
            this.$message.success('上传成功！');
        },
    }
};
</script>

<style scoped>
.header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    background-color: #3f51b5;
    color: white;
    padding: 10px 20px;
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

.main {
    padding: 20px;
    background-color: #f5f5f5;
    height: calc(100vh - 64px);
    overflow-y: auto;
}
</style>
