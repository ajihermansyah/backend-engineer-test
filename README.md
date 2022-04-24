# Backend-Engineer-Test
a set of APIs to be used by backend-end engineers to develop an application for REST API generate JWT Token, create new user register, claims jwt token by user login, fetching data commodity from site API, get data commodity aggregate, Tech stack : Go, Docker, Gin Framework, JWT Token

## App-Auth
`app-auth` the application REST API for manage new user (register), password, and JWT generate process to csv file as based file database.

### App-Auth Setup
Set up the application App-Auth, you need to have already installed golang version `go1.16` or latest version, and `docker` installed in your local machine.If you already have them in your local machine, set up first app configation environment values in `app-auth/.env`
* `SERVER_PORT` to define which port that the application will listen to 
	* The default port is `8080`
* `FILE_PATH` to define the path to where the file to store user data
* `SECRET` to define which secret to use to generate/parse JWT

### App-Auth Build and Run
For run and build this app, there are two ways to run this after the configuration values for the app is set up in file `app-auth/.env`.

*The first method is to run the application directly in local machine by using :
```
$ cd app-auth
$ go mod tidy
// run 'go get xxx' if needed to pull the needed library/package as manualy installed
$ go run main.go
```

*The other method is to build this app docker image and run a container based on the image. To build the image, execute these in your machine :
```
$ cd app-auth
$ docker build -t <image-name> .
```

then to run the container based on the image, execute the below command after the image is built in docker (make sure that you set the docker port to `SERVER_PORT` value in `app-auth/.env`) :
```
$ docker run -d -p <host port>:<docker port> <image-name>
```
note: 
`-d` means the docker will run detached/in background  
`-p` means the docker will publish and listen to those port


### App-Auth Example
```
$ cd app-auth
$ docker build -t app-auth .
$ docker run -d -p 8080:8080 app-auth
$ curl localhost:8080/ping
```

### App-Auth System Design
``` 
Using Gin Golang Framework

              HTTP Response
                    A
                    |
                    |
HTTP Request ---> Controllers <---> Repository <---> Database
              1                 2                3     
            (Model)          (Model)          (Model)    

1. HTTP Request is parsed and assigned into a context variable which will be handled by Controllers functions.
    a. Controllers will parse the request body and header and assign those values to variables/struct which is taken from the 'model' package that is shared through all levels/layers of functions.
    b. Controllers will also return an HTTP Response after processing the received/parsed data. Whether it's an error/failed or successful process.
    c. Controllers will only handle request parsing, response, auth, and user/client communication related task.
2. Controllers will call repository as logic functions that will do data processing with the parsed values from request in variables/struct from 'model' package as the functions parameters.
    a. repository will only do data processing. All the data needed for the process are either given by Controllers as parameters or fetched by Repostitory from Database
3. Controllers will calls Repository functions to fetch more data that is needed for the process
4. Repository will fetch data from all external sources, in this case a database.
    a. App-Auth is using a csv file for the database store, repostitory functions are built in to translate logic and read or write csv format files.                       
```

The API serve 4 endpoint routes for App-auth.
- GET  [/ping] for checking server healt check 
- POST [/auth/register] for register new user
- GET  [/auth/generate-token] for generate token by phone and password by user registered
- GET  [/auth/claims-token] fot get data user by jwt token


## App-Fetch
`app-fetch` the application REST API for fetching and process data resources

### App-Fetch Setup
Set up the application `App-Fetch`, you need to have already installed golang version `go1.16` or latest version, and `docker` installed in your local machine.If you already have them in your local machine, set up first app configation environment values in `app-fetch/.env`
* `SERVER_PORT` to define which port that the application will listen to 
	* The default port is `1000`
* `FILE_PATH` to define the path to where the file to store user data
* `SECRET` to define which secret to use to generate/parse JWT
* `KEY` to be used when getting currency conversion rate from https://free.currencyconverterapi.com
    * You need to register and verify your email first to get the API Key
* `RATE_CURRENCY_RATIO` to be used when passing param for get currency conversion rate IDR to USD
* `CACHE_EXPIRED` to define how long the cache expired (set in seconds)
    * The number that is set will be counted in `seconds` (ie. `CACHE_EXPIRED=60` means the cache will be clear/reset after time cache expired setup)


### App-Fetch Build and Run
For run and build this app, there are two ways to run this after the configuration values for the app is set up in file `app-fetch/.env`.

