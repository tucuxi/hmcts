#!/bin/bash

API=http://localhost:3000/cases

for id in {1..1000}
do
  curl -s -X POST --json '{"id":'$id',"caseNumber":"cn","title":"t","status":"new","createdDate":"2025-04-20"}' "$API"
  echo 
done

for id in {1..1000}
do
  curl -s -X DELETE "$API/$id"
done
