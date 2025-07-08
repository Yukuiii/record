/**
 * 全局通知管理组合式函数
 */

export const useNotification = () => {
  // 通知状态
  const notifications = useState("notifications", () => []);

  // 添加通知
  const addNotification = (notification) => {
    const id = Date.now() + Math.random();
    const newNotification = {
      id,
      type: "info", // info, success, warning, error
      message: "",
      duration: 3000, // 显示时长，0 表示不自动关闭
      ...notification,
    };

    notifications.value.push(newNotification);

    // 自动移除通知
    if (newNotification.duration > 0) {
      setTimeout(() => {
        removeNotification(id);
      }, newNotification.duration);
    }

    return id;
  };

  // 移除通知
  const removeNotification = (id) => {
    const index = notifications.value.findIndex((n) => n.id === id);
    if (index > -1) {
      notifications.value.splice(index, 1);
    }
  };

  // 清除所有通知
  const clearNotifications = () => {
    notifications.value = [];
  };

  // 便捷方法
  const showSuccess = (message, duration = 3000) => {
    return addNotification({ type: "success", message, duration });
  };

  const showError = (message, duration = 5000) => {
    return addNotification({ type: "error", message, duration });
  };

  const showWarning = (message, duration = 4000) => {
    return addNotification({ type: "warning", message, duration });
  };

  const showInfo = (message, duration = 3000) => {
    return addNotification({ type: "info", message, duration });
  };

  return {
    notifications: readonly(notifications),
    addNotification,
    removeNotification,
    clearNotifications,
    showSuccess,
    showError,
    showWarning,
    showInfo,
  };
};
