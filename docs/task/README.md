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
      "team_id": 1
    },
    {
      "id": 2,
      "title": "Design UI",
      "description": "Create user interface mockups",
      "status": "todo",
      "assignee_id": 2,
      "team_id": 2
    }
  ]
}
```

---

### Görev Detayı
```http
GET /tasks/:id
```

Belirli bir görevin detaylarını getirir.

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
    "team_id": 1
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
  "team_id": 1
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
    "team_id": 1
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
  "assignee_id": 2
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
    "team_id": 1
  }
}
```

**Hata Durumları:**
- `404 Not Found` - Görev bulunamadı
- `500 Internal Server Error` - Güncelleme hatası

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

### Update Task Request
- `title`: Opsiyonel, minimum 3 karakter
- `description`: Opsiyonel
- `status`: Opsiyonel, geçerli status değeri
- `assignee_id`: Opsiyonel, geçerli kullanıcı ID'si

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
    "team_id": 1
  }'
```

**Görev güncelle:**
```bash
curl -X PUT http://localhost:8080/api/v1/tasks/1 \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -H "Content-Type: application/json" \
  -d '{
    "status": "done"
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