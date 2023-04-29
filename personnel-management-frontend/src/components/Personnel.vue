<template>
  <div id="personnel">
    <n-image :width="100" :height="140" :src="imageUrl" :lazy="true" />
    <n-grid :="gridProps">
      <n-gi class="n-gi" :span="12" v-for="(gItem, i) in gItems" :key="i">
        {{ gItem.label }} : {{ gItem.value }}
      </n-gi>
      <!-- <n-gi class="n-gi" :span="24">
        <n-button @click="handleToInput">
          修改
        </n-button>
      </n-gi> -->
    </n-grid>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref, ExtractPropTypes, watch, onMounted, } from 'vue'
import { NImage, NGrid, NGridItem, NGi, NButton } from 'naive-ui'
import { personnel } from '@/types/Personnel'
import { useRouter } from 'vue-router'
import { a_getImage, a_getPersonnel, getImage, getPersonnel } from '@/utils/request'
import { formateDate } from '@/utils/date'

type gridItemList = {
  [key in keyof (personnel)]: {
    label: string,
    value: string | number
  }
}

// type gItemKeys = keyof gridItemList

export default defineComponent({
  name: 'Personnel',
  components: {
    NImage,
    NGrid,
    NGridItem,
    NGi,
    NButton
  },
  props: {
    userName: {
      type: String,
      required: true
    }
  },
  setup(props) {
    const router = useRouter()

    const imageUrl = ref('')
    const gridProps: ExtractPropTypes<typeof NGrid> = reactive({
      cols: 24,
      xGap: 24,
      yGap: 15
    })
    const gItems: gridItemList = reactive({
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
    })

    function handleToInput(e: MouseEvent) {
      e.preventDefault()
      router.push({ name: 'formList' })
    }

    // onMounted(() => {
    //   getPersonnel().then((res: any) => {
    //     Object.keys(gItems).forEach((key: string) => {
    //       gItems[key as keyof gridItemList].value = res.data.personnel[key]
    //     })
    //     gItems['birthday'].value = formateDate(gItems['birthday'].value as number)
    //     gItems['entryTime'].value = formateDate(gItems['entryTime'].value as number)
    //     gItems['regularTime'].value = formateDate(gItems['regularTime'].value as number)
    //   }).catch(() => {
    //     return Promise.resolve()
    //   })
    //   getImage().then((res: any) => {
    //     imageUrl.value = res.data.image + '?_t' + new Date().getTime()
    //   }).catch(() => {
    //     imageUrl.value = "http://localhost:5173/src/assets/img/Avatar.png"
    //     return Promise.resolve()
    //   })
    // })

    console.log(props.userName);
    if (props.userName === sessionStorage.getItem('userName')) {
      
      getPersonnel().then((res: any) => {
        Object.keys(gItems).forEach((key: string) => {
          gItems[key as keyof gridItemList].value = res.data.personnel[key]
        })
        gItems['birthday'].value = formateDate(gItems['birthday'].value as number)
        gItems['entryTime'].value = formateDate(gItems['entryTime'].value as number)
        gItems['regularTime'].value = formateDate(gItems['regularTime'].value as number)
      }).catch(() => {
        return Promise.resolve()
      })
      getImage().then((res: any) => {
        imageUrl.value = res.data.image + '?_t' + new Date().getTime()
      }).catch(() => {
        return Promise.resolve()
      })
    } else {
      a_getPersonnel(props.userName).then((res: any) => {
        Object.keys(gItems).forEach((key: string) => {
          gItems[key as keyof gridItemList].value = res.data.personnel[key]
        })
        gItems['birthday'].value = formateDate(gItems['birthday'].value as number)
        gItems['entryTime'].value = formateDate(gItems['entryTime'].value as number)
        gItems['regularTime'].value = formateDate(gItems['regularTime'].value as number)
      }).catch(() => {
        return Promise.resolve()
      })
      a_getImage(props.userName).then((res: any) => {
        imageUrl.value = res.data.image + '?_t' + new Date().getTime()
      }).catch(() => {
        return Promise.resolve()
      })
    }

    // const mediaQuery = window.matchMedia('(max-width: 700px)')      
    // if (mediaQuery.matches) {
    //   gridProps.cols = 20
    // } else {
    //   gridProps.cols = 24
    // }

    return {
      gridProps,
      gItems,
      imageUrl,
      handleToInput
    }
  },
})
</script>

<style scoped lang="scss">
#personnel {

  display: flex;
  justify-content: center;
  margin-left: 50px;
  color: #000;
  font: 800 16px/1.5 '黑体';
  
  .n-image {
    // width: 100px;
    // height: 140px;
    display: block;
    margin-right: 50px;
    margin-left: 0px;
    // overflow: hidden;
    // img{
    // width: 100%;
    // }
  }


  .n-gi {
    padding-bottom: 50px;
  }

  .n-button {
    margin: 0 50%;
    transform: translateX(-50%);
    color: #fff;
    background-color: rgba($color: #000, $alpha: 0.6);
  }
}
</style>
