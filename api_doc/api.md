## 域名
https://www.onb.io/

## 接口汇总

### 1.注册接口

#### 接口地址:
/api/user/register
#### 请求方式：
POST
#### 请求参数：
|  参数名   | 类型  | 是否必须   | 说明 |
|  ----  | ----  | ----  | ----  |
| user_name  | string | Y  | 用户名 |
| email  | string |Y  | 邮箱 |
| user_pwd  | string |Y  | 密码 |
#### 返回
```
{
    "ret":0,
    "msg":"succ",
    "data":""
}
```

### 2.发送注册邮件接口

#### 接口地址:
/api/user/send_validate_email
#### 请求方式：
POST
#### 请求参数：
|  参数名   | 类型  | 是否必须   | 说明 |
|  ----  | ----  | ----  | ----  |
|  user_email  | string  | Y  | 用户邮箱  |

#### 返回
```json
{
    "ret":0,
    "msg":"succ",
    "data":""
}
```

### 3.验证邮箱接口

#### 接口地址:
/api/user/validate_email
#### 请求方式：
GET
#### 请求参数：
|  参数名   | 类型  | 是否必须   | 说明 |
|  ----  | ----  | ----  | ----  |
|  code  | string  | 是  | 验证code  |

#### 返回
```json
{
    "ret":0,
    "msg":"succ",
    "data":""
}
```

### 4.登录接口

#### 接口地址:
/api/user/login
#### 请求方式：
Form 表单提交
#### 请求参数：
|  参数名   | 类型  | 是否必须   | 说明 |
|  ----  | ----  | ----  | ----  |
|  user_name  | string  | 是  | 用户名  |
|  user_pwd  | string  | 是  | 密码  |
|  is_remembered  | int  | 否  | 是否要记住  |

#### 返回
```json
跳转个人主页
```

### 5.发送密码重置邮件

#### 接口地址:
/api/user/send_reset_pwd
#### 请求方式：
POST
#### 请求参数：
|  参数名   | 类型  | 是否必须   | 说明 |
|  ----  | ----  | ----  | ----  |
|  user_email  | string  | 是  | 用户邮箱  |

#### 返回
```json
{
    "ret":0,
    "msg":"succ",
    "data":""
}
```

### 6.密码重置请求

#### 接口地址:
/api/user/reset_pwd
#### 请求方式：
POST
#### 请求参数：
|  参数名   | 类型  | 是否必须   | 说明 |
|  ----  | ----  | ----  | ----  |
|  code  | string  | 是  | 上一个请求会在链接里埋一个code，在这里带过来，验证是否来自正常的重置流程  |
|  new_pwd  | string  | 是  | 新密码  |

#### 返回
```json
{
    "ret":0,
    "msg":"succ",
    "data":""
}
```

### 7.个人主页

#### 接口地址:
/api/user/userinfo
#### 请求方式：
GET
#### 请求参数：
|  参数名   | 类型  | 是否必须   | 说明 |
|  ----  | ----  | ----  | ----  |
|  user_name  | string  | 是  | 用户名  |

#### 返回
```json
{
    "ret":0,
    "msg":"succ",
    "data":{
        "user_name":"kris",
        "user_avatar":"用户头像",
        "user_link":"用户主页",
        "is_confirmed":false,
        "user_extra":"其他附加信息"
    }
}
```

### 8.更新个人头像

#### 接口地址:
/api/link/updateinfo
#### 请求方式：
GET
#### 请求参数：
|  参数名   | 类型  | 是否必须   | 说明 |
|  ----  | ----  | ----  | ----  |
|  user_img  | string  | 是  | 新头像链接  |

#### 返回
```json
{
    "ret":0,
    "msg":"succ",
    "data":""
}
```


### 9.个人链接列表

#### 接口地址:
/api/link/userlink
#### 请求方式：
GET
#### 请求参数：
|  参数名   | 类型  | 是否必须   | 说明 |
|  ----  | ----  | ----  | ----  |

#### 返回
```json
{
    "ret":0,    
    "msg":"succ",
    "data":[
        {
            "count":1,
            "page":1,
            "page_size":10,
            "list":[
                {
                    "link_id":111,
                    "link_url":"http://www.qq.com",
                    "link_desc":"etss",
                    "link_img":"首图链接",
                    "is_valid":0,
                    "is_special":1,
                    "position":12
                }
            ]

        }
    ]
}
```

### 10.创建个人链接

#### 接口地址:
/api/link/createlink
#### 请求方式：
POST
#### 请求参数：
|  参数名   | 类型  | 是否必须   | 说明 |
|  ----  | ----  | ----  | ----  |
|  link_url  | string  | 是  | url  |
|  position  | int  | 是  | position  |
#### 返回
```json
{
    "ret":0,    
    "msg":"succ",
    "data":
    {
        "link_id":111,
        "link_url":"http://www.qq.com",
        "link_desc":"etss",
        "link_img":"首图链接",
        "is_valid":0,
        "is_special":1,
        "position":12
    }
}
```

### 10.更新个人链接

#### 接口地址:
/api/link/updatelink
#### 请求方式：
POST
#### 请求参数：
|  参数名   | 类型  | 是否必须   | 说明 |
|  ----  | ----  | ----  | ----  |
|  id  | int  | 是  | id  |
|  link_url  | string  | 是  | url  |
|  link_desc  | string  | 是  | desc  |
|  link_img  | string  | 是  | img  |
|  is_valid  | int  | 是  | 是否有效  |
|  is_special  | int  | 是  | 是否有特效  |
|  position  | int  | 是  | 位置  |

#### 返回
```json
{
    "ret":0,    
    "msg":"succ",
    "data":{}
}
```
### 11.个人链接列表（不需要登录态，只获取个人编辑的有效链接）

#### 接口地址:
/api/user/userlinklist
#### 请求方式：
GET
#### 请求参数：
|  参数名   | 类型  | 是否必须   | 说明 |
|  ----  | ----  | ----  | ----  |
|  user_name  | string  | 是  | 用户名  |
|  page  | int  | 是  | 页码  |
|  page_size  | int  | 是  | 页大小  |
#### 返回
```json
{
    "ret":0,    
    "msg":"succ",
    "data":[
        {
            "count":1,
            "page":1,
            "page_size":10,
            "list":[
                {
                    "link_id":111,
                    "link_url":"http://www.qq.com",
                    "link_desc":"etss",
                    "link_img":"首图链接",
                    "is_special":1,
                    "position":12
                }
            ]

        }
    ]
}
```

### 12.删除个人链接

#### 接口地址:
/api/link/deletelink
#### 请求方式：
POST
#### 请求参数：
|  参数名   | 类型  | 是否必须   | 说明 |
|  ----  | ----  | ----  | ----  |
|  id  | int  | 是  | id  |

#### 返回
```json
{
    "ret":0,    
    "msg":"succ",
    "data":{}
}
```


### 13.上传图片

#### 接口地址:
/api/link/upload
#### 请求方式：
POST multipart/form-data 表单
#### 请求参数：
|  参数名   | 类型  | 是否必须   | 说明 |
|  ----  | ----  | ----  | ----  |
|  upload  | file  | 是  | file  |

#### 返回
```json
{
    "ret":0,    
    "msg":"succ",
    "url":"http://www.qq.com"
}
```

### 14.登出

#### 接口地址:
/api/link/logout
#### 请求方式：
GET 
#### 请求参数：
|  参数名   | 类型  | 是否必须   | 说明 |
|  ----  | ----  | ----  | ----  |


#### 返回
跳转首页

