<template>
  <div style="height: 100%;width:100%">
    <!--  展示指令服务器列表  -->
    <el-table :data="gmServerTableData" style="width: 100vh;min-height: 10vh;" size="large" stripe="true">
      <el-table-column fixed prop="name" label="指令服务器" width="150">
      </el-table-column>
      <el-table-column fixed prop="desc" label="描 述" width="250">
      </el-table-column>
      <el-table-column fixed prop="addr" label="地 址" width="250">
      </el-table-column>
      <el-table-column fixed prop="func" label="功 能">
        <template #default="scope">
          <el-button @click="handleOperation('operation', scope.$index, scope.row)" size="large" type="success">
            <el-icon style="vertical-align: middle">
              <Operation/>
            </el-icon>
            <span>指令列表</span>
          </el-button>
          <el-button @click="handleOperation('like', scope.$index, scope.row)" size="large" type="primary">
            <el-icon style="vertical-align: middle">
              <Operation/>
            </el-icon>
            <span>我的收藏</span>
          </el-button>
          <el-button @click="handleOperation('history', scope.$index, scope.row)" size="large" type="info">
            <el-icon style="vertical-align: middle">
              <Operation/>
            </el-icon>
            <span>执行历史</span>
          </el-button>
        </template>
      </el-table-column>
    </el-table>

    <!--  点击某个服务器执行列表展示后出现的弹窗  -->
    <el-dialog v-model="dialogData.operVisible" modal="true" style="width: 100vh;height: 70vh">
      <template #header>
        <span style="font-size: 40px">{{curEnv}}/{{dialogData.gmServerData.name}}/GM指令操作列表</span>
      </template>
      <component :is="cmdServerComponent" :project="project" :curEnv="curEnv" :commandserver="dialogData.gmServerData" :commandlist="dialogData.commandlist"></component>
    </el-dialog>

    <el-dialog v-model="dialogData.likeVisible" modal="true" style="width: 100vh;height: 70vh">
      <template #header>
        <span style="font-size: 40px">{{curEnv}}/{{dialogData.gmServerData.name}}/收藏列表</span>
      </template>
      <el-table :data="dialogData.likeList" table-layout="auto" stripe>
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
        <el-table-column prop="command.base_req_params.fields" label="参数" width="250">
          <template #default="scope">
            <el-input :autosize="{minRows: 2, maxRows: 6}" type="textarea" :placeholder="JSON.stringify(scope.row.command.base_req_params.fields)"></el-input>
          </template>
        </el-table-column>
        <el-table-column prop="command.date" label="创建时间" width="150">
        </el-table-column>
        <el-table-column>
          <template #default="scope">
            <el-button  @click="handleShowExecCommand(dialogData.gmServerData, scope.row.command)" type="success" size="large">
              <el-icon style="vertical-align: middle">
                <Operation/>
              </el-icon>
              <span>执行</span>
            </el-button>
<!--            <el-button @click="handleDislike(props.$index, props.row, scope.$index, scope.row)" type="danger" size="large">-->
<!--              <el-icon style="vertical-align: middle">-->
<!--                <Notebook/>-->
<!--              </el-icon>-->
<!--              <span>删除</span>-->
<!--            </el-button>-->
          </template>
        </el-table-column>
      </el-table>
    </el-dialog>

    <el-dialog v-model="dialogData.historyVisible" modal="true" style="width: 100vh;height: 70vh">
      <template #header>
        <span style="font-size: 40px">{{curEnv}}/{{dialogData.gmServerData.name}}/执行历史列表</span>
      </template>
      <el-table :data="dialogData.historyRecordList" style="width: 100%;height: 60vh" table-layout="auto" stripe>
        <el-table-column prop="request_info.env" label="环境">
        </el-table-column>
        <el-table-column prop="request_info.command_server_name" label="指令服">
        </el-table-column>
        <el-table-column prop="request_info.base_req_params.name" label="指令">
        </el-table-column>
        <el-table-column prop="request_info.user" label="用户">
        </el-table-column>
        <el-table-column label="参数" width="80">
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
            <el-input v-model="scope.row.exec_res" :autosize="{minRows: 1, maxRows: 6}" type="textarea" placeholder="执行结果"></el-input>
          </template>
        </el-table-column>
        <el-table-column prop="request_info.date" label="执行时间">
        </el-table-column>
        <el-table-column label="操作">
          <template #default="scope">
            <el-button  @click="handleShowExecCommand(dialogData.gmServerData, scope.row.request_info)" type="primary" size="large">
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

    <el-dialog v-model="dialogData.execShowVisible" modal="true" destroy-on-close style="width: 100vh;height: 70vh">
      <template #header>
        <span style="font-size: 40px">{{curEnv}}/{{dialogData.gmServerData.name}}/[{{dialogData.preExecCommandHistory.base_req_params.name}}]/指令执行</span>
      </template>
      <component
          :is="execCommandComponent"
          :curEnv="curEnv"
          :commandserver="dialogData.gmServerData"
          :command="dialogData.preExecCommandHistory"
          :project="project"
      ></component>
    </el-dialog>
  </div>
