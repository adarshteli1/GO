JWT + Casbin Authorization System in Golang

This project is a complete authentication and authorization system built using Golang, Gin, JWT, and Casbin. It demonstrates how modern backend applications secure APIs using token-based authentication and role-based access control (RBAC).

Users can register and log in through public routes, where JWT access and refresh tokens are generated after successful credential verification. Protected routes use custom authentication middleware to validate JWT tokens and extract user claims such as username, email, and role. After authentication, Casbin middleware enforces RBAC policies by checking whether a user's role is allowed to access a specific route and HTTP method according to the rules defined in policy.csv.

For example:

ADMIN users can create and read notes.
USER users can only read notes.

The project demonstrates:

JWT Authentication
Refresh Tokens
Gin Middleware
Protected Routes
Casbin RBAC Authorization
Role-Based API Access
Secure API Design in Golang

This setup reflects how production-grade backend systems handle authentication and authorization for scalable and secure APIs.
