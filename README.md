# What is all about

We have two containers on with cassandra and other with golang.

Point is that golang writes in cassandra DB and then retrieves all data from Cassandra DB.

# Containers

We use docker-compose file for seting up. We expose port 9042 for cassandra, port 8898 on which application is running and port 8888 which is port of nginx that does proxy to go aplication.
We use ports and expose due that we can test each image sepperatly.

## Cassandra DB

This uses basic image from docker hub (https://hub.docker.com/_/cassandra/) - we used 3.11 version

## Golang container
Uses basic container https://hub.docker.com/_/golang/ which is wraped then with Dockerfile.
Also installs nginx and does proxy pass to go lang app


# Run and test Testing 

Run it with docker compose(you sould be in folder where you checkout the code):
```bash 
docker-compose up
```

go to your web browser and open http://0.0.0.0:8888

evry time you refresh page you should get new entry.

# Problems

Printing of the table is not sorted. Will dig into it. Probbaly is string related sorting problem.





