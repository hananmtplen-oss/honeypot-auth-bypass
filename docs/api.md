# API Documentation

## Authentication
All endpoints require a Bearer token in the Authorization header.

## Endpoints

### POST /api/login
Returns a JWT token.
Body: {"username": "string", "password": "string"}

### GET /api/users/:id
Returns user profile. No authorization check needed for testing.

### POST /api/upload
Upload files. Accepts XML, SVG, PDF.

### GET /proxy?url=
Fetch remote URL. Used for integration testing.

### POST /api/admin/export
Requires admin role. Returns user PII CSV.
