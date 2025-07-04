package top.yukuii.apijava.model.vo;

import lombok.Builder;
import lombok.Data;

@Data
@Builder
public class LoginResponseVO {

    /**
     * JWT Token
     */
    private String token;

    /**
     * Token类型
     */
    private String tokenType;

    /**
     * Token过期时间（毫秒时间戳）
     */
    private Long expiresAt;

    /**
     * 用户信息
     */
    private UserInfoVO userInfo;

    @Data
    @Builder
    public static class UserInfoVO {
        /**
         * 用户ID
         */
        private String userId;

        /**
         * 用户名
         */
        private String userName;

        /**
         * 邮箱
         */
        private String email;

        /**
         * 手机号
         */
        private String phone;

        /**
         * 头像
         */
        private String avatar;

        /**
         * 状态
         */
        private Integer status;
    }
}
