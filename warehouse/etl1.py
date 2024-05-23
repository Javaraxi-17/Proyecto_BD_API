from pyspark.sql import SparkSession

# Create a Spark session
spark = SparkSession.builder \
    .appName("Hive Integration") \
    .enableHiveSupport() \
    .getOrCreate()

# Create a Hive table
create_table_query = """
CREATE TABLE IF NOT EXISTS employee (
    id INT,
    name STRING,
    age INT,
    salary DOUBLE
)
STORED AS PARQUET
"""
spark.sql(create_table_query)

# Example data
data = [
    (1, 'John Doe', 30, 50000.0),
    (2, 'Jane Smith', 25, 60000.0),
    (3, 'Alice Johnson', 28, 70000.0)
]

# Define schema
columns = ["id", "name", "age", "salary"]

# Create DataFrame
df = spark.createDataFrame(data, columns)

# Show the DataFrame
df.show()

# Write the DataFrame to the Hive table
df.write.mode("append").insertInto("employee")

# Query the Hive table to verify data insertion
spark.sql("SELECT * FROM employee").show()

# Stop the Spark session
spark.stop()
