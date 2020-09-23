module github.com/nekishdev/graphql-example-catalog

go 1.14

require (
	github.com/99designs/gqlgen v0.13.0
	github.com/jinzhu/gorm v1.9.16
	github.com/joho/godotenv v1.3.0
	github.com/vektah/gqlparser/v2 v2.1.0
	gorm.io/gorm v1.20.1
)

replace (
	github.com/99designs/gqlgen v0.13.0 => github.com/99designs/gqlgen v0.12.2
)