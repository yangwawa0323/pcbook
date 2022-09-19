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