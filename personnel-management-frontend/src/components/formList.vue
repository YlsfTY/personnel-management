<template>
  <div id="formTest">
    <idPhoto :userName="$props.userName" />
    <n-form ref="formRef" :="formProps">
      <n-grid :="gridProps">
        <n-form-item-gi class="n-gi" v-for="(fItem, i) in fItemGis" :key="i" :="fItem.Props">
          <component v-if="fItem.sonForm" :is="fItem.sonForm.Type" :="fItem.sonForm.Props"
          v-model:value="data[fItem.Props.path]">
        </component>
        </n-form-item-gi>
        <n-gi class="n-gi" :span="24">
          <n-button @click="handlePersonnel">
            提交
          </n-button>
        </n-gi>
      </n-grid>
    </n-form>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref, ExtractPropTypes, watch, onMounted, } from 'vue'
import { NImage, NGrid, NGridItem, NGi, NForm, NFormItem, NFormItemGi, FormInst, } from 'naive-ui'
import { NInput, NInputNumber, NSelect, NDatePicker, NButton, NUpload, NUploadFileList, NUploadDragger, NUploadTrigger } from 'naive-ui'
import { personnel } from '@/types/Personnel'
import { FormItemRule } from 'naive-ui/lib'
import idPhoto from '@/components/idPhoto.vue'
import { getPersonnel, a_getPersonnel, a_createPersonnel } from '@/utils/request'
import { createPersonnel } from '@/utils/request'
import { useRouter } from 'vue-router'

type perKeyName = keyof personnel

interface FItemGiProps extends ExtractPropTypes<typeof NFormItemGi> {
  path: perKeyName
}

interface FItemGi {
  Props: FItemGiProps,
  sonForm: {
    Type: string,
    Props: {},
  } | null
}

