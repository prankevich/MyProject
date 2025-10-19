Профилирование функции UpdateEmployeesByID с помощью pprof, запросы PUT localhost:7777/api/employees/1 направлялись через POSTMAN 
Добавлена искусственная нагрузка в виде цикла 


Результаты СPU:
Total: 150ms
Разделение нагрузки по слоям cpu.pprof:
Controller :60ms 
Service: 60m s
Repository: 20ms
Из них 10ms — искусственная нагрузка
10ms — SQL-запрос

Результаты HEAP:

Total: 12.7 MB
Основные потребители:
Swagger   и  golang.org/x/net/webdav — около 10 MB
Остальное — системное выделение 

Выводы 

Искусственная нагрузка отразилась в профиле в размере +- 7% CPU
Swagger сильно влияет на потребление памяти

 
