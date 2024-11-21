package middleware

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/africhild/common/pkg/injection"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Transaction(d *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		tx := d.Begin()
		if tx.Error != nil {
			log.Print(tx.Error)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		injection.SetSQL(c, tx)

		defer func() {
			if err := recover(); err != nil {
				encodedError, marshalError := json.Marshal(err)
				if marshalError != nil {
					encodedError = []byte("Unable to marshall error")
				}

				log.Print(encodedError)

				err := tx.Rollback().Error
				if err != nil {
					log.Print("Cannot rollback transaction", err)
					return
				}
			} else if len(c.Errors) > 0 {
				err := tx.Rollback().Error
				if err != nil {
					log.Print("Cannot rollback transaction", err)
					return
				}
			} else if c.Writer.Status() < 200 || c.Writer.Status() >= 300 {
				err := tx.Rollback().Error
				if err != nil {
					log.Print("Cannot rollback transaction", err)
					return
				}
			} else {
				err := tx.Commit().Error
				if err != nil {
					log.Print(err, "Cannot commit transaction")
					return
				}
			}
		}()

		c.Next()
	}
}
