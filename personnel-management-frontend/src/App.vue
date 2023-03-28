<template>
  <n-message-provider>
    <useMessageComponent/>
    <router-view/>
  </n-message-provider>
</template>

<script lang="ts">
import {NMessageProvider} from 'naive-ui'
import {defineComponent} from 'vue'
import useMessageComponent from '@/components/useMessageComponents.vue'

export default defineComponent({
  name: 'App',
  components: {
    NMessageProvider,
    useMessageComponent
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