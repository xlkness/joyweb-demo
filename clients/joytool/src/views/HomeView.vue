<template>
  <div>
    <el-container>
      <el-header>
        <HeaderView :username="username"/>
      </el-header>
      <el-container>
        <el-aside>
          <el-menu
              active-text-color="#bdf2ff"
              class="el-menu"
              background-color="#ffffff"
              :default-active="activeIndex"
              text-color="black"
              router
          >
            <!--最多支持3层前套-->
            <template v-for="(routeItem, idx) in routeList">
              <el-sub-menu class="sub-menu1" :index="routeItem.meta.name + String(idx)" v-if="(routeItem.children && routeItem.children.length > 0)">
                <template  #title>
                  <div class="sub-menu1-title">
                    <el-icon>
                      <component :is="routeItem.meta.icon"></component>
                    </el-icon>
                    <span>{{routeItem.meta.name}}</span>
                  </div>
                </template>
                <template v-for="(subRouteItem, idx1) in routeItem.children">
                  <el-menu-item :index="subRouteItem.path"
                                class="menu-item2"
                                :route="{path: subRouteItem.path, query: subRouteItem.query}"
                                :title="subRouteItem.meta.name"
                                v-if="!subRouteItem.children || subRouteItem.children.length <= 0">
                    <el-icon>
                      <component class="el-icon" :is="subRouteItem.meta.icon"></component>
                    </el-icon>
                    <span>{{subRouteItem.meta.name}}</span>
                  </el-menu-item>

                  <el-sub-menu class="sub-menu2" :index="subRouteItem.meta.name + String(idx1)" v-else>
                    <template #title>
                      <div class="sub-menu2-title">
                        <el-icon>
                          <component :is="subRouteItem.meta.icon"></component>
                        </el-icon>
                        <span>{{subRouteItem.meta.name}}</span>
                      </div>
                    </template>
                    <el-menu-item v-for="(subSubRouteItem, idx2) in subRouteItem.children"
                                  class="menu-item3"
                                  :index="subSubRouteItem.path"
                                  :route="{path: subSubRouteItem.path, query: subSubRouteItem.query}"
                                  :title="subSubRouteItem.meta.name">
                      <el-icon>
                        <component class="el-icon" :is="subSubRouteItem.meta.icon"></component>
                      </el-icon>
                      <span>{{subSubRouteItem.meta.name}}</span>
                    </el-menu-item>
                  </el-sub-menu>
                </template>
              </el-sub-menu>

              <el-menu-item class="menu-item1" :index="routeItem.path" :route="{path: routeItem.path, query: routeItem.query}" v-if="!routeItem.children || routeItem.children.length == 0">
                <el-icon>
                  <component class="el-icon" :is="routeItem.meta.icon"></component>
                </el-icon>
                <span>{{routeItem.meta.name}}</span>
              </el-menu-item>
            </template>
          </el-menu>
        </el-aside>
        <el-main>
          <router-view :key="routerKey"></router-view>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script lang="ts">
import {computed, defineComponent} from "vue";
import {useRoute} from "vue-router";
import router, {homeRoutes} from "@/router";
import LocalCache from "@/stores/localCache";

export default defineComponent({
  setup() {
    const useRouter = useRoute()
    const routerKey = computed(() => {
      return useRouter.path + Math.random()
    })
    const routeList = homeRoutes
    const userInfo = LocalCache.getCache("userInfo")
    // console.log("userinfo", userInfo.username)
    const username = userInfo.username
    return {routeList, router, routerKey, username}
  }
})
</script>

<style scoped>
.sub-menu1-title {
  font-size: 35px !important;
  font-family: YouYuan,  Microsoft YaHei, sans-serif;
}

.sub-menu1:hover {
  background-color: #445ebf !important;
}

:deep() {
  .sub-menu1:hover .el-sub-menu__title {
    background-color: #445ebf !important;
  }
}


.sub-menu2 {
  font-size: 30px;
  font-family: YouYuan,  Microsoft YaHei, sans-serif;
}

.menu-item1 {
  font-size: 35px;
  font-family: YouYuan,  Microsoft YaHei, sans-serif;
}

.menu-item2 {
  font-size: 30px;
  font-family: YouYuan,  Microsoft YaHei, sans-serif;
}

.menu-item3 {
  font-size: 25px;
  font-family: YouYuan,  Microsoft YaHei, sans-serif;
}

</style>

<style lang="scss" scoped>

.el-header {
  //height: 10vh;
  background-color: white;
  margin-bottom: 50px;
}

/*:deep(.el-menu) {*/
/*  font-size: 100px;*/
/*  font-family: YouYuan,Arial,Microsoft YaHei,sans-serif !important;*/
/*}*/

/*.home-menu .el-menu {*/
/*  font-size: 100px;*/
/*  font-family: YouYuan,Arial,Microsoft YaHei,sans-serif !important;*/
/*}*/

:deep() {
  .el-menu {
    font-family: YouYuan,Arial,Microsoft YaHei,sans-serif !important;
  }
  .el-sub-menu__title {
    font-family: YouYuan,Arial,Microsoft YaHei,sans-serif !important;
    font-size: 30px !important;
    background-color: #ffffff !important;
  }
  .el-sub-menu__title:hover {
    font-family: YouYuan,Arial,Microsoft YaHei,sans-serif !important;
    outline: 0 !important;
    color: white !important;
    //border-radius: 5%;
    background-color: #445ebf !important;
  }
  .el-menu-item {
    font-family: YouYuan,Arial,Microsoft YaHei,sans-serif !important;
    font-size: 30px;
    background-color: #ffffff !important;
  }
  .el-menu-item:hover {
    font-family: YouYuan,Arial,Microsoft YaHei,sans-serif !important;
    outline: 0 !important;
    color: white !important;
    //border-radius: 6%;
    background-color: #445ebf !important;
  }
  .el-menu-item.is-active {
    font-family: YouYuan,Arial,Microsoft YaHei,sans-serif !important;
    color: #1890ff !important;
  }
}

</style>
