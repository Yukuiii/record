package top.yukuii.apijava.controller;

import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import top.yukuii.apijava.common.Result;
import top.yukuii.apijava.model.dto.LoginRequestDTO;
import top.yukuii.apijava.model.vo.LoginResponseVO;
import top.yukuii.apijava.service.AuthService;
import top.yukuii.apijava.util.TokenManager;

/**
 * 认证控制器
 */
@RestController
@RequestMapping("/api/auth")
@RequiredArgsConstructor
public class AuthController {

    private final AuthService authService;
    private final TokenManager tokenManager;

    /**
     * 用户登录
     * 支持用户名、邮箱、手机号登录
     */
    @PostMapping("/login")
    public Result<LoginResponseVO> login(@RequestBody @Valid LoginRequestDTO request) {
        LoginResponseVO response = authService.login(request);
        return Result.success(response);
    }

    /**
     * 用户登出
     */
    @PostMapping("/logout")
    public Result<String> logout() {
        // 直接登出当前用户，无需传递任何参数
        tokenManager.logoutCurrent();
        return Result.success("登出成功");
    }

    /**
     * 验证Token是否有效
     * 注意：此接口不被拦截器拦截，用于前端检查Token状态
     */
    @PostMapping("/validate")
    public Result<Boolean> validateToken() {
        // 获取当前Token
        String token = tokenManager.getCurrentToken();
        if (token == null) {
            return Result.success(false);
        }

        // 验证Token
        boolean isValid = tokenManager.isTokenValid(token);
        return Result.success(isValid);
    }

    /**
     * 刷新Token
     */
    @PostMapping("/refresh")
    public Result<LoginResponseVO> refresh() {
        try {
            // 直接刷新当前用户Token，无需传递任何参数
            LoginResponseVO response = tokenManager.refreshCurrentToken();
            return Result.success(response);
        } catch (Exception e) {
            return Result.error("Token刷新失败");
        }
    }
}
