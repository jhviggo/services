INSERT INTO users (displayName, username, passwd, createdAt)
VALUES (
  'Viggo',
  'viggolicious',
  '$2a$10$QTTpTQPN16XjTNrLceWF0.e/BezkpsIN9u5qR4VfQ4NkASySmt3jK',
  NOW()
);

INSERT INTO vehicles (userId, model, estimatedKML, createdAt)
VALUES (
  (SELECT id FROM users WHERE username = 'viggolicious'),
  'Suzuki GS 500',
  25,
  NOW()
);

INSERT INTO refuels (userId, vehicleId, totalKM, tripKM, liters, cost, currency, createdAt)
VALUES (
  (SELECT id FROM users WHERE username = 'viggolicious'),
  (SELECT id FROM vehicles WHERE model = 'Suzuki GS 500' LIMIT 1),
  32031.9, -- total
  315.7, -- trip
  13.03, -- liters
  23.44, -- cost
  'eur',
  NOW()
), (
  (SELECT id FROM users WHERE username = 'viggolicious'),
  (SELECT id FROM vehicles WHERE model = 'Suzuki GS 500' LIMIT 1),
  32400.1, -- total
  367.1, -- trip
  14.73, -- liters
  27.09, -- cost
  'eur',
  NOW()
);
