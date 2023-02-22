# Account Management Fantastic Auth

본 서비스는 인증 및 권한을 관리하는 서비스이다.

크게 두 가지의 기능을 수행한다.

- Authentication : 인증
  - 로그인 자격을 증명하여 로그인한 사용자를 인식
  - 로그인한 사용자에게 JWT을 발행
  - JWT 기반으로 다른 API 호출
  - 잘못된 JWT, 만료된 JWT로 API를 호출하게 되면 로그아웃 처리 후 JWT를 재발행

- Authorization : 권한
  - 로그인한 사용자가 무엇을 할 수 있는지 규칙 및 권한 정의
  - 기본적으로 로그인한 사용자에 대한 CRUD(Create, Read, Update, Delete) 권한을 소유

    `* 위 구조는 변경 될 수 있음을 명시한다.`

## Project Layout

---
go 기반 프로젝트 레이아웃에 대해 정의 한다. (TODO)

## Swagger command

---
swaggo 라는 go 기반 swagger를 install 해준다.

```shell
go instatll github.com/swaggo/swag/cmd/swag@latest
```

swag init을 통해 docs 문서를 만들어준다.

```shell
swag init -g ./cmd/ams-fantastic-auth/main.go
```

`* -g flag의 경우 main pacakge가 정의된 파일을 지정해준다.`

## go command

---
go run을 통해 바로 ams-fantastic-auth app을 실행 시켜줄수 있다.

```shell
go run cmd/ams-fantastic-auth/main.go

```

## docker build command

```shell
 docker build -t ams-fantastic-auth:1.0.0 -f build/Dockerfile .
 ```
