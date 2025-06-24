# Backend Assignment

## 서버 실행 방법

```bash
go run .
```

## API 테스트 방법

```bash
go test ./... -v # -v 옵션을 추가하면 로그가 출력됩니다.
```

## 구현 내용

- Available endpoints
  - PATCH /issue:id 는 시간이 부족해서 구현 완료 안됨.
- Test
  - Integration test에 다음 endpoint들을 추가하였습니다.
    - POST /issue
    - GET /issues
