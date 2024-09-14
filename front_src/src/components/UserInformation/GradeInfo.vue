<template>
    <div class="container">
        <table>
            <thead>
                <th>学年</th>
                <th>学期</th>
                <th>课程代码</th>
                <th>课程名称</th>
                <th>课程性质</th>
                <th>考试性质</th>
                <th>学分</th>
                <th>成绩</th>
                <th>绩点</th>
                <th>成绩备注</th>
            </thead>
            <tbody>
                <tr v-for="grade in GradeInformation" :key="grade.id">
                    <td>{{ grade.academic_year }}</td>
                    <td>{{ grade.semester }}</td>
                    <td>{{ grade.course_code }}</td>
                    <td>{{ grade.course_name }}</td>
                    <td>{{ grade.course_nature }}</td>
                    <td>{{ grade.exam_nature}}</td>
                    <td>{{ grade.credits }}</td>
                    <td>{{ grade.grade }}</td>
                    <td>{{ grade.grade_points}}</td>
                    <!-- <td>{{ grade.grade_remark }}</td> -->
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
            GradeInformation: []
        }
    },
    mounted() {
        this.fetchGradeInformation();
    },
    methods: {
        async fetchGradeInformation() {
            const account = this.$store.getters.getLoginData.account; // 当前登录用户的学号
            try {
                const response = await axios.get('http://localhost:8081/admin/GetGradeInformationHandler', {
                params: {
                    account: account
                }
                });
                this.GradeInformation = response.data;
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