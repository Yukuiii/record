package top.yukuii.apijava.service;

import org.springframework.stereotype.Service;

import com.baomidou.mybatisplus.core.conditions.query.LambdaQueryWrapper;
import com.baomidou.mybatisplus.extension.plugins.pagination.Page;

import cn.hutool.core.bean.BeanUtil;
import cn.hutool.core.util.ObjectUtil;
import cn.hutool.core.util.StrUtil;
import lombok.RequiredArgsConstructor;
import top.yukuii.apijava.common.BusinessException;
import top.yukuii.apijava.mapper.TransactionMapper;
import top.yukuii.apijava.model.dto.CreateTransactionDTO;
import top.yukuii.apijava.model.entity.Transaction;
import top.yukuii.apijava.model.vo.GetTransactionVO;
import top.yukuii.apijava.util.TokenManager;

@Service
@RequiredArgsConstructor
public class TransactionService {

    private final TransactionMapper transactionMapper;

    public Page<GetTransactionVO> getTransactionsPage(String type, Long categoryId, Long startDate, Long endDate, 
                                                      String keyword, String sort, String order, Integer page, Integer pageSize) {
        
        // 设置默认分页参数
        int currentPage = page != null && page > 0 ? page : 1;
        int size = pageSize != null && pageSize > 0 ? pageSize : 10;
        
        Page<Transaction> pageParam = new Page<>(currentPage, size);
        
        // 构建查询条件
        LambdaQueryWrapper<Transaction> wrapper = new LambdaQueryWrapper<>();
        
        // 类型过滤
        if (StrUtil.isNotBlank(type)) {
            wrapper.eq(Transaction::getType, type);
        }
        
        // 分类过滤
        if (categoryId != null) {
            wrapper.eq(Transaction::getCategoryId, categoryId);
        }
        
        // 日期范围过滤
        if (startDate != null) {
            wrapper.ge(Transaction::getTransactionDate, startDate);
        }
        if (endDate != null) {
            wrapper.le(Transaction::getTransactionDate, endDate);
        }
        
        // 关键词搜索
        if (StrUtil.isNotBlank(keyword)) {
            wrapper.like(Transaction::getDescription, keyword);
        }
        
        // 排序
        if (StrUtil.isNotBlank(sort)) {
            boolean isAsc = !"desc".equalsIgnoreCase(order);
            switch (sort) {
                case "amount":
                    wrapper.orderBy(true, isAsc, Transaction::getAmount);
                    break;
                case "date":
                    wrapper.orderBy(true, isAsc, Transaction::getTransactionDate);
                    break;
                case "createTime":
                    wrapper.orderBy(true, isAsc, Transaction::getCreateTime);
                    break;
                default:
                    wrapper.orderByDesc(Transaction::getCreateTime);
                    break;
            }
        } else {
            wrapper.orderByDesc(Transaction::getCreateTime);
        }
        
        // 执行分页查询
        Page<Transaction> result = transactionMapper.selectPage(pageParam, wrapper);
        
        // 转换为VO
        Page<GetTransactionVO> voPage = new Page<>();
        BeanUtil.copyProperties(result, voPage);
        
        // 转换记录
        voPage.setRecords(result.getRecords().stream()
            .map(transaction -> {
                GetTransactionVO vo = new GetTransactionVO();
                BeanUtil.copyProperties(transaction, vo);
                return vo;
            })
            .toList());
        
        return voPage;
    }

    /**
     * 创建账单记录
     * @param createTransactionDTO
     */
    public void createTransaction(CreateTransactionDTO createTransactionDTO){        
        if(ObjectUtil.isNull(createTransactionDTO)){
            throw new BusinessException("交易数据不能为空");
        }

        Transaction transaction = new Transaction();
        BeanUtil.copyProperties(createTransactionDTO, transaction);
        transaction.setUserId(Long.parseLong(TokenManager.getCurrentUserId()));
        transaction.setTransactionDate(System.currentTimeMillis());
        transaction.setTags(createTransactionDTO.getTags());
        transaction.setCreateBy(Long.parseLong(TokenManager.getCurrentUserId()));
        transaction.setCreateTime(System.currentTimeMillis());
        transactionMapper.insert(transaction);
    }
}