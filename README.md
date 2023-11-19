## Readme Top

<a name="readme-top"></a>

<!-- PROJECT LOGO -->
<br />
<div align="center">

<h3 align="center">ecommerce crud</h3>

  <p align="center">
    Developer Guide To Start The Project!
  </p>
</div>

<!-- Seed Database -->

# Seed Database

To seed

- fill up the .env file with proper configuration
- go to root of the application
- run the following command in the terminal

```
$ make seed
```

<!-- Start Application -->

# Start Application

To start the application locally

- fill up the .env file with proper configuration
- go to root of the application
- run the following command in the terminal

```
$ make serve
```

- if you want to run with docker-compose

* fill up the .env file with proper configuration
* go to root of the application
* run the following commands

```
$ docker-compose build
$ docker-compose up
```

# SQL

```
  DROP TABLE IF EXISTS product_stocks;
	DROP TABLE IF EXISTS products;
	DROP TABLE IF EXISTS brands;
	DROP TABLE IF EXISTS categories;
	DROP TABLE IF EXISTS suppliers;

	CREATE TABLE IF NOT EXISTS brands (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		name VARCHAR(255) NOT NULL,
		status_id INTEGER NOT NULL,
		created_at BIGINT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS categories (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		name VARCHAR(255) NOT NULL,
		parent_id UUID,
		sequence INTEGER,
		status_id INTEGER NOT NULL,
		created_at BIGINT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS suppliers (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL,
		phone VARCHAR(20),
		status_id INTEGER NOT NULL,
		is_verified_supplier BOOLEAN NOT NULL,
		created_at BIGINT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS products (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		name VARCHAR(255) NOT NULL,
		description TEXT,
		specifications TEXT,
		brand_id UUID REFERENCES brands(id) NOT NULL,
		category_id UUID REFERENCES categories(id) NOT NULL,
		supplier_id UUID REFERENCES suppliers(id) NOT NULL,
		unit_price NUMERIC NOT NULL,
		discount_price NUMERIC,
		tags VARCHAR(255)[],
		status_id INTEGER NOT NULL,
		created_at BIGINT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS product_stocks (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		product_id UUID REFERENCES products(id) NOT NULL,
		stock_quantity INTEGER NOT NULL,
		updated_at BIGINT NOT NULL
	);
```

# Project: ecommerce-crud

# 📁 Collection: Brand

## End-point: Create brand

### Method: POST

> ```
> http://localhost:5000/api/brands
> ```

### Body (**raw**)

```json
{
  "name": "Apple1",
  "status_id": 1
}
```

⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Get brand

### Method: GET

> ```
> http://localhost:5000/api/brands/eac728f6-8057-4525-8754-1bad2e44b5ec
> ```

⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Update Brand

### Method: PUT

> ```
> http://localhost:5000/api/brands/eac728f6-8057-4525-8754-1bad2e44b5ec
> ```

### Body (**raw**)

```json
{
  "name": "Apple2",
  "status_id": 1
}
```

⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Delete Brand

### Method: DELETE

> ```
> http://localhost:5000/api/brands/28228a5c-a926-4709-86fb-6fba77e99d75
> ```

⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Get brands

### Method: GET

> ```
> http://localhost:5000/api/brands?page=1&limit=2
> ```

### Query Params

| Param | value |
| ----- | ----- |
| page  | 1     |
| limit | 2     |

⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

# 📁 Collection: Product

## End-point: Create product

### Method: POST

> ```
> http://localhost:5000/api/products
> ```

### Body (**raw**)

```json
{
  "name": "mlenovo think v2",
  "description": "abc",
  "brand_id": "5f2dc58e-d3a8-4580-b4fb-0e72d93f0afe",
  "category_id": "8ace9e3f-3bca-4deb-8128-e0f67b0c0924",
  "supplier_id": "5d96a2df-370b-4afd-a7c4-cfcc1e7241d2",
  "unit_price": 50.05,
  "discount_price": 12.54,
  "tags": ["abc", "xyz"],
  "status_id": 1,
  "stock_quantity": 100
}
```

⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Get product

### Method: GET

> ```
> http://localhost:5000/api/products/9b097254-60c8-42d7-9fcf-5857a64ef360
> ```

⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Update product

### Method: PUT

> ```
> http://localhost:5000/api/products/9b097254-60c8-42d7-9fcf-5857a64ef360
> ```

### Body (**raw**)

```json
{
  "name": "mlenovo think v2",
  "description": "abc",
  "brand_id": "5f2dc58e-d3a8-4580-b4fb-0e72d93f0afe",
  "category_id": "8ace9e3f-3bca-4deb-8128-e0f67b0c0924",
  "supplier_id": "5d96a2df-370b-4afd-a7c4-cfcc1e7241d2",
  "unit_price": 50.05,
  "discount_price": 12.54,
  "tags": ["abc", "xyz"],
  "status_id": 1,
  "stock_quantity": 100
}
```

⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Delete product

### Method: DELETE

> ```
> http://localhost:5000/api/products/9b097254-60c8-42d7-9fcf-5857a64ef360
> ```

⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Get products

### Method: GET

> ```
> http://localhost:5000/api/products?page=1&limit=20
> ```

### Query Params

| Param | value |
| ----- | ----- |
| page  | 1     |
| limit | 20    |

⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

# 📁 Collection: Supplier

## End-point: Create supplier

### Method: POST

> ```
> http://localhost:5000/api/suppliers
> ```

### Body (**raw**)

```json
{
  "name": "Habibur Rahman",
  "email": "habib@gmail.com",
  "phone": "01723403221",
  "status_id": 1,
  "is_verified_supplier": true
}
```

⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Get supplier

### Method: GET

> ```
> http://localhost:5000/api/suppliers/ffb80702-6d3c-400b-bf6b-08ccf85017cf
> ```

⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Update supplier

### Method: PUT

> ```
> http://localhost:5000/api/suppliers/ffb80702-6d3c-400b-bf6b-08ccf85017cf
> ```

### Body (**raw**)

```json
{
  "name": "Habibur",
  "email": "habib@gmail.com",
  "phone": "01723403221",
  "status_id": 1,
  "is_verified_supplier": true
}
```

⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Delete supplier

### Method: DELETE

> ```
> http://localhost:5000/api/suppliers/b10c8de7-f6a8-4fc9-a5da-61420f3096dc
> ```

⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Get suppliers

### Method: GET

> ```
> http://localhost:5000/api/suppliers?page=1&limit=5
> ```

### Query Params

| Param | value |
| ----- | ----- |
| page  | 1     |
| limit | 5     |

⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

# 📁 Collection: Category

## End-point: Create category

### Method: POST

> ```
> http://localhost:5000/api/categories
> ```

### Body (**raw**)

```json
{
  "name": "Mobile 5",
  "parent_id": "fef438e9-2c04-4e12-961d-d35e2d75e5cd",
  "status_id": 1
}
```

⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Get category

### Method: GET

> ```
> http://localhost:5000/api/categories/4bdce61a-4fd4-4218-910c-08f9a313ad72
> ```

⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Update categories

### Method: PUT

> ```
> http://localhost:5000/api/categories/460c3893-d368-4cdb-b407-cd267c32e36d
> ```

### Body (**raw**)

```json
{
  "name": "Phone",
  "status_id": 1
}
```

⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Delete Brand

### Method: DELETE

> ```
> http://localhost:5000/api/categories/4bdce61a-4fd4-4218-910c-08f9a313ad72
> ```

⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Get categories

### Method: GET

> ```
> http://localhost:5000/api/categories?page=1&limit=5
> ```

### Query Params

| Param | value |
| ----- | ----- |
| page  | 1     |
| limit | 5     |

⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Get category tree

### Method: GET

> ```
> http://localhost:5000/api/categories/tree
> ```

⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

---
