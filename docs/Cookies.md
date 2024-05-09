主要有四个cookie

## \_maq
这个没啥用 是用来统计的(代码注释有写)
```js 
// 来自 https://app.readoor.cn/app/dt/pd/1544059443?s=1
if(!empty(stats_url)){
   //统计用
   var _maq = _maq || [];
    _maq.push(['_uid', _userId]);
   _maq.push(['_appId', app_id]);
   _maq.push(['_cid', _companyId]);
   _maq.push(['_isMobile', is_mobile]);  //1:is_mobile;0:PC
   _maq.push(['_isWeixin', is_weixin_]);  //3:小程序；2：微信；0：其他
   _maq.push(['_accessType', ""]); //微库id
   _maq.push(['_keyValue', ""]); //内容
   //add by tianjun 20180718 start
  _maq.push(['_appver', '2.29.1']);
  _maq.push(['_apiver', '1.0']);
  //add by tianjun 20180718 end
            document.cookie="_maq="+ _maq.toString() +"; path=/";
    
```

## csrf_cookie_name

这个是服务端返回的cookie 如果你没有就会返回给你新的 
如果已经有了 下次服务端返回的时候就会续签(延长cookie时间)
这个挺重要的 会在很多地方用上 例如 
get请求的参数
post表单的数据

## vpapp_session

里面有 session一词就大概能猜出来这个是 sessionID
关于什么是 SessionID 自行wiki

## 171db5f7d616b........
这么长一串 我也不知道是啥用

它的 value 是 a%3A1%3A%7Bi%3A0%3Bs%3A5%3A%2231446%22%3B%7D

是个人都看得出来这个是url编码的字符串

解析出来 
```text
a:1:{i:0;s:5:"31446";}
```
问了问gpt
```text
这段看起来像是 PHP 的序列化数据。  
根据这个推测，它可能表示一个包含一个数组的关联数组，其中键为0，值是一个包含一个字符串的数组，字符串长度为5，内容为"31446"。
```
具体啥用 是啥 咱也不知道
