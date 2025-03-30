# 🔗 Shortly

Shortly es un acortador de URLs rápido y simple, ahora con soporte para autenticación mediante JWT.  
Ideal para desarrolladores que buscan gestionar enlaces personales de forma segura.

---

## 🚀 Funcionalidades

- ✅ Acortamiento rápido de enlaces
- 🔁 Redirección automática
- 👤 Registro e inicio de sesión con JWT
- 🔐 Asociación de URLs por usuario
- 📊 Estadísticas de uso (clics)
- 🧠 Almacenamiento en PostgreSQL
- 🐳 Despliegue listo con Docker

---

## 📦 Requisitos

- Go 1.21 o superior
- PostgreSQL 13+
- (Opcional) Docker y Docker Compose

---

## ⚙️ Variables de entorno (.env)

Crea un archivo `.env` en la raíz con el siguiente contenido:

```env
BASE_URL=http://localhost:8080
DB_URL=postgres://usuario:contraseña@localhost:5432/shortly
JWT_SECRET=super_secreta
```

---

## 🛠️ Instalación local

```bash
git clone https://github.com/HarryAlvarado28/shortly.git
cd shortly
go mod tidy
go run .
```

---

## 🐳 Uso con Docker

```bash
docker build -t shortly .
docker run -p 8080:8080 --env-file .env shortly
```

---

## 📌 Endpoints principales

### 🔐 Autenticación

#### Registro

```bash
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"username": "harry", "email": "harry@example.com", "password": "12345678"}'
```

#### Login

```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"email": "harry@example.com", "password": "12345678"}'
```

✅ Devuelve un token JWT para usar en las siguientes rutas

---

### ✂️ Acortar URL (requiere token)

```bash
curl -X POST http://localhost:8080/shorten \
  -H "Authorization: Bearer-TU_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"url": "https://example.com"}'
```

---

### 📄 Ver tus URLs (autenticado)

```bash
curl http://localhost:8080/my/urls \
  -H "Authorization: Bearer-TU_TOKEN"
```

---

### 🔁 Redirección

```bash
curl -L http://localhost:8080/abc123
```

---

### 📊 Ver estadísticas

```bash
curl http://localhost:8080/stats/abc123
```

---

## 🧪 Pruebas y depuración

El sistema de variables usa `GetEnv`, por lo que verás logs como:

```
[ENV] JWT_SECRET = super_secreta (desde jwt.go:25 → utils.GenerateJWT)
```

---

## 📄 Licencia

MIT © 2025 - HarryLab28

---

## 🌟 ¿Te gustó el proyecto?

Dale ⭐ en GitHub o contribuye con ideas nuevas 😄
