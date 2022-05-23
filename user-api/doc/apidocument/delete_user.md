# user Interface

#### scene
delete single user information.

#### api
DELETE {apiaddress}/v1/user

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


#### example
`
{
    "id": 1,
    "name":"evan"
}


{
    "code": 0,
    "message": "ok"
}
`

#### note
