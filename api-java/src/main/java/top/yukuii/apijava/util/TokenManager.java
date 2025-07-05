package top.yukuii.apijava.util;

import java.util.Date;
import java.util.HashMap;
import java.util.Map;

import org.springframework.web.context.request.RequestContextHolder;
import org.springframework.web.context.request.ServletRequestAttributes;

import jakarta.servlet.http.HttpServletRequest;
import lombok.extern.slf4j.Slf4j;
import top.yukuii.apijava.common.BusinessException;
import top.yukuii.apijava.model.entity.User;
import top.yukuii.apijava.model.vo.LoginResponseVO;

/**
 * Token管理工具类
 * 统一封装所有Token相关操作
 */
@Slf4j
public class TokenManager {

    /**
     * 用户登录 - 创建Token
     * 
     * @param user 用户信息
     * @return 登录响应对象
     */
    public static LoginResponseVO login(User user) {
        return login(user, null);
    }

    /**
     * 用户登录 - 创建Token（带额外声明）
     * 
     * @param user 用户信息
     * @param extraClaims 额外声明
     * @return 登录响应对象
     */
    public static LoginResponseVO login(User user, Map<String, Object> extraClaims) {
        try {
            // 创建Token
            String token = JwtUtil.createToken(user.getUserId(), extraClaims);
            
            log.info("用户登录成功，用户ID: {}", user.getUserId());
            
            // 构建响应
            return LoginResponseVO.builder()
                    .token(token)
                    .tokenType("Bearer")
                    .expiresAt(System.currentTimeMillis() + 24 * 60 * 60 * 1000L)
                    .userInfo(LoginResponseVO.UserInfoVO.builder()
                            .userId(user.getUserId())
                            .userName(user.getUserName())
                            .email(user.getEmail())
                            .phone(user.getPhone())
                            .avatar(user.getAvatar())
                            .status(user.getStatus())
                            .build())
                    .build();
        } catch (Exception e) {
            log.error("创建Token失败，用户ID: {}", user.getUserId(), e);
            throw new BusinessException("登录失败");
        }
    }

    /**
     * 用户登出 - 将Token加入黑名单
     * 
     * @param token JWT Token
     */
    public static void logout(String token) {
        try {
            // 验证Token是否有效
            if (!isTokenValid(token)) {
                throw new BusinessException("无效的Token");
            }

            // 检查Token是否已经在黑名单中
            if (TokenBlacklistUtil.isBlacklisted(token)) {
                throw new BusinessException("Token已失效");
            }

            // 将Token加入黑名单
            TokenBlacklistUtil.addToBlacklist(token);
            
            // 获取用户ID用于日志记录
            String userId = getUserIdFromToken(token);
            log.info("用户登出成功，用户ID: {}", userId);
            
        } catch (BusinessException e) {
            throw e;
        } catch (Exception e) {
            log.error("用户登出失败", e);
            throw new BusinessException("登出失败");
        }
    }

    /**
     * 刷新Token
     * 
     * @param oldToken 原Token
     * @return 新的登录响应对象
     */
    public static LoginResponseVO refreshToken(String oldToken) {
        try {
            // 验证原Token是否有效
            if (!isTokenValid(oldToken)) {
                throw new BusinessException("Token已失效，请重新登录");
            }

            // 刷新Token
            String newToken = JwtUtil.refreshToken(oldToken);
            
            // 构建响应
            return LoginResponseVO.builder()
                    .token(newToken)
                    .tokenType("Bearer")
                    .expiresAt(System.currentTimeMillis() + 24 * 60 * 60 * 1000L)
                    .build();
        } catch (Exception e) {
            log.error("Token刷新失败", e);
            throw new BusinessException("Token刷新失败");
        }
    }

    /**
     * 验证Token是否有效（包括黑名单检查）
     * 
     * @param token JWT Token
     * @return 是否有效
     */
    public static boolean isTokenValid(String token) {
        try {
            // 1. 验证Token格式和签名
            if (!JwtUtil.validateToken(token)) {
                return false;
            }

            // 2. 检查是否在黑名单中
            if (TokenBlacklistUtil.isBlacklisted(token)) {
                return false;
            }

            // 3. 检查是否过期
            if (JwtUtil.isTokenExpired(token)) {
                return false;
            }

            return true;
        } catch (Exception e) {
            log.warn("Token验证失败: {}", e.getMessage());
            return false;
        }
    }

