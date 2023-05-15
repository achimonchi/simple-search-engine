# Comparison Search Engine using Meilisearch and Typesense
Simple search engine using Meilisearch and Typesense


## Techstack
- Docker
- Docker Compose
- Golang
- PostgreSQL
- Meilisearch (docker version)
- Typesense (docker version)
## How to Run
Running all dependencies first using :
```bash
make run-infra
```
This will running `docker-compose` inside folder **deploy**. 

Migrate our data :
```bash
make migrate-up
```
This will running migration and insert several data that already we provided in **deploy/data.json**.

And then, you can run the apps using ;
```bash
make run
```
