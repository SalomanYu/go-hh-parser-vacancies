INSERT INTO public."position" (id,name,other_names,"level",parent_id,parsed) VALUES
	 (1415,'Водитель-испытатель','',0,0,false),
	 (1416,'Водитель испытатель мототранспортных средств','Водитель испытатель|Водитель - испытатель|Водитель -испытатель|Водитель- испытатель|Водитель-испытатель мототранспортных средств|Водитель- испытатель мототранспортных средств|Водитель -испытатель мототранспортных средств|Водитель - испытатель мототранспортных средств',0,1415,false),
	 (1417,'Водитель испытатель легковых автомобилей','Водитель-испытатель легковых автомобилей|Водитель - испытатель легковых автомобилей|Водитель -испытатель легковых автомобилей|Водитель- испытатель легковых автомобилей',0,1415,false),
	 (1418,'Водитель испытатель пассажирского транспорта','Водитель-испытатель пассажирского транспорта|Водитель - испытатель пассажирского транспорта|Водитель- испытатель пассажирского транспорта|Водитель -испытатель пассажирского транспорта|Водитель испытатель пасажирского транспорта|Водитель-испытатель пасажирского транспорта|Водитель - испытатель пасажирского транспорта|Водитель- испытатель пасажирского транспорта|Водитель -испытатель пасажирского транспорта',0,1415,false),
	 (1419,'Водитель испытатель грузового транспорта','Водитель-испытатель грузового транспорта|Водитель - испытатель грузового транспорта|Водитель -испытатель грузового транспорта|Водитель- испытатель грузового транспорта',0,1415,false),
	 (1420,'Испытатель','',0,0,false),
	 (1421,'Испытатель двигателей','Инженер-испытатель|Инженер по испытаниям|Испытатель-механик двигателей|Инженер-испытатель (двигателей и агрегатов)',0,1420,false),
	 (1422,'Инженер-испытатель','Инженер испытатель',0,4483,false),
	 (1423,'Испытатель на герметичность','',0,1420,false),
	 (1424,'Контролер качества на линии','Контролер качества на производственной линии',0,1420,false);
INSERT INTO public."position" (id,name,other_names,"level",parent_id,parsed) VALUES
	 (1425,'Контролер сварочных работ','Контролер сварочных и наплавочных работ|Контролер ОТК сварочных работ|Контролер ОТК (сварочных работ)|Контролер ОТК сварочных работ (машиностроение)',0,1420,false),
	 (1426,'Контролер ОТК','',0,0,false),
	 (1427,'Дефектоскопист','Контролер ОТК / Инженер ОТК|Контролер качества (ОТК)|Инженер-дефектоскопист',0,1426,false),
	 (1428,'Инженер по гарантии','',0,1428,false),
	 (1429,'Механик по металлоконструкциям','',0,1429,false),
	 (1430,'Техник','Механик по металлоконструкциям',0,1430,false),
	 (1431,'Антикоррозийщик','',0,1431,false),
	 (1432,'Дизайнер автомобилестроения','Дизайнер средств транспорта',0,1432,false),
	 (1433,'Специалист окрасочного производства','',0,1433,false),
	 (1434,'Оператор окрасочно-сушильных линий','',0,1434,false);
INSERT INTO public."position" (id,name,other_names,"level",parent_id,parsed) VALUES
	 (1435,'Станочник','',0,0,false),
	 (1436,'Станочник металлообрабатывающих станков','Разнорабочий-станочник|Станочник широкого профиля|Станочник широкого профиля на токарные и фрезерные станки',0,1435,false),
	 (1437,'Станочник широкого профиля','Станочник|Станочники и наладчики металлообрабатывающих станков',0,1435,false),
	 (1438,'Наладчик','',0,0,false),
	 (1439,'Наладчик металлообрабатывающих станков','Наладчик производственного оборудования|Оператор-наладчик|Инженер-наладчик|Наладчик станков |Наладчик технологического оборудования|Оператор - наладчик|Механик-наладчик |Оператор наладчик|Механик-наладчик оборудования',0,1438,false),
	 (1440,'Наладчик оборудования','Оператор - наладчик|Механик-наладчик |Оператор наладчик|Наладчик|Механик-наладчик оборудования|Оператор-наладчик|Инженер-наладчик',0,1438,false),
	 (1442,'Наладчик производственного оборудования','Инженер-наладчик производственного оборудования|Механик-наладчик производственного оборудования|Оператор производственной линии / Наладчик оборудования',0,1438,false),
	 (1443,'Наладчик станков с ЧПУ','Оператор-наладчик станка ЧПУ|Наладчик - оператор станков с ЧПУ|Наладчик (оператор) станков с ЧПУ',0,1438,false),
	 (1444,'Наладчик шлифовального оборудования','',0,1438,false),
	 (1445,'Наладчик-штамповщик','Штамповчик',0,1438,false);
