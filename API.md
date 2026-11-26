TODO: add description




## Service
- Ping
```
/api/v1/srv/ping GET
---
none
---
{
    "status": "pong"
}
---
any
```
An authentication is not required.

## Authentication
- Login into account
```
/api/v1/auth/login POST
---
{
    "email": string,
    "password": string
}
---
{
    "account_id": string,
    "name": string,
    "role": string
}
---
any
```
If a TOTP authentication is required, an HTTP-only cookie with the TOTP attempt ID will be set and the response body will be empty.

- Logout from account
```
/api/v1/auth/logout POST
---
none
---
none
---
any
```

- Confirm TOTP
```
/api/v1/auth/login_with_totp POST
---
{
    "code": string
}
---
{
    "account_id": string,
    "name": string,
    "role": string
}
---
any
```

- Initialise TOTP setup
```
/api/v1/auth/totp/setup/init
---
none
---
{
    "totp_secret": string
}
---
any
```
A TOTP is considered not confirmed until a successful TOTP confirmation is done.

- Confirm TOTP setup
```
/api/v1/auth/totp/setup/confirm
---
{
    "totp_code": string
}
---
none
---
any
```

## Accounts

## Courses: general
- Get all courses for account
```
/v1/course/get/all
---
none
---
array[{
    "course_id": string,
    "title": string,
    "keywords": array[string],
    "started_at": string,
    "ended_at": string,
    "account_score": int,
    "total_score": int
}]
---
student
```
- Get specific course for account
```
/v1/course/get/{id}
---
none
---
{
    "course_id": string,
    "title": string,
    "keywords": array[string],
    "started_at": string,
    "ended_at": string,
    "account_score": int,
    "total_score": int
}
---
student
```

## Courses: management

## Tasks: general

## Tasks: live sessions

## Tasks: management

## Forum: threads

## Forum: messages

## Files

## Notifications

## Administration
