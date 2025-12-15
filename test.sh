#!/bin/bash

API_KEY="my-secret-key"

echo "CREATE"
curl -X POST http://localhost:8080/students \
-H "Content-Type: application/json" \
-H "X-API-Key: $API_KEY" \
-d '{"name":"Ajith","age":29}'

echo -e "\n\nGET ALL"
curl -X GET http://localhost:8080/students \
-H "X-API-Key: $API_KEY"

echo -e "\n\nUPDATE"
curl -X PUT http://localhost:8080/students/1 \
-H "Content-Type: application/json" \
-H "X-API-Key: $API_KEY" \
-d '{"name":"Ajith Kumar","age":30}'

echo -e "\n\nGET BY ID"
curl -X GET http://localhost:8080/students/1 \
-H "X-API-Key: $API_KEY"

echo -e "\n\nDELETE"
curl -X DELETE http://localhost:8080/students/1 \
-H "X-API-Key: $API_KEY"

echo -e "\n\nGET AGAIN"
curl -X GET http://localhost:8080/students \
-H "X-API-Key: $API_KEY"
