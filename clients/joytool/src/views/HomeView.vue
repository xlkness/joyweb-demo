<template>
  <div>
    <el-container>
      <el-header>
        <HeaderView/>
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
              <el-sub-menu :index="routeItem.meta.name" v-if="(routeItem.children && routeItem.children.length > 0)">
                <template #title>
                  <el-icon>
                    <component :is="routeItem.meta.icon"></component>
                  </el-icon>
                  <span>{{routeItem.meta.name}}</span>
                </template>
                <template v-for="(subRouteItem, idx1) in routeItem.children">
                  <el-menu-item :index="subRouteItem.path"
                                :route="{query: subRouteItem.query}"
                                :title="subRouteItem.meta.name"
                                v-if="!subRouteItem.children || subRouteItem.children.length <= 0">
                    <el-icon>
                      <component class="el-icon" :is="subRouteItem.meta.icon"></component>
                    </el-icon>
                    <span>{{subRouteItem.meta.name}}</span>
                  </el-menu-item>

                  <el-sub-menu :index="subRouteItem.meta.name" v-else>
                    <template #title>
                      <el-icon>
                        <component :is="subRouteItem.meta.icon"></component>
                      </el-icon>
                      <span>{{subRouteItem.meta.name}}</span>
                    </template>
                    <el-menu-item v-for="(subSubRouteItem, idx2) in subRouteItem.children"
                                  :index="subSubRouteItem.path"
                                  :route="{query: subSubRouteItem.query}"
                                  :title="subSubRouteItem.meta.name">
                      <el-icon>
                        <component class="el-icon" :is="subSubRouteItem.meta.icon"></component>
                      </el-icon>
                      <span>{{subSubRouteItem.meta.name}}</span>
                    </el-menu-item>
                  </el-sub-menu>
                </template>
              </el-sub-menu>

              <el-menu-item :index="routeItem.path" :route="{query: routeItem.query}" v-else>
                <el-icon>
                  <component class="el-icon" :is="routeItem.meta.icon"></component>
                </el-icon>
                <span>{{routeItem.meta.name}}</span>
              </el-menu-item>
            </template>
          </el-menu>
        </el-aside>
        <el-main>
          <router-view :key="routerKey" :to="router.currentRoute.value.query"></router-view>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script lang="ts">
import {computed, defineComponent} from "vue";
import {useRoute} from "vue-router";
import router, {homeRoutes} from "@/router";

export default defineComponent({
  setup() {
    const useRouter = useRoute()
    const routerKey = computed(() => {
      return useRouter.path + Math.random()
    })
    const routeList = homeRoutes
    return {routeList, router, routerKey}
  }
})
</script>

<style scoped>

.el-header {
  height: 8vh;
  background-color: white;
}

</style>