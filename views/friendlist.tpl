<!DOCTYPE html>

<html>
<head>
  <title>好友列表</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
   <link href="/static/css/bootstrap.css" rel="stylesheet" type="text/css"/>
   
   <script src="https://code.jquery.com/jquery-3.2.1.min.js"></script>
   <script>
       $(document).ready(function() {
           currentName = "";
           
           $("#SearchUser").click(function(e) {
               var newName = $("#NameForSearch").val();
               if (newName != "") {
                   $.ajax({
                       type: "get",
                       url: "/friend/queryname?name=" + newName,
                       async: true,
                       success: function(data) {
                           if (data != null) {
                               $("#NewFriendName").html(data.Name);
                               $("#SearchUserResult").modal("show");
                           }
                       }
                   });
               }
               else {
                   alert("Please input new friend name to search!");
               }
           });
           
           $("#AddNewFriend").click(function(e) {
               var newName = $("#NewFriendName").html();
               if (newName != "") {
                   $.ajax({
                       type: "post",
                       url: "/friend/add?name=" + newName,
                       async: true,
                       success: function(data) {
                           if (data != null && data.Code == 0) {
                               alert("Succeeded");
                               $("#SearchUserResult").modal("hide");
							   var str = '<li class="list-group-item context friend-item"><span class="badge unread-count">0</span><span class="friendname" data-name="' + newName + '">' + newName + '</span></li>'
							   $("#FriendList").append(str);
							   FuncFriendItem();
                           } else {
						       alert(data.Msg)
						   }
                       }
                   });
               }
           });
           
           $("#SendNewMessage").click(function(e) {
               if (currentName == "") {
                   alert("Please select a friend.");
               }
               var msg = $("#NewMessage").val();
               if (msg == "") {
                   alert("Please input message first.");
               }
               $.ajax({
                   type: "post",
                   url: "/message/new?name=" + currentName + "&msg=" + msg,
                   async: true,
                   success: function(data) {
                       if (data.Code == 0) {
                           var str = '<li class="list-group-item msg-item list-group-item-success" style="text-align:right;">' + msg + "</li>";
                           $("#MsgList").append(str);
                           $("#NewMessage").val("");
                       } else {
                           alert("Failed to send message. Detail: " + data.Msg)
                       }
                   }
               });
           });
		   
		   FuncFriendItem();
       });
	   
	   function FuncMsgContextMenu() {
	       $('li.self-msg-context').contextmenu({
             target:'#msg-context-menu', 
             before: function(e,context) {
               // execute code before context menu if shown
               //alert($(e.target).attr("data-msg-id"));
               $("#DeleteMessage").attr("data-msg-id", $(e.target).attr("data-msg-id"));
			   $("#DeleteMessage").attr("data-msg", $(e.target).html());
             },
             onItem: function(context,e) {
               // execute on menu item selection
               var msgId = $(e.target).attr("data-msg-id");
			   var msg = $(e.target).attr("data-msg");
               if (confirm("Are you sure to delete message " + msg) == true) {
                 $.ajax({
                   type: "post",
                   url: "/message/remove?msgId=" + msgId,
                   async: true,
                   success: function(data) {
                     if (data != null && data.Code == 0) {
                        alert("Succeeded");
                        var selectedLi = $("li.msg-item[data-msg-id='" + msgId + "']", $("ul#MsgList"));
                        $(selectedLi).remove();
                     }
                   }
                 });
               }
             }
           });
	   }
	   
	   function FuncFriendItem() {
	       $("li.friend-item").click(function(e) {
             var name = $("span.friendname", this).attr("data-name");
             currentName = name
		     $("#CurrentChatName").html(name)
             $("#MsgList").children().remove();
			 $("span.unread-count", this).html('0');
             $.ajax({
               type: "get",
               url: "/message/all?name=" + name,
               async: true,
               success: function(data) {
                 if (data.length > 0) {
                   $.each(data, function(i, n) {
                     var isSelf = n.From.Name != name;
                     var str = '<li class="list-group-item msg-item';
                     if(isSelf){
                         str += ' list-group-item-success self-msg-context" style="text-align:right;"';
                     } else {
                         str += ' list-group-item-warning"';
                     }
                     str += ' data-msg-id="' + n.Id + '">';
                     str += n.Msg;
                     str += '</li>'
                     $("#MsgList").append(str);
                   });
				   FuncMsgContextMenu();
                 }
               }
             });
           });
		   
		   $('.context').contextmenu({
             target:'#context-menu', 
             before: function(e,context) {
               // execute code before context menu if shown
               // alert($("span.friendname", $(e.target)).text());
               $("#DeleteFriend").attr("data-name", $("span.friendname", $(e.target)).text());
             },
             onItem: function(context,e) {
               // execute on menu item selection
               var fname = $(e.target).attr("data-name");
               if (confirm("Are you sure to delete friend " + fname) == true) {
                 $.ajax({
                   type: "post",
                   url: "/friend/remove?name=" + fname,
                   async: true,
                   success: function(data) {
                     if (data != null && data.Code == 0) {
                        alert("Succeeded");
                        var selectedSpan = $("span.friendname[data-name='" + fname + "']", $("ul#FriendList"));
                        var selectedLi = $(selectedSpan).parents("li");
                        $(selectedLi).remove();
                     } else {
						alert(data.Msg)
					 }
                   }
                 });
               }
             }
           });
	   }
   </script>
   <style>
   </style>
