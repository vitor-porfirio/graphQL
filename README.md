## Simple GraphQL project using postgreSQL
## Full Cycle Course - Communication between systems

## Quick start
1. Start the graphql server

       go run cmd/server/server.go
       
2. Open playground on browser

       http://localhost:8080/#
      
      
3. Create and configure your database cmd/server/server.go

	host     = "localhost"
	port     = 5432
       user     = "postgres"
       password = "YOUR-PASSWORD"
       dbname   = "YOUR-DATABASE"

4. You can use the following code on GraphQL playground:
     
```graphql
mutation createCategory {
  createCategory(input: {name: "Technology", description: "Technology courses"}) {
    id
    name
    description
  }
}

mutation createCourse {
  createCourse(input: {name: "Full Cycle", description: "Technology courses", categoryId: "1"}) {
    id
    name
  } {
    id
    name
    description
  }
}

query QueryCategories {
  categories {
    id
    name
    description
  }
}
