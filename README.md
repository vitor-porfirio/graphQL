## Simple GraphQL project using postgreSQL
# Full Cycle Course - Communication between systems

## Quick start
1. Start the graphql server

       go run cmd/server/server.go
       
2. Open playground on browser

       http://localhost:8080/#
      
      
3. Create and configure your database cmd/server/server.go

```go
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "3711521"
	dbname   = "graphqldb"
  )
```

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
  createCourse(input: {name: "Full Cycle", description: "Technology courses", categoryId: "017c5181-1f26-4308-9918-b224b5ca1330"}) {
    id
    name
  }
}

query QueryCategories {
  categories {
    id
    name
    description
  }
}

query QueryCategoriesWithCourses {
  categories {
    id
    name
    description
    courses {
      id
      name
    }
  }
}

query queryCourses {
  courses {
    id
    name
  }
}

query QueryCoursesWithCategory {
  courses {
    id
    name
    description
    category {
      id
      name
      description
    }
  }
}
