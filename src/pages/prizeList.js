const MAX_TOP = 300; // 弹幕的最大垂直位置
const MAX_WIDTH = document.body.clientWidth; // 文档主体的最大宽度

let defaultType = 0; // 默认奖项类型
let prizes = []; // 奖项数组
const DEFAULT_MESSAGES = [
  // 默认弹幕消息列表
  '我是该抽中一等奖还是一等奖呢，纠结ing...',
  '听说要提前一个月吃素才能中大奖喔！',
  '好想要一等奖啊！！！',
  '一等奖有没有人想要呢？',
  '五等奖也不错，只要自己能中奖就行',
  '祝大家新年快乐！',
  '中不中奖不重要，大家吃好喝好。',
  '新年，祝福大家事事顺遂。',
  '作为专业陪跑的我，我就看看你们有谁跟我一样',
  '新的一年祝福大家越来越好！',
  '来年再战！！！',
];

let lastDanMuList = []; // 最近的弹幕消息列表

/**
 * 弹幕类，用于创建和控制弹幕效果
 */
class DanMu {
  /**
   * 构造函数
   * @param {Object} options - 配置对象
   * @param {string} options.text - 弹幕文本
   * @param {Function} options.onComplete - 弹幕完成后的回调函数
   */
  constructor({ text, onComplete }) {
    this.position = {}; // 弹幕的位置
    this.text = text; // 弹幕文本
    this.onComplete = onComplete; // 完成后的回调函数
    this.init(); // 初始化弹幕
  }

  /**
   * 初始化弹幕元素
   */
  init() {
    this.element = document.createElement('div'); // 创建弹幕容器
    this.element.className = 'dan-mu'; // 设置类名
    document.body.appendChild(this.element); // 添加到文档主体
    this.start(); // 开始弹幕动画
  }

  /**
   * 设置弹幕文本
   * @param {string} text - 新的弹幕文本
   */
  setText(text) {
    this.text = text || this.text; // 更新文本
    this.element.textContent = this.text; // 设置元素文本内容
    this.width = this.element.clientWidth + 100; // 计算弹幕宽度
  }

  /**
   * 开始弹幕动画
   */
  start() {
    const speed = Math.floor(Math.random() * 10000) + 6000; // 随机速度
    this.position = { x: MAX_WIDTH }; // 初始位置
    this.setText(); // 设置文本
    this.element.style.transform = `translateX(${this.position.x}px)`; // 设置初始位置
    this.element.style.top = `${Math.floor(Math.random() * MAX_TOP) + 10}px`; // 设置随机垂直位置
    this.element.classList.add('active'); // 添加激活类

    new TWEEN.Tween(this.position) // 使用Tween库创建动画
      .to({ x: -this.width }, speed) // 动画目标位置
      .onUpdate(() => this.render()) // 更新位置
      .onComplete(() => this.onComplete?.()) // 动画完成后执行回调
      .start(); // 启动动画
  }

  /**
   * 渲染弹幕位置
   */
  render() {
    this.element.style.transform = `translateX(${this.position.x}px)`; // 更新弹幕位置
  }
}

/**
 * 气泡类，用于创建和控制气泡效果
 */
class Qipao {
  /**
   * 构造函数
   * @param {Object} options - 配置对象
   * @param {string} options.text - 气泡文本
   * @param {Function} options.onComplete - 气泡完成后的回调函数
   */
  constructor({ text, onComplete }) {
    this.text = text; // 气泡文本
    this.onComplete = onComplete; // 完成后的回调函数
    this.$par = document.querySelector('.qipao-container') || createContainer(); // 获取或创建容器

    this.init(); // 初始化气泡
  }

  /**
   * 初始化气泡元素
   */
  init() {
    this.element = document.createElement('div'); // 创建气泡容器
    this.element.className = 'qipao animated'; // 设置类名
    this.$par.appendChild(this.element); // 添加到容器
    this.start(); // 开始气泡动画
  }

  /**
   * 设置气泡文本
   * @param {string} text - 新的气泡文本
   */
  setText(text) {
    this.text = text || this.text; // 更新文本
    this.element.textContent = this.text; // 设置元素文本内容
  }

