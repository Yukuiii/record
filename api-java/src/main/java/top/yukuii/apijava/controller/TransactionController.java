package top.yukuii.apijava.controller;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import com.baomidou.mybatisplus.extension.plugins.pagination.Page;

import lombok.RequiredArgsConstructor;
import top.yukuii.apijava.common.Result;
import top.yukuii.apijava.model.dto.CreateTransactionDTO;
import top.yukuii.apijava.model.vo.GetTransactionVO;
import top.yukuii.apijava.service.TransactionService;

@RestController
@RequestMapping("/api/transaction") 
@RequiredArgsConstructor
public class TransactionController {

    private final TransactionService transactionService;

    /**
     * 获取当前登录用户账单记录
     * @param type
     * @param categoryId
     * @param startDate
     * @param endDate
     * @param keyword
     * @param sort
     * @param order
     * @param page
     * @param pageSize
     * @return
     */
    @GetMapping("/page")
    public Result<Page<GetTransactionVO>> getTransactionsPage(@RequestParam(required = false) String type,
            @RequestParam(required = false) Long categoryId,
            @RequestParam(required = false) Long startDate,
            @RequestParam(required = false) Long endDate,
            @RequestParam(required = false) String keyword,
            @RequestParam(required = false) String sort,
            @RequestParam(required = false) String order,
            @RequestParam(required = false) Integer page,
            @RequestParam(required = false) Integer pageSize) {
        return Result.success(transactionService.getTransactionsPage(type, categoryId, startDate, endDate, keyword, sort, order, page, pageSize));
    }

    /**
     * 创建账单记录
     * @param createTransactionDTO
     * @return
     */
    @PostMapping("/create")
    public Result<Void> createTransaction(@RequestBody CreateTransactionDTO createTransactionDTO) {
        transactionService.createTransaction(createTransactionDTO);
        return Result.success();
    }
    
}
