import { createRouter, createWebHistory } from "vue-router";
import HomeView from "../views/HomeView.vue";
import BodyLayout from "../layouts/BodyLayout.vue";
import SubLayout from "../layouts/SubLayout.vue";
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/login",
      name: "login",
      component: SubLayout,
    },
    {
      path: "/",
      name: "layout",
      component: BodyLayout,
      children:[
        {
          path: "/cluster",
          name: "cluster",
          component: () =>import("../views/HomeView.vue")
        },
      ]
    },
  ],
});

export default router;
