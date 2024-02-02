package contact

import (
	"net/http"

	"github.com/EMSI-zero/go-chat/domain/contact"
	"github.com/EMSI-zero/go-chat/infra/httputils"
	"github.com/EMSI-zero/go-chat/registry"
	"github.com/gin-gonic/gin"
)

type ContactController struct {
	contactController contact.ContactService
}

func NewContactController(sr registry.ServiceRegistry) *ContactController {
	return &ContactController{
		contactController: sr.GetContactService(),
	}
}

func (cc *ContactController) GetContacts(c *gin.Context) {
	ctx := c.Request.Context()

	res, err := cc.contactController.GetContacts(ctx)
	if err != nil {
		httputils.NewError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (cc *ContactController) AddContact(c *gin.Context) {
	ctx := c.Request.Context()

	req := new(contact.AddContectRequest)
	err := c.Bind(req)
	if err != nil {
		httputils.NewError(c, http.StatusBadRequest, err)
		return
	}

	err = cc.contactController.AddContact(ctx, req)
	if err != nil {
		httputils.NewError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (cc *ContactController) RemoveContact(c *gin.Context) {
	ctx := c.Request.Context()

	req := new(contact.RemoveContactRequest)
	err := c.Bind(req)
	if err != nil {
		httputils.NewError(c, http.StatusBadRequest, err)
		return
	}

	err = cc.contactController.DeleteContact(ctx, req)
	if err != nil {
		httputils.NewError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}
