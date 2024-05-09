# 流程
先说一下流程 后面的api也会按照这个流程来
 - 先获取 一个csrf_cookie_name 和 vpapp_session 
 - 设定_maq 模拟真人 ( 可有可无)
 - 登陆
 - 刷时长
     - 先去网页端找到一个题做一遍 记录几个ajax 的请求
     - 这里注意一下 下面几个链接的请求体 可以保存下来
     - https://mrr.readoor.cn/api/qb/producer
     - https://stat.readoor.cn/index.php/materialData/insertMaterialRecord
     - https://stat.readoor.cn/index.php/materialData/getMaterialRecord
     - 最后提交一下就行 就能刷时长了
我们的最终目的是提交上我们的时长 提交前需要一些数据 要通过请求才能获得

# API
### 先设置那个随机的cookie
GET https://app.readoor.cn/app/cs/addBookShelf/1544059443
Args:

| KEY           | Description                 | Example                |
| ------------- | --------------------------- | ---------------------- |
| csrf_app_name | 与 csrf-cookie-name 值相同      | 0adff426ccea9e3173b... |
| course        | 这个是这本书的 book_id base64 后的结果 | MzE0NDY=               |
Response:
会在响应头添加 2 个 Set-Cookies 头
第一个是删除已有的 第二个是添加新的

### 获取做题页面 获取idf begTime
GET https://app.readoor.cn/app/qb/spd/1544059443?lib=MzA2Mjc=&group=MjE4&s=1&dis=1&sr=1&cl=MzE0NDZfNjg3NTQ=&ci=12233&pid=122622&tid=10418
Args:

| KEY   | Description                      | Example          |
| ----- | -------------------------------- | ---------------- |
| lib   |                                  | MzA2Mjc=         |
| group | 书的 group_id                      | MjE4             |
| s     |                                  | 1                |
| dis   |                                  | 1                |
| sr    |                                  | 1                |
| cl    | 解析出来 是一个下划线分开的 前面的数字是book_id 后面？ | MzE0NDZfNjg3NTQ= |
| ci    | company_id                       | 12233            |
| pid   | 可以研究一下获取书的api 这个是 data-pid 标签的   | 122622           |
| tid   | 同上 data-training_id              | 10418            |
Response:
返回一个html页面 我们要从这个页面提取信息
主要是给后面提交用
提交需要几个 参数 
- uidffqbsub
- referrer
- url
- begTime
这些参数都需要从页面中提取 它们是随时刷新的且加密的 下一次请求就会有一个新的值
同时下一次请求过后 上一次的值就无法使用了
我们可以用正则表达式把这些提取出来

url是你请求页面的url
其他都可以在 返回的html的 一个 script 标签里面找到
begTime 可以通过搜索 Mooc 来找到 其他的也在附近


### 提交前的准备
POST https://stat.readoor.cn/index.php/materialData/insertMaterialRecord
Body:

| KEY                  | Description                                | Example                                                                                                                                                                                                                                                                                                    |
| -------------------- | ------------------------------------------ | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| csrf_app_name        | 与 csrf-cookie-name 值相同                     | 0adff426ccea9e3173b...                                                                                                                                                                                                                                                                                     |
| previousMaterialList | 还记得之前说记下来请求记录吗 这个就是要用的 它是一个json 发送的是这个书的内容 | [  <br>    {  <br>        "ui": "",  <br>        "bi": "",  <br>        "bci": "",  <br>        "bt": "",  <br>        "bst": "",  <br>        "cui": ,  <br>        "cpi": ,  <br>        "mid": "",  <br>        "mi": ,  <br>        "tm":   <br>    }  <br>]                                           |
| materialRecordList   | 和上一个差不多 只多了几个键 想要知道是什么可以去翻翻网页端的代码 这些键都是缩写  | [  <br>    {  <br>        "ui": "",  <br>        "bi": "",  <br>        "bci": "",  <br>        "bt": "",  <br>        "bst": "",  <br>        "lid": "",  <br>        "li": ,  <br>        "lt": "",  <br>        "mp": ,  <br>        "msp": ,  <br>        "md": ,  <br>        "fs":  <br>    }  <br>] |
Response:
没啥返回的

### 提交做题记录
POST https://mrr.readoor.cn/api/qb/producer
Body:

| KEY           | Description                                                                                                                  | Example                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
| ------------- | ---------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| csrf_app_name | 与 csrf-cookie-name 值相同                                                                                                       | 0adff426ccea9e3173b...                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| data          | 包括了做题的信息 idf 做题时间 做题独一的time-based的id 这里讲几个重要的键 1-time_consuming 这个是做题时间 单位 秒 2-unique_idf 这个是你当时从第一个页面获取到的 unique_idf 要复制到这来 | {<br>    "result": {},<br>    "base_data": {<br>        "app_id": "",<br>        "course_id": "",<br>        "lesson_id": "",<br>        "source": "",<br>        "user_id": "",<br>        "time_consuming": ,<br>        "answer_status": ,<br>        "theory_total_num": ,<br>        "group_id": "",<br>        "qb_id": "",<br>        "is_elective": "",<br>        "banji_class_id": "",<br>        "project_id": "",<br>        "train_id": "",<br>        "unique_idf": "",<br>        "submit_unique_num": <br>    }<br>} |
| cid           | 不是company_id                                                                                                                 | XAA                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| begTime       | 开始时间 已加密 和前面的MoocBegTime 相同                                                                                                  | VQ4FU1ZWDQWGBg                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| url           | 你开始获取信息页面的url                                                                                                                | https://app.readoor.cn/app/qb/spd/1544059443....                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| referrer      | 前面获取的                                                                                                                        | https://app.readoor.cn/app/qb/spd/1544059443....                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| system        |                                                                                                                              | 3                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| uidffqbsub    | 独立id                                                                                                                         | Bg0BAFxRWlQBAFN.....                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
Response:
这里没保存下来 大概就是会返回一个实际消耗的时间 和状态 还有时间这些
如果你按标准流程来了 应该是可以记录成功的

# 注意事项
你在提交的时候 它会去后台查记录 你请求里面的time_consuming 比上一次请求的时间还要长 他不会记录进去 所以实际还是需要那么多时间

