<template>
  <el-menu-item class="sub-menu" :index="menuItem.path">
    <el-icon>
      <component class="el-icon" :is="menuItem.meta.icon"></component>
    </el-icon>
    <span>{{menuItem.meta.name}}</span>
  </el-menu-item>
</template>


<script lang="ts">
import {defineComponent} from "vue";

export default defineComponent({
  props: ['menuItem'],
  setup(props) {
    const menuItem = props.menuItem
    return {menuItem}
  }
})
</script>

<style scoped>

</style>