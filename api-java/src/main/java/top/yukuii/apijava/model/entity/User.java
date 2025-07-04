package top.yukuii.apijava.model.entity;

import com.baomidou.mybatisplus.annotation.IdType;
import com.baomidou.mybatisplus.annotation.TableId;
import com.baomidou.mybatisplus.annotation.TableName;

import lombok.Builder;
import lombok.Data;
import lombok.EqualsAndHashCode;

@Data
@EqualsAndHashCode(callSuper=false)
@Builder
@TableName("user")
public class User extends BaseEntity {

    @TableId(type = IdType.AUTO)
    private Long id;

    @TableId(type = IdType.ASSIGN_UUID)
    private String userId;

    private String userName;

    private String password;

    private String email;

    private Integer status;

    private String phone;
    
    private String avatar;

    private String remark;

}
