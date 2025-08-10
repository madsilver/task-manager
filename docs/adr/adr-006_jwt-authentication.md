# ADR 006: JWT Authentication Strategy

## Status

Accepted

## Context

The API requires an authentication mechanism that is stateless, portable across services, and compatible with API Gateway integration. This authentication method should allow clients to receive a signed token after providing valid credentials and use that token to authenticate future requests.

## Decision

We chose **JWT (JSON Web Token)** as the authentication strategy for the following reasons:

* **Stateless**: No need to persist sessions on the server
* **Portable**: Token can be used across different microservices or external systems
* **Standardized**: Widely supported and well-documented
* **Flexible claims**: Allows encoding of user ID, roles, and permissions

### Token Generation

Upon successful login, the server generates a JWT using the following claim structure:

```json
{
  "sub": "<user-id>",
  "email": "user@example.com",
  "roles": ["admin"],
  "exp": 1691234567
}
```

This token is signed using HMAC with a secret key from the environment:

```env
JWT_SECRET=your_super_secret
```

### Token Signing Code

```go
token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    "sub": user.ID,
    "email": user.Email,
    "roles": user.Roles,
    "exp": time.Now().Add(24 * time.Hour).Unix(),
})
tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
```

## Consequences

* No server-side session management required
* Requires clients to securely store and transmit the JWT
* JWT expiration must be enforced during validation
* Only the signature is verified by the API (validation is skipped if an API Gateway handles it)

## Alternatives Considered

* **Session cookies**: Not suitable for stateless APIs or mobile clients
* **OAuth2**: Overkill for this use case and introduces additional complexity
* **API Key per request**: Less flexible and secure than JWT with claims

## Future Considerations

* Implement token revocation via blacklist (e.g., Redis)
* Support refresh tokens for long-lived sessions
* Consider asymmetric signing (e.g., RS256) if verifying tokens externally

## References

* [https://jwt.io](https://jwt.io)
* [https://auth0.com/docs/secure/tokens/json-web-tokens](https://auth0.com/docs/secure/tokens/json-web-tokens)

---

**Date:** 2025-08-05  
**Author:** Rodrigo Prata  
**Status:** Recorded  