</head>

<body>
  <div class="row">
      <div class="col-md-8 col-md-offset-1">
         <h1 class="logo">{{.UserName}}, Welcome back!</h1>
      </div>
  </div>
  <hr />
  <div class="row" style="height:600px;">
      <div class="col-md-3 col-md-offset-1">
          <div class="input-group">
            <input type="text" class="form-control" id="NameForSearch" placeholder="Search for new friend...">
            <span class="input-group-btn">
              <button class="btn btn-default" id="SearchUser" type="button">Search!</button>
            </span>
          </div>
          
          <h3>My Friends</h3>
      
          <ul class="list-group" id="FriendList">
              {{range .FriendList}}
                  <li class="list-group-item context friend-item">
                      <span class="badge unread-count">{{.UnreadCount}}</span>
                      <span class="friendname" data-name="{{.Friend.Friend.Name}}">{{.Friend.Friend.Name}}</span>
                  </li>
              {{end}}
          </ul>
      </div>
      
      <div class="col-md-7">
	    <h5 id="CurrentChatName">&nbsp</h5>
        <div style="height:450px;overflow-y:scroll;border:1px solid #DDDDDD;">
            <ul class="list-group" id="MsgList">
                
            </ul>
        </div>
        <div style="height:150px;border:1px solid #DDDDDD;">
            <textarea style="height:100px;width:100%;resize:none;" id="NewMessage"></textarea>
            <div style="height:50px;line-height:50px;width:100%;position:relative;text-align:right;">
                <button class="btn btn-info" id="SendNewMessage" style="">Send message</button>
            </div>
        </div>
      </div>
  </div>
  <br />
  
  <div class="backdrop"></div>

  <div class="modal fade" tabindex="-1" role="dialog" id="SearchUserResult">
    <div class="modal-dialog" role="document">
      <div class="modal-content">
        <div class="modal-header">
          <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
          <h4 class="modal-title">New Friend</h4>
        </div>
        <div class="modal-body">
          <p id="NewFriendName"></p>
        </div>
        <div class="modal-footer">
          <button type="button" class="btn btn-default" data-dismiss="modal">Close</button>
          <button type="button" class="btn btn-primary" id="AddNewFriend">Add friend</button>
        </div>
      </div><!-- /.modal-content -->
    </div><!-- /.modal-dialog -->
  </div><!-- /.modal -->
  
  <div id="context-menu">
    <ul class="dropdown-menu" role="menu">
      <li><a tabindex="-1" href="#" operator="top" id="DeleteFriend">Delete friend</a></li>
    </ul>
   </div>
   
   <div id="msg-context-menu">
    <ul class="dropdown-menu" role="menu">
      <li><a tabindex="-1" href="#" operator="top" id="DeleteMessage">Delete message</a></li>
    </ul>
   </div>
    
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-contextmenu.js"></script>
  <script src="http://cdnjs.cloudflare.com/ajax/libs/prettify/r224/prettify.js"></script>
  <script src="/static/js/reload.min.js"></script>
</body>
</html>
