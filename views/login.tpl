<!DOCTYPE html>

<html>
<head>
  <title>Web Messager Login</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <link href="/static/css/bootstrap.min.css" rel="stylesheet" type="text/css"/>
</head>

<body>
  <header>
    <h2>Login</h2>
  </header>
  {{if .HasError}}
  <div>
      <div class="alert alert-danger" role="alert">{{.ErrorMsg}}</div>
  </div>
  {{end}}
  <div class="row">
    <div class="col-md-6">
        <form action="/account/loginuser" method="post">
            <div class="form-group">
              <label for="UserName">User name</label>
              <input type="text" class="form-control" id="UserName" name="username" placeholder="User name" value="{{.Name}}">
            </div>
            <div class="form-group">
              <label for="InputPassword1">Password</label>
              <input type="password" class="form-control" id="InputPassword1" name="password" placeholder="Password">
            </div>
            <input type="submit" class="btn btn-info" value="Login" />
        </form>
    </div>
  </div>
  <div class="backdrop"></div>

  <script src="/static/js/reload.min.js"></script>
</body>
</html>
