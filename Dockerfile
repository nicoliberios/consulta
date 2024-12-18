# Usa una imagen base de nginx para servir archivos estáticos
FROM nginx:alpine

# Copia el archivo index.html al contenedor
COPY index.html /usr/share/nginx/html/index.html

# Expone el puerto 80
EXPOSE 80

# El contenedor se ejecutará con nginx por defecto, que ya está configurado para servir archivos en el directorio /usr/share/nginx/html
