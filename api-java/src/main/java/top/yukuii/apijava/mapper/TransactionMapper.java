package top.yukuii.apijava.mapper;

import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import top.yukuii.apijava.model.entity.Transaction;
import org.apache.ibatis.annotations.Mapper;

@Mapper
public interface TransactionMapper extends BaseMapper<Transaction> {
}