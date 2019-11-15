<html>
<script src="https://code.jquery.com/jquery-3.2.1.slim.min.js" integrity="sha384-KJ3o2DKtIkvYIK3UENzmM7KCkRr/rE9/Qpg6aAZGJwFDMVNA/GpGFF93hXpG5KkN" crossorigin="anonymous"></script>
<!-- Latest compiled and minified CSS -->
<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css" integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">

<!-- Optional theme -->
<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap-theme.min.css" integrity="sha384-rHyoN1iRsVXV4nD0JutlnGaslCJuC7uwjduW9SVrLvRYooPp2bWYgmgJQIXwl/Sp" crossorigin="anonymous">

<!-- Latest compiled and minified JavaScript -->
<script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js" integrity="sha384-Tc5IQib027qvyjSMfHjOMaLkfuWVxZxUPnCJA7l2mCWNIpG9mGCD8wGNIcPD7Txa" crossorigin="anonymous"></script>
    <head>
        <title>Register</title>
    </head>
    <body> 
        <div class="col-lg-12">
        <form action="/register" method="post">
            <div class="form-group">
                <div class="col-sm-2">First Name</div>
                <div class="col-sm-4"><input type="text" class="form-control" name="firstname" required></div>
            </div>
            <div class="form-group">
                <div class="col-sm-2">Last Name</div>
                <div class="col-sm-4"><input type="text" class="form-control" name="lastname" required></div>
            </div>
            <div class="form-group">
                <div class="col-sm-2">Username</div>
                <div class="col-sm-4"><input type="text" class="form-control" name="username" required></div>
            </div>
            <div class="row"></div>
            <div class="row"></div>
            <div class="form-group">
                <div class="col-sm-2">Password</div>
                <div class="col-sm-4"><input type="password" class="form-control" name="password" required></div>
            </div>

            <div class="row"></div>
            <!--div class="form-group">
                
                <div class="col-sm-2">Gender</div>

                <div class="col-sm-2">
                 <input type="radio" name="gender" value="1"> <label>Male</label>
                 <input type="radio" name="gender" value="2"> <label>Female</label>
                 </div>
            </div-->
            
            <div class="col-md-12">
            <input type="submit" value="Register">
            </div>
        </form>
        </div>
    </body>
</html>