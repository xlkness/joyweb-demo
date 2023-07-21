<template>
  <div style="width: 100%;height: 60vh" >
    <el-table :data="historyRecordList" style="width: 100%;height: 60vh" table-layout="auto" stripe>
      <el-table-column prop="history.request_info.env" label="环境">
      </el-table-column>
      <el-table-column prop="history.request_info.command_server_name" label="指令服">
      </el-table-column>
      <el-table-column prop="history.request_info.base_req_params.name" label="指令">
      </el-table-column>
      <el-table-column prop="history.request_info.user" label="用户">
      </el-table-column>
      <el-table-column label="参数" width="80">
        <template #default="scope">
          <el-popover
              placement="right-start"
              title="参数"
              :width="200"
              trigger="hover"
              :content="scope.row.history.request_info.base_req_params.fields"
          >
            <template #reference>
              <el-button type="textarea">...</el-button>
            </template>
          </el-popover>
        </template>
      </el-table-column>
      <el-table-column prop="history.exec_res_code" label="返回码">
      </el-table-column>
      <el-table-column prop="history.exec_res" label="执行结果">
        <template #default="scope">
          <el-input v-model="scope.row.history.exec_res" :autosize="{minRows: 1, maxRows: 6}" type="textarea" placeholder="执行结果"></el-input>
        </template>
      </el-table-column>
      <el-table-column prop="history.request_info.date" label="执行时间">
      </el-table-column>
      <el-table-column label="操作">
        <template #default="scope">
          <el-button  @click="handleShowExecCommand({name: scope.row.history.request_info.command_server_name}, scope.row.history.request_info)" type="primary" size="large">
            <el-icon style="vertical-align: middle">
              <Operation/>
            </el-icon>
            <span>再次执行</span>
          </el-button>
          <el-button @click="handleLike(scope.$index, scope.row)" type="success" size="large">
            <el-icon style="vertical-align: middle">
              <Notebook/>
            </el-icon>
            <span>收藏</span>
          </el-button>
          <el-button @click="handleDeleteExecHistory( scope.$index, scope.row)" type="danger" size="large">
            <el-icon style="vertical-align: middle">
              <Delete/>
            </el-icon>
            <span>删除</span>
          </el-button>
        </template>
      </el-table-column>
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
import {deleteExecHistory, gmtool, likeCommand} from "@/requests/gmtool";
import EnvView from "@/components/gmtool/EnvView.vue";
import {ElNotification} from "element-plus";
import {Delete} from "@element-plus/icons-vue";
import CmdExecuteWithHistory from "@/components/gmtool/CmdExecuteWithHistory.vue";

export default defineComponent({
  components: {Delete},
  props: ['project'],
  setup(props) {
    const project = props.project
    var historyRecordList = ref([])
    const dialogData = ref({
      execShowVisible: false,
      gmServerData: {},
      preExecCommandHistory: {},
    })

    gmtool({project: project}).then((res)=> {
      // 获取到gmtool信息，刷新页面
      var commandServerList = res.payload.command_server_list || []
      var envList = res.payload.env_list || []

      for (let i=0; i<commandServerList.length; i++) {
        if (!commandServerList[i].exec_history_list) {
          continue
        }
        for (let j=0; j<commandServerList[i].exec_history_list.length; j++) {
          historyRecordList.value.push({
            index: j,
            cmdServer: commandServerList[i],
            history: commandServerList[i].exec_history_list[j]
          })
        }
      }
      // console.log("commandServerList", commandServerList, envList)
    }, (err) => {
      console.log("request error:", err)
    })

    const handleShowExecCommand = (cmdServer, command) => {
      dialogData.value.gmServerData = cmdServer
      dialogData.value.preExecCommandHistory = command
      dialogData.value.execShowVisible = true
    }

    const handleLike = (index, row) => {
      let fieldsValue = []
      for (let i=0; row.history.request_info.base_req_params.fields && i < row.history.request_info.base_req_params.fields.length; i++) {
        const field = row.history.request_info.base_req_params.fields[i]
        fieldsValue.push({name: field.name, value: field.value})
      }
      let baseExecInfo = {
        name: row.history.request_info.base_req_params.name,
        fields: fieldsValue,
      }
      let execCommandInfo = {
        project: project,
        command_server_name: row.history.request_info.command_server_name,
        env: row.history.request_info.base_req_params.env,
        base_info: baseExecInfo,
      }
      likeCommand(execCommandInfo).then((res)=>{
        console.log("exec command resp:", res)
        ElNotification({
          title: "收藏结果通知",
          message: "收藏指令[" + row.history.request_info.base_req_params.name + "]成功！",
          type: 'success',
          duration: 4000,
          "show-close": true,
        })
        // commandExecResult.value = JSON.stringify(res)
      }, (err) => {
        console.log("request error:", err)
      })
    }

    const handleDeleteExecHistory = (index, row) => {
      console.log(row)
      deleteExecHistory({project: project, name: row.cmdServer.name, index: row.index}).then((res) => {
        ElNotification({
          title: "删除结果通知",
          message: "删除历史[" + row.index + "]成功！",
          type: 'success',
          duration: 4000,
          "show-close": true,
        })
        historyRecordList.value.splice(index, 1)
      }, (err) => {

      })
    }

    const execCommandComponent = CmdExecuteWithHistory

    return {historyRecordList, dialogData, handleShowExecCommand,handleLike, handleDeleteExecHistory, execCommandComponent}
  }
})
</script>

<style scoped>

</style>