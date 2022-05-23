# login

#### scene
登录获取token

#### api
POST {apiaddress}/v1/login

#### request param
| field name |  type   |  required  |  description  |  example |
| ---------  | ------  | -------| --------- |------------ |
| phone       | string  |   Y    | phone     |          |
| verifi_code      | string  |   Y    | verification code     |          |


#### response param
| field name |  type   |  required  |  description  |  example |
| ---------  | ------  | -------| --------- |------------ |
| code       | int     |   Y    |  0:success other:fail    |         |
| message    | string  |   N    | error message     |          |
| data       | object  |   N    | user info     |          |
| data.token | string  |   Y    |      |          |

#### example
`
{
    "phone":"evan",
    "verifi_code":"2020-01-01"
}


{
    "code": 0,
    "message": "ok",
    "data":{
        "token":"xxxxxx"
    }
}
`

#### note
