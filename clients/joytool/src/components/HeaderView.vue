<template>
  <div>
    <el-row :gutter="5" style="border-bottom: dashed 1px gray;display: flex;justify-content: center;align-items: center">
      <el-col :span="3" style="display: flex;justify-content: center">
        <div>
          <el-icon color="#41c9c7" size="8vh" style="display: flex;align-items: center">
            <HomeFilled />
          </el-icon>
        </div>
      </el-col>
      <el-col :span="15" :offset="0" style="display: flex;justify-content: center;align-items: center">
        <span style="font-size:55px; font-weight: bold; color: #445ebf">创 乐 汇</span> <span style="font-size:50px; color:#41c9c7">后 台 工 具 管 理 系 统</span>
      </el-col>
      <el-col :span="4" :offset="2">
        <el-row style="display: flex;justify-content: flex-end;">
          <el-col class="avatar" :span="24" style="display: flex;justify-content: flex-end">
            <el-avatar :icon="UserFilled" :src="defaultAvatar" :size="80" style="margin-right: 5px"></el-avatar>
            <el-dropdown class="dropdown" :hide-on-click="false" size="large" @command="handleClick">
              <span class="avatar-username">
                {{username}}<el-icon class="el-icon--right"><arrow-down/></el-icon>
              </span>
              <template #dropdown>
                <el-dropdown-menu>
                  <el-dropdown-item :icon="User" command="profile">个人中心</el-dropdown-item>
                  <el-dropdown-item :icon="SwitchButton" divided command="exit">退出登陆</el-dropdown-item>
                </el-dropdown-menu>
              </template>
            </el-dropdown>
          </el-col>
        </el-row>
      </el-col>
    </el-row>
  </div>
</template>

<script lang="ts">
import {defineComponent} from "vue";
import {ArrowDown, Help, HomeFilled, Plus, SwitchButton, User, UserFilled} from "@element-plus/icons-vue";
import defaultAvatar from '@/assets/img/default-avatar.gif';
import {logout} from "@/requests/user";
import {useRouter} from "vue-router";
import {ElNotification} from "element-plus";
import LocalCache from "@/stores/localCache";
import ExpireCache from "@/stores/expireCache";

export default defineComponent({
  components: {ArrowDown},
  computed: {
    SwitchButton() {
      return SwitchButton
    },
    User() {
      return User
    },
    Plus() {
      return Plus
    },
    UserFilled() {
      return UserFilled
    }
  },
  props: ["username"],
  setup(props) {
    const username = props.username || "guest"
    const router = useRouter()

    const handleClick = (command: string) => {
      if (command == "exit") {
        const userInfo = LocalCache.getCache('userInfo')
        logout().then((res) => {
          ElNotification({
            title: "退出登陆通知",
            message: "取消用户[" + userInfo.username + "]成功！",
            type: 'success',
            duration: 4000,
            "show-close": true,
          })
          LocalCache.deleteCache('token')
          LocalCache.deleteCache('userInfo')
          ExpireCache.deleteCache("token")
          router.push('/login')
        })
      }
    }

    return {username, defaultAvatar, handleClick}
  }
})
</script>

<style lang="scss" scoped>

:deep() {
  .avatar{
    cursor: pointer;
    position: relative;
    display: flex;
    align-items: center;
    justify-content: center;
  }
  .el-tooltip__trigger {
    cursor: pointer !important;
  }
  .avatar-username{
    font-size: 25px;
    //color: var(--el-color-primary);
    cursor: pointer !important;
    display: flex;
    align-items: center;
  }
}




</style>