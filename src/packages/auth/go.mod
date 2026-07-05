module github.com/CSVaishakh/QuickHand/src/packages/auth

go 1.26.3

require github.com/CSVaishakh/QuickHand/src/packages/db v0.0.0

require golang.org/x/crypto v0.53.0

require github.com/golang-jwt/jwt/v5 v5.3.1 // direct

require (
	github.com/google/uuid v1.6.0 // direct
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/text v0.38.0 // indirect
	gorm.io/gorm v1.31.1 // direct
)

replace github.com/CSVaishakh/QuickHand/src/packages/db => ../db
