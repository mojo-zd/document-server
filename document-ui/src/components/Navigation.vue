<template>
    <div class="navigation">
        <el-menu theme="dark" class="el-menu-demo" mode="horizontal">
        <el-menu-item index="1"><a class="none-underline" href="/">首页</a> </el-menu-item>
        <el-menu-item index="2">文档处理</el-menu-item>
        <el-submenu index="3">
            <template slot="title">我的工作台</template>
            <el-menu-item index="2-1">选项1</el-menu-item>
            <el-menu-item index="2-2">选项2</el-menu-item>
            <el-menu-item index="2-3">选项3</el-menu-item>
        </el-submenu>
        <el-menu-item index="4"><a class="none-underline" href="https://www.ele.me" target="_blank">订单管理</a></el-menu-item>
        </el-menu>
    </div>  
</template>