  /**
   * 开始气泡动画
   */
  start() {
    this.setText(); // 设置文本
    this.element.classList.replace('bounceOutRight', 'bounceInRight'); // 替换类名开始动画

    setTimeout(() => {
      this.element.classList.replace('bounceInRight', 'bounceOutRight'); // 结束动画
      this.onComplete?.(); // 执行回调函数
    }, 10000); // 10秒后结束动画
  }
}

/**
 * 创建气泡容器
 * @returns {HTMLElement} 返回创建的容器元素
 */
function createContainer() {
  const container = document.createElement('div'); // 创建容器
  container.className = 'qipao-container'; // 设置类名
  document.body.appendChild(container); // 添加到文档主体
  return container; // 返回容器
}

const qipaoPool = []; // 气泡池

/**
 * 添加气泡
 * @param {string} text - 气泡文本
 */
function addQipao(text) {
  const qipao =
    qipaoPool.pop() || // 从池中获取或新建气泡
    new Qipao({
      onComplete: () => qipaoPool.push(qipao), // 回收气泡到池中
    });
  qipao.start(text); // 开始气泡动画
}

/**
 * 设置奖项数据
 * @param {Array} newPrizes - 新的奖项数组
 */
function setPrizes(newPrizes) {
  prizes = newPrizes; // 更新奖项数组
  defaultType = prizes[0].type; // 设置默认奖项类型
  updateLastPrizeIndex(prizes.length - 1); // 更新最后一个奖项索引
}

/**
 * 显示奖项列表
 * @param {number} currentPrizeIndex - 当前奖项索引
 */
function showPrizeList(currentPrizeIndex) {
  const currentPrize = prizes[currentPrizeIndex]; // 获取当前奖项
  const remainingCount = currentPrize.count === '不限制' ? '不限制' : currentPrize.count; // 计算剩余数量

  const htmlCode = ` // 构建HTML代码
    <div class="prize-mess">
      正在抽取<label id="prizeType" class="prize-shine">${currentPrize.text}</label>
      <label id="prizeText" class="prize-shine">${currentPrize.title}</label>，
      剩余<label id="prizeLeft" class="prize-shine">${remainingCount}</label>个
    </div>
    <ul class="prize-list">
      ${prizes
        .map(
          (item, index) => ` // 映射奖项到HTML
        <li id="prize-item-${item.type}" class="prize-item ${index === currentPrizeIndex ? 'shine' : ''}">
          <span></span><span></span><span></span><span></span>
          <div class="prize-img"><img src="${item.img}" alt="${item.title}"></div>
          <div class="prize-text">
            <h5 class="prize-title">${item.text} ${item.title}</h5>
            <div class="prize-count">
              <div class="progress">
                <div id="prize-bar-${
                  item.type
                }" class="progress-bar progress-bar-danger progress-bar-striped active" style="width: 100%;">
                </div>
              </div>
              <div id="prize-count-${item.type}" class="prize-count-left">${item.count}/${item.count}</div>
            </div>
          </div>
        </li>
      `
        )
        .join('')} // 连接所有奖项HTML
    </ul>
  `;

  document.getElementById('prizeBar').innerHTML = htmlCode; // 插入到DOM中
}

/**
 * 重置奖项
 * @param {number} currentPrizeIndex - 当前奖项索引
 */
function resetPrize(currentPrizeIndex) {
  clearPrizeElements(); // 清除奖项元素
  updateLastPrizeIndex(currentPrizeIndex); // 更新最后一个奖项索引
  showPrizeList(currentPrizeIndex); // 显示奖项列表
}

const prizeElements = {}; // 奖项元素缓存

/**
 * 获取奖项元素
 * @param {number} type - 奖项类型
 * @returns {Object} 返回奖项元素对象
 */
function getPrizeElement(type) {
  return (
    prizeElements[type] || // 如果已存在则返回
    (prizeElements[type] = {
      // 否则创建并缓存
      box: document.getElementById(`prize-item-${type}`),
      bar: document.getElementById(`prize-bar-${type}`),
      text: document.getElementById(`prize-count-${type}`),
    })
  );
}

/**
 * 清除奖项元素缓存
 */
function clearPrizeElements() {
  Object.keys(prizeElements).forEach(key => delete prizeElements[key]); // 删除所有缓存
}

/**
 * 更新最后一个奖项索引
 * @param {number} index - 新的奖项索引
 */
