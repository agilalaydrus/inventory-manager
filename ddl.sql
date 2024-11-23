CREATE DATABASE inventory_management;

USE inventory_management;

CREATE TABLE Products (
                          ID INT AUTO_INCREMENT PRIMARY KEY,
                          name VARCHAR(255) NOT NULL,
                          description TEXT,
                          price DECIMAL(10, 2) NOT NULL,
                          category VARCHAR(100)
);

CREATE TABLE Inventory (
                           product_id INT,
                           quantity INT NOT NULL,
                           location VARCHAR(255),
                           FOREIGN KEY (product_id) REFERENCES Products(ID)
);

CREATE TABLE Orders (
                        order_id INT AUTO_INCREMENT PRIMARY KEY,
                        product_id INT,
                        quantity INT NOT NULL,
                        order_date DATETIME DEFAULT CURRENT_TIMESTAMP,
                        FOREIGN KEY (product_id) REFERENCES Products(ID)
);

INSERT INTO Products (name, description, price, category) VALUES
                                                              ('Product A', 'Description A', 10000, 'Category 1'),
                                                              ('Product B', 'Description B', 20000, 'Category 2');

INSERT INTO Inventory (product_id, quantity, location) VALUES
                                                           (1, 50, 'Warehouse 1'),
                                                           (2, 30, 'Warehouse 2');

INSERT INTO Orders (product_id, quantity) VALUES
                                              (1, 2),
                                              (2, 1);