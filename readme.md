# Projet [![CI-Test](https://github.com/NicolasMartino/simplebank/actions/workflows/ci.yaml/badge.svg?branch=master&event=push)](https://github.com/NicolasMartino/simplebank/actions/workflows/ci.yaml)

# Docker: 
## ops
- docker pull postgres:12-alpine
- docker run --name postgres12 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:12-alpine
- docker exec -it postgres12 psql -U root
- docker logs postgres12
- docker stop postgres12
## Db operations:
- docker exec -it postgres12 createdb --username=root --owner=root simple_bank
- docker exec -it postgres12 psql simple_bank

# Migration:
Install scoop: https://scoop.sh/

Install make: scoop install make

Install migrate: scoop install migrate

    https://github.com/golang-migrate/migrate

- migrate create -ext sql -dir db/migration -seq init_schema

Tutorial:
https://www.youtube.com/watch?v=rx6CPDK_5mU&ab_channel=TECHSCHOOL

https://www.youtube.com/watch?v=rx6CPDK_5mU&ab_channel=TECHSCHOOL

https://www.youtube.com/watch?v=Q9ipbLeqmQo&ab_channel=TECHSCHOOL

https://www.youtube.com/watch?v=0CYkrGIJkpw&ab_channel=TECHSCHOOL

https://www.youtube.com/watch?v=prh0hTyI1sU&ab_channel=TECHSCHOOL
