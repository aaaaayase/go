package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"net/http"
	"strconv"
	"time"
	pb "week1/auth-server/proto"
)

func main() {
	router := gin.Default()
	router.POST("/v1/client/sign-in", func(c *gin.Context) {
		// 获取Authorization
		authorization := c.GetHeader("Authorization")
		// 如果Authorization为空 返回401
		if authorization == "" {
			c.JSON(http.StatusUnauthorized, gin.H{})
			return
		}
		// 获取x-timestamp-unix
		sendTimeStr := c.GetHeader("x-timestamp-unix")
		if sendTimeStr == "" {
			c.JSON(http.StatusBadRequest, gin.H{})
			return
		}
		sendTimeUnix, err := strconv.ParseInt(sendTimeStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{})
			return
		}
		// 转换为 time.Time
		sendTime := time.Unix(sendTimeUnix, 0)
		// 计算时间差
		currentTime := time.Now()
		timeDiff := currentTime.Sub(sendTime).Abs()
		if timeDiff > 5*time.Minute {
			c.JSON(http.StatusBadRequest, gin.H{})
			return
		}

		// 获取credential
		req := SignInRequest{}
		err = c.ShouldBindBodyWithJSON(&req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{})
			return
		}
		// 验证credential
		if req.XiaomiCredential.Session == "" || req.XiaomiCredential.Uid == "" {
			c.JSON(http.StatusUnauthorized, gin.H{})
			return
		}

		// 调用grpc
		comboId := signIn(c, &req)
		if comboId != "" {
			c.JSON(http.StatusOK, gin.H{"combo_Id": comboId})
		}

	})

	router.Run(":8081")
}

func signIn(c *gin.Context, req *SignInRequest) string {
	// grpc调用
	conn, err := grpc.Dial("127.0.0.1:8082", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return ""
	}
	defer conn.Close()
	// 跟服务端建立连接
	client := pb.NewAuthClient(conn)
	// 调用
	// 先构造需要的结构体
	// Credential
	pbCredential := pb.Credential{}
	pbCredential.GetXiaomi().Uid = req.XiaomiCredential.Uid
	pbCredential.GetXiaomi().Session = req.XiaomiCredential.Session
	pbCredential.Idp = pb.IdentityProvider_XIAOMI
	// ClientMeta
	clientMeta := pb.ClientMeta{}
	resp, err := client.SignIn(context.Background(), &pb.SignInRequest{GameId: "ggg",
		Credential: &pbCredential, ClientMeta: &clientMeta, UpushDeviceToken: req.UpushDeviceToken})
	if resp == nil && err != nil {
		if status.Code(err) == codes.Unauthenticated {
			c.JSON(http.StatusUnauthorized, gin.H{})
			return ""
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{})
			return ""
		}
	}
	return resp.ComboId
}
