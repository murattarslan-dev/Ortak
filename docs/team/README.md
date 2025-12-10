# Team Management API

Takım yönetimi için CRUD işlemleri.

## Base URL
```
http://localhost:8080/api/v1
```

> **Not:** Tüm endpoints JWT token gerektirir.  
> Header: `Authorization: Bearer <token>`

## Endpoints

### Takımları Listele
```http
GET /teams
```

Tüm takımları getirir.

**Response (200 OK):**
```json
{
  "success": true,
  "message": "Teams retrieved successfully",
  "data": [
    {
      "id": 1,
      "name": "Development Team",
      "description": "Backend development team",
      "owner_id": 1
    },
    {
      "id": 2,
      "name": "Design Team",
      "description": "UI/UX design team",
      "owner_id": 2
    }
  ]
}
```

---

### Takım Detayı
```http
GET /teams/:id
```

Belirli bir takımın detaylarını getirir.

**Parameters:**
- `id` (path): Takım ID'si

**Response (200 OK):**
```json
{
  "success": true,
  "message": "Team retrieved successfully",
  "data": {
    "id": 1,
    "name": "Development Team",
    "description": "Backend development team",
    "owner_id": 1,
    "members": [
      {
        "id": 1,
        "user_id": 2,
        "team_id": 1,
        "role": "developer"
      },
      {
        "id": 2,
        "user_id": 3,
        "team_id": 1,
        "role": "admin"
      }
    ]
  }
}
```

**Hata Durumları:**
- `404 Not Found` - Takım bulunamadı

---

### Takım Oluştur
```http
POST /teams
```

Yeni takım oluşturur. Oluşturan kullanıcı otomatik olarak owner olur ve takıma "owner" rolü ile üye olarak eklenir.

**Request Body:**
```json
{
  "name": "QA Team",
  "description": "Quality Assurance team"
}
```

**Response (201 Created):**
```json
{
  "success": true,
  "message": "Team created successfully",
  "data": {
    "id": 3,
    "name": "QA Team",
    "description": "Quality Assurance team",
    "owner_id": 1
  }
}
```

> **Not:** Takım oluşturulduktan sonra, oluşturan kullanıcı otomatik olarak "owner" rolü ile takıma üye olarak eklenir.

---

### Takım Güncelle
```http
PUT /teams/:id
```

Mevcut takımı günceller. Sadece takım sahibi güncelleyebilir.

**Parameters:**
- `id` (path): Takım ID'si

**Request Body:**
```json
{
  "name": "Updated Team Name",
  "description": "Updated team description"
}
```

**Response (200 OK):**
```json
{
  "success": true,
  "message": "Team updated successfully",
  "data": {
    "id": 1,
    "name": "Updated Team Name",
    "description": "Updated team description",
    "owner_id": 1
  }
}
```

**Hata Durumları:**
- `403 Forbidden` - Sadece takım sahibi güncelleyebilir
- `404 Not Found` - Takım bulunamadı

---

### Takım Sil
```http
DELETE /teams/:id
```

Takımı siler. Sadece takım sahibi silebilir.

**Parameters:**
- `id` (path): Takım ID'si

**Response (200 OK):**
```json
{
  "success": true,
  "message": "Team deleted successfully"
}
```

**Hata Durumları:**
- `403 Forbidden` - Sadece takım sahibi silebilir
- `404 Not Found` - Takım bulunamadı

## Validation Rules

### Create Team Request
- `name`: Zorunlu, minimum 3 karakter
- `description`: Opsiyonel

### Update Team Request
- `name`: Opsiyonel, minimum 3 karakter
- `description`: Opsiyonel

## Örnek Kullanım

### cURL Örnekleri

**Tüm takımları listele:**
```bash
curl -X GET http://localhost:8080/api/v1/teams \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

**Takım detayı:**
```bash
curl -X GET http://localhost:8080/api/v1/teams/1 \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

