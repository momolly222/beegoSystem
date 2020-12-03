<!DOCTYPE html>
<html>

<head>
    <meta charset="UTF-8">
    <title>登陆</title>
</head>

<body>
    <form action="/doLogin" method="post">
        用户名 <input type="text" name="username" value={{.username}}><br><br>
        密 码 <input type="password" name="password" /><br><br>
        <input type="submit" value="提交" />
    </form>
</body>

</html>
