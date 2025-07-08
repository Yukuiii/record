<template>
  <div>
    <NuxtPage />
  </div>
</template>

<script setup>
// 应用启动时检查认证状态
const { checkAuth } = useAuth()

// 在应用挂载时检查认证状态
onMounted(async () => {
  await checkAuth()
})
</script>
