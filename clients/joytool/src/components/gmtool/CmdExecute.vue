<template>
  <div>
    <el-scrollbar max-height="40vh">
      <el-form
          ref="commandFormRef"
          :model="commandFormData"
          label-width="80px"
      >
        <el-form-item
            v-for="(cmdField, index) in commandFormData.fields"
            :label="cmdField.key"
            :key="cmdField.key"
            :prop="'fields.'+index+'.value'"
            :rules="cmdField.rule"
            label-width="120px"
            require-asterisk-position="right"
        >
          <el-tooltip effect="light" :content="cmdField.desc" placement="bottom-start" v-if="(cmdField.choices.length > 0)" >
            <el-select placeholder="可选项" v-model="cmdField.value">
              <el-option v-for="option in cmdField.choices" :value="option"></el-option>
            </el-select>
          </el-tooltip>
          <el-input v-model="cmdField.value" :placeholder="cmdField.desc" v-else>
          </el-input>
        </el-form-item>
        <el-form-item>
          <el-row :gutter="20">
            <el-col :span="12">
              <el-button @click="submitExecForm(commandFormRef)" size="large" type="success">执 行</el-button>
            </el-col>
            <el-col :span="12">
              <el-button  @click="likeExecForm(commandFormRef)" size="large" type="primary">收 藏</el-button>
            </el-col>
          </el-row>
        </el-form-item>
      </el-form>
    </el-scrollbar>

    <el-divider content-position="left">执行结果</el-divider>
    <el-input v-model="commandExecResult" :rows=3 type="textarea"></el-input>

  </div>
</template>

<script lang="ts">
import {defineComponent, ref} from "vue";
import {ElNotification, FormInstance} from "element-plus";
import {execCmdServerCommand, likeCommand} from "@/requests/gmtool";
import {Operation} from "@element-plus/icons-vue";

export default defineComponent({
  components: {Operation},
  props: ['project', 'curEnv', 'commandserver', 'command'],
  setup(props) {
    const project = props.project
    const curEnv = props.curEnv
    const commandserver = props.commandserver
    const command = props.command
    const commandFormRef = ref<FormInstance>()
    const commandFormData = ref({
      fields: [],
    })
    const commandExecResult = ref('')

    console.log("cur command:", command, command.action)

    for (let i=0; i<command.fields.length; i++) {
      const fieldData = {
        key: command.fields[i].name,
        fieldType: command.fields[i]['type'],
        value: '',
        choices: [],
        rule: "{}",
        desc: command.fields[i].help_text,
      }

      if (command.fields[i].choices && command.fields[i].choices.length > 0) {
        fieldData.choices = command.fields[i].choices
        fieldData["rule"] = {required: true, message: '必填', trigger: 'blur'}
      } else if (!command.fields[i].black) {
        fieldData["rule"] = {required: true, message: '必填', trigger: 'blur'}
      }

      commandFormData.value.fields.push(fieldData)
    }

    const submitExecForm = (formEl: FormInstance | undefined) => {
      if (!formEl) {
        console.log("formref null")
        return
      }
      // 对表单的内容进行验证
      formEl.validate((valid, fields) => {
        const execFormData = commandFormData.value
        if (valid) {
          commandExecResult.value = "执行中。。。"
          let fieldsValue = []
          for (let i=0; i < execFormData.fields.length; i++) {
            const field = execFormData.fields[i]
            fieldsValue.push({name: field.key, value: field.value})
          }
          let baseExecInfo = {
            name: command.name,
            fields: fieldsValue,
          }
          let execCommandInfo = {
            project: project,
            command_server_name: commandserver.name,
            env: curEnv,
            base_info: baseExecInfo,
          }
          console.log('exec params:', execCommandInfo, execFormData)
          execCmdServerCommand(execCommandInfo).then((res)=>{
            console.log("exec command resp:", res)
            commandExecResult.value = JSON.stringify(res)
          }, (err) => {
            console.log("request error:", err)
          })
        } else {
          console.log('error submit!', execFormData)
          return false
        }
      })
    }

    const likeExecForm = (formEl: FormInstance | undefined) => {
      if (!formEl) {
        console.log("formref null")
        return
      }
      // 对表单的内容进行验证
      formEl.validate((valid, fields) => {
        const execFormData = commandFormData.value
        if (valid) {
          let fieldsValue = []
          for (let i=0; i < execFormData.fields.length; i++) {
            const field = execFormData.fields[i]
            fieldsValue.push({name: field.key, value: field.value})
          }
          let baseExecInfo = {
            name: command.name,
            fields: fieldsValue,
          }
          let execCommandInfo = {
            project: project,
            command_server_name: commandserver.name,
            env: curEnv,
            base_info: baseExecInfo,
          }
          likeCommand(execCommandInfo).then((res)=>{
            console.log("exec command resp:", res)
            ElNotification({
              title: "收藏结果通知",
              message: "收藏指令[" + command.name + "]成功！",
              type: 'success',
              duration: 4000,
              "show-close": true,
            })
            // commandExecResult.value = JSON.stringify(res)
          }, (err) => {
            console.log("request error:", err)
          })
        } else {
          console.log('error submit!', execFormData)
          return false
        }
      })
    }

    return {commandserver, command, commandFormRef, commandFormData, commandExecResult, submitExecForm, likeExecForm}
  }
})
</script>

<style scoped>

.breadcrumb {
  font-family: '微软雅黑';
  font-size: 50px;
}

</style>