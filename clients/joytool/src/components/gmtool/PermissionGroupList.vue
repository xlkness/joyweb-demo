<template>
  <el-container>
    <el-header>
      <el-button @click="dialogOperationVisible = true" size="large" type="primary">添加环境</el-button>
    </el-header>
    <el-main>
      <el-table :data="permissionGroupList" :default-sort="{prop:'index', order: 'ascending'}"  table-layout="auto" stripe style="width: 98vh;height: 60vh">
        <el-table-column prop="name" label="名 字">
        </el-table-column>
        <el-table-column prop="desc" label="描 述">
        </el-table-column>
        <el-table-column prop="created_at" label="创建时间">
        </el-table-column>
        <el-table-column prop="func" label="功 能">
          <template #default="scope">
            <el-button @click="handleEdit( scope.$index, scope.row)" size="large" type="success">
              <el-icon style="vertical-align: middle">
                <Operation/>
              </el-icon>
              <span>编辑</span>
            </el-button>
            <el-button  @click="handleDelete(scope.$index, scope.row)" size="large" type="danger">
              <el-icon style="vertical-align: middle">
                <Operation/>
              </el-icon>
              <span>删除</span>
            </el-button>
          </template>
        </el-table-column>
      </el-table>

      <el-dialog v-model="dialogOperationVisible" title="添加环境" modal="true" :before-close="handleCloseDialog" destroy-on-close>
        <el-form
            ref="dialogOperationFormRef"
            :model="dialogOperationFormData"
        >
          <el-form-item
              label="名字"
              prop="name"
              :rules="{required: true, message: '必填', trigger: 'blur'}"
          >
            <el-input v-model="dialogOperationFormData.name" :disabled="dialogOperationFormData.oldData.name"></el-input>
          </el-form-item>
          <el-form-item
              label="描述"
              prop="desc"
          >
            <el-input v-model="dialogOperationFormData.desc"></el-input>
          </el-form-item>
          <el-form-item label="权限">
            <el-scrollbar max-height="30vh" style="width: 400px">
              <el-tree ref="treeRef" :data="tree" show-checkbox node-key="id" default-expand-all :default-checked-keys="dialogOperationFormData.checkedKeys"
                       props="{children: 'children', label: 'label" @check-change="handleCheckChange">

              </el-tree>
            </el-scrollbar>
          </el-form-item>
        </el-form>

        <el-button @click="submitOperation(dialogOperationFormRef)" type="primary" size="large">提交</el-button>
      </el-dialog>
    </el-main>
  </el-container>
</template>

<script lang="ts">
import {defineComponent, ref} from "vue";
import {createPermissionGroup, deleteEnv, deletePermissionGroup, editPermissionGroup, gmtool} from "@/requests/gmtool";
import EnvView from "@/components/gmtool/EnvView.vue";
import {FormInstance, ElTree, ElNotification} from "element-plus";

