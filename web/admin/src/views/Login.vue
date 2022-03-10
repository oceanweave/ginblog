<template>
    <div class="container">
        <div class="loginBox">
            <a-form-model ref="loginFormRef" :rules="rules" :model="formdata" class="loginForm">

                <a-form-model-item prop="username">
                    <a-input v-model="formdata.username" placeholder="请输入用户名">
                        <a-icon slot="prefix" type="user" style="color:rgba(0,0,0,.25)" />
                    </a-input>
                </a-form-model-item>

                <a-form-model-item prop="password">
                    <a-input v-model="formdata.password" placeholder="请输入密码" type="password"
                    v-on:keyup.enter="login">
                    <!--  v-on:keyup.enter="login" 输入回车相当于登陆 -->
                        <a-icon slot="prefix" type="lock" style="color:rgba(0,0,0,.25)" />
                    </a-input>
                </a-form-model-item>

                <a-form-model-item class="loginButton">
                    <a-button type="primary" style="margin: 10px" @click="login">登录</a-button>
                    <a-button type="info" style="margin: 10px" @click="resetForm" >取消</a-button>
                </a-form-model-item>

             </a-form-model>
        </div>      
    </div>
</template>

<script>
export default {
    data(){
        return {
          formdata: {
            username: '',
            password: ''
          },
          rules: {
            username: [
                { required: true, message: '请输入用户名', trigger: 'blur' },
                { min: 4, max: 12, message: '用户名必须在4-12个字符之间', trigger: 'blur' },
             ],
            password: [
                { required: true, message: '请输入密码', trigger: 'blur' },
                { min: 4, max: 12, message: '密码必须在6-12个字符之间', trigger: 'blur' },
            ],
          }
        }
    },
    methods:{
        resetForm(){
            // 清空表单 取消按键的作用
            this.$refs.loginFormRef.resetFields()
            // console.log(this.$refs)
        },
        login(){
            this.$refs.loginFormRef.validate(async (valid) => {
                if (!valid) {
                    return this.$message.error("输入非法数据，请重新输入")
                }
                // this.formdata 是从表单获取到的用户登录信息
                // res 是响应详细
                const res = await this.$http.post('login', this.formdata)
                // 此处需要考虑 后端设置的跨域
                // 打印日志，可以看到响应信息，及其结构
                // console.log(res)
                // 因此需要根据具体响应结构做调整，我的浏览器 层级是res.data.message 别的浏览器层级可能是 res.message
                if (res.data.status != 200 ) return this.$message.error(res.data.message)
                // sessionStorage 是关闭浏览器之后销毁，localStorage 一直存储在浏览器，需要手动清理
                // 存储 token
                window.sessionStorage.setItem('token',res.data.token)
                // 跳转到管理页面
                this.$router.push('admin')
            })
        }
    } 
}
</script>

<!-- // 只对本页面有效 -->
<style scoped>
.container{
    height: 100%;
    background-color:burlywood;
    /* display: flex;
    justify-content: center;
    align-items: center; */
}

.loginBox{
    width: 450px;
    height: 300px;
    background-color: aliceblue;
    /* // 居中 */
    position: absolute;
    top: 50%;
    left: 70%;
    transform: translate(-50%,-50%);
    border-radius: 9px; /*圆角*/
}

.loginForm{
    width: 100%;
    position: absolute;
    bottom: 30%;
    padding: 0 20px;
    box-sizing: border-box;
}

.loginButton{
    display: flex;
    justify-content: flex-end;
    position: absolute;
    /* bottom: 10%; */
    right: 5%;
}

</style>
