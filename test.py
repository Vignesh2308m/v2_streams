import duckdb


db=duckdb.connect("test.db")

print(db.sql("select * from test.test").df)


