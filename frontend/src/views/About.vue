<template>
  <div class="about-page">
    <div class="page-header">
      <h1>关于我</h1>
      <p>了解更多关于我的信息</p>
    </div>

    <div class="container">
      <div class="about-content">
        <div class="profile-section">
          <div class="profile-avatar">
            <img src="https://via.placeholder.com/200" alt="头像" />
          </div>
          <div class="profile-info">
            <h2>你好，我是 [你的名字]</h2>
            <p class="bio">
              这里是你的个人简介。可以介绍你的职业、技能、兴趣爱好等。 这是一个使用 Go + Vue 3
              构建的个人网站，用于展示博客文章和作品集。
            </p>
            <div class="social-links">
              <a href="https://github.com/yourusername" target="_blank" class="social-link">GitHub</a>
              <a href="mailto:your@email.com" class="social-link">Email</a>
            </div>
          </div>
        </div>

        <div class="skills-section">
          <h3>技能栈</h3>
          <div class="skills-grid">
            <div class="skill-category">
              <h4>前端</h4>
              <div class="skill-tags">
                <span>Vue.js</span>
                <span>React</span>
                <span>TypeScript</span>
                <span>CSS/SCSS</span>
              </div>
            </div>
            <div class="skill-category">
              <h4>后端</h4>
              <div class="skill-tags">
                <span>Go</span>
                <span>Node.js</span>
                <span>Python</span>
                <span>SQL</span>
              </div>
            </div>
            <div class="skill-category">
              <h4>工具</h4>
              <div class="skill-tags">
                <span>Git</span>
                <span>Docker</span>
                <span>Linux</span>
                <span>VS Code</span>
              </div>
            </div>
          </div>
        </div>

        <div class="contact-section">
          <h3>联系方式</h3>
          <p>如果你有任何问题或合作意向，欢迎通过以下方式联系我：</p>
          <ul>
            <li>📧 Email: your@email.com</li>
            <li>💻 GitHub: github.com/yourusername</li>
          </ul>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
// 可以在这里添加动态数据
</script>

<style lang="scss" scoped>
.about-page {
  .about-content {
    padding: 3rem 0;
    max-width: 800px;
    margin: 0 auto;
  }
}

.profile-section {
  display: flex;
  gap: 3rem;
  align-items: center;
  margin-bottom: 4rem;

  .profile-avatar {
    flex-shrink: 0;

    img {
      width: 200px;
      height: 200px;
      border-radius: 50%;
      object-fit: cover;
      box-shadow: 0 8px 24px rgba(0, 0, 0, 0.15);
    }
  }

  .profile-info {
    h2 {
      font-size: 2rem;
      margin-bottom: 1rem;
      color: #2d3748;
    }

    .bio {
      color: #718096;
      line-height: 1.8;
      margin-bottom: 1.5rem;
    }

    .social-links {
      display: flex;
      gap: 1rem;

      .social-link {
        padding: 0.5rem 1rem;
        background: #667eea;
        color: white;
        border-radius: 6px;
        text-decoration: none;
        transition: background 0.2s;

        &:hover {
          background: #5a67d8;
        }
      }
    }
  }
}

.skills-section {
  margin-bottom: 4rem;

  h3 {
    font-size: 1.5rem;
    margin-bottom: 1.5rem;
    color: #2d3748;
  }

  .skills-grid {
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    gap: 2rem;
  }

  .skill-category {
    h4 {
      margin-bottom: 1rem;
      color: #4a5568;
    }

    .skill-tags {
      display: flex;
      flex-wrap: wrap;
      gap: 0.5rem;

      span {
        padding: 0.4rem 0.8rem;
        background: #eef2ff;
        color: #667eea;
        border-radius: 6px;
        font-size: 0.9rem;
      }
    }
  }
}

.contact-section {
  background: #f7fafc;
  padding: 2rem;
  border-radius: 12px;

  h3 {
    font-size: 1.5rem;
    margin-bottom: 1rem;
    color: #2d3748;
  }

  p {
    color: #718096;
    margin-bottom: 1rem;
  }

  ul {
    list-style: none;

    li {
      padding: 0.5rem 0;
      color: #4a5568;
    }
  }
}

@media (max-width: 768px) {
  .profile-section {
    flex-direction: column;
    text-align: center;

    .social-links {
      justify-content: center;
    }
  }

  .skills-section .skills-grid {
    grid-template-columns: 1fr;
  }
}
</style>
