
import psycopg2 as pg
import random 

conn = pg.connect(

    host='localhost', #localhost replace from our server 
    database ='000001_init',  #need change this name
    port = 5432,
    user ='postgres',
    password = 'root'
)
cursor = conn.cursor()


cursor.execute(f"SELECT id FROM offices")
id_office = cursor.fetchall()
cursor.execute(f"SELECT id FROM services")
id_services = cursor.fetchall()

for key in id_office:
    for key_services in id_services:
        
        if bool(random.getrandbits(1)) == True:
            sql = f"INSERT INTO offices_services (office_id, service_id) VALUES {(key[0], key_services[0])}"
            cursor.execute(sql)
            conn.commit()
            
    
