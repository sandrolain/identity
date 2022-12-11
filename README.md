# Identity

Identity Service as GoLang Exercise


## Premises

- Zero trust
- Minimal
- Environment variables as Configuration input
- Strong inputs validation
- Required 2FA 
- Authenticate gRPC as communication channel
- Separated gRPC ports for admin and clients
- Session IDs as authentication Keys for everything
  - For sending 2FA without resend credentials
  - For Users' short session
  - For Machines' long session
- Entities:
  - Admin: Login via dedicated channel
  - User: Short session expire (redis)
  - Machine: Long session expire (mongodb for persistence, redis for cache and fast data access)
- Users and machines have dedicated APIs
- Roles management
- Dedicated Roles:
  - ADMIN
    - can authenticate in admin APIs
    - can manage other users
    - can manage machines and Sessions
- Username must be a valid email
- Keys rotation for User Session JWT sign
- Technologies:
  - JWT for session ID tokens
  - TOTP for 2FA
  - WebAuthn for Users authentication
  - Redis for short persistence and caching
  - MongoDB for long persistence
    - Users
    - Machines keys
