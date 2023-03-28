<template lang="">
  <div class="login">
    <!-- 注册 -->
    <div class="register-box"
      :class="{ 'slide-up':isLoginMode }" 
      ref="registerBox"
    >
      <h2 class="register-title"
        @click="toggleMode"
      >
        <span>没有账号，去</span>注册
      </h2>
      <div class="input-box">
        <input v-for="(input, index) in inputs" 
          :key="index"
          :type="input.type"
          :placeholder="input.placeholder"
          :minlength="input.minlength"
          :maxlength="input.maxlength"
          :class="{
            'error':input.isError,
            'input-shake':input.isShake
          }"
          v-model="data[input.name]"
          @blur="handleBlur(input)"
        >
      </div>
      <button @click="handleRegister">注册</button>
    </div>

    <!-- 登录 -->
    <div class="login-box"
      :class="{ 'slide-up':!isLoginMode }"
      ref="loginBox"
    >
      <div class="center">
        <h2 class="login-title"
          @click="toggleMode"
        >
          <span>已有账号，去</span>登录
        </h2>
        <div class="input-box">
          <input
            v-for="(input, index) in inputs" end="2" 
            :key="index"
            :type="input.type"
            :placeholder="input.placeholder"
            :minlength="input.minlength"
            :maxlength="input.maxlength"
            :class="{
              'error':input.isError,
              'input-shake':input.isShake
            }"
            v-model="data[input.name]"
            @blur="handleBlur(input)"
          >
        </div>
        <button @click="handleLogin">登录</button>
      </div>
    </div>

  </div>
</template>

<script lang="ts">
import { ref, reactive, watch, nextTick } from 'vue';
import { defineComponent } from 'vue';
import { userLogin, userRegister } from '../utils/request';

interface UserData {
  userName: string;
  userPassword: string;
  verifyPassword: string;
}

interface inputCell {
  type: "text" | "password";
  name: "userName" | "userPassword" | "verifyPassword";
  placeholder: string;
  isError: boolean;
  isShake: boolean;
  minlength: number;
  maxlength: number;
}

export default defineComponent({
  components: {},
  setup() {
    const data = reactive<UserData>({
      userName: '',
      userPassword: '',
      verifyPassword: '',
    });

    const inputs: inputCell[] = reactive([
      {
        type: "text",
        placeholder: "用户名",
        isError: false,
        isShake: false,
        name: "userName",
        minlength: 1,
        maxlength: 16,
      },
      {
        type: "password",
        placeholder: "密码",
        isError: false,
        isShake: false,
        name: "userPassword",
        minlength: 6,
        maxlength: 16,
      },
      {
        type: "password",
        placeholder: "重复密码",
        isError: false,
        isShake: false,
        name: "verifyPassword",
        minlength: 6,
        maxlength: 16,
      },

    ])

    const isLoginMode = ref(false);

    function toggleMode() {
      isLoginMode.value = !isLoginMode.value;
      inputs[0].placeholder = "用户名"
      inputs[1].placeholder = "密码"
      inputs[2].placeholder = "重复密码"
      resetErrors();
    }

    function resetErrors() {
      inputs.forEach(input => {
        data[input.name] = ''
        input.isError = false
        input.isShake = false
      })
    }
    watch([
      () => inputs[0].isShake,
      () => inputs[1].isShake,
      () => inputs[2].isShake,
    ],
      ([newIsShake_0, newIsShake_1, newIsShake_2]) => {
        if (newIsShake_0 === true) {
          setTimeout(() => {
            inputs[0].isShake = false
          }, 820)
        }
        if (newIsShake_1 === true) {
          setTimeout(() => {
            inputs[1].isShake = false
          }, 820)
        }
        if (newIsShake_2 === true) {
          setTimeout(() => {
            inputs[2].isShake = false
          }, 820)
        }
      },
    )


    function handleBlur(input: inputCell): void {
      const value = data[input.name]
      switch (input.name) {
        case 'userName': {
          if (value.length < input.minlength || value.length > input.maxlength) {
            input.placeholder = `请输入${input.minlength}到${input.maxlength}的用户名`
            input.isError = true
            input.isShake = true
            data[input.name] = ''
          } else {
            input.isError = false
            input.isShake = false
          }
          break
        }
        case 'userPassword': {
          if (value.length < input.minlength || value.length > input.maxlength) {
            input.placeholder = `请输入${input.minlength}到${input.maxlength}的密码`
            input.isError = true
            input.isShake = true
            data[input.name] = ''
          } else {
            input.isError = false
            input.isShake = false
          }
          break
        }
        case 'verifyPassword': {
          if (value != data['userPassword']) {
            input.placeholder = "两次密码不一致"
            input.isError = true
            input.isShake = true
            data[input.name] = ''
          } else if (value.length == 0) {
            input.placeholder = "请重复密码"
            input.isError = true
            input.isShake = true
          } else {
            input.isError = false
            input.isShake = false
          }
          break
        }
        default: {
          break;
        }
      }
    }
    function handleLogin() {
      for (let i = 0; i < 2; i++) {
        handleBlur(inputs[i])
        if (data[inputs[i].name].length == 0){
          return
        }
      }
      userLogin({
        userName: data.userName,
        userPassword: data.userPassword,
      }).then((res) => {
        console.log(res);
      });
    }

    function handleRegister() {
      inputs.forEach(input => {
        handleBlur(input)
        if (data[input.name].length == 0){
          return
        }
      })
      userRegister({
        userName: data.userName,
        userPassword: data.userPassword,
      }).then((res) => {
        console.log(res);
      });
    }

    return {
      data,
      inputs,
      isLoginMode,
      toggleMode,
      handleBlur,
      handleRegister,
      handleLogin
    }
  },
})
</script>

<style lang="scss">
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
      background: linear-gradient(rgba($color: #fff, $alpha: 0.5), rgba($color: #556e8f, $alpha: 0.1) 30%);
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

  .error {
    &::placeholder {
      color: red;
    }
  }

  @keyframes shake {

    10%,
    90% {
      transform: translate3d(-1px, 0, 0);
    }

    20%,
    80% {
      transform: translate3d(2px, 0, 0);
    }

    30%,
    50%,
    70% {
      transform: translate3d(-4px, 0, 0);
    }

    40%,
    60% {
      transform: translate3d(4px, 0, 0);
    }
  }

  .input-shake {
    animation: shake 0.82s cubic-bezier(0.36, 0.07, 0.19, 0.97) both;
  }
}
</style>