export default defineComponent({
  name: 'FormTest',
  components: {
    NImage,
    NGrid,
    NGridItem,
    NGi,
    NForm,
    NFormItem,
    NFormItemGi,
    NInput,
    NSelect,
    NInputNumber,
    NDatePicker,
    NButton,
    NUpload,
    NUploadFileList,
    NUploadDragger,
    NUploadTrigger,
    idPhoto
  },

  props: {
    userName: {
      type: String,
      required: true
    }
  },

  setup(props) {

    const router = useRouter()

    const data: personnel = reactive({
      name: '',
      age: 0,
      phone: '',
      birthday: 0,
      IDNumber: '',
      residence: '',
      education: '',
      school: '',
      department: '',
      entryTime: 0,
      regularTime: 0,
      salary: 0,
      additional: '',
    })

    const formRef = ref<FormInst | null>(null)

    const formRules = {
      name: {
        required: true,
        // message: '请输入姓名',
        trigger: 'blur',
        validator: (rule: FormItemRule, value: string) => {
          if (value == '') {
            return new Error('请输入姓名')
          }
          const pattern = /^[\u4E00-\u9FA5]{2,4}$/;
          if (pattern.test(value)) {
            return true
          } else {
            return new Error('请输入2到4位的汉字')
          }
        }
      },
      phone: {
        required: true,
        // message: '请输入电话号码',
        trigger: 'blur',
        validator: (rule: FormItemRule, value: string) => {
          if (value == '') {
            return new Error('请输入电话号码')
          }
          const pattern = /^1[3-9]\d{9}$/
          if (pattern.test(value)) {
            return true
          } else {
            console.log('###')
            return new Error('电话号码格式错误')
          }
        }
      },
      age: {
        required: true,
        // message: '请输入年龄',
        trigger: ['input', 'blur'],
        validator: (rule: FormItemRule, value: number) => {
          if (!value) {
            return new Error('请输入年龄')
          }
          if (value >= 16) {
            return true
          } else {
            return new Error('年龄不符合规范')
          }
        }
      },
      birthday: {
        required: true,
        // message: '请输入生日',
        trigger: 'blur',
        validator: (rule: FormItemRule, value: number) => {
          if (!value) {
            return new Error('请输入生日')
          }
          if (new Date() < new Date(value)) {
            return new Error('出生日期不能超过今天')
          } else {
            return true
          }
        }
      },
      ID: {
        required: true,
        // message: '请输入身份证号',
        trigger: 'blur',
        validator: (rule: FormItemRule, value: string) => {
          if (!value) {
            return new Error('请输入身份证号')
          }
          const reg = /\d{17}(\d|X|x)$/
          if (!reg.test(value)) {
            return new Error('身份证号格式错误')
          }
          const weights = [7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2, 1]
          const codes = '10X98765432'
          let sum = 0
          for (let i = 0; i < 17; i++) {
            sum += parseInt(value.charAt(i)) * weights[i]
          }
          const index = sum % 11
          const code = codes.charAt(index)
          if (code !== value.charAt(17).toUpperCase()) {
            return new Error('身份证号格式错误')
          }
          return true
        }
      },
      residence: {
        required: true,
        message: '请输入居住地',
        trigger: 'blur',
      },
      educationBackground: {
        required: true,
        message: '请输入学历',
        trigger: 'blur',
      },
      graduateInstitutions: {
        required: true,
        message: '请输入毕业院校',
        trigger: 'blur',
      },
      department: {
        required: true,
        message: '请输入部门',
        trigger: 'blur',
      },
      entryTime: {
        required: true,
        // message: '请输入入职日期',
        trigger: 'blur',
        validator: (rule: FormItemRule, value: number) => {
          if (!value) {
            return new Error('请输入入职日期')
          }
          if (new Date() < new Date(value)) {
            return new Error('入职日期不能超过今天')
          } else {
            return true
          }
        }
      },
      regularTime: {
        required: false,
        // message: '请输入转正日期',
        trigger: 'blur',
        validator: (rule: FormItemRule, value: number) => {
          return true
        }
      },
      pay: {
        required: true,
        message: '请输入薪资',
        trigger: 'blur',
      },
      information: {
        required: false,
        message: '请输入补充信息',
        trigger: 'blur',
      },
    }

    const formProps: ExtractPropTypes<typeof NForm> = ({
      // 禁用
      disabled: false,
      // 行内
      inline: false,
      // 标签宽度
      // "label-width": undefined,
      labelWidth: undefined,
      // 文本对齐方式
      // "label-align": "right",
      labelAlign: "left",
      // 标签显示位置
      // "label-placement": "top",
      labelPlacement: "top",
      // 表单收集对象
      model: data,
      // 表单验证规则对象
      rules: formRules,
      // 校验反馈是否展示
      // "show-feedback": true,
      showFeedback: true,
      // 是否展示标签
      // "show-label": true,
      showLabel: true,
      // 是否显示必填*号
      // "show-require-mark": true,
      showRequireMark: true,
      // *号位置
      // "require-mark-placement": "right",
      RequireMarkPlacement: "right",
      // 尺寸
      size: "medium",
      // async-validator
      // "validate-messages": undefined
      ValidateMessages: undefined,
    })

    const gridProps: ExtractPropTypes<typeof NGrid> = ({
      cols: 24,
      xGap: 24,
      yGap: 15
    })

    const fItemGis: FItemGi[] = ([
      {
        Props: {
          span: 12,
          label: "姓名",
          path: "name",
        },
        sonForm: {
          Type: 'NInput',
          Props: {
            placeholder: "请输入姓名",
          }
        },
      },
      {
        Props: {
          span: 12,
          label: "电话号码",
          path: "phone",
        },
        sonForm: {
          Type: 'NInput',
          Props: {
            placeholder: "请输入电话号码",
          }
        },
      },
      {
        Props: {
          span: 12,
          label: "生日",
          path: "birthday",
        },
        sonForm: {
          Type: 'NDatePicker',
          Props: {
            type: 'date',
          }
        },
      },
      {
        Props: {
          span: 12,
          label: "年龄",
          path: "age",
        },
        sonForm: {
          Type: 'NInputNumber',
          Props: {
            placeholder: "请输入年龄",
            min: '16',
            max: '120',
            showButton: false,
          }
        },
      },
      {
        Props: {
          span: 12,
          label: "身份证号",
          path: "IDNumber",
        },
        sonForm: {
          Type: 'NInput',
          Props: {
            placeholder: "请输入身份证号",
          }
        },
      },
      {
        Props: {
          span: 12,
          label: "居住地",
          path: "residence",
        },
        sonForm: {
          Type: 'NInput',
          Props: {
            placeholder: "请输入居住地"
          }
        },
      },
      {
        Props: {
          span: 12,
          label: "学历",
          path: "education",
        },
        sonForm: {
          Type: 'NSelect',
          Props: {
            options: [
              {
                label: "点击选择学历",
                value: '',
                disabled: true
              },
              {
                label: "中学",
                value: "中学"
              },
              {
                label: "专科",
                value: "专科"
              },
              {
                label: "本科",
                value: "本科",
              },
              {
                label: "研究生",
                value: "研究生"
              },
            ]
          }
        },
      },
      {
        Props: {
          span: 12,
          label: "毕业学校",
          path: "school",
        },
        sonForm: {
          Type: 'NInput',
          Props: {
            placeholder: "请输入毕业学校"
          }
        },
      },
      {
        Props: {
          span: 12,
          label: "工作部门",
          path: "department",
        },
        sonForm: {
          Type: 'NSelect',
          Props: {
            options: [
              {
                label: "点击选择部门",
                value: '',
                disabled: true
              },
              {
                label: "A部门",
                value: "A部门",
              },
              {
                label: "B部门",
                value: "B部门",
              },
              {
                label: "C部门",
                value: "C部门",
              },
              {
                label: "D部门",
                value: "D部门",
              },
            ]
          }
        },
      },
      {
        Props: {
          span: 12,
          label: "入职日期",
          path: "entryTime",
        },
        sonForm: {
          Type: 'NDatePicker',
          Props: {
            type: 'date',
          }
        },
      },
      {
        Props: {
          span: 12,
          label: "转正日期",
          path: "regularTime",
          showRequireMark: false,
        },
        sonForm: {
          Type: 'NDatePicker',
          Props: {
            type: 'date',
          }
        },
      },
      {
        Props: {
          span: 12,
          label: "薪资",
          path: "salary",
        },
        sonForm: {
          Type: 'NInputNumber',
          Props: {
            placeholder: "请输入薪资",
            showButton: false,
          }
        },
      },
      {
        Props: {
          span: 24,
          label: "补充信息",
          path: "additional",
          showRequireMark: false,
        },
        sonForm: {
          Type: 'NInput',
          Props: {
            placeholder: "请输入补充信息",
            type: "textarea"
          }
        },
      },
    ])

    watch(() => data['age'], (newValue) => {
      if (newValue < 16) {
        return
      }
      const today = new Date()
      const todayYear = today.getFullYear()
      const birthYear = todayYear - newValue
      const inputBirthYear = new Date(data['birthday']).getFullYear()
      if (inputBirthYear && (inputBirthYear === birthYear - 1 || birthYear === inputBirthYear)) {
        return
      } else {
        data['birthday'] = (new Date(birthYear, 0, 1)).getTime()
        return
      }
    })

    watch(() => data['birthday'], (newValue) => {
      const today = new Date().getTime()
      const birthday = newValue
      const age = new Date(today - birthday).getFullYear() - 1970
      const inputAge = data['age']
      if (age !== inputAge) {
        data['age'] = age
      }
      return
    })

    function handlePersonnel(e: MouseEvent) {
      e.preventDefault()
      formRef.value?.validate((err) => {
        if (err) {
          window.$message.error('验证失败')
        } else {
          if (props.userName === sessionStorage.getItem('userName')) {
            createPersonnel(data).then((res: any) => {
              window.$message.success('提交成功')
              router.push({
                name: 'personnel',
                params: {
                  userName: props.userName
                }
              })
            }).catch(() => {
              window.$message.success('提交失败')
              return Promise.resolve()
            })
          } else {
            a_createPersonnel(data,props.userName).then((res: any) => {
              window.$message.success('提交成功')
              
            }).catch(() => {
              return Promise.resolve()
            })
          }
        }
      })
    }

    if (props.userName === sessionStorage.getItem('userName')) {
      getPersonnel().then((res: any) => {
        (Object.keys(data) as Array<perKeyName>)
          .forEach(key => {
            (data[key] as string | number) = (res.data.personnel as personnel)[key]
          });
      }).catch(() => {
        return Promise.resolve()
      })
    } else {
      a_getPersonnel(props.userName).then((res: any) => {
        (Object.keys(data) as Array<perKeyName>)
          .forEach(key => {
            (data[key] as string | number) = (res.data.personnel as personnel)[key]
          })
      }).catch(() => {
        return Promise.resolve()
      })
    }

    return {
      formRef,
      formProps,
      gridProps,
      formRules,
      fItemGis,
      data,
      handlePersonnel
    }
  },
})
</script>

<style scoped lang="scss">
#formTest {
  display: flex;
  justify-content: center;

  .n-date-picker,
  .n-input-number {
    width: 100%;
  }

  .n-button {
    margin: 0 50%;
    transform: translateX(-50%);
    color: #fff;
    background-color: rgba($color: #000, $alpha: 0.6);
  }
}
</style>
