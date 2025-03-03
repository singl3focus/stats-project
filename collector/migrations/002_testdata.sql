-- Начало транзакции для атомарности
BEGIN;

-- Заполнение services
INSERT INTO services (name, description)
VALUES 
  ('Flight Booking', 'Авиабилеты'),
  ('Hotel Reservation', 'Отели'),
  ('Travel Insurance', 'Путешествия')
ON CONFLICT DO NOTHING;

-- Заполнение users
INSERT INTO users (name)
VALUES 
  ('Санечка'),
  ('Макс'),
  ('Даня'),
  ('Соня')
ON CONFLICT DO NOTHING;

-- Заполнение stats (только после заполнения users и services!)
INSERT INTO stats (user_id, service_id, count)
VALUES
  (1, 1, 5),   -- Санечка использовал поиск авиабилетов 5 раз
  (1, 2, 3),   -- Санечка искал отели 3 раза
  (2, 1, 12),  -- Макс часто искал авиабилеты
  (3, 3, 2),   -- Даня оформил страховку 2 раза
  (4, 1, 7),   -- Соня активно использует поиск авиабилетов
  (4, 2, 4),   -- Соня искала отели 4 раза
  (4, 3, 1)    -- Соня оформила страховку 1 раз
ON CONFLICT (user_id, service_id) 
DO UPDATE SET count = excluded.count;

-- Фиксация изменений
COMMIT;