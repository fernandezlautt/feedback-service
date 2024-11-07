# Arquitectura de microservicios

- Alumno: Lautaro Fernandez
- Año: 2024

# Microservicio de User Feed

### Introducción

El microservicio de User Feed permite a los usuarios del ecommerce dejar comentarios sobre los artículos que han comprado.
Los comentarios proporcionan retroalimentación a otros usuarios sobre la experiencia de compra.
Cada artículo puede tener múltiples comentarios de diferentes usuarios. Si un comentario es negativo (es decir, tiene una calificación baja),
se generará una notificación asincrónica a través de RabbitMQ para que el equipo de atención al cliente envíe un email de seguimiento.

### Lenguaje

- Go
- TypeScript

## Casos de uso

### CU: Crear Feedback

- Se valida que el usuario haya comprado el artículo antes de dejar un comentario.
- Cualquier usuario autenticado que haya comprado el artículo puede dejar un comentario.

### CU: Consultar Feedback

- Se listan los comentarios de un artículo específico mediante su ID.
- Los comentarios incluyen información sobre el usuario y la calificación del artículo.

### CU: Notificar Feedback Negativo

- Si el comentario incluye una calificación baja, se emite un mensaje de notificación a través de RabbitMQ para activar un envío de email de seguimiento.

## Modelo de datos

**Feedback**

- id: ObjectId
- articleId: string
- customerName: String
- feedbackInfo: String
- rating: number (de 1 a 5)
- creationDate: Date

**Notification**

- id: ObjectId
- feedbackId: string
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
	"userId": {"id del usuario que genera el feedback"},
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
            "customerName": {"customerName"},
            "feedbackInfo": {"comentario del usuario"},
            "rating": {"calificación"},
            "creationDate": {"creationDate"}
      }
]
```

## Interfaz asincrónica (RabbitMQ)

**Validación de artículo**

- Envía una solicitud de verificación de existencia del artículo, devolviendo si existe o no.

* Escucha de mensajes article-exist \*

* Envía mensaje a catalog para comprobar que existe un artículo\*

**Notificación de feedback negativo**

- Envía notificación de feedback negativo si el rating es bajo (1 o 2) a Mail Notifications Service\*
