Project:
    name:
        Credentials Authentication
    author:
        Michael Machado
    language:
        Go
    libs:
	    github.com/gin-gonic/gin v1.10.0
	    github.com/golang-jwt/jwt v3.2.2+incompatible
	    github.com/lib/pq v1.10.9
	    golang.org/x/crypto v0.23.0
    tools:
        docker-composer
        make
    infra:
        Redis 
    goal:
        create a service for authentication,
        and apply this service for anyone other project as a middleware.
    knowledge:
        learning about Go in general
        working with pointers
        API in Go
        Go patterns
        Makefile
        docker-compose

