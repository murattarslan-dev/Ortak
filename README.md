# Ortak - Takım ve Görev Yönetim API'si

Modern, ölçeklenebilir takım ve görev yönetim sistemi. Go dilinde yazılmış RESTful API ile mikroservis mimarisine uygun modüler yapı.

## 🚀 Özellikler

- **JWT Tabanlı Kimlik Doğrulama** - Güvenli token-based authentication
- **Modüler Mimari** - Mikroservis geçişine hazır clean architecture
- **Standart Response Format** - Tutarlı API response yapısı
- **Comprehensive Middleware** - Logging, error handling, recovery
- **Unit Testing** - Mock repository pattern ile test edilebilir kod
- **Memory Storage** - Development için in-memory database

## 🛠 Teknolojiler

| Kategori | Teknoloji | Açıklama |
|----------|-----------|----------|
| **Backend** | Go 1.21+ | Ana programlama dili |
| **Web Framework** | Gin | HTTP router ve middleware |
| **Authentication** | JWT | Token-based kimlik doğrulama |
| **Security** | bcrypt | Şifre hashleme |
| **Testing** | Go Testing | Unit test framework |
| **Architecture** | Clean Architecture | Modüler ve test edilebilir yapı |

## 📁 Proje Yapısı

```
Ortak/
├── cmd/api/                    # Uygulama giriş noktası
│   └── main.go
├── internal/                   # İç modüller
│   ├── auth/                   # Kimlik doğrulama modülü
│   │   ├── handler/           # HTTP handlers
│   │   ├── service/           # Business logic
│   │   ├── repository/        # Data access layer
│   │   └── model.go           # Data models
│   ├── user/                   # Kullanıcı yönetimi modülü
│   ├── team/                   # Takım yönetimi modülü
│   ├── task/                   # Görev yönetimi modülü
│   ├── middleware/             # HTTP middleware'ler
│   │   ├── auth.go            # JWT middleware
│   │   ├── logger.go          # Request/response logging
│   │   ├── error.go           # Error handling
│   │   └── formatter.go       # Response formatting
│   └── db/                     # Database bağlantısı
├── pkg/                        # Paylaşılan paketler
│   ├── utils/                 # Yardımcı fonksiyonlar
│   └── response/              # Standart response yapısı
├── docs/                       # API dokümantasyonu
│   ├── auth/                  # Auth API docs
│   ├── user/                  # User API docs
│   ├── team/                  # Team API docs
│   └── task/                  # Task API docs
└── README.md                   # Bu dosya
```

## 🏗 Mimari

### Clean Architecture Katmanları

1. **Handler Layer** - HTTP request/response handling
2. **Service Layer** - Business logic ve validation
3. **Repository Layer** - Data access abstraction
4. **Model Layer** - Data structures

### Dependency Injection

```go
// Repository Interface
type Repository interface {
    GetAll() []User
    GetByID(id string) *User
    Create(user User) *User
}

// Service uses Repository
type Service struct {
    repo Repository
}

// Handler uses Service
type Handler struct {
    service *Service
}
```

## 🚦 Hızlı Başlangıç

### Gereksinimler
- Go 1.21+
- Git

### Kurulum

1. **Projeyi klonlayın:**
```bash
git clone <repository-url>
cd Ortak
```

2. **Bağımlılıkları yükleyin:**
```bash
go mod tidy
```

3. **Uygulamayı çalıştırın:**
```bash
go run cmd/api/main.go
```

4. **API'yi test edin:**
```bash
curl http://localhost:8080/api/v1/health
```

## 📚 API Dokümantasyonu

Her modül için detaylı API dokümantasyonu:

- **[Authentication API](docs/auth/README.md)** - Kayıt, giriş, çıkış
- **[User Management API](docs/user/README.md)** - Kullanıcı CRUD işlemleri
- **[Team Management API](docs/team/README.md)** - Takım CRUD işlemleri
- **[Task Management API](docs/task/README.md)** - Görev CRUD işlemleri

### Base URL
```
http://localhost:8080/api/v1
```

### Standart Response Format

**Başarılı Response:**
```json
{
  "success": true,
  "message": "Operation completed successfully",
  "data": { ... }
}
```

**Hata Response:**
```json
{
  "success": false,
  "message": "Error description"
}
```

## 🧪 Testing

### Unit Testleri Çalıştırma

```bash
# Tüm testler
go test ./...

# Belirli modül
go test ./internal/user/handler -v

# Coverage raporu
go test -cover ./...
```

### Test Yapısı

- **Mock Repository Pattern** - Test için sahte data layer
- **HTTP Test** - Gin test context ile endpoint testleri
- **Service Tests** - Business logic unit testleri

## 🔧 Geliştirme

### Yeni Modül Ekleme

1. `internal/` altında yeni klasör oluşturun
2. Handler, Service, Repository katmanlarını ekleyin
3. Model'leri tanımlayın
4. Unit testleri yazın
5. `main.go`'da route'ları ekleyin

### Middleware Ekleme

1. `internal/middleware/` altında yeni dosya oluşturun
2. Gin middleware pattern'ini kullanın
3. `main.go`'da middleware'i kaydedin

## 🚀 Deployment

### Development
```bash
go run cmd/api/main.go
```

### Production Build
```bash
go build -o ortak cmd/api/main.go
./ortak
```

## 🤝 Katkıda Bulunma

1. Fork edin
2. Feature branch oluşturun (`git checkout -b feature/amazing-feature`)
3. Commit edin (`git commit -m 'Add amazing feature'`)
4. Push edin (`git push origin feature/amazing-feature`)
5. Pull Request oluşturun

## 📄 Lisans

Bu proje eğitim amaçlı geliştirilmiştir.

---

**Geliştirici:** Murat Arslan  
**Versiyon:** 1.0.0  
**Son Güncelleme:** Aralık 2024