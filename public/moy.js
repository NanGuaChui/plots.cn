// ==UserScript==
// @name         自动点击卡片按钮
// @namespace    http://tampermonkey.net/
// @version      1.0
// @description  监听卡片变化并自动点击符合条件的按钮
// @author       You
// @match        **moyu-idle.com/*
// @match        *://*moyu-idle.com/*
// @match        *://www.moyu-idle.com/*
// @grant        GM_addStyle
// @grant        GM_getValue
// @grant        GM_setValue
// @grant        unsafeWindow
// @grant        GM_xmlhttpRequest
// @grant        GM_setClipboard
// @run-at       document-start
// ==/UserScript==

(function () {
  'use strict';

  // --- 配置常量 ---
  const CONFIG = {
    selectors: {
      cards: '.el-main .el-card .el-card',
      cardTitle: '.truncate',
      actionButton: '.el-button.el-button--success',
    },
    keywords: {
      targetPrefix: '采集',
      exclude: ['云絮', '彩虹', '种植'],
    },
    delay: {
      min: 500,
      max: 1500,
    },
    colors: {
      success: { bg: '#67C23A', border: '#5DAF34' },
      warning: { bg: '#E6A23C', border: '#CF9236' },
      error: { bg: '#F56C6C', border: '#DD6161' },
      info: { bg: '#409EFF', border: '#3A8EE6' },
    },
  };

  // --- 工具函数 ---
  const log = (msg, ...args) => {
    console.log(`[MoyuScript] ${msg}`, ...args);
  };

  // --- 状态管理 ---
  let activeObservers = [];

  // --- 核心逻辑 ---
  // ====== 功能1：自动采集卡片 ======
  /**
   * 检查卡片是否处于“已完成/运行中”状态
   * @param {HTMLElement} card
   * @returns {boolean} true 表示卡片正在采集（不需要点击），false 表示需要点击
   */
  const isCardRunning = (card) => {
    const titleElement = card.querySelector(CONFIG.selectors.cardTitle);
    if (!titleElement) return true; // 找不到标题视为已完成，避免误操作

    const title = (titleElement.textContent || titleElement.innerText).trim();
    const { targetPrefix, exclude } = CONFIG.keywords;

    // 逻辑：标题以 '采集' 开头，且不包含排除关键词
    const isTarget = title.startsWith(targetPrefix);
    const isExcluded = exclude.some((keyword) => title.includes(keyword));

    return isTarget && !isExcluded;
  };

  /**
   * 尝试点击卡片上的按钮
   * @param {HTMLElement} card
   * @returns {boolean} 是否成功找到并点击了按钮
   */
  const tryClickButton = (card) => {
    const button = card.querySelector(CONFIG.selectors.actionButton);
    if (button) {
      button.click();
      return true;
    }
    return false;
  };

  /**
   * 获取随机延迟时间
   */
  const getRandomDelay = () => {
    return Math.random() * (CONFIG.delay.max - CONFIG.delay.min) + CONFIG.delay.min;
  };

  /**
   * 处理单个卡片的自动化逻辑
   * @param {HTMLElement} card
   * @param {number} index
   */
  const processCard = (card, index) => {
    // 如果卡片已经是运行状态，直接跳过
    if (isCardRunning(card)) return null;

    let isWaiting = false; // 简单的防抖锁，防止短时间内重复点击

    const observer = new MutationObserver(() => {
      if (isCardRunning(card)) {
        log(`✅ 卡片 ${index + 1} 已进入目标状态，停止监听`);
        observer.disconnect();
        // 从全局数组中移除
        activeObservers = activeObservers.filter((obs) => obs !== observer);

        if (activeObservers.length === 0) {
          showToast('所有任务已完成！', 'success');
        }
      } else {
        if (isWaiting) return; // 如果正在等待点击，则忽略新的变动

        isWaiting = true;
        log(`⏳ 卡片 ${index + 1} 未完成，准备点击...`);

        setTimeout(() => {
          tryClickButton(card);
          isWaiting = false;
        }, getRandomDelay());
      }
    });

    observer.observe(card, {
      childList: true,
      subtree: true,
      characterData: true,
    });

    // 立即触发一次点击
    tryClickButton(card);

    return observer;
  };

  const executeAutoCollect = () => {
    log('🚀 开始执行自动采集...');

    // 清理旧的 observers
    activeObservers.forEach((obs) => obs.disconnect());
    activeObservers = [];

    const allCards = Array.from(document.querySelectorAll(CONFIG.selectors.cards));
    // 筛选出需要处理的卡片（即 !isCardRunning 的卡片）
    const targetCards = allCards.filter((card) => !isCardRunning(card));

    log(`🔍 找到 ${targetCards.length} 张待处理卡片`, targetCards);

    if (targetCards.length === 0) {
      showToast('未找到符合条件的卡片', 'warning');
      return;
    }

    targetCards.forEach((card, index) => {
      const observer = processCard(card, index);
      if (observer) {
        activeObservers.push(observer);
      }
    });

    showToast(`开始处理 ${targetCards.length} 张卡片`, 'info');
  };

  // ====== 功能2：快速开始任务 ======

  /**
   * 执行快速开始（点击快速开始按钮并确认）
   */
  const executeQuickStart = async () => {
    log('🚀 开始执行快速开始...');

    const cards = Array.from(document.querySelectorAll(CONFIG.selectors.cards));

    if (!cards.length) {
      showToast('未找到任务卡片', 'warning');
      return;
    }

    let processedCount = 0;

    for (const card of cards) {
      try {
        const btn = card.querySelector('.el-button.el-button--primary.el-button--small');

        if (btn && btn.innerText === '快速开始') {
          btn.click();
          await new Promise((resolve) => setTimeout(resolve, 300));
        }
      } catch (error) {
        log('❌ 快速开始出错:', error);
      }
    }

    showToast(`快速开始完成！处理了 ${processedCount} 个任务`, 'success');
  };

  // ====== 功能3：自动使用技能书 ======

  /**
   * 使用单个技能书卡片
   * @param {HTMLElement} card
   */
  const useSkillCard = async (card) => {
    return new Promise(async (resolve) => {
      card.click();
      await new Promise((r) => setTimeout(r, 500));
      const dialogEl = card.parentElement.nextElementSibling;
      const buttons = dialogEl.querySelectorAll('.el-button.el-button--primary.el-button--small');
      if (buttons[3]) {
        buttons[3].click();
        await new Promise((r) => setTimeout(r, 500));
        const useBtn = dialogEl.querySelector('.el-button.el-button--primary.w-full');
        if (useBtn) {
          resolve(useBtn.click());
        } else {
          resolve();
        }
      } else {
        resolve();
      }
    });
  };

  /**
   * 执行自动使用技能书
   */
  const executeAutoUseSkills = async () => {
    log('🚀 开始执行自动使用技能书...');

    const skills = Array.from(document.querySelectorAll('.w-full.font-bold.text-lg')).filter((el) => {
      return el.innerText === '技能书';
    });

    if (!skills.length) {
      showToast('未找到技能书区域', 'warning');
      return;
    }

    const skillSection = skills[0].nextElementSibling;
    const skillCards = skillSection.querySelectorAll(':scope > div > .item-card');

    if (!skillCards.length) {
      showToast('未找到技能书卡片', 'warning');
      return;
    }

    showToast(`找到 ${skillCards.length} 张技能书，开始使用...`, 'info');

    for (const card of skillCards) {
      try {
        await useSkillCard(card);
        await new Promise((r) => setTimeout(r, 300));
      } catch (error) {
        log('❌ 使用技能书出错:', error);
      }
    }

    showToast('技能书使用完成！', 'success');
  };

  // --- UI 与 工具函数 ---

  /**
   * 创建功能菜单
   */
  const createMenu = () => {
    if (document.getElementById('moyu-menu')) return;

    const menu = document.createElement('div');
    menu.id = 'moyu-menu';
    menu.style.display = 'none'; // 默认隐藏

    Object.assign(menu.style, {
      position: 'fixed',
      bottom: '180px',
      right: '30px',
      backgroundColor: '#fff',
      borderRadius: '8px',
      boxShadow: '0 4px 16px rgba(0, 0, 0, 0.2)',
      zIndex: '9998',
      overflow: 'hidden',
      minWidth: '160px',
    });

    const menuItems = [
      { text: '🌾 自动采集', action: executeAutoCollect },
      { text: '⚡ 快速开始', action: executeQuickStart },
      { text: '📚 使用技能书', action: executeAutoUseSkills },
    ];

    menuItems.forEach((item, index) => {
      const menuItem = document.createElement('div');
      menuItem.textContent = item.text;
      menuItem.className = 'moyu-menu-item';

      Object.assign(menuItem.style, {
        padding: '12px 20px',
        cursor: 'pointer',
        fontSize: '14px',
        color: '#333',
        transition: 'background-color 0.2s',
        borderBottom: index < menuItems.length - 1 ? '1px solid #eee' : 'none',
      });

      menuItem.addEventListener('mouseenter', () => {
        menuItem.style.backgroundColor = '#f5f5f5';
      });

      menuItem.addEventListener('mouseleave', () => {
        menuItem.style.backgroundColor = 'transparent';
      });

      menuItem.addEventListener('click', () => {
        log(`🖱️ 点击菜单项: ${item.text}`);
        toggleMenu(false); // 关闭菜单
        item.action(); // 执行对应功能
      });

      menu.appendChild(menuItem);
    });

    document.body.appendChild(menu);
    return menu;
  };

  /**
   * 切换菜单显示/隐藏
   * @param {boolean} show - 是否显示菜单
   */
  const toggleMenu = (show) => {
    const menu = document.getElementById('moyu-menu') || createMenu();
    menu.style.display = show ? 'block' : 'none';
  };

  const showToast = (message, type = 'info') => {
    log(`📢 显示提示: ${message}`);

    // 1. 移除已存在的 Toast (单例模式)
    const existingToast = document.querySelector('.moyu-toast');
    if (existingToast) {
      existingToast.remove();
    }

    const toast = document.createElement('div');
    toast.className = 'moyu-toast'; // 添加类名以便查找
    toast.textContent = message;

    const color = CONFIG.colors[type] || CONFIG.colors.info;

    Object.assign(toast.style, {
      position: 'fixed',
      top: '20px',
      left: '50%',
      transform: 'translateX(-50%)',
      padding: '12px 24px',
      backgroundColor: color.bg,
      color: '#fff',
      borderRadius: '4px',
      border: `2px solid ${color.border}`,
      boxShadow: '0 4px 12px rgba(0, 0, 0, 0.15)',
      fontSize: '14px',
      fontWeight: '500',
      zIndex: '10000',
      opacity: '0',
      transition: 'opacity 0.3s ease',
      pointerEvents: 'none', // 防止遮挡点击
    });

    document.body.appendChild(toast);

    // 强制重绘以触发 transition
    requestAnimationFrame(() => {
      toast.style.opacity = '1';
    });

    // 3秒后自动消失
    setTimeout(() => {
      toast.style.opacity = '0';
      setTimeout(() => {
        if (toast.parentNode) toast.remove();
      }, 300);
    }, 3000);
  };

  const createFloatingButton = () => {
    if (document.getElementById('auto-click-floating-btn')) return;

    const button = document.createElement('div');
    button.id = 'auto-click-floating-btn';
    button.innerHTML = '🚀';
    button.title = '点击打开菜单';

    Object.assign(button.style, {
      position: 'fixed',
      bottom: '110px',
      right: '30px',
      width: '60px',
      height: '60px',
      borderRadius: '50%',
      backgroundColor: '#409EFF',
      color: '#fff',
      fontSize: '30px',
      display: 'flex',
      alignItems: 'center',
      justifyContent: 'center',
      cursor: 'pointer',
      boxShadow: '0 4px 12px rgba(0, 0, 0, 0.3)',
      zIndex: '9999',
      transition: 'transform 0.2s, box-shadow 0.2s',
      userSelect: 'none',
    });

    button.addEventListener('mouseenter', () => {
      button.style.transform = 'scale(1.1)';
      button.style.boxShadow = '0 6px 16px rgba(0, 0, 0, 0.4)';
    });

    button.addEventListener('mouseleave', () => {
      button.style.transform = 'scale(1)';
      button.style.boxShadow = '0 4px 12px rgba(0, 0, 0, 0.3)';
    });

    // 点击按钮切换菜单显示/隐藏
    button.addEventListener('click', (e) => {
      e.stopPropagation(); // 防止事件冒泡
      log('🖱️ 用户点击了悬浮按钮');

      const menu = document.getElementById('moyu-menu');
      const isMenuVisible = menu && menu.style.display === 'block';
      toggleMenu(!isMenuVisible);

      // 简单的点击动画
      button.style.transform = 'scale(0.9)';
      setTimeout(() => (button.style.transform = 'scale(1)'), 150);
    });

    document.body.appendChild(button);

    // 点击页面其他地方关闭菜单
    document.addEventListener('click', (e) => {
      const menu = document.getElementById('moyu-menu');
      if (menu && !menu.contains(e.target) && e.target !== button) {
        toggleMenu(false);
      }
    });
  };

  // 初始化
  if (document.readyState === 'loading') {
    document.addEventListener('DOMContentLoaded', () => {
      createFloatingButton();
      createMenu();
    });
  } else {
    createFloatingButton();
    createMenu();
  }
})();
