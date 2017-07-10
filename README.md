# Web私信系统
## 1 语言、系统和框架
1 采用go语言开发系统的后台。Web框架采用beego，原因是其文档较为完善，上手比较快，同时也支持MVC。Web服务器架设在Windows系统上。当然架设到Linux系统上也没问题。<br /><br />
2 数据库采用了mysql，部署在centos7系统上。<br />

## 2 数据库设计
数据库脚本放在 script/mysql/deploy.sql文件中。总共有6个表。<br /><br />
user表用于保存注册用户的数据。其中name列保存用户名，password_md5保存用户密码的md5（暂时保存原文）。<br /><br />
friend表用于保存朋友关系，主动加好友者为user_id，被动好友者为friend_id。有个状态列friendstatus_id用户指示主动好友者是否删除被动好友者，是指向friendstatus表的外键。<br /><br />
friendstatus表是domain table，用于保存好友状态名。<br /><br />
message表用于保存消息，from_id是发送者，to_id是接收者，msg是消息内容，messagestatus_id是消息状态，用于指示是否被删除。<br /><br />
messagestatus表也是domain table，用于保存消息的状态。<br /><br />
lastreadmessagetime表用于保存用户对好友消息最后的阅读时间。在此时间之后发送的消息被视为未读消息。<br /><br />

## 3 Web
采用beego的MVC框架。<br /><br />

### Controller
主要写了3个controller，分别是用于处理账号的AccountController，用于处理好友关系的FriendController，以及用于处理消息的MessageController。<br /><br />
#### AccountController
[get] account/register 打开注册页面 <br />
[get] account/login    打开登录页面 <br />
[post] account/registeruser 处理注册请求 <br />
[post] account/loginuser 处理登录请求 <br />

#### FriendController
[get] friend/list     打开好友页面 <br />
[web api][get] friend/queryname  通过账号查询好友 <br />
[web api][post] friend/add       通过账号名添加好友 <br />
[web api][post] friend/delete    通过账号名删除好友 <br />

#### MessageController
[web api][get] message/all       获取所有消息 <br />
[web api][post] message/new      发送消息 <br />
[web api][post] message/remove   删除自己已发送的消息 <br />

### Models
基本的struct放在 models/models.go文件中。主要包括 User, Friendstatus, Friend, FriendWithUnReadCount, Messagestatus, Message, Lastreadmessagetime。<br /><br />
另外写了三个manager用于处理对数据的请求。分别是用于处理账号的AccountManager， 处理好友关系的FriendManager， 处理消息的MessageManager。<br /><br />
#### AccountManager
AddNewFriend用于添加好友（底层调用FriendManager的CreateOrUpdateFriend）。<br /><br />
DeleteFriendByName用于删除好友（底层调用FriendManager的DeleteFriend）。<br /><br />
SendMessage用于发送消息（底层调用MessageManager的SendMessage）。<br /><br />
DeleteMessage用于删除消息（底层调用MessageManager的UpdateMessageStatus）。<br /><br />
GetMessagesByPage用于获取消息（底层调用MessageManager的GetMessagesByPage）。<br /><br />

#### FriendManager
CreateOrUpdateFriend用于添加或更新好友。<br /><br />
DeleteFriendByName用于更新好友状态为Deleted。<br /><br />
GetFriends用于获取好友。<br /><br />

#### MessageManager
SendMessage用于发送消息。<br /><br />
UpdateMessageStatus用于更新消息的状态。<br /><br />
SetLastReadTime用于更新最后阅读时间。<br /><br />
GetUnReadMessageCount用于获取未读消息数量。<br /><br />
GetMessagesByPage用于获取消息。（本想实现分页功能，但时间来不及，后续有时间再加上）<br /><br />
