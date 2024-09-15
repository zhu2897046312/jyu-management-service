<template>
    <div>
      <h1>学生成绩总表</h1>
      <button @click="downloadExcel">下载 Excel</button>
    </div>
</template>

<script>
export default {
methods: {
    // 下载 Excel 文档
    downloadExcel() {
    // 发送 GET 请求以生成 Excel
    this.downloadFile('http://localhost:8081/generate_excel', 'student_grades.xlsx');
    },
    // 通用文件下载方法
    downloadFile(url, filename) {
    // 创建一个隐藏的链接
    const link = document.createElement('a');
    link.href = url;
    link.download = filename;

    // 模拟点击链接以触发下载
    document.body.appendChild(link);
    link.click();

    // 移除链接元素
    document.body.removeChild(link);
    }
}
}
</script>

<style scoped>
h1 {
margin-bottom: 20px;
}

button {
margin-right: 10px;
padding: 10px 20px;
font-size: 16px;
cursor: pointer;
}
</style>
  