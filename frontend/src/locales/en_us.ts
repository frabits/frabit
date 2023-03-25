import cluster from "./en_us/cluster";
import common from "./en_us/common";
import database from "./en_us/database";
import instance from "./en_us/instance";
import permission from "./en_us/permission";
import project from "./en_us/project";
import role from "./en_us/role";
import settings from "./en_us/settings";
import user from "./en_us/user";
import workspace from "./en_us/workspace";

import antdEnUS from "ant-design-vue/es/locale/en_US";

const components = {
  antLocale: antdEnUS,
  DayJsName: "en-us",
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
