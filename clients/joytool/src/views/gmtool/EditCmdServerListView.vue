<template>
  <div>
    <el-empty description="当前组权限不足无法查看！" v-if="(permission == false)">
    </el-empty>
    <el-container v-else>
      <el-header>
        <el-button @click="dialogAddVisible = true" size="large" type="primary">添加指令服务器</el-button>
      </el-header>
      <el-main>
        <el-table :data="commandServerList" style="width: 100%"  table-layout="auto" stripe>
          <el-table-column prop="name" label="名 字">
          </el-table-column>
          <el-table-column prop="desc" label="描 述">
          </el-table-column>
          <el-table-column prop="addr" label="地 址">
          </el-table-column>
          <el-table-column prop="envs" label="绑定环境">
          </el-table-column>
          <el-table-column prop="created_at" label="创建时间" width="150">
          </el-table-column>
          <el-table-column prop="func" label="功 能">
            <template #default="scope">
              <el-button  size="large" type="success" @click="handleEdit('operation', scope.$index, scope.row)">
                <el-icon style="vertical-align: middle">
                  <Operation/>
                </el-icon>
                <span>编辑</span>
              </el-button>
              <el-button  size="large" type="danger" @click="handleDelete('operation', scope.$index, scope.row)">
                <el-icon style="vertical-align: middle">
                  <Operation/>
                </el-icon>
                <span>删除</span>
              </el-button>
            </template>
          </el-table-column>
        </el-table>

        <el-dialog v-model="dialogAddVisible" title="添加指令服务器" modal="true" :before-close="handleCloseDialog" destroy-on-close>
          <el-form
              ref="dialogAddFormRef"
              :model="dialogAddFormData"
              label-width="80px"
          >
            <el-form-item
                label="名字"
                prop="name"
                :rules="{required: true, message: '必填', trigger: 'blur'}"
            >
              <el-input v-model="dialogAddFormData.name"></el-input>
            </el-form-item>
            <el-form-item
                label="描述"
                prop="desc"
            >
              <el-input v-model="dialogAddFormData.desc"></el-input>
            </el-form-item>
            <el-form-item
                label="地址"
                prop="addr"
                :rules="{required: true, message: '必填', trigger: 'blur'}"
            >
              <el-input v-model="dialogAddFormData.addr">
                <template #prepend>http://</template>
              </el-input>
            </el-form-item>
            <el-form-item
                label="环境"
                prop="envs"
            >
<!--              <el-checkbox-group v-model="dialogAddFormData.envs" v-for="env in envList">-->
<!--                <el-checkbox :label="env.name"/>-->
<!--              </el-checkbox-group>-->
              <el-checkbox-group v-model="dialogAddFormData.envs" size="large">
                <el-checkbox-button v-for="env in envList" :key="env.name" :label="env.name">
                  {{env.name}}
                </el-checkbox-button>
              </el-checkbox-group>
            </el-form-item>
            <el-form-item>
              <el-button @click="submitAddCmdServer(dialogAddFormRef)" size="large" type="primary">提交</el-button>
            </el-form-item>
          </el-form>
        </el-dialog>

        <el-dialog v-model="dialogEditVisible" title="编辑指令服务器" modal="true" :before-close="handleCloseDialog" destroy-on-close>
          <el-form
              ref="dialogEditFormRef"
              :model="dialogEditFormData"
              label-width="80px"
          >
            <el-form-item
                label="名字"
                prop="name"
                :rules="{required: true, message: '必填', trigger: 'blur'}"
            >
              <el-input v-model="dialogEditFormData.name" :placeholder="dialogEditFormData.oldCommandServerData.name" disabled="true"></el-input>
            </el-form-item>
            <el-form-item
                label="描述"
                prop="desc"
            >
              <el-input v-model="dialogEditFormData.desc" :placeholder="dialogEditFormData.oldCommandServerData.desc"></el-input>
            </el-form-item>
            <el-form-item
                label="地址"
                prop="addr"
                :rules="{required: true, message: '必填', trigger: 'blur'}"
            >
              <el-input v-model="dialogEditFormData.addr" :placeholder="dialogEditFormData.oldCommandServerData.addr"></el-input>
            </el-form-item>
            <el-form-item
                label="环境"
                prop="envs"

            >
              <el-checkbox-group v-model="dialogEditFormData.envs" size="large">
                <el-checkbox-button v-for="env in envList" :key="env.name" :label="env.name">
                  {{env.name}}
                </el-checkbox-button>
              </el-checkbox-group>
            </el-form-item>
            <el-form-item>
              <el-button @click="submitEditCmdServer(dialogEditFormRef, dialogEditFormData.oldCommandServerData)" size="large" type="primary">提交</el-button>
            </el-form-item>
          </el-form>
        </el-dialog>
      </el-main>
    </el-container>
  </div>
</template>

