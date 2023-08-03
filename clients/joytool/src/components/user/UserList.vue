<template>
  <div>
    <el-container>
      <el-header>
        <el-button @click="handleAddUser" size="large" type="primary">添加用户</el-button>
      </el-header>
      <el-main>
        <el-table :data="userList" table-layout="auto" stripe style="width: 98vh;height: 60vh">
          <el-table-column prop="username" label="名 字">
          </el-table-column>
          <el-table-column prop="isAdminStr" label="是否管理员">
          </el-table-column>
          <el-table-column prop="group" label="所属组">
          </el-table-column>
          <el-table-column prop="createdAt" label="创建时间">
          </el-table-column>
          <el-table-column prop="func" label="功 能">
            <template #default="scope">
              <el-button @click="handleEditUser(scope.$index, scope.row)" size="large" type="success" :disabled="scope.row.isAdmin && !scope.row.isMyInfo">
                <el-icon style="vertical-align: middle">
                  <Operation/>
                </el-icon>
                <span>编辑</span>
              </el-button>
              <el-button  @click="handleDeleteUser(scope.$index, scope.row)" size="large" type="danger" :disabled="scope.row.isAdmin">
                <el-icon style="vertical-align: middle">
                  <Operation/>
                </el-icon>
                <span>删除</span>
              </el-button>
            </template>
          </el-table-column>
        </el-table>

        <el-dialog v-model="dialogAddUserVisible" :title="subsystemName+'/添加用户'" modal="true" :before-close="handleCloseDialog" destroy-on-close>
          <template #header>
            <span v-if="dialogUserData.oldData && dialogUserData.oldData.username">{{subsystemName}}/修改用户</span>
            <span v-else>{{subsystemName}}/添加用户</span>
          </template>
          <component :is="dialogAddUser" :subsystem="subsystem" :subsystemName="subsystemName" :subsystemGroups="subsystemGroups"
                     :dialogUserData="dialogUserData" :userList="userList" @update:userList="userList=$event"
                     :dialogAddUserVisible="dialogAddUserVisible" @update:dialogAddUserVisible="dialogAddUserVisible=$event"></component>
        </el-dialog>

      </el-main>
    </el-container>
  </div>
</template>

<script lang="ts">
import {defineComponent, ref} from "vue";
import {createUser, deleteUser, listUsers} from "@/requests/user";
import {ElNotification, FormInstance} from "element-plus";
import {addEnv} from "@/requests/gmtool";
import OperateUser from "@/components/user/OperateUser.vue";
import LocalCache from "@/stores/localCache";

export default defineComponent({
  props: ['subsystem', 'subsystemName', 'subsystemGroups'],
  setup(props) {

    const subsystem = props.subsystem || "user"
    const subsystemName = props.subsystemName || "用户系统"
    const subsystemGroups = props.subsystemGroups || ["用户管理员", "普通用户"]
    // console.log("sub system:", subsystem, props.subsystem)
    // console.log("sub system name:", subsystemName, props.subsystemName)
    // console.log("groups:", subsystemGroups, props.subsystemGroups,)
    const userInfo = LocalCache.getCache("userInfo")

    const userList = ref([])
    listUsers({system: subsystem}).then((res) => {
      // console.log('userlist', res.payload)
      for (let i=0; i<res.payload.length; i++) {
        const curUser = res.payload[i]
        const row = {
          username: curUser.username,
          isAdmin: curUser.is_admin,
          isAdminStr : curUser.is_admin ? "是" : "否",
          group: curUser.group,
          isMyInfo: userInfo.username == curUser.username,
          createdAt: curUser.created_at,
        }
        // console.log("row", row)
        userList.value.push(row)
      }
    }, (err) => {
      console.log(err)
    })

    const dialogAddUser = OperateUser
    const dialogAddUserVisible = ref(false)
    const dialogUserData = ref({
      index: -1,
      oldData: {},
    })

    const handleAddUser = () => {
      dialogUserData.value.oldData = {}
      dialogAddUserVisible.value = true
    }
    const handleEditUser = (index, row) => {
      dialogAddUserVisible.value = true
      dialogUserData.value.index = index
      dialogUserData.value.oldData = row
    }
    const handleCloseDialog = () => {
      dialogAddUserVisible.value = false
      dialogUserData.value = {index: -1, oldData: {}}
    }
    const handleDeleteUser = (index, row) => {
      deleteUser({username: row.username}).then((res) => {
        ElNotification({
          title: "删除结果通知",
          message: "删除用户[" + row.username + "]成功！",
          type: 'success',
          duration: 4000,
          "show-close": true,
        })
        userList.value.splice(index, 1)
      })
    }

    return {subsystem, subsystemName, subsystemGroups, handleAddUser,
      handleCloseDialog, dialogAddUserVisible, handleEditUser, userList, dialogAddUser, dialogUserData, handleDeleteUser}
  }
})
</script>

<style scoped>

</style>