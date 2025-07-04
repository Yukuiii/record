package top.yukuii.apijava.util;

import java.util.Date;
import java.util.HashMap;
import java.util.Map;

import javax.crypto.SecretKey;

import io.jsonwebtoken.Claims;
import io.jsonwebtoken.ExpiredJwtException;
import io.jsonwebtoken.JwtBuilder;
import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.MalformedJwtException;
import io.jsonwebtoken.UnsupportedJwtException;
import io.jsonwebtoken.security.Keys;
import lombok.extern.slf4j.Slf4j;

/**
 * JWT工具类
 * 基于jjwt库实现
 */
@Slf4j
public class JwtUtil {

    /**
     * JWT密钥 - 生产环境应该从配置文件读取
     */
    private static final String SECRET_KEY = "your-super-secret-key-for-jwt-token-generation-must-be-at-least-256-bits";

    /**
     * Token过期时间（毫秒）- 24小时
     */
    private static final long EXPIRE_TIME = 24 * 60 * 60 * 1000L;

    /**
     * Token前缀
     */
    public static final String TOKEN_PREFIX = "Bearer ";

    /**
     * Header中Token的key
     */
    public static final String HEADER_TOKEN_KEY = "Authorization";

    /**
     * 生成安全密钥
     */
    private static SecretKey getSecretKey() {
        return Keys.hmacShaKeyFor(SECRET_KEY.getBytes());
    }

    /**
     * 创建JWT Token
     *
     * @param userId 用户ID
     * @return JWT Token
     */
    public static String createToken(String userId) {
        return createToken(userId, null);
    }

    /**
     * 创建JWT Token（带额外声明）
     *
     * @param userId 用户ID
     * @param claims 额外声明
     * @return JWT Token
     */
    public static String createToken(String userId, Map<String, Object> claims) {
        Date now = new Date();
        Date expiration = new Date(now.getTime() + EXPIRE_TIME);

        JwtBuilder builder = Jwts.builder()
                .subject(userId)
                .issuedAt(now)
                .expiration(expiration)
                .signWith(getSecretKey());

        // 添加额外声明
        if (claims != null && !claims.isEmpty()) {
            builder.claims(claims);
        }

        String token = builder.compact();
        log.debug("创建JWT Token成功，用户ID: {}", userId);
        return token;
    }

    /**
     * 验证JWT Token
     *
     * @param token JWT Token
     * @return 是否有效
     */
    public static boolean validateToken(String token) {
        try {
            parseToken(token);
            return true;
        } catch (Exception e) {
            log.warn("JWT Token验证失败: {}", e.getMessage());
            return false;
        }
    }

    /**
     * 解析JWT Token
     *
     * @param token JWT Token
     * @return Claims
     */
    public static Claims parseToken(String token) {
        try {
            return Jwts.parser()
                    .verifyWith(getSecretKey())
                    .build()
                    .parseSignedClaims(token)
                    .getPayload();
        } catch (ExpiredJwtException e) {
            log.warn("JWT Token已过期");
            throw new RuntimeException("Token已过期", e);
        } catch (UnsupportedJwtException e) {
            log.warn("不支持的JWT Token");
            throw new RuntimeException("不支持的Token", e);
        } catch (MalformedJwtException e) {
            log.warn("JWT Token格式错误");
            throw new RuntimeException("Token格式错误", e);
        } catch (SecurityException e) {
            log.warn("JWT Token签名验证失败");
            throw new RuntimeException("Token签名验证失败", e);
        } catch (IllegalArgumentException e) {
            log.warn("JWT Token参数错误");
            throw new RuntimeException("Token参数错误", e);
        }
    }
    /**
     * 从Token中获取用户ID
     *
     * @param token JWT Token
     * @return 用户ID
     */
    public static String getUserId(String token) {
        Claims claims = parseToken(token);
        return claims.getSubject();
    }

    /**
     * 从Token中获取过期时间
     *
     * @param token JWT Token
     * @return 过期时间
     */
    public static Date getExpiration(String token) {
        Claims claims = parseToken(token);
        return claims.getExpiration();
    }

    /**
     * 判断Token是否过期
     *
     * @param token JWT Token
     * @return 是否过期
     */
    public static boolean isTokenExpired(String token) {
        try {
            Date expiration = getExpiration(token);
            return expiration.before(new Date());
        } catch (Exception e) {
            return true;
        }
    }

    /**
     * 刷新Token（重新生成）
     *
     * @param token 原Token
     * @return 新Token
     */
    public static String refreshToken(String token) {
        try {
            Claims claims = parseToken(token);
            String userId = claims.getSubject();

            // 移除时间相关的声明
            Map<String, Object> newClaims = new HashMap<>(claims);
            newClaims.remove("iat");
            newClaims.remove("exp");
            newClaims.remove("sub");

            return createToken(userId, newClaims);
        } catch (Exception e) {
            log.error("刷新Token失败", e);
            throw new RuntimeException("刷新Token失败", e);
        }
    }

    /**
     * 从请求头中提取Token
     *
     * @param authHeader Authorization头的值
     * @return 提取的Token，如果格式不正确返回null
     */
    public static String extractToken(String authHeader) {
        if (authHeader != null && authHeader.startsWith(TOKEN_PREFIX)) {
            return authHeader.substring(TOKEN_PREFIX.length());
        }
        return null;
    }
}
