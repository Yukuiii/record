package top.yukuii.apijava.controller;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import lombok.RequiredArgsConstructor;
import top.yukuii.apijava.common.Result;
import top.yukuii.apijava.util.TokenManager;

/**
 * 用户控制器
 * 演示如何使用JWT拦截器和用户上下文
 */
@RestController
@RequestMapping("/api/user")
@RequiredArgsConstructor
public class UserController {

    private final TokenManager tokenManager;

    /**
     * 获取当前用户信息
     * 拦截器会自动验证Token，并将用户ID存储到请求属性中
     */
    @GetMapping("/profile")
    public Result<String> getCurrentUserProfile() {
        // 直接从TokenManager获取当前用户ID，无需手动验证Token
        String userId = tokenManager.getCurrentUserId();

        // TODO: 根据userId查询用户详细信息
        return Result.success("当前用户ID: " + userId);
    }

    /**
     * 更新用户信息
     */
    @PutMapping("/profile")
    public Result<String> updateProfile(@RequestBody String updateData) {
        // 获取当前用户ID（必须存在）
        String userId = tokenManager.requireCurrentUserId();

        // TODO: 更新用户信息逻辑
        return Result.success("用户 " + userId + " 信息更新成功");
    }

    /**
     * 获取用户设置
     */
    @GetMapping("/settings")
    public Result<String> getUserSettings() {
        String userId = tokenManager.getCurrentUserId();

        // TODO: 查询用户设置
        return Result.success("用户 " + userId + " 的设置信息");
    }

    /**
     * 检查登录状态
     */
    @GetMapping("/status")
    public Result<Boolean> checkLoginStatus() {
        boolean isLoggedIn = tokenManager.isLoggedIn();
        return Result.success(isLoggedIn);
    }
}
