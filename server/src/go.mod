module server

go 1.15

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-contrib/cors v1.3.0
	github.com/gin-contrib/sessions v0.0.3
	github.com/gin-gonic/gin v1.5.0
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/google/uuid v1.1.1
	github.com/joho/godotenv v1.3.0
	golang.org/x/crypto v0.0.0-20191205180655-e7c4368fe9dd
	gopkg.in/go-playground/validator.v8 v8.18.2
	gopkg.in/yaml.v2 v2.2.2
	gorm.io/driver/mysql v1.0.3
	gorm.io/gorm v1.20.8
)
