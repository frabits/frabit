<template>
  <a-config-provider :locale="lang">
    <RouterView />
  </a-config-provider>
</template>

<script setup lang="ts">
import zhCN from "ant-design-vue/es/locale/zh_CN";
import enUS from "ant-design-vue/es/locale/en_US";
import { locale } from "dayjs";
import { defaultLang } from "@/locales";

locale(defaultLang);
const lang = defaultLang === "en_US" ? enUS : zhCN;
</script>
