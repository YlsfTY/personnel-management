<!-- 登录页 -->
<template lang="">
  <div id="index">
    <div id="index-top">
      <div id="welcome">
        <Welcome/>
      </div>
      <div id="login">
        <Login/>
      </div>
    </div>
    <div id="about">
      <h1>介绍</h1>
    </div>
  </div>

</template>

<script lang="ts">
import Login from '../components/Login.vue';
import Welcome from '../components/Welcome.vue'
import { defineComponent,} from 'vue'

export default defineComponent({
  name:"Index",
  components:{
    Welcome,
    Login,
  },
  setup(props) {
    
  } 
})
</script>

<style lang="scss" scoped>
  #index {
    display: flex;
    flex-direction: column;
    justify-content: center;
    // align-items: center;
    height: 100vh;
    background-image: url(../assets/img/bgc.png);
    background-color: rgba($color: #000000, $alpha: 0.3);
    
    #index-top {
      display: flex;
      flex-direction: row;
      justify-content: center;
      #welcome,#login{
        // width: 50vw;
        width: 575px;
        // height: 60vh;
        height: 601px;
      }
    }

    #about {
      margin: 0 auto;
    }
  }

</style>