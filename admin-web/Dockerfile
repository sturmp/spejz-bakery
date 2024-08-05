FROM node:22 AS build

ARG VITE_API_URL
ARG VITE_API_AUTH_TOKEN
ENV VITE_API_URL=$VITE_API_URL
ENV VITE_API_AUTH_ADMINTOKEN=$VITE_API_AUTH_ADMINTOKEN

WORKDIR /app
COPY package.json ./
RUN npm install
COPY . ./
RUN npm run build

FROM nginx:1.27
COPY --from=build /app/dist /usr/share/nginx/html
CMD ["nginx", "-g", "daemon off;"]