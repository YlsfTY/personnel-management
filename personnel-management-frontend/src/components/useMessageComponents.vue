<script lang="ts">
  import { defineComponent,onUnmounted } from 'vue';
  import {useMessage} from 'naive-ui'

  export default defineComponent({
    setup(){
      // window.$message = useMessage()
      Reflect.defineProperty(window,"$message",{
      value:useMessage(),
      })
      onUnmounted(() => {
      Reflect.deleteProperty(window,"$message")
      })
    }
  })
</script>