-- SQL (Structured Query Language) is a language to interact with relational database management system (DBMS)
-- CRUD -> 4 main operation
-- 2 types of databases: Relational Database (SQL) vs Non-Relational (NoSQL)
-- Relational Database -> organize data into one or more tables, unique key identifies each row
-- -- Non-Relational -> any database that's not relational, i.e. Documents (JSON, XML, etc), key-value stores, graphs, etc

-- Column -> single attributes
-- Row -> single instance

-- Every table should have a primary key as unique value
-- primary key can be anything, number, string
-- surrogate key is basically a key that has no mapping to anything in the real world
-- so a lot of ID is a surrogate key then?
-- natural key is key that has a mapping or has a purpose in the real world, not just in database i.e. NIP
-- foreign key stores of the primary key of a row in another database table i.e. a basketball player table
-- have a team_id as foreign key
-- foreign key is the one that define relationship between table
-- one table can have more than one foreign key
-- composite key is a primary key that consist of more than 1 column

-- best practice for sql is snake case

CREATE DATABASE backend_intro;

-- Main Datatype
-- INT              integer
-- DECIMAL (M,N)    float, M being total digits and N being total digits of decimal number (?)
-- VARCHAR (l)      string with length of l
-- BLOB             binary large object
-- DATE             yyyy-mm-dd
-- TIMESTAMP        yyyy-mm-dd hh:mm:ss

-- usually people write it in all uppercase so stick with it
CREATE TABLE;

-- student_id INT PRIMARY KEY is equal to PRIMARY KEY(student_id)
CREATE TABLE student (
    student_id INT PRIMARY KEY,
    name VARCHAR(20),
    major VARCHAR(20)
);

DESCRIBE student;


DROP TABLE student;

ALTER TABLE student ADD gpa DECIMAL(3,2);

ALTER TABLE student DROP COLUMN gpa;

-- inserting data
INSERT INTO student VALUES(
    1,
    "Jack",
    'Biology'
);

INSERT INTO student VALUES(
    2,
    "Kate",
    'Sociology'
);

INSERT INTO student(student_id, name) VALUES(
    3,
    "Claire"
);

-- the sequence doesn't matter
INSERT INTO student(name) VALUES(
    "Jim"
);

INSERT INTO student VALUES(
    5,
    "Jack",
    'Biology'
);

INSERT INTO student VALUES(
    6,
    "Mike",
    'Computer Science'
);
INSERT INTO student VALUES(
    7,
    "Claude",
    'Chemistry'
);

SELECT * FROM student;


DROP TABLE student;

-- Constraint
-- NOT NULL
-- UNIQUE
-- DEFAULT
-- AUTO_INCREMENT

CREATE TABLE student(
    student_id INT AUTO_INCREMENT,
    name VARCHAR(20) NOT NULL,
    major VARCHAR(20) UNIQUE,
    -- gpa DECIMAL(3,2) DEFAULT 2.00,
    PRIMARY KEY(student_id)
);

-- primary key basically just not null and unique column

-- update and deleting
UPDATE student
-- SET major = "Bio"
-- WHERE  major = "Biology";
-- WHERE student_id = 4;
-- SET major = "Biochemistry"
-- WHERE major = "Bio" OR major = "Chemistry";
SET name = 'Tom', major = "undecided"
WHERE student_id = 1;
-- WHERE is optional

SELECT * FROM student;

DELETE FROM student; -- this is to delete all record
-- WHERE student_id = 5;
-- WHERE name = "Tom" AND major = "undecided";

-- basic queries
-- SELECT name, major
SELECT student.name, student.major
FROM student
ORDER BY name DESC
-- ORDER BY name ASC;
-- ORDER BY major, student_id; --order by major then student_id
-- ORDER BY major, student_id DESC; --order by major ascending the student_id descending
LIMIT 2;

SELECT *
FROM student
-- WHERE major = "biology" OR major = "chemistry"; -- wait, sql data isn't case sensitive too biology = Biology
WHERE name IN ("claire", "kate", "mike");

