package top.yukuii.apijava.mapper;

import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import top.yukuii.apijava.model.entity.Category;
import org.apache.ibatis.annotations.Mapper;

@Mapper
public interface CategoryMapper extends BaseMapper<Category> {
}