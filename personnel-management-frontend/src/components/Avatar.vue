<template lang="">
  <div id="avatar">
    <n-avatar 
      :src="imageUrl"
    round/>
    <span class="name">{{ userName }} 你好</span>
  </div>
</template>
<script lang="ts">
import { defineComponent, onMounted, ref } from 'vue';
import { NAvatar } from 'naive-ui';
import { getImage, userInfo } from '@/utils/request';
export default defineComponent({
  name: 'Avatar',
  components: {
    NAvatar
  },
  setup() {

    const userName = ref('');
    const imageUrl = ref('');

    onMounted(() => {
      userInfo().then((res: any) => {
        userName.value = res.data.user
      })

      getImage().then((res: any) => {
        imageUrl.value = res.data.image + '?_t' + new Date().getTime()
      }).catch(() => {
        imageUrl.value = 'http://127.0.0.1:5173/src/img/Avatar.png'
        return Promise.resolve()
      })
    })

    return {
      userName,
      imageUrl,
    }
  }
})
</script>
<style lang="scss" scoped>
#avatar {
  display: flex;
  align-items: center;
  position: relative;

  span {
    line-height: 100%;

    &.name {
      margin: 0 10px;
    }
  }
}</style>