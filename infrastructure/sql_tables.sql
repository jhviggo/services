USE main;

CREATE TABLE users (
  id INT AUTO_INCREMENT PRIMARY KEY,
  displayName NVARCHAR(50) NOT NULL,
  username NVARCHAR(100) NOT NULL,
  passwd NVARCHAR(255) NOT NULL,
  createdAt DATETIME NOT NULL,
  INDEX (id)
  UNIQUE (username)
);

CREATE TABLE vehicles (
  id INT AUTO_INCREMENT PRIMARY KEY,
  userId INT NOT NULL,
  model VARCHAR(50) NOT NULL,
  createdAt DATETIME NOT NULL,
  FOREIGN KEY (userId) REFERENCES users(id),
  INDEX (userId)
);

CREATE TABLE refuels (
  id INT AUTO_INCREMENT PRIMARY KEY,
  userId INT NOT NULL,
  vehicleId INT NOT NULL,
  createdAt DATETIME NOT NULL,
  FOREIGN KEY (userId) REFERENCES users(id),
  FOREIGN KEY (vehicleId) REFERENCES vehicles(id),
  INDEX (userId)
);

CREATE TABLE logs (
  id INT AUTO_INCREMENT PRIMARY KEY,
  serviceName VARCHAR(100),
  functionName VARCHAR(100),
  logMessage TEXT,
  createdAt DATETIME NOT NULL,
);


