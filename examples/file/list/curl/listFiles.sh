curl "https://api.m3o.com/v1/file/List" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer $MICRO_API_TOKEN" \
-d '{
  "project": "examples"
}'