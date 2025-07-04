package top.yukuii.apijava.controller;

import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import lombok.RequiredArgsConstructor;
import top.yukuii.apijava.common.Result;
import top.yukuii.apijava.model.dto.LoginRequestDTO;
import top.yukuii.apijava.model.vo.LoginResponseVO;
import top.yukuii.apijava.service.AuthService;

@RestController
@RequestMapping("/auth")
@RequiredArgsConstructor
public class AuthController {

    private final AuthService authService;


    @PostMapping("/login")
    public Result<LoginResponseVO> login(@RequestBody LoginRequestDTO request) {
        return Result.success(authService.login(request));
    }
}
