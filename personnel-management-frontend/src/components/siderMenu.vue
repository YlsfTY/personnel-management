<template>
  <n-menu :options="menuOptions" id="siderMenu" />
</template>

<script lang="ts">
import { defineComponent, h } from 'vue'
import { RouterLink, useRouter } from 'vue-router'
import { NMenu } from 'naive-ui'
// import type { MenuOption } from 'naive-ui'


export default defineComponent({
  name: 'siderMenu',
  components: {
    NMenu
  },
  setup() {
    const router = useRouter()
    const menuOptions = [
      // {
      //   label: () =>
      //     h(
      //       RouterLink,
      //       {
      //         to: {
      //           name: 'home'
      //         }
      //       },
      //       // '主页'
      //       {
      //         default: () => '主页'
      //       }
      //     ),
      //   key: 'go-home',
      // },
      {
        label: () =>
          h(
            RouterLink,
            {
              to: {
                name: 'personnel'
              }
            },
            // '查看信息'
            {
              default: () => '查看信息'
            }
          ),
        key: 'go-personnel',
      },
      {
        label: () =>
          h(
            RouterLink,
            {
              to: {
                name: 'formList'
              }
            },
            // '编辑信息'
            {
              default: () => '编辑信息'
            }
          ),
        key: 'go-formList',
      },
      {
        label: () =>
          h(
            RouterLink,
            {
              to: {
                name: 'index'
              },
              onClick: (e: MouseEvent) => {
                localStorage.removeItem('token')
              },
            },
            {
              default: () => '退出登录'
            }
          ),
        key: 'go-login',
      }
    ]
    return {
      menuOptions
    }
  },
})
</script>

<style lang="scss">
#siderMenu.n-menu {
  background-color: rgba($color: #657487, $alpha: 0.5);

  .n-menu-item a {
    font-weight: bold;
    color: #000;
  }
}
</style>
