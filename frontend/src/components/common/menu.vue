<template>
  <a-menu v-model:selectedKeys="selectedKey" mode="inline" @click="changeMenu">
    <a-menu-item key="/home">
      <HomeOutlined />
      <span> {{ $t("menu.home") }}</span>
    </a-menu-item>
    <a-menu-item key="/home">
      <HomeOutlined />
      <span> {{ $t("menu.home") }}</span>
    </a-menu-item>
    <a-menu-item key="/home">
      <HomeOutlined />
      <span> {{ $t("menu.home") }}</span>
    </a-menu-item>
    <a-menu-item key="/home">
      <HomeOutlined />
      <span> {{ $t("menu.home") }}</span>
    </a-menu-item>
    <a-menu-item key="/home">
      <HomeOutlined />
      <span> {{ $t("menu.home") }}</span>
    </a-menu-item>
    <a-sub-menu title="$t('menu.deploy')">
      <template #icon>
        <AreaChartOutlined />
      </template>
      <a-menu-item key="/server/deploy/mysql">
        <ToolOutlined />
        <span>{{ $t("menu.deploy.mysql") }}</span>
      </a-menu-item>
    </a-sub-menu>
  </a-menu>
</template>

<script lang="ts" setup>
import { ref } from "vue";
import router from "@/router";
import { AreaChartOutlined, ToolOutlined } from "@ant-design/icons-vue";

const emit = defineEmits(["close"]);

const selectedKey = ref("");
const changeMenu = (vl: { keyPath: string[]; key: string }) => {};
</script>
