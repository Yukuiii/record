<template>
  <div class="fixed top-4 right-4 z-50 space-y-2">
    <TransitionGroup
      name="notification"
      tag="div"
      class="space-y-2"
    >
      <div
        v-for="notification in notifications"
        :key="notification.id"
        :class="getNotificationClass(notification.type)"
        class="min-w-80 max-w-md p-4 rounded-lg shadow-lg backdrop-blur-sm border flex items-start gap-3"
      >
        <!-- 图标 -->
        <div class="flex-shrink-0 mt-0.5">
          <Icon
            :name="getIconName(notification.type)"
            :class="getIconClass(notification.type)"
            size="20"
          />
        </div>
        
        <!-- 消息内容 -->
        <div class="flex-1 min-w-0">
          <p class="text-sm font-medium text-gray-900">
            {{ notification.message }}
          </p>
        </div>
        
        <!-- 关闭按钮 -->
        <button
          @click="removeNotification(notification.id)"
          class="flex-shrink-0 text-gray-400 hover:text-gray-600 transition-colors"
        >
          <Icon name="material-symbols:close" size="16" />
        </button>
      </div>
    </TransitionGroup>
  </div>
</template>

<script setup>
const { notifications, removeNotification } = useNotification();

// 获取通知样式类
const getNotificationClass = (type) => {
  const baseClass = 'bg-white/90 border-l-4';
  
  switch (type) {
    case 'success':
      return `${baseClass} border-l-green-500 border-green-200`;
    case 'error':
      return `${baseClass} border-l-red-500 border-red-200`;
    case 'warning':
      return `${baseClass} border-l-yellow-500 border-yellow-200`;
    case 'info':
    default:
      return `${baseClass} border-l-blue-500 border-blue-200`;
  }
};

// 获取图标名称
const getIconName = (type) => {
  switch (type) {
    case 'success':
      return 'material-symbols:check-circle';
    case 'error':
      return 'material-symbols:error';
    case 'warning':
      return 'material-symbols:warning';
    case 'info':
    default:
      return 'material-symbols:info';
  }
};

// 获取图标样式类
const getIconClass = (type) => {
  switch (type) {
    case 'success':
      return 'text-green-500';
    case 'error':
      return 'text-red-500';
    case 'warning':
      return 'text-yellow-500';
    case 'info':
    default:
      return 'text-blue-500';
  }
};
</script>

<style scoped>
.notification-enter-active,
.notification-leave-active {
  transition: all 0.3s ease;
}

.notification-enter-from {
  opacity: 0;
  transform: translateX(100%);
}

.notification-leave-to {
  opacity: 0;
  transform: translateX(100%);
}

.notification-move {
  transition: transform 0.3s ease;
}
</style>