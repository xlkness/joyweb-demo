<template>
  <div>
    <el-row :gutter="20" style="margin-bottom: 20px">
      <el-col :span="10" :offset="6">
        <h4 style="font-size:70px">请选择执行的项目</h4>
      </el-col>
    </el-row>
    <el-row :gutter="20">
      <el-col :span="6" :offset="3">
        <el-button class="projectButton" color="#38b6ff" @click="handleSelectPro('足球')">
          <span style="font-size:80px">足球</span>
        </el-button>
      </el-col>
      <el-col :span="6">
        <el-button class="projectButton" color="#ffbd59" @click="handleSelectPro('SIM')">
          <span style="font-size:80px">SIM</span>
        </el-button>
      </el-col>
      <el-col :span="6">
        <el-button class="projectButton" color="#d9d9d9">
          敬请期待
        </el-button>
      </el-col>
    </el-row>
  </div>
</template>

<script lang="ts">
import {defineComponent, ref} from "vue";
import ProjectShowView from "@/views/gmtool/ProjectShowView.vue";

export default defineComponent({
  props: ['project', 'currentComponent',],
  setup(props, ctx) {
    const handleSelectPro = (pro: string) => {
      // console.log("select pro", pro)
      ctx.emit('update:project', pro)
      ctx.emit('update:currentComponent', ProjectShowView)
    }
    return {handleSelectPro}
  }
})
</script>

<style scoped>

.projectButton {
  min-width: 300px;
  min-height: 150px;
}

</style>