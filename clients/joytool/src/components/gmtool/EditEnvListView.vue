<template>
  <div>
    <el-container>
      <el-header>
        <el-button @click="dialogAddVisible = true" size="large" type="primary">添加环境</el-button>
      </el-header>
      <el-main>
        <el-table :data="envList" :default-sort="{prop:'index', order: 'ascending'}"  table-layout="auto" stripe style="width: 98vh;height: 60vh">
          <el-table-column prop="index" label="序 号">
          </el-table-column>
          <el-table-column prop="name" label="名 字">
          </el-table-column>
          <el-table-column prop="desc" label="描 述">
          </el-table-column>
          <el-table-column prop="created_at" label="创建时间">
          </el-table-column>
          <el-table-column prop="func" label="功 能">
            <template #default="scope">
              <el-button @click="handleEdit('operation', scope.$index, scope.row)" size="large" type="success">
                <el-icon style="vertical-align: middle">
                  <Operation/>
                </el-icon>
                <span>编辑</span>
              </el-button>
              <el-button  @click="handleDelete('operation', scope.$index, scope.row)" size="large" type="danger">
                <el-icon style="vertical-align: middle">
                  <Operation/>
                </el-icon>
                <span>删除</span>
              </el-button>
            </template>
          </el-table-column>
        </el-table>

        <el-dialog v-model="dialogAddVisible" title="添加环境" modal="true" :before-close="handleCloseDialog" destroy-on-close>
          <el-form
              ref="dialogAddFormRef"
              :model="dialogAddFormData"
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
                label="序号"
                prop="index"
                :rules="{required: true, message: '必填', trigger: 'blur'}"
            >
              <el-input v-model="dialogAddFormData.index" type="number"></el-input>
            </el-form-item>
          </el-form>
          <el-button @click="submitAddEnv(dialogAddFormRef)" type="primary" size="large">提交</el-button>
        </el-dialog>

        <el-dialog v-model="dialogEditVisible" title="编辑环境" modal="true" :before-close="handleCloseDialog" destroy-on-close>
          <el-form
              ref="dialogEditFormRef"
              :model="dialogEditFormData"
          >
            <el-form-item
                label="名字"
                prop="name"
                :rules="{required: true, message: '必填', trigger: 'blur'}"
            >
              <el-input v-model="dialogEditFormData.name" :placeholder="dialogEditFormData.oldData.name" disabled="true"></el-input>
            </el-form-item>
            <el-form-item
                label="描述"
                prop="desc"
            >
              <el-input v-model="dialogEditFormData.desc" :placeholder="dialogEditFormData.oldData.desc"></el-input>
            </el-form-item>
            <el-form-item
                label="序号"
                prop="index"
                :rules="{required: true, message: '必填', trigger: 'blur'}"
            >
              <el-input v-model="dialogEditFormData.index" :placeholder="dialogEditFormData.oldData.index"  type="number"></el-input>
            </el-form-item>
          </el-form>
          <el-button type="success" size="large" @click="submitEditEnv(dialogEditFormRef, dialogEditFormData.oldData)">提交</el-button>
        </el-dialog>
      </el-main>
    </el-container>
  </div>
</template>

<script lang="ts">
import {defineComponent, ref} from "vue";
import {addCmdServer, addEnv, deleteEnv, editCmdServer, editEnv, gmtool} from "@/requests/gmtool";
import {ElNotification, FormInstance} from "element-plus";

export default defineComponent({
  props: ['project'],
  setup(props) {
    const project = props.project
    const dialogAddVisible = ref(false)
    const dialogEditVisible = ref(false)
    const dialogAddFormRef =  ref<FormInstance>()
    const dialogEditFormRef =  ref<FormInstance>()
    const dialogAddFormData = ref({
      project: project,
      name: '',
      desc: '',
      index: -1,
    })
    const dialogEditFormData = ref({
      oldIndex: -1,
      oldData: {},
      project: project,
      name: '',
      desc: '',
      index: -1,
    })
    const handleCloseDialog = (done: () => void) => {
      dialogAddVisible.value = false
      dialogEditVisible.value = false
    }

    var commandServerList = ref([])
    var envList = ref([])

    gmtool({project: project}).then((res)=> {
      // 获取到gmtool信息，刷新页面
      var commandServerListTmp = res.payload.command_server_list || []
      var envListTmp = res.payload.env_list || []

      for (let i=0; i < commandServerListTmp.length; i++) {
        commandServerList.value.push(commandServerListTmp[i])
      }
      for (let i=0; i < envListTmp.length; i++) {
        envList.value.push(envListTmp[i])
      }
    }, (err) => {
      console.log("request error:", err)
    })

    const handleEdit = (oper: string, index: number, row: Object) => {
      dialogEditFormData.value.oldData = row
      dialogEditFormData.value.oldIndex = index
      dialogEditFormData.value.name = row.name
      dialogEditFormData.value.desc = row.desc
      dialogEditVisible.value = true
    }

    const handleDelete = (oper: string, index: number, row: Object) => {
      deleteEnv({project: project, name: row.name}).then((res) => {
        ElNotification({
          title: "删除结果通知",
          message: "删除环境[" + row.name + "]成功！",
          type: 'success',
          duration: 4000,
          "show-close": true,
        })
        envList.value.splice(index, 1)
      })
    }

    function compare(a, b) {

      return 0
      if (a.index == b.index) {
        if (a.name < b.name) {
          return 1;
        }
      } else if (a.index > b.index) return 1

      return 0
    }

    const submitAddEnv = (formEl: FormInstance | undefined) => {
      if (!formEl) {
        console.log("formref null")
        return
      }
      // 对表单的内容进行验证
      formEl.validate((valid, fields) => {
        dialogAddFormData.value.index = parseInt(dialogAddFormData.value.index)
        const execFormData = dialogAddFormData.value
        if (valid) {
          // console.log("add command server valid:", execFormData)
          addEnv(execFormData).then((res) => {
            ElNotification({
              title: "添加结果通知",
              message: "添加环境[" + execFormData.name + "]成功！",
              type: 'success',
              duration: 4000,
              "show-close": true,
            })
            dialogAddVisible.value = false
            envList.value.push(res.payload)
            // envList.value.sort(compare)
          }, (err) => {
            console.log("request error:", err)
          })
        } else {
          console.log('error submit!', execFormData)
          return false
        }
      })
    }

    const submitEditEnv = (formEl: FormInstance | undefined, oldData) => {
      if (!formEl) {
        console.log("formref null")
        return
      }
      // 对表单的内容进行验证
      formEl.validate((valid, fields) => {
        dialogEditFormData.value['index'] = parseInt(dialogEditFormData.value['index'])
        const execFormData = dialogEditFormData.value
        if (valid) {
          editEnv(execFormData).then((res) => {
            ElNotification({
              title: "修改结果通知",
              message: "修改环境[" + execFormData.name + "]成功！",
              type: 'success',
              duration: 4000,
              "show-close": true,
            })
            // console.log("edit command server valid:", execFormData, oldData, res)
            envList.value[dialogEditFormData.value.oldIndex] = res.payload
            // envList.value.sort(compare)
            // commandServerList.value.splice(dialogEditFormData.value.oldCommandServerIndex, 0, res.payload)
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
      dialogAddVisible,
      dialogEditVisible,
      handleCloseDialog,
      dialogAddFormRef,
      dialogAddFormData,
      dialogEditFormRef,
      dialogEditFormData,
      submitAddEnv,
      submitEditEnv,
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
