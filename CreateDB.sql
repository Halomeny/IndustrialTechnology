CREATE TABLE Users (
    Username VARCHAR(255) PRIMARY KEY,
    Password VARCHAR(255) NOT NULL,
    DeletionMark BOOLEAN
);

CREATE TABLE Customers (
    Customer_ID SERIAL PRIMARY KEY,
    Name VARCHAR(50),
    Phone VARCHAR(15) UNIQUE NOT NULL,
    Email VARCHAR(50),
    DeletionMark BOOLEAN,
    User_name VARCHAR(255),
    FOREIGN KEY (User_name) REFERENCES Users(Username)
);

CREATE TABLE Products (
    Name VARCHAR(100) UNIQUE PRIMARY KEY,
    Price DECIMAL(10, 2),
    Description Text,
    DeletionMark BOOLEAN,
    Category_ID INT
);

CREATE TABLE Product_Categories (
    Category_ID SERIAL PRIMARY KEY,
    Category_Name VARCHAR(50),
    DeletionMark BOOLEAN
);

CREATE TABLE Employees (
    Employee_ID SERIAL PRIMARY KEY,
    Name VARCHAR(50),
    Phone VARCHAR(15) UNIQUE NOT NULL, 
    Position VARCHAR(50),
    DeletionMark BOOLEAN,
    User_name VARCHAR(255),
    FOREIGN KEY (User_name) REFERENCES Users(Username)
);


CREATE TABLE Orders (
    Order_ID SERIAL PRIMARY KEY,
    Customer_ID INT,
    Product_Name VARCHAR(100),
    Employee_ID INT,
    Quantity INT,
    Order_Date DATE,
    DeletionMark BOOLEAN,
    FOREIGN KEY (Customer_ID) REFERENCES Customers(Customer_ID),
    FOREIGN KEY (Product_Name) REFERENCES Products(Name),
    FOREIGN KEY (Employee_ID) REFERENCES Employees(Employee_ID)
);


