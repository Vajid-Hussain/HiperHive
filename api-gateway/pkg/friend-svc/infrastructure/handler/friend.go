package handler_friend_svc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	requestmodel_friend_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/friend-svc/infrastructure/model/requestModel"
	responsemodel_friend_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/friend-svc/infrastructure/model/responseModel"
	"github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/friend-svc/pb"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type FriendSvc struct {
	clind pb.FriendServiceClient
	redis *Caching
}

func NewFriendSvc(clind pb.FriendServiceClient, redis *Caching) *FriendSvc {
	return &FriendSvc{clind: clind,
		redis: redis}
}

func (h *FriendSvc) FriendRequest(ctx echo.Context) error {
	var req requestmodel_friend_svc.FriendRequest
	ctx.Bind(&req)

	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	result, err := h.clind.FriendRequest(context, &pb.FriendRequestRequest{
		UserID:   ctx.Get("userID").(string),
		FriendID: req.FriendID,
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responsemodel_friend_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}
	return ctx.JSON(http.StatusOK, responsemodel_friend_svc.Responses(http.StatusOK, "Friend request send succesfully", result, nil))
}

func (h *FriendSvc) GetFriends(ctx echo.Context) error {

	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	result, err := h.clind.FriendList(context, &pb.FriendListRequest{Friend: &pb.GetPendingListRequestModel{
		UserID: ctx.Get("userID").(string),
		OffSet: ctx.QueryParam("page"),
		Limit:  ctx.QueryParam("limit"),
	}})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responsemodel_friend_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}
	return ctx.JSON(http.StatusOK, responsemodel_friend_svc.Responses(http.StatusOK, "", result, nil))
}

func (h *FriendSvc) GetReceivedFriendRequest(ctx echo.Context) error {
	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	result, err := h.clind.GetReceivedFriendRequest(context, &pb.GetReceivedFriendRequestRequest{
		Received: &pb.GetPendingListRequestModel{
			UserID: ctx.Get("userID").(string),
			OffSet: ctx.QueryParam("page"),
			Limit:  ctx.QueryParam("limit"),
		},
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responsemodel_friend_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}
	return ctx.JSON(http.StatusBadRequest, responsemodel_friend_svc.Responses(http.StatusOK, "", result, nil))
}

func (h *FriendSvc) GetSendFriendRequest(ctx echo.Context) error {
	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	result, err := h.clind.GetSendFriendRequest(context, &pb.GetSendFriendRequestRequest{
		Send: &pb.GetPendingListRequestModel{
			UserID: ctx.Get("userID").(string),
			OffSet: ctx.QueryParam("page"),
			Limit:  ctx.QueryParam("limit"),
		},
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responsemodel_friend_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}
	return ctx.JSON(http.StatusBadRequest, responsemodel_friend_svc.Responses(http.StatusOK, "", result, nil))
}

func (h *FriendSvc) GetBlockFriendRequest(ctx echo.Context) error {
	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	result, err := h.clind.GetBlockFriendRequest(context, &pb.GetBlockFriendRequestRequest{
		Block: &pb.GetPendingListRequestModel{
			UserID: ctx.Get("userID").(string),
			OffSet: ctx.QueryParam("page"),
			Limit:  ctx.QueryParam("limit"),
		},
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responsemodel_friend_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}
	return ctx.JSON(http.StatusBadRequest, responsemodel_friend_svc.Responses(http.StatusOK, "", result, nil))
}

func (h *FriendSvc) UpdateFriendshipStatus(ctx echo.Context) error {
	var req requestmodel_friend_svc.FrendShipStatusUpdate
	ctx.Bind(&req)

	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := h.clind.UpdateFriendshipStatus(context, &pb.UpdateFriendshipStatusRequest{

		FriendShipID: req.FrendShipID,
		Status:       ctx.QueryParam("action"),
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responsemodel_friend_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}
	return ctx.JSON(http.StatusBadRequest, responsemodel_friend_svc.Responses(http.StatusOK, "succesfully updated as "+ctx.QueryParam("action"), "", nil))
}

var upgrade = websocket.Upgrader{}

type Message struct {
	SenderID    int       `json:"sender_id"`
	RecipientID int       `json:"recipient_id"`
	Content     string    `json:"content"`
	Timestamp   time.Time `json:"timestamp"`
}

type WebSocketInfo struct {
	RemoteAddr string `json:"remote_addr"`
}

func (h *FriendSvc) FriendMessage(ctx echo.Context) error {
	fmt.Println("message called")

	conn, err := upgrade.Upgrade(ctx.Response(), ctx.Request(), nil)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responsemodel_friend_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}
	defer fmt.Println("closed")
	defer conn.Close()

	err = h.redis.CachingUserConnection(ctx.Get("userID").(string), conn)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responsemodel_friend_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}

	for {
		err := conn.WriteMessage(websocket.TextMessage, []byte("hellow world"))
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, responsemodel_friend_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
		}

		_, msg, err := conn.ReadMessage()
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, responsemodel_friend_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
		}
		fmt.Println(string(msg))

		var message Message
		if err := json.Unmarshal([]byte(msg), &message); err != nil {
			return err
		}

		err = h.redis.SendMessage(message)
		if err != nil {
			return err
		}

		err = json.Unmarshal(msg, &message)
		if err != nil {
			fmt.Println("--", err)
		}
		fmt.Println(message, message)

	}

	return errors.New("work done")
}
