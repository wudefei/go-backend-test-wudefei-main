# sms

#### scene
发送短信验证码

#### api
POST {apiaddress}/v1/sms

#### request param
| field name |  type   |  required  |  description  |  example |
| ---------  | ------  | -------| --------- |------------ |
| phone       | string  |   Y    | phone     |          |


#### response param
| field name |  type   |  required  |  description  |  example |
| ---------  | ------  | -------| --------- |------------ |
| code       | int     |   Y    |  0:success other:fail    |         |
| message    | string  |   N    | error message     |          |

#### example
`
{
    "phone":"XXXX"
}


{
    "code": 0,
    "message": "ok"
}
`

#### note
