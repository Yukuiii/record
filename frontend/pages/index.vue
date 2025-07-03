<template>
  <div class="min-h-screen bg-gray-50">
    <!-- 移动端状态栏 -->
    <div class="bg-white px-4 py-2 flex justify-between items-center text-sm font-medium">
      <div class="flex items-center space-x-1">
        <span class="text-black">17:23</span>
      </div>
      <div class="flex items-center space-x-1">
        <div class="flex space-x-1">
          <div class="w-1 h-3 bg-black rounded-full"></div>
          <div class="w-1 h-3 bg-black rounded-full"></div>
          <div class="w-1 h-3 bg-gray-300 rounded-full"></div>
          <div class="w-1 h-3 bg-gray-300 rounded-full"></div>
        </div>
        <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
          <path fill-rule="evenodd"
            d="M17.778 8.222c-4.296-4.296-11.26-4.296-15.556 0A1 1 0 01.808 6.808c5.076-5.077 13.308-5.077 18.384 0a1 1 0 01-1.414 1.414zM14.95 11.05a7 7 0 00-9.9 0 1 1 0 01-1.414-1.414 9 9 0 0112.728 0 1 1 0 01-1.414 1.414zM12.12 13.88a3 3 0 00-4.24 0 1 1 0 01-1.415-1.414 5 5 0 017.07 0 1 1 0 01-1.415 1.414zM9 16a1 1 0 011-1h.01a1 1 0 110 2H10a1 1 0 01-1-1z"
            clip-rule="evenodd"></path>
        </svg>
        <div class="bg-green-500 text-white text-xs px-1.5 py-0.5 rounded">89%</div>
      </div>
    </div>

    <!-- 头部导航 -->
    <div class="bg-white px-4 py-3 flex justify-between items-center border-b border-gray-100">
      <div class="flex items-center space-x-6">
        <h1 class="text-xl font-bold text-gray-900">记录</h1>
        <span class="text-gray-500">洞察</span>
      </div>
      <div class="flex items-center space-x-4">
        <div class="flex items-center space-x-1">
          <span class="text-gray-700">全部</span>
          <svg class="w-4 h-4 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path>
          </svg>
        </div>
        <div class="w-8 h-8 bg-gray-300 rounded-full flex items-center justify-center">
          <svg class="w-5 h-5 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
              d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"></path>
          </svg>
        </div>
      </div>
    </div>

    <!-- 日历组件 -->
    <div class="bg-blue-50 mx-4 mt-4 rounded-2xl p-4">
      <div class="flex justify-between items-center mb-2">
        <h2 class="text-lg font-semibold text-gray-800">{{ currentDate.year }}年{{ currentDate.month }}月</h2>
      </div>

      <!-- 日历网格 -->
      <div class="grid grid-cols-10 gap-1.5">
        <!-- 日期数字 -->
        <div v-for="day in calendarDays" :key="day.date" @click="handleDateClick(day)" :class="[
          'h-7 w-7 flex items-center justify-center text-sm font-medium rounded-lg transition-all duration-200',
          selectedDate === day.day ? 'bg-blue-500 text-white' :
            (day.isToday && selectedDate === null) ? 'bg-blue-500 text-white' :
              day.isFuture ? 'bg-white text-gray-400' : 'bg-white text-gray-800 hover:bg-blue-50 cursor-pointer active:scale-95'
        ]">
          {{ day.day }}
        </div>
      </div>
    </div>

    <!-- 今天日期 -->
    <div class="px-4 mt-6 mb-4">
      <div class="flex items-center space-x-2">
        <Icon name="material-symbols:calendar-today-outline-rounded" size="24" class="text-blue-500" />
        <span class="text-lg font-medium text-gray-800">今天 {{ todayText }}</span>
      </div>
      <!-- 今日收支统计 -->
      <div class="ml-8">
        <span class="text-sm text-gray-600">支出 ¥{{ todayExpense }} | 收入 ¥{{ todayIncome }}</span>
      </div>
    </div>

    <!-- 记录卡片 -->
    <div class="px-4 space-y-4">
      <div class="bg-blue-50 rounded-2xl p-4 flex items-start space-x-3">
        <div class="flex-shrink-0">
          <!-- 人物形象 -->
          <div class="w-28 h-38">
            <img src="https://wp-cdn.yukuii.top/v2/zu7m9gq.png" alt="人物形象" class="w-full h-full object-cover" />
          </div>
        </div>
        <div class="flex-1">
          <h3 class="text-blue-600 font-medium mb-2">记点什么好呢:</h3>
          <div class="space-y-1 text-sm text-gray-700">
            <p>"今天午饭30元，用的支付宝"</p>
            <p>"明天下午6点开会，提前5分钟提醒我"</p>
            <p>"抢到演唱会票了，激动到我睡不着！"</p>
          </div>
        </div>
      </div>

      <!-- 自动记录功能卡片 -->
      <div class="mb-6">
        <div class="flex items-center justify-between mb-3">
          <span class="text-blue-600 text-sm">自动记录也超方便 点我去体验 ></span>
        </div>
        <div class="grid grid-cols-3 gap-3">
          <!-- 睡眠截屏 -->
          <div class="bg-white rounded-2xl p-4 text-center shadow-sm">
            <div class="w-12 h-12 bg-blue-100 rounded-2xl flex items-center justify-center mx-auto mb-2 relative">
              <!-- 睡觉的人 -->
              <div class="w-8 h-6 bg-orange-300 rounded-full relative">
                <!-- 头部 -->
                <div class="w-4 h-4 bg-orange-400 rounded-full absolute -top-1 left-1"></div>
                <!-- Z字符表示睡觉 -->
                <div class="absolute -top-2 -right-1 text-xs text-blue-500 font-bold">Z</div>
              </div>
            </div>
            <p class="text-xs text-gray-600">就一般睡眠截屏</p>
          </div>

          <!-- 小白点截屏 -->
          <div class="bg-white rounded-2xl p-4 text-center shadow-sm">
            <div class="w-12 h-12 bg-blue-100 rounded-2xl flex items-center justify-center mx-auto mb-2 relative">
              <!-- 手指点击 -->
              <div class="w-6 h-8 bg-orange-300 rounded-t-full relative">
                <!-- 指尖 -->
                <div class="w-2 h-2 bg-orange-400 rounded-full absolute bottom-0 left-2"></div>
                <!-- 点击效果 -->
                <div class="absolute -bottom-1 left-1 w-4 h-1 bg-blue-500 rounded-full opacity-60"></div>
              </div>
            </div>
            <p class="text-xs text-gray-600">小白点截屏</p>
          </div>

          <!-- iPhone快捷键 -->
          <div class="bg-white rounded-2xl p-4 text-center shadow-sm">
            <div class="w-12 h-12 bg-blue-100 rounded-2xl flex items-center justify-center mx-auto mb-2 relative">
              <!-- 手机 -->
              <div class="w-6 h-9 bg-gray-700 rounded-lg relative">
                <!-- 屏幕 -->
                <div class="w-5 h-7 bg-blue-400 rounded-md absolute top-1 left-0.5"></div>
                <!-- 按键指示 -->
                <div class="absolute -top-1 -right-1 w-2 h-2 bg-orange-400 rounded-full"></div>
                <div class="absolute -bottom-1 -right-1 w-2 h-2 bg-orange-400 rounded-full"></div>
              </div>
            </div>
            <p class="text-xs text-gray-600">iPhone快捷键</p>
          </div>
        </div>
      </div>
    </div>

    <!-- 底部上拉指示器 -->
    <div class="flex justify-center py-6">
      <svg class="w-6 h-6 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 15l7-7 7 7"></path>
      </svg>
    </div>

    <!-- 底部输入区域 -->
    <div class="fixed bottom-0 left-0 right-0 bg-white border-t border-gray-200 p-4">
      <div class="flex items-center space-x-3">
        <button class="w-8 h-8 bg-gray-100 rounded-full flex items-center justify-center">
          <svg class="w-5 h-5 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
          </svg>
        </button>
        <div class="flex-1 bg-gray-100 rounded-full px-4 py-2">
          <input type="text" placeholder="一句话轻松记录~"
            class="w-full bg-transparent text-gray-600 placeholder-gray-400 outline-none">
        </div>
        <button class="w-10 h-10 bg-blue-500 rounded-full flex items-center justify-center">
          <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
              d="M19 11a7 7 0 01-7 7m0 0a7 7 0 01-7-7m7 7v4m0 0H8m4 0h4m-4-8a3 3 0 01-3-3V5a3 3 0 116 0v6a3 3 0 01-3 3z">
            </path>
          </svg>
        </button>
      </div>

      <!-- 底部指示条 -->
      <div class="flex justify-center mt-2">
        <div class="w-32 h-1 bg-black rounded-full"></div>
      </div>
    </div>

    <!-- 底部安全区域 -->
    <div class="h-20"></div>
  </div>
