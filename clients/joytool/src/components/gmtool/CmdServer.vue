<template>
  <div>
    <el-table :data="commandlist" style="width: 98vh;height: 60vh">
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
          <el-button @click="handleOperation('operation', scope.$index, scope.row)" type="primary" size="large">
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
  </div>
</template>

<script lang="ts">
import {defineComponent, ref} from "vue";
import {Operation} from "@element-plus/icons-vue";
import CmdExecute from "@/components/gmtool/CmdExecute.vue";

export default defineComponent({
  components: {Operation},
  props: ['project', 'curEnv', 'commandserver', 'commandlist',],
  setup(props) {
    const project = props.project
    const curEnv = props.curEnv
    const commandserver = props.commandserver
    const commandlist = props.commandlist

    const executeCommandDialogData = ref({
      execCommand: {},
      execCommandVisible: false,
    })

    const handleShowExecCommand = (index: number, row: Object) => {
      executeCommandDialogData.value.execCommand = row
      executeCommandDialogData.value.execCommandVisible = true
    }

    const execCommandComponent = CmdExecute

    return {project, curEnv, commandserver, commandlist, handleShowExecCommand, execCommandComponent, executeCommandDialogData}
  }
})
</script>

<style scoped>

</style>