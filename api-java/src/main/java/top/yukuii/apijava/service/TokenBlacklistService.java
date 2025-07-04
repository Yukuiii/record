package top.yukuii.apijava.service;

import java.util.Date;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.Executors;
import java.util.concurrent.ScheduledExecutorService;
import java.util.concurrent.TimeUnit;

import org.springframework.stereotype.Service;

import lombok.extern.slf4j.Slf4j;
import top.yukuii.apijava.util.JwtUtil;

/**
 * Token黑名单服务
 * 用于管理已登出的Token
 */
@Slf4j
@Service
public class TokenBlacklistService {

    /**
     * Token黑名单存储（生产环境建议使用Redis）
     * Key: Token, Value: 过期时间
     */
    private final ConcurrentHashMap<String, Long> blacklist = new ConcurrentHashMap<>();

    /**
     * 定时清理过期Token的调度器
     */
    private final ScheduledExecutorService scheduler = Executors.newScheduledThreadPool(1);

    public TokenBlacklistService() {
        // 每小时清理一次过期的Token
        scheduler.scheduleAtFixedRate(this::cleanExpiredTokens, 1, 1, TimeUnit.HOURS);
    }

    /**
     * 将Token加入黑名单
     * 
     * @param token JWT Token
     */
    public void addToBlacklist(String token) {
        try {
            // 获取Token的过期时间
            Date expiration = JwtUtil.getExpiration(token);
            blacklist.put(token, expiration.getTime());
            log.info("Token已加入黑名单，过期时间: {}", expiration);
        } catch (Exception e) {
            log.warn("添加Token到黑名单失败: {}", e.getMessage());
            // 如果无法获取过期时间，设置默认过期时间（24小时后）
            long defaultExpiration = System.currentTimeMillis() + 24 * 60 * 60 * 1000L;
            blacklist.put(token, defaultExpiration);
        }
    }

    /**
     * 检查Token是否在黑名单中
     * 
     * @param token JWT Token
     * @return true表示在黑名单中（已登出），false表示不在黑名单中
     */
    public boolean isBlacklisted(String token) {
        Long expiration = blacklist.get(token);
        if (expiration == null) {
            return false;
        }
        
        // 如果Token已过期，从黑名单中移除
        if (expiration < System.currentTimeMillis()) {
            blacklist.remove(token);
            return false;
        }
        
        return true;
    }

    /**
     * 清理过期的Token
     */
    private void cleanExpiredTokens() {
        long currentTime = System.currentTimeMillis();
        int sizeBefore = blacklist.size();

        // 清理过期的Token
        blacklist.entrySet().removeIf(entry -> entry.getValue() < currentTime);

        int sizeAfter = blacklist.size();
        int removedCount = sizeBefore - sizeAfter;

        if (removedCount > 0) {
            log.info("清理了 {} 个过期的黑名单Token", removedCount);
        }
    }

    /**
     * 获取黑名单大小（用于监控）
     */
    public int getBlacklistSize() {
        return blacklist.size();
    }

    /**
     * 清空黑名单（谨慎使用）
     */
    public void clearBlacklist() {
        blacklist.clear();
        log.warn("黑名单已被清空");
    }
}
