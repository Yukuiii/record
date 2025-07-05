package top.yukuii.apijava.controller;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import top.yukuii.apijava.common.Result;
import top.yukuii.apijava.model.dto.LoginRequestDTO;
import top.yukuii.apijava.model.dto.RegisterRequestDTO;
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
     * 用户注册
     */
    @PostMapping("/register")
    public Result<LoginResponseVO> register(@RequestBody @Valid RegisterRequestDTO request) {
        LoginResponseVO response = authService.register(request);
        return Result.success(response);
    }

    /**
     * 用户登出
     */
    @PostMapping("/logout")
    public Result<String> logout() {
        // 直接登出当前用户，无需传递任何参数
        TokenManager.logoutCurrent();
        return Result.success("登出成功");
    }

    /**
     * 验证Token是否有效
     * 注意：此接口不被拦截器拦截，用于前端检查Token状态
     */
    @PostMapping("/validate")
    public Result<Boolean> validateToken() {
        // 获取当前Token
        String token = TokenManager.getCurrentToken();
        if (token == null) {
            return Result.success(false);
        }

        // 验证Token
        boolean isValid = TokenManager.isTokenValid(token);
        return Result.success(isValid);
    }

    /**
     * 刷新Token
     */
    @PostMapping("/refresh")
    public Result<LoginResponseVO> refresh() {
        try {
            // 直接刷新当前用户Token，无需传递任何参数
            LoginResponseVO response = TokenManager.refreshCurrentToken();
            return Result.success(response);
        } catch (Exception e) {
            return Result.error("Token刷新失败");
        }
    }

    /**
     * 检查用户名是否可用
     */
    @GetMapping("/check-username")
    public Result<Boolean> checkUserName(@RequestParam String userName) {
        boolean available = authService.isUserNameAvailable(userName);
        return Result.success(available);
    }

    /**
     * 检查邮箱是否可用
     */
    @GetMapping("/check-email")
    public Result<Boolean> checkEmail(@RequestParam String email) {
        boolean available = authService.isEmailAvailable(email);
        return Result.success(available);
    }

    /**
     * 检查手机号是否可用
     */
    @GetMapping("/check-phone")
    public Result<Boolean> checkPhone(@RequestParam String phone) {
        boolean available = authService.isPhoneAvailable(phone);
        return Result.success(available);
    }
}
