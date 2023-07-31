<template>
  <div>
    <el-empty description="无数据，请添加环境和指令服！" v-if="envTabs.length <= 0">
    </el-empty>
    <el-tabs v-model="envTabsValue" type="border-card" class="env-tabs" v-else>
      <el-tab-pane class="env-tabs-pane"
                   v-for="item in envTabs"
                   :key="item.name"
                   :label="item.title"
                   :name="item.name"
      >
        <component
            :is="item.content"
            :curEnv="item.name"
            :commandServerList="item.commandServerList"
            :envList="item.envList"
            :project="project"
            :likeList="item.likeList"
            :permissionList="item.permissionList"
            v-if="(envTabsValue == item.name)"
        ></component>
        <!--            {{item.content}}-->
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script lang="ts">
import {defineComponent, onMounted, ref} from "vue";
import {gmtool} from "@/requests/gmtool";
import EnvView from "@/components/gmtool/EnvView.vue";
import {useRouter} from "vue-router";
export default defineComponent({
  setup() {
    const useRoute = useRouter()
    const project = useRoute.currentRoute.value.query.project
    var envTabs = ref([])
    var envTabsValue = ref('0')

    // console.log("project", project)

    gmtool({project: project}).then((res)=> {
      // 获取到gmtool信息，刷新页面
      var commandServerList = res.payload.command_server_list || []
      var envList = res.payload.env_list || []
      var likeList = res.payload.like_list || []
      console.log("gmtool", res.payload)
      for (let i=0; i<envList.length; i++) {
        if (res.payload.is_admin) {
          envTabs.value.push({
            "title": String(envList[i].name),
            "name": String(envList[i].name),
            "content": EnvView,
            "commandServerList": commandServerList,
            "envList": envList,
            "likeList": likeList,
            "permissionList": res.payload.permission.permissions,
          })
          continue
        }
        for (let j=0; j<res.payload.permission.permissions.length; j++) {
          if (res.payload.permission.permissions[j].env == envList[i].name) {
            envTabs.value.push({
              "title": String(envList[i].name),
              "name": String(envList[i].name),
              "content": EnvView,
              "commandServerList": commandServerList,
              "envList": envList,
              "likeList": likeList,
              "permissionList": res.payload.permission.permissions,
            })
            break
          }
        }
      }

      if (envList.length > 0) {
        envTabsValue.value = envList[0].name
      }

      // console.log("commandServerList", commandServerList, envList)
    }, (err) => {
      console.log("request error:", err)
    })

    return {project, envTabs, envTabsValue}
  }
})
</script>

<style lang="scss" scoped>


:deep() {
//.el-tabs--border-card {
//  font-size: 100px;
//}
//.el-tabs__item {
//  font-size: 100px;
//}
  .el-tabs--border-card > .el-tabs__content {
    padding: 32px;
    color: #6b778c;
    /*background-color: aqua;*/
    /*color: #42b983;*/
    font-size: 20px;
    font-weight: 600;
  }
  .el-tabs--border-card .el-tabs__header {
    height: auto;
  }
  .el-tabs--border-card .el-tabs__item {
    font-size: 20px;
    height: auto;
  }
}

</style>