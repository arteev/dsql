//Create the databases for exsamples 
//MUST Envroment variables: ISC_USER, ISC_PASSWORD


SET SQL DIALECT 3;

create database '/tmp/dsql.exsample1.fdb'
PAGE_SIZE = 4096
DEFAULT CHARACTER SET UTF8;
commit;

connect '/tmp/dsql.exsample1.fdb';

CREATE TABLE Employee (
  id INTEGER not null,  
  NAME varchar(255) default NULL,
  EmpNo varchar(13) default NULL,
  Birth varchar(255),
  Salary varchar(50) default NULL,
  PRIMARY KEY (id)
);

INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (1,'Jasper Hobbs','1662051221599','1488672867',20464);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (2,'Daniel Sykes','1666020429599','1454200362',58055);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (3,'Lucius Pena','1608021141199','1441394009',88194);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (4,'Channing Fuentes','1636071470099','1483911609',51984);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (5,'Porter Hooper','1653020100499','1478594691',28450);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (6,'Gray Delaney','1685020652099','1465660439',34489);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (7,'Brenden Jordan','1672051445099','1446324166',66523);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (8,'Magee Santiago','1612061784499','1483808639',91002);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (9,'Colin Tate','1667082314199','1489007353',32905);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (10,'Kirk Sullivan','1621010984499','1434493248',42662);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (11,'Sylvester Haynes','1613082107599','1473307179',37051);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (12,'Ryder Rush','1648050734999','1489028357',46668);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (13,'Jack Odonnell','1634111431799','1483490738',28258);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (14,'Craig Hess','1621030200499','1488277959',91844);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (15,'Matthew Crane','1675011679899','1446846160',12840);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (16,'Clinton William','1669030117099','1442208368',80229);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (17,'Kamal Henry','1635040158699','1436856755',58428);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (18,'Nero Durham','1603030776699','1472700685',25592);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (19,'Chaim Lara','1605030689699','1469148126',60512);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (20,'Hunter Boyle','1691120340299','1480533982',30129);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (21,'Ulric Bridges','1665040608299','1439605950',25121);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (22,'Zeph Wagner','1681011434399','1443390181',46066);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (23,'Bruce Carver','1688031135899','1451201033',96494);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (24,'Zahir Glass','1678042183499','1435067037',34524);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (25,'Macaulay Solis','1670111174499','1459045487',65250);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (26,'Channing Conley','1656042833899','1456635107',48600);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (27,'Jason Yang','1609111559199','1443568617',9616);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (28,'Laith Dickerson','1621101548299','1443012132',35746);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (29,'Colt Banks','1632031620599','1467247024',24698);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (30,'Grady Gallagher','1643020258199','1431761141',74136);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (31,'Devin Duran','1682061506099','1467773268',47087);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (32,'Barrett Castaneda','1686062087999','1446637622',51729);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (33,'Lee Spence','1622020596399','1467283065',44499);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (34,'Rafael Higgins','1666053060799','1442750666',84857);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (35,'Lester Mccray','1646122243599','1479938702',23779);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (36,'Lucius Delaney','1694122331599','1429088093',88825);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (37,'Scott Mcleod','1639091725699','1439759248',44938);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (38,'Asher Joyner','1669030592899','1469165482',80019);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (39,'Jameson Maynard','1653070200799','1442882228',80228);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (40,'Chadwick Harmon','1653061676599','1461810936',7558);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (41,'Steel Keller','1690102552399','1448550163',10704);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (42,'Zane Mcfadden','1681050747199','1466805070',9557);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (43,'Hayes Mcclain','1687082389799','1438477141',84225);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (44,'Axel Meadows','1689011118799','1450241998',93113);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (45,'Aladdin Dorsey','1654021339299','1446012638',84610);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (46,'Wyatt Hernandez','1680090213599','1465736052',83013);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (47,'John Jacobson','1616080638899','1443693592',49291);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (48,'Plato Goodman','1608110264799','1441999298',71070);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (49,'Yardley Padilla','1609030613199','1439336836',17095);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (50,'Griffith Ayers','1607111245099','1482724331',74899);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (51,'Oren Clark','1698093048599','1471254806',8978);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (52,'Quentin Pruitt','1606082659199','1445939136',94936);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (53,'Edan Madden','1660122665499','1435191044',27685);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (54,'Russell Sutton','1688080370499','1486003254',92444);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (55,'Porter Bell','1626102923099','1468253964',65461);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (56,'Xander Webster','1666091868399','1439631805',12232);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (57,'Amos Bray','1693122420299','1465406155',46269);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (58,'Aaron Bauer','1666092057899','1448657106',68415);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (59,'Nicholas Marquez','1647100896499','1477698386',59445);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (60,'Lee Roy','1627090807199','1442289355',93909);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (61,'Geoffrey England','1621041354599','1427886214',78911);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (62,'Blake Weber','1633071774199','1469782064',74201);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (63,'Bevis Whitfield','1613112635199','1439537332',12858);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (64,'Rahim Schneider','1613111171299','1443120513',96919);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (65,'Harrison Solis','1610072687099','1457290994',3176);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (66,'Nathan Rodriquez','1612100561799','1454008298',42665);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (67,'Kaseem Cole','1680090653799','1449067154',34495);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (68,'Merrill Knight','1689102870399','1468053878',17086);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (69,'Harper Moses','1664111290699','1486066234',72161);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (70,'Theodore Vargas','1615052506899','1473717697',62851);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (71,'Addison Brown','1664022838999','1434619793',22443);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (72,'Kenneth Moody','1639081240799','1466134058',36728);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (73,'Quinn Hodges','1633112448499','1458491467',5780);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (74,'Walker Tran','1691102400299','1470954736',25035);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (75,'Tad Munoz','1614122490699','1475039593',99304);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (76,'Jermaine Robertson','1620032800299','1476681070',23644);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (77,'Evan Conrad','1684121017199','1440261584',63747);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (78,'Steven May','1632032689499','1480761027',16108);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (79,'Herrod Callahan','1618092867699','1445262835',63593);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (80,'Zane Talley','1695042428299','1459793220',39031);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (81,'Aristotle Oconnor','1631081358199','1468942372',39085);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (82,'Burke Cardenas','1602111394299','1454902606',79532);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (83,'Graiden Tanner','1688090247299','1470559522',67096);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (84,'Dolan Albert','1626052469599','1489156311',38276);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (85,'Scott Sears','1691120240799','1449103505',2806);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (86,'Chester Humphrey','1615062014099','1489637597',29146);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (87,'Laith Nolan','1663101372199','1440028266',3669);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (88,'Noah Larson','1654100378699','1430715038',32282);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (89,'Gannon Harrison','1668012762499','1479161194',28535);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (90,'Bevis Ortiz','1642081540399','1462668362',77542);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (91,'Lucian Gibbs','1622071026099','1467648932',16636);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (92,'Ulric Whitaker','1654092690299','1447241908',63647);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (93,'Hiram Landry','1636052322799','1470967062',96605);