function updateLastPrizeIndex(index) {
  window.lastPrizeIndex = index; // 更新全局变量
}

/**
 * 设置奖项数据
 * @param {number} currentPrizeIndex - 当前奖项索引
 * @param {number} count - 剩余数量
 * @param {boolean} isInit - 是否初始化
 */
function setPrizeData(currentPrizeIndex, count, isInit) {
  const currentPrize = prizes[currentPrizeIndex]; // 获取当前奖项
  const { type } = currentPrize; // 获取奖项类型
  const elements = getPrizeElement(type); // 获取奖项元素

  if (isInit) {
    // 如果是初始化
    for (let i = prizes.length - 1; i > currentPrizeIndex; i--) {
      // 遍历后续奖项
      const prizeType = prizes[i].type; // 获取奖项类型
      document.getElementById(`prize-item-${prizeType}`).classList.remove('shine'); // 移除高亮
      document.getElementById(`prize-bar-${prizeType}`).style.width = '0%'; // 设置进度条宽度
      document.getElementById(`prize-count-${prizeType}`).textContent = `0/${prizes[i].count}`; // 设置剩余数量
    }
  }

  if (window.lastPrizeIndex !== currentPrizeIndex) {
    // 如果奖项变化
    const lastPrize = prizes[window.lastPrizeIndex]; // 获取上一个奖项
    document.getElementById(`prize-item-${lastPrize.type}`).classList.remove('shine'); // 移除上一个奖项高亮
    elements.box.classList.add('shine'); // 设置当前奖项高亮

    document.getElementById('prizeType').textContent = currentPrize.text; // 设置奖项类型文本
    document.getElementById('prizeText').textContent = currentPrize.title; // 设置奖项标题文本
    updateLastPrizeIndex(currentPrizeIndex); // 更新最后一个奖项索引
  }

  if (currentPrizeIndex === 0) {
    // 如果是额外奖
    document.getElementById('prizeType').textContent = '额外奖'; // 设置奖项类型文本
    document.getElementById('prizeText').textContent = ''; // 清空奖项标题文本
    document.getElementById('prizeLeft').textContent = '不限制'; // 设置剩余数量为不限制
    return;
  }

  count = Math.max(currentPrize.count - count, 0); // 计算剩余数量
  const percentage = ((count / currentPrize.count) * 100).toFixed(2); // 计算百分比

  elements.bar.style.width = `${percentage}%`; // 设置进度条宽度
  elements.text.textContent = `${count}/${currentPrize.count}`; // 设置剩余数量文本
  document.getElementById('prizeLeft').textContent = count; // 设置剩余数量文本
}

/**
 * 开始弹幕效果
 */
function startMaoPao() {
  const messageCount = DEFAULT_MESSAGES.length; // 默认消息数量
  const danmuCount = 5; // 弹幕数量
  let currentIndex = Math.floor(Math.random() * messageCount); // 随机起始索引
  const danmus = []; // 弹幕实例数组

  /**
   * 重启弹幕
   */
  function restart() {
    danmus.forEach(danmu => {
      // 遍历所有弹幕
      const text = lastDanMuList.length > 0 ? lastDanMuList.shift() : DEFAULT_MESSAGES[currentIndex++]; // 获取下一个消息
      danmu.start(text); // 开始弹幕动画
      currentIndex %= messageCount; // 循环索引
    });
  }

  for (let i = 0; i < danmuCount; i++) {
    // 创建指定数量的弹幕
    setTimeout(() => {
      danmus.push(
        new DanMu({
          text: DEFAULT_MESSAGES[currentIndex++], // 设置初始消息
          onComplete: () =>
            setTimeout(() => {
              danmu.start(DEFAULT_MESSAGES[currentIndex++]); // 重新开始弹幕动画
              currentIndex %= messageCount; // 循环索引
            }, 1000), // 1秒后重新开始
        })
      );
      currentIndex %= messageCount; // 循环索引
    }, 1500 * i); // 间隔1.5秒创建下一个弹幕
  }
}

/**
 * 添加弹幕消息
 * @param {string} text - 弹幕文本
 */
function addDanMu(text) {
  lastDanMuList.push(text); // 添加到最近消息列表
}

export { startMaoPao, showPrizeList, setPrizeData, addDanMu, setPrizes, resetPrize, addQipao }; // 导出公共方法
