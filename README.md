readme
===================
1. 充值: https://www.skrill.com/fileadmin/content/pdf/Skrill_Quick_Checkout_Guide.pdf
2. 提现: https://www.skrill.com/fileadmin/content/pdf/Skrill_Automated_Payments_Interface_Guide.pdf


鉴权
==============
1. 充值: 是把对应的信息都发给了前端, 前端构造了form表单提交,  随后psp做跳转到收银台让用户支付. (form中的信息都是一些公开的账号信息,不包含任何秘钥等信息)
2. 充值回调: 回调信息中有一个字段md5签名, 需要对之进行鉴权
3. 提现: 是一个api请求, server构造了一个request, 里边包含了自己的账号/密码信息(以此让skrill来认可请求的合法性),此接口不需要任何回调

联系方式
=============
merchantservices@skrill.com


回调地址 
==============
在下单时，参数里传递进来的(动态可修改)


Comment
===============
1. both support deposit && withdrawl