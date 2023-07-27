<template>
  <component :is="userListComp" :subsystem="user"/>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import UserList from "@/components/user/UserList.vue";

export default defineComponent({
  setup() {
    const userListComp = UserList
    return {userListComp}
  }
})
</script>

<style scoped>

</style>