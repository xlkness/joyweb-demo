import {createRouter, createWebHashHistory, createWebHistory, RouteRecordRaw} from 'vue-router'
import UserView from "@/views/user/UserView.vue";
import LoginView from "@/views/LoginView.vue";
import NewHomeView from "@/views/NewHomeView.vue";
import GMToolHomeView from "@/views/gmtool_old/GMToolHomeView.vue";
import GMToolMenuItem from "@/views/gmtool_old/GMToolMenuItem.vue";
import UserListView from "@/views/user/UserListView.vue";
import TestView from "@/views/TestView.vue";
import HomeView from "@/views/HomeView.vue";

const getGMToolChildrenMenuItems = (enProject: string, project: string) => {
  const query = () => {
    return {
      project: project,
      t: Date.now(),
    }
  }
  return [
    {
      path: "/" + enProject + "/gmtool/operation",
      meta: {
        "key": enProject + "/gmtool/operation",
        "name": "指令服操作",
        "icon": "Tools",
      },
      component: TestView,
      query: query(),
    },
    {
      path: "/" + enProject + "/gmtool/environments",
      meta: {
        key: enProject + "/gmtool/environments",
        "name": "环境管理",
        "icon": "Tools",
      },
      component: TestView,
      query: query(),
    },
    {
      path: "/" + enProject + "/gmtool/command_server_list",
      meta: {
        key: enProject + "/gmtool/command_server_list",
        "name": "指令服管理",
        "icon": "Tools",
      },
      component: TestView,
      query: query(),
    },
    {
      path: "/" + enProject + "/gmtool/like_list",
      meta: {
        key: enProject + "/gmtool/like_list",
        "name": "收藏管理",
        "icon": "Tools",
      },
      component: TestView,
      query: query(),
    },
    {
      path: "/" + enProject + "/gmtool/history_list",
      meta: {
        key: enProject + "/gmtool/history_list",
        "name": "执行历史管理",
        "icon": "Tools",
      },
      component: TestView,
      query: query(),
    },
    {
      path: "/" + enProject + "/gmtool/permission_group_list",
      meta: {
        key: enProject + "/gmtool/permission_group_list",
        "name": "权限组管理",
        "icon": "Tools",
      },
      component: TestView,
      query: query(),
    },
    {
      path: "/" + enProject + "/gmtool/user_list",
      meta: {
        key: enProject + "/gmtool/user_list",
        "name": "用户管理",
        "icon": "Tools",
      },
      component: TestView,
      query: query(),
    },
  ]
}

export let homeRoutes = [
  {
    path: "/gmtool",
    meta: {
      "key": "/gmtool",
      'name': 'GM管理系统',
      'icon': "Tools",
    },
    children: [
      {
        path: "/gmtool/football",
        meta: {
          "key": "/gmtool/football",
          "name": "足球项目",
          "icon": "Tools",
        },
        children: getGMToolChildrenMenuItems("football", "足球"),
      },
      {
        path: "/gmtool/sim",
        meta: {
          "key": "/gmtool/sim",
          "name": "SIM项目",
          "icon": "Tools",
        },
        children: getGMToolChildrenMenuItems("sim", "SIM"),
      },
    ],
  },
  {
    path: "/users",
    meta: {
      "key": "/users",
      'name': '用户管理',
      'icon': "Tools",
    },
    children: [
      {
        path: "/userlist",
        name: "userlist",
        meta: {
          'name': '用户列表',
          'icon': "Tools",
        },
        component: UserListView,
      },
    ],
  },
  {
    path: "/test",
    component: TestView,
    meta: {
      name: "/test1",
      project: "足球",
      name: "test1",
      icon: "Tools",
    },
    query: {
      project: "足球",
      t: Date.now(),
    }
  },
  {
    path: "/test",
    meta: {
      "key": "test2",
      project: "SIM",
      name: "test2",
      icon: "Tools",
    },
    component: TestView,
    query: {
      project: "SIM",
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
    component: HomeView,
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
