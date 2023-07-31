<template>
  <div>
    <el-row :gutter="20">
      <el-col :span="2">
        <div>
          <el-icon color="#41c9c7" size="8vh" style="display: flex;align-items: center">
            <HomeFilled />
          </el-icon>
        </div>
      </el-col>
      <el-col :span="10" :offset="1">
        <h2 style="margin-top: 22px;color:#41c9c7">创乐汇后台工具管理系统</h2>
      </el-col>
      <el-col :span="5" :offset="6">
        <el-row>
          <el-col :span="8">
<!--            <el-icon size="8vh" color="#65d9d6">-->
<!--              <Help></Help>-->
<!--            </el-icon>-->
<!--            <el-avatar src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"></el-avatar>-->
            <el-avatar :icon="UserFilled" :size="60"></el-avatar>
          </el-col>
          <el-col :span="8">
            <span>{{username}}</span>
          </el-col>
          <el-col :span="8">

          </el-col>
        </el-row>
      </el-col>
    </el-row>
  </div>
</template>

<script lang="ts">
import {defineComponent} from "vue";
import {Help, HomeFilled, UserFilled} from "@element-plus/icons-vue";

export default defineComponent({
  computed: {
    UserFilled() {
      return UserFilled
    }
  },
  props: ["username"],
  setup(props) {
    const username = props.username || "guest"
    return {username}
  }
})
</script>

<style scoped>

</style>