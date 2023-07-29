<template>
  <a-layout style="min-height: 100vh">
    <a-layout-sider
      style="background-color: #293b5f"
      v-model:collapsed="collapsed"
      collapsible
    >
      <div class="logo">
        <img :src="logoUrl" style="width: 150px" />
      </div>
      <a-menu v-model:selectedKeys="selectedKeys" theme="dark" mode="inline">
        <a-menu-item key="1">
          <home-outlined />
          <span>Home</span>
        </a-menu-item>
        <a-menu-item key="sub1">
          <template #title>
            <span>
              <desktop-outlined />
              <span>Cluster</span>
            </span>
          </template>
          <a-menu-item key="3">MySQL</a-menu-item>
          <a-menu-item key="4">Redis</a-menu-item>
          <a-menu-item key="5">MongoDB</a-menu-item>
          <a-menu-item key="6">ClickHouse</a-menu-item>
        </a-menu-item>
        <a-sub-menu key="sub2">
          <template #title>
            <span>
              <partition-outlined />
              <span>Deploy</span>
            </span>
          </template>
          <a-menu-item key="7">Import</a-menu-item>
          <a-menu-item key="8">Add</a-menu-item>
        </a-sub-menu>
        <a-sub-menu key="sub3">
          <template #title>
            <span>
              <setting-outlined />
              <span>Settings</span>
            </span>
          </template>
          <a-menu-item key="6">Team 1</a-menu-item>
          <a-menu-item key="8">Team 2</a-menu-item>
        </a-sub-menu>
        <a-menu-item key="9">
          <file-outlined />
          <span>File</span>
        </a-menu-item>
      </a-menu>
    </a-layout-sider>
    <a-layout>
      <a-layout-header style="background: #3ffdcd; padding: 0" />
      <a-space size="1000"></a-space>
      <a-layout-content style="margin: 0 16px">
        <a-breadcrumb style="margin: 16px 0">
          <a-breadcrumb-item>User</a-breadcrumb-item>
          <a-breadcrumb-item>Bill</a-breadcrumb-item>
        </a-breadcrumb>
        <div
          :style="{ padding: '24px', background: '#fff', minHeight: '360px' }"
        >
          Bill is a cat.
        </div>
      </a-layout-content>
      <a-divider type="vertical" />
      <a-layout-footer style="text-align: center">
        <h5 style="text-align: center">Copyrights©2022-2023 Frabit LLC</h5>
      </a-layout-footer>
    </a-layout>
  </a-layout>
</template>
<script lang="ts" setup>
import {
  HomeOutlined,
  DesktopOutlined,
  PartitionOutlined,
  SettingOutlined,
  FileOutlined,
} from "@ant-design/icons-vue";
import { ref } from "vue";
const collapsed = ref<boolean>(false);
const selectedKeys = ref<string[]>(["1"]);
const logoUrl =
  localStorage.getItem("theme") === "light"
    ? new URL("../assets/logo.png", import.meta.url).href
    : new URL("../assets/logo.png", import.meta.url).href;
</script>

<style scoped>
#components-layout-demo-side .logo {
  height: 32px;
  margin: 16px;
  background: rgba(255, 255, 255, 0.3);
}

.site-layout .site-layout-background {
  background: #fff;
}
[data-theme="dark"] .site-layout .site-layout-background {
  background: #0000;
}
</style>
