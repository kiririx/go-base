package go_base

import (
	_ "github.com/Shopify/sarama"
	_ "github.com/gin-contrib/static"
	_ "github.com/gin-gonic/gin"
	_ "github.com/go-redis/redis/v9"
	_ "github.com/golang-jwt/jwt"
	_ "github.com/gorilla/websocket"
	_ "github.com/kiririx/krutils/strx"
	_ "github.com/satori/go.uuid"
	_ "gorm.io/driver/mysql"
	_ "gorm.io/driver/sqlite"
)
