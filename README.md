# simple-grpc-app #
A simple app created to attempt to create frontend & backend apis which communicate using grpc. Having heard much about using it and that it's much faster than using json http requests I decided to have a go at implementing a service using this. Having never used it before it took a little while to understand how to use it but although simplistic, it worked!

It's still a work in progress, started whilst on furlough during the summer to also keep my skills as a microservices engineer fresh & explore something I hadn't had the opportunity to use at work.


## BookMyPlace is a simple webapp for booking properties ##
Trying to think of an idea for a project without having one set for you is tough, make it too challenging and you'll achieve a lot less, but too easy and you might loose interest. I chose to develop an AirBnB-esque (cliche I know :) ) to allow users to search for properties and book them. Again simplicity was important as I'd never used grpc (did I say that already??)

This also made it important to try and keep the models simple as in order to get something working it's much easier to start small, so the 3 main models you will find in the database are User, Property & Booking. A Booking represents an actual booking and therefore has a forgein key relationship with a User & a Property.

This is represented bookings.proto with simialar data structures as well as a UserProperty booking which is composed of a User, Property & Booking so that all required data for a booking can be returned, either when creating a booking or retrieving a booking, although the latter is yet to be implemented.

This is a work in progress, bakanced with juggling my numerous other personal projects, such as learning c++, working on a clothing design/software design project, and also a little project I started attempting to write a 3D perspective game for my nephews in Python.

So how to run it for yourself...


## Running BookMyPlace Locally

In order to run the services you must have docker, docker-compose & golang installed.

Once installed, from the root folder simply enter the following command into the command line interface of your choice, ensuring that Docker is up & running first:

`docker-compose up --build`

Once the containers are all up, n.b. the migrate container should exit but none of the others (This simply spins up and applies any migrations to the database, in this case creating the models and prepopulating the database with some entries)

Then using an api client such as postman, or insomnia (or even a browser for GET requests) you can test the landing page by using the URL http://localhost:8080/

In order to get a list of properties the http://localhost:8080/properties accepts GET requests and will return all properties in the database. (I know this is not ideal but I haven't yet implented pagination - TODO)

You can also add query paramaters, e.g. to search by country you could do:

http://localhost:8080/properties?country=Spain

The resulting json response body would be returned (abbreviated):

```json
[
  {
    "id": 6,
    "doorNumber": "16",
    "address": "Calle Menorca",
    "city": "Madrid",
    "country": "Spain"
  },
  {
    "id": 7,
    "doorNumber": "78",
    "address": "Travessera de Gracia",
    "city": "Barcelona",
    "country": "Spain"
  },
  {
    "id": 8,
    "doorNumber": "19",
    "address": "Calle Rio Seco",
    "city": "Alicante",
    "country": "Spain"
  }...
]
```

You can also making a booking by sending a post request to:

http://localhost:8080/booking

the request should have `Content-Type: application/json` in the header.

An example would request body would be:

```json
{
	"propertyId": 6,
	"userId": 2,
	"startDate": "2020-07-20T13:17:49.5197009Z",
	"endDate": "2020-07-27T13:17:49.5197009Z"
}
```

The response would have http status code **201**, and the example response would look like:

```json
{
  "property": {
    "id": 6,
    "doorNumber": "16",
    "address": "Calle Menorca",
    "city": "Madrid",
    "country": "Spain"
  },
  "user": {
    "id": 2,
    "firstName": "Billy-Bob",
    "surname": "Thornton"
  },
  "startDate": "2020-07-20T13:17:49.519701Z",
  "endDate": "2020-07-27T13:17:49.519701Z",
  "createdAt": "2020-11-14T15:13:21.087881Z"
}
```