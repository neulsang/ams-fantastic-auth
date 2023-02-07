# Account Management Fantastic Auth

본 서비스는 인증 및 권한을 관리하는 서비스입니다.

크게 두가지의 기능을 수행합니다.

- Authentication : 인증
    * 로그인 자격을 증명하여 로그인한 사용자를 인식하는 것입니다.
        - 로그인 (Post /api/v1/auth/signin)
        - 로그아웃 (Delete /api/v1/auth/signout)

- Authorization : 권한
    * 로그인한 사용자가 무엇을 할 수 있는 규칙을 정하는 것입니다. 기본적으로 로그인한 사용자에 대한 CRUD(Create, Read, Update, Delete) 권한을 갖게 됩니다.
        - 회원가입 (Post /api/v1/users/signup)
        - 회원정보 조회 (Get /api/v1/users/id)
        - 회원정보 수정 (Put /api/v1/users/id)
        - 회원 탈퇴 (Delete /api/v1/user/id)

```
위 구조는 변경 될 수 있음을 명시한다.
```