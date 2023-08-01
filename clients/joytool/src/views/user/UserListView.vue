<template>
  <el-empty description="您不是管理员！" v-if="(isAdmin == false)">
  </el-empty>
  <component :is="userListComp" :subsystem="user" v-else/>
</template>

<script lang="ts">
import {defineComponent, ref} from "vue";
import UserList from "@/components/user/UserList.vue";
import LocalCache from "@/stores/localCache";

export default defineComponent({
  setup() {
    const userListComp = UserList
    const userInfo = LocalCache.getCache("userInfo")
    const isAdmin = ref(false)
    if (userInfo.is_admin) {
      isAdmin.value = true
    }
    return {userListComp, isAdmin}
  }
})
</script>

<style scoped>

</style>