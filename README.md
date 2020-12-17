# 铁路12306
<div align=center>
    <img src="https://github.com/mamachengcheng/12306/blob/main/docs/logo.png" width=300" height="300" />
</div>

铁路12306是基于gin开发的售票系统，实现了用户信息注册、余票查询、售票、退票和改签等功能。

## 快速部署
### 安装docker

[安装](https://docs.docker.com/engine/install/)

### 安装docker-compose

[安装](https://docs.docker.com/compose/install/)

### 部署项目

```shell
git clone https://github.com/mamachengcheng/12306.git && docker-compose up
```

## 项目文档

### 需求文档

[需求文档](https://github.com/mamachengcheng/12306/blob/main/docs/PRD/PRD.md)

### 数据模型文档

[数据模型文档](https://github.com/mamachengcheng/12306/blob/main/docs/model/model.md)

### 公共部分

| 编号 | 字段       | 数据类型  | 解释     | 备注 | 字段标签 |
| ---- | ---------- | --------- | -------- | ---- | -------- |
| 1    | id         | uint      | ID       |      |          |
| 2    | created_at | time.Time | 创建时间 |      |          |
| 3    | updated_at | time.Time | 更新时间 |      |          |
| 4    | deleted_at | time.Time | 删除时间 |      |          |



### 用户部分

#### User 用户

| 编号 | 列名                 | 数据类型 | 解释               | 备注                                                         | 字段标签                              |
| ---- | -------------------- | -------- | ------------------ | ------------------------------------------------------------ | ------------------------------------- |
| 1    | uid                  | string   | 用户ID             |                                                              | primaryKey;unique;default:;not null;- |
| 2    | user_name            | string   | 用户名             | 字母数字或者下划线"-"，6-30位                                | not null;check:;                      |
| 3    | password             | string   | 密码               | 不少于6位                                                    | not null;check:;                      |
| 4    | name                 | string   | 姓名               | 真实姓名                                                     | not null;check:;                      |
| 5    | country              | uint     | 国家               | +86                                                          | not null;check:;                      |
| 6    | certificate_type     | uint     | 证件类型           | 中国居民身份证、港澳居民来往内地通行证、台湾居民来往大陆通行证和护照 | default:                              |
| 7    | sex                  | bool     | 性别               | 男和女                                                       |                                       |
| 8    | birthday             | date     | 出生日期           |                                                              |                                       |
| 9    | certificate_deadline | date     | 证件有效期截至日期 |                                                              |                                       |
| 10   | certificate          | string   | 证件号码           |                                                              | unique                                |
| 11   | mobile_phone         | string   | 手机号码           | 默认+86                                                      | unique                                |
| 12   | fixed_phone          | string   | 固定电话           |                                                              |                                       |
| 13   | mail                 | string   | 电子邮箱           |                                                              | unique                                |
| 14   | address              | string   | 地址               | 省-市-区-街道-楼牌号（）                                     |                                       |
| 15   | postcode             | uint     | 邮编               |                                                              |                                       |
| 16   | passenger_type       | uint     | 旅客类型           | 成人、儿童、学生和参军                                       |                                       |
| 17   | user_type            | uint     | 用户类型           |                                                              |                                       |
| 18   | check_status         | uint     | 审核状态           | 审核中、审核通过和审核失败                                   |                                       |
| 19   | user_status          | uint     | 用户状态           | 可用、禁用                                                   |                                       |

#### Passenger 乘客

| 编号 | 字段                 | 数据类型 | 解释               | 备注                       | 字段标签 |
| ---- | -------------------- | -------- | ------------------ | -------------------------- | -------- |
| 1    | name                 | string   | 姓名               | 真实姓名                   |          |
| 2    | certificate_type     | uint     | 证件类型           |                            |          |
| 3    | sex                  | bool     | 性别               | 男和女                     |          |
| 4    | birthday             | data     | 出生日期           |                            |          |
| 5    | certificate_deadline | data     | 证件有效期截至日期 |                            |          |
| 6    | certificate          | string   | 证件号码           |                            |          |
| 7    | certificate_type     | uint     | 旅客类型           |                            |          |
| 8    | mobile_phone         | string   | 手机号码           | 默认+86                    |          |
| 9    | mail                 | string   | 电子邮箱           |                            |          |
| 10   | check_status         | uint     | 审核状态           | 审核中、审核通过和审核失败 |          |
| 11   | user_status          | uint     | 用户状态           | 可用、禁用                 |          |

### 列车部分

#### Train 列车

| 编号 | 字段       | 数据类型 | 解释     | 备注 | 字段标签 |
| ---- | ---------- | -------- | -------- | ---- | -------- |
| 1    | train_code | string   | 列车代号 |      |          |
| 2    | train_type | uint     | 列车类型 |      |          |
| 3    | seat       | Seat     | 座位     |      |          |

#### Seat 座位

| 编号 | 字段        | 数据类型 | 解释     | 备注 | 字段标签 |
| ---- | ----------- | -------- | -------- | ---- | -------- |
| 1    | seat_number | string   | 座位编号 |      |          |
| 2    | car_number  | uint     | 车厢编号 |      |          |
| 3    | train       | Train    | 列车     |      |          |
| 4    | price       | float    | 票价     |      |          |
| 5    | seat_status | uint     | 座位状态 |      |          |
| 6    | seat_type   | uint     | 座位类型 |      |          |

#### Schedule 班次

| 编号 | 字段          | 数据类型 | 解释     | 备注 | 字段标签 |
| ---- | ------------- | -------- | -------- | ---- | -------- |
| 1    | train         | Train    | 列车     |      |          |
| 2    | start_station | Station  | 终点站   |      |          |
| 3    | end_station   | Sation   | 起点站   |      |          |
| 4    | start_time    | time     | 出发时间 |      |          |
| 5    | end_time      | time     | 到达时间 |      |          |
| 6    | stop          | Stop     | 经停站   |      |          |
| 7    | duration      | uint     | 沿途时间 |      |          |

#### Station 车站

| 编号 | 字段         | 数据类型 | 解释     | 备注 | 字段标签 |
| ---- | ------------ | -------- | -------- | ---- | -------- |
| 1    | station_name | string   | 车站名   |      |          |
| 2    | city         | string   | 城市名   |      |          |
| 3    | city_pinyin  | string   | 拼音     |      |          |
| 4    | first_pinyin | string   | 首字母   |      |          |
| 5    | city_code    | string   | 城市代码 |      |          |
| 6    | station_code | string   | 车站代码 |      |          |

#### Stop 经停站

| 编号 | 字段          | 数据类型  | 解释     | 备注 | 字段标签 |
| ---- | ------------- | --------- | -------- | ---- | -------- |
| 1    | schedule      | Schedule  | 班次     |      |          |
| 2    | serial_number | uint      | 站序     |      |          |
| 3    | start_station | Station   | 起始站   |      |          |
| 4    | end_station   | Station   | 终止站   |      |          |
| 5    | start_time    | time.Time | 发车时间 |      |          |
| 6    | end_time      | time.Time | 到达时间 |      |          |
| 7    | duration      | uint      | 沿途时间 |      |          |

### 订单部分

#### Order 订单

| 编号 | 字段      | 数据类型  | 解释     | 备注 | 字段标签 |
| ---- | --------- | --------- | -------- | ---- | -------- |
| 1    | trade_no  | string    | 交易编号 |      |          |
| 2    | passenger | Passenger | 乘客     |      |          |
| 3    | user      | User      | 用户     |      |          |
| 4    | schedule  | Schedule  | 班次     |      |          |
| 5    | status    | uint      | 订单状态 |      |          |
| 6    | seat      | Seat      | 座位     |      |          |



### 接口文档

#### 内部接口

[内部接口文档](https://github.com/mamachengcheng/12306/blob/main/docs/api/api.md)

#### 外部接口

[外部接口文档](https://belugahub.postman.co/build/workspace/Team-Workspace~7003af59-00c2-4a32-8d4d-098d1af5422a/request/13390250-aaed9fc2-ee89-4de2-b56b-93acfe09943c)
