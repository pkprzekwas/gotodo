### Local setup

Starting postgres in docker:
```bash
docker run --name todo-postgres -e POSTGRES_PASSWORD=secret_pass -p 5432:5432 -d postgres
```

Connecting to database through psql:
```bash
psql -h localhost -U postgres
```

### Project structure
The structure of project was inspired by succeeding medium post - [link](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1)
