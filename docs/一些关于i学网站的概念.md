
# 是session 还是 JWT

直接告诉你 是 session+cookie 的方式 不是jwt

但是它相较于传统的
通常只有一个 sessionID 这种
它多了 一个新的cookie 叫做 csrf-cookie-name
看名字就知道它和csrf有关 这个也是要用的 也要拿去做身份认证

也很重要 会被带到很多请求里面去 作为get请求参数和post表单数据之类的

# i学的加密相关的

i学的加密 可以说很多(但是大部分感觉都是用来对付产品经理要求的加密之类的)

我现在看见大部分都是base64 加密的方法
有些地方会用到双次base64 （但是双层加密 用的base64 那不就跟没加密一样吗?）
~~这里吐槽一下 翻倒一个代码注释估计直接cv的 从解密复制过去没改注释 加密部分注释也是 解密~~
还有些地方是他们自己内部的加密算法之类的 还有一些随机算法

登陆会用到 1024bit PCCS#1 的RSA 算法 (听起来很严谨是吧 但是前端无秘密)
这个RSA算法会提供给你一个公钥 加密过后送给服务端(有了https还要这个干什么？)
这个公钥是固定的 你去主页面的html就能找到 这里我直接贴出来
```text
-----BEGIN PUBLIC KEY-----  
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC0Llg1bVZhnyslfezwfeOkvnXW  
q59bDtmQyHvxkP/38Fw8QQXBfROCgzGc+Te6pOPl6Ye+vQ1rAnisBaP3rMk40i3O  
pallzVkuwRKydek3V9ufPpZEEH4eBgInMSDiMsggTWxcI/Lvag6eHjkSc67RTrj9  
6oxj0ipVRqjxW4X6HQIDAQAB  
-----END PUBLIC KEY-----
```

# i学怎么通过网站去逆向
直接浏览器打开开发者工具 一步一步来就行 
挺简单的 它甚至还贴心的为你写了注释( laugh)

# 我经常看见的有一串数字什么意思
这里是不是说 1544059443 
这个好像是i学appid 这个网站不是只有i学一个app 还有其他很多app 一个域名下
这个数字不用太管 记得它是appid就行


