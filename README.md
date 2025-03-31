# 🔗 Shortly

Shortly es un acortador de URLs rápido y simple, ahora con soporte para autenticación mediante JWT.  
Ideal para desarrolladores que buscan gestionar enlaces personales de forma segura.

---

## 🚀 Funcionalidades

- ✅ Acortamiento rápido de enlaces
- 🔁 Redirección automática
- 👤 Registro e inicio de sesión con JWT
- 🔐 Asociación de URLs por usuario
- ✨ Sesión anónima automática con JWT
- 📊 Estadísticas de uso (clics)
- ⏳ Expiración de enlaces anónimos (15 días)
- 🧠 Almacenamiento en PostgreSQL
- 🐳 Despliegue listo con Docker

---

## 📦 Requisitos

- Go 1.21 o superior
- PostgreSQL 13+
- (Opcional) Docker y Docker Compose

---

## ⚙️ Variables de entorno (.env)

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

#### Sesión anónima

```bash
curl -X POST http://localhost:8080/anon
```

---

### ✂️ Acortar URL (requiere token)

```bash
curl -X POST http://localhost:8080/shorten \
  -H "Authorization: Bearer TU_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"url": "https://example.com"}'
```

---

### 📄 Ver tus URLs (autenticado)

```bash
curl http://localhost:8080/my/urls \
  -H "Authorization: Bearer TU_TOKEN"
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

El sistema de variables usa `GetEnvCached`, por lo que verás logs como:

```
[ENV] JWT_SECRET = super_secreta (desde jwt.go:25 → utils.GenerateJWT)
```

---

## 🧭 Roadmap de próximas funcionalidades

Estas son algunas ideas futuras para seguir mejorando Shortly:

- 🔑 Recuperación de contraseña por email
- 📁 Historial de URLs creadas por el usuario
- 🔒 Eliminación de enlaces creados (para usuarios autenticados)
- 📊 Estadísticas avanzadas (clics por día, geolocalización, referer)
- 🧾 Vista previa de enlaces antes de redirigir
- 🎨 Diseño visual con temas claros/oscuro
- 🌍 Localización (multi-idioma)
- 🧪 Tests automatizados y CI/CD

¿Tienes una idea o sugerencia? ¡Contribuye al proyecto o abre un issue! 🚀

---

## 🔖 Release actual: `v1.1.0`

Versión estable con:
- Soporte para usuarios anónimos automáticos con JWT
- Asociación de URLs a usuarios (registrados o anónimos)
- Frontend limpio y funcional en HTML + JS
- Expiración automática de enlaces anónimos a los 15 días
- Backend con Go + PostgreSQL
- Despliegue sencillo con Docker y Render

---

## 📄 Licencia

MIT © 2025 - HarryLab28

---

## 🌟 ¿Te gustó el proyecto?

Dale ⭐ en GitHub o contribuye con ideas nuevas 😄