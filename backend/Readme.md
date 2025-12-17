# Backend — Test Webdev Suiten 25

Backend API menggunakan **Go (net/http)**, **Gorilla Mux**, **GORM**, dan dokumentasi **Swagger (swaggo)**.

## Requirements
- Go 1.24+
- PostgreSQL
- (Optional) `swag` CLI untuk generate docs

## Setup Environment
Aplikasi membaca environment dari `.env` (lokal) atau environment variables (deploy).

Minimal yang dibutuhkan:
- Database config (tergantung implementasi `config.LoadConfig()` / `database.Connect()` kamu)
- `PORT` (untuk deploy seperti Railway)

### Railway notes
- Railway memberikan port via env `PORT` (wajib dipakai).
- Untuk Swagger supaya tombol **Try it out** memanggil domain Railway (bukan `localhost`), set:
    - `SWAGGER_HOST=<your-app>.up.railway.app` (tanpa `https://`)

## Run (Local)
Dari folder `backend`:

```
bash go mod tidy go run
```

## Swagger
Swagger UI tersedia di:

- `http://localhost:<port>/swagger/index.html`

### Generate / Update Swagger Docs
Install `swag` CLI sekali:

```
bash go install github.com/swaggo/swag/cmd/swag@latest
```

Generate docs (Windows-friendly):

```
bash cd cmd swag init -g main.go -d .,..\internal\controller,..\internal\routes,..\internal\models,..\internal\service,..\internal\repository,..\internal\util -o ..\docs cd .
```

> Setelah generate, folder `docs/` akan berisi `docs.go`, `swagger.json`, `swagger.yaml`.

## Base URL
Semua endpoint berada di prefix:

- `/api`

## Data Models (DTO)

### DivisionDTO

```
json { "id": 1, "name": "Tukang Kayu" }
```

### EmployeeDTO

```
json { "id": 1, "name": "Nurhadi", "division": { "id": 1, "name": "Tukang Kayu" }, "phone_number": "081234567890", "account_number": "1234567890", "bank_name": "BCA", "shift": "PAGI", "salary": 5000000, "period_salary": "MONTHLY", "daily_salary": 200000, "meal_allowance": 15000, "red_meal_allowance": 0, "overtime": 20000, "red_overtime": 0 
```

### EmployeeInputModifyDTO (Request Body)

```
json { "division_id": 1, "name": "Nurhadi", "phone_number": "081234567890", "account_number": "1234567890", "bank_name": "BCA", "shift": "PAGI", "salary": 5000000, "period_salary": "MONTHLY", "daily_salary": 200000, "meal_allowance": 15000, "red_meal_allowance": 0, "overtime": 20000, "red_overtime": 0 }
```


### AttendanceDTO
> Field `date` bertipe `time.Time`, jadi format JSON-nya **RFC3339** (contoh: `"2025-10-03T00:00:00Z"`).

```
json { "id": 10, "employee_id": 1, "employee_name": "Nurhadi", "division_id": 1, "division_name": "Tukang Kayu", "date": "2025-10-03T00:00:00Z", "home_time": "21:00", "total_overtime": "1+4", "note": "-" }
```

### BulkAttendanceUpsertDTO (Request Body)
```
json [ { "employee_id": 1, "date": "2025-10-03T00:00:00Z", "home_time": "21:00" }, { "employee_id": 2, "date": "2025-10-03T00:00:00Z", "home_time": "21:00" } ]
```

## Endpoints (from Swagger)

### Divisions
#### List divisions
- `GET /api/divisions`
- Response: `200` → `[]DivisionDTO`

Example:

```
bash curl "[http://localhost:8080/api/divisions](http://localhost:8080/api/divisions)"
```

#### Get division by ID
- `GET /api/divisions/{id}`
- Response: `200` → `DivisionDTO`

Example:

```
bash curl "[http://localhost:8080/api/divisions/1](http://localhost:8080/api/divisions/1)"
```

#### Create division
- `POST /api/divisions`
- Body:

```
json { "name": "Tukang Kayu" }
```

- Response: `201` → `DivisionDTO`

Example:

```
bash curl -X POST "[http://localhost:8080/api/divisions](http://localhost:8080/api/divisions)"
-H "Content-Type: application/json"
-d '{"name":"Tukang Kayu"}'
```

#### Update division
- `PUT /api/divisions/{id}`
- Body:

