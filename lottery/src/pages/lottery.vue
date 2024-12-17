<template>
  <div class="canvas-box">
    <canvas id="canvas">你的浏览器不支持canvas</canvas>
  </div>

  <div class="music">
    <audio id="music" :src="music" class="music-item" loop></audio>
    <div id="musicBox" class="music-box" title="播放/暂停背景音乐">Music</div>
  </div>

  <div id="title">Merry Christmas and Happy New Year Gathering</div>

  <div id="prizeBar"></div>

  <div id="container"></div>

  <div id="menu">
    <button id="enter">进入抽奖</button>
    <div id="lotteryBar" class="none">
      <button id="lottery">开始抽奖</button>
      <button id="reLottery">重新抽奖</button>
      <div class="fixed-bar">
        <!-- <button id="save" class="fixed-btn">导出抽奖结果</button> -->
        <button id="reset" class="fixed-btn">重置</button>
      </div>
    </div>
  </div>

  <div class="qipao-container"></div>

  <div id="setting" class="setting">
    <img @click="showCategory = !showCategory" :src="settingIcon" alt="" />
    <ul v-show="showCategory" class="dropdown"></ul>
  </div>
</template>

<script setup>
import './lottery.css';
import initLottery from './lottery.js';
import music from '../assets/audio/music.mp3';
import { onMounted, ref } from 'vue';
import settingIcon from '../assets/img/setting.png';

const showCategory = ref(false);

onMounted(() => {
  initLottery();
});
</script>

<style lang="scss">
.setting {
  position: fixed;
  left: 1vw;
  bottom: 1vw;
  width: 2vw;
  height: 2vw;
  opacity: 0.6;
  cursor: pointer;
  z-index: 999;
  img {
    display: block;
    width: 100%;
    height: 100%;
  }
  .dropdown {
    position: absolute;
    margin: 0;
    padding: 0;
    left: 0;
    bottom: 2.5vw;
    list-style: none;
    li {
      margin: 0;
      button {
        margin: 0.5vw 0 0;
        padding: 0.3vw 1vw;
        word-break: keep-all;
      }
    }
  }
}
</style>
