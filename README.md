# Project
short url 을 생성, 관리, 연결 해주는 웹사이트 

# Components
````
Golang
Beego
angularJS
Redis
````

# dependencies
````
Redis v3.2.8
Go 1.8

````


## Installation & Execution
````

$go get github.com/HJaeH/short.ly
$cd $GOPATH/src/github.com/short.ly
$go get -v
$go run main.go

http://127.0.0.1:8080

````



## Test
````
$cd $GOPATH/src/short.ly
$go test ./...

````

## 고려 사항
````
 주로  short URL 생성은 hash함수를 사용하는데,
 
 shortURL서비스는 hash 함수는 가지는 특성을 활용하기
 
 적합하지 않습니다.
 
 보안 hash나 데이터 검증 hash함수들의 특성인 원본 url에 대한 복호 난이도, 
 
 또는 check sum, crc check 등의  데이터 검증도 필요하지 않으며 
 
 오히려 충돌상황에 대한 고려가 필요하게 됩니다. 
 
 base62의 중복하지않은 단순한 값이 url생성에 필요한 것으로 보입니다.
 
 0부터 등록되는 url마다 1씩 카운트 업을 하여 base62로 표현 및 map으로 원본 url에 접근
 
 할 수 있도록 한다면 모든 기능을 구현 가능하다고 생각하였습니다.
```` 

## 해결 방법
````
 다만 실 서비스에서 활용되기 위한 무작위한 hash형태의 값을 short url로 전달하기 위해서
 
 redis list에 0 부터 일정 수의 값을 미리 등록
  
 랜덤으로 하나의 인덱스 값을 얻고, 해당 인덱스를 삭제 후 얻은 값을 base62 encode
 
 하는 방식으로 해시 충돌에 대한 고민 없이, 중복 검사에 대한 리소스 없이
 
 빠르게 중복되지 않는 shorturl값을 생성 하도록 하였습니다.
 
   
 
 
 
````
