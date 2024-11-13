# Arquitectura de microservicios

- Alumno: Lautaro Fernandez
- Año: 2024

# Microservicio de User Feed

### Introducción

El microservicio de User Feed permite a los usuarios del ecommerce dejar comentarios sobre los artículos que han comprado.
Los comentarios proporcionan retroalimentación a otros usuarios.
Cada artículo puede tener múltiples comentarios de diferentes usuarios.
Se pueden deshabilitar los comentarios de un artículo si se considera necesario.

### Lenguaje

- Go

## Casos de uso

### CU: Crear Feedback

- Se valida que el usuario haya comprado el artículo antes de dejar un comentario.
- Cualquier usuario autenticado puede dejar un comentario.

### CU: Consultar Feedback

- Se listan los comentarios de un artículo específico mediante su ID.
- Los comentarios incluyen información sobre el usuario y la calificación del artículo.

### CU: Deshabitación de Feedback

- Se deshabilita el feedback por la razón que se especifique.


## Modelo de datos

**Feedback**

- id: ObjectId
- articleId: string
- customerName: string
- feedbackInfo: string
- rating: number (de 1 a 5)
- status: string
- reason: string
- creationDate: Date

## Interfaz REST

**Creación de feedback de un artículo**
`POST /v1/feedback/{articleId}`

#### Headers

| Cabecera                  | Contenido            |
| ------------------------- | -------------------- |
| Authorization: Bearer xxx | Token en formato JWT |

_Body_

```json
{
	"feedbackInfo": {"comentario del usuario"},
	"rating": {"calificación del artículo (1-5)"}
}
```

_Response_
`200 OK` si el feedback fue creado con éxito | `400 BAD REQUEST`

**Consulta de feedback de un artículo**
`GET /v1/feedback/{articleId}`

#### Headers

| Cabecera                  | Contenido            |
| ------------------------- | -------------------- |
| Authorization: Bearer xxx | Token en formato JWT |

_Response_
`200 OK` si existe por lo menos un feedback en ese artículo | `404 NOT FOUND`

```json
[
    {
        "id": "ID del feedback",
        "articleId": "ID del artículo",
        "customerName": "Nombre del cliente",
        "customerId": "ID del cliente",
        "feedbackInfo": "Información del feedback",
        "rating": "Calificación del artículo (1-5)",
        "creationDate": "Fecha de creación del feedback",
        "status": "Estado del feedback",
        "reason": "Razón del estado del feedback"
    }
]
```

**Deshabilitación de feedback de un artículo**
`POST /v1/feedback/disable/{articleId}`

#### Headers

| Cabecera                  | Contenido            |
| ------------------------- | -------------------- |
| Authorization: Bearer xxx | Token en formato JWT |

_Body_

```json
{
	"reason": "Razón de la deshabilitación"
}
```

_Response_
`200 OK` si el feedback fue deshabilitado con éxito | `400 BAD REQUEST`


## Interfaz asincrónica (RabbitMQ)

**Validación de artículo**

- Envía una solicitud de verificación de existencia del artículo, devolviendo si existe o no.

- Escucha de mensajes catalog_article_exist_receive

- Actualiza el estado del feedback a "confirmed"