DROP TABLE student;
-- <, >, <=, >=, =, <>, AND, OR
-- <> is not equal to lol

-- COMPANY DATABASE USECASE
CREATE TABLE employee(
    emp_id INT PRIMARY KEY,
    first_name VARCHAR(40),
    last_name VARCHAR(40),
    birth_date DATE,
    sex VARCHAR(1),
    salary INT,
    -- can't define super_id and branch_id as foreign key yet since the table aren't created
    super_id INT,
    branch_id INT
);

CREATE TABLE branch(
    branch_id INT PRIMARY KEY,
    branch_name VARCHAR(40),
    mgr_id INT,
    mgr_start_date DATE,
    FOREIGN KEY(mgr_id) REFERENCES employee(emp_id) ON DELETE SET NULL
);

ALTER TABLE employee
ADD FOREIGN KEY(branch_id)
REFERENCES branch(branch_id)
ON DELETE SET NULL;

ALTER TABLE employee
ADD FOREIGN KEY(super_id)
REFERENCES employee(emp_id)
ON DELETE SET NULL;


CREATE TABLE client(
    client_id INT PRIMARY KEY,
    client_name VARCHAR(40),
    branch_id INT,
    FOREIGN KEY(branch_id) REFERENCES branch(branch_id) ON DELETE SET NULL
);

CREATE TABLE works_with(
    emp_id INT,
    client_id INT,
    total_sales INT,
    PRIMARY KEY(emp_id, client_id),
    FOREIGN KEY(emp_id) REFERENCES employee(emp_id) ON DELETE CASCADE,
    FOREIGN KEY(client_id) REFERENCES client(client_id) ON DELETE CASCADE
);

CREATE TABLE branch_supplier(
    branch_id INT,
    supplier_name VARCHAR(40),
    supply_type VARCHAR(40),
    PRIMARY KEY(branch_id, supplier_name),
    FOREIGN KEY(branch_id) REFERENCES branch(branch_id) ON DELETE CASCADE
);

-- Corporate
INSERT INTO employee VALUES(100, 'David', 'Wallace', '1967-11-17', 'M', 250000, NULL, NULL);

INSERT INTO branch VALUES(1, 'Corporate', 100, '2006-02-09');

UPDATE employee
SET branch_id=1
WHERE emp_id = 100;

INSERT INTO employee VALUES(101, 'Jan', 'Levinson', '1961-05-11', 'F', 110000, 100, 1);

-- Scranton
INSERT INTO employee VALUES(102, 'Michael', 'Scott', '1964-03-15', 'M', 75000, 100, NULL);

INSERT INTO branch VALUES(2, 'Scranton', 102, '1992-04-06');

UPDATE employee
SET branch_id = 2
WHERE emp_id = 102;

INSERT INTO employee VALUES(103, 'Angela', 'Martin', '1971-06-25', 'F', 63000, 102, 2);
INSERT INTO employee VALUES(104, 'Kelly', 'Kapoor', '1980-02-05', 'F', 55000, 102, 2);
INSERT INTO employee VALUES(105, 'Stanley', 'Hudson', '1958-02-19', 'M', 69000, 102, 2);

-- Stamford
INSERT INTO employee VALUES(106, 'Josh', 'Porter', '1969-09-05', 'M', 78000, 100, NULL);

INSERT INTO branch VALUES(3, 'Stamford', 106, '1998-02-13');

UPDATE employee
SET branch_id = 3
WHERE emp_id = 106;

INSERT INTO employee VALUES(107, 'Andy', 'Bernard', '1973-07-22', 'M', 65000, 106, 3);
INSERT INTO employee VALUES(108, 'Jim', 'Halpert', '1978-10-01', 'M', 71000, 106, 3);


