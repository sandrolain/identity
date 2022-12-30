# Identity

Identity Service as GoLang Exercise

## ToDo

- [x] Env Config: Password Strength validation
- [ ] Machine APIs: IP validation
- [ ] Machine APIs: Issuer validation
- WebAuthn
  - [ ] WebAuthn: Association
  - [ ] WebAuthn: Login
- [ ] gRPC: mTLS
- [ ] Roles package
- [ ] Admin APIs authorized with roles
- User:
  - [ ] Password reset
  - [ ] Logout
- Admin:
  - [ ] Edit user
  - [ ] Edit machine
  - [ ] List Entities
  - [ ] Create Admin
  - [ ] Edit Admin
  - [ ] Reset machine sessions
  - [ ] List machine sessions
  - [ ] Logout

## Premises

- Zero trust
- Opinionated approach
- Environment variables as Configuration input
- Strong inputs validation
- Required 2FA with TOTP
  - 2FA Secret regenerated until first successful 2FA login
- Authenticated gRPC as communication channel
- Separated gRPC ports for administration and public clients
- Session IDs as authentication Keys for everything
  - For sending 2FA without resend credentials
  - For Users' short session
  - For Users' token association (WebAuthn)
  - For Users' password reset
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
- WebAuthn
- Technologies:
  - JWT for session ID tokens
  - WebAuthn for Users authentication
  - Redis for short persistence and caching
  - MongoDB for long persistence
    - Users
    - Machines keys
