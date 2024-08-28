# tucows-challenge
This is a Coding challenge to show up a glimpse of my Technical Skills as a Software Engineer for Tucows (Wavelo) 

## Challenge's description
This small project is split in two parts. We need to create RESTful API (WEB Server) and a CLI (Client side) that intact with the API.
There's not a mandatory topic so feel free to choose any topic. 

As the position use heavily Golang, we recommend to use Go as a primary language to develop the challenge. 

### Functional Requirement
- Multiple Endpoints (aka GET, POST, PUT, etc)
- CRUD interaction with a Database (Postgres)
- JWT Authentication (aka Token)
- Concurrency

### Non-Functional Requirement
- Self-Certificate (aka HTTPS)
- Docker Containers
- Container Orchestration (aka k8s, docker-compose)
- Clean Code
- Documentation

_______

## Brainstorming (You can skip this section)
Now let's Think and Build this Challenge. :) 

As mentioned, First we need to think what project can be a good idea to satisfy all the Requirements. That's why I proceed to 'Brainstorm' and see the Trade-off.  
I believe it will be nice to build a 'Coffee Shop' system, because I love Coffee. <br>
![Task Requirements](https://i.postimg.cc/dVWhdW9n/Requirements-and-Possible-Ideas.jpg)

Later, I moved to visualize how the API will interact with the DB and Worker Pool. And which methods we would like to expose to our Client (CLI). <br>
![Task Requirements](https://i.postimg.cc/xdccPKxQ/Functions-and-Infra.jpg)


Here we can see a Diagram how the Data is inserted and handle through the system. <br>
![Task Requirements](https://i.postimg.cc/pd4ycp4F/Diagram.jpg)


_______

## Develop

To build the systems we need some tools, as well if you want to run it you will have to install:
- Golang https://go.dev/doc/install
- Docker https://docs.docker.com/engine/install/
- Postgres https://www.postgresql.org/download/

Then feel free to Clone this repo. You will notice we have to projects/folder:
- api 
- client

We need to pull and install dependencies, so run the following command inside the main repo and each project (api and client):
```go mod tidy```

Now you have two chooses on how to build and run the API Server. 

### Run API - Using Docker compose

Docker compose will read the Dockerfile for the API and Build a container and Bridge to connect with the Database. We just need to run: 
```docker-compose up --build```

It will start running and listening in port 8080 
```https://localhost:8080/tucows-coffee/``` 


### Run API - Manually
As the API Server depends on the DB, we need the Database running so let's just start postgres by running the following docker command:
```
docker run \
--rm --name postgres \
-e POSTGRES_USER=user \
-e POSTGRES_PASSWORD=password \
-e POSTGRES_DB=mydb \
-p 5432:5432 \
-d postgres:latest
```

Then we just need to move in to the API and Run the main file. 
```
cd tucows-challenge/api/cmd
go run *.go
```

#### Terminal Output
Here we can see the Endpoints
![API Server - Running and Listening](https://i.postimg.cc/xjHgRmhL/Screenshot-2024-08-28-at-14-52-15.png)


### Run CLI - Manually
Let's now run our Interface to interact and request data from the API. So let's move in to the CLI and Run the main file.
```
cd tucows-challenge/client
go run *.go
```

#### Terminal Output
We will see a Console, we will have to interact by input the options showed in the Menu. 
![API Server - Running and Listening](https://i.postimg.cc/FsCvVw2f/Screenshot-2024-08-28-at-14-04-32.png)




## How to use the App (It's super User-Friendly but just in case...)
Let's talk about it in Reverse ↓

```3. Exit```<br>
It's simple, it will close the CLI (system). 

```2. Show Products Menu```<br>
This options is Public and don't need Auth so you can Fetch our Products and check their Prices.


```1. Login as Employee``` <br>
The CLI actually needs Authentication in order to interact with Orders, Wait... but you don't have an Account with us. Luckily for you, you can *Log in* by sending the same value for username and password. Ex: admin / admin

After you log in as an Employee, you will find a *Friendly Menu* with options to handle 'Orders'


![Employee Menu](https://i.postimg.cc/xT8D8tc9/Screenshot-2024-08-28-at-14-17-30.png)

Order can modify but as soon as they got *Confirmed* will be sent to the 'Kitchen' and a Worker will prepare it. 

## Kitchen? What is this? 

Kitchen is a [Thread Pool](https://en.wikipedia.org/wiki/Thread_pool) and also known Worker Pool in Golang. Meaning that we have some Workers taking Order and Preparing them so the client can enjoy their Coffee.
And to show this pattern, We Integrated an "Open Kitchen" so you see all the workers preparing the Orders. 

In the next image check how we trigger all the pending (Pre Orders) and send them to the Kitchen. Then every order is handle by a different worker.


![Kitchen Worker Pool](https://i.postimg.cc/MZr3Pxp8/Screenshot-2024-08-28-at-14-31-05.png)




## Thank you and Enjoy your Coffee!
![Kitchen Gif](https://cdn.dribbble.com/users/939968/screenshots/2362151/chefs-cooking.gif)

