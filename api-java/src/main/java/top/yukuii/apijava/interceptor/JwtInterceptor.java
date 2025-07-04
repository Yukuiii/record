package top.yukuii.apijava.interceptor;

import org.springframework.stereotype.Component;
import org.springframework.web.servlet.HandlerInterceptor;

import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import top.yukuii.apijava.service.AuthService;
import top.yukuii.apijava.util.JwtUtil;

/**
 * JWT拦截器
 * 自动验证请求中的Token
 */
@Slf4j
@Component
@RequiredArgsConstructor
public class JwtInterceptor implements HandlerInterceptor {

    private final AuthService authService;

    @Override
    public boolean preHandle(HttpServletRequest request, HttpServletResponse response, Object handler) throws Exception {
        // 获取请求路径
        String requestURI = request.getRequestURI();
        log.debug("拦截请求: {}", requestURI);

        // 从请求头获取Token
        String authHeader = request.getHeader("Authorization");
        String token = JwtUtil.extractToken(authHeader);

        if (token == null) {
            log.warn("请求缺少Token: {}", requestURI);
            sendErrorResponse(response, 401, "缺少认证Token");
            return false;
        }

        // 验证Token
        if (!authService.isTokenValid(token)) {
            log.warn("Token验证失败: {}", requestURI);
            sendErrorResponse(response, 401, "Token无效或已过期");
            return false;
        }

        // 将用户ID存储到请求属性中，供Controller使用
        try {
            String userId = JwtUtil.getUserId(token);
            request.setAttribute("userId", userId);
            log.debug("Token验证成功，用户ID: {}", userId);
        } catch (Exception e) {
            log.error("获取用户ID失败", e);
            sendErrorResponse(response, 401, "Token解析失败");
            return false;
        }

        return true;
    }

    /**
     * 发送错误响应
     */
    private void sendErrorResponse(HttpServletResponse response, int status, String message) throws Exception {
        response.setStatus(status);
        response.setContentType("application/json;charset=UTF-8");
        
        String jsonResponse = String.format(
            "{\"code\":%d,\"message\":\"%s\",\"data\":null}",
            status, message
        );
        
        response.getWriter().write(jsonResponse);
    }
}
