<template>
  <el-empty description="当前组权限不足无法查看！" v-if="(permission == false)">
  </el-empty>
  <component :is="RealUserList"
             :subsystem="subsystem"
             :subsystemName="subsystemName"
             :subsystemGroups="subsystemGroups"
             v-else
  >
  </component>
</template>

<script lang="ts">
import {defineComponent, ref} from "vue";
import {permissionGroupList} from "@/requests/gmtool";
import UserList from "@/components/user/UserList.vue";
import {useRouter} from "vue-router";

export default defineComponent({
  setup() {
    const useRoute = useRouter()
    const project = useRoute.currentRoute.value.query.project
    const subsystemGroups = ref([])
    const permission = ref(false)

    permissionGroupList({project: project}).then((res)=>{
      // console.log('permission group list', res.payload)
      if (res.payload.permission_group_list && res.payload.permission_group_list.length > 0) {
        subsystemGroups.value = []
        res.payload.permission_group_list.forEach(function (value, index, obj) {
          subsystemGroups.value.push(value.name)
        })
      }
      if (res.payload.is_admin) {
        permission.value = true
      }
    }, (err) => {
      console.log("groupList return error", err)
    })

    const RealUserList = UserList
    const subsystem = project + "-gmtool"
    const subsystemName = project + "-GM管理系统"
    return {permission, project, subsystemGroups, RealUserList, subsystem, subsystemName}
  }
})
</script>

<style scoped>

</style>