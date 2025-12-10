# Task Management API

Görev yönetimi için CRUD işlemleri.

## Base URL
```
http://localhost:8080/api/v1
```

> **Not:** Tüm endpoints JWT token gerektirir.  
> Header: `Authorization: Bearer <token>`

## Endpoints

### Görevleri Listele
```http
GET /tasks
```

Tüm görevleri getirir.

**Response (200 OK):**
```json
{
  "success": true,
  "message": "Tasks retrieved successfully",
  "data": [
    {
      "id": 1,
      "title": "Setup API",
      "description": "Create REST API endpoints",
      "status": "in_progress",
      "assignee_id": 1,
      "team_id": 1,
      "tags": ["backend", "api"],
      "comment_count": 3
    },
    {
      "id": 2,
      "title": "Design UI",
      "description": "Create user interface mockups",
      "status": "todo",
      "assignee_id": 2,
      "team_id": 2,
      "tags": ["frontend", "design"],
      "comment_count": 1
    }
  ]
}
```

---

### Görev Detayı
```http
GET /tasks/:id
```

Belirli bir görevin detaylarını getirir. Görev detayında tüm yorumlar da döner.

**Parameters:**
- `id` (path): Görev ID'si

**Response (200 OK):**
```json
{
  "success": true,
  "message": "Task retrieved successfully",
  "data": {
    "id": 1,
    "title": "Setup API",
    "description": "Create REST API endpoints",
    "status": "in_progress",
    "assignee_id": 1,
    "team_id": 1,
    "tags": ["backend", "api"],
    "comment_count": 2,
    "comments": [
      {
        "id": 1,
        "task_id": 1,
        "comment": "API endpoints are looking good!",
        "created_at": "2024-12-20T10:30:00Z",
        "user": {
          "id": 2,
          "username": "john_doe",
          "email": "john@example.com"
        }
      },
      {
        "id": 2,
        "task_id": 1,
        "comment": "Need to add validation",
        "created_at": "2024-12-20T11:15:00Z",
        "user": {
          "id": 3,
          "username": "jane_smith",
          "email": "jane@example.com"
        }
      }
    ]
  }
}
```

**Hata Durumları:**
- `404 Not Found` - Görev bulunamadı

---

### Görev Oluştur
```http
POST /tasks
```

Yeni görev oluşturur.

**Request Body:**
```json
{
  "title": "Write Tests",
  "description": "Write unit tests for API endpoints",
  "assignee_id": 1,
  "team_id": 1,
  "tags": ["testing", "backend"]
}
```

**Response (201 Created):**
```json
{
  "success": true,
  "message": "Task created successfully",
  "data": {
    "id": 3,
    "title": "Write Tests",
    "description": "Write unit tests for API endpoints",
    "status": "todo",
    "assignee_id": 1,
    "team_id": 1,
    "tags": ["testing", "backend"],
    "comment_count": 0
  }
}
```

---

### Görev Güncelle
```http
PUT /tasks/:id
```

Mevcut görevi günceller.

**Parameters:**
- `id` (path): Görev ID'si

**Request Body:**
```json
{
  "title": "Updated Task Title",
  "description": "Updated task description",
  "status": "in_progress",
  "assignee_id": 2,
  "tags": ["backend", "api", "urgent"]
}
```

**Response (200 OK):**
```json
{
  "success": true,
  "message": "Task updated successfully",
  "data": {
    "id": 1,
    "title": "Updated Task Title",
    "description": "Updated task description",
    "status": "in_progress",
    "assignee_id": 2,
    "team_id": 1,
    "tags": ["backend", "api", "urgent"],
    "comment_count": 3
  }
}
```

**Hata Durumları:**
- `404 Not Found` - Görev bulunamadı
- `500 Internal Server Error` - Güncelleme hatası

---

### Görev Durumu Güncelle
```http
PUT /tasks/:id/status
```

Sadece görev durumunu günceller.

**Parameters:**
- `id` (path): Görev ID'si

**Request Body:**
```json
{
  "status": "in_progress"
}
```

**Response (200 OK):**
```json
{
  "success": true,
  "message": "Task status updated successfully",
  "data": {
    "id": 1,
    "title": "Setup API",
    "description": "Create REST API endpoints",
    "status": "in_progress",
    "assignee_id": 1,
    "team_id": 1,
    "tags": ["backend", "api"],
    "comment_count": 3
  }
}
```

