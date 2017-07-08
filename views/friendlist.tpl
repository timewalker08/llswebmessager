<!DOCTYPE html>

<html>
<head>
  <title>好友列表</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
   <link href="/static/css/bootstrap.min.css" rel="stylesheet" type="text/css"/>
   
   <script src="https://code.jquery.com/jquery-3.2.1.min.js"></script>
   <script>
       $(document).ready(function() {
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
                           }
                       }
                   });
               }
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
                     }
                   }
                 });
               }
             }
           })
       });
   </script>
</head>

<body>
  <header>
    <h1 class="logo">Welcome back!</h1>
  </header>
  <hr />
  <div class="row">
      <div class="col-md-3">
          <div class="input-group">
            <input type="text" class="form-control" id="NameForSearch" placeholder="Search for new friend...">
            <span class="input-group-btn">
              <button class="btn btn-default" id="SearchUser" type="button">Search!</button>
            </span>
          </div><!-- /input-group -->
      </div><!-- /.col-lg-6 -->
  </div>
  <br />
  <h3>My Friends</h3>
  <div class="row">
    <div class="col-md-3">
        <ul class="list-group" id="FriendList">
            {{range .FriendList}}
                <li class="list-group-item context">
                    <span class="badge">14</span>
                    <span class="friendname" data-name="{{.Friend.Name}}">{{.Friend.Name}}</span>
                </li>
            {{end}}
        </ul>
    </div>
  </div>
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
    
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-contextmenu.js"></script>
  <script src="http://cdnjs.cloudflare.com/ajax/libs/prettify/r224/prettify.js"></script>
  <script src="/static/js/reload.min.js"></script>
</body>
</html>
