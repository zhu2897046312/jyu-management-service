<template>
    <div class="container">
        <table>
            <thead>
                <th>学年</th>
                <th>学期</th>
                <th>课程名称</th>
                <th>开课学院</th>
                <th>课程类别</th>
                <th>学分</th>
                <th>教学班名称</th>
                <th>教师名称</th>
                <th>上课时间</th>
                <th>上课地点</th>
            </thead>
            <tbody>
                <tr v-for="course in courses" :key="course.id">
                    <td>{{ course.academic_year }}</td>
                    <td>{{ course.semester }}</td>
                    <td>{{ course.courseName }}</td>
                    <td>{{ course.commencement_academy }}</td>
                    <td>{{ course.course_type }}</td>
                    <td>{{ course.credits }}</td>
                    <td>{{ course.class_name }}</td>
                    <td>{{ course.teacher_name }}</td>
                    <td>{{ course.class_time }}-{{ course.endTime }}</td>
                    <td>{{ course.class_address }}</td>
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
            courses: []
        }
    },
    mounted() {
        this.fetchCourses();
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
            } catch (error) {
                console.error('获取课程数据失败:', error);
            }
        },
    }
}
</script>

<style>
.container {
    overflow: auto;
}
</style>