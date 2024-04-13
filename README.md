# My personal go web api template

This is not meant to be a guide on how to build apis, i am primaryly a .net developer, however i recently decided to switch my primary language to go for OSS work. 
So this template will change over the next few months or years, so don't take it as concrete example of how to build anything, it's just a base that i use to build upon

The purpose of this template is to avoid bullshit like clean architecture for apis, keep things simple yet isolated and at the same time easy to understand so if you have complaints read the contribution section and offer a better solution.

## Features

-  Authentication middleware
-  Jwt Service to issue tokens
-  Configuration Service to manage api settings
-  Example repository for authentication
-  Authentication Service based on ED25519 Signatures (passkeys), not compatible with FIDO2 (sort of can be with a bit more work, we don't care, it's for internal use only)

# Planned features

- Email service to send emails and recover accounts
- SMS Service
- MFA Service
- Authorization Service to manage security roles (Currently there is only one role, "user", which is assigned after the signature is verified)
- Logging service

### Contributions

If you wish to make suggestions, point out flaws or tell me how stupid i did it, go ahead in the issue section and provide a better alternative of your making.
