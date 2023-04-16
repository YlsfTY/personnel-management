<template lang="">
  <div id="avatar" @click="TestClick">
    <n-avatar round>
      U
    </n-avatar>
    <span class="name">username你好</span>
  </div>
</template>
<script lang="ts">
import { defineComponent } from 'vue';
import { NAvatar } from 'naive-ui';
export default defineComponent({
  name: 'Avatar',
  components: {
    NAvatar
  },
  setup() {
    function TestClick() {
      console.log(123);
      
    }
  }
})
</script>
<style lang="scss" scoped>
#avatar {
  display: flex;
  align-items: center;
  span {
    line-height: 100%;
    &.name {
      margin: 0 10px;
    }
  }
}
</style>