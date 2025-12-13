# Upload API Dokümantasyonu

Ortak projesi için dosya yükleme ve yönetim API'si.

## 🚀 Özellikler

- **Güvenli Upload**: JWT tabanlı kimlik doğrulama
- **Otomatik Thumbnail**: Resimler için 200x200 thumbnail oluşturma
- **Metadata Extraction**: Dosya boyutu, MIME type, resim boyutları
- **Docker Volume**: Kalıcı dosya depolama
- **Database Tracking**: Tüm upload bilgileri DB'de saklanır
- **Admin Panel**: Admin kullanıcılar tüm dosyaları görebilir

## 📁 Desteklenen Formatlar

| Kategori | Formatlar | Thumbnail |
|----------|-----------|-----------|
| **Resim** | JPG, PNG, GIF | ✅ |
| **Dokument** | PDF, DOC, DOCX | ❌ |
| **Diğer** | TXT, ZIP, vb. | ❌ |

## 🔧 Konfigürasyon

### Environment Variables

```bash
UPLOAD_DIR=/app/uploads          # Upload klasörü
MAX_FILE_SIZE=52428800          # 50MB (bytes)
```

### Docker Volume

```yaml
volumes:
  - upload_data:/app/uploads     # Kalıcı depolama
```

## 📚 API Endpoints

### Base URL
```
http://localhost:8080/api/v1
```

---

## 📤 Dosya Yükleme

### `POST /upload`

**Headers:**
```
Authorization: Bearer JWT_TOKEN
Content-Type: multipart/form-data
```

**Body:**
```
file: [binary file]
```

**Response:**
```json
{
  "success": true,
  "message": "Dosya başarıyla yüklendi",
  "data": {
    "file_id": "550e8400-e29b-41d4-a716-446655440000",
    "file_name": "image.jpg",
    "size": 1024000,
    "url": "/uploads/550e8400-e29b-41d4-a716-446655440000.jpg",
    "thumb_url": "/uploads/thumbs/550e8400-e29b-41d4-a716-446655440000_thumb.jpg",
    "mime_type": "image/jpeg",
    "width": 1920,
    "height": 1080
  }
}
```

**cURL Örneği:**
```bash
curl -X POST http://localhost:8080/api/v1/upload \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -F "file=@/path/to/image.jpg"
```

---

## 📋 Upload Bilgileri

### `GET /uploads`

Kullanıcının yüklediği dosyaları listeler. Admin kullanıcılar tüm dosyaları görebilir.

**Headers:**
```
Authorization: Bearer JWT_TOKEN
```

**Response (Normal User):**
```json
{
  "success": true,
  "data": [
    {
      "id": "550e8400-e29b-41d4-a716-446655440000",
      "original_name": "my_document.pdf",
      "file_name": "550e8400-e29b-41d4-a716-446655440000.pdf",
      "size": 2048000,
      "mime_type": "application/pdf",
      "uploaded_by": "user-123",
      "uploaded_at": "2024-12-01T10:30:00Z"
    }
  ]
}
```

**Response (Admin User):**
```json
{
  "success": true,
  "data": [
    {
      "id": "550e8400-e29b-41d4-a716-446655440000",
      "original_name": "user1_image.jpg",
      "uploaded_by": "user-123",
      "uploaded_at": "2024-12-01T10:30:00Z"
    },
    {
      "id": "660f9511-f3ac-52e5-b827-557766551111",
      "original_name": "user2_document.pdf", 
      "uploaded_by": "user-456",
      "uploaded_at": "2024-12-01T09:15:00Z"
    }
  ]
}
```

---

### `GET /uploads?id=FILE_ID`

Belirli bir dosyanın detay bilgilerini getirir.

**Headers:**
```
Authorization: Bearer JWT_TOKEN
```

**Response:**
```json
{
  "success": true,
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "original_name": "image.jpg",
    "file_name": "550e8400-e29b-41d4-a716-446655440000.jpg",
    "size": 1024000,
    "content_type": "image/jpeg",
    "path": "/app/uploads/550e8400-e29b-41d4-a716-446655440000.jpg",
    "uploaded_by": "user-123",
    "uploaded_at": "2024-12-01T10:30:00Z",
    "mime_type": "image/jpeg",
    "width": 1920,
    "height": 1080,
    "thumb_url": "/uploads/thumbs/550e8400-e29b-41d4-a716-446655440000_thumb.jpg"
  }
}
```

