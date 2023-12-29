USE main;

CREATE TABLE users (
  id INT AUTO_INCREMENT PRIMARY KEY,
  displayName NVARCHAR(50) NOT NULL,
  username NVARCHAR(100) NOT NULL,
  passwd NVARCHAR(255) NOT NULL,
  createdAt DATETIME NOT NULL,
  INDEX (id),
  UNIQUE (username)
);

CREATE TABLE vehicles (
  id INT AUTO_INCREMENT PRIMARY KEY,
  userId INT NOT NULL,
  model VARCHAR(50) NOT NULL,
  estimatedKML INT,
  createdAt DATETIME NOT NULL,
  FOREIGN KEY (userId) REFERENCES users(id),
  INDEX (userId)
);

CREATE TABLE refuels (
  id INT AUTO_INCREMENT PRIMARY KEY,
  userId INT NOT NULL,
  vehicleId INT NOT NULL,
  totalKM FLOAT NOT NULL,
  tripKM FLOAT NOT NULL,
  liters FLOAT NOT NULL,
  cost FLOAT NOT NULL,
  currency VARCHAR(10) NOT NULL,
  createdAt DATETIME NOT NULL,
  FOREIGN KEY (userId) REFERENCES users(id),
  FOREIGN KEY (vehicleId) REFERENCES vehicles(id),
  INDEX (userId)
);
