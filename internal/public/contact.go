package public

import (
	"bookieguardserver/config"
	"bookieguardserver/internal/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ContactForm struct {
	Name    string `form:"name" json:"name" xml:"name"`
	Email   string `form:"email" json:"email" xml:"email" binding:"required"`
	Subject string `form:"subject" json:"subject" xml:"subject" binding:"required"`
	Message string `form:"message" json:"message" xml:"message" binding:"required"`
}

func Contact(c *gin.Context) {

	if c.Request.Method == "GET" {

		c.HTML(http.StatusOK, "public.contact.html", gin.H{
			"title":     "Contact",
			"canonical": "/contact",
		})
	}

	if c.Request.Method == "POST" {
		var json ContactForm
		if err := c.ShouldBind(&json); err != nil {
			c.String(http.StatusBadRequest, "%v", gin.H{"error": err.Error()})
			return
		}

		contact := models.Contact{}
		contact.Name = json.Name
		contact.Email = json.Email
		contact.Subject = json.Subject
		contact.Message = json.Message

		contact.Create()

		fmt.Println(contact)

		c.JSON(200, config.BodyStructure{
			"success": "true",
			"message": "Form submitted successfully.",
		})

	}
}
