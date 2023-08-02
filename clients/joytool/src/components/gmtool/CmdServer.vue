<template>
  <div>
    <el-table :data="commandlist" style="width: 98vh;height: 57vh">
      <el-table-column fixed prop="name" label="路径">

      </el-table-column>
      <el-table-column fixed prop="desc" label="描述">

      </el-table-column>
      <el-table-column fixed prop="detail" label="详情">
        <template #default="scope">
          <el-button  @click="handleShowExecCommand(scope.$index, scope.row)" type="success" size="large">
            <el-icon style="vertical-align: middle">
              <Operation/>
            </el-icon>
            <span>执行指令</span>
          </el-button>
          <el-button @click="handleShowHistory(scope.$index, scope.row)" type="primary" size="large">
            <el-icon style="vertical-align: middle">
              <Notebook/>
            </el-icon>
            <span>最近使用</span>
          </el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-dialog v-model="executeCommandDialogData.execCommandVisible" modal="true" destroy-on-close style="width: 100vh;height: 70vh">
      <template #header>
        <span style="font-size: 40px">{{curEnv}}/{{commandserver.name}}/[{{executeCommandDialogData.execCommand.name}}]/指令执行</span>
      </template>
      <component
          :is="execCommandComponent"
          :curEnv="curEnv"
          :commandserver="commandserver"
          :command="executeCommandDialogData.execCommand"
          :project="project"
      ></component>
    </el-dialog>

    <el-dialog v-model="historyDialogData.historyCommandVisible" modal="true" style="width: 100vh;height: 70vh">
      <template #header>
        <span style="font-size: 40px">{{curEnv}}/{{commandserver.name}}/[{{historyDialogData.execCommand.name}}]/执行历史列表</span>
      </template>
      <el-table :data="historyDialogData.historyRecordList" size="small" style="width: 100%;height: 57vh" table-layout="auto" stripe>
        <el-table-column prop="request_info.env" label="环境">
        </el-table-column>
        <el-table-column prop="request_info.command_server_name" label="指令服">
        </el-table-column>
        <el-table-column prop="request_info.base_req_params.name" label="指令">
        </el-table-column>
        <el-table-column prop="request_info.user" label="用户">
        </el-table-column>
        <el-table-column label="参数">
          <template #default="scope">
            <el-popover
                placement="right-start"
                title="参数"
                :width="200"
                trigger="hover"
                :content="scope.row.request_info.base_req_params.fields"
            >
              <template #reference>
                <el-button type="textarea">...</el-button>
              </template>
            </el-popover>
          </template>
        </el-table-column>
        <el-table-column prop="exec_res_code" label="返回码">
        </el-table-column>
        <el-table-column prop="exec_res" label="执行结果">
          <template #default="scope">
            <el-popover
                placement="right-start"
                title="执行结果"
                :width="100"
                trigger="hover"
                :content="scope.row.exec_res"
            >
              <template #reference>
                <el-button type="textarea">...</el-button>
              </template>
            </el-popover>
            <!-- <el-input v-model="scope.row.exec_res" :autosize="{minRows: 1, maxRows: 6}" type="textarea" placeholder="执行结果"></el-input> -->
          </template>
        </el-table-column>
        <el-table-column prop="request_info.date" label="执行时间">
        </el-table-column>
        <el-table-column label="操作">
          <template #default="scope">
            <el-button  @click="handleHistoryShowExecCommand(commandserver, scope.row.request_info)" type="primary" size="large">
              <el-icon style="vertical-align: middle">
                <Operation/>
              </el-icon>
              <span>执行</span>
            </el-button>
            <el-button @click="handleLike(scope.$index, scope.row)" type="success" size="large">
              <el-icon style="vertical-align: middle">
                <Notebook/>
              </el-icon>
              <span>收藏</span>
            </el-button>
