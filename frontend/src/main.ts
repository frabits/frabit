import { createApp } from "vue";
import Antd from "ant-design-vue";
import { createPinia } from "pinia";

import i18n from "./locales";
import App from "./App.vue";
import router from "./router";

if (localStorage.getItem("theme") === "null") {
  import("@/style/dark.less");
} else {
  if (localStorage.getItem("theme") === "dark") {
    import("@/style/dark.less");
  } else {
    import("@/style/light.less");
  }
}

const app = createApp(App);

app.use(createPinia());
app.use(router);
app.use(i18n);
app.use(Antd);

app.mount("#app");
