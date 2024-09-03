<template>
  <div class="calendar-container">
    <div class="calendar-header">
      <span class="term-info">{{ semesterInfo }}</span>
    </div>
    <div class="calendar-body">
      <table class="calendar-table">
        <thead>
          <tr>
            <th v-for="day in weekdays" :key="day" class="calendar-weekday">{{ day }}</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(week, index) in calendarWeeks" :key="index">
            <td v-for="date in week" :key="date" :class="['calendar-date', date.class]">
              {{ date.day }}
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      semesterInfo: '2024-2025学年1学期(2024-09-01至2025-01-18)',
      weekdays: ['日', '一', '二', '三', '四', '五', '六'],
      year: 2024,
      month: 9, // 九月
      calendarWeeks: [] // 用于存放每周的数据
    };
  },
  mounted() {
    this.generateCalendar(this.year, this.month);
  },
  methods: {
    generateCalendar(year, month) {
      const daysInMonth = new Date(year, month, 0).getDate(); // 当月总天数
      const firstDay = new Date(year, month - 1, 1).getDay(); // 当月1号是星期几
      
      let days = [];
      let week = [];

      // 填充空白天数（上个月）
      for (let i = 0; i < firstDay; i++) {
        week.push({ day: '', class: 'empty' });
      }

      // 填充当月天数
      for (let i = 1; i <= daysInMonth; i++) {
        const dateClass = this.getDateClass(i);
        week.push({ day: i, class: dateClass });

        // 每周结束，将本周加入 weeks
        if (week.length === 7) {
          days.push(week);
          week = [];
        }
      }

      // 如果最后一周没有填满，补上空白
      while (week.length < 7) {
        week.push({ day: '', class: 'empty' });
      }
      days.push(week);

      this.calendarWeeks = days;
    },
    getDateClass(day) {
      const today = new Date();
      const isToday = day === today.getDate() && this.month === today.getMonth() + 1 && this.year === today.getFullYear();
      const isWeekend = [0, 6].includes(new Date(this.year, this.month - 1, day).getDay());

      if (isToday) return 'today';
      if (isWeekend) return 'weekend';
      return '';
    }
  }
};
</script>

<style scoped>
.calendar-container {
  width: 100%;
  background-color: #f9f9f9;
  border-radius: 4px;
}

.calendar-header {
  padding: 10px;
  text-align: center;
  background-color: #fff;
  border-bottom: 1px solid #ddd;
}

.term-info {
  font-weight: bold;
  font-size: 14px;
}

.calendar-body {
  padding: 15px;
}

.calendar-table {
  width: 100%;
  border-collapse: collapse;
  text-align: center;
}

.calendar-weekday {
  background-color: #eaeaea;
  font-weight: bold;
  padding: 8px 0;
  border-bottom: 1px solid #ddd;
}

.calendar-date {
  padding: 10px 0;
  border-bottom: 1px solid #ddd;
}

.empty {
  background-color: #f4f4f4;
}

.today {
  background-color: #3a86ff;
  color: #fff;
}

.weekend {
  color: red;
}

</style>
