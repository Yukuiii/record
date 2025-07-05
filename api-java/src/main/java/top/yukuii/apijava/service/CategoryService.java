package top.yukuii.apijava.service;

import java.util.List;

import org.springframework.stereotype.Service;

import com.baomidou.mybatisplus.core.conditions.query.LambdaQueryWrapper;

import cn.hutool.core.bean.BeanUtil;
import lombok.RequiredArgsConstructor;
import top.yukuii.apijava.mapper.CategoryMapper;
import top.yukuii.apijava.model.entity.Category;
import top.yukuii.apijava.model.vo.CategoryTreeVO;

@Service
@RequiredArgsConstructor
public class CategoryService {

    private final CategoryMapper categoryMapper;

    /**
     * 获取用户可见的分类列表（系统分类 + 用户自定义分类）
     */
    public List<Category> getUserCategories(Long userId, String type) {
        LambdaQueryWrapper<Category> wrapper = new LambdaQueryWrapper<>();
        
        // 查询系统分类或用户自定义分类
        wrapper.and(w -> w.isNull(Category::getUserId).or().eq(Category::getUserId, userId));
        
        // 按类型过滤
        if (type != null) {
            wrapper.eq(Category::getType, type);
        }
        
        // 只查询启用的分类
        wrapper.eq(Category::getStatus, 1);
        
        // 按排序序号和创建时间排序
        wrapper.orderByAsc(Category::getSortOrder, Category::getCreateTime);
        
        return categoryMapper.selectList(wrapper);
    }

    /**
     * 获取分类树结构
     */
    public List<CategoryTreeVO> getCategoryTree(Long userId, String type) {
        List<Category> allCategories = getUserCategories(userId, type);
        
        // 先获取一级分类并转换为VO
        List<CategoryTreeVO> topCategories = allCategories.stream()
            .filter(category -> category.getParentId() == null)
            .map(category -> {
                CategoryTreeVO vo = new CategoryTreeVO();
                BeanUtil.copyProperties(category, vo);
                return vo;
            })
            .toList();
        
        // 为每个一级分类设置子分类
        topCategories.forEach(topCategory -> {
            List<CategoryTreeVO> children = allCategories.stream()
                .filter(category -> topCategory.getId().equals(category.getParentId()))
                .map(category -> {
                    CategoryTreeVO childVO = new CategoryTreeVO();
                    BeanUtil.copyProperties(category, childVO);
                    return childVO;
                })
                .toList();
            topCategory.setChildren(children);
        });
        
        return topCategories;
    }

    /**
     * 创建分类
     */
    public Category createCategory(Category category) {
        // 设置层级
        if (category.getParentId() == null) {
            category.setLevel(1);
        } else {
            category.setLevel(2);
        }
        
        // 设置默认状态
        if (category.getStatus() == null) {
            category.setStatus(1);
        }
        
        // 设置为用户自定义分类
        if (category.getIsSystem() == null) {
            category.setIsSystem(0);
        }
        
        categoryMapper.insert(category);
        return category;
    }

    /**
     * 更新分类
     */
    public Category updateCategory(Category category) {
        categoryMapper.updateById(category);
        return category;
    }

    /**
     * 删除分类
     */
    public void deleteCategory(Long id) {
        // 检查是否有子分类
        LambdaQueryWrapper<Category> wrapper = new LambdaQueryWrapper<>();
        wrapper.eq(Category::getParentId, id);
        Long childCount = categoryMapper.selectCount(wrapper);
        
        if (childCount > 0) {
            throw new RuntimeException("该分类下存在子分类，无法删除");
        }
        
        // TODO: 检查是否有关联的交易记录
        
        categoryMapper.deleteById(id);
    }

    /**
     * 根据ID获取分类
     */
    public Category getCategoryById(Long id) {
        return categoryMapper.selectById(id);
    }
}