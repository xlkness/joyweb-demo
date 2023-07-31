<template>
  <div>
      {{project}}-test-{{this.$router.query}}
  </div>
</template>

<script lang="ts">
import {defineComponent} from "vue";
import {useRouter} from "vue-router";
import router from "@/router";

export default defineComponent({
  setup() {
    console.log("entry test")
    const useRoute = useRouter()
    // const project = this.$route.query.project
    const project = useRoute.currentRoute.value.query.project
    console.log('path', useRoute.currentRoute.value.fullPath)
    console.log('1', useRoute.currentRoute.value)
    console.log('2', router.currentRoute.value)
    return {project}
  }
})
</script>

<style scoped>

</style>