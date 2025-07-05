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
import top.yukuii.apijava.model.dto.RegisterRequestDTO;
import top.yukuii.apijava.model.entity.User;
import top.yukuii.apijava.model.vo.LoginResponseVO;
import top.yukuii.apijava.util.TokenManager;

@Slf4j
@Service
@RequiredArgsConstructor
public class AuthService {

    private final UserMapper userMapper;

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
        return TokenManager.login(user);
    }

    /**
     * 用户注册
     *
     * @param request 注册请求
     * @return 注册响应
     */
    public LoginResponseVO register(RegisterRequestDTO request) {
        // 1. 参数验证
        validateRegisterRequest(request);

        // 2. 检查用户名是否已存在
        if (isUserNameExists(request.getUserName())) {
            throw new BusinessException("用户名已存在");
        }

        // 3. 检查邮箱是否已存在
        if (ObjectUtil.isNotEmpty(request.getEmail()) && isEmailExists(request.getEmail())) {
            throw new BusinessException("邮箱已被注册");
        }

        // 4. 检查手机号是否已存在
        if (ObjectUtil.isNotEmpty(request.getPhone()) && isPhoneExists(request.getPhone())) {
            throw new BusinessException("手机号已被注册");
        }

        // 5. 创建用户
        User user = createUser(request);

        // 6. 保存用户
        userMapper.insert(user);

        // 7. 自动登录并返回Token
        return TokenManager.login(user);
    }

    /**
     * 验证注册请求参数
     */
    private void validateRegisterRequest(RegisterRequestDTO request) {
        if (ObjectUtil.isEmpty(request.getUserName()) || ObjectUtil.isEmpty(request.getPassword())) {
            throw new BusinessException("用户名和密码不能为空");
        }

        if (!request.getPassword().equals(request.getConfirmPassword())) {
            throw new BusinessException("两次输入的密码不一致");
        }

        // 至少需要邮箱或手机号其中一个
        if (ObjectUtil.isEmpty(request.getEmail()) && ObjectUtil.isEmpty(request.getPhone())) {
            throw new BusinessException("邮箱和手机号至少填写一个");
        }
    }

    /**
     * 检查用户名是否已存在
     */
    private boolean isUserNameExists(String userName) {
        User existUser = userMapper.selectOne(
            new LambdaQueryWrapper<User>().eq(User::getUserName, userName)
        );
        return ObjectUtil.isNotEmpty(existUser);
    }

    /**
     * 检查邮箱是否已存在
     */
    private boolean isEmailExists(String email) {
        User existUser = userMapper.selectOne(
            new LambdaQueryWrapper<User>().eq(User::getEmail, email)
        );
        return ObjectUtil.isNotEmpty(existUser);
    }

    /**
     * 检查手机号是否已存在
     */
    private boolean isPhoneExists(String phone) {
        User existUser = userMapper.selectOne(
            new LambdaQueryWrapper<User>().eq(User::getPhone, phone)
        );
        return ObjectUtil.isNotEmpty(existUser);
    }

    /**
     * 创建用户对象
     */
    private User createUser(RegisterRequestDTO request) {
        // 加密密码
        String hashedPassword = BCrypt.hashpw(request.getPassword(), BCrypt.gensalt());

        long currentTime = System.currentTimeMillis();

        User user = User.builder()
                .userName(request.getUserName())
                .password(hashedPassword)
                .email(request.getEmail())
                .phone(request.getPhone())
                .avatar(request.getAvatar())
                .remark(request.getRemark())
                .status(1) // 默认状态为正常
                .build();

        // 设置审计字段
        user.setCreateTime(currentTime);
        user.setUpdateTime(currentTime);

        return user;
    }

    /**
     * 用户登出
     *
     * @param token JWT Token
     */
    public void logout(String token) {
        TokenManager.logout(token);
    }

    /**
     * 验证Token是否有效（包括黑名单检查）
     *
     * @param token JWT Token
     * @return 是否有效
     */
    public boolean isTokenValid(String token) {
        return TokenManager.isTokenValid(token);
    }

    /**
     * 检查用户名是否可用
     *
     * @param userName 用户名
     * @return true表示可用，false表示已被占用
     */
    public boolean isUserNameAvailable(String userName) {
        return !isUserNameExists(userName);
    }

    /**
     * 检查邮箱是否可用
     *
     * @param email 邮箱
     * @return true表示可用，false表示已被占用
     */
    public boolean isEmailAvailable(String email) {
        return !isEmailExists(email);
    }

    /**
     * 检查手机号是否可用
     *
     * @param phone 手机号
     * @return true表示可用，false表示已被占用
     */
    public boolean isPhoneAvailable(String phone) {
        return !isPhoneExists(phone);
    }

}
