<template>
  <div style="width: 98vh;height: 60vh">
    <el-table
        :data="likeRecordList"
        row-key="cmdServer.name"
        :expand-row-keys="expandRowKeys"
        style="width: 98vh;height: 60vh;font-size: 23px !important;" table-layout="auto"
        stripe
        @row-click="clickRow"
    >
      <el-table-column type="expand" >
        <template #default="props">
          <div m="4">
            <el-table :data="props.row.subList" table-layout="auto" stripe>
              <el-table-column prop="id" label="ID">
              </el-table-column>
              <!--            <el-table-column prop="command.env" label="环境">-->
              <!--            </el-table-column>-->
              <!--            <el-table-column prop="command.command_server_name" label="指令服">-->
              <!--            </el-table-column>-->
              <el-table-column prop="command.base_req_params.name" label="指令">
              </el-table-column>
              <el-table-column prop="command.user" label="用户">
              </el-table-column>
              <el-table-column prop="command.base_req_params.fields" label="参数">
                <template #default="scope">
                  <el-popover
                      placement="right-start"
                      title="参数"
                      :width="200"
                      trigger="hover"
                      :content="JSON.stringify(scope.row.command.base_req_params.fields)"
                  >
                    <template #reference>
                      <el-button type="textarea">...</el-button>
                    </template>
                  </el-popover>
                  <!-- <el-input :autosize="{minRows: 2, maxRows: 6}" type="textarea" :placeholder="JSON.stringify(scope.row.command.base_req_params.fields)"></el-input> -->
                </template>
              </el-table-column>
              <el-table-column prop="command.date" label="创建时间" width="150">
              </el-table-column>
              <el-table-column>
                <template #default="scope">
                  <el-button  @click="handleShowExecCommand(props.row.cmdServer, scope.row.command)" type="primary" size="large">
                    <el-icon style="vertical-align: middle">
                      <Operation/>
                    </el-icon>
                    <span>执行</span>
                  </el-button>
                  <el-button @click="handleDislike(props.$index, props.row, scope.$index, scope.row)" type="danger" size="large">
                    <el-icon style="vertical-align: middle">
                      <Notebook/>
                    </el-icon>
                    <span>删除</span>
                  </el-button>
                </template>
              </el-table-column>
            </el-table>
          </div>

        </template>
      </el-table-column>
      <el-table-column prop="cmdServer.name" label="指令服"></el-table-column>
      <el-table-column prop="cmdServer.envs" label="环境"></el-table-column>
      <el-table-column prop="cmdServer.desc" label="描述"></el-table-column>
    </el-table>

    <el-dialog v-model="dialogData.execShowVisible" modal="true" destroy-on-close style="width: 100vh;height: 70vh">
      <template #header>
        <span style="font-size: 40px">{{dialogData.preExecCommandHistory.env}}/{{dialogData.gmServerData.name}}/[{{dialogData.preExecCommandHistory.base_req_params.name}}]/指令执行</span>
      </template>
      <component
          :is="execCommandComponent"
          :curEnv="dialogData.preExecCommandHistory.env"
          :commandserver="dialogData.gmServerData"
          :command="dialogData.preExecCommandHistory"
          :project="project"
      ></component>
    </el-dialog>
  </div>
</template>

<script lang="ts">
import {defineComponent, ref} from "vue";
import {dislikeCommand, gmtool} from "@/requests/gmtool";
import {ElNotification} from "element-plus";
import CmdExecuteWithHistory from "@/components/gmtool/CmdExecuteWithHistory.vue";
import {useRouter} from "vue-router";

export default defineComponent({
  setup() {
    const useRoute = useRouter()
    const project = useRoute.currentRoute.value.query.project
    var likeRecordList = ref([])
    const dialogData = ref({
      execShowVisible: false,
      gmServerData: {},
      preExecCommandHistory: {},
    })
    const expandRowKeys = ref([])

    gmtool({project: project}).then((res)=> {
      // 获取到gmtool信息，刷新页面
      var commandServerList = res.payload.command_server_list || []
      var envList = res.payload.env_list || []
      var likeList = res.payload.like_list || []

      let index = -1
      for (let i=0; i<commandServerList.length; i++) {
        let subList = []
        for (let j=0; j<likeList.length; j++) {
          if (commandServerList[i].name == likeList[j].command.command_server_name) {
            subList.push(likeList[j])
          }
        }
        if (subList.length > 0) {
          index++
          likeRecordList.value.push({
            index: index,
            cmdServer: commandServerList[i],
            subList: subList,
          })
        }
      }

      // console.log("likeList", likeRecordList.value)
    }, (err) => {
      console.log("request error:", err)
    })

    const handleShowExecCommand = (cmdServer, command) => {
      console.log(cmdServer, command)
      dialogData.value.gmServerData = cmdServer
      dialogData.value.preExecCommandHistory = command
      dialogData.value.execShowVisible = true
    }

    const handleDislike = (index, row, childIndex, childRow) => {
      // console.log('click', index, row, childIndex, childRow, likeRecordList.value)
      dislikeCommand({project: project, id: String(childRow.id)}).then((res) => {
        ElNotification({
          title: "取消收藏通知",
          message: "取消收藏[" + row.id + "]成功！",
          type: 'success',
          duration: 4000,
          "show-close": true,
        })

        let child = likeRecordList.value[index]
        child.subList.splice(childIndex, 1)
        likeRecordList.value[index] = child

        // likeRecordList.value.splice(index, 1)
      }, (err) => {
        console.log("dislike error:", err)
      })
    }

    const execCommandComponent = CmdExecuteWithHistory

    const clickRow = (row, column, event) => {
      for (let i=0; i < expandRowKeys.value.length; i++) {
        if (expandRowKeys.value[i] == row.cmdServer.name) {
          expandRowKeys.value.splice(i, 1)
          return
        }
      }
      expandRowKeys.value.push(row.cmdServer.name)
    }

    return {likeRecordList, expandRowKeys, dialogData, execCommandComponent, clickRow, handleShowExecCommand, handleDislike}
  }
})
</script>

<style scoped>

</style>