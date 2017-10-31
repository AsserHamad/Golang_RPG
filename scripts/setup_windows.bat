REM Install beego (Requires git)
go get github.com/astaxie/beego

REM In case https validation not working
REM git config --global http.sslVerify false

REM Installing bee
go get github.com/beego/bee

REM Govendor simulates package.json and helps in installing the required packages
go get -u github.com/kardianos/govendor

REM MySQL Database driver vroom vroom
go get -u github.com/go-sql-driver/mysql
