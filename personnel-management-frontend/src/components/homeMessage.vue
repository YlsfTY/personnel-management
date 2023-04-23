<template>
  <div>
    <h1>基本信息</h1>
    <button
      @click="testForm"
    >跳转Form</button>
    <button
      @click="testList"
    >跳转List</button>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import { useRouter } from 'vue-router';

export default defineComponent({
  name: 'homeMessage',
  setup() {
    const router = useRouter()
    function testForm(event:MouseEvent){
      router.push('formList')
    }
    function testList(event:MouseEvent){
      router.push('personnel')
    }
    return {
      testForm,
      testList
    }
  },
})
</script>

<style scoped>
</style>