**Steps**
*git clone [https://github.com/ajihermansyah/backend-engineer-test.git](https://github.com/ajihermansyah/backend-engineer-test.git)
*The first method is to run the application directly in local machine by using :
```
$ cd app-fetch
$ go mod tidy
// run 'go get xxx' if needed to pull the needed library/package as manualy installed
$ go run main.go
```

*The other method is to build this app docker image and run a container based on the image. To build the image, execute these in your machine :
```
$ cd app-fetch
$ docker build -t <image-name> .
```

then to run the container based on the image, execute the below command after the image is built in docker (make sure that you set the docker port to `SERVER_PORT` value in `app-fetch/.env`) :
```
$ docker run -d -p <host port>:<docker port> <image-name>
```
note: 
`-d` means the docker will run detached/in background  
`-p` means the docker will publish and listen to those port


### App-Fetch Example
```
$ cd app-fetch
$ docker build -t app-fetch .
$ docker run -d -p 1000:1000 app-fetch
$ curl localhost:1000/ping
```

### App-Auth System Design
``` 
Using Gin Golang Framework
Using memory/RAM based cache

                         HTTP Response                          Cache
                               A                                  A
                               |                                  |
                               |                                  V
HTTP Request ---> HTTP Helper ---> Controllers <---> Repository <---> External Data Sources (API site)
              1                2                 3                4          5
            (Model)         (Model)           (Model)          (Model)    (Model)

1. HTTP Request will be validated and authorized by the HTTP Helper before be forwarded to Controllers. Especially the JWT token inside header request.
    a. Failure on validating or authorizing HTTP Request will make HTTP Helper to return an HTTP Response to user/client.
2. HTTP Request is parsed and assigned into a context variable which will be handled by Controllers functions.
    a. Controllers will parse the request body and header and assign those values to variables/struct which is taken from the 'model' package that is shared through all levels/layers of functions.
    b. Controllers will also return an HTTP Response after processing the received/parsed data. Whether it's an error/failed or successful process.
    c. Controllers will only handle request parsing, response, auth, and user/client communication related task.
3. Controllers will call Repository that will do data processing with the parsed values from request in variables/struct from 'model' package as the functions parameters.
    a. Repository functions will only do data processing. All the data needed for the process are either given by Controllers as parameters or fetched by Repostitory from Database.
4. Controllers will call Repository functions to fetch more data that is needed for the process.
    a. Logic in repository will try to get the data from hit api site resource for the first time,and the cache will be set after fetching data first time.
    b. The next request fetch data, logic in repository will try to check data on cache, before cache data expired. if the cache has been expired,
   	   the logic will hit back repository for fetching data to api site.
	c. The cache will be reset/clear after expired time cache setup from environment
5. The repo will fetch data from all external sources, in this case an external HTTP server that stores the data.
    a. App-Feth gets data from external HTTP server.Repository functions are built to send HTTP requests to the server and parse the response
	   for data before returning data to the function that called it.                    
```

The API serve 4 endpoint routes for App-Fetch.
- GET  [/ping] for checking server healt check 
- GET  [/auth/claims-token] fot get data user by jwt token

- GET  [/commodities] for get list data commodity form external source API site
- GET  [/commodities/aggregate] for get lisr of data commodity aggregate

## Directory Structure
```
backend-engineer-test
    |--app-auth
		|config                 	- to initialize environment apps
			|--config.go            - for getting environment config apps variables
		|--controllers              - to store package controllers
			|--token        		- to handle token controllers
				|--token.go         - for handle generate JWT token
			|--user             	- to handle user controllers
				|--user.go      	- for handle register new user and generate password
		|--data-store               - folder data-store for manage base file database
			|--data.csv 			- file for store data user register 
       
		|--models                   - to store package models for object and mysql query
			|--response         	- for define struct response
				|--response.go
			|--token.go             - for define struct token claims 	
			|--user.go              - for define struct user
		|--repository               - folder containing logic for each entity
			|--token
				|--token.go         - file containing various logical sets for entity token
			|--user
				|--user.go	        - file containing various logical sets for entity user
			|--repository.go        - file containing various functions that represent various entities
		|--.env                     - file containing app configuration variables
        |--dockerfile               - file configruation docker image
		|--main.go      
		
	|--app-fetch
		|config                 				- handle for cache data process
			|--cache.go             
		|config                 				- to initialize environment apps
			|--config.go            			- for getting environment config apps variables
		|--controllers             	 			- to store package controllers
			|--commodity        				- to handle commodity controllers
				|--commodity.go     			- for fetch data commodity 
			|--token             				- to handle token controllers
				|--token.go      				- for handle claims JWT token
		|--helper                   			- folder for manage helper functions
			|--auth 			    			- for manage helper authentication header request
				|--authentication_helper.go
			|--http
				|--http_helper.go				- for manage helper json response http
			helper.go 							- for manage general helper (i.e parse string to time)
		|--models                   			- to store package models for object and mysql query
			|--response         				- for define struct response
				|--response.go
			|--commodity.go						- for define struct commodity
			|--token.go                         - for define struct token claims 	
			|--currency.go                      - for define struct currency
		|--repository                           - folder containing logic for each entity
			|--token
				|--token.go        			    - file containing various logical sets for entity token
			|--commodity
				|--commodity.go	        		- file containing various logical sets for entity commodity
			|--currency
				|--currency.go					- file containing various logical sets for entity currency
			|--repository.go        			- file containing various functions that represent various entities
		|--.env                   			  	- file containing app configuration variables
        |--dockerfile               			- file configruation docker image
		|--main.go                           

  
```

## API Documentations
See [API Documentation](https://documenter.getpostman.com/view/6937298/UyrBhvSQ) on how to use it.