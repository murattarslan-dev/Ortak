# Ortak API - Çalıştırma Scriptleri

## Kullanılabilir Scriptler

### 1. API Server Başlatma

#### `build_and_run.bat`
- API'yi derler (ortak-api.exe)
- Port çakışmalarını çözer
- Yeni terminalde çalıştırır
- Renkli çıktı

### 2. Test Çalıştırma

#### `run_tests.bat`
- Tüm unit testleri çalıştırır
- Mavi renk teması
- Test sonuçlarını gösterir

## API Endpoints

- **POST** `/api/v1/register` - Kullanıcı kaydı
- **POST** `/api/v1/login` - Giriş
- **DELETE** `/api/v1/logout` - Çıkış
- **GET** `/api/v1/users` - Kullanıcıları listele
- **POST** `/api/v1/users` - Kullanıcı oluştur
- **GET** `/api/v1/teams` - Takımları listele
- **POST** `/api/v1/teams` - Takım oluştur
- **GET** `/api/v1/tasks` - Görevleri listele
- **POST** `/api/v1/tasks` - Görev oluştur
- **GET** `/api/v1/health` - Sistem durumu

Server: `http://localhost:8080`