export default defineComponent({
  props: ['project'],
  setup(props) {
    const project = props.project
    const permissionGroupList = ref([])

    const treeRef = ref<InstanceType<typeof ELTree>>()
    const tree = ref([])
    const duplicateRecords = {}
    const permissionsConfRecords = []

    gmtool({project: project}).then((res)=> {
      // 获取到gmtool信息，刷新页面
      // var commandServerList = res.payload.command_server_list || []
      // var envList = res.payload.env_list || []
      // var likeList = res.payload.like_list || []

      for (let i=0; res.payload.permission_group_list && i<res.payload.permission_group_list.length; i++) {
        const curPermissionGroup = res.payload.permission_group_list[i]
        permissionGroupList.value.push(curPermissionGroup)
      }

      let envIncrID = 1
      let cmdServerIncrID = 100
      let rwIncrID = 1000

      for (let i=0; res.payload.env_list && i< res.payload.env_list.length; i++) {
        const curEnv = res.payload.env_list[i]
        const curEnvTreeNode = {
          id: envIncrID++,
          label: curEnv.name,
          children: [],
        }
        for (let j=0; res.payload.command_server_list && j < res.payload.command_server_list.length; j++) {
          const curCmdServer = res.payload.command_server_list[j]
          if (curCmdServer.envs.filter((data) => {return data == curEnv.name}).length == 0) {
            continue
          }

          const curCmdServerActionReadNode = {id: rwIncrID++, label: "读", name: "read", env: curEnv.name, cmd_server: curCmdServer.name}
          const curCmdServerActionWriteNode = {id: rwIncrID++, label: "写", name: "write", env: curEnv.name, cmd_server: curCmdServer.name}

          permissionsConfRecords.push(curCmdServerActionReadNode)
          permissionsConfRecords.push(curCmdServerActionWriteNode)

          const curCmdServerTreeNode = {
            id: cmdServerIncrID++,
            label: curCmdServer.name,
            children: [curCmdServerActionReadNode, curCmdServerActionWriteNode]
          }

          let dupReadCount = duplicateRecords[curCmdServer.name + curCmdServerActionReadNode.label]
          if (!dupReadCount) {
            dupReadCount = []
          }
          dupReadCount.push(curCmdServerActionReadNode.id)
          duplicateRecords[curCmdServer.name + curCmdServerActionReadNode.label] = dupReadCount

          let dupWriteCount = duplicateRecords[curCmdServer.name + curCmdServerActionWriteNode.label]
          if (!dupWriteCount) {
            dupWriteCount = []
          }
          dupWriteCount.push(curCmdServerActionWriteNode.id)
          duplicateRecords[curCmdServer.name + curCmdServerActionWriteNode.label] = dupWriteCount

          curEnvTreeNode.children.push(curCmdServerTreeNode)
        }
        tree.value.push(curEnvTreeNode)
      }

      // console.log('tree', tree.value)

      // console.log("commandServerList", commandServerList, envList)
    }, (err) => {
      console.log("request error:", err)
    })

    const dialogOperationVisible = ref(false)
    const dialogOperationFormRef = ref<FormInstance>()
    const dialogOperationFormData = ref({
      name: "",
      desc: '',
      oldIndex: -1,
      oldData: {},
      checkedKeys: [],
    })

    const handleCloseDialog = (done: () => void) => {
      dialogOperationVisible.value = false
      dialogOperationFormData.value.oldData = {}
      dialogOperationFormData.value.checkedKeys = []
    }

    const handleEdit = (index, row) => {
      dialogOperationVisible.value = true
      dialogOperationFormData.value.name = row.name
      dialogOperationFormData.value.desc = row.desc
      dialogOperationFormData.value.oldIndex = index
      dialogOperationFormData.value.oldData = row

      for (let i=0; row.permissions && i < row.permissions.length; i++) {
        let curEditPermission = row.permissions[i]
        for (let j=0; j < permissionsConfRecords.length; j++) {
          let curConfPermission = permissionsConfRecords[j]
          if (curEditPermission.env == curConfPermission.env &&
            curEditPermission.command_server == curConfPermission.cmd_server &&
            curEditPermission.action == curConfPermission.name) {
            dialogOperationFormData.value.checkedKeys.push(curConfPermission.id)
          }
        }
      }
      dialogOperationFormData.value.checkedKeys
    }

    const handleDelete = (index, row) => {
      deletePermissionGroup({project: project, name: row.name}).then((res) => {
        ElNotification({
          title: "删除结果通知",
          message: "删除权限组[" + row.name + "]成功！",
          type: 'success',
          duration: 4000,
          "show-close": true,
        })
        permissionGroupList.value.splice(index, 1)
      })
    }

    const handleCheckChange = (node, isChecked, hasChildChecked) => {
      return
      if (node.id < 1000) {
        return
      }

      const dupCount = duplicateRecords[node.cmd_server + node.label]

      if (isChecked) {
        let allSelectedIDs = treeRef.value!.getCheckedKeys(false).filter((data) => {return data>=1000})
        // console.log("勾选", node.id, "当前选中id", allSelectedIDs)
        // 将同名不同环境的勾选中
        treeRef.value!.setCheckedKeys([...new Set([].concat(dupCount, allSelectedIDs))], false)
      } else {
        let allSelectedIDs = treeRef.value!.getCheckedKeys(false).filter((data) => {return data>=1000})
        // 将同名不同环境的不勾选
        // console.log("=========================")
        // console.log("去掉勾选", node.id)
        // console.log("所有相关id", dupCount)
        // console.log("当前选中id", allSelectedIDs)
        let newFilterdIDs = allSelectedIDs.filter((data) => {return dupCount.indexOf(data) == -1})
        // console.log("过滤去组后的ids", newFilterdIDs)
        treeRef.value!.setCheckedKeys(newFilterdIDs, false)
      }
    }

    const submitOperation = (formEl: FormInstance | undefined) => {
      if (!formEl) {
        console.log("formref null")
        return
      }
      // 对表单的内容进行验证
      formEl.validate((valid, fields) => {
        if (valid) {
          let allSelectedNode = new Map()
          treeRef.value!.getCheckedKeys(false).filter((data) => {return data>=1000}).
            forEach((key) => {
              let node = treeRef.value!.getNode(key).data
              let pushKey = node.cmd_server + node.label
              if (true || !allSelectedNode.has(pushKey)) {
                allSelectedNode.set(pushKey, {
                  "env": node.env,
                  "command_server": node.cmd_server,
                  "action": node.name,
                })
              }
            })

          let allSelectedNode1 = [...allSelectedNode.values()]

          const reqData = {
            project: project,
            name: dialogOperationFormData.value.name,
            desc: dialogOperationFormData.value.desc,
            permissions: allSelectedNode1,
          }

          if (!dialogOperationFormData.value.oldData.name) {
            createPermissionGroup(reqData).then((res) => {
              ElNotification({
                title: "添加结果通知",
                message: "添加权限组[" + dialogOperationFormData.value.name + "]成功！",
                type: 'success',
                duration: 4000,
                "show-close": true,
              })
              dialogOperationVisible.value = false
              permissionGroupList.value.push(res.payload)
            })
          } else {
            editPermissionGroup(reqData).then((res) => {
              ElNotification({
                title: "修改结果通知",
                message: "修改权限组[" + dialogOperationFormData.value.name + "]成功！",
                type: 'success',
                duration: 4000,
                "show-close": true,
              })
              dialogOperationVisible.value = false
              permissionGroupList.value[dialogOperationFormData.value.oldIndex] = res.payload
            })
          }

        }
      })
    }

    return {project, permissionGroupList, dialogOperationVisible, dialogOperationFormRef, dialogOperationFormData,
      handleCloseDialog, handleEdit, handleDelete, treeRef, tree, handleCheckChange, submitOperation}
  }
})
</script>

<style scoped>

</style>