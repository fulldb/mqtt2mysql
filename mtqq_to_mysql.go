package main

import (
    "fmt"
    "strconv"
    //import the Paho Go MQTT library
    MQTT "github.com/eclipse/paho.mqtt.golang"
    "os"
    "time"
    "database/sql"
     _ "github.com/go-sql-driver/mysql"
    cf "conf"
)

//定义mysql连接对象
var db *sql.DB

//define a function for the default message handler
var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
    mysqlemq(msg.Topic(),msg.Payload())
    fmt.Printf("TOPIC: %s\n", msg.Topic())
    fmt.Printf("MSG: %s\n", msg.Payload())
}

func main() {
    initmysql()
    //create a ClientOptions struct setting the broker address, clientid, turn
    //off trace output and set the default message handler
    //读取配置文件 
    myConfig := new(cf.Config)
    myConfig.InitConfig("config.txt")
    emqqurl := myConfig.Read("emqq", "emqqurl")
    emqquser := myConfig.Read("emqq", "emqquser")
    emqqpasswd := myConfig.Read("emqq", "emqqpasswd")
    fmt.Printf("emqqurl: %s\n", emqqurl)
    opts := MQTT.NewClientOptions().AddBroker(emqqurl)
    opts.SetClientID("go-simple")
    opts.SetUsername(emqquser)
    opts.SetPassword(emqqpasswd)
    opts.SetDefaultPublishHandler(f)
    
    //create and start a client using the above ClientOptions
    c := MQTT.NewClient(opts)
    if token := c.Connect(); token.Wait() && token.Error() != nil {
      panic(token.Error())
    }
    
    //subscribe to the topic /go-mqtt/sample and request messages to be delivered
    //at a maximum qos of zero, wait for the receipt to confirm the subscription
    if token := c.Subscribe("/#", 0, nil); token.Wait() && token.Error() != nil {
      fmt.Println(token.Error())
      os.Exit(1)
    }
    
    //Publish 5 messages to /go-mqtt/sample at qos 1 and wait for the receipt
    //from the server after sending each message
    for i := 0; i < 5; i++ {
      text := fmt.Sprintf("this is msg%d!", i)
      token := c.Publish("/test", 0, false, text)
      token.Wait()
    }
    
    time.Sleep(60*60*24*360 * time.Second)
    
    //unsubscribe from /go-mqtt/sample
    if token := c.Unsubscribe("go-mqtt/sample"); token.Wait() && token.Error() != nil {
      fmt.Println(token.Error())
      os.Exit(1)
    }
    
    c.Disconnect(250)
}


/***
*mtqq数据插入mysql数据库
***/
func mysqlemq(x string, y []byte) {
    stmt, err := db.Prepare("INSERT INTO user_info SET username=?,departname=?,create_time=?")
    fmt.Println("mysqlemq1")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(time.Now())
    res, err := stmt.Exec(x, y, time.Now())
    id, err := res.LastInsertId()
    if err != nil {
        panic(err)
    }
    defer db.Ping()
    fmt.Println("id:",id)
}
  
/***
*mysql连接池初始化

***/
func initmysql() {    
    myConfig := new(cf.Config)
	myConfig.InitConfig("config.txt")
	mysqlurl := myConfig.Read("mysql", "mysqlurl")
    fmt.Printf("mysqlurl: %s\n", mysqlurl)
    db,_= sql.Open("mysql", mysqlurl)
    //fmt.Println(db.Ping())  检查是否连接成功数据库
    //从配置文件中读取maxopenconns最大打开的连接数
    maxopenconns:= myConfig.Read("mysql", "maxopenconns")
    //从配置文件中读取maxidleconns最大空闲连接数
    maxidleconns:= myConfig.Read("mysql", "maxidleconns")
    maxopenconnsint,err:=strconv.Atoi(maxopenconns)
    maxidleconnsint,err:=strconv.Atoi(maxidleconns)
    if err != nil {
        panic(err)
    }
    db.SetMaxOpenConns(maxopenconnsint)
    db.SetMaxIdleConns(maxidleconnsint)
    db.Ping()
    fmt.Println("mysql连接池初始化成功")
}
