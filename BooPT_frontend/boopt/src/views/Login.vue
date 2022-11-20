<template>
    <div id="loginBackground">
        <el-container>
            <el-main>
                <img src="../assets/minilogo.png"></img>
                <el-tabs v-model="activeName" @tab-click="handleClick">
                    <el-tab-pane label="登录" name="logTab">
                        <el-input id = "usrInput" v-model="usrInput" placeholder="学号/电子邮箱"></el-input>
                        <el-input id = "pwdInput" v-model="pwdInput" placeholder="密码" show-password></el-input>
                        <el-button id = "logButton" v-model="logButton" @click="logEvent" type="success" size = "medium" icon = "el-icon-s-custom" plain>登录</el-button>
                    </el-tab-pane>
                    <el-tab-pane label="注册" name="regTab">
                        <el-input id = "stuidInput" v-model="stuidInput" placeholder="学号（LDAP用户名）"></el-input>
                        <el-input id = "ldappwdInput" v-model="ldappwdInput" placeholder="认证码（LDAP用户密码）" show-password></el-input>
                        <el-input id = "emailInput" v-model="emailInput" placeholder="电子邮件" ></el-input>
                        <el-input id = "usrInput" v-model="nameInput" placeholder="用户名" ></el-input>
                        <el-input id = "pwdInput" v-model="pwdInput2" placeholder="设置密码" show-password></el-input>
                        <el-button id = "regButton" @click="regEvent" type="primary" size = "medium" icon = "el-icon-circle-plus" plain>注册</el-button>
                    </el-tab-pane>
                </el-tabs>
            </el-main>
      </el-container>
    </div>
</template>

<script>
  export default {
    name: "Login.vue",
    data() {
      return {
        usrInput:"2021211315",
        pwdInput:"yigepassword",
        stuidInput:"",
        ldappwdInput:"",
        emailInput:"",
        nameInput:"",
        pwdInput2:"",
        activeName: "logTab"
      }
    },
    methods:{
      logEvent(){
        const self = this
        this.$axios({
          method: 'post',
          url: 'http://localhost:3000/api/account/login',
          params: {
            user: this.usrInput,
            password: this.pwdInput
          }
        })
        .then(function (response) {
          console.log(response);
          window.sessionStorage.setItem('token', response.data.token);
          window.sessionStorage.setItem('stuid', self.usrInput);
          self.$message({
            message: '登录成功！',
            type: 'success'
          });
          self.$router.push({
            path: '/Homepage',
            params: response.data
          })
        })
        .catch(function (error) {
            console.log(error);
            if(error.response.data.error === "Invalid Student ID or password") {
              self.$message.error("学号或密码错误");
            }
            else if(error.response.data.error === "Missing username or password") {
              self.$message.error("请输入用户名或密码");
            }
            else if(error.response.data.error === "Invalid email or password") {
              self.$message.error("邮箱或密码错误");
            }
            else {
              self.$message.error("内部服务错误");
            }
        });
      },
      regEvent(){
        const self = this
        this.$axios({
          method: 'post',
          url: 'http://localhost:3000/api/account/register',
          params: {
            stu_id: this.stuidInput,
            ldap_password: this.ldappwdInput,
            email: this.emailInput,
            name: this.nameInput,
            password: this.pwdInput2
          }
        })
        .then(function (response) {
            console.log(response);
            self.$message({
            message: '注册成功！',
            type: 'success'
          });
          
        })
        .catch(function (error) {
            console.log(error);
            if(error.response.data.error === "Invalid Student ID") {
              self.$message.error("学号格式错误！请输入数字格式的学号");
            }
            else if(error.response.data.error === "Invalid Student ID or password") {
              self.$message.error("学号不存在或认证密码错误");
            }
            else if(error.response.data.error === "Your student id has been used.") {
              self.$message.error("该学号已注册过账号");
            }
            else {
              self.$message.error("内部服务错误");
            }
        });
      }
    }
  }
</script>

<style>

  #loginBackground {
    background:url("../assets/background.png");
    width:100%;	
    height:100%;	
    position:fixed;
    background-size:100% 
  }
  .el-main {
    background-color: #E9EEF3;
    text-align: center;
    position: absolute;
    width: 24%;
    top: 20%;
    left: 38%;
    color: #333;
  }
  #usrInput {
    margin-right: 35px;
    margin-top: 20px;
  }
  #pwdInput {
    margin-right: 35px;
    margin-top: 20px;
  }
  #stuidInput{
    margin-right: 35px;
    margin-top: 20px;
  }
  #logButton{
    margin-top: 20px;
  }
  #regButton{
    margin-top: 20px;
  }
  #ldappwdInput{
    margin-top: 20px;
  }
  #emailInput{
    margin-top: 20px;
  }
</style>