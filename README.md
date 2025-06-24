# Backend Assignment

## 서버 실행 방법

```bash
go run .
```

## API 테스트 방법

- `./test` 폴더에는 통합 테스트 (integration test, end-to-end test) 코드들이
  있습니다.
- `./internal/server/update_issue_test` 파일에는 issue 수정 로직 테스트 코드가
  있습니다.

한꺼번에 테스트 코드 실행. (통합 테스트는 서버가 실행되어있어야 합니다.)

```bash
go test ./... -v # -v 옵션을 추가하면 로그가 출력됩니다.
```

또는 다음과 같이 각각 테스트를 실행할 수 있습니다.

```bash
go test ./test -v
go test ./internal/server -v
```

## Implementation

- 제공해주신 API 명세를 따라 구현을 완료하였습니다.
- 통합테스트는 예시 코드를 사용하여 구현하였습니다.
