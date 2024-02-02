package chat

import (
	"net/http"
	"strconv"

	"github.com/EMSI-zero/go-chat/domain/chat"
	"github.com/EMSI-zero/go-chat/infra/httputils"
	"github.com/EMSI-zero/go-chat/registry"
	"github.com/gin-gonic/gin"
)

type ChatController struct {
	chatService chat.ChatService
}

func NewChatController(r registry.ServiceRegistry) *ChatController {
	return &ChatController{chatService: r.GetChatService()}
}

func (cc *ChatController) GetChats(c *gin.Context) {
	ctx := c.Request.Context()

	chats, err := cc.chatService.FetchConversations(ctx)
	if err != nil {
		httputils.NewError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, chats)
}

func (cc *ChatController) GetChat(c *gin.Context) {
	ctx := c.Request.Context()

	cid, err := strconv.Atoi(c.Param("chat_id"))
	if err != nil {
		httputils.NewError(c, http.StatusBadRequest, err)
	}

	chat, err := cc.chatService.FetchConversation(ctx, int64(cid))
	if err != nil {
		httputils.NewError(c, http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, chat)
}

func (cc *ChatController) NewChat(c *gin.Context) {
	ctx := c.Request.Context()

	req := new(chat.NewChatRequest)
	err := c.Bind(req)
	if err != nil {
		httputils.NewError(c, http.StatusBadRequest, err)
	}

	err = cc.chatService.CreatePrivateConversation(ctx, req)
	if err != nil {
		httputils.NewError(c, http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, nil)
}

func (cc *ChatController) LeaveChat(c *gin.Context) {
	ctx := c.Request.Context()
	cid, err := strconv.Atoi(c.Param("chat_id"))
	if err != nil {
		httputils.NewError(c, http.StatusBadRequest, err)
	}

	err = cc.chatService.RemovePrivateConversation(ctx, int64(cid))
	if err != nil {
		httputils.NewError(c, http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, nil)
}

func (cc *ChatController) DeleteMessage(c *gin.Context) {
	ctx := c.Request.Context()
	cid, err := strconv.Atoi(c.Param("chat_id"))
	if err != nil {
		httputils.NewError(c, http.StatusBadRequest, err)
	}

	mid, err := strconv.Atoi(c.Param("message_id"))
	if err != nil {
		httputils.NewError(c, http.StatusBadRequest, err)
	}

	req := new(chat.RemoveMessageRequest)
	req.ConversationID = int64(cid)
	req.MessageID = int64(mid)

	err = cc.chatService.RemoveMessage(ctx, req)
	if err != nil {
		httputils.NewError(c, http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, nil)

}
