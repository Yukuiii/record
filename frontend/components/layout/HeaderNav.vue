<template>
  <!-- 头部导航 -->
  <div class="bg-white px-4 py-3 flex justify-between items-center border-b border-gray-100">
    <div class="flex items-center space-x-6">
      <h1 class="text-xl font-bold text-gray-900">记录</h1>
      <!-- <span class="text-gray-500">洞察</span> -->
    </div>
    <div class="flex items-center space-x-4">
      <!-- <div class="flex items-center space-x-1">
        <span class="text-gray-700">全部</span>
        <svg class="w-4 h-4 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path>
        </svg>
      </div> -->
      <div class="w-8 h-8 bg-gray-300 rounded-full flex items-center justify-center">
        <svg class="w-5 h-5 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
            d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"></path>
        </svg>
      </div>
    </div>
  </div>
</template>

<script setup>
// 头部导航组件
// 可以在这里添加导航相关的逻辑，比如：
// - 用户头像点击事件
// - 筛选下拉菜单
// - 导航切换逻辑等

// 如果需要接收父组件传递的数据，可以使用 defineProps
// const props = defineProps({
//   activeTab: {
//     type: String,
//     default: '记录'
//   }
// })

// 如果需要向父组件发送事件，可以使用 defineEmits
// const emit = defineEmits(['tab-change', 'user-click'])

// 处理用户头像点击
const handleUserClick = () => {
  console.log('用户头像被点击')
  // emit('user-click')
}

// 处理筛选下拉点击
const handleFilterClick = () => {
  console.log('筛选下拉被点击')
  // 这里可以添加筛选逻辑
}
</script>

<style scoped>
/* 头部导航样式 */
/* 如果需要特殊样式可以在这里添加 */
</style>
