<template>
  <div id="personnel">
    <n-image width="100" src="http://localhost:5173/src/assets/img/Avatar.png" />
    <n-grid :="gridProps">
      <n-gi class="n-gi" :span="12" v-for="(gItem, i) in gItems" :key="i">
        {{ gItem.label }} : {{ gItem.value }}
      </n-gi>
      <n-gi class="n-gi" :span="24">
        <n-button @click="handleToInput">
          修改
        </n-button>
      </n-gi>
    </n-grid>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref, ExtractPropTypes, watch, onMounted, } from 'vue'
import { NImage, NGrid, NGridItem, NGi,NButton } from 'naive-ui'
import { personnel } from '@/types/Personnel'
import { useRouter } from 'vue-router'

type gridItemList = {
  [key in keyof (personnel)]: {
    label: string,
    value: string | number
  }
}

export default defineComponent({
  name: 'Personnel',
  components: {
    NImage,
    NGrid,
    NGridItem,
    NGi,
    NButton
  },
  setup() {
    const router = useRouter()
    const gridProps: ExtractPropTypes<typeof NGrid> = {
      cols: 24,
      xGap: 24,
      yGap: 15
    }
    const gItems: gridItemList = {
      name: {
        label: '姓名',
        value: ''
      },
      age: {
        label: '年龄',
        value: ''
      },
      phone: {
        label: '电话号码',
        value: ''
      },
      birthday: {
        label: '生日',
        value: ''
      },
      IDNumber: {
        label: '身份证号',
        value: ''
      },
      residence: {
        label: '居住地',
        value: ''
      },
      education: {
        label: '学历',
        value: ''
      },
      school: {
        label: '毕业学校',
        value: ''
      },
      department: {
        label: '部门',
        value: ''
      },
      entryTime: {
        label: '入职日期',
        value: ''
      },
      regularTime: {
        label: '转正日期',
        value: ''
      },
      salary: {
        label: '薪资',
        value: ''
      },
      additional: {
        label: '其它信息',
        value: ''
      }
    }

    function handleToInput(e:MouseEvent){
      e.preventDefault()
      router.push({name:'formList'})
      
    }

    return {
      gridProps,
      gItems,
      handleToInput
    }
  },
})
</script>

<style scoped lang="scss">
#personnel {

  display: flex;
  justify-content: center;

  .n-image {
    display: block;
    margin-right: 50px;
    margin-left: 118px;
  }

  .n-gi {
    padding-bottom: 50px;
    color: #000;
    font: 800 16px/1.5 '黑体'
  }
  .n-button{
    margin: 0 50%;
    transform: translateX(-50%);
    color: #fff;
    background-color: rgba($color: #000, $alpha: 0.6);
  }
}
</style>