---

## 🌐 Dosya Erişimi

### Public URL Access

Yüklenen dosyalar public URL ile erişilebilir:

```
# Orijinal dosya
http://localhost:8080/uploads/FILE_NAME.ext

# Thumbnail (sadece resimler)
http://localhost:8080/uploads/thumbs/FILE_ID_thumb.ext
```

**Örnek:**
```
http://localhost:8080/uploads/550e8400-e29b-41d4-a716-446655440000.jpg
http://localhost:8080/uploads/thumbs/550e8400-e29b-41d4-a716-446655440000_thumb.jpg
```

---

## 🔒 Güvenlik

### Kimlik Doğrulama
- Tüm upload işlemleri JWT token gerektirir
- Token `Authorization: Bearer TOKEN` header'ında gönderilir

### Yetkilendirme
- **Normal User**: Sadece kendi dosyalarını görebilir
- **Admin User**: Tüm kullanıcıların dosyalarını görebilir

### Dosya Güvenliği
- Maksimum dosya boyutu: 50MB (konfigüre edilebilir)
- Dosya adları UUID ile güvenli hale getirilir
- MIME type kontrolü yapılır

---

## 🗄️ Database Schema

### Files Tablosu

```sql
CREATE TABLE files (
    id VARCHAR(36) PRIMARY KEY,
    original_name VARCHAR(255) NOT NULL,
    file_name VARCHAR(255) NOT NULL,
    size BIGINT NOT NULL,
    content_type VARCHAR(100),
    path TEXT NOT NULL,
    uploaded_by VARCHAR(36),
    uploaded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    mime_type VARCHAR(100),
    width INTEGER DEFAULT 0,
    height INTEGER DEFAULT 0,
    thumb_url TEXT
);
```

---

## 🐳 Docker Deployment

### Volume Yönetimi

```bash
# Volume listesi
docker volume ls

# Backup
docker run --rm -v upload_data:/data -v $(pwd):/backup alpine \
  tar czf /backup/upload_backup.tar.gz /data

# Restore
docker run --rm -v upload_data:/data -v $(pwd):/backup alpine \
  tar xzf /backup/upload_backup.tar.gz -C /
```

---

## 🚨 Hata Kodları

| Kod | Açıklama |
|-----|----------|
| `400` | Dosya yüklenemedi / Boyut çok büyük |
| `401` | Yetkisiz erişim (JWT token gerekli) |
| `404` | Dosya bulunamadı |
| `500` | Sunucu hatası |

---

## 📊 Örnek Kullanım Senaryoları

### 1. Profil Fotoğrafı Yükleme
```bash
curl -X POST http://localhost:8080/api/v1/upload \
  -H "Authorization: Bearer JWT_TOKEN" \
  -F "file=@profile.jpg"
```

### 2. Döküman Yükleme
```bash
curl -X POST http://localhost:8080/api/v1/upload \
  -H "Authorization: Bearer JWT_TOKEN" \
  -F "file=@document.pdf"
```

### 3. Upload Geçmişi
```bash
curl -H "Authorization: Bearer JWT_TOKEN" \
  http://localhost:8080/api/v1/uploads
```

### 4. Admin Panel
```bash
# Admin token ile tüm dosyalar
curl -H "Authorization: Bearer ADMIN_JWT_TOKEN" \
  http://localhost:8080/api/v1/uploads
```

---

## 🔧 Geliştirici Notları

### Thumbnail Oluşturma
- Sadece JPG ve PNG formatları desteklenir
- 200x200 piksel boyutunda oluşturulur
- Aspect ratio korunur (Lanczos3 algoritması)

### Performans
- Gin Static middleware ile optimize edilmiş file serving
- Database indexleri upload performansını artırır
- Docker volume ile I/O performansı

### Monitoring
- Tüm upload işlemleri loglanır
- Database'de audit trail tutulur
- File system ve DB senkronizasyonu

---

**Geliştirici:** Murat Arslan  
**Versiyon:** 1.0.0  
**Son Güncelleme:** Aralık 2024