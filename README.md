# Welcome to exoplanets

This app is for managing different exoplanets.


## API

Below is a list of API endpoints with their respective input and output.

### Store Exoplanet

#### Endpoint

```
POST
/exoplanets/store
```

#### Input

```json
{
    "name": "Proxima Centauri b",
    "description": "Proxima Centauri b is the closest exoplanet to Earth",
    "distance_from_earth": 4.2,
    "radius": 0.94,
    "mass": 1.27,
    "type": "GasGiant"
}
```

#### Output

200 OK

### Get Exoplanet

#### Endpoint

```
GET
/exoplanets/read/<name>
```

`name`: A string value, e.g. `Proxima Centauri b`

#### Output

```json
{
    "name": "Proxima Centauri b",
    "description": "Proxima Centauri b is the closest exoplanet to Earth",
    "distance_from_earth": 4.2,
    "radius": 0.94,
    "mass": 1.27,
    "type": "GasGiant"
}
```

### Get All Exoplanets

#### Endpoint

```
GET
/exoplanets/read-all
```

#### Output

```json
[
    {
        "name": "Proxima Centauri b",
        "description": "Proxima Centauri b is the closest exoplanet to Earth",
        "distance_from_earth": 4.2,
        "radius": 0.94,
        "mass": 1.27,
        "type": "GasGiant"
    }
]
```

### Update Exoplanet
 
#### Endpoint

```
PUT
/exoplanets/update
```

#### Input

```json
{
    "name": "Proxima Centauri b",
    "description": "Proxima Centauri b is the closest exoplanet to Earth",
    "distance_from_earth": 4.2,
    "radius": 0.94,
    "mass": 1.28,
    "type": "GasGiant"
}
```

#### Output

200 OK
 
### Delete Exoplanet

#### Endpoint

```
DELETE
/exoplanets/delete/<name>
```

`name`: A string value, e.g. `Proxima Centauri b`

#### Output

200 OK

### Fuel Estimation

#### Endpoint

```
GET
/exoplanets/fuel-estimation/<name>/<crew-capacity>
```

`name`: A string value, e.g. `Proxima Centauri b`
`crew-capacity`: A string value, e.g. `2`

#### Output

26.233165055999997


## Useful commands

In "exoplanets\cmd\server" directory, execute below commands

## go build will generate binary with name server
```
go build
```
## Then execute the binary
./server
```