<template>
  <n-upload 
    id="photo"
    @preview="handlePreview"
    :="photoProps"
  >
    点击或拖拽上传一寸照
  </n-upload>
  <n-modal 
    v-model:show="showModal"
    preset="card"
    style="width: 600px"
  >
    <img :src="previewImageUrl" style="width: 100%">
  </n-modal>
</template>

<script lang="ts">
import { ExtractPropTypes, defineComponent, ref, reactive } from 'vue'
import { UploadFileInfo, NUpload, NModal, } from 'naive-ui'

export default defineComponent({
  name: 'idPhoto',
  components: {
    NUpload,
    NModal
  },
  setup() {
    const showModalRef = ref(false)
    const previewImageUrlRef = ref('')

    const photoList: UploadFileInfo[] = reactive([
    ])

    const photoProps: ExtractPropTypes<typeof NUpload> = reactive({
      action: "http://127.0.0.1:8000/api/personnel/uploadImage",
      defaultFileList: photoList,
      listType: "image-card",
      max: 1,
      headers: {
        Authorization: 'Bearer ' + localStorage.getItem('token'),
      }
    })

    function handlePreview(file: UploadFileInfo) {
      const { url } = file
      previewImageUrlRef.value = url as string
      showModalRef.value = true
    }
    return {
      showModal: showModalRef,
      previewImageUrl: previewImageUrlRef,
      photoProps,
      handlePreview,
    }
  }
})

</script>

<style lang="scss">
#photo.n-upload {
  width: 100px;
  height: 140px;
  margin-right: 50px;
  margin-left: -50px;

  .n-upload-trigger {
    width: 100px;
    height: 140px;
  }

  .n-upload-file-list {
    display: flex;

    .n-upload-file {
      width: 100px;
      height: 140px;

      .n-upload-file-info {
        width: 100%;
        height: 100%;

        .n-image {
          width: 100%;
          height: 100%;
        }
      }
    }
  }
}
</style>