</template>

<script lang="ts">
import {defineComponent, ref} from "vue";
import {likeCommand, queryCmdServerCommandList} from "@/requests/gmtool";
import CmdServer from "@/components/gmtool/CmdServer.vue";
import {ElNotification} from "element-plus";
import CmdExecute from "@/components/gmtool/CmdExecute.vue";
import CmdExecuteWithHistory from "@/components/gmtool/CmdExecuteWithHistory.vue";

export default defineComponent({
  props: ['project', 'curEnv', 'commandServerList', 'envList', 'likeList'],
  setup(props) {

    // console.log("refresh env view")
    const project = props.project
    const curEnv = props.curEnv
    const allCommandServerList = props.commandServerList || []
    const allEnvs = props.envList
    const likeList = props.likeList
    const gmServerTableData = ref([])

    // 初始化指令服务器列表
    for (let i=0; i < allCommandServerList.length; i++) {
      for (let j=0; allCommandServerList[i].envs && j < allCommandServerList[i].envs.length; j++) {
        if (curEnv == allCommandServerList[i].envs[j]) {
          gmServerTableData.value.push({
            "name": allCommandServerList[i].name,
            "addr": allCommandServerList[i].addr,
            "desc": allCommandServerList[i].desc,
            "detail": allCommandServerList[i],
          })
        }
      }
    }

    const dialogData = ref({
      operVisible: false,
      likeVisible: false,
      historyVisible: false,
      execShowVisible: false,
      gmServerData: {},
      commandlist: null,
      likeList: [],
      historyRecordList: [],
      preExecCommandHistory: {},
    })

    const allDisable = () => {
      dialogData.value.operVisible = false
      dialogData.value.likeVisible = false
      dialogData.value.historyVisible = false
      dialogData.value.execShowVisible = false
    }
    const operEnable = () => {
      dialogData.value.operVisible = true
      dialogData.value.likeVisible = false
      dialogData.value.historyVisible = false
    }
    const likeEnable = () => {
      dialogData.value.operVisible = false
      dialogData.value.likeVisible = true
      dialogData.value.historyVisible = false
    }
    const historyEnable = () => {
      dialogData.value.operVisible = false
      dialogData.value.likeVisible = false
      dialogData.value.historyVisible = true
    }

    const handleOperation = (oper: string, index: number, row: Object) => {
      dialogData.value.gmServerData=row

      switch (oper) {
        case "operation":
          queryCmdServerCommandList({project: project, name: row.detail.name, env: curEnv}).then((res)=> {
            // console.log('req:', row, 'json res:', res)
            dialogData.value.commandlist = res.payload
            operEnable()
          }, (err)=>{
            console.log("queryCmdServerCommandList error:", err)
          })
          // const url = "http://" + row.addr + `/commandlist?format=jsonp&` + `gm_id=` + `likun`
          // jsonp(url, {format: 'jsonp'}).then((res)=> {
          //
          // })
          break;
        case "like":
          let curCmdServerLikeList = []
          for (let i=0; i < likeList.length; i++) {
            let curLike = likeList[i]
            if (curLike.command.command_server_name == row.name) {
              curCmdServerLikeList.push(curLike)
            }
          }
          dialogData.value.likeList = curCmdServerLikeList
          likeEnable()
          break;
        case "history":
            dialogData.value.historyRecordList = []
            for (let i=0; i < allCommandServerList.length; i++) {
              let curCmdServer = allCommandServerList[i]
              if (curCmdServer.name == row.name) {
                dialogData.value.historyRecordList = curCmdServer.exec_history_list
                break;
              }
            }
          historyEnable()
          break;
      }
      // console.log("点击操作：", oper, ",索引：", index, ",行数据：", row.name, row.addr, row.desc)
    }

    const handleClose = (done: ()=>void) => {
      allDisable()
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

    const handleShowExecCommand = (cmdServer, command) => {
      dialogData.value.gmServerData = cmdServer
      dialogData.value.preExecCommandHistory = command
      dialogData.value.execShowVisible = true
    }

    const cmdServerComponent = CmdServer
    const execCommandComponent = CmdExecuteWithHistory

    return {project, curEnv, gmServerTableData, handleOperation, handleClose, dialogData, cmdServerComponent, execCommandComponent, handleLike, handleShowExecCommand}
  }
})
</script>

<style scoped>

</style>