-- BRANCH SUPPLIER
INSERT INTO branch_supplier VALUES(2, 'Hammer Mill', 'Paper');
INSERT INTO branch_supplier VALUES(2, 'Uni-ball', 'Writing Utensils');
INSERT INTO branch_supplier VALUES(3, 'Patriot Paper', 'Paper');
INSERT INTO branch_supplier VALUES(2, 'J.T. Forms & Labels', 'Custom Forms');
INSERT INTO branch_supplier VALUES(3, 'Uni-ball', 'Writing Utensils');
INSERT INTO branch_supplier VALUES(3, 'Hammer Mill', 'Paper');
INSERT INTO branch_supplier VALUES(3, 'Stamford Lables', 'Custom Forms');

-- CLIENT
INSERT INTO client VALUES(400, 'Dunmore Highschool', 2);
INSERT INTO client VALUES(401, 'Lackawana Country', 2);
INSERT INTO client VALUES(402, 'FedEx', 3);
INSERT INTO client VALUES(403, 'John Daly Law, LLC', 3);
INSERT INTO client VALUES(404, 'Scranton Whitepages', 2);
INSERT INTO client VALUES(405, 'Times Newspaper', 3);
INSERT INTO client VALUES(406, 'FedEx', 2);

-- WORKS_WITH
INSERT INTO works_with VALUES(105, 400, 55000);
INSERT INTO works_with VALUES(102, 401, 267000);
INSERT INTO works_with VALUES(108, 402, 22500);
INSERT INTO works_with VALUES(107, 403, 5000);
INSERT INTO works_with VALUES(108, 403, 12000);
INSERT INTO works_with VALUES(105, 404, 33000);
INSERT INTO works_with VALUES(107, 405, 26000);
INSERT INTO works_with VALUES(102, 406, 15000);
INSERT INTO works_with VALUES(105, 406, 130000);


SELECT *
FROM EMPLOYEE;

SELECT *
FROM client;

SELECT *
FROM employee
ORDER BY salary DESC;

SELECT *
FROM employee
ORDER BY sex, first_name, last_name;


SELECT *
FROM employee
LIMIT 5;

SELECT first_name AS forename, last_name AS surname
FROM employee;

SELECT DISTINCT sex AS gender
FROM employee;

-- FUNCTION
SELECT COUNT(emp_id)
FROM employee;

-- THIS IS WRONG, WHY?
SELECT COUNT(emp_id)
FROM employee
WHERE super_id <> NULL;

SELECT COUNT(super_id)
FROM employee;

SELECT COUNT(emp_id)
-- SELECT *
FROM employee
WHERE sex = 'F' AND birth_date >= '1971-01-01';

SELECT AVG(salary)
FROM employee
WHERE sex = 'M';

SELECT SUM(salary)
FROM employee;

SELECT COUNT(sex), sex
FROM employee
GROUP BY sex;

SELECT works_with.total_sales
FROM works_with, employee
WHERE employee.sex = 'M';

SELECT emp_id, SUM(total_sales)
FROM works_with
GROUP BY emp_id;

SELECT client_id, SUM(total_sales)
FROM works_with
GROUP BY client_id;

-- WILDCARD
-- % =  any # characters, _ = one character
SELECT *
FROM client
WHERE client_name LIKE '%LLC';

SELECT *
FROM branch_supplier
WHERE supplier_name LIKE '%LAB%';

SELECT *
FROM employee
WHERE birth_date LIKE '____-10%';

SELECT *
FROM client
WHERE client_name LIKE '%school%';

-- UNION
SELECT employee.*, branch.branch_name
FROM employee, branch
WHERE employee.branch_id = branch.branch_id;

SELECT first_name
FROM employee
UNION
SELECT branch_name
FROM branch
UNION
SELECT client_name
FROM client;

-- to use union, both n_column from select should be equal, same datatype
SELECT salary
FROM employee
UNION
SELECT total_sales
FROM works_with;

-- JOIN
INSERT INTO branch VALUES(4, 'Buffalo', NULL, NULL);

SELECT employee.emp_id, employee.first_name, branch.branch_name
FROM employee
JOIN branch
ON employee.emp_id = branch.mgr_id;

