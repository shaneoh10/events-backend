# Events Backend (Go)

This project is a RESTful backend service for managing events, registrations, and attendee-facing event data. It is written in **Golang** and designed to provide a clean, scalable API for event-driven applications (web, mobile, or admin dashboards).

The service is intended to handle:

- Event creation and management
- Event discovery/listing
- Attendee registration and cancellation

---

## API Overview

### Users

- `POST /signup`
	- Create a new user.

- `POST /login`
	- User login.

---

### Events

- `GET /events`
	- List all events.

- `GET /events/{id}`
	- Get one event by ID.

- `POST /events`
	- Create a new event.

- `PUT /events/{id}`
	- Update an existing event.

- `DELETE /events/{id}`
	- Delete an event.

---

### Registrations / Attendees

- `POST events/{id}/register`
	- Register an attendee for an event.

- `DELETE events/{id}/register`
	- Cancel/remove a registration.
