package go_base

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/websocket"
	"github.com/kiririx/krutils/strx"
	uuid "github.com/satori/go.uuid"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	_ = redis.Client{}
	_ = mysql.Config{}
	_ = gin.H{}
	_ = gorm.Model{}
	_ = websocket.Conn{}
	_ = uuid.DomainOrg
	_ = sqlite.DriverName
	_ = strx.ToStr("")
	_ = static.INDEX
	_ = jwt.ErrECDSAVerification
)
