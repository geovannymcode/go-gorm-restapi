# Go GORM REST API

Este proyecto es una API REST desarrollada en Go utilizando el framework GORM para la gestión de una base de datos PostgreSQL. Está basada en el enrutador `gorilla/mux` y utiliza `gorm.io/gorm` como ORM.

## 📌 Tecnologías utilizadas

- **Go 1.21.6**
- **GORM** (ORM para Go)
- **PostgreSQL** (como base de datos)
- **gorilla/mux** (enrutador HTTP)
- **Echo** (framework web ligero)

## 🚀 Instalación y configuración

### 1️⃣ Clonar el repositorio
```bash
git clone https://github.com/geovannymcode/go-gorm-restapi.git
cd go-gorm-restapi
```

### 2️⃣ Configurar las dependencias

Asegúrate de tener Go instalado y ejecuta el siguiente comando para descargar las dependencias:
```bash
go mod tidy
```

### 3️⃣ Configurar la base de datos

Asegúrate de tener un servidor PostgreSQL en ejecución y crea una base de datos para la aplicación. Modifica las credenciales en el archivo de configuración correspondiente.

4️⃣ Ejecutar la aplicación

Para iniciar el servidor, ejecuta:
```bash
go run main.go
```

📌 Endpoints disponibles

La API REST maneja operaciones CRUD para una entidad base. A continuación, se presentan los endpoints principales:

| Método | Endpoint        | Descripción                      |
|--------|---------------|--------------------------------|
| GET    | `/items`       | Obtiene todos los ítems       |
| GET    | `/items/{id}`  | Obtiene un ítem por ID        |
| POST   | `/items`       | Crea un nuevo ítem           |
| PUT    | `/items/{id}`  | Actualiza un ítem existente  |
| DELETE | `/items/{id}`  | Elimina un ítem por ID       |

🔧 Estructura del proyecto
```plainText
/go-gorm-restapi
│── main.go             # Punto de entrada de la aplicación
│── config/             # Configuración de la base de datos
│── models/             # Definición de estructuras y modelos
│── routes/             # Definición de rutas HTTP
│── controllers/        # Controladores para la API
│── database/           # Conexión y migraciones de la BD
│── go.mod              # Módulo y dependencias del proyecto
│── go.sum              # Hash de dependencias
```

🛠️ Dependencias principales
```go
require (
    github.com/labstack/gommon v0.4.2
    github.com/gorilla/mux v1.8.1
    gorm.io/driver/postgres v1.5.6
    gorm.io/gorm v1.25.7
)
```

📝 Licencia

Este proyecto está bajo la licencia MIT. Consulta el archivo LICENSE para más detalles.

✍️ Autor: Geovanny Mendoza
