import json
import random
import psycopg2 as pg
import os

conn = pg.connect(
    host= 'localhost', #localhost replace from our server 
    database= 'moretech5',  #need change this name
    port= 5432,
    user= 'andrey',
    password= 'andrey'
)
cursor = conn.cursor()


        
def OfficeParse(cursor):
    file = './python/parse_data/offices.json'
    with open(file, encoding='utf-8') as file:
        data = json.load(file)
    keys_to_skip = ["openHours","rko","openHoursIndividual","salePointFormat", "hasRamp","distance","kep","myBranch","suoAvailability","network", "salePointCode"]
    table_name = "offices"
    table_name_time = "works_time_offices"
    ord_to_day = {0: 'пн', 1: 'вт', 2: 'ср', 3: 'чт', 4: 'пт', 5: 'сб', 6: 'вс'}
    day_to_ord = {'пн': 0, 'вт': 1, 'ср': 2, 'чт': 3, 'пт': 4, 'сб': 5, 'вс': 6}
    flag_not_work = True
    ans = []

    for obj in data:
    
        if "ДО" not in obj["salePointName"] and "ВТБ" not in obj["salePointName"]:
            continue
        if "network" not in obj.keys():
            obj["network"] = None
        if "salePointCode" not in obj.keys():
            obj["salePointCode"] = None
        
        obj["status"] = obj["status"] == 'открытая'
        #columns = ", ".join(key for key in obj.keys() if key not in keys_to_skip)
        placeholders = ", ".join("%s" for _ in range(len(obj) - len(keys_to_skip)))
        values = [obj[key] for key in obj if key not in keys_to_skip]

        sql = f"INSERT INTO {table_name} (sale_point_name, address, status, office_type, latitude, longitude, metro_station) VALUES ({placeholders})"
        cursor.execute(sql, values)
        conn.commit()

        cursor.execute("SELECT id FROM offices WHERE id=(select max(id) from offices);")
        count = cursor.fetchone()[0]
        
        if  'Не' in obj['openHours'][0]['days']:
            flag_not_work = False
            cursor.execute(f"INSERT INTO {table_name_time} (office_id, weeks_day, work_time_start, work_time_end, lunch_time_start, lunch_time_end, is_legal) VALUES (%s, %s, %s, %s, %s, %s, %s)",(count, None, None, None,None,None, None))
            continue

        if len(obj['openHours']) == 7:
            ans.append(obj['openHours'].copy())
            for d in obj['openHours']: 
                if d["hours"] == 'выходной':
                    cursor.execute(f"INSERT INTO {table_name_time} (office_id, weeks_day, work_time_start, work_time_end, lunch_time_start, lunch_time_end, is_legal) VALUES (%s, %s, %s, %s, %s, %s, %s)",(count,d["days"],None, None, None, None, True))
                    conn.commit()
                    continue
                cursor.execute(f"INSERT INTO {table_name_time} (office_id, weeks_day, work_time_start, work_time_end, lunch_time_start, lunch_time_end, is_legal) VALUES (%s, %s, %s, %s, %s, %s, %s)",(count,d["days"],d["hours"][:5],d["hours"][6:] , None, None, True))
                conn.commit()
            continue

        new_dict_list = list()
        for l in obj['openHours']:
            if l["days"] == 'перерыв':
                for i in range(len(new_dict_list)):
                    new_dict_list[i]['hours'] = new_dict_list[i]['hours'][:5] + '-' + l['hours'][:5] + ', ' + l['hours'][6:] + '-' + new_dict_list[i]['hours'][6:]
            elif len(l["days"]) > 2:
                s_front = l["days"][:2]
                s_back = l["days"][3:]

                for i in range(day_to_ord[s_front], day_to_ord[s_back] + 1):
                    if i + 1 <= len(new_dict_list):
                        continue
                    new_dict_list.append({'days': ord_to_day[i], "hours": l["hours"]})
            else:
                new_dict_list.append({"days": l["days"], "hours": l["hours"]})

            for d in new_dict_list: 
                if d["hours"] == 'выходной':
                    cursor.execute(f"INSERT INTO {table_name_time} (office_id, weeks_day, work_time_start, work_time_end, lunch_time_start, lunch_time_end, is_legal) VALUES (%s, %s, %s, %s, %s, %s, %s)",(count,d["days"],None, None, None, None, True))
                    conn.commit()
                    
                elif len(d["hours"]) == 11:
                    cursor.execute(f"INSERT INTO {table_name_time} (office_id, weeks_day, work_time_start, work_time_end, lunch_time_start, lunch_time_end, is_legal) VALUES (%s, %s, %s, %s, %s, %s, %s)",(count,d["days"],d["hours"][:5],d["hours"][6:] , None, None, True))
                    conn.commit()
                    
                else:
                    cursor.execute(f"INSERT INTO {table_name_time} (office_id, weeks_day, work_time_start, work_time_end, lunch_time_start, lunch_time_end, is_legal) VALUES (%s, %s, %s, %s, %s, %s, %s)",(count,d["days"],d["hours"][:5],d["hours"][19:],d["hours"][6:11] ,d["hours"][13:18] , True))
                    conn.commit()
        
        if  'Не' in obj['openHoursIndividual'][0]['days']:
            flag_not_work = False
            cursor.execute(f"INSERT INTO {table_name_time} (office_id, weeks_day, work_time_start, work_time_end, lunch_time_start, lunch_time_end, is_legal) VALUES (%s, %s, %s, %s, %s, %s, %s)",(count, None, None, None,None,None, None))
            continue

        if len(obj['openHoursIndividual']) == 7:
            ans.append(obj['openHoursIndividual'].copy())
            for d in obj['openHoursIndividual']: 
                if d["hours"] == 'выходной':
                    cursor.execute(f"INSERT INTO {table_name_time} (office_id, weeks_day, work_time_start, work_time_end, lunch_time_start, lunch_time_end, is_legal) VALUES (%s, %s, %s, %s, %s, %s, %s)",(count,d["days"],None, None, None, None, False))
                    conn.commit()
                    continue
                cursor.execute(f"INSERT INTO {table_name_time} (office_id, weeks_day, work_time_start, work_time_end, lunch_time_start, lunch_time_end, is_legal) VALUES (%s, %s, %s, %s, %s, %s, %s)",(count,d["days"],d["hours"][:5],d["hours"][6:] , None, None, False))
                conn.commit()
            continue

        new_dict_list = list()
        for l in obj['openHoursIndividual']:
            if l["days"] == 'перерыв':
                for i in range(len(new_dict_list)):
                    new_dict_list[i]['hours'] = new_dict_list[i]['hours'][:5] + '-' + l['hours'][:5] + ', ' + l['hours'][6:] + '-' + new_dict_list[i]['hours'][6:]
            elif len(l["days"]) > 2:
                s_front = l["days"][:2]
                s_back = l["days"][3:]

                for i in range(day_to_ord[s_front], day_to_ord[s_back] + 1):
                    if i + 1 <= len(new_dict_list):
                        continue
                    new_dict_list.append({'days': ord_to_day[i], "hours": l["hours"]})
            else:
                new_dict_list.append({"days": l["days"], "hours": l["hours"]})

            for d in new_dict_list: 
                if d["hours"] == 'выходной':
                    cursor.execute(f"INSERT INTO {table_name_time} (office_id, weeks_day, work_time_start, work_time_end, lunch_time_start, lunch_time_end, is_legal) VALUES (%s, %s, %s, %s, %s, %s, %s)",(count,d["days"],None, None, None, None, False))
                    conn.commit()
                    
                elif len(d["hours"]) == 11:
                    cursor.execute(f"INSERT INTO {table_name_time} (office_id, weeks_day, work_time_start, work_time_end, lunch_time_start, lunch_time_end, is_legal) VALUES (%s, %s, %s, %s, %s, %s, %s)",(count,d["days"],d["hours"][:5],d["hours"][6:] , None, None, False))
                    conn.commit()
                    
                else:
                    cursor.execute(f"INSERT INTO {table_name_time} (office_id, weeks_day, work_time_start, work_time_end, lunch_time_start, lunch_time_end, is_legal) VALUES (%s, %s, %s, %s, %s, %s, %s)",(count,d["days"],d["hours"][:5],d["hours"][19:],d["hours"][6:11] ,d["hours"][13:18] , True))
                    conn.commit()
        
        count += 1
            #sqltimework = f"INSERT INTO {table_name_work} (id_office, days, timeworks, is_legal) VALUES(%s %s %s)"
    file.close()