    /**
     * 从Token中获取用户ID
     * 
     * @param token JWT Token
     * @return 用户ID
     */
    public static String getUserIdFromToken(String token) {
        try {
            return JwtUtil.getUserId(token);
        } catch (Exception e) {
            log.warn("从Token获取用户ID失败: {}", e.getMessage());
            throw new BusinessException("Token解析失败");
        }
    }

    /**
     * 获取当前请求的Token
     *
     * @return 当前请求的JWT Token，如果没有返回null
     */
    public static String getCurrentToken() {
        try {
            ServletRequestAttributes attributes = (ServletRequestAttributes) RequestContextHolder.getRequestAttributes();
            if (attributes != null) {
                HttpServletRequest request = attributes.getRequest();
                String authHeader = request.getHeader("Authorization");
                return extractToken(authHeader);
            }
        } catch (Exception e) {
            log.warn("获取当前Token失败: {}", e.getMessage());
        }
        return null;
    }

    /**
     * 获取当前登录用户ID
     *
     * @return 用户ID，如果未登录返回null
     */
    public static String getCurrentUserId() {
        return UserContext.getCurrentUserId();
    }

    /**
     * 获取当前登录用户ID（必须存在）
     * 
     * @return 用户ID
     * @throws BusinessException 如果用户未登录
     */
    public static String requireCurrentUserId() {
        String userId = getCurrentUserId();
        if (userId == null) {
            throw new BusinessException("用户未登录");
        }
        return userId;
    }

    /**
     * 检查当前是否已登录
     *
     * @return true表示已登录，false表示未登录
     */
    public static boolean isLoggedIn() {
        return UserContext.isLoggedIn();
    }

    /**
     * 获取当前Token的过期时间
     *
     * @return 过期时间，如果获取失败返回null
     */
    public static Date getCurrentTokenExpiration() {
        String token = getCurrentToken();
        if (token != null) {
            try {
                return getTokenExpiration(token);
            } catch (Exception e) {
                log.warn("获取当前Token过期时间失败: {}", e.getMessage());
            }
        }
        return null;
    }

    /**
     * 获取当前用户的角色信息
     *
     * @return 角色数组，如果获取失败返回空数组
     */
    public static String[] getCurrentUserRoles() {
        String token = getCurrentToken();
        if (token != null) {
            return getRolesFromToken(token);
        }
        return new String[0];
    }

    /**
     * 刷新当前用户的Token
     *
     * @return 新的登录响应对象
     * @throws BusinessException 如果刷新失败
     */
    public static LoginResponseVO refreshCurrentToken() {
        String token = getCurrentToken();
        if (token == null) {
            throw new BusinessException("当前请求中没有Token");
        }
        return refreshToken(token);
    }

    /**
     * 登出当前用户
     *
     * @throws BusinessException 如果登出失败
     */
    public static void logoutCurrent() {
        String token = getCurrentToken();
        if (token == null) {
            throw new BusinessException("当前请求中没有Token");
        }
        logout(token);
    }

    /**
     * 从请求头提取Token
     * 
     * @param authHeader Authorization头的值
     * @return 提取的Token，如果格式不正确返回null
     */
    public static String extractToken(String authHeader) {
        return JwtUtil.extractToken(authHeader);
    }

    /**
     * 获取Token的过期时间
     * 
     * @param token JWT Token
     * @return 过期时间
     */
    public static Date getTokenExpiration(String token) {
        try {
            return JwtUtil.getExpiration(token);
        } catch (Exception e) {
            log.warn("获取Token过期时间失败: {}", e.getMessage());
            throw new BusinessException("Token解析失败");
        }
    }

    /**
     * 创建带有角色信息的Token
     * 
     * @param user 用户信息
     * @param roles 用户角色列表
     * @return 登录响应对象
     */
    public static LoginResponseVO loginWithRoles(User user, String... roles) {
        Map<String, Object> claims = new HashMap<>();
        claims.put("roles", roles);
        return login(user, claims);
    }

    /**
     * 从Token中获取角色信息
     * 
     * @param token JWT Token
     * @return 角色数组
     */
    @SuppressWarnings("unchecked")
    public static String[] getRolesFromToken(String token) {
        try {
            var claims = JwtUtil.parseToken(token);
            var roles = claims.get("roles");
            if (roles instanceof String[]) {
                return (String[]) roles;
            }
            return new String[0];
        } catch (Exception e) {
            log.warn("从Token获取角色信息失败: {}", e.getMessage());
            return new String[0];
        }
    }

    /**
     * 私有构造函数，防止实例化
     */
    private TokenManager() {
        throw new UnsupportedOperationException("This is a utility class and cannot be instantiated");
    }
}
