<template>
  <div>
    <el-container>
      <el-header>
        <el-page-header style="font-size: 100px; line-height: 100px; margin-left: 20px" @back="handleBackHome">
          <template #icon>
            <el-icon size="28px" style="margin-bottom: 8px">
              <DArrowLeft></DArrowLeft>
            </el-icon>
          </template>
          <template #title>
            <span style="font-size: 30px">返回选择项目</span>
          </template>
          <template #content>
            <span class="text-large font-600 mr-3" style="font-size: 30px">
              {{project}}项目GM管理系统
            </span>
          </template>
        </el-page-header>
      </el-header>
      <el-main class="functionMain">
        <el-tabs type="border-card" v-model="selectTabName" class="function-tab" :tab-position="tabPosition" @tab-click="handleClick">
          <el-tab-pane v-for="tabPane in TabList" :label="tabPane.name" :name="tabPane.key" :key="tabPane.keyCount">
            <component :is="tabPane.component" v-if="(selectTabName == tabPane.key)" :project="project"
                       subsystem="gmtool" subsystemName="GM管理系统" :subsystemGroups="subsystemGroups"></component>
          </el-tab-pane>
        </el-tabs>
      </el-main>
    </el-container>
  </div>
</template>

<script lang="ts">
import {defineComponent, ref} from "vue";
import ProjectSelectView from "@/views/gmtool_old/ProjectSelectView.vue";
import EnvViews from "@/components/gmtool/EnvViews.vue";
import EditCmdServerListView from "@/components/gmtool/EditCmdServerListView.vue";
import EditEnvListView from "@/components/gmtool/EditEnvListView.vue";
import EditLikeList from "@/components/gmtool/EditLikeList.vue";
import EditExecHistoryList from "@/components/gmtool/EditExecHistoryList.vue";
import UserList from "@/components/user/UserList.vue";
import PermissionGroupList from "@/components/gmtool/PermissionGroupList.vue";
import {gmtool, permissionGroupList} from "@/requests/gmtool";

export default defineComponent({
  props: ['project', 'currentComponent',],
  setup(props, ctx) {
    let project = props.project
    const tabPosition = ref('left')
    // console.log('receive pro:', project)

    const selectTabName = ref('exec')
    const subsystemGroups = ref([])

    const projectList = new Set
    gmtool({project: project}).then((res)=> {
      // 获取到gmtool信息，刷新页面
      // var commandServerList = res.payload.command_server_list || []
      // var envList = res.payload.env_list || []
      // var likeList = res.payload.like_list || []

      var permissionList = res.payload.permission && res.payload.permission.permissions || []
      for (let i=0; i < permissionList.length; i++) {
        var curPermission = permissionList[i]
        projectList.add(curPermission.project)
      }

    })

    const TabList = ref([
      {
        name: '执行指令',
        key: "exec",
        keyCount : 0,
        component: EnvViews,
      },
      {
        name: '环境列表',
        key: "envList",
        keyCount : 0,
        component: EditEnvListView,
      },
      {
        name: '指令服列表',
        key: "cmdServerList",
        keyCount : 0,
        component: EditCmdServerListView,
      },
      {
        name: '收藏列表',
        key: "likeList",
        keyCount : 0,
        component: EditLikeList,
      },
      {
        name: '执行历史',
        key: "historyList",
        keyCount : 0,
        component: EditExecHistoryList,
      },
      {
        name: "权限组管理",
        key: "permissions",
        keyCount: 0,
        component: PermissionGroupList,
      },
      {
        name: '用户管理',
        key: "userList",
        keyCount : 0,
        component: UserList,
      }
    ])

    const handleBackHome = () => {
      ctx.emit('update:currentComponent', ProjectSelectView)
      ctx.emit('update:project', '')
    }


    const handleClick = (paneCtx, event) => {
      // paneCtx.props.key = count.value
      let curTabData = TabList.value.filter(item => item.key == paneCtx.props.name)[0]
      curTabData.keyCount++
      selectTabName.value = paneCtx.props.name

      permissionGroupList({project: project}).then((res)=>{
        if (res.payload) {
          subsystemGroups.value = []
          res.payload.forEach(function (value, index, obj) {
            subsystemGroups.value.push(value.name)
          })
          console.log("groups:", subsystemGroups.value)
        }
      }, (err) => {
        console.log("groupList return error", err)
      })
      // console.log('click pane', paneCtx.props.name)
    }

    permissionGroupList({project: project}).then((res)=>{
      if (res.payload) {
        subsystemGroups.value = []
        res.payload.forEach(function (value, index, obj) {
          subsystemGroups.value.push(value.name)
        })
      }
    }, (err) => {
      console.log("groupList return error", err)
    })

    return {handleBackHome, TabList, selectTabName, handleClick, project, tabPosition, subsystemGroups}
  }
})
</script>

<style lang="scss">

.functionMain {
  margin-top: 10px;
  //background-color: green;
}

.el-tabs {

}

.el-tabs__nav {
  background-color: white !important;
}

.el-tabs__item {
  font-size: 30px !important;
}


</style>