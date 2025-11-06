# Go Event Booking API

Ein modulares **Event-Buchungssystem in Go**, bestehend aus mehreren Microservices.  
Derzeit umfasst das Projekt folgende Services:

- **User-Service** – Authentifizierung und Registrierung  
- **Event-Service** – Verwaltung von Events und Teilnehmer-Buchungen  

Jeder Service läuft als eigenständiger Container im Docker-Compose-Setup und nutzt eine eigene SQLite-Datenbank (PostgreSQL-Unterstützung ist geplant).

---

## Features

- Saubere Service-Struktur in Go mit `gin-gonic`
- Gemeinsame `.env`-Konfiguration (über `.env.example`)
- Multi-Service-Setup mit **Docker Compose**
- Bereit für zukünftige **Kubernetes-Deployments**
- Optionale **Swagger-Dokumentation** (in Planung)

---

## Installation & Setup

1. **Repository klonen**
   ```
   git clone https://github.com/rearatrox/go-event-booking-api.git
   cd go-event-booking-api
   ```

2. **.env-Dateien anpassen**  
   Erstelle aus der `.env.example` eine `.env`-Datei und passe sie an:
   ```
   cp .env.example .env
   ```

3. **Container starten**
   ```
   docker compose up -d
   ```

4. **Services testen**   
   - Event-Service: [http://localhost:8081 (oderer anderer gewählter Port)](http://localhost:8081)
   - User-Service: [http://localhost:8082 (oderer anderer gewählter Port)](http://localhost:8082) 

---

## API-Routen (aktuell verfügbar)

### 👤 User-Service (`:USERSERVICE_PORT`)
| Methode | Route | Beschreibung |
|----------|--------|--------------|
| `GET` | `API_PFREIX/users` | Listet alle User auf |
| `GET` | `API_PFREIX/users/:id` | Zeigt einen bestimmten User |
| `POST` | `API_PFREIX/users/signup` | Erstellt einen neuen Benutzer |
| `POST` | `API_PFREIX/users/login` | Authentifiziert einen Benutzer |

### Event-Service (`:EVENTSERVICE_PORT`)
| Methode | Route | Beschreibung |
|----------|--------|--------------|
| `POST` | `API_PFREIX/events` | Erstellt ein neues Event |
| `GET` | `API_PFREIX/events` | Listet alle Events auf |
| `GET` | `API_PFREIX/events/:id` | Zeigt ein bestimmtes Event |
| `PUT` | `API_PFREIX/events/:id` | Aktualisiert ein bestimmtes Event (Auth) |
| `DELETE` | `API_PFREIX/events/:id` | Löscht ein bestimmtes Event (Auth) |
| `POST` | `API_PFREIX/events/:id/register` | Bucht ein Event für einen Benutzer (Auth) |
| `DELETE` | `API_PFREIX/events/:id/delete` | Löscht ein Event für einen Benutzer (Auth) |

---

## Swagger

Eine Integration mit [**swaggo/gin-swagger**](https://github.com/swaggo/gin-swagger) ist vorgesehen.  

---

## Kubernetes 

In Zukunft werden Kubernetes-Manifeste unter  
`/k8s/` bereitgestellt, um eine einfache Bereitstellung der Services auf einem Cluster zu ermöglichen.

---

## Grundlage 

Als Grundlage des Projekts diente der folgende Udemy-Kurs: [Go - The Complete Guide](https://www.udemy.com/course/go-the-complete-guide/)
