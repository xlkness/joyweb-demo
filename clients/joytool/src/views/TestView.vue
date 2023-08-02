<template>
  <div>
<!--      {{project}}-test-{{this.$router.query}}-->
      <n-menu class="menu" :options="menuOptions"/>
  </div>
</template>

<script lang="ts">
import {defineComponent, h} from "vue";
import {RouterLink, useRouter} from "vue-router";
import router from "@/router";
import type {MenuOption} from 'naive-ui';
import {NMenu} from "naive-ui";

export default defineComponent({
  components: {
    NMenu
  },
  setup() {
    console.log("entry test")
    const useRoute = useRouter()
    // const project = this.$route.query.project
    const project = useRoute.currentRoute.value.query.project
    console.log('path', useRoute.currentRoute.value.fullPath)
    console.log('1', useRoute.currentRoute.value)
    console.log('2', router.currentRoute.value)

    const menuOptions: MenuOption[] = [
      {
        label: "子菜单1abcABCV",
        key: "item1",
        children: [
          {
            type: 'group',
            label: '子菜单11abhcABCV',
            key: 'item11',
            children: [
              {
                label: '子菜单111asdfABCV',
                key: 'item111',
              }
            ]
          }
        ]
      }
    ]
    return {project, menuOptions}
  }
})
</script>

<style lang="scss" scoped>

.menu{
  //font-size: 35px;
  //font-family: "Microsoft YaHei", YouYuan, FangSong, "Microsoft YaHei", serif;
  width: 35vh;
}
.n-menu-item {
  font-size: 60px;
  font-family: YouYuan, FangSong, "Microsoft YaHei", serif;
}
.n-menu-item-content {
  font-size: 55px;
  font-family: YouYuan, FangSong, "Microsoft YaHei", serif;
}
.menu .n-menu-item-content-header {
  font-size: 60px;
  font-family: YouYuan, FangSong, "Microsoft YaHei", serif;
}

</style>