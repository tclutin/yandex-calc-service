# yandex-calc-service
## Description
Exam assignments from Yandex Lyceum. The service for processing simple mathematical expressions
## Install

1.You need Go version 1.23 or higher. [Download](https://go.dev/dl/)

2.Clone this repository
```bash
git clone https://github.com/tclutin/yandex-calc-service
```
3.Navigate to this directory
```bash
cd yandex-calc-service
```
4.Run the service
```bash
go run ./cmd/app/main.go
```
The service will be available [here](http://localhost:8080)

## Endpoints
| **Endpoint**       | **Method** | **Status** | **Windows curl**                                                                 | 
|--------------------|------------|------------|----------------------------------------------------------------------------------|
| `/api/v1/calculate` | `POST`     | `200`      | `curl --location "http://localhost:8080/api/v1/calculate" --header "Content-Type: application/json" --data "{ \"expression\": \"2+2*2\" }"` | 
```json
{
  "result": 6
}
