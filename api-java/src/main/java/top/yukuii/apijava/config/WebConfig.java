package top.yukuii.apijava.config;

import org.springframework.context.annotation.Configuration;
import org.springframework.web.servlet.config.annotation.InterceptorRegistry;
import org.springframework.web.servlet.config.annotation.WebMvcConfigurer;

import lombok.RequiredArgsConstructor;
import top.yukuii.apijava.interceptor.JwtInterceptor;

/**
 * Web配置类
 * 配置拦截器和其他Web相关设置
 */
@Configuration
@RequiredArgsConstructor
public class WebConfig implements WebMvcConfigurer {

    private final JwtInterceptor jwtInterceptor;

    @Override
    public void addInterceptors(InterceptorRegistry registry) {
        registry.addInterceptor(jwtInterceptor)
                // 拦截所有API请求
                .addPathPatterns("/api/**")
                // 排除认证相关接口
                .excludePathPatterns(
                        "/api/auth/login",          // 登录接口
                        "/api/auth/register",       // 注册接口
                        "/api/auth/validate",       // Token验证接口
                        "/api/auth/check-username", // 检查用户名接口
                        "/api/auth/check-email",    // 检查邮箱接口
                        "/api/auth/check-phone",    // 检查手机号接口
                        "/api/health",              // 健康检查接口
                        "/api/public/**"            // 公开接口
                );
    }
}
