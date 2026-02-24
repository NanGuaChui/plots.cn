<template>
  <div id="app">
    <nav class="navbar" v-if="!isAdminRoute">
      <div class="container">
        <router-link to="/" class="logo">plots</router-link>
        <div class="nav-links">
          <router-link to="/">首页</router-link>
          <router-link to="/articles">博客</router-link>
          <router-link to="/portfolio">作品集</router-link>
          <router-link to="/about">关于</router-link>
        </div>
      </div>
    </nav>

    <main>
      <router-view />
    </main>

    <footer class="footer" v-if="!isAdminRoute">
      <div class="container">
        <p>&copy; {{ currentYear }} plots. All rights reserved.</p>
      </div>
    </footer>
  </div>
</template>

<script setup>
import { computed } from 'vue';
import { useRoute } from 'vue-router';

const route = useRoute();
const currentYear = new Date().getFullYear();

const isAdminRoute = computed(() => {
  return route.path.startsWith('/admin');
});
</script>

<style lang="scss">
#app {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

main {
  flex: 1;
}

.navbar {
  background: #fff;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06);
  position: sticky;
  top: 0;
  z-index: 100;

  .container {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 1rem 2rem;
    max-width: 1200px;
    margin: 0 auto;
  }

  .logo {
    font-size: 1.5rem;
    font-weight: 700;
    color: #667eea;
    text-decoration: none;
  }

  .nav-links {
    display: flex;
    gap: 2rem;

    a {
      color: #4a5568;
      text-decoration: none;
      font-weight: 500;
      transition: color 0.2s;

      &:hover,
      &.router-link-active {
        color: #667eea;
      }
    }
  }
}

.footer {
  background: #1a202c;
  color: #a0aec0;
  padding: 2rem;
  text-align: center;

  p {
    margin: 0;
  }
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 1rem;
}
</style>
