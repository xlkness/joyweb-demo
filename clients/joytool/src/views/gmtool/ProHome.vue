<template>
  <div class="home">
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
          <el-menu-item v-for="(routeItem, idx) in functionList" class="sub-menu" :index="routeItem.path">
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
  </div>
</template>

<script lang="ts">
import {defineComponent, ref} from "vue";
import {Operation} from "@element-plus/icons-vue";

export default defineComponent({
  components: {Operation},
  props: [ 'functionList'],
  watch: {
    functionList: function (val) {
      // console.log('functionList', val)
    },
  },
  setup() {
    return {}
  }
})
</script>

<style lang="scss" scoped>

.home {
  width: 100%;
  height: 100%;
}

.el-menu {
  height: 100%;
}

.el-menu-item {
  font-size: 100px;
}

</style>