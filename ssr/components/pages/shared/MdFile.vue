<template>
  <div class="md-file">
    <md-icon :class="iconClass" @click.native="openPicker">cloud_upload</md-icon>

    <input
      class="md-input"
      v-model="model"
      v-bind="{ disabled, required, placeholder }"
      @focus="onFocus"
      @blur="onBlur"
      >

    <input type="file" ref="inputFile" v-bind="attributes" @change="onFileSelected" />
  </div>
</template>

<script>
  import MdUuid from '~/node_modules/vue-material/src/core/utils/MdUuid'
  import MdFieldMixin from '~/node_modules/vue-material/src/components/MdField/MdFieldMixin'
  export default {
    name: 'MdFile',
    props: {
      id: {
        type: String,
        default: () => 'md-file-' + MdUuid()
      },
      type: String,
      name: String
    },
    computed: {
      iconClass () {
        return {
          'md-disabled': this.disabled
        }
      }
    },
    mixins: [MdFieldMixin],
    inject: ['MdField'],
    methods: {
      openPicker () {
        this.$refs.inputFile.click()
      },
      onFileSelected ({ target, dataTransfer }) {
        const files = target.files || dataTransfer.files
        let formData = new FormData()
        formData.append('file', files[0])
        let that = this
        
        this.$axios({
          method: 'post',
          url: `${process.env.apiUrl}/${that.type == 'pdf' ? 'pdf' : 'image'}/upload`, 
          data: formData
        })
        .then(function(response) {
          that.model = `${process.env.apiUrl}/${response.data.location}`
        })
        .catch(function(error) {
          that.$store.commit('setSnackbarMessage', 'Đã xảy ra sự cố, vui lòng thử lại.')
          that.$store.commit('toggleSnackbar', true)
        })
      }
    },
    created () {
      this.MdField.file = true
    },
    beforeDestroy () {
      this.MdField.file = false
    }
  }
</script>

<style lang="scss">
  .md-file {
    display: flex;
    flex: 1;
    input[type="file"] {
      width: 1px;
      height: 1px;
      margin: -1px;
      padding: 0;
      overflow: hidden;
      position: absolute;
      clip: rect(0 0 0 0);
      border: 0;
    }
    .md-icon {
      position: relative;
      top: 5px;
      &:not(.md-disabled) {
        cursor: pointer;
      }
      &.md-disabled {
        pointer-events: none;
      }
    }
  }
</style>