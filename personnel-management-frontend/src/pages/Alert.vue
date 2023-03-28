<template>
  <div class="alert-container">
    <n-alert
      v-for="(alert, index) in alerts"
      :key="index"
      :title="alert.title"
      :type="alert.type"
      :showIcon="alert.showIcon"
      v-show=alert.show
      closable
    >
      {{ alert.content }}
    </n-alert>
  </div>
  <button @click="alertsTest">
    点击
  </button>
</template>

<script lang="ts">
import { useMessage } from 'naive-ui'
import { defineComponent } from 'vue'
import {ref,reactive} from 'vue'
import { NAlert, NIcon } from 'naive-ui'
import { IosAirplane } from '@vicons/ionicons4'
import test from 'node:test';

interface Alert {
  title: string;
  type: "default" | "error" | "info" | "success" | "warning";
  closable: boolean;
  showIcon: boolean;
  content: string;
  show?: boolean;
}

export default defineComponent({
  components: {
    NAlert,
    NIcon,
    IosAirplane,
  },
  setup(props) {
    const alerts: Alert[] = reactive([
    {
      title: 'Info 类型',
      type: 'info',
      closable: true,
      showIcon: true,
      content: 'Gee it\'s good to be back home',
      show:true
    },
    {
      title: 'Success 类型',
      type: 'success',
      closable: true,
      showIcon: true,
      content: 'Leave it till tomorrow to unpack my case',
      show:true
    },
    {
      title: 'Warning 类型',
      type: 'warning',
      closable: true,
      showIcon: true,
      content: 'Honey disconnect the phone',
      show:true
    },
    {
      title: 'Error 类型',
      type: 'error',
      closable: true,
      showIcon: true,
      content: 'I\'m back in the U.S.S.R.',
      show:true
    },
  ]);
  const message = useMessage()
  const alertsTest = ()=>{
    message.loading(
      '加载中'
    )
  }
  return {
    alerts,
    alertsTest
  };
  }
})
</script>

<style lang="scss" scoped>
.alert-container {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
}
button {
  margin-top: 300px;
} 
</style>