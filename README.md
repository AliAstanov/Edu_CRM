# EduCRM — O‘quv markazlari uchun CRM tizimi

**EduCRM** — bu o‘quv markazlari, repetitorlik kurslari va ta’lim tashkilotlari uchun mo‘ljallangan, zamonaviy va kengaytiriladigan **CRM tizimi**. Platforma yordamida o‘quvchilarni, to‘lovlarni, dars davomati va ustozlarni boshqarish osonlashadi. Shuningdek, tizimda sun’iy intellekt yordamida tahlil va tavsiyalar ham qo‘shiladi.

---

## ✨ Asosiy imkoniyatlar

- 👨‍🎓 O‘quvchilarni ro‘yxatdan o‘tkazish va boshqarish
- 👨‍🏫 Ustozlar va fanlar moduli
- 👥 Guruhlar va dars jadvali
- ✅ Davomat (attendances) yuritish
- 💵 To‘lovlar monitoringi va eslatmalar
- 📊 Hisobotlar: oylik, dars qoldirish statistikasi
- 🤖 AI yordamida tavsiyalar va tahlil (GPT integratsiyasi)

---

## 🧱 Texnologiyalar

| Qism       | Texnologiya            |
|------------|------------------------|
| Til        | Go (Golang)            |
| Framework  | Gin Web Framework      |
| Ma’lumotlar bazasi | PostgreSQL      |
| AI         | OpenAI API (GPT-4o)    |
| Konteynerlash | Docker, Docker Compose |
| Frontend   | Keyinchalik qo‘shiladi |

---

## 🗂️ Ma’lumotlar bazasi jadvallari

### 🧍‍♂️ 1. `students` — O‘quvchilar

| Field      | Type      | Description                           |
|------------|-----------|---------------------------------------|
| id         | UUID      | Primary key                           |
| full_name  | TEXT      | O‘quvchi ismi                         |
| phone      | TEXT      | Telefon raqami                        |
| group_id   | UUID      | FK → `groups.id`                      |
| created_at | TIMESTAMP | Ro‘yxatdan o‘tgan vaqt                |

---

### 👨‍🏫 2. `teachers` — Ustozlar

| Field      | Type      | Description                           |
|------------|-----------|---------------------------------------|
| id         | UUID      | Primary key                           |
| full_name  | TEXT      | Ustoz ismi                            |
| phone      | TEXT      | Telefon raqami                        |
| created_at | TIMESTAMP | Qo‘shilgan vaqt                       |

---

### 📚 3. `subjects` — Fanlar

| Field      | Type      | Description                           |
|------------|-----------|---------------------------------------|
| id         | UUID      | Primary key                           |
| name       | TEXT      | Masalan: Ingliz tili, Matematika      |

---

### 🏫 4. `groups` — O‘quvchilar guruhi

| Field      | Type      | Description                           |
|------------|-----------|---------------------------------------|
| id         | UUID      | Primary key                           |
| name       | TEXT      | Guruh nomi                            |
| subject_id | UUID      | FK → `subjects.id`                    |
| teacher_id | UUID      | FK → `teachers.id`                    |
| created_at | TIMESTAMP | Yaratilgan vaqt                       |

---

### ✅ 5. `attendances` — Davomat

| Field      | Type      | Description                           |
|------------|-----------|---------------------------------------|
| id         | UUID      | Primary key                           |
| student_id | UUID      | FK → `students.id`                    |
| group_id   | UUID      | FK → `groups.id`                      |
| date       | DATE      | Dars kuni                             |
| status     | TEXT      | `present` yoki `absent`               |

---

### 💵 6. `payments` — To‘lovlar

| Field      | Type      | Description                           |
|------------|-----------|---------------------------------------|
| id         | UUID      | Primary key                           |
| student_id | UUID      | FK → `students.id`                    |
| amount     | INT       | To‘langan summa                       |
| month      | TEXT      | `YYYY-MM` shaklida (masalan: 2025-07) |
| paid_at    | TIMESTAMP | To‘lov vaqti                          |

---

### 🤖 7. `ai_logs` — Sun’iy intellekt so‘rovlar

| Field       | Type    | Description                            |
|-------------|---------|----------------------------------------|
| id          | UUID    | Primary key                            |
| input_type  | TEXT    | `student_report`, `payment_trend` ...  |
| input_data  | JSONB   | AI’ga uzatilgan malumotlar             |
| response    | TEXT    | AI javobi                              |
| created_at  | TIMESTAMP | So‘rov vaqti                         |

---

## ⚙️ Loyihani ishga tushirish (Local Development)

### 1. PostgreSQL sozlash

```bash
CREATE EXTENSION IF NOT EXISTS "pgcrypto";
