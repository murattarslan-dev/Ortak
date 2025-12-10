# Authentication API

Kullanıcı kimlik doğrulama işlemleri için API endpoints.

## Base URL
```
http://localhost:8080/api/v1
```

## Endpoints

### Kayıt Ol
```http
POST /register
```

Yeni kullanıcı kaydı oluşturur.

**Request Body:**
```json
{
  "username": "johndoe",
  "email": "john@example.com",
  "password": "password123"
}
```

**Response (201 Created):**
```json
{
  "success": true,
  "message": "User registered successfully",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": 1,
      "username": "johndoe",
      "email": "john@example.com"
    }
  }
}
```

**Hata Durumları:**
- `400 Bad Request` - Geçersiz veri formatı
- `500 Internal Server Error` - Sunucu hatası

---

### Giriş Yap
```http
POST /login
```

Mevcut kullanıcı girişi yapar.

**Request Body:**
```json
{
  "email": "john@example.com",
  "password": "password123"
}
```

**Response (200 OK):**
```json
{
  "success": true,
  "message": "Login successful",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": 1,
      "username": "johndoe",
      "email": "john@example.com"
    }
  }
}
```

**Hata Durumları:**
- `400 Bad Request` - Geçersiz veri formatı
- `401 Unauthorized` - Hatalı kimlik bilgileri

---

### Çıkış Yap
```http
DELETE /logout
```

Kullanıcı oturumunu sonlandırır.

**Headers:**
```
Authorization: Bearer <token>
```

**Response (200 OK):**
```json
{
  "success": true,
  "message": "Logged out successfully"
}
```

**Hata Durumları:**
- `400 Bad Request` - Authorization header gerekli
- `401 Unauthorized` - Geçersiz token
- `500 Internal Server Error` - Sunucu hatası

## Validation Rules

### Register Request
- `username`: Zorunlu, minimum 3 karakter
- `email`: Zorunlu, geçerli email formatı
- `password`: Zorunlu, minimum 6 karakter

### Login Request
- `email`: Zorunlu, geçerli email formatı
- `password`: Zorunlu

## Örnek Kullanım

### cURL Örnekleri

**Kayıt:**
```bash
curl -X POST http://localhost:8080/api/v1/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "email": "test@example.com",
    "password": "password123"
  }'
```

**Giriş:**
```bash
curl -X POST http://localhost:8080/api/v1/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "password123"
  }'
```

**Çıkış:**
```bash
curl -X DELETE http://localhost:8080/api/v1/logout \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```