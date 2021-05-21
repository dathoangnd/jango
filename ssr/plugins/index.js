import Vue from 'vue'
import VueMaterial from "vue-material"
import Vuelidate from 'vuelidate'
import Editor from '@tinymce/tinymce-vue';

import 'vue-material/dist/vue-material.min.css'
import '~/assets/scss/theme.scss'

Vue.use(VueMaterial)
Vue.use(Vuelidate)
Vue.component('editor',Editor)