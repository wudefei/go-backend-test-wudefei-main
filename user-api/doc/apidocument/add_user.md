# add user

#### scene
add single user information.

#### api
POST {apiaddress}/v1/user

#### request header
| field name |  type   |  required  |  description  |  example |
| ---------  | ------  | -------| --------- |------------ |
| token       | string  |   Y    |      |          |

#### request param
| field name |  type   |  required  |  description  |  example |
| ---------  | ------  | -------| --------- |------------ |
| name       | string  |   Y    | user name     |    "evan"      |
| birth      | string  |   N    | user birth     |   "2020-01-01"       |
| address    | string  |   N    | user address     |          |
| description| string  |   N    | user description     |          |

#### response param
| field name |  type   |  required  |  description  |  example |
| ---------  | ------  | -------| --------- |------------ |
| code       | int     |   Y    |  0:success other:fail    |         |
| message    | string  |   N    | error message     |          |


#### example
`
{
    "name":"evan",
    "birth":"2020-01-01",
    "address":"深圳",
    "description":""
}


{
    "code": 0,
    "message": "ok"
}
`

#### note