INSERT INTO public."position" (id,name,other_names,"level",parent_id,parsed) VALUES
	 (1446,'Оператор  ','',0,0,false),
	 (1447,'Оператор автоматических и полуавтоматических линий станков и установок','Оператор металлорежущих станков-автоматов',0,1446,false),
	 (1448,'Оператор металлообрабатывающего оборудования','',0,1446,false),
	 (1449,'Оператор лазерного станка','',0,1446,false),
	 (1450,'Оператор линии производства','',0,1446,false),
	 (1451,'Оператор листогибочного пресса','',0,1446,false),
	 (1452,'Оператор станка','',0,1446,false),
	 (1453,'Оператор токарного станка','',0,1446,false),
	 (1454,'Кладовщик','',0,0,false),
	 (1455,'Кладовщик-грузчик','Кладовщик- грузчик|Кладовщик -грузчик|Кладовщик - грузчик|Кладовщик грузчик',0,1454,false);
INSERT INTO public."position" (id,name,other_names,"level",parent_id,parsed) VALUES
	 (1456,'Комплектовщик-транспортировщик','Комплектовщик- транспортировщик|Комплектовщик -транспортировщик|Комплектовщик - транспортировщик|Комплектовщик транспортировщик',0,1454,false),
	 (1457,'Кладовщик-комплектовщик','Кладовщик- комплектовщик|Кладовщик -комплектовщик|Кладовщик комплектовщик|Кладовщик - комплектовщик',0,1454,false),
	 (1458,'Кладовщик-контролер','Кладовщик- контролер|Кладовщик -контролер|Кладовщик контролер|Кладовщик - контролер',0,1454,false),
	 (1459,'Кладовщик-наборщик','Кладовщик -наборщик|Кладовщик- наборщик|Кладовщик - наборщик|Кладовщик наборщик ',0,1454,false),
	 (1460,'Кладовщик-оператор','Кладовщик оператор|Кладовщик -оператор|Кладовщик- оператор|Кладовщик - оператор',0,1454,false),
	 (1461,'Кладовщик-приемщик','Кладовщик- приемщик|Кладовщик -приемщик|Кладовщик - приемщик|Кладовщик приемщик',0,1454,false),
	 (1462,'Кладовщик склада','Кладовщик / сотрудник на склад|Сотрудник на склад|Работник склада (кладовщик)|Специалист склада / кладовщик|Кладовщик на склад|Сотрудник склада',0,1454,false),
	 (1463,'Кладовщик-экспедитор','Кладовщик- экспедитор|Кладовщик -экспедитор|Кладовщик - экспедитор|Кладовщик экспедитор|Кладовщик/экспедитор',0,1454,false),
	 (1464,'Менеджер склада','Менеджер на склад',0,1454,false),
	 (1465,'Оператор логистических работ','',0,1446,false);
INSERT INTO public."position" (id,name,other_names,"level",parent_id,parsed) VALUES
	 (1466,'Оператор склада','',0,1446,false),
	 (1467,'Оператор-упаковщик','',0,1446,false),
	 (1468,'Приемщик','',0,0,false),
	 (1469,'Приемщик-сдатчик','Приемщик - сдатчик|Приемщик- сдатчик|Приемщик -сдатчик|Приемщик сдатчик',0,1468,false),
	 (1470,'Водитель','',0,0,false),
	 (1471,'Водитель погрузчика','Водитель-погрузчика|Водитель- погрузчика|Водитель -погрузчика|Водитель - погрузчика|Водитель погрузчика (карщик)|Водитель погрузчика-грузчик',0,1470,false),
	 (1472,'Водитель технологического транспорта','',0,1470,false),
	 (1473,'Водитель-грузчик','Водитель- грузчик|Водитель -грузчик|Водитель - грузчик|Водитель грузчик',0,1470,false),
	 (1474,'Водитель-менеджер','',0,1470,false),
	 (1475,'Водитель ПРТ','Водитель-ПРТ|Водитель- ПРТ|Водитель -ПРТ|Водитель - ПРТ',0,1470,false);
