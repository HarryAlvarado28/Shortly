# 📘 Shortly API Reference

Este documento describe los endpoints disponibles en la API del proyecto Shortly, incluyendo los detalles de las rutas, métodos, parámetros y respuestas esperadas.

---

## 🔐 Autenticación

### 📥 `POST /register`

**Descripción:** Registra un nuevo usuario.

**Body JSON:**
```json
{
  "username": "harry",
  "email": "harry@example.com",
  "password": "12345678"
}
```

**Respuesta:**
```json
{
  "message": "Usuario registrado con éxito"
}
```

---

### 🔑 `POST /login`

**Descripción:** Inicia sesión con email y contraseña. Devuelve un token JWT.

**Body JSON:**
```json
{
  "email": "harry@example.com",
  "password": "12345678"
}
```

**Respuesta:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6..."
}
```

---

## ✂️ URLs

### 🔗 `POST /shorten` (autenticado)

**Descripción:** Acorta una URL. Si estás autenticado, se asocia a tu usuario.

**Header:**
```
Authorization: Bearer TU_TOKEN
```

**Body JSON:**
```json
{
  "url": "https://example.com"
}
```

**Respuesta:**
```json
{
  "short_url": "http://localhost:8080/abc123"
}
```

---

### 📄 `GET /my/urls` (autenticado)

**Descripción:** Retorna todas las URLs creadas por el usuario autenticado.

**Header:**
```
Authorization: Bearer TU_TOKEN
```

**Respuesta:**
```json
[
  {
    "short_id": "abc123",
    "original_url": "https://example.com",
    "clicks": 2,
    "user_id": 1,
    ...
  }
]
```

---

### 🔁 `GET /{short_id}`

**Descripción:** Redirecciona a la URL original asociada al identificador corto.

**Ejemplo:**
```
GET /abc123 → 302 Found → https://example.com
```

---

### 📊 `GET /stats/{short_id}`

**Descripción:** Muestra estadísticas del enlace.

**Respuesta:**
```json
{
  "short_id": "abc123",
  "original_url": "https://example.com",
  "clicks": 5,
  "created_at": "2025-03-30T12:00:00Z",
  "expires_at": null
}
```

---

## ⚠️ Errores comunes

| Código | Causa                          |
|--------|---------------------------------|
| 400    | Datos inválidos                |
| 401    | Token ausente o inválido       |
| 404    | URL no encontrada              |
| 500    | Error del servidor o base de datos |

---

## 🔐 Seguridad

- El token JWT tiene una validez de 72 horas.
- Todas las rutas protegidas requieren el header: `Authorization: Bearer TU_TOKEN`

---

## 🧪 Probar con curl

Consulta el `README.md` para ejemplos listos de uso con `curl`.
