import {createRouter, createWebHashHistory, createWebHistory, RouteRecordRaw} from 'vue-router'
import UserView from "@/views/user/UserView.vue";
import LoginView from "@/views/LoginView.vue";
import NewHomeView from "@/views/NewHomeView.vue";
import GMToolHomeView from "@/views/gmtool/GMToolHomeView.vue";
import GMToolMenuItem from "@/views/gmtool/GMToolMenuItem.vue";
import UserListView from "@/views/user/UserListView.vue";
import TestView from "@/views/TestView.vue";

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
    children: [
      {
        path: "/userlist",
        name: "userlist",
        meta: {
          'name': '用户管理',
          'icon': "Tools",
          'useMenuItem': false,
          // 'menuItem': GMToolMenuItem,
        },
        component: UserListView,
      },
    ],
    // component: UserView,
  },
  {
    path: "/test/:project",
    name: "test1",
    component: TestView,
    meta: {
      name: "test1",
    },
    query: {
      t: Date.now()
    }
  },
  {
    path: "/test/:project",
    name: "test2",
    meta: {
      name: "test2",
    },
    component: TestView,
    query: {
      t: Date.now()
    }
  }
]

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    redirect: '/login',
  },
  {
    path: '/home',
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
  history: createWebHashHistory(process.env.BASE_URL),
  routes
})

export default router
