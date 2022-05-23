# query user

#### scene
query single user information.

#### api
GET {apiaddress}/v1/user

#### request header
| field name |  type   |  required  |  description  |  example |
| ---------  | ------  | -------| --------- |------------ |
| token       | string  |   Y    |      |          |

#### request param
| field name |  type   |  required  |  description  |  example |
| ---------  | ------  | -------| --------- |------------ |
| id       | int  |   Y    | user ID     |    1      |
| name       | string  |   N    | user name     |    "evan"      |


#### response param
| field name |  type   |  required  |  description  |  example |
| ---------  | ------  | -------| --------- |------------ |
| code       | int     |   Y    |  0:success other:fail    |         |
| message    | string  |   N    | error message     |          |
| data       | object  |   N    | user info     |          |
| data.id       | int  |   Y    | user id     |          |
| data.name       | string  |   Y    | user name     |          |
| data.birth      | string  |   Y    | user birth     |          |
| data.address    | string  |   Y    | user address     |          |
| data.description| string  |   Y    | user description     |          |
| data.created_at | string  |   Y    | user created_at     |          |
| data.updated_at | string  |   Y    | user updated_at     |          |


#### example
`
GET {apiaddress}/v1/user?id=1&name=evan


{
    "code": 0,
    "message": "ok",
    "data":{
        "id":1,
        "name":"evan",
        "birth":"2020-01-01",
        "address":"深圳",
        "description":"",
        "created_at":"2022-05-01 10:00:00",
        "updated_at":"2022-05-01 10:00:00"
    }
}
`

#### note
