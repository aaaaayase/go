package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-viper/mapstructure/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"time"
	pb "week1/auth-server/proto"
	"week1/global"
)

func (server *Server) SignIn(c context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {
	// 使用小米SDK来验证登录
	// 名称	重要性	说明
	//	appId	必须	游戏ID
	//	session	必须	用户sessionID
	//	uid	必须	用户ID
	//	signature	必须	签名,签名方法见后面说明
	// 拼接要加密的字符串
	str := "appId=" + req.GameId + "&" + "session=" + req.Credential.GetXiaomi().Session + "&" + "uid=" + req.Credential.GetXiaomi().Uid
	encrypt, err := HmacSHA1Encrypt(str, SecretKey)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// 准备请求的数据
	var reqData = XiaomiRequest{
		AppID:     req.GameId,
		Session:   req.Credential.GetXiaomi().Session,
		Uid:       req.Credential.GetXiaomi().Uid,
		Signature: encrypt,
	}
	// 序列化为json
	jsonData, err := json.Marshal(reqData)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// 创建HTTP请求
	url := "https://mis.migc.xiaomi.com/api/biz/service/loginvalidate"
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Request creation error:", err)
		return nil, err
	}
	// 设置请求头
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// 创建HTTP客户端并发送请求
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("Request sending error:", err)
		return nil, err
	}
	defer resp.Body.Close()
	// 读取响应
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Response reading error:", err)
		return nil, err
	}
	// 解析响应
	var response XiaomiResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("JSON unmarshal error:", err)
		return nil, err
	}
	if response.Errcode != 200 {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized")
	}

	var count int64
	global.Db.AutoMigrate(&ComboUser{})
	global.Db.AutoMigrate(&RealNameCache{})
	// 判断是否是第一次登录
	// 将传入数据作为唯一标识查询记录数
	global.Db.Where("external_id==?", req.Credential.Idp.String()).Count(&count)
	var comboIdR string
	if count == 0 {
		// 获取comboId
		err, comboIdT := getComboId(req)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		// 缓存表
		cache := RealNameCache{
			ExternalID:  req.Credential.Idp.String(),
			Idp:         8,
			PersonID:    8,
			DateOfBirth: time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC),
		}
		comboId, err := strconv.ParseInt(comboIdT, 10, 64)

		// 对象转map
		var meta map[string]string
		mapstructure.Decode(req.ClientMeta, &meta)
		personId := uint32(100)
		date := time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)
		user := ComboUser{
			ComboID:      uint64(comboId),
			Idp:          8,
			ExternalID:   req.Credential.Idp.String(),
			ExternalName: req.Credential.Idp.String(),
			GameID:       "ggg",
			Meta:         meta,
			ClientMeta:   meta,
			PersonID:     &personId,
			DateOfBirth:  &date,
		}

		// 插入
		global.Db.Create(&user)
		global.Db.Create(&cache)
		comboIdR = comboIdT
		signInResp := pb.SignInResponse{
			ComboId:   comboIdR,
			Identity:  &pb.Identity{},
			IsNewUser: true,
		}
		return &signInResp, nil
	}
	global.Db.Select("combo_id").Where("external_id=?", req.Credential.Idp.String()).Scan(&count)
	signInResp := pb.SignInResponse{
		ComboId:   comboIdR,
		Identity:  &pb.Identity{},
		IsNewUser: false,
	}
	return &signInResp, nil
}

func main() {
	// rpc服务dd
	// 开启端口
	listen, _ := net.Listen("tcp", ":8082")
	// 创建grpc服务
	newServer := grpc.NewServer()
	// 在grpc服务中去注册我们自己的服务
	pb.RegisterAuthServer(newServer, &Server{})
	// 启动服务
	err := newServer.Serve(listen)
	if err != nil {
		fmt.Println("failed to serve:%v", err)
		return
	}
}

func getComboId(req *pb.SignInRequest) (error, string) {
	// 获取当前 UTC 时间
	utcTime := time.Now().UTC()

	// 格式化为 yyMMdd
	datePart := utcTime.Format("060102") // "060102" → "yyMMdd"

	// 计算当天已过去的秒数（0-86400）
	secondsSinceMidnight := utcTime.Hour()*3600 + utcTime.Minute()*60 + utcTime.Second()

	// 格式化为 5 位数字（不足补零）
	secondsPart := fmt.Sprintf("%05d", secondsSinceMidnight)

	// 组合成 yyMMddsssss
	timeStr := datePart + secondsPart

	// 拼接redis key
	key := "auth:id:" + timeStr
	pipeline := global.RedisDB.Pipeline()
	ret := pipeline.Incr(key)
	pipeline.Expire(key, 86400*time.Second)
	pipeline.Exec()
	if ret.Val() >= 10000 {
		return errors.New("too many requests"), ""
	}
	return nil, timeStr + fmt.Sprintf("%04d", ret.Val())

}
