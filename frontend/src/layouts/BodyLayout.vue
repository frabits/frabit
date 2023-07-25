<template>
  <a-layout>
    <a-layout-header class="header">
      <div class="logo" />
      <a-menu
        v-model:selectedKeys="selectedKeys1"
        theme="dark"
        mode="horizontal"
        :style="{ lineHeight: '64px' }"
      >
        <a-menu-item key="1">nav 1</a-menu-item>
        <a-menu-item key="2">nav 2</a-menu-item>
        <a-menu-item key="3">nav 3</a-menu-item>
      </a-menu>
    </a-layout-header>
    <a-layout-content style="padding: 0 50px">
      <a-breadcrumb style="margin: 16px 0">
        <a-breadcrumb-item>Home</a-breadcrumb-item>
        <a-breadcrumb-item>List</a-breadcrumb-item>
        <a-breadcrumb-item>App</a-breadcrumb-item>
      </a-breadcrumb>
      <a-layout style="padding: 24px 0; background: #fff">
        <a-layout-sider width="200" style="background: #fff">
          <a-menu
            v-model:selectedKeys="selectedKeys2"
            v-model:openKeys="openKeys"
            mode="inline"
            style="height: 100%"
          >
            <a-sub-menu key="sub1">
              <template #title>
                <span>
                  <home-outlined />
                  Home
                </span>
              </template>
            </a-sub-menu>
            <a-sub-menu key="sub2">
              <template #title>
                <span>
                  <laptop-outlined />
                  Cluster
                </span>
              </template>
              <a-menu-item key="5">MySQL</a-menu-item>
              <a-menu-item key="6">Redis</a-menu-item>
              <a-menu-item key="7">MongoDB</a-menu-item>
              <a-menu-item key="8">ClickHouse</a-menu-item>
            </a-sub-menu>
            <a-sub-menu key="sub3">
              <template #title>
                <span>
                  <partition-outlined />
                  Deploy
                </span>
              </template>
              <a-menu-item key="9">Import</a-menu-item>
              <a-menu-item key="10">Add</a-menu-item>
            </a-sub-menu>
            <a-sub-menu key="sub4">
              <template #title>
                <span>
                  <setting-outlined />
                  Settings
                </span>
              </template>
              <a-menu-item key="9">Account</a-menu-item>
              <a-menu-item key="10">Hosts</a-menu-item>
              <a-menu-item key="11">Integrations</a-menu-item>
              <a-menu-item key="12">Licences</a-menu-item>
            </a-sub-menu>
            <a-sub-menu key="sub5">
              <template #title>
                <span>
                  <setting-outlined />
                  Help
                </span>
              </template>
              <a-menu-item key="9">Docs</a-menu-item>
              <a-menu-item key="10">Community</a-menu-item>
              <a-menu-item key="11">Feedback</a-menu-item>
              <a-menu-item key="12">Roadmap</a-menu-item>
            </a-sub-menu>
          </a-menu>
        </a-layout-sider>
        <a-layout-content :style="{ padding: '0 24px', minHeight: '280px' }">
          Content
        </a-layout-content>
      </a-layout>
    </a-layout-content>
    <a-layout-footer style="text-align: center">
      <h4 style="text-align: center">Frabit LLC ©2022-2023 All Rights Reserved</h4>
      <span style="text-align: center"
        >Frabit is a trademark of Frabit LLC, registered in CHN. and in other
        countries.
      </span>
    </a-layout-footer>
  </a-layout>
</template>
<script lang="ts" setup>
import { ref } from 'vue';
import {SettingOutlined, LaptopOutlined, PartitionOutlined,HomeOutlined } from '@ant-design/icons-vue';
const selectedKeys1 = ref<string[]>(["2"]);
const selectedKeys2 = ref<string[]>(["1"]);
const openKeys = ref<string[]>(["sub1"]);
</script>
<style scoped>
#components-layout-demo-top-side .logo {
  float: left;
  width: 120px;
  height: 31px;
  margin: 16px 24px 16px 0;
  background: rgba(255, 255, 255, 0.3);
}

.ant-row-rtl #components-layout-demo-top-side .logo {
  float: right;
  margin: 16px 0 16px 24px;
}

.site-layout-background {
  background: #fff;
}
</style>