def AtmParse(cursor):       
    file = './python/parse_data/atms.json'
    # Открываем файл с указанием кодировки
    with open(file, encoding='utf-8') as file:
        # Загружаем JSON данные
        data = json.load(file)

    data = data['atms']
    keys_to_skip = ["services","allDay"]
    table_name = "atms"
    for obj in data:
        if obj["allDay"] == False:
            obj['work_time_start'] = '10:00'
            obj['work_time_end'] = '22:00'
        else:
            obj['work_time_start'] = '00:00'
            obj['work_time_end'] = '24:00'
        #columns = ", ".join(key for key in obj.keys() if key not in keys_to_skip)
        placeholders = ", ".join("%s" for _ in range(len(obj) - len(keys_to_skip)))
        values = [obj[key] for key in obj if key not in keys_to_skip]
        sql = f"INSERT INTO {table_name} (address, latitude, longitude, work_time_start, work_time_end) VALUES ({placeholders})"
        cursor.execute(sql, values)
        conn.commit()
    file.close()

def ServiceParse(cursor):
    file = './python/parse_data/atms.json'
    # Открываем файл с указанием кодировки
    with open(file, encoding='utf-8') as file:
        # Загружаем JSON данные
        data = json.load(file)
    data = data['atms']

    key_to_pars = "services"
    table_name = "services"
    service_set = set()
    for obj in data:
        dictionary = obj[key_to_pars]
        for key in dictionary.keys():
            service_set.add(key)
    service_list = list(service_set)
    file.close()
    keys_to_skip = ["PP","JP", "rko","salePointFormat", "hasRamp","distance","kep","suoAvailability"]
    service_list = service_list + keys_to_skip
    for s in service_list:
        sql = f"INSERT INTO {table_name} (service_name) VALUES ('{s}')"
        cursor.execute(sql)
        conn.commit()   

def AtmsServices(cursor):
    file = './python/parse_data/atms.json'
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

def OfficesServices(cursor):
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


print("I'm here")
OfficeParse(cursor)
AtmParse(cursor)
ServiceParse(cursor)
AtmsServices(cursor)
OfficesServices(cursor)
#timeParsligal(cursor)
print("Good")

