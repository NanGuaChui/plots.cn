<template>
  <div class="not-found-page">
    <div class="container">
      <h1>404</h1>
      <p>抱歉，您访问的页面不存在</p>
      <router-link to="/" class="btn btn-primary">返回首页</router-link>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.not-found-page {
  min-height: 60vh;
  display: flex;
  align-items: center;
  justify-content: center;
  text-align: center;

  h1 {
    font-size: 8rem;
    color: #667eea;
    margin-bottom: 1rem;
  }

  p {
    font-size: 1.5rem;
    color: #718096;
    margin-bottom: 2rem;
  }
}
</style>
