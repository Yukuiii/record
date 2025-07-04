package top.yukuii.apijava.service;

import org.springframework.stereotype.Service;

import com.baomidou.mybatisplus.core.conditions.query.LambdaQueryWrapper;

import cn.hutool.core.util.ObjectUtil;
import cn.hutool.crypto.digest.BCrypt;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import top.yukuii.apijava.common.BusinessException;
import top.yukuii.apijava.mapper.UserMapper;
import top.yukuii.apijava.model.dto.LoginRequestDTO;
import top.yukuii.apijava.model.entity.User;
import top.yukuii.apijava.model.vo.LoginResponseVO;
import top.yukuii.apijava.util.TokenManager;

@Slf4j
@Service
@RequiredArgsConstructor
public class AuthService {

    private final UserMapper userMapper;
    private final TokenManager tokenManager;

    public LoginResponseVO login(LoginRequestDTO request) {
        // 1. 参数验证
        if (ObjectUtil.isEmpty(request.getUserName()) || ObjectUtil.isEmpty(request.getPassword())) {
            throw new BusinessException("用户名和密码不能为空");
        }

        // 2. 查询用户（支持用户名、邮箱、手机号登录）
        User user = userMapper.selectOne(
            new LambdaQueryWrapper<User>()
                .eq(User::getUserName, request.getUserName())
                .or()
                .eq(User::getEmail, request.getUserName())
                .or()
                .eq(User::getPhone, request.getUserName())
        );

        if (ObjectUtil.isEmpty(user)) {
            throw new BusinessException("用户不存在");
        }

        // 3. 检查用户状态
        if (user.getStatus() == null || user.getStatus() != 1) {
            throw new BusinessException("用户已被禁用");
        }

        // 4. 验证密码
        if (!BCrypt.checkpw(request.getPassword(), user.getPassword())) {
            throw new BusinessException("密码错误");
        }

        // 5. 使用TokenManager创建Token和构建响应
        return tokenManager.login(user);
    }

    /**
     * 用户登出
     *
     * @param token JWT Token
     */
    public void logout(String token) {
        tokenManager.logout(token);
    }

    /**
     * 验证Token是否有效（包括黑名单检查）
     *
     * @param token JWT Token
     * @return 是否有效
     */
    public boolean isTokenValid(String token) {
        return tokenManager.isTokenValid(token);
    }

}
