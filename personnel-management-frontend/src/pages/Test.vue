<script lang="ts">
  import { defineComponent, reactive, ref } from 'vue'
  import { NConfigProvider, NInput, NDatePicker, NSpace } from 'naive-ui'
  // theme
  import { createTheme, inputDark, datePickerDark } from 'naive-ui'
  // locale & dateLocale
  import { zhCN, dateZhCN } from 'naive-ui'

  export default defineComponent({
    components: {
      NConfigProvider,
      NInput,
      NDatePicker,
      NSpace
    },
    setup() {

      const model:{value:{[key:string]:number}} = ref({
        a:1,
        b:2,
        c:3
      })

      Object.keys(model.value).forEach((key)=>{
        console.log(model.value[key]);
      })
      Object.values(model.value).forEach((value:any)=>{
        console.log(value)
      })

      return {
        darkTheme: createTheme([inputDark, datePickerDark]),
        zhCN,
        dateZhCN
      }
    }
  })
</script>

<template>
  <n-config-provider :theme="darkTheme" :locale="zhCN" :date-locale="dateZhCN">
    <n-space vertical>
      <n-input />
      <n-date-picker />
    </n-space>
  </n-config-provider>
</template>

<style>
  body {
    background: black;
  }
</style>