<template>
  <div class="home">
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
            <el-menu-item v-for="(routeItem, idx) in newRouteList" class="sub-menu" :index="routeItem.path">
              <el-icon>
                <component class="el-icon" :is="routeItem.meta.icon"></component>
              </el-icon>
              <span>{{routeItem.meta.name}}</span>
            </el-menu-item>
          </el-menu>
        </el-aside>
        <el-main>
          <router-view></router-view>
        </el-main>
      </el-container>
    </el-container>
  </div>

</template>

<script lang="ts">
import {defineComponent} from "vue";
import {homeRoutes} from '../router/index';
import GMToolMenuItem from "@/views/gmtool/GMToolMenuItem.vue";

export default defineComponent({
  components: {GMToolMenuItem},
  setup() {

    const newRouteList = homeRoutes
    return {newRouteList, }
  }
})
</script>

<style lang="scss" scoped>

.home {
  width: 100%;
  height: 100%;
  background-color: #f1f5fe;
  background-size: cover;
}

.container {
  height: 100%;
}

.el-header {
  height: 8vh;
  background-color: white;
}

.el-aside {
  height: 92vh;
}

.el-menu {
  height: 100%;
}

.el-menu-item {
  font-size: 30px;
  background-color: #ffffff !important;
}

.el-menu-item:hover {
  outline: 0 !important;
  color: white !important;
  border-radius: 6%;
  background-color: #445ebf !important;
}

.el-menu-item.is-active {
  color: #1890ff !important;
}
</style>