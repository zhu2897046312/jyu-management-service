<template>
    <div class="container">
        <header>
            <h3>学号: <span>{{ account }}</span></h3>
            <nav>
                <ul class="menu">
                    <li v-for="item in menuItems" :key="item">
                        <a href="#" @click.prevent="setActiveComponent(item)">{{ item }}</a>
                    </li>
                </ul>
            </nav>
        </header>

        <main>
            <!-- 根据 activeComponent 动态加载组件 -->
            <component :is="activeComponent"></component>
        </main>
    </div>
</template>

<script>
// 导入需要动态加载的组件
import BasicInfo from '../components/UserInformation/BasicInfo.vue'
import ContactInfo from '../components/UserInformation/ContactInfo.vue';
// import StudentAcademicInfo from '@/components/UserInformation/StudentAcademicInfo.vue';
// import OtherInfo from '@/components/UserInformation/OtherInfo.vue';
// import RewardInfo from '@/components/UserInformation/RewardInfo.vue';
// import GradesInfo from '@/components/UserInformation/GradesInfo.vue';
// import CourseInfo from '@/components/UserInformation/CourseInfo.vue';
// import TrainingPlan from '@/components/UserInformation/TrainingPlan.vue';
// import StatusChange from '@/components/UserInformation/StatusChange.vue';
// import AnnualEvaluation from '@/components/UserInformation/AnnualEvaluation.vue';

export default {
    data() {
        return {
            account: this.$store.getters.getLoginData.account,
            menuItems: [
                "基本信息", "学籍信息", "其他信息", "联系方式",
                "奖惩信息", "成绩信息", "选课信息", "培养方案",
                "学籍异动", "学年鉴定"
            ],
            activeComponent: 'BasicInfo', // 默认显示的组件
        };
    },
    methods: {
        setActiveComponent(menuItem) {
            switch(menuItem) {
                case "基本信息":
                    this.activeComponent = 'BasicInfo';
                    break;
                case "学籍信息":
                    this.activeComponent = 'StudentAcademicInfo';
                    break;
                case "其他信息":
                    this.activeComponent = 'OtherInfo';
                    break;
                case "联系方式":
                    this.activeComponent = 'ContactInfo';
                    break;
                case "奖惩信息":
                    this.activeComponent = 'RewardInfo';
                    break;
                case "成绩信息":
                    this.activeComponent = 'GradesInfo';
                    break;
                case "选课信息":
                    this.activeComponent = 'CourseInfo';
                    break;
                case "培养方案":
                    this.activeComponent = 'TrainingPlan';
                    break;
                case "学籍异动":
                    this.activeComponent = 'StatusChange';
                    break;
                case "学年鉴定":
                    this.activeComponent = 'AnnualEvaluation';
                    break;
                default:
                    this.activeComponent = 'StudentBasicInfo';
            }
        }
    },
    components: {
        BasicInfo,
        ContactInfo,
        // StudentAcademicInfo,
        // OtherInfo,
        // RewardInfo,
        // GradesInfo,
        // CourseInfo,
        // TrainingPlan,
        // StatusChange,
        // AnnualEvaluation
    }
}
</script>

<style>
.container {
    width: 80%;
    margin: auto;
}

.menu {
    display: flex;
    gap: 15px;
}

li {
    list-style-type: none;
    padding: 8px;
}

a {
    text-decoration: none;
    color: black;
    cursor: pointer;
}

a:hover {
    color: blue;
}
</style>
