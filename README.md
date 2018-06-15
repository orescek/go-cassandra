# What is all about

We have two containers on with cassandra and other with golang.

Point is that golang writes in cassandra DB and then retrieves all data from Cassandra DB.

#Containers

We use docker-compose file for seting up. We expose all ports for cassandra(probably we need just 9042), and port on which application is running(hardcoded on 8898).


## Cassandra DB

This uses basic image from docker hub (https://hub.docker.com/_/cassandra/) - we used 3.11 version

## Golang container
Uses basic container https://hub.docker.com/_/golang/ which is wraped then with Dockerfile.


#Testing

Run it with docker compose(you sould be in folder where you checkout the code):
docker-compose up

go to your web browser and open http://0.0.0.0:8898/test

evry time you refresh page you should get new entry.

#Problems

Printing of the table is not sorted. Will dig into it.