```
json { "name": "Tukang Kayu Senior" 
```

- Response: `200` → `DivisionDTO`

Example:

```
bash curl -X PUT "http://localhost:8080/api/divisions/1"
-H "Content-Type: application/json"
-d '{"name":"Tukang Kayu Senior"}
```

#### Delete division
- `DELETE /api/divisions/{id}`
- Response: `204 No Content`

Example:

```
bash curl -X DELETE "[http://localhost:8080/api/divisions/1](http://localhost:8080/api/divisions/1)"
```

---

### Employees
> Note: path yang dipakai adalah `/api/employee` (singular).

#### List employees
- `GET /api/employee`
- Response: `200` → `[]EmployeeDTO`

Example:

```
bash curl "[http://localhost:8080/api/employee](http://localhost:8080/api/employee)"
```

#### Get employee by ID
- `GET /api/employee/{id}`
- Response: `200` → `EmployeeDTO`

Example:

```
bash curl "[http://localhost:8080/api/employee/1](http://localhost:8080/api/employee/1)"
```

#### Create employee
- `POST /api/employee`
- Body: `EmployeeInputModifyDTO`
- Response: `201` → `EmployeeDTO`

Example:

```
bash curl -X POST "[http://localhost:8080/api/employee](http://localhost:8080/api/employee)"
-H "Content-Type: application/json"
-d '{ "division_id": 1, "name": "Nurhadi", "phone_number": "081234567890", "account_number": "1234567890", "bank_name": "BCA", "shift": "PAGI", "salary": 5000000, "period_salary": "MONTHLY", "daily_salary": 200000, "meal_allowance": 15000, "red_meal_allowance": 0, "overtime": 20000, "red_overtime": 0 }'
```

#### Update employee
- `PUT /api/employee/{id}`
- Body: `EmployeeInputModifyDTO`
- Response: `200` → `EmployeeDTO`

Example:

```
bash curl -X PUT "[http://localhost:8080/api/employee/1](http://localhost:8080/api/employee/1)"
-H "Content-Type: application/json"
-d '{ "division_id": 1, "name": "Nurhadi (Updated)", "phone_number": "081234567890", "account_number": "1234567890", "bank_name": "BCA", "shift": "PAGI", "salary": 5000000, "period_salary": "MONTHLY", "daily_salary": 200000, "meal_allowance": 15000, "red_meal_allowance": 0, "overtime": 20000, "red_overtime": 0 }'
```

#### Delete employee
- `DELETE /api/employee/{id}`
- Response: `204 No Content`

Example:

```
bash curl -X DELETE "[http://localhost:8080/api/employee/1](http://localhost:8080/api/employee/1)"
```

---

### Attendance

#### List attendance by date & division
- `GET /api/attendance?date=YYYY-MM-DD&division_id={id}`
- Query:
    - `date` (required) format `YYYY-MM-DD`
    - `division_id` (required) integer
- Response: `200` → `[]AttendanceDTO`

Example:

```
bash curl "[http://localhost:8080/api/attendance?date=2025-10-03&division_id=1](http://localhost:8080/api/attendance?date=2025-10-03&division_id=1)"
```


#### Bulk upsert attendance (bulk save)
- `POST /api/attendance/bulk`
- Body: `BulkAttendanceUpsertDTO`
- Response: `200` → `{ "updated": <count> }`

Example:

```
bash curl -X POST "http://localhost:8080/api/attendance/bulk" -H "Content-Type: application/json" -d '{ "items": }'
```

##### Overtime rule
Saat mengisi `home_time` (jam pulang), backend menghitung `total_overtime` dengan ketentuan:
- Pulang `17:00` → `1+0` (1 hari kerja + 0 jam lembur)
- Pulang `21:00` → `1+4`
- Pulang `23:00` → `2+1`  
  (setiap 5 jam lembur dihitung sebagai 1 hari kerja tambahan)

## Error Response
Jika terjadi error, response umumnya berbentuk:

```
json { "error": "message" }
```

## Deploy (Railway)
1. Set database env vars sesuai kebutuhan config kamu.
2. Pastikan `PORT` dipakai (Railway akan inject).
3. Set `SWAGGER_HOST` agar Swagger “Try it out” tidak memanggil `localhost`:
    - `SWAGGER_HOST=<your-app>.up.railway.app`
4. Deploy.