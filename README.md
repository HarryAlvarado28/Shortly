# Shortly 🔗 ![REST API](https://img.shields.io/badge/type-REST%20API-blue)

Acorta URLs largas en versiones cortas y fáciles de compartir.  
Redirige automáticamente al enlace original cuando visitas la URL generada.

---

## 🚀 Funcionalidades

- ✅ Acortamiento rápido de enlaces
- 🔁 Redirección automática
- 🧠 Almacenamiento en memoria
- 🐳 Despliegue listo con Docker

---

## 📦 Instalación local

### 🔧 Requisitos

- Go 1.18 o superior
- (Opcional) Docker

### 🛠️ Clonar el repositorio

```bash
git clone https://github.com/HarryAlvarado28/shortly.git
cd shortly
go run main.go
```

---

## 📬 Endpoints de la API
### 🔗 POST /shorten
Acorta una URL larga y devuelve una versión corta.

### 📥 Ejemplo con curl:

```bash
curl -X POST http://localhost:8080/shorten \
  -H "Content-Type: application/json" \
  -d '{"url":"https://example.com/mi-url-larga"}'
```

### 📤 Respuesta esperada:

```bash
{
  "short_url": "http://localhost:8080/abc123"
}
```

### 🚀 GET /{id}
Redirige automáticamente a la URL original. El parámetro {id} es el código generado por el acortador.

🔄 Ejemplo con curl:

```bash
curl -L http://localhost:8080/abc123
```
