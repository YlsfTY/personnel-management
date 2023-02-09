<template lang="">
  <div class="welcome">
    <div class="box">
      <span 
        class="text"
        v-for="(item, index) in sum" :key="index"
      >
        Welcome
      </span>
    </div>
  </div>
</template>

<script>
import { ref } from 'vue'

export default {
  name:'Welcome',
  setup() {
    const sum = ref(10)
    return {
      sum,
    }
  }

}
</script>

<style lang="scss" scoped>

  .welcome {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100%;
    .box {
      transform-style: preserve-3d;
      perspective: 800px;
      position: relative;
      animation: slogan 5s linear;
      animation-iteration-count: infinite;

      .text{
        position: absolute;
        font-size: 5vw;
        font-weight: 800;
        opacity: 0.8;
        font-family: monospace;
        top: 0;
        left: 0;
      }
      :nth-child(1){
        transform: translateZ(-1px);
        position: static;
      }
      :nth-child(2){
        transform: translateZ(-2px);
      }
      :nth-child(3){
        transform: translateZ(-3px);
      }
      :nth-child(4){
        transform: translateZ(-4px);
      }
      :nth-child(5){
        transform: translateZ(-5px);
      }
      :nth-child(6){
        transform: translateZ(-6px);
      }
      :nth-child(7){
        transform: translateZ(-7px);
      }
      :nth-child(8){
        transform: translateZ(-8px);
      }
      :nth-child(9){
        transform: translateZ(-9px);
      }
      :nth-child(10){
        transform: translateZ(-10px);
      }
    }

  }

  @keyframes slogan {
    from {
      transform: rotateY(0deg);
    }
    to {
      transform: rotateY(360deg);
    }
  }
</style>