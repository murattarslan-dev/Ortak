# User Management API

Kullanıcı yönetimi için CRUD işlemleri.

## Base URL
```
http://localhost:8080/api/v1
```

> **Not:** Tüm endpoints JWT token gerektirir.  
> Header: `Authorization: Bearer <token>`

## Endpoints

### Kullanıcıları Listele
```http
GET /users
```

Tüm kullanıcıları getirir.

**Response (200 OK):**
```json
{
  "success": true,
  "message": "Users retrieved successfully",
  "data": [
    {
      "id": 1,
      "username": "johndoe",
      "email": "john@example.com"
    },
    {
      "id": 2,
      "username": "janedoe",
      "email": "jane@example.com"
    }
  ]
}
```

---

### Kullanıcı Detayı
```http
GET /users/:id
```

Belirli bir kullanıcının detaylarını getirir.

**Parameters:**
- `id` (path): Kullanıcı ID'si

**Response (200 OK):**
```json
{
  "success": true,
  "message": "User retrieved successfully",
  "data": {
    "id": 1,
    "username": "johndoe",
    "email": "john@example.com"
  }
}
```

**Hata Durumları:**
- `404 Not Found` - Kullanıcı bulunamadı

---

### Kullanıcı Oluştur
```http
POST /users
```

Yeni kullanıcı oluşturur.

**Request Body:**
```json
{
  "username": "newuser",
  "email": "newuser@example.com",
  "password": "password123"
}
```

**Response (201 Created):**
```json
{
  "success": true,
  "message": "User created successfully",
  "data": {
    "id": 3,
    "username": "newuser",
    "email": "newuser@example.com"
  }
}
```

---

### Kullanıcı Güncelle
```http
PUT /users/:id
```

Mevcut kullanıcıyı günceller. Partial update desteklenir.

**Parameters:**
- `id` (path): Kullanıcı ID'si

**Request Body:**
```json
{
  "username": "updateduser",
  "email": "updated@example.com",
  "password": "newpassword123"
}
```

**Response (200 OK):**
```json
{
  "success": true,
  "message": "User updated successfully",
  "data": {
    "id": 1,
    "username": "updateduser",
    "email": "updated@example.com"
  }
}
```

**Hata Durumları:**
- `404 Not Found` - Kullanıcı bulunamadı
- `500 Internal Server Error` - Güncelleme hatası

---

### Kullanıcı Sil
```http
DELETE /users/:id
```

Kullanıcıyı siler.

**Parameters:**
- `id` (path): Kullanıcı ID'si

**Response (200 OK):**
```json
{
  "success": true,
  "message": "User deleted successfully"
}
```

**Hata Durumları:**
- `404 Not Found` - Kullanıcı bulunamadı
- `500 Internal Server Error` - Silme hatası

## Validation Rules

### Create User Request
- `username`: Zorunlu, minimum 3 karakter
- `email`: Zorunlu, geçerli email formatı
- `password`: Zorunlu, minimum 6 karakter

### Update User Request
- `username`: Opsiyonel, minimum 3 karakter
- `email`: Opsiyonel, geçerli email formatı
- `password`: Opsiyonel, minimum 6 karakter

## Örnek Kullanım

### cURL Örnekleri

**Tüm kullanıcıları listele:**
```bash
curl -X GET http://localhost:8080/api/v1/users \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

**Kullanıcı detayı:**
```bash
curl -X GET http://localhost:8080/api/v1/users/1 \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

**Kullanıcı oluştur:**
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "newuser",
    "email": "newuser@example.com",
    "password": "password123"
  }'
```

**Kullanıcı güncelle:**
```bash
curl -X PUT http://localhost:8080/api/v1/users/1 \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "updateduser"
  }'
```

**Kullanıcı sil:**
```bash
curl -X DELETE http://localhost:8080/api/v1/users/1 \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```