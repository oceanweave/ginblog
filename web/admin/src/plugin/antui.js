import Vue from 'vue'
import { Button,FormModel,Input,Icon,message } from 'ant-design-vue'

message.config({
    top: `100px`,
    duration: 2,
    maxCount: 3
})

// 全局提示消息 因此要引入为全局变量
Vue.prototype.$message = message

Vue.use(Button)
Vue.use(FormModel)
Vue.use(Input)
Vue.use(Icon)