**Takım oluştur:**
```bash
curl -X POST http://localhost:8080/api/v1/teams \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "New Team",
    "description": "Team description"
  }'
```

**Takım güncelle:**
```bash
curl -X PUT http://localhost:8080/api/v1/teams/1 \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Updated Team Name"
  }'
```

**Takım sil:**
```bash
curl -X DELETE http://localhost:8080/api/v1/teams/1 \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

## Takım Üyeleri Yönetimi

### Takıma Üye Ekleme
```http
POST /teams/:id/members
```

Takıma yeni üye ekler. Sadece takım sahibi ekleyebilir.

**Parameters:**
- `id` (path): Takım ID'si

**Request Body:**
```json
{
  "user_id": 2,
  "role": "developer"
}
```

**Response (201 Created):**
```json
{
  "success": true,
  "message": "Team member added successfully",
  "data": {
    "id": 1,
    "user_id": 2,
    "team_id": 1,
    "role": "developer"
  }
}
```

**Hata Durumları:**
- `403 Forbidden` - Sadece takım sahibi üye ekleyebilir
- `404 Not Found` - Takım bulunamadı

---

### Takımdan Üye Çıkarma
```http
DELETE /teams/:id/members/:userId
```

Takımdan üye çıkarır. Sadece takım sahibi çıkarabilir.

**Parameters:**
- `id` (path): Takım ID'si
- `userId` (path): Çıkarılacak kullanıcı ID'si

**Response (200 OK):**
```json
{
  "success": true,
  "message": "Team member removed successfully"
}
```

**Hata Durumları:**
- `403 Forbidden` - Sadece takım sahibi üye çıkarabilir
- `404 Not Found` - Takım veya üye bulunamadı

---

### Üye Rolü Güncelleme
```http
PUT /teams/:id/members/:userId/role
```

Takım üyesinin rolünü günceller. Sadece takım sahibi güncelleyebilir.

**Parameters:**
- `id` (path): Takım ID'si
- `userId` (path): Güncellenecek kullanıcı ID'si

**Request Body:**
```json
{
  "role": "admin"
}
```

**Response (200 OK):**
```json
{
  "success": true,
  "message": "Member role updated successfully",
  "data": {
    "id": 1,
    "user_id": 2,
    "team_id": 1,
    "role": "admin"
  }
}
```

**Hata Durumları:**
- `403 Forbidden` - Sadece takım sahibi rol güncelleyebilir
- `404 Not Found` - Takım veya üye bulunamadı

## Üye Rolleri

| Rol | Açıklama |
|-----|----------|
| `owner` | Takım sahibi - Tüm yetkiler |
| `admin` | Yönetici - Üye yönetimi hariç tüm yetkiler |
| `developer` | Geliştirici - Görev oluşturma ve düzenleme |
| `designer` | Tasarımcı - Tasarım görevleri |
| `viewer` | Görüntüleyici - Sadece okuma yetkisi |

## Örnek Kullanım

### cURL Örnekleri

**Takıma üye ekleme:**
```bash
curl -X POST http://localhost:8080/api/v1/teams/1/members \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": 2,
    "role": "developer"
  }'
```

**Üye rolü güncelleme:**
```bash
curl -X PUT http://localhost:8080/api/v1/teams/1/members/2/role \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -H "Content-Type: application/json" \
  -d '{
    "role": "admin"
  }'
```

**Üye çıkarma:**
```bash
curl -X DELETE http://localhost:8080/api/v1/teams/1/members/2 \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

## İleri Seviye Özellikler

> **Not:** Bu özellikler gelecek versiyonlarda eklenecektir.

- **Takım İstatistikleri** - Görev tamamlama oranları
- **Takım Avatarı** - Profil resmi upload
- **Rol Bazlı Yetkiler** - Detaylı yetki sistemi
- **Üye Davetleri** - Email ile üye davet sistemi