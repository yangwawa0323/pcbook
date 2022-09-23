## GRPC and GORM project demo

# Run GRPC server

In the root folder of this project. and run the following command.

```shell
	make server
```


# Run GRPC client

run the following command

```shell
   make client
```

> Note: By default, we are using MySQL server as the backend database.
> You can use the sqlite as the backend database as well.

# Change the backend database

modify the `sql/db_manager.go` file.
comment the `gorm.Open` and import package name to 

```go
import (


-  // "gorm.io/driver/mysql
+  "gorm.io/driver/sqlite"

)

func InitDB() (*gorm.DB, error) {
	infoLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Info,
			Colorful: true,
		},
	)

-	// db, err := gorm.Open(mysql.Open(GetMySqlDSN()), &gorm.Config{Logger: infoLogger})
+	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{Logger: infoLogger})
	if err != nil {
		log.Fatal(out.Panic("cannot connect to MySQL database: %v", err))
		return nil, err
	}
	return db, err
}

```

After change the two place in the code. run 

```shell
	go mod tidy
```
to get the addition package as nessesary.

> Note: 
> You are in the windows environment, the minimal requirement must met
>  *  gcc you can download
>     [TDM64-gcc](https://github.com/jmeubank/tdm-gcc/releases/download/v10.3.0-tdm-1/tdm-gcc-10.3.0.exe)
>  * make 
	[make for windows version](http://gnuwin32.sourceforge.net/downlinks/make.php)


## Fixed : using buf.build generate the new hierarchy package struct 

See official docs [buf](https://docs.buf.build)

For the reason of diffrence revision of the protobuffer, for example:

* `techschool.pcbook.pb.user.v1`
* `techschool.pcbook.pb.user.v1alpha`
* `techschool.pcbook.pb.user.v1beta`
* `techschool.pcbook.pb.user.v2`

wrote the `buf.yaml` configuration file to instead of `protoc -I <package_name>`
wrote the `directories` in the `buf.work.yaml` configuration file.
finally. wrote the `buf.gen.yaml` for `buf generate` command-line.