JSCH_ENV=production
HTTP_PORT=8085

JWT_SIGNING_KEY=test
AUTH_PASSWORD_SALT=salt
AUTH_VERIFICATION_CODE_LENGTH=8

# For connect to exist mongo db
JSCH_MONGO_URI=test
JSCH_MONGO_USER=test
JSCH_MONGO_PASSWORD=test
#JSCH_MONGO_DATABASE=

# For initialization mongo db
MONGO_INITDB_ROOT_USERNAME=admin
MONGO_INITDB_ROOT_PASSWORD=qwerty
MONGODB_DATABASE=jobSchedule

.DEFAULT_GOAL := run

.PHONY:
.SILENT:
.EXPORT_ALL_VARIABLES:

build:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/app/main.go

run: build 
	docker-compose up --remove-orphans app
