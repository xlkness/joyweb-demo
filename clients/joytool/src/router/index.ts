import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import UserView from "@/views/user/UserView.vue";
import LoginView from "@/views/LoginView.vue";
import NewHomeView from "@/views/NewHomeView.vue";
import GMToolHomeView from "@/views/gmtool/GMToolHomeView.vue";
import GMToolMenuItem from "@/views/gmtool/GMToolMenuItem.vue";

export let homeRoutes = [
  {
    path: "/gmtool",
    name: "gmtool",
    meta: {
      'name': 'GM管理系统',
      'icon': "Tools",
      'useMenuItem': true,
      'menuItem': GMToolMenuItem,
    },
    component: GMToolHomeView,
  },
  {
    path: "/users",
    name: "users",
    meta: {
      'name': '用户管理',
      'icon': "Tools",
      'useMenuItem': false,
    },
    component: UserView,
  },
]

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'home',
    component: NewHomeView,
    children: homeRoutes,
  },
  {
    path: '/login',
    name: 'login',
    component: LoginView,
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