INSERT INTO public."position" (id,name,other_names,"level",parent_id,parsed) VALUES
	 (1476,'Перегонщик автомобилей','Техник-перегонщик|Техник-перегонщик автомобилей|Приемщик-перегонщик автомобилей|Техник перегонщик автомобилей',0,1470,false),
	 (1477,'Грузчик','',0,0,false),
	 (1478,'Грузчик-комплектовщик','Водитель-грузчик|Грузчик на склад|Грузчик- комплектовщик|Грузчик -комплектовщик|Грузчик - комплектовщик|Грузчик комплектовщик|Грузчик-комплектовщик на склад',0,1477,false),
	 (1479,'Грузчик-стропальщик','Грузчик -стропальщик|Грузчик- стропальщик|Грузчик - стропальщик|Грузчик стропальщик',0,1477,false),
	 (1480,'Грузчик-такелажник','Грузчик- такелажник|Грузчик -такелажник|Грузчик - такелажник|Грузчик такелажник ',0,1477,false),
	 (1481,'Грузчик-транспортировщик','Грузчик- транспортировщик|Грузчик -транспортировщик|Грузчик - транспортировщик|Грузчик транспортировщик|Транспортировщик-грузчик
|Транспортировщик- грузчик|Транспортировщик -грузчик|Транспортировщик - грузчик|Транспортировщик грузчик',0,1477,false),
	 (1482,'Грузчик-экспедитор','Грузчик -экспедитор|Грузчик- экспедитор|Грузчик - экспедитор|Грузчик экспедитор',0,1477,false),
	 (10,'Диспетчер','',0,0,false),
	 (1483,'Диспетчер-логист','Распределитель работ (диспетчер)|Диспетчер -логист|Диспетчер- логист|Диспетчер - логист|Диспетчер логист|Логист-диспетчер|Логист -диспетчер|Логист- диспетчер|Логист - диспетчер|Логист диспетчер|Диспетчер отдела логистики',0,10,false),
	 (1484,'Диспетчер производственно-диспетчерской службы','Диспетчер производства|Диспетчер производственно- диспетчерской службы|Диспетчер производственно -диспетчерской службы|Диспетчер производственно - диспетчерской службы|Диспетчер производственно диспетчерской службы',0,10,false);
INSERT INTO public."position" (id,name,other_names,"level",parent_id,parsed) VALUES
	 (1485,'Логист','Менеджер по логистике|Специалист по логистике|Транспортный логист|Диспетчер-логист|Логист-диспетчер|Специалист отдела логистики|Менеджер отдела логистики|Специалист по транспортной логистике|Менеджер по транспортной логистике|Логист-аналитик',0,4451,false),
	 (1486,'Логист на склад','Логист склада',0,1485,false),
	 (1487,'Менеджер по закупкам','Специалист по закупкам|Менеджер отдела снабжения|Менеджер отдела закупок',0,1485,false),
	 (1488,'Менеджер по импорту','Менеджер по работе с импортом',0,1485,false),
	 (1489,'Менеджер по международным перевозкам','Менеджер по международной логистике',0,1485,false),
	 (1490,'Менеджер по перевозкам','Менеджер по перевозкам/логист',0,1485,false),
	 (27,'Водитель','',0,22,false),
	 (1491,'Кузнец','',0,0,false),
	 (1492,'Кузнец-штамповщик','Кузнец- штамповщик|Кузнец -штамповщик|Кузнец - штамповщик|Кузнец штамповщик',0,1491,false),
	 (1493,'Кузнец на молотах и прессах','Кузнец на молотах и прессах (на горячих участках работ)',0,1491,false);
INSERT INTO public."position" (id,name,other_names,"level",parent_id,parsed) VALUES
	 (1494,'Резчик','',0,1491,false),
	 (1495,'Резчик металла на ножницах и прессах','',0,1494,false),
	 (1496,'Резчик на пилах, ножовках и станках','Резчик на пилах,ножовках и станках',0,1494,false),
	 (1497,'Резьбонарезчик','Резьбонарезчик на спецстанках|Резьбонарезчик на специальных станках',0,1494,false),
	 (1498,'Гибщик','',0,0,false),
	 (1499,'Гибщик-вальцовщик','Гибщик листового метала|Вальцовщик - Гибщик листового металла|Гибщик -вальцовщик|Гибщик- вальцовщик|Гибщик - вальцовщик|Гибщик вальцовщик|Вальцовщик/гибщик',0,1498,false),
	 (1500,'Гибщик труб','',0,1498,false),
	 (1501,'Инженер технолог','Инженер-технолог|Инженер -технолог|Инженер- технолог|Инженер - технолог',0,1501,false),
	 (1502,'Термист','',0,4505,false),
	 (1503,'Техник-конструктор','Техник- конструктор|Техник -конструктор|Техник - конструктор|Техник конструктор',0,1503,false);
INSERT INTO public."position" (id,name,other_names,"level",parent_id,parsed) VALUES
	 (1504,'Литейщик','',0,0,false),
	 (1505,'Литейщик пласмас','',0,1504,false),
	 (1506,'Литейщик металов и сплавов','',0,1504,false),
	 (13267,'Литейщик пластмасс','Формовщик-литейщик пластмасс|Литейщик изделий из пластмасс',0,13267,false),
	 (16478,'Литейщик металлов и сплавов','',0,1504,false),
	 (1507,'Химик-технолог','',0,0,false),
	 (1508,'Лаборант химического анализа','Химик- технолог|Химик -технолог|Химик - технолог|Химик технолог',0,1507,false),
	 (1509,'Лаборант спектрального анализа','',0,1507,false),
	 (1510,'Лаборант металлографического анализа','',0,1507,false),
	 (1511,'Лаборант неразрушающего метода контроля','',0,1507,false);
INSERT INTO public."position" (id,name,other_names,"level",parent_id,parsed) VALUES
	 (1512,'Лаборант механических испытаний','Лаборант по физико-механическим испытаниям',0,1507,false),
	 (1513,'Бригадир','Бригадир объектов',0,1507,false),
	 (1514,'Слесарь','',0,0,false),
	 (1515,'Слесарь механосборочных работ','Слесарь механосборочных работ (МСР)',0,1514,false),
	 (1516,'Слесарь-моторист','Слесарь -моторист|Слесарь- моторист|Слесарь - моторист|Слесарь моторист|Моторист/слесарь',0,1514,false),
	 (1517,'Слесарь МСР','Слесарь механосборочных работ|Слесарь механосборочных работ (МСР)',0,1514,false),
	 (1518,'Слесарь-универсал','Слесарь -универсал|Слесарь- универсал|Слесарь - универсал|Слесарь универсал',0,1514,false),
	 (1519,'Автомаляр','',0,0,false),
	 (1520,'Автомаляр-подготовщик','Автомаляр -подготовщик|Автомаляр- подготовщик|Автомаляр - подготовщик',0,1519,false),
	 (1521,'Маляр по металлу','',0,1519,false);
INSERT INTO public."position" (id,name,other_names,"level",parent_id,parsed) VALUES
	 (1522,'Колорист','Маляр-колорист|Колорист автоэмалей',0,1522,false),
	 (1523,'Художник-аэрографист','Художник- аэрографист|Художник -аэрографист|Художник - аэрографист|Художник аэрографист',0,1523,false),
	 (1524,'Полировщик','Полировщик автомобилей',0,1524,false),
	 (1525,'Слесарь по ремонту автомобилей','Автослесарь',0,1514,false),
	 (1526,'Автомеханик','',0,0,false),
	 (1527,'Автомеханик-автослесарь','Мастер ремонтной зоны по техническому обслуживанию, ремонту автотранспортных средств (АТС) и их компонентов|Автомеханик -автослесарь|Автомеханик- автослесарь|Автомеханик - автослесарь|Автомеханик автослесарь|Автомеханик / Автослесарь|Автомеханик (автослесарь)
|Автомеханик / автослесарь|Автослесарь-автомеханик
|Автослесарь -автомеханик|Автослесарь- автомеханик|Автослесарь - автомеханик|Автослесарь автомеханик',0,1526,false),
	 (1528,'Электромеханик по средствам автоматики и приборам технологического оборудования','',0,1526,false),
	 (1529,'Мехатроник','Специалист по мехатронным системам автомобиля|Автомехатроник по ремонту автотранспортных средств (АТС)|Автомехатроник|Специалист по мехатронике в автомобилестроении|Техник-мехатроник',0,1529,false),
	 (1530,'Электромонтер','Дежурный электромонтер|Электромонтер по ремонту и обслуживанию электрооборудования',0,1530,false),
	 (16653,'Супервайзер','Супервайзер|Супервайзер по бронированию|Супервайзер отдела приема и размещения гостей|Супервайзер/администратор|Супервайзер отдела гостиничного хозяйства|Супервайзер областных продаж|Супервайзер гостевой зоны|Supervisor',0,0,false);
INSERT INTO public."position" (id,name,other_names,"level",parent_id,parsed) VALUES
	 (11789,'Хостес','Хостес|Администратор службы приема и размещения|Менеджер по встрече гостей',0,11789,false),
	 (332,'Администратор','Администратор|Системный администратор (системный инженер)|Системный администратор|Системный инженер|Системный администратор Linux|Сетевой / системный администратор|System Administrator|System Engineer|Администратор баз данных|Дежурный администратор баз данных|Администратор баз данных SQL|Database administrator|Database administrator MS SQL|DBA - Database Administrator|Database administrator / DBA|Oracle DBA|Администратор сайта|Администратор сайта интернет-магазина|Site Administrator|Администратор 1С|Администратор IT и 1С|Системный администратор 1С|Администратор службы приема и размещения|Администратор службы приема гостей|Администратор на ресепшен|Ночной администратор|Администратор гостиницы|Администратор в сеть апарт-отелей|Администратор отеля|Администратор ресепшен|Администратор на ресепшен отеля|Администратор на базу отдыха|Администратор в мини-отель|Администратор в сеть апартаментов',0,0,false),
	 (16654,'Администратор гостиницы','Администратор гостиницы|Администратор|Администратор службы приема и размещения|Администратор службы приема гостей|Администратор на ресепшен|Ночной администратор|Администратор в сеть апарт-отелей|Администратор отеля|Администратор ресепшен|Администратор на ресепшен отеля|Администратор на базу отдыха|Администратор в мини-отель|Администратор в сеть апартаментов',0,332,false),
	 (16655,'Администратор жилого комплекса','Администратор жилого комплекса|Администратор гостиницы|Администратор|Администратор службы приема и размещения|Администратор службы приема гостей|Администратор на ресепшен|Ночной администратор|Администратор в сеть апарт-отелей|Администратор отеля|Администратор ресепшен|Администратор на ресепшен отеля|Администратор на базу отдыха|Администратор в мини-отель|Администратор в сеть апартаментов',0,332,false),
	 (16656,'Менеджер','Менеджер|Менеджер по работе с горячими заявками|Менеджер по работе с клиентами',0,0,false),
	 (16657,'Менеджер гостиничного сервиса','Менеджер гостиничного сервиса',0,16656,false),
	 (16658,'Менеджер по работе с клиентами гостиницы','Менеджер по работе с клиентами гостиницы',0,16656,false),
	 (16659,'Руководитель/управляющий гостиничного комплекса/сети гостиниц','Старший администратор гостиницы|Старший администратор в отель|Старший Администратор Комплекса Отдыха',0,16656,false),
	 (5561,'Специалист по санаторно-курортному делу','Специалист по санаторно-курортному делу',0,0,false),
	 (16660,'Горничная','Горничная|Горничная в отель|Горничная номерного фонда',0,0,false);
INSERT INTO public."position" (id,name,other_names,"level",parent_id,parsed) VALUES
	 (16661,'Работник по приему и размещению гостей','Работник по приему и размещению гостей',0,0,false),
	 (16662,'Супервайзер номерного фонда','Супервайзер номерного фонда|Супервайзер|Супервайзер по бронированию|Супервайзер отдела приема и размещения гостей|Супервайзер/администратор|Супервайзер отдела гостиничного хозяйства|Супервайзер областных продаж|Супервайзер гостевой зоны|Supervisor',0,0,false),
	 (16663,'Швейцар','Швейцар|Bellman|Подносчик багажа|Швейцар в гостинице|Ночной швейцар|Багажист|Bell-Boy|Швейцар-портье',0,0,false),
	 (16664,'Аппаратчик бельевых сушильных установок','Аппаратчик бельевых сушильных установок',0,0,false),
	 (16665,'Гладильщик','Гладильщик|Гладильщик белья|Гладильщица|Гладильщик в химчистку',0,0,false),
	 (16666,'Комплектовщик белья','Комплектовщик белья',0,0,false),
	 (16667,'Контролер качества обработки одежды и белья','Контролер качества обработки одежды и белья',0,0,false),
	 (16668,'Оператор прачечной самообслуживания','Оператор прачечной самообслуживания',0,0,false),
	 (16669,'Оператор стиральных машин','Оператор стиральных машин|Сотрудник прачечной|Оператор прачечной|Оператор стиральных машин в прачечную',0,0,false),
	 (16670,'Отжимщик белья на центрифугах','Отжимщик белья на центрифугах',0,0,false);
INSERT INTO public."position" (id,name,other_names,"level",parent_id,parsed) VALUES
	 (16671,'Подготовитель белья для глажения','Подготовитель белья для глажения',0,0,false),
	 (16672,'Приготовитель стиральных растворов','Приготовитель стиральных растворов',0,0,false),
	 (16673,'Аппаратчик химической чистки','Аппаратчик химической чистки|Пятновыводчик химической чистки и Аква чистки|Аппаратчик - пятновыводчик|Технолог химчистки',0,0,false),
	 (16674,'Аппаратчик чистки ковровых изделий','Аппаратчик чистки ковровых изделий',0,0,false),
	 (16675,'Аппаратчик чистки пухо-перовых изделий','Аппаратчик чистки пухо-перовых изделий',0,0,false),
	 (16676,'Комплектовщик изделий','Комплектовщик изделий',0,0,false),
	 (16677,'Контролер качества обработки изделий','Контролер качества обработки изделий',0,0,false),
	 (16678,'Красильщик','Красильщик',0,0,false),
	 (16679,'Отделочник головных уборов','Отделочник головных уборов',0,0,false),
	 (16680,'Отпарщик-прессовщик','Отпарщик-прессовщик',0,0,false);
INSERT INTO public."position" (id,name,other_names,"level",parent_id,parsed) VALUES
	 (16681,'Пятновыводчик','Пятновыводчик|Аппаратчик - пятновыводчик |Технолог химчистки',0,0,false),
	 (9998,'Сушильщик изделий','Сушильщик изделий',0,9998,false),
	 (16682,'Менеджер по туризму','Менеджер по туризму|Специалист по туризму|Турагент|Онлайн-турагент|Турагент (онлайн)|Менеджер по туризму / турагент',0,0,false),
	 (16683,'Менеджер по внутреннему туризму','Менеджер по внутреннему туризму|Менеджер по внутреннему и въездному туризму|Менеджер по туризму|Специалист по туризму|Турагент|Онлайн-турагент|Турагент (онлайн)|Менеджер по туризму / турагент',0,16682,false),
	 (16684,'Менеджер по въездному туризму','Менеджер по въездному туризму|Менеджер по внутреннему и въездному туризму|Менеджер по бронированию в отдел внутреннего и въездного туризма|Менеджер отдела по въездному туризму|Менеджер по туризму|Специалист по туризму|Турагент|Онлайн-турагент|Турагент (онлайн)|Менеджер по туризму / турагент',0,16682,false),
	 (16685,'Менеджер по выездному туризму','Менеджер по выездному туризму|Менеджер по туризму|Специалист по туризму|Турагент|Онлайн-турагент|Турагент (онлайн)|Менеджер по туризму / турагент',0,16682,false),
	 (16686,'Тревел-консультант','Тревел-консультант|Тревел-консультант (удаленно)',0,0,false),
	 (16687,'Менеджер по продажам (туров)','Менеджер по продажам (туров)|Менеджер по продажам туров (удаленно)|Менеджер по продажам туров',0,0,false),
	 (16688,'Менеджер по продажам в сфере туризма','Менеджер по продажам в сфере туризма',0,0,false),
	 (1441,'Наладчик КИПиА','Инженер-наладчик КИПиА |Слесарь-наладчик КИП и А|Наладчик кипиа ',0,1438,true);
