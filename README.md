## 1
随便打开一个学习页面
https://app.readoor.cn/app/qb/spd/1544059443?lib=MzA2Mjc=&group=MjE0&s=1&dis=1&sr=1&cl=MzE0NDZfNjg3NTE=&ci=12233&pid=122622&tid=10418

发现可以正常进入
cookie: normal_1.json

```text
其中有一个键值对是 171db5f7d616bdf56f1d4e0fa546468b=a%3A1%3A%7Bi%3A0%3Bs%3A5%3A%2231446%22%3B%7D
这里url解码出来是 a:1:{i:0;s:5:"31446";}

问了问gpt
这段看起来像是 PHP 的序列化数据。
根据这个推测，它可能表示一个包含一个数组的关联数组，其中键为0，值是一个包含一个字符串的数组，字符串长度为5，内容为"31446"。
```

再开一个private 页面进去

这里发现只有一个随机生成的key没了 然后_maq 里面的内容 _uid 为空 这个时候就需要重新登陆了

这里试试不带上随机生成的key 只带上_maq 里面的 _uid
用postman测试还是同样的页面 说明需要那个随机的cookie

## 2
带上随机的cookie进入
内容还是要你登陆 说明那个session和 那个随机的cookie可能是绑定在一起的
