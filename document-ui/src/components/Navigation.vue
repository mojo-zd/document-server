<template>
    <div class="navigation">
        <el-menu theme="dark" class="el-menu-demo" mode="horizontal">
        <el-menu-item index="1">
            <router-link :to="{ name: 'dashboard'}" class="none-underline">首页</router-link>
        </el-menu-item>
        <el-menu-item index="2">
            <router-link :to="{ name: 'documents'}" class="none-underline">文档处理</router-link>
        </el-menu-item>
        <el-menu-item index="4"><a class="none-underline" href="https://www.ele.me" target="_blank">订单管理</a></el-menu-item>
        </el-menu>
    </div>  
</template>
