import csv
from datetime import datetime

import pymongo
import psycopg2

def read_csv(filename):
	with open(filename, mode='r') as csv_file:
		csv_reader = csv.DictReader(csv_file)
		data = []
		for row in csv_reader:
			data.append(row)

	return data

def execute_sql(connection, query, records=[]):
	cursor = connection.cursor()
	if records:
		cursor.executemany(query, records)
	else:
		cursor.execute(query)

	connection.commit()
	cursor.close()

def restore_orders(connection):
	filename = "Test task - Postgres - orders.csv"
	data = read_csv(filename)
	create_table_query = '''CREATE TABLE IF NOT EXISTS orders ( id INT PRIMARY KEY NOT NULL, created_at TIMESTAMP NOT NULL, order_name VARCHAR(50) NOT NULL, customer_id VARCHAR(50) NOT NULL ); '''
	execute_sql(connection, create_table_query)
	records = [(int(row["id"]), row["created_at"], row["order_name"], row["customer_id"]) for row in data]
	sql_insert_query = """ INSERT INTO orders (id, created_at, order_name, customer_id) VALUES (%s,%s,%s, %s) """
	execute_sql(connection, sql_insert_query, records)
	connection.commit()

def restore_order_items(connection):
	filename = "Test task - Postgres - order_items.csv"
	data = read_csv(filename)

	create_table_query = '''CREATE TABLE IF NOT EXISTS order_items ( id INT PRIMARY KEY NOT NULL, order_id INT NOT NULL, price_per_unit DECIMAL DEFAULT NULL, quantity INT NOT NULL, product VARCHAR(100) NOT NULL, FOREIGN KEY (order_id) REFERENCES orders(id) ); '''
	execute_sql(connection, create_table_query)
	records = [(int(row["id"]), row["order_id"], float(row["price_per_unit"]) if row["price_per_unit"] else None, row["quantity"], row["product"]) for row in data]
	sql_insert_query = """ INSERT INTO order_items (id, order_id, price_per_unit, quantity, product) VALUES (%s,%s,%s, %s, %s) """
	execute_sql(connection, sql_insert_query, records)
	connection.commit()

def restore_deliveries(connection):
	filename = "Test task - Postgres - deliveries.csv"
	data = read_csv(filename)

	create_table_query = '''CREATE TABLE IF NOT EXISTS deliveries ( id INT PRIMARY KEY NOT NULL, order_item_id INT NOT NULL, delivered_quantity INT NOT NULL, FOREIGN KEY (order_item_id) REFERENCES order_items(id) ); '''
	execute_sql(connection, create_table_query)
	records = [(int(row["id"]), row["order_item_id"], int(row["delivered_quantity"])) for row in data]
	sql_insert_query = """ INSERT INTO deliveries (id, order_item_id, delivered_quantity) VALUES (%s,%s,%s) """
	execute_sql(connection, sql_insert_query, records)
	connection.commit()

def restore_customers(mongo_db):
	filename = "Test task - Mongo - customers.csv"
	data = read_csv(filename)
	customers_collection = mongo_db["customers"]
	x = customers_collection.insert_many(data)

def restore_customer_companies(mongo_db):
	filename = "Test task - Mongo - customer_companies.csv"
	data = read_csv(filename)
	companies_collection = mongo_db["customer_companies"]
	x = companies_collection.insert_many(data)


if __name__ == "__main__": 
	myclient = pymongo.MongoClient("mongodb://localhost:27017/")
	mongo_db = myclient["geektest"]
	restore_customers(mongo_db)
	restore_customer_companies(mongo_db)

	try:
		connection = psycopg2.connect(user="root",password="geektest",host="127.0.0.1",port="5432",database="geektest")
		restore_orders(connection)
		restore_order_items(connection)
		restore_deliveries(connection)

	except (Exception, psycopg2.Error) as error :
		print ("Error while connecting to PostgreSQL", error)
	finally:
			if(connection):
				connection.close()
				print("PostgreSQL connection is closed")