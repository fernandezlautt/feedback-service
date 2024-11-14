# Arquitectura de microservicios

- Alumno: Lautaro Fernandez
- Año: 2024

# Microservicio de User Feed

### Introducción

El microservicio de User Feed permite a los usuarios del ecommerce dejar comentarios sobre los artículos que han comprado.
Los comentarios proporcionan retroalimentación a otros usuarios.
Cada artículo puede tener múltiples comentarios de diferentes usuarios.
Se pueden deshabilitar los comentarios de un artículo si se considera necesario.

### Tecnologías

- Go
- MongoDB
- RabbitMQ

## Casos de uso

### CU: Crear Feedback

- Cualquier usuario autenticado puede dejar un comentario.
- Se puede calificar el artículo con un número del 1 al 5.
- Se valida que el articulo exista, por lo tanto se establece en pending el estado del feedback hasta que `cataloggo` responda.
- Se guarda el userId (para implementacion de notificación) y el nombre del usuario que dejo el comentario para poder mostrarlo en la interfaz.

### CU: Consultar Feedback

- Se listan los comentarios de un artículo específico mediante su ID )enviado por query param).
- Los comentarios incluyen el nombre del usuario, el comentario, la calificación del artículo y demas datos del artículo.

### CU: Deshabitación de Feedback

- Se deshabilita el feedback por la razón que se especifique.
- La deshabilitacion es a nivel feedback no artículo, sirve, por ejemplo, para comentarios inapropiados.


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
`POST /v1/feedback`

#### Headers

| Cabecera                  | Contenido            |
| ------------------------- | -------------------- |
| Authorization: Bearer xxx | Token en formato JWT |

_Body_

```json
{
	"feedbackInfo": {"comentario del usuario"},
	"rating": {"calificación del artículo (1-5)"},
	"articleId": {"ID del artículo a comentar"}
}
```

Ejemplo:
```json
{
  "feedbackInfo": "Muy buen producto",
  "rating": 5,
  "articleId": "60f3b3b3b3b3b3b3b3b3b3"
}
```

_Response_


`200 OK` si el feedback fue creado con éxito | `400 BAD REQUEST`

**Consulta de feedback de un artículo**
`GET /v1/feedback`

#### Query Params

| Parámetro | Descripción     |
| --------- | ----------------|
| articleId | ID del artículo |
| page      | Página a consultar |
| size      | Cantidad de elementos por página |

#### Headers

| Cabecera                  | Contenido            |
| ------------------------- | -------------------- |
| Authorization: Bearer xxx | Token en formato JWT |

_Response_


`200 OK` si existe por lo menos un feedback en ese artículo | `404 NOT FOUND` si no se encontro feedback para el articulo | `400 BAD REQUEST` si no se agrego articleId o la request fue malformada.

Es importante remarcar que solo los feedbacks con estado "confirmed" son devueltos.

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
Ejemplo:
```json
[
    {
        "id": "60f3b3b3b3b3b3b3b3b3b3b3",
        "articleId": "60f3b3b3b3b3b3b3b3b3b3",
        "customerName": "Lautaro Fernandez",
        "customerId": "60f3b3b3b3b3b3b3b3b3b3",
        "feedbackInfo": "Muy buen producto",
        "rating": 5,
        "creationDate": "2024-07-18T00:00:00Z",
        "status": "confirmed",
        "reason": ""
    }
]
```

**Deshabilitación de feedback de un artículo**
`PATCH /v1/feedback/{feedbackId}/disable`

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
Ejemplo:
```json
{
  "reason": "Comentario inapropiado"
}
```

_Response_


`200 OK` si el feedback fue deshabilitado con éxito | `400 BAD REQUEST`


## Interfaz asincrónica (RabbitMQ)

**Validación de artículo**

- Envía una solicitud de verificación de existencia del artículo al microservicio `cataloggo`.

- Escucha de mensajes de validez en `catalog_article_exist_receive`.

- Actualiza el estado del feedback a `"confirmed"` si `valid` es true.

_Body enviado a catalog_

```json
{
    "correlation_id": "Id de correlación, utilizado para rastrear la solicitud y la respuesta",
    "message": {
        "articleId": "Id del artículo, utilizado para identificar el artículo en cuestión",
        "referenceId": "Id de referencia del objeto remoto"
    },
    "routing_key": "Clave de enrutamiento remoto para responder",
    "exchange": "Intercambio remoto para responder"
}
```
Ejemplo:
```json
{
    "correlation_id": "60f3b3b3b3b3b3b3b3b3b3",
    "message": {
        "articleId": "60f3b3b3b3b3b3b3b3b3b3",
        "referenceId": "60f3b3b3b3b3b3b3b3b3b3"
    },
    "routing_key": "catalog_article_exist_send",
    "exchange": "catalog"
}
```
_Body recibido de catalog_
```json
{
    "correlation_id": "Id de correlación, utilizado para rastrear la solicitud y la respuesta",
    "message": {
        "articleId": "Id del artículo, utilizado para identificar el artículo en cuestión",
        "price": "Precio del artículo",
        "referenceId": "Id de referencia del objeto remoto",
        "stock": "Stock del artículo",
        "valid": "Estado de validez del artículo"
    }
}
```
Ejemplo:
```json
{
    "correlation_id": "60f3b3b3b3b3b3b3b3b3b3",
    "message": {
        "articleId": "60f3b3b3b3b3b3b3b3b3b3",
        "referenceId": "60f3b3b3b3b3b3b3b3b3b3"
    },
    "routing_key": "catalog_article_exist_response",
    "exchange": "catalog"
}
```
