# Ortak - Takım ve Görev Yönetim API'si

Eğitim amaçlı geliştirilmiş bir backend projesi. Go dilinde yazılmış RESTful API ile takım ve görev yönetim-takip uygulaması.

## Özellikler

- JWT tabanlı kimlik doğrulama
- Kullanıcı yönetimi (CRUD)
- Takım yönetimi (CRUD)
- Görev yönetimi (CRUD)
- PostgreSQL veritabanı desteği
- Docker containerization
- Environment variables ile yapılandırma

## Teknolojiler

- **Go 1.21**
- **Gin Web Framework** - HTTP router ve middleware
- **PostgreSQL** - Veritabanı
- **JWT** - Kimlik doğrulama
- **bcrypt** - Şifre hashleme
- **Docker** - Containerization

## Kurulum

### Gereksinimler
- Go 1.21+
- PostgreSQL
- Git

### Adımlar

1. Projeyi klonlayın:
```bash
git clone <repository-url>
cd Ortak
```

2. Bağımlılıkları yükleyin:
```bash
go mod tidy
```

3. Environment dosyasını oluşturun:
```bash
cp .env.example .env
```

4. `.env` dosyasını düzenleyin:
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=ortak
JWT_SECRET=your-secret-key
```

5. PostgreSQL veritabanını oluşturun:
```sql
CREATE DATABASE ortak;
```

6. Uygulamayı çalıştırın:
```bash
go run cmd/api/main.go
```

Uygulama `http://localhost:8080` adresinde çalışacaktır.

## Docker ile Çalıştırma

```bash
docker build -t ortak .
docker run -p 8080:8080 --env-file .env ortak
```

## API Endpoints

### Base URL
```
http://localhost:8080/api/v1
```

### Kimlik Doğrulama

#### Kayıt Ol
```http
POST /register
```

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
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": 1,
    "username": "johndoe",
    "email": "john@example.com"
  }
}
```

**Hata Durumları:**
- `400 Bad Request` - Geçersiz veri formatı
- `500 Internal Server Error` - Sunucu hatası

#### Giriş Yap
```http
POST /login
```

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
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": 1,
    "username": "johndoe",
    "email": "john@example.com"
  }
}
```

**Hata Durumları:**
- `400 Bad Request` - Geçersiz veri formatı
- `401 Unauthorized` - Hatalı kimlik bilgileri

### Kullanıcı Yönetimi

> **Not:** Tüm kullanıcı endpoints'leri JWT token gerektirir.
> Header: `Authorization: Bearer <token>`

#### Kullanıcıları Listele
```http
GET /users
```

**Response (200 OK):**
```json
[
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
```

#### Kullanıcı Oluştur
```http
POST /users
```

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
  "id": 3,
  "username": "newuser",
  "email": "newuser@example.com"
}
```

### Takım Yönetimi

> **Not:** Tüm takım endpoints'leri JWT token gerektirir.

#### Takımları Listele
```http
GET /teams
```

**Response (200 OK):**
```json
[
  {
    "id": 1,
    "name": "Development Team",
    "description": "Backend development",
    "owner_id": 1
  },
  {
    "id": 2,
    "name": "Design Team",
    "description": "UI/UX design",
    "owner_id": 2
  }
]
```

#### Takım Oluştur
```http
POST /teams
```

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
  "id": 3,
  "name": "QA Team",
  "description": "Quality Assurance team",
  "owner_id": 1
}
```

### Görev Yönetimi

> **Not:** Tüm görev endpoints'leri JWT token gerektirir.

#### Görevleri Listele
```http
GET /tasks
```

**Response (200 OK):**
```json
[
  {
    "id": 1,
    "title": "Setup API",
    "description": "Create REST API",
    "status": "in_progress",
    "assignee_id": 1,
    "team_id": 1
  },
  {
    "id": 2,
    "title": "Design UI",
    "description": "Create user interface",
    "status": "todo",
    "assignee_id": 2,
    "team_id": 2
  }
]
```

#### Görev Oluştur
```http
POST /tasks
```

**Request Body:**
```json
{
  "title": "Write Tests",
  "description": "Write unit tests for API",
  "assignee_id": 1,
  "team_id": 1
}
```

**Response (201 Created):**
```json
{
  "id": 3,
  "title": "Write Tests",
  "description": "Write unit tests for API",
  "status": "todo",
  "assignee_id": 1,
  "team_id": 1
}
```

## Hata Kodları

| Kod | Açıklama |
|-----|----------|
| 200 | Başarılı |
| 201 | Oluşturuldu |
| 400 | Hatalı İstek |
| 401 | Yetkisiz Erişim |
| 500 | Sunucu Hatası |

## Görev Durumları

- `todo` - Yapılacak
- `in_progress` - Devam Ediyor
- `done` - Tamamlandı

## Proje Yapısı

```
/Ortak
  /cmd/api
    main.go                # Uygulama giriş noktası
  /internal
    /auth                  # Kimlik doğrulama
      handler.go
      service.go
      model.go
    /user                  # Kullanıcı CRUD
      handler.go
      service.go
      model.go
    /team                  # Takım CRUD
      handler.go
      service.go
      model.go
    /task                  # Görev CRUD
      handler.go
      service.go
      model.go
    /db                    # Veritabanı bağlantısı
      db.go
    /middleware            # Middleware'ler
      auth.go
  /pkg
    /utils                 # Yardımcı fonksiyonlar
      jwt.go
      hash.go
  go.mod
  go.sum
  Dockerfile
  .env
  .env.example
  README.md
```

## Geliştirme

### Yeni Özellik Ekleme

1. `internal/` klasörü altında yeni modül oluşturun
2. Handler, Service ve Model dosyalarını ekleyin
3. `main.go` dosyasında route'ları tanımlayın

### Test Etme

API'yi test etmek için Postman, curl veya benzeri araçları kullanabilirsiniz.

**Örnek curl komutu:**
```bash
# Kayıt ol
curl -X POST http://localhost:8080/api/v1/register \
  -H "Content-Type: application/json" \
  -d '{"username":"test","email":"test@example.com","password":"password123"}'

# Giriş yap
curl -X POST http://localhost:8080/api/v1/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"password123"}'

# Takımları listele (token gerekli)
curl -X GET http://localhost:8080/api/v1/teams \
  -H "Authorization: Bearer <your-token>"
```

## Katkıda Bulunma

1. Fork edin
2. Feature branch oluşturun (`git checkout -b feature/amazing-feature`)
3. Commit edin (`git commit -m 'Add amazing feature'`)
4. Push edin (`git push origin feature/amazing-feature`)
5. Pull Request oluşturun

## Lisans

Bu proje eğitim amaçlı geliştirilmiştir.
