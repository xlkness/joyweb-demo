<template>
  <div class="login-box">
    <div :class="{ active: currentModel, container: true,animate__animated :true, animate__flipInX :true}" >
      <div class="form-container sign-in-container">
        <el-form
            ref="ruleFormRef"
            :model="ruleForm"
            status-icon
            :rules="rules"
            class="form"
        >
          <el-form-item class="form-item" prop="username">
            <el-input v-model="ruleForm.username" placeholder="用户名" autocomplete="off" />
          </el-form-item>
          <el-form-item class="form-item" prop="password">
            <el-input
                v-model="ruleForm.password"
                placeholder="密码"
                type="password"
                autocomplete="off"
            />
          </el-form-item>
<!--            <el-form-item>-->
            <el-button class="theme-button" type="primary" @click="submitForm(ruleFormRef)" @keydown.enter="keyDown()">登 陆</el-button>
            <p class="forget"
                @click="
                currentModel = !currentModel;
                isForget = true;
            ">
              --- Forget Passwoed----
            </p>

<!--            </el-form-item>-->
        </el-form>
      </div>
      <!--覆盖容器-->
      <div class="overlay_container">
        <div class="overlay">
          <div class="overlay_panel overlay_right_container">
            <h2 class="container-title">hello friend!</h2>
            <p>输入您的个人信息，以便使用创乐汇后台管理系统</p>
            <button
                type="button"
                class="theme-button"
                id="sign-up"
                @click="
                currentModel = !currentModel;
                isForget = false;
                "
            >
              注册
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import {defineComponent, reactive, toRefs, ref } from "vue";
import { LoginData } from "../types/Login";
import type { FormInstance } from "element-plus";
import {login} from '../requests/user';
import {useRouter} from "vue-router";
import LocalCache from "@/stores/localCache";
import ExpireCache from "@/stores/expireCache";

export default defineComponent({
  mounted() {
    window.addEventListener("keydown", this.keyDown)
  },
  unmounted() {
    window.removeEventListener("keydown", this.keyDown, false)
  },
  setup() {
    const router=useRouter()
    const routeList=router.getRoutes().filter(v=>v.meta.menu_tab)
    const routeNameList: string[] = [];
    for (let i = 0; i < routeList.length; i++) {
      routeNameList.push(String(routeList[i].name))
    }
    const data=reactive(new LoginData())

    const rules={
      username: [
        {
          required: true,
          message: "请输入用户名",
          trigger: "blur",
        },
        {
          min: 3,
          max: 15,
          message: "用户名长度不合法，长度3-15",
          trigger: "blur",
        }
      ],
      password: [
        {
          required: true,
          message: "请输入密码",
          trigger: "blur",
        },
        {
          min: 3,
          max: 15,
          message: "密码长度不合法，长度3-15",
          trigger: "blur",
        }
      ]
    }

    // 登陆
    const ruleFormRef = ref<FormInstance>()
    const submitForm = (formEl: FormInstance | undefined) => {
      if (!formEl) return
      // 对表单的内容进行验证
      formEl.validate((valid) => {
        if (valid) {
          login({base_info: data.ruleForm}).then((res)=>{
            // 保存token，跳转
            // console.log('登陆成功，用户名：', res.payload.username)
            // console.log('登陆成功，token：', res.payload.token)
            LocalCache.setCache('token', res.payload.token)
            LocalCache.setCache('userInfo', res.payload)
            ExpireCache.setCache("token", res.payload.token, res.payload.expire)
            router.push('/home')
          }, (res) => {
            console.log("login response error:")
            console.log(res)
          })
        } else {
          console.log('error submit!')
          return false
        }
      })
    }

    const resetForm=()=>{
      data.ruleForm.username=""
      data.ruleForm.password=""
    }

    const keyDown = (e) => {
      if (e.keyCode === 13 || e.keyCode === 100) {
        submitForm(ruleFormRef.value)
      }
    }

    return {...toRefs(data), rules, resetForm, ruleFormRef, submitForm, keyDown};
  }
})


</script>

<style lang="scss" scoped>
  .login-box {
    width: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
    background-color: #f6f5f7;
    background-size: cover;
    background-image: url("@/assets/img/login/background.svg");
  }

  .container {
    position: relative;
    width: 768px;
    height: 480px;
    background-color: white;
    box-shadow: 0 14px 28px rgba(0, 0, 0, 0.25), 0 10px 10px rgba(0, 0, 0, 0.2);
    border-radius: 10px;
    overflow: hidden;
  }

  .form-container {
    position: absolute;
    top: 0;
    width: 50%;
    height: 100%;
    /*border: solid;*/
    background-color: white;
    transition: all 0.6s ease-in-out;
  }

  .form {
    display: flex;
    text-align: -webkit-center;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    height: 100%;
    padding: 0 50px;
    /*border: solid;*/
  }

  .form-item {
    width: 100%;
    /*border: solid;*/
    :deep(.el-input-group__append) {
      :deep(.el-button) {
        width: 60px;
      }
    };
  }

  :deep(.el-input__wrapper) {
    border-radius: 30px;
  }

  :deep(.el-input-group__append) {
    border-radius: 20px;
    margin-left: 10px;
    background: white;
    color: #fafafa;
  }

  .forget {
    font-size: 12px;
    color: rgba(50, 50, 50, 0.6);
    position: relative;
    top: 50px;
  }

  button:active {
    transform: scale(0.95);
  }

  .container.active .sign-up-container {
    transform: translateX(100%);
    z-index: 5;
  }

  .container.active .sign-in-container {
    transform: translateX(100%);
  }

  .container.active .overlay_container {
    transform: translateX(-100%);
  }

  .container.active .overlay {
    transform: translateX(50%);
  }

  .theme-button {
    background: rgba(64, 158, 255, 0.7);
    padding: 10px 50px;
    border: 1px solid transparent;
    border-radius: 20px;
    text-transform: uppercase;
    color: white;
    margin-top: 10px;
    outline: none;
    transition: transform 80;
  }

  .overlay_container {
    position: absolute;
    top: 0;
    width: 50%;
    height: 100%;
    z-index: 100;
    right: 0;
    overflow: hidden;
    transition: all 0.6s ease-in-out;
  }

  .overlay {
    position: absolute;
    width: 200%;
    height: 100%;
    left: -100%;
    background-color: #42b983;
  }

  .overlay_panel {
    position: absolute;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    width: 50%;
    height: 100%;
    color: white;
    padding: 0 0px;
    text-align: center;
  }

  .overlay_panel button {
    background-color: transparent;
    border: 1px solid white;
  }

  .overlay_panel p {
    font-size: 12px;
    margin: 10px 0 15px 0;
  }

  .overlay_right_container {
    right: 0;
  }
</style>