**Hata Durumları:**
- `400 Bad Request` - Geçersiz status değeri
- `404 Not Found` - Görev bulunamadı
- `500 Internal Server Error` - Güncelleme hatası

---

### Görev Yorumu Ekle
```http
POST /tasks/:id/comments
```

Göreve yorum ekler.

**Parameters:**
- `id` (path): Görev ID'si

**Request Body:**
```json
{
  "comment": "Bu görev için API dokümantasyonu da hazırlanmalı"
}
```

**Response (201 Created):**
```json
{
  "success": true,
  "message": "Comment added successfully",
  "data": {
    "id": 3,
    "task_id": 1,
    "comment": "Bu görev için API dokümantasyonu da hazırlanmalı",
    "created_at": "2024-12-20T12:00:00Z",
    "user": {
      "id": 1,
      "username": "admin",
      "email": "admin@example.com"
    }
  }
}
```

**Hata Durumları:**
- `400 Bad Request` - Geçersiz istek formatı
- `401 Unauthorized` - Kimlik doğrulama hatası
- `404 Not Found` - Görev bulunamadı
- `500 Internal Server Error` - Yorum ekleme hatası

---

### Görev Sil
```http
DELETE /tasks/:id
```

Görevi siler.

**Parameters:**
- `id` (path): Görev ID'si

**Response (200 OK):**
```json
{
  "success": true,
  "message": "Task deleted successfully"
}
```

**Hata Durumları:**
- `404 Not Found` - Görev bulunamadı
- `500 Internal Server Error` - Silme hatası

## Görev Durumları

| Status | Açıklama |
|--------|----------|
| `todo` | Yapılacak - Henüz başlanmamış |
| `in_progress` | Devam Ediyor - Üzerinde çalışılıyor |
| `done` | Tamamlandı - Görev bitirildi |

## Validation Rules

### Create Task Request
- `title`: Zorunlu, minimum 3 karakter
- `description`: Opsiyonel
- `assignee_id`: Zorunlu, geçerli kullanıcı ID'si
- `team_id`: Zorunlu, geçerli takım ID'si
- `tags`: Opsiyonel, string array (etiketler)

### Update Task Request
- `title`: Opsiyonel, minimum 3 karakter
- `description`: Opsiyonel
- `status`: Opsiyonel, geçerli status değeri
- `assignee_id`: Opsiyonel, geçerli kullanıcı ID'si
- `tags`: Opsiyonel, string array (etiketler)

### Update Task Status Request
- `status`: Zorunlu, geçerli status değeri (todo, in_progress, done)

### Add Comment Request
- `comment`: Zorunlu, minimum 1 karakter

## Örnek Kullanım

### cURL Örnekleri

**Tüm görevleri listele:**
```bash
curl -X GET http://localhost:8080/api/v1/tasks \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

**Görev detayı:**
```bash
curl -X GET http://localhost:8080/api/v1/tasks/1 \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

**Görev oluştur:**
```bash
curl -X POST http://localhost:8080/api/v1/tasks \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "New Task",
    "description": "Task description",
    "assignee_id": 1,
    "team_id": 1,
    "tags": ["backend", "api"]
  }'
```

**Görev güncelle:**
```bash
curl -X PUT http://localhost:8080/api/v1/tasks/1 \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Updated Task",
    "status": "done",
    "tags": ["backend", "completed"]
  }'
```

**Görev durumu güncelle:**
```bash
curl -X PUT http://localhost:8080/api/v1/tasks/1/status \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -H "Content-Type: application/json" \
  -d '{
    "status": "done"
  }'
```

**Görev yorumu ekle:**
```bash
curl -X POST http://localhost:8080/api/v1/tasks/1/comments \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -H "Content-Type: application/json" \
  -d '{
    "comment": "Great progress on this task!"
  }'
```

**Görev sil:**
```bash
curl -X DELETE http://localhost:8080/api/v1/tasks/1 \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

## İleri Seviye Özellikler

> **Not:** Bu özellikler gelecek versiyonlarda eklenecektir.

- **Görev Atama** - Kullanıcıya görev atama/değiştirme
- **Status Workflow** - Durum geçiş kuralları
- **Görev Yorumları** - Görev üzerinde yorum sistemi
- **Dosya Ekleri** - Görevlere dosya ekleme
- **Deadline Yönetimi** - Son teslim tarihi
- **Görev Etiketleri** - Kategorilendirme için etiketler