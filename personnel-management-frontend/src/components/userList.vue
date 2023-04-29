<template>
  <n-data-table :="dataTableProps" :data="tableList">
  </n-data-table>
  <n-config-provider :theme="darkTheme">
    <n-drawer v-model:show="l_DrawerShow" placement="left" width="750px" class="left-drawer">
      <n-drawer-content>
        <template #header>
          <h3>{{ userName }}的职员信息</h3>
        </template>
        <personnel :userName="userName" />
      </n-drawer-content>
    </n-drawer>
    <n-drawer v-model:show="r_DrawerShow" placement="right" width="50%">
      <n-drawer-content>
        <template #header>
          <h3>编辑{{ userName }}的职员信息</h3>
        </template>
        <FormList :userName="userName" />
      </n-drawer-content>
    </n-drawer>
  </n-config-provider>

  <n-modal v-model:show="showModal" preset="dialog" title="确认" :content="`确认删除${userName}的职员信息?`" positive-text="确认"
    negative-text="取消" @positive-click="positiveClick" @negative-click="negativeClick">
  </n-modal>
</template>

<script lang="ts">
import { NButton, NConfigProvider, NDataTable, NDrawer, NDrawerContent, dataTableProps, darkTheme, NModal } from 'naive-ui';
import { ExtractPropTypes, defineComponent, h, markRaw, ref } from 'vue'
import { a_deletePersonnel, a_getListData } from '@/utils/request'
import { useRouter } from 'vue-router';
import Personnel from './Personnel.vue';
import FormList from './formList.vue';

export default defineComponent({
  name: 'userList',
  components: {
    NDataTable,
    NDrawer,
    Personnel,
    NDrawerContent,
    NConfigProvider,
    FormList,
    NModal
  },

  props:{
    
  },

  setup() {
    const router = useRouter()

    const l_DrawerShow = ref(false)
    const r_DrawerShow = ref(false)
    const showModal = ref(false)
    const userName = ref("")

    const dataTableProps: ExtractPropTypes<typeof NDataTable> = {
      columns: [
        {
          title: '用户',
          key: 'userName',
          align: 'center',
        },
        {
          title: '姓名',
          key: 'name',
          align: 'center',
        },
        {
          title: '部门',
          key: 'department',
          align: 'center',
        },
        {
          title: '操作',
          align: 'center',
          width: 200,
          render: (row: any) => {
            return h(
              'div',
              {
                class: 'button-item'
              },
              [
                h(
                  NButton,
                  {
                    type: 'primary',
                    style: {
                      marginRight: '4px'
                    },
                    onClick: () => {
                      l_DrawerShow.value = true
                      userName.value = row.userName
                    }
                  },
                  { default: () => '查看' }
                ),
                h(
                  NButton,
                  {
                    type: 'warning',
                    style: {
                      marginRight: '4px'
                    },
                    onClick: () => {
                      r_DrawerShow.value = true
                      userName.value = row.userName
                    }
                  },
                  { default: () => '编辑' }
                ),
                h(
                  NButton,
                  {
                    type: 'error',
                    onClick: () => {
                      userName.value = row.userName
                      showModal.value = true
                    }
                  },
                  { default: () => '删除' }
                )
              ]
            )
          }
        }
      ],
      // 斑马纹
      striped: true,
      // 右边框
      singleLine: false,
      pagination: {
        pageSize: 10
      }
    }

    const tableList = ref(markRaw([]))

    function setTableData() {
      a_getListData().then(res => {
        tableList.value = markRaw(res.data.dataList)
      }).catch(err => {
        return Promise.resolve()
      })
    }
    setTableData()

    const positiveClick = () => {
      a_deletePersonnel(userName.value).then(res => {
        window.$message.success("删除成功")
        setTableData()
      }).catch(err => {
        window.$message.error("删除失败")
        return Promise.resolve()
      })
    }

    const negativeClick = () => {
      return
    }

    return {
      dataTableProps,
      tableList,
      l_DrawerShow,
      r_DrawerShow,
      userName,
      darkTheme,
      showModal,
      positiveClick,
      negativeClick,
    }
  },
})
</script>

<style lang="scss">
.left-drawer {
  display: flex;

  #personnel {
    margin: 0 auto;

    .n-gi {
      color: #fff;
    }
  }
}
</style>
