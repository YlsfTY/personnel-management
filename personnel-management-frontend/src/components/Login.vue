<template lang="">
  <div class="login">
    <!-- 注册 -->
    <div class="register-box"
      :class="{ 'slide-up':isLog }" 
      ref="registerBox"
    >
      <h2 class="register-title"
        @click="goRegister"
      >
        <span>没有账号，去</span>注册
      </h2>
      <div class="input-box">
        <input type="text" placeholder="用户名">
        <input type="password" placeholder="密码">
        <input type="password" placeholder="确认密码">
      </div>
      <button>注册</button>
    </div>

    <!-- 登录 -->
    <div class="login-box"
      :class="{ 'slide-up':!isLog }"
      ref="loginBox"
    >
      <div class="center">
        <h2 class="login-title"
          @click="goLogin"
        >
          <span>已有账号，去</span>登录
        </h2>
        <div class="input-box">
          <input type="text" placeholder="用户名">
          <input type="password" placeholder="密码">
        </div>
        <button>登录</button>
      </div>
    </div>

  </div>
</template>

<!-- <script>
import { ref,reactive, toRefs } from 'vue'

export default {
  name: 'Login',
  setup(props) {
    // 用户数据
    const user = reactive({
      userName: '',//用户名 
      userPassword: undefined,//密码
    })

    // 获取节点
    const registerTitle = ref()

    // 点击滑动事件
    function goRegister(event){
      console.log(registerTitle);
    }


    return {
      ...toRefs(user),
      registerTitle,
      goRegister

    }
  }
}
</script> -->
<script setup>
import { ref,reactive, toRefs } from 'vue'

  // 用户数据
  const {userName,userPassword} = reactive({
    userName: '',//用户名 
    userPassword: undefined,//密码
  })

  //控制滑动
  const isLog = ref(false) 

  // 点击滑动事件
  function goRegister(){
    if(isLog.value === true)
      isLog.value = false
  }

  function goLogin(){
    if(isLog.value === false)
      isLog.value = true
  }

</script>

<style lang="scss">
// * {
// margin: 0;
// padding: 0;
// }
.login {
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: rgba($color: #556e8f, $alpha: 0.7);
  border: 1px solid #556e8f;
  border-radius: 15px;
  height: 100%;
  overflow: hidden;
  position: relative;

  // 注册区，及一些通用样式
  .register-box {
    // display: none;

    width: 70%;
    position: absolute;
    z-index: 1;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    // border: 1px solid #000;
    transition: 0.3s ease;

  }

  .register-title,
  .login-title {
    color: #FFF;
    font-size: 27px;
    text-align: center;

    span {
      color: rgba($color: #000000, $alpha: 0.6);
      display: none;
    }
  }

  .input-box {
    background-color: #fff;
    border-radius: 15px;
    overflow: hidden;
    margin-top: 1%;
    visibility: visible;
    transition: 0.6s ease;
  }

  input {
    width: 100%;
    height: 30px;
    border: none;
    border-bottom: 1px solid rgba($color: #000000, $alpha: 0.1);
    font-size: 12px;
    padding: 8px 0;
    text-indent: 15px;
    outline: none;

    &:last-child {
      border-bottom: none;
    }

    &::placeholder {
      color: rgba($color: #000000, $alpha: 0.4);
    }
  }

  button {
    width: 100%;
    padding: 15px 45px;
    margin: 15px 0px;
    background-color: rgba($color: #000000, $alpha: 0.4);
    border: none;
    border-radius: 15px;
    color: rgba($color: #fff, $alpha: 0.8);
    font-size: 13px;
    font-weight: bold;
    cursor: pointer;
    opacity: 1;
    visibility: visible;
    transition: 0.3 ease;

    &:hover {
      background-color: rgba($color: #000000, $alpha: 0.8);
    }
  }

  // 登录区
  .login-box {
    // display: none;

    position: absolute;
    inset: 0;
    top: 20%;
    z-index: 2;
    // background-color: #556e8f;

    transition: 0.3s ease;

    &::before {
      content: "";
      // background-color: #455c77;
      background: linear-gradient(rgba($color: #fff, $alpha: 0.5),rgba($color: #556e8f, $alpha: 0.1) 30%);
      width: 200%;
      height: 250px;
      border-radius: 50%;
      position: absolute;
      top: -20px;
      left: 50%;
      transform: translateX(-50%);
    }

    .center {
      width: 70%;
      position: absolute;
      z-index: 3;
      left: 50%;
      top: 40%;
      transform: translate(-50%, -50%);

      .login-title {
        color: #000000;
      }

      .input-box {
        border: 1px solid rgba($color: #000000, $alpha: 0.1);
      }

      button {
        background-color: #233d54;
      }
    }
  }

  // 滑动收起
  .slide-up {
    // top: 90%;
    &.login-box {
      top: 90%;
    }

    &.register-box {
      top: 28%;
      // top: 17vh;
    }

    .center {
      top: 10%;
      transform: translate(-50%, 0%);
    }

    .login-title,
    .register-title {
      font-size: 16px;
      cursor: pointer;

      span {
        margin-right: 5px;
        display: inline-block;
      }
    }

    .input-box,
    button {
      opacity: 0;
      visibility: hidden;
    }
  }

}
</style>