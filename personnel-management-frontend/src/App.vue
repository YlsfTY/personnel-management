<template>
  <n-message-provider>
    <router-view/>
  </n-message-provider>
</template>

<script lang="ts">
import {NMessageProvider} from 'naive-ui'
import {defineComponent} from 'vue'

export default defineComponent({
  name: 'App',
  components: {
    NMessageProvider
  }
})
</script>

<style>
  *{
    margin: 0;
    padding: 0;
  }
  
  html,body {
    margin: 0;
    height: 100%;
    width: 100%;
  }
</style>