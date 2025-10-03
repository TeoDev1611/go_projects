# 🔐 GoPass - Generador de Contraseñas Seguras

Un generador de contraseñas aleatorias y criptográficamente seguras desarrollado en Go, con interfaz interactiva en la terminal.

## ✨ Características

- **Generación criptográficamente segura**: Utiliza `crypto/rand` para garantizar aleatoriedad de alta calidad
- **Interfaz interactiva**: Menú intuitivo basado en terminal con la biblioteca Survey
- **Personalizable**: Configura exactamente qué tipos de caracteres incluir
- **Exclusión de caracteres ambiguos**: Opción para evitar caracteres que se confunden visualmente (como `1`, `l`, `I`, `0`, `O`)
- **Sin sesgo estadístico**: Implementa un algoritmo que elimina el sesgo modular en la generación aleatoria

## 🚀 Instalación

### Prerrequisitos

- Go 1.16 o superior instalado en tu sistema

### Pasos

1. Clona el repositorio:
```bash
git clone <url-del-repositorio>
cd gopass
```

2. Instala las dependencias:
```bash
go get github.com/AlecAivazis/survey/v2
```

3. Compila el proyecto:
```bash
go build -o gopass
```

## 💻 Uso

Ejecuta el programa:

```bash
./gopass
```

### Opciones de configuración

El programa te presentará un menú interactivo donde podrás seleccionar:

- ✅ **Minúsculas (a-z)** - Letras en minúscula
- ✅ **Mayúsculas (A-Z)** - Letras en mayúscula
- ✅ **Dígitos (0-9)** - Números del 0 al 9
- ⬜ **Símbolos (!@#$%...)** - Caracteres especiales
- ⬜ **Excluir caracteres ambiguos** - Evita caracteres como `1lI0Oo8B` que pueden confundirse

*Por defecto, se seleccionan minúsculas, mayúsculas y dígitos.*

### Ejemplo de uso

```
 ::::::::  :::::::: :::::::::    :::     ::::::::  ::::::::  
:+:    :+::+:    :+::+:    :+: :+: :+:  :+:    :+::+:    :+: 
+:+       +:+    +:++:+    +:++:+   +:+ +:+       +:+        
:#:       +#+    +:++#++:++#++#++:++#++:+#++:++#+++#++:++#++ 
+#+   +#+#+#+    +#++#+      +#+     +#+       +#+       +#+ 
#+#    #+##+#    #+##+#      #+#     #+##+#    #+##+#    #+# 
 ########  ######## ###      ###     ### ########  ########  

? Selecciona las opciones para el pool de caracteres:
  ◉ Usar minúsculas (a-z)
  ◉ Usar mayúsculas (A-Z)
  ◉ Usar dígitos (0-9)
  ◯ Usar símbolos (!@#$%...)
  ◉ Excluir caracteres ambiguos (l,1,O,0...)

📊 Resumen de selección:
  Minúsculas: true
  Mayúsculas: true
  Dígitos: true
  Símbolos: false
  Excluir ambiguos: true

🎯 Pool generado (54 caracteres):
abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789

? Ingresa la longitud de tu contraseña: 20

Contraseña generada: mK7p3nRx9TqW5zFh4YvG
```

## 🏗️ Estructura del código

### Funciones principales

- **`NuevosConjuntos()`**: Define los conjuntos de caracteres disponibles
- **`CrearPool()`**: Genera el pool de caracteres según las opciones seleccionadas
- **`CrearPwsd()`**: Genera la contraseña aleatoria
- **`indiceAleatorioSeguro()`**: Implementa generación de índices sin sesgo modular
- **`menuPool()`**: Interfaz interactiva para selección de opciones
- **`pedirNumero()`**: Solicita la longitud deseada de la contraseña

## 🔒 Seguridad

Este generador implementa las siguientes medidas de seguridad:

1. **Uso de `crypto/rand`**: Fuente de aleatoriedad criptográficamente segura proporcionada por el sistema operativo
2. **Eliminación de sesgo modular**: El algoritmo rechaza valores que podrían introducir sesgo estadístico en la distribución
3. **Sin almacenamiento**: Las contraseñas no se guardan en ningún archivo ni historial

## 📋 Dependencias

- [github.com/AlecAivazis/survey/v2](https://github.com/AlecAivazis/survey) - Para la interfaz interactiva en terminal

## 🤝 Contribuciones

Las contribuciones son bienvenidas. Por favor:

1. Haz fork del proyecto
2. Crea una rama para tu feature (`git checkout -b feature/AmazingFeature`)
3. Commit tus cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push a la rama (`git push origin feature/AmazingFeature`)
5. Abre un Pull Request

## 👨‍💻 Autor

Desarrollado con ❤️ usando Go en Ecuador.