<!--            <el-button @click="handleDeleteExecHistory( scope.$index, scope.row)" type="danger" size="large">-->
<!--              <el-icon style="vertical-align: middle">-->
<!--                <Delete/>-->
<!--              </el-icon>-->
<!--              <span>删除</span>-->
<!--            </el-button>-->
          </template>
        </el-table-column>
      </el-table>
    </el-dialog>

    <el-dialog v-model="historyDialogData.execCommandVisible" modal="true" destroy-on-close style="width: 100vh;height: 70vh">
      <template #header>
        <span style="font-size: 40px">{{curEnv}}/{{commandserver.name}}/[{{historyDialogData.preExecCommandHistory.base_req_params.name}}]/指令执行</span>
      </template>
      <component
          :is="execCommandComponentHistory"
          :curEnv="curEnv"
          :commandserver="commandserver"
          :command="historyDialogData.preExecCommandHistory"
          :project="project"
      ></component>
    </el-dialog>
  </div>
</template>

<script lang="ts">
import {defineComponent, ref} from "vue";
import {Operation} from "@element-plus/icons-vue";
import CmdExecute from "@/components/gmtool/CmdExecute.vue";
import CmdExecuteWithHistory from "./CmdExecuteWithHistory.vue";
import { likeCommand } from "@/requests/gmtool";
import LocalCache from "@/stores/localCache";

export default defineComponent({
  components: {Operation},
  props: ['project', 'curEnv', 'commandserver', 'commandlist', 'isAdmin'],
  setup(props) {
    const project = props.project
    const curEnv = props.curEnv
    const commandserver = props.commandserver
    const commandlist = props.commandlist
    const isAdmin = props.isAdmin || false
    const userInfo = LocalCache.getCache("userInfo")

    const executeCommandDialogData = ref({
      execCommand: {},
      execCommandVisible: false,
    })

    const historyDialogData = ref({
      execCommand: {},
      historyRecordList: [],
      historyCommandVisible: false,
      execCommandVisible: false,
      preExecCommandHistory: {},
    })

    const handleShowExecCommand = (index: number, row: Object) => {
      executeCommandDialogData.value.execCommand = row
      executeCommandDialogData.value.execCommandVisible = true
    }

    const handleShowHistory = (index: number, row: Object) => {
      historyDialogData.value.execCommand = row
      historyDialogData.value.historyRecordList = []
      for (let i=0; i < commandserver.detail.exec_history_list.length; i++) {
        let curHistory = commandserver.detail.exec_history_list[i]
        // console.log(curHistory, commandserver)
        if (isAdmin || (curHistory.request_info.base_req_params.name == row.name &&
        curHistory.request_info.env == commandserver.name && curHistory.request_info.user == userInfo.username)) {
          historyDialogData.value.historyRecordList.push(curHistory)
        }
      }
      
      historyDialogData.value.historyCommandVisible = true
    }
    
    const handleHistoryShowExecCommand = (index: number, row: Object) => {
      historyDialogData.value.execCommand = row
      historyDialogData.value.preExecCommandHistory = row
      historyDialogData.value.execCommandVisible = true
    }

    const handleLike = (index, row) => {
      let fieldsValue = []
      for (let i=0; row.request_info.base_req_params.fields && i < row.request_info.base_req_params.fields.length; i++) {
        const field = row.request_info.base_req_params.fields[i]
        fieldsValue.push({name: field.key, value: field.value})
      }
      let baseExecInfo = {
        name: row.request_info.base_req_params.name,
        fields: fieldsValue,
      }
      let execCommandInfo = {
        project: project,
        command_server_name: row.request_info.command_server_name,
        env: row.request_info.base_req_params.env,
        base_info: baseExecInfo,
      }
      likeCommand(execCommandInfo).then((res)=>{
        console.log("exec command resp:", res)
        ElNotification({
          title: "收藏结果通知",
          message: "收藏指令[" + row.request_info.base_req_params.name + "]成功！",
          type: 'success',
          duration: 4000,
          "show-close": true,
        })
        // commandExecResult.value = JSON.stringify(res)
      }, (err) => {
        console.log("request error:", err)
      })
    }

    const execCommandComponent = CmdExecute
    const execCommandComponentHistory = CmdExecuteWithHistory

    return {project, curEnv, commandserver, commandlist, handleShowHistory, handleHistoryShowExecCommand,
       handleShowExecCommand, execCommandComponent, execCommandComponentHistory, executeCommandDialogData, historyDialogData, handleLike}
  }
})
</script>

<style scoped>

</style>