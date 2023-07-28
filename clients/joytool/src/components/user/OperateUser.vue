<template>
  <div>
    <el-form
        ref="dialogAddFormRef"
        :model="dialogAddFormData"
    >
      <el-form-item
          label="名字"
          prop="username"
          :rules="{required: true, message: '必填', trigger: 'blur'}"
      >
        <el-input v-model="dialogAddFormData.username" :disabled="dialogAddFormData.username"></el-input>
      </el-form-item>
      <el-form-item
          label="密码"
          prop="passwd"
          :rules="{required: true, message: '必填', trigger: 'blur'}"
      >
        <el-input type="password" v-model="dialogAddFormData.passwd"></el-input>
      </el-form-item>
      <el-form-item
          label="是否管理员"
      >
        <el-switch
            v-model="dialogAddFormData.isAdmin"
            size="large"
            class="ml-2"
            inline-prompt
            active-text="是"
            inactive-text="否"
        ></el-switch>
      </el-form-item>
      <el-form-item
          label="执行组"
      >
        <!--              <el-tree-->
        <!--                  :data="editPermissionData.permissions"-->
        <!--                  :show-checkbox="editPermissionData.isAdmin == false"-->
        <!--                  node-key="id"-->
        <!--                  :default-checked-keys="editPermissionData.checkedIds"-->
        <!--                  :props="{children: 'children', label: 'label'}"-->
        <!--              >-->
        <!--              </el-tree>-->
        <el-radio-group v-model="dialogAddFormData.group" size="large" :disabled="dialogAddFormData.isAdmin == true">
          <el-radio-button v-for="item in subsystemGroups" :label="item"></el-radio-button>
        </el-radio-group>
      </el-form-item>
      <el-form-item>
        <el-button type="success" size="large" @click="submitCreateUser(dialogAddFormRef)">提交</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script lang="ts">
import {defineComponent, ref} from "vue";
import {createUser, editUser, listUsers} from "@/requests/user";
import {ElNotification, FormInstance} from "element-plus";
import ProjectSelectView from "@/views/gmtool_old/ProjectSelectView.vue";

export default defineComponent({
  props: ['subsystem', 'subsystemName', 'subsystemGroups', 'dialogUserData', 'dialogAddUserVisible', 'userList'],
  setup(props, ctx) {
    const subsystem = props.subsystem || "user"
    const subsystemName = props.subsystemName || "用户系统"
    const subsystemGroups = props.subsystemGroups || ["用户管理员", "普通用户"]
    const dialogUserData = props.dialogUserData || {}
    const userList = props.userList

    const dialogAddVisible = ref(false)
    const dialogAddFormRef =  ref<FormInstance>()
    const dialogAddFormData = ref({
      username: dialogUserData.oldData.username || '',
      passwd: dialogUserData.oldData.password || '',
      isAdmin: false,
      group: '',
    })

    console.log('old data', dialogUserData.oldData)

    for (let i=0; dialogUserData.oldData.systems && i < dialogUserData.oldData.systems.length; i++) {
      const curSystem = dialogUserData.oldData.systems[i]
      if (curSystem.name == subsystem) {
        dialogAddFormData.value.group = curSystem.group
        dialogAddFormData.value.isAdmin = curSystem.is_admin
        break
      }
    }

    const submitCreateUser = (formEl: FormInstance | undefined) => {
      if (!formEl) {
        console.log("formref null")
        return
      }
      // 对表单的内容进行验证
      formEl.validate((valid, fields) => {
        if (valid) {
          const createRequest = {
            base_info: {
              username: dialogAddFormData.value.username,
              password: dialogAddFormData.value.passwd,
            },
            group_list: [{name: subsystem, is_admin: dialogAddFormData.value.isAdmin, group: dialogAddFormData.value.group}]
          }
          if (props.dialogUserData.oldData.username) {
            editUser(createRequest).then((res) => {
              ElNotification({
                title: "修改结果通知",
                message: "修改用户[" + dialogAddFormData.value.username + "]成功！",
                type: 'success',
                duration: 4000,
                "show-close": true,
              })
              dialogAddVisible.value = false
              // userList.value.push(res.payload)
              userList[props.dialogUserData.index] = res.payload
              ctx.emit('update:userList', userList)
              // envList.value.sort(compare)
            }, (err) => {
              console.log("request error:", err)
            })
          } else {
            createUser(createRequest).then((res) => {
              ElNotification({
                title: "添加结果通知",
                message: "添加用户[" + dialogAddFormData.value.username + "]成功！",
                type: 'success',
                duration: 4000,
                "show-close": true,
              })
              dialogAddVisible.value = false
              console.log("user list", userList)
              userList.push(res.payload)
              ctx.emit('update:userList', userList)
              // envList.value.sort(compare)
            }, (err) => {
              console.log("request error:", err)
            })
          }

          ctx.emit('update:dialogAddUserVisible', false)
        } else {
          console.log('error submit!', dialogAddFormData)
          return false
        }
      })
    }

    return {subsystemName, userList, dialogAddVisible, dialogAddFormRef, dialogAddFormData, subsystemGroups, submitCreateUser}
  }
})
</script>

<style scoped>

</style>