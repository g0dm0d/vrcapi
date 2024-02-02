# Setup 2fa auth

Go to [profile](https://vrchat.com/home/profile), then scroll to two-factor authentication and click enable. 

Now in step 2 download the QR code and get the text stored in it for example using [this site](https://qrcoderaptor.com)

![step2](https://raw.githubusercontent.com/g0dm0d/vrcapi/master/examples/2fa_login/src/step2.png)

then you will get your 2fa secret by which TOTP is generated.

![secret](https://raw.githubusercontent.com/g0dm0d/vrcapi/master/examples/2fa_login/src/qr_decode.png)

It needs to be used here.

```go
vrc.Auth("username", "password", "JYYGCWLJLJ2VAVBYKF3HIQLOLJMVIMSV")
```

The bot automatically logs into the account and automatically confirms login.