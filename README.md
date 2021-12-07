# go_movies
## _A React js + GraphQL supported with backend in GoLang.

This app is an attempt towards using go lang with graphql data fetch in react front end. As graphql is popular amongt the front end developers. They can ask for additional models in a single request without backend team doing work to accomodate changes. Front end developer have lot of control over the kind of data sent by the backend.

Golang being popular in microservices space and its performant. This app has backend built together with Postgres DB as persistent layer. All configs are in the repo and you need to run the docker command for development box.

## Deploy 
#   Development

Move over to root. Install the dependencies and devDependencies and start the server.
```sh
docker-compose -f docker-compose.dev.yaml up --build
```

This should start few services (in order). The names are indicative as to their work:
- movies_db
- movies_api
- movies_client
- nginx

## Deploy 
#   Production

We use Kind tool for local kubernetes cluster using the standard install docker infrastructure setup and Make commands to run executable commands easier to use. These are target based, dependency chain all packed together. So a single make target should kich start the clusters and running the app.
