package top.yukuii.apijava.util;

import org.springframework.web.context.request.RequestContextHolder;
import org.springframework.web.context.request.ServletRequestAttributes;

import jakarta.servlet.http.HttpServletRequest;

/**
 * 用户上下文工具类
 * 用于在Controller中方便地获取当前用户信息
 */
public class UserContext {

    /**
     * 获取当前用户ID（可能为null）
     * 
     * @return 用户ID，如果未登录返回null
     */
    public static String getCurrentUserId() {
        try {
            ServletRequestAttributes attributes = (ServletRequestAttributes) RequestContextHolder.getRequestAttributes();
            if (attributes != null) {
                HttpServletRequest request = attributes.getRequest();
                return (String) request.getAttribute("userId");
            }
        } catch (Exception e) {}
        return null;
    }

    /**
     * 检查是否已登录
     * 
     * @return true表示已登录，false表示未登录
     */
    public static boolean isLoggedIn() {
        return getCurrentUserId() != null;
    }
}
