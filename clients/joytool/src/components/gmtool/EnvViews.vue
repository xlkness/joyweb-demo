<template>
  <div>
    <el-tabs v-model="envTabsValue" type="border-card" class="env-tabs">
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
export default defineComponent({
  props: ['project'],
  setup(props) {
    const project = props.project
    var envTabs = ref([])
    var envTabsValue = ref('0')

    gmtool({project: project}).then((res)=> {
      // 获取到gmtool信息，刷新页面
      var commandServerList = res.payload.command_server_list || []
      var envList = res.payload.env_list || []
      var likeList = res.payload.like_list || []

      for (let i=0; i<envList.length; i++) {
        envTabs.value.push({
          "title": String(envList[i].name),
          "name": String(envList[i].name),
          "content": EnvView,
          "commandServerList": commandServerList,
          "envList": envList,
          "likeList": likeList,
        })
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

<style scoped>

</style>