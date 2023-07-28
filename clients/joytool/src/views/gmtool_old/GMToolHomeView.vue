<template>
  <div class="home">
    <component
        :is="currentComponent"
        :project="project"
        @update:project="project=$event"
        :currentComponent="currentComponent"
        @update:currentComponent="currentComponent=$event"
    ></component>
  </div>
</template>

<script lang="ts">
import {defineComponent, ref, watch} from "vue";
import ProjectSelectView from "@/views/gmtool_old/ProjectSelectView.vue";

export default defineComponent({
  setup() {
    const currentComponent = ref(ProjectSelectView)
    const project = ref('')
    // console.log("project", project)
    return {currentComponent, project}
  }
})
</script>

<style scoped>
.home {
  width: 100%;
  height: 100%;
  background-color: white;
  background-size: cover;
}

</style>