commit;

--drop database;


-----------------------------------------------------

create database '/tmp/dsql.exsample2.fdb'
PAGE_SIZE = 4096
DEFAULT CHARACTER SET UTF8;
commit;

connect '/tmp/dsql.exsample2.fdb';

CREATE TABLE Employee (
  id INTEGER not null,  
  NAME varchar(255) default NULL,
  EmpNo varchar(13) default NULL,
  Birth varchar(255),
  Salary varchar(50) default NULL,
  PRIMARY KEY (id)
);
commit;
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (1,'Jasper Hobbs','1662051221599','1488672867',20464);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (2,'Daniel Sykes','1666020429599','1454200362',58055);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (3,'Lucius Pena','1608021141199','1441394009',88194);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (4,'Channing Fuentes','1636071470099','1483911609',51984);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (5,'Porter Hooper','1653020100499','1478594691',28450);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (6,'Gray Delaney','1685020652099','1465660439',34489);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (7,'Brenden Jordan','1672051445099','1446324166',66523);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (8,'Magee Santiago','1612061784499','1483808639',91002);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (9,'Colin Tate','1667082314199','1489007353',32905);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (10,'Kirk Sullivan','1621010984499','1434493248',42662);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (11,'Sylvester Haynes','1613082107599','1473307179',37051);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (12,'Ryder Rush','1648050734999','1489028357',46668);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (13,'Jack Odonnell','1634111431799','1483490738',28258);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (14,'Craig Hess','1621030200499','1488277959',91844);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (15,'Matthew Crane','1675011679899','1446846160',12840);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (16,'Clinton William','1669030117099','1442208368',80229);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (17,'Kamal Henry','1635040158699','1436856755',58428);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (18,'Nero Durham','1603030776699','1472700685',25592);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (19,'Chaim Lara','1605030689699','1469148126',60512);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (20,'Hunter Boyle','1691120340299','1480533982',30129);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (21,'Ulric Bridges','1665040608299','1439605950',25121);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (22,'Zeph Wagner','1681011434399','1443390181',46066);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (23,'Bruce Carver','1688031135899','1451201033',96494);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (24,'Zahir Glass','1678042183499','1435067037',34524);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (25,'Macaulay Solis','1670111174499','1459045487',65250);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (26,'Channing Conley','1656042833899','1456635107',48600);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (27,'Jason Yang','1609111559199','1443568617',9616);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (28,'Laith Dickerson','1621101548299','1443012132',35746);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (29,'Colt Banks','1632031620599','1467247024',24698);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (30,'Grady Gallagher','1643020258199','1431761141',74136);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (31,'Devin Duran','1682061506099','1467773268',47087);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (32,'Barrett Castaneda','1686062087999','1446637622',51729);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (33,'Lee Spence','1622020596399','1467283065',44499);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (34,'Rafael Higgins','1666053060799','1442750666',84857);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (35,'Lester Mccray','1646122243599','1479938702',23779);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (36,'Lucius Delaney','1694122331599','1429088093',88825);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (37,'Scott Mcleod','1639091725699','1439759248',44938);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (38,'Asher Joyner','1669030592899','1469165482',80019);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (39,'Jameson Maynard','1653070200799','1442882228',80228);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (40,'Chadwick Harmon','1653061676599','1461810936',7558);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (41,'Steel Keller','1690102552399','1448550163',10704);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (42,'Zane Mcfadden','1681050747199','1466805070',9557);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (43,'Hayes Mcclain','1687082389799','1438477141',84225);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (44,'Axel Meadows','1689011118799','1450241998',93113);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (45,'Aladdin Dorsey','1654021339299','1446012638',84610);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (46,'Wyatt Hernandez','1680090213599','1465736052',83013);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (47,'John Jacobson','1616080638899','1443693592',49291);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (48,'Plato Goodman','1608110264799','1441999298',71070);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (49,'Yardley Padilla','1609030613199','1439336836',17095);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (50,'Griffith Ayers','1607111245099','1482724331',74899);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (51,'Oren Clark','1698093048599','1471254806',8978);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (52,'Quentin Pruitt','1606082659199','1445939136',94936);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (53,'Edan Madden','1660122665499','1435191044',27685);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (54,'Russell Sutton','1688080370499','1486003254',92444);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (55,'Porter Bell','1626102923099','1468253964',65461);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (56,'Xander Webster','1666091868399','1439631805',12232);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (57,'Amos Bray','1693122420299','1465406155',46269);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (58,'Aaron Bauer','1666092057899','1448657106',68415);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (59,'Nicholas Marquez','1647100896499','1477698386',59445);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (60,'Lee Roy','1627090807199','1442289355',93909);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (61,'Geoffrey England','1621041354599','1427886214',78911);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (62,'Blake Weber','1633071774199','1469782064',74201);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (63,'Bevis Whitfield','1613112635199','1439537332',12858);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (64,'Rahim Schneider','1613111171299','1443120513',96919);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (65,'Harrison Solis','1610072687099','1457290994',3176);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (66,'Nathan Rodriquez','1612100561799','1454008298',42665);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (67,'Kaseem Cole','1680090653799','1449067154',34495);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (68,'Merrill Knight','1689102870399','1468053878',17086);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (69,'Harper Moses','1664111290699','1486066234',72161);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (70,'Theodore Vargas','1615052506899','1473717697',62851);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (71,'Addison Brown','1664022838999','1434619793',22443);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (72,'Kenneth Moody','1639081240799','1466134058',36728);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (73,'Quinn Hodges','1633112448499','1458491467',5780);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (74,'Walker Tran','1691102400299','1470954736',25035);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (75,'Tad Munoz','1614122490699','1475039593',99304);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (76,'Jermaine Robertson','1620032800299','1476681070',23644);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (77,'Evan Conrad','1684121017199','1440261584',63747);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (78,'Steven May','1632032689499','1480761027',16108);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (79,'Herrod Callahan','1618092867699','1445262835',63593);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (80,'Zane Talley','1695042428299','1459793220',39031);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (81,'Aristotle Oconnor','1631081358199','1468942372',39085);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (82,'Burke Cardenas','1602111394299','1454902606',79532);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (83,'Graiden Tanner','1688090247299','1470559522',67096);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (84,'Dolan Albert','1626052469599','1489156311',38276);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (85,'Scott Sears','1691120240799','1449103505',2806);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (86,'Chester Humphrey','1615062014099','1489637597',29146);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (87,'Louis Bullock','1693071702799','1456655796',17292);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (88,'Graham Ward','1693080125399','1451048168',25269);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (89,'Merritt Mcdonald','1645070379699','1459320798',90329);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (90,'Fuller Alvarez','1627040266099','1463661228',23022);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (91,'Brennan Carroll','1683062596299','1468949603',2400);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (92,'Allen Reyes','1678122565999','1436401953',98706);
INSERT INTO Employee (ID,NAME,EmpNo,Birth,Salary) VALUES (93,'Hoyt Hubbard','1697012318399','1487590540',87869);
commit;

--drop database;
