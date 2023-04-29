<template>
  <n-menu :options="menuOptions" id="siderMenu" />
</template>

<script lang="ts">
import { defineComponent, h, ref } from 'vue'
import { RouterLink, useRouter } from 'vue-router'
import { NMenu } from 'naive-ui'
import { adminInfo } from '@/utils/request'

export default defineComponent({
  name: 'siderMenu',
  components: {
    NMenu
  },
  async setup() {
    const router = useRouter()

    let userListShow = false
    await adminInfo().then(res => {
      userListShow = true
    }).catch(err => {
      userListShow = false
      return Promise.resolve()
    })

    const menuOptions = [
      {
        label: () =>
          h(
            RouterLink,
            {
              to: {
                name: 'personnel',
                params: {
                  userName: sessionStorage.getItem('userName')
                }
              }
            },
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
                name: 'formList',
                path: '/formList',
                params: {
                  userName: sessionStorage.getItem('userName')
                }
              }
            },
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
                name: 'userList'
              },
            },
            {
              default: () => '管理用户'
            }
          ),
        key: 'go-userList',
        show: userListShow,
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
      },
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
