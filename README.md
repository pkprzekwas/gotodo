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

### Usage

##### Initialization:
```bash
go run main.go init 
```
NOTE: Database must be available on localhost:5432

##### Listing
```bash
go run main.go list
```

##### Creating todo:
```bash
go run main.go create --title "example todo" --description "Example of very important task"
```

##### Getting todo:
```bash
go run main.go get --id 3
```

##### Deliting todo:
```bash
go run main.go delete --id 3
```

### Todo:
- reading from configuration file
- hading task status
- filtering by status
