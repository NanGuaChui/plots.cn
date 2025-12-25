// ==UserScript==
// @name         摸鱼自动化助手
// @namespace    http://tampermonkey.net/
// @version      2.0
// @description  自动采集卡片、快速开始任务、使用技能书
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

  // ========== CONFIG ==========
  const CONFIG = {
    // DOM 选择器
    selectors: {
      cards: '.el-main .el-card .el-card',
      cardTitle: '.truncate',
      actionButton: '.el-button.el-button--success',
      quickStartBtn: '.el-button.el-button--primary.el-button--small',
      sectionTitle: '.w-full.font-bold.text-lg',
      skillCard: '.item-card',
      dialogBtn: '.el-button.el-button--primary.el-button--small',
      useSkillBtn: '.el-button.el-button--primary.w-full',
    },

    // 关键词配置
    keywords: {
      targetPrefix: '采集',
      exclude: ['云絮', '彩虹', '种植'],
      quickStart: '快速开始',
      skillBook: '技能书',
    },

    // 延迟配置 (ms)
    delay: {
      min: 300,
      max: 800,
      short: 300,
      medium: 500,
      long: 1000,
    },

    // Toast 颜色配置
    colors: {
      success: { bg: '#67C23A', border: '#5DAF34' },
      warning: { bg: '#E6A23C', border: '#CF9236' },
      error: { bg: '#F56C6C', border: '#DD6161' },
      info: { bg: '#409EFF', border: '#3A8EE6' },
    },

    // 提示消息
    messages: {
      collectDone: '采集完成，正在启动快速开始...',
      noCollectCards: '无待采集卡片，启动快速开始...',
      noTaskCards: '未找到任务卡片',
      quickStartDone: (count) => `快速开始完成！处理了 ${count} 个任务`,
      processing: (count) => `开始处理 ${count} 张卡片`,
      noSkillSection: '未找到技能书区域',
      noSkillCards: '未找到技能书卡片',
      foundSkills: (count) => `找到 ${count} 张技能书，开始使用...`,
      skillsDone: '技能书使用完成！',
    },

    // 菜单配置（已移除"快速开始"，因为它会在采集完成后自动执行）
    menuItems: [
      { icon: '🌾', text: '自动采集', action: 'autoCollect' },
      { icon: '📚', text: '使用技能书', action: 'useSkills' },
    ],

    // UI 样式
    styles: {
      menu: {
        position: 'fixed',
        bottom: '180px',
        right: '30px',
        backgroundColor: '#fff',
        borderRadius: '8px',
        boxShadow: '0 4px 16px rgba(0, 0, 0, 0.2)',
        zIndex: '9998',
        overflow: 'hidden',
        minWidth: '160px',
      },
      menuItem: {
        padding: '12px 20px',
        cursor: 'pointer',
        fontSize: '14px',
        color: '#333',
        transition: 'background-color 0.2s',
      },
      floatingButton: {
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
      },
      toast: {
        position: 'fixed',
        top: '20px',
        left: '50%',
        transform: 'translateX(-50%)',
        padding: '12px 24px',
        color: '#fff',
        borderRadius: '4px',
        boxShadow: '0 4px 12px rgba(0, 0, 0, 0.15)',
        fontSize: '14px',
        fontWeight: '500',
        zIndex: '10000',
        opacity: '0',
        transition: 'opacity 0.3s ease',
        pointerEvents: 'none',
      },
    },
  };

  // ========== UTILS ==========
  /**
   * 控制台日志
   */
  const log = (msg, ...args) => console.log(`[MoyuScript] ${msg}`, ...args);

  /**
   * 延迟执行
   * @param {number} ms - 毫秒数
   */
  const sleep = (ms) => new Promise((resolve) => setTimeout(resolve, ms));

  /**
   * 获取随机延迟时间
   */
  const randomDelay = () => Math.random() * (CONFIG.delay.max - CONFIG.delay.min) + CONFIG.delay.min;

  /**
   * 批量应用样式
   * @param {HTMLElement} el - DOM 元素
   * @param {Object} styles - 样式对象
   */
  const applyStyles = (el, styles) => Object.assign(el.style, styles);

  /**
   * 简化的选择器查询
   * @param {string} selector - CSS 选择器
   * @param {HTMLElement} parent - 父元素，默认 document
   */
  const $ = (selector, parent = document) => parent.querySelector(selector);
  const $$ = (selector, parent = document) => Array.from(parent.querySelectorAll(selector));

  // ========== UI ==========
  /**
   * 显示 Toast 提示
   * @param {string} message - 提示消息
   * @param {string} type - 类型: success/warning/error/info
   */
  const showToast = (message, type = 'info') => {
    log(`📢 显示提示: ${message}`);

    // 移除已存在的 Toast (单例模式)
    const existingToast = $('.moyu-toast');
    if (existingToast) existingToast.remove();

    const toast = document.createElement('div');
    toast.className = 'moyu-toast';
    toast.textContent = message;

    const color = CONFIG.colors[type] || CONFIG.colors.info;
    applyStyles(toast, {
      ...CONFIG.styles.toast,
      backgroundColor: color.bg,
      border: `2px solid ${color.border}`,
    });

    document.body.appendChild(toast);

    // 强制重绘以触发 transition
    requestAnimationFrame(() => (toast.style.opacity = '1'));

    // 3秒后自动消失
    setTimeout(() => {
      toast.style.opacity = '0';
      setTimeout(() => toast.parentNode && toast.remove(), 300);
    }, 3000);
  };

  /**
   * 切换菜单显示/隐藏
   * @param {boolean} show - 是否显示菜单
   */
  const toggleMenu = (show) => {
    const menu = $('#moyu-menu') || createMenu();
    menu.style.display = show ? 'block' : 'none';
  };

  /**
   * 创建功能菜单
   */
  const createMenu = () => {
    if ($('#moyu-menu')) return $('#moyu-menu');

    const menu = document.createElement('div');
    menu.id = 'moyu-menu';
    menu.style.display = 'none';
    applyStyles(menu, CONFIG.styles.menu);

    // 任务动作映射
    const actionMap = {
      autoCollect: executeAutoCollect,
      useSkills: executeAutoUseSkills,
    };

    CONFIG.menuItems.forEach((item, index) => {
      const menuItem = document.createElement('div');
      menuItem.textContent = `${item.icon} ${item.text}`;
      menuItem.className = 'moyu-menu-item';

      applyStyles(menuItem, {
        ...CONFIG.styles.menuItem,
        borderBottom: index < CONFIG.menuItems.length - 1 ? '1px solid #eee' : 'none',
      });

      menuItem.addEventListener('mouseenter', () => (menuItem.style.backgroundColor = '#f5f5f5'));
      menuItem.addEventListener('mouseleave', () => (menuItem.style.backgroundColor = 'transparent'));
      menuItem.addEventListener('click', () => {
        log(`🖱️ 点击菜单项: ${item.text}`);
        toggleMenu(false);
        actionMap[item.action]?.();
      });

      menu.appendChild(menuItem);
    });

    document.body.appendChild(menu);
    return menu;
  };

  /**
   * 创建悬浮按钮
   */
  const createFloatingButton = () => {
    if ($('#auto-click-floating-btn')) return;

    const button = document.createElement('div');
    button.id = 'auto-click-floating-btn';
    button.innerHTML = '🚀';
    button.title = '点击打开菜单';

    applyStyles(button, CONFIG.styles.floatingButton);

    button.addEventListener('mouseenter', () => {
      button.style.transform = 'scale(1.1)';
      button.style.boxShadow = '0 6px 16px rgba(0, 0, 0, 0.4)';
    });

    button.addEventListener('mouseleave', () => {
      button.style.transform = 'scale(1)';
      button.style.boxShadow = '0 4px 12px rgba(0, 0, 0, 0.3)';
    });

    button.addEventListener('click', (e) => {
      e.stopPropagation();
      log('🖱️ 用户点击了悬浮按钮');

      const menu = $('#moyu-menu');
      const isMenuVisible = menu && menu.style.display === 'block';
      toggleMenu(!isMenuVisible);

      // 点击动画
      button.style.transform = 'scale(0.9)';
      setTimeout(() => (button.style.transform = 'scale(1)'), 150);
    });

    document.body.appendChild(button);

    // 点击页面其他地方关闭菜单
    document.addEventListener('click', (e) => {
      const menu = $('#moyu-menu');
      if (menu && !menu.contains(e.target) && e.target !== button) {
        toggleMenu(false);
      }
    });
  };

  // ========== TASKS ==========
  // 状态管理
  let activeObservers = [];

  /**
   * 检查卡片是否处于"已完成/运行中"状态
   * @param {HTMLElement} card
   * @returns {boolean} true 表示卡片正在采集，false 表示需要点击
   */
  const isCardRunning = (card) => {
    const titleElement = $(CONFIG.selectors.cardTitle, card);
    if (!titleElement) return true;

    const title = (titleElement.textContent || titleElement.innerText).trim();
    const { targetPrefix, exclude } = CONFIG.keywords;

    const isTarget = title.startsWith(targetPrefix);
    const isExcluded = exclude.some((keyword) => title.includes(keyword));

    return isTarget && !isExcluded;
  };

  /**
   * 尝试点击卡片上的按钮
   * @param {HTMLElement} card
   * @returns {boolean} 是否成功点击
   */
  const tryClickButton = (card) => {
    const button = $(CONFIG.selectors.actionButton, card);
    if (button) {
      button.click();
      return true;
    }
    return false;
  };

  /**
   * 处理单个卡片的自动化逻辑
   * @param {HTMLElement} card
   * @param {number} index
   */
  const processCard = (card, index) => {
    if (isCardRunning(card)) return null;

    let isWaiting = false;

    const observer = new MutationObserver(async () => {
      if (isCardRunning(card)) {
        log(`✅ 卡片 ${index + 1} 已进入目标状态，停止监听`);
        observer.disconnect();
        activeObservers = activeObservers.filter((obs) => obs !== observer);

        if (activeObservers.length === 0) {
          showToast(CONFIG.messages.collectDone, 'success');
          await sleep(CONFIG.delay.long);
          executeQuickStart();
        }
      } else {
        if (isWaiting) return;

        isWaiting = true;
        log(`⏳ 卡片 ${index + 1} 未完成，准备点击...`);

        setTimeout(() => {
          tryClickButton(card);
          isWaiting = false;
        }, randomDelay());
      }
    });

    observer.observe(card, {
      childList: true,
      subtree: true,
      characterData: true,
    });

    tryClickButton(card);
    return observer;
  };

  /**
   * 执行自动采集
   */
  const executeAutoCollect = () => {
    log('🚀 开始执行自动采集...');

    // 清理旧的 observers
    activeObservers.forEach((obs) => obs.disconnect());
    activeObservers = [];

    const allCards = $$(CONFIG.selectors.cards);
    const targetCards = allCards.filter((card) => !isCardRunning(card));

    log(`🔍 找到 ${targetCards.length} 张待处理卡片`, targetCards);

    if (targetCards.length === 0) {
      showToast(CONFIG.messages.noCollectCards, 'warning');
      executeQuickStart();
      return;
    }

    targetCards.forEach((card, index) => {
      const observer = processCard(card, index);
      if (observer) activeObservers.push(observer);
    });

    showToast(CONFIG.messages.processing(targetCards.length), 'info');
  };

  /**
   * 执行快速开始（内部函数，由自动采集完成后调用）
   */
  const executeQuickStart = async () => {
    log('🚀 开始执行快速开始...');

    const cards = $$(CONFIG.selectors.cards);

    if (!cards.length) {
      showToast(CONFIG.messages.noTaskCards, 'warning');
      return;
    }

    let processedCount = 0;

    for (const card of cards) {
      try {
        const btn = $(CONFIG.selectors.quickStartBtn, card);

        if (btn && btn.innerText === CONFIG.keywords.quickStart) {
          btn.click();
          processedCount++;
          await sleep(CONFIG.delay.short);
        }
      } catch (error) {
        log('❌ 快速开始出错:', error);
      }
    }

    showToast(CONFIG.messages.quickStartDone(processedCount), 'success');
  };

  /**
   * 使用单个技能书卡片
   * @param {HTMLElement} card
   */
  const useSkillCard = async (card) => {
    card.click();
    await sleep(CONFIG.delay.medium);

    const dialogEl = card.parentElement.nextElementSibling;
    const buttons = $$(CONFIG.selectors.dialogBtn, dialogEl);

    if (buttons[3]) {
      buttons[3].click();
      await sleep(CONFIG.delay.medium);

      const useBtn = $(CONFIG.selectors.useSkillBtn, dialogEl);
      if (useBtn) useBtn.click();
    }
  };

  /**
   * 执行自动使用技能书
   */
  const executeAutoUseSkills = async () => {
    log('🚀 开始执行自动使用技能书...');

    const skills = $$(CONFIG.selectors.sectionTitle).filter((el) => el.innerText === CONFIG.keywords.skillBook);

    if (!skills.length) {
      showToast(CONFIG.messages.noSkillSection, 'warning');
      return;
    }

    const skillSection = skills[0].nextElementSibling;
    const skillCards = $$(':scope > div > ' + CONFIG.selectors.skillCard, skillSection);

    if (!skillCards.length) {
      showToast(CONFIG.messages.noSkillCards, 'warning');
      return;
    }

    showToast(CONFIG.messages.foundSkills(skillCards.length), 'info');

    for (const card of skillCards) {
      try {
        await useSkillCard(card);
        await sleep(CONFIG.delay.short);
      } catch (error) {
        log('❌ 使用技能书出错:', error);
      }
    }

    showToast(CONFIG.messages.skillsDone, 'success');
  };

  // ========== INIT ==========
  const init = () => {
    createFloatingButton();
    createMenu();
    log('✅ 摸鱼自动化助手已加载');
  };

  if (document.readyState === 'loading') {
    document.addEventListener('DOMContentLoaded', init);
  } else {
    init();
  }
})();