</template>

<script setup>
// 使用移动端布局
definePageMeta({
  layout: 'mobile'
})

// 设置页面元信息
useHead({
  title: '记录 - 个人记账应用',
  meta: [
    { name: 'description', content: '个人记账应用首页，轻松记录您的财务' },
    { name: 'viewport', content: 'width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=no' }
  ]
})

// 获取当前日期信息
const getCurrentDate = () => {
  const now = new Date()
  return {
    year: now.getFullYear(),
    month: now.getMonth() + 1, // getMonth() 返回 0-11，需要 +1
    day: now.getDate(),
    weekDay: now.getDay() // 0=周日, 1=周一, ..., 6=周六
  }
}

// 获取星期几的中文表示
const getWeekDayText = (weekDay) => {
  const weekDays = ['周日', '周一', '周二', '周三', '周四', '周五', '周六']
  return weekDays[weekDay]
}

// 生成今天的文本显示
const currentDate = getCurrentDate()
const todayText = `${currentDate.month}月${currentDate.day}日(${getWeekDayText(currentDate.weekDay)})`

// 生成日历数据
const generateCalendarDays = () => {
  const { year, month, day: today } = currentDate

  // 获取当月最后一天，确定这个月有多少天
  const lastDay = new Date(year, month, 0)
  const daysInMonth = lastDay.getDate()

  const days = []

  // 只添加当月的日期
  for (let day = 1; day <= daysInMonth; day++) {
    days.push({
      day,
      date: `${year}-${month}-${day}`,
      isCurrentMonth: true,
      isToday: day === today,
      isFuture: day > today // 判断是否是未来日期
    })
  }

  return days
}

