import cluster from "./zh_cn/cluster";
import common from "./zh_cn/common";
import database from "./zh_cn/database";
import instance from "./zh_cn/instance";
import permission from "./zh_cn/permission";
import project from "./zh_cn/project";
import role from "./zh_cn/role";
import settings from "./zh_cn/settings";
import user from "./zh_cn/user";
import workspace from "./zh_cn/workspace";

import antdZhCN from "ant-design-vue/es/locale/zh_CN";

const components = {
  antLocale: antdZhCN,
  DayJsName: "zh-cn",
};

export default {
  ...components,
  ...cluster,
  ...common,
  ...database,
  ...instance,
  ...permission,
  ...project,
  ...role,
  ...settings,
  ...user,
  ...workspace,
};
