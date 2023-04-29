<!-- 首页 -->
<template lang="">
  <n-space vertical size="large">
    <n-layout id="home">
      <n-layout-header>
        <div class="top-title">
          <h1>PERSONNEL-MANAGEMENT</h1>
          <span>-基于Vue+Gin的人事管理系统-</span>
        </div>
        <div class="top-user">
          <Avatar/>
        </div>
      </n-layout-header>
      <n-layout has-sider id="center">
        <n-layout-content content-style="padding: 24px;">
          <router-view>
          </router-view>
        </n-layout-content>
        <n-layout-sider
          content-style="padding: 24px;"
          :width="`100%`"
          style="max-width:200px;min-width:0px"
        >
        <Suspense>
          <siderMenu/>
        </Suspense>
        </n-layout-sider>
      </n-layout>
      <n-layout-footer>
        <h1>这是底部</h1>
      </n-layout-footer>
    </n-layout>
  </n-space>
</template>

<script lang="ts" >
import { defineComponent, Suspense } from 'vue';
import { NSpace, NLayout, NLayoutContent, NLayoutHeader, NLayoutFooter, NLayoutSider } from 'naive-ui'
import { RouterView } from 'vue-router';
import Avatar from '@/components/Avatar.vue';
import siderMenu from '@/components/siderMenu.vue';
export default defineComponent({
  name: 'Home',
  components: {
    NLayout,
    NSpace,
    NLayoutContent,
    NLayoutFooter,
    NLayoutSider,
    NLayoutHeader,
    RouterView,
    Avatar,
    siderMenu,
  },
  setup(props) {

  }
})
</script>

<style lang="scss" scoped>
.n-layout {
  background: rgba($color: #000000, $alpha: 0.0);

  // 整体的layout
  &#home {
    height: 100vh;
    background-image: url(../assets/img/bgc.png);
    background-color: rgba($color: #000000, $alpha: 0.3);

    .n-layout-header {
      display: flex;
      align-items: center;
      justify-content: space-between;
      background-color: rgba($color: #000000, $alpha: 0.6);
      padding: 24px;
      overflow: hidden;
      color: #FFF;

      .top-title {
        span {
          color: rgba($color: #fff, $alpha: 0.5);
        }
      }

      .top-user {
        height: 100%;
      }
    }
  }
  
  // 包含sider与Content的部分
  &#center {
    max-width: 1200px;
    margin: 0 auto;


    .n-layout-content {
      // width: 60vw;
      background: rgba($color: #d0beb4, $alpha: 0.4);
      color: #000;
      font-weight: 800;
      min-width: 375px;
      min-height: 667px;
    }

    .n-layout-sider {
      background: rgba(128, 128, 128, 0.3);
      color: #FFF;
      // background-color: red;
      width: 200px;
      overflow: auto;
    }
  }

}

.n-layout-footer {
  // background: rgba(128, 128, 128, 0.2);
  background-color: rgba($color: #000, $alpha: 0.6);
  padding: 24px;
  color: #fff;
  text-align: center;
}
</style>