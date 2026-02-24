#!/usr/bin/env bash

BASE_URL="http://localhost:8080/users"

echo "----------------------------------------"
echo "1) GET all users (initial state)"
echo "----------------------------------------"
curl -s -X GET "$BASE_URL" | jq .
echo -e "\n"

echo "----------------------------------------"
echo "2) POST create user #1"
echo "----------------------------------------"
curl -s -X POST "$BASE_URL" \
  -H "Content-Type: application/json" \
  -d '{"name":"Rakesh","email":"rakesh@gmail.com"}' | jq .
echo -e "\n"

echo "----------------------------------------"
echo "3) POST create user #2"
echo "----------------------------------------"
curl -s -X POST "$BASE_URL" \
  -H "Content-Type: application/json" \
  -d '{"name":"John","email":"john@example.com"}' | jq .
echo -e "\n"

echo "----------------------------------------"
echo "4) GET all users (after inserts)"
echo "----------------------------------------"
curl -s -X GET "$BASE_URL" | jq .
echo -e "\n"

echo "----------------------------------------"
echo "5) POST invalid request (missing email)"
echo "----------------------------------------"
curl -s -X POST "$BASE_URL" \
  -H "Content-Type: application/json" \
  -d '{"name":"InvalidUser"}' | jq .
echo -e "\n"

echo "----------------------------------------"
echo "Demo complete"
echo "----------------------------------------"