// 日历数据
const calendarDays = ref(generateCalendarDays())

// 选中的日期
const selectedDate = ref(null)

// 今日收支数据
const todayExpense = ref('0.00')
const todayIncome = ref('0.00')

// 处理日期点击事件
const handleDateClick = (day) => {
  // 只有今天和之前的日期可以点击
  if (day.isFuture) {
    return // 未来日期不可点击
  }

  // 设置选中的日期
  selectedDate.value = day.day

  console.log('点击了日期:', day.date)
  console.log('日期信息:', {
    年: currentDate.year,
    月: currentDate.month,
    日: day.day,
    是否今天: day.isToday,
    选中状态: selectedDate.value === day.day
  })

  // 可以在这里添加更多功能，比如：
  // navigateTo(`/records/${day.date}`) // 跳转到该日期的记录页面
  // 或者触发一个事件来显示该日期的记录
}
</script>

<style scoped>
/* 移动端优化样式 */
@media (max-width: 768px) {
  .min-h-screen {
    min-height: 100vh;
    min-height: 100dvh;
    /* 动态视口高度，更适合移动端 */
  }
}

/* 确保在移动端有合适的字体大小 */
html {
  font-size: 16px;
}

/* 优化触摸体验 */
button,
.cursor-pointer {
  touch-action: manipulation;
}

/* 防止文本选择 */
.select-none {
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
}
</style>