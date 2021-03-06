## Stacks
This project was build using React, Go and MongoDB <br />
<img src="https://www.vectorlogo.zone/logos/mongodb/mongodb-ar21.svg"  width="210px" height="135px">
<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/a/a7/React-icon.svg/512px-React-icon.svg.png"  width="200px" height="140px">
<img src="https://camo.githubusercontent.com/98ed65187a84ecf897273d9fa18118ce45845057/68747470733a2f2f7261772e6769746875622e636f6d2f676f6c616e672d73616d706c65732f676f706865722d766563746f722f6d61737465722f676f706865722e706e67" width="10%" height="5%">


## Before running the server, install the dependencies

### Server

`go get go.mongodb.org/mongo-driver`

<br />

`go get -u github.com/gorilla/mux`

### Client

`npm install`

## Available Scripts

In the project directory, you can run:

### `yarn start`

yarn start runs the client app in the development mode.<br />
Open [http://localhost:3000](http://localhost:3000) to view it in the browser.

The page will reload if you make edits.<br />
You will also see any lint errors in the console.

### `go run main.go`

Runs the server app in the development mode.<br />
Use [http://localhost:3333](http://localhost:333) to test it with postman or insomnia.

After any changes made, it`s needed to re-run the server.


## Api Features
Create:  [http://localhost:3333/api/planet](http://localhost:333/api/planet) method: POST

> Saves the data provided into mongo db <br />
> JSON { <br />
>   name: (Planet Name) <br />
>   weather: (Planet Weather)  <br />
>   terrain: (Planet Terrain Type) <br />
> } 


Index:  [http://localhost:3333/api/planet](http://localhost:333/api/planet) method: GET

> This route indexes the data inside mongo db <br />
> If neither query or id is provided [http://localhost:3333/api/planet](http://localhost:3333/api/planet) it will index all the data avaible <br />
> If filter query is provided [http://localhost:3333/api/planet?name=PlanetName](http://localhost:3333/api/planet?name=PlanetName) it will index all the data that matches the filter <br />
> If id is provided [http://localhost:3333/api/planet/{id}](http://localhost:3333/api/planet/{id}) it will index only the data that matches the id <br />


Delete:  [http://localhost:3333/api/planet/del/{id}](http://localhost:3333/api/planet/del/{id}) method: DELETE

> This route only works with id, so its necessary to provide an id to delete the data[http://localhost:3333/api/planet/{id}](http://localhost:3333/api/planet/{id}) 

## Client Features (Client language: PT-BR )

About Page: 

> Introtudory page, that tells the api features availables and Developer name <br />
> Contain 2 links, Register Page(Create) and List Planets(Index)

Register Page:  [http://localhost:3000/register](http://localhost:3000/register) method: POST

> Saves the data provided into mongo db <br />
>   Name: (Planet Name) <br />
>   Weather: (Planet Weather)  <br />
>   Terrain: (Planet Terrain Type)  <br />


Planet List Page:  [http://localhost:3000/planet](http://localhost:3000/planet) method: GET

> This page indexes all the data inside mongo db  <br />
> The planet modal has a delete button, so its possible to delete any planet displayed

## License (MIT)

<img src="https://data.gopher.se/gopher/viking-gopher.svg" align="right" width="30%" height="300">

Copyright (c) 2014-2020 [Peter Hellberg](https://c7.se)

> Permission is hereby granted, free of charge, to any person obtaining
> a copy of this software and associated documentation files (the "Software"),
> to deal in the Software without restriction, including without limitation
> the rights to use, copy, modify, merge, publish, distribute, sublicense,
> and/or sell copies of the Software, and to permit persons to whom the
> Software is furnished to do so, subject to the following conditions:
>
> The above copyright notice and this permission notice shall be included
> in all copies or substantial portions of the Software.
>
> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
> EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
> OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
> IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
> DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
> TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
> OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.