-- INNER JOIN, LEFT JOIN, RIGHT JOIN
-- FULL OUTER JOIN CAN'T BE DONE IN MYSQL IT'S LEFT JOIN + RIGHT JOIN

-- NESTED QUERIES

SELECT employee.first_name
FROM employee
JOIN works_with
ON employee.emp_id = works_with.emp_id
WHERE works_with.total_sales > 30000;

SELECT employee.first_name, employee.last_name
FROM employee
WHERE employee.emp_id IN (
    SELECT works_with.emp_id
    FROM works_with
    WHERE works_with.total_sales > 30000
);


SELECT client.client_name
FROM client
WHERE client.branch_id IN(
    SELECT branch_id
    FROM employee
    -- WHERE employee.first_name = 'Michael' AND employee.last_name = 'Scott'
    WHERE employee.emp_id = 102
);

-- if the foreign key are part of primary key ON DELETE CASCADE since primary key can't null
-- else ON DELETE SET NULL

-- TRIGGER
CREATE TABLE trigger_test (
    message VARCHAR(100)
);

-- U NEED TO DO IT IN SQL TO CHANGE DELIMITER FIRST BEFORE YOU CAN USE TRIGGER

DELIMITER $$
CREATE
    TRIGGER my_trigger1 BEFORE INSERT
    ON employee
    FOR EACH ROW BEGIN
        -- INSERT INTO trigger_test VALUES('added new employee');
        INSERT INTO trigger_test VALUES(NEW.first_name);
    END$$
DELIMITER ;

INSERT INTO employee VALUES(109, 'Oscar', 'Martinez', '1968-02-19', 'M', 69000, 106, 3);

SELECT * FROM trigger_test;

INSERT INTO employee VALUES(110, 'Kevin', 'Malone', '1978-02-19', 'M', 69000, 106, 3);

DELETE FROM employee WHERE emp_id = 110;

-- HOW TO SEE LIST OF TABLES, LIST OF TRIGGERS THAT WE'VE BEEN CREATED

DELIMITER $$
CREATE
    TRIGGER my_trigger2 BEFORE INSERT
    -- BEFORE | AFTER
    -- INSERT | UPDATE | DELETE
    ON employee
    FOR EACH ROW BEGIN
        IF NEW.sex = 'M' THEN
            INSERT INTO trigger_test VALUES('ADDED MALE EMPLOYEE');
        ELSEIF NEW.sex = 'F' THEN
            INSERT INTO trigger_test VALUES('ADDED MALE EMPLOYEE');
        ELSE
            INSERT INTO trigger_test VALUES('ADDED OTHER EMPLOYEE');
        END IF;
    END$$
DELIMITER ;

DROP TRIGGER my_trigger;

-- ER (Entity Relationship) DIAGRAM 
-- Entity -> an object we want to model & store information about (table basically (?))
-- Attributes -> specific information about an entity
-- Primary Key -> an attribute(s) that uniquely identify an entry in the database table (usually underlined)
-- Composite Attribute -> an attribute that can be broken up into sub-attributes (name->fname, lname)
-- Multi-valued Attribute -> example one person can have more than one child, children is multi-valued attributes
-- Derived Attributes -> an attribute that can be derived from the other attributes
-- i.e. has_honors in student can be derived from gpa
-- Multiple Entities -> you can define more than one entity in a diagram (more than one table)
-- Relationships -> defines a relationship between to entities
-- Total Participation -> All members must participate in the relationship (all class need to have a student)
-- Partial Participation -> (Not all student need to take all class)
-- Relationship Attributs -> grades is a relationship attributes for student that take a class
-- Relationship Cardinality -> the number of instances of an entity from a relation that can be associated with the relation
-- 1:1, 1:N, N:M
-- Weak Entity -> entity that cannot be uniquely identified by it's attribute alone
-- Identifying relationship -> a relationship that serves to uniquely identify the weak entity