<script lang="ts">
import {defineComponent, h, ref} from "vue";
import {addCmdServer, deleteCmdServer, editCmdServer, execCmdServerCommand, gmtool} from "@/requests/gmtool";
import {ElMessage, ElNotification, FormInstance} from "element-plus";
import {useRouter} from "vue-router";

export default defineComponent({
  setup() {
    const useRoute = useRouter()
    const project = useRoute.currentRoute.value.query.project
    const dialogAddVisible = ref(false)
    const dialogEditVisible = ref(false)
    const dialogAddFormRef =  ref<FormInstance>()
    const dialogEditFormRef =  ref<FormInstance>()
    let emptyAddFormData = {
      project: project,
      name: '',
      desc: '',
      addr: '',
      envs: [],
    }
    const dialogAddFormData = ref(emptyAddFormData)
    let emptyEditFormData = {
      oldCommandServerIndex: -1,
      oldCommandServerData: {},
      project: project,
      name: '',
      desc: '',
      addr: '',
      envs: [],
    }
    const dialogEditFormData = ref(emptyEditFormData)
    const handleCloseDialog = (done: () => void) => {
      dialogAddVisible.value = false
      dialogEditVisible.value = false
      dialogAddFormData.value = emptyAddFormData
      dialogAddFormData.value.envs = []
      dialogEditFormData.value = emptyEditFormData
      dialogAddFormData.value.envs = []
    }

    var commandServerList = ref([])
    var envList = ref([])

    const permission = ref(false)

    gmtool({project: project}).then((res)=> {
      // 获取到gmtool信息，刷新页面
      var commandServerListTmp = res.payload.command_server_list || []
      var envListTmp = res.payload.env_list || []
      var permissionListTmp = res.payload.permission

      for (let i=0; i < commandServerListTmp.length; i++) {
        commandServerList.value.push(commandServerListTmp[i])
      }
      for (let i=0; i < envListTmp.length; i++) {
        envList.value.push(envListTmp[i])
      }

      console.log('permission', res.payload)

      if (res.payload.is_admin) {
        permission.value = true
      }
    }, (err) => {
      console.log("request error:", err)
    })

    const handleEdit = (oper: string, index: number, row: Object) => {
      dialogEditFormData.value.oldCommandServerData = row
      dialogEditFormData.value.oldCommandServerIndex = index
      dialogEditFormData.value.name = row.name
      dialogEditFormData.value.desc = row.desc
      dialogEditFormData.value.addr = row.addr
      dialogEditFormData.value.envs = row.envs || []
      dialogEditVisible.value = true
    }

    const handleDelete = (oper: string, index: number, row: Object) => {
      deleteCmdServer({project: project, name: row.name}).then((res) => {
        ElNotification({
          title: "删除结果通知",
          message: "删除指令服务器[" + row.name + "]成功！",
          type: 'success',
          duration: 4000,
          "show-close": true,
        })
        commandServerList.value.splice(index, 1)
      })
    }

    const submitAddCmdServer = (formEl: FormInstance | undefined) => {
      if (!formEl) {
        console.log("formref null")
        return
      }
      // 对表单的内容进行验证
      formEl.validate((valid, fields) => {
        dialogAddFormData.value.addr = "http://" + dialogAddFormData.value.addr
        const execFormData = dialogAddFormData.value
        if (valid) {
          // console.log("add command server valid:", execFormData)
          addCmdServer(execFormData).then((res) => {
            ElNotification({
              title: "添加结果通知",
              message: "添加指令服务器[" + execFormData.name + "]成功！",
              type: 'success',
              duration: 4000,
              "show-close": true,
            })
            commandServerList.value.push(res.payload)
            dialogAddVisible.value = false
          }, (err) => {
            console.log("request error:", err)
          })
        } else {
          console.log('error submit!', execFormData)
          return false
        }
      })
    }

    const submitEditCmdServer = (formEl: FormInstance | undefined, oldData) => {
      if (!formEl) {
        console.log("formref null")
        return
      }
      // 对表单的内容进行验证
      formEl.validate((valid, fields) => {
        var execFormData = dialogEditFormData.value
        if (valid) {
          editCmdServer(execFormData).then((res) => {
            ElNotification({
              title: "修改结果通知",
              message: "修改指令服务器[" + execFormData.name + "]成功！",
              type: 'success',
              duration: 4000,
              "show-close": true,
            })
            console.log("edit command server valid:", execFormData, oldData, res)
            commandServerList.value[dialogEditFormData.value.oldCommandServerIndex] = res.payload
            dialogEditVisible.value = false
          }, (err) => {
            console.log("request error:", err)
          })
        } else {
          console.log('error submit!', execFormData)
          return false
        }
      })
    }

    return {
      permission,
      dialogAddVisible,
      dialogEditVisible,
      handleCloseDialog,
      dialogAddFormRef,
      dialogAddFormData,
      dialogEditFormRef,
      dialogEditFormData,
      submitAddCmdServer,
      submitEditCmdServer,
      commandServerList,
      envList,
      handleEdit,
      handleDelete,
    }
  }
})
</script>

<style scoped>

</style>