import json
import psycopg2 as pg



conn = pg.connect(

    host='localhost', #localhost replace from our server 
    database ='000001_init',  #need change this name
    port = 5432,
    user ='postgres',
    password = 'root'
)
cursor = conn.cursor()

file = 'atms.json'
# Открываем файл с указанием кодировки
with open(file, encoding='utf-8') as file:
    # Загружаем JSON данные
    data = json.load(file)

data = data['atms']

#keys_to_skip = ["services","allDay"]
table_name = "atms_services"


for obj in data:
    d = obj["services"]
    latitude = obj["latitude"]
    longitude = obj["longitude"]
    cursor.execute(f"SELECT id FROM atms WHERE latitude = '{latitude}' AND longitude = '{longitude}';")
    id_atms = cursor.fetchone()[0]
    
    for f in d:
        cursor.execute(f"SELECT id FROM services WHERE service_name = '{f}';")
        id_service = cursor.fetchone()[0]
        if d[f]['serviceCapability'] == "SUPPORTED":
            sql = f"INSERT INTO atms_services (atm_id, service_id) VALUES {(id_atms, id_service)}"
            cursor.execute(sql)
            conn.commit()
    


    