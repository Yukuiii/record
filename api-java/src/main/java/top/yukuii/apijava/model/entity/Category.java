package top.yukuii.apijava.model.entity;

import com.baomidou.mybatisplus.annotation.IdType;
import com.baomidou.mybatisplus.annotation.TableId;
import com.baomidou.mybatisplus.annotation.TableName;

import lombok.Builder;
import lombok.Data;
import lombok.EqualsAndHashCode;

/**
 * 分类实体类
 */
@Data
@EqualsAndHashCode(callSuper=false)
@Builder
@TableName("category")
public class Category extends BaseEntity {

    /**
     * 分类主键ID
     */
    @TableId(type = IdType.AUTO)
    private Long id;

    /**
     * 用户ID，系统分类为null，用户自定义分类关联用户
     */
    private Long userId;

    /**
     * 分类名称
     */
    private String name;

    /**
     * 分类编码，用于前端标识
     */
    private String code;

    /**
     * 父分类ID，为空表示一级分类
     */
    private Long parentId;

    /**
     * 分类层级：1-一级分类，2-二级分类
     */
    private Integer level;

    /**
     * 分类类型：income-收入分类，expense-支出分类
     */
    private String type;

    /**
     * 分类图标
     */
    private String icon;

    /**
     * 分类颜色（用于前端显示）
     */
    private String color;

    /**
     * 排序序号
     */
    private Integer sortOrder;

    /**
     * 是否系统内置分类：0-用户自定义，1-系统内置
     */
    private Integer isSystem;

    /**
     * 分类状态：0-禁用，1-启用
     */
    private Integer status;

    /**
     * 分类描述
     */
    private String description;

}