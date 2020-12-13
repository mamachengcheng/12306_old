# 铁路12306数据模型文档

## 公共

| 编号 | 字段       | 数据类型 | 解释 | 字段标签 | 备注 |
| ---- | ---------- | -------- | ---- | -------- | ---- |
| 1    | id         |          |      |          |      |
| 2    | created_at |          |      |          |      |
| 3    | updated_at |          |      |          |      |
| 4    | deleted_at |          |      |          |      |



## user 用户

| 编号 | 列名                 | 数据类型 | 解释               | 字段标签                              | 备注                                                         |
| ---- | -------------------- | -------- | ------------------ | ------------------------------------- | ------------------------------------------------------------ |
|      | uid                  |          | 用户ID             | primaryKey;unique;default:;not null;- |                                                              |
|      | user_name            | string   | 用户名             | not null;check:;                      | 字母数字或者下划线"-"，6-30位                                |
|      | password             | string   | 密码               | not null;check:;                      | 不少于6位                                                    |
|      | name                 | string   | 姓名               | not null;check:;                      | 真实姓名                                                     |
|      | country              | uint     | 国家               | not null;check:;                      | +86                                                          |
|      | certificate_type     | uint     | 证件类型           | default:                              | 中国居民身份证、港澳居民来往内地通行证、台湾居民来往大陆通行证和护照 |
|      | sex                  | bool     | 性别               |                                       | 男和女                                                       |
|      | birthday             | date     | 出生日期           |                                       |                                                              |
|      | certificate_deadline | date     | 证件有效期截至日期 |                                       |                                                              |
|      | certificate          | string   | 证件号码           | unique                                |                                                              |
|      | mobile_phone         | string   | 手机号码           | unique                                | 默认+86                                                      |
|      | fixed_phone          | string   | 固定电话           |                                       |                                                              |
|      | mail                 | string   | 电子邮箱           | unique                                |                                                              |
|      | address              | string   | 地址               |                                       | 省-市-区-街道-楼牌号（）                                     |
|      | postcode             | uint     | 邮编               |                                       |                                                              |
|      | passenger_type       | uint     | 旅客类型           |                                       | 成人、儿童、学生和参军                                       |
|      | user_type            | uint     | 用户类型           |                                       |                                                              |
|      | check_status         | uint     | 审核状态           |                                       | 审核中、审核通过和审核失败                                   |
|      | user_status          | uint     | 用户状态           |                                       | 可用、禁用                                                   |



## passenger 乘客

| 编号 | 字段                 | 数据类型 | 解释               | 备注                       |
| ---- | -------------------- | -------- | ------------------ | -------------------------- |
|      | name                 | string   | 姓名               | 真实姓名                   |
|      | certificate_type     | uint     | 证件类型           |                            |
|      | sex                  | bool     | 性别               | 男和女                     |
|      | birthday             | data     | 出生日期           |                            |
|      | certificate_deadline | data     | 证件有效期截至日期 |                            |
|      | certificate          | string   | 证件号码           |                            |
|      | certificate_type     | uint     | 旅客类型           |                            |
|      | mobile_phone         | string   | 手机号码           | 默认+86                    |
|      | mail                 | string   | 电子邮箱           |                            |
|      | check_status         | uint     | 审核状态           | 审核中、审核通过和审核失败 |
|      | user_status          | uint     | 用户状态           | 可用、禁用                 |

