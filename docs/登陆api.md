
# Api
### 主登陆

POST https://app.readoor.cn/app/cm/login
Content-Type: application/x-www-form-urlencoded
Args: ( 是的没错 ，post请求也有url的query 参数 用法和GET 请求一样)

| KEY  | Description | Example    |
| ---- | ----------- | ---------- |
| s_id | appid       | 1544059443 |

Body:

| KEY           | Description                                            | Example                |
| ------------- | ------------------------------------------------------ | ---------------------- |
| account       | 这个是加密后的数据 原数据可以是手机号或者邮箱 这里使用的RSA 加密 详细看 一些关于i学网站的概念.md | gaGpivTFpnEx88...      |
| password      | 这个也是加密后的数据 输入的是密码                                      | V5F2X6+CBEduD...       |
| verify        | 这个也是加密后的数据 输入的是验证码                                     | oaqrZ6p3jeBqdd....     |
| csrf_app_name | 与 csrf-cookie-name 值相同                                 | 0adff426ccea9e3173b... |
Response :
- 成功登陆
返回一个json字符串 包括了这个账号的信息 具体每个键什么意识 可以自己去请求一下 一下就能看懂
```json
{
    "status": 1,
    "user": {
        "user_id": "",
        "user_identifier": "",
        "user_account": "",
        "user_type": "",
        "user_mail": "",
        "user_pn": "",
        "user_nickname": "",
        "user_sex": "",
        "user_birthday": "",
        "user_province_id": "",
        "user_register_time": "",
        "user_logo_url": "",
        "user_login_token": "",
        "user_student_number": "",
        "user_guid": "",
        "balance": "",
        "wx_user_type": "",
        "wx_user_pn": "",
        "user_real_name": "",
        "user_info": [

        ],
        "user_grade": null,
        "vip_type": null,
        "vip_start_time": null,
        "vip_end_time": null,
        "isCanvasser": false,
        "forceLogin": 0,
        "role": [

        ]
    },
    "errorMessage": [

    ],
    "time": 0
}
```

- 登陆失败
也是返回json 只不过结构不一样了
```json
{  
    "distributor_id": "",  
    "promotion_id": "",  
    "user": {  
        "user_identifier": "",  
        "user_id": ""  
    },  
    "status": -1,  
    "errorMessage": {  
        "account": 1111 
    },  
    "time": 0 
}
```

其中 errorMessage.account 是错误代码 这里只总结两个
1111 验证码错误
1110 账密错误

### 获取验证码
GET https://app.readoor.cn/app/napp/getImageCode/login/4/80/40
Args:

| KEY  | Description     | Example            |
| ---- | --------------- | ------------------ |
| s_id | appid           | 1544059443         |
| rand | 一个低于1的浮点数 共16 位 | 0.6082343041477555 |
Response:
- 成功
会返回 一个图形验证码 注意 : 随机数的参数和验证码无关 但是最好还是随机一下
- 失败
.. 还没遇到过