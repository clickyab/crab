
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
UPDATE users SET city_id=NULL;
UPDATE campaign_attributes SET region=NULL;
ALTER TABLE users DROP FOREIGN KEY user_city_id_fk;
ALTER TABLE provinces DROP FOREIGN KEY provinces_countries_id_fk;
ALTER TABLE cities DROP FOREIGN KEY cities_provinces_id_fk;
DROP TABLE cities;
DROP TABLE provinces;

CREATE TABLE `provinces` (
`code` char(5) NOT NULL,
`name` varchar(50) NOT NULL,
`country_id` int(10) unsigned NOT NULL,
`fa_name` varchar(50) DEFAULT NULL,
PRIMARY KEY (`name`),
UNIQUE KEY `provinces_code_uindex` (`code`),
KEY `provinces_countries_id_fk` (`country_id`),
CONSTRAINT `provinces_countries_id_fk` FOREIGN KEY (`country_id`) REFERENCES `countries` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `cities` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `province` varchar(50) NOT NULL,
  `name` varchar(50) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `cities_provinces_test_name_fk` (`province`),
  CONSTRAINT `cities_provinces_test_name_fk` FOREIGN KEY (`province`) REFERENCES `provinces` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

ALTER TABLE users
  ADD CONSTRAINT users_cities_id_fk
FOREIGN KEY (city_id) REFERENCES cities (id);


INSERT INTO provinces (code, name, country_id, fa_name) VALUES ('IR-01', 'Azarbayjane-e Sharqi', 1, 'آذربايجان شرقي');
INSERT INTO provinces (code, name, country_id, fa_name) VALUES ('IR-02', 'Azarbayjane-e Gharbi', 1, 'آذربايجان غربي');
INSERT INTO provinces (code, name, country_id, fa_name) VALUES ('IR-03', 'Ardabil', 1, 'اردبيل');
INSERT INTO provinces (code, name, country_id, fa_name) VALUES ('IR-04', 'Esfahan', 1, 'اصفهان');
INSERT INTO provinces (code, name, country_id, fa_name) VALUES ('IR-05', 'Ilam', 1, 'ايلام');
INSERT INTO provinces (code, name, country_id, fa_name) VALUES ('IR-06', 'Bushehr', 1, 'بوشهر');
INSERT INTO provinces (code, name, country_id, fa_name) VALUES ('IR-07', 'Tehran', 1, 'تهران');
INSERT INTO provinces (code, name, country_id, fa_name) VALUES ('IR-08', 'Chahar Mahal va Bakhtiari', 1, 'چهارمحال بختياري');
INSERT INTO provinces (code, name, country_id, fa_name) VALUES ('IR-10', 'Khuzestan', 1, 'خوزستان');
INSERT INTO provinces (code, name, country_id, fa_name) VALUES ('IR-11', 'Zanjan', 1, 'زنجان');
INSERT INTO provinces (code, name, country_id, fa_name) VALUES ('IR-12', 'Semnan', 1, 'سمنان');
INSERT INTO provinces (code, name, country_id, fa_name) VALUES ('IR-13', 'Sistan va Baluchestan', 1, 'سيستان و بلوچستان');
INSERT INTO provinces (code, name, country_id, fa_name) VALUES ('IR-14', 'Fars', 1, 'فارس');
INSERT INTO provinces (code, name, country_id, fa_name) VALUES ('IR-15', 'Kerman', 1, 'كرمان');
INSERT INTO provinces (code, name, country_id, fa_name) VALUES ('IR-16', 'Kordestan', 1, 'كردستان');
INSERT INTO provinces (code, name, country_id, fa_name) VALUES ('IR-17', 'Kermanshah', 1, 'كرمانشاه');
INSERT INTO provinces (code, name, country_id, fa_name) VALUES ('IR-18', 'Kohgiluye va Bowyer Ahmad', 1, 'كهكيلويه و بويراحمد');
INSERT INTO provinces (code, name, country_id, fa_name) VALUES ('IR-19', 'Gilan', 1, 'گيلان');
INSERT INTO provinces (code, name, country_id, fa_name) VALUES ('IR-20', 'Lorestan', 1, 'لرستان');
INSERT INTO provinces (code, name, country_id, fa_name) VALUES ('IR-21', 'Mazandaran', 1, 'مازندران');
INSERT INTO provinces (code, name, country_id, fa_name) VALUES ('IR-22', 'Markazi', 1, 'مركزي');
INSERT INTO provinces (code, name, country_id, fa_name) VALUES ('IR-23', 'Hormozgan', 1, 'هرمزگان');
INSERT INTO provinces (code, name, country_id, fa_name) VALUES ('IR-24', 'Hamadan', 1, 'همدان');
INSERT INTO provinces (code, name, country_id, fa_name) VALUES ('IR-25', 'Yazd', 1, 'يزد');
INSERT INTO provinces (code, name, country_id, fa_name) VALUES ('IR-26', 'Qom', 1, 'قم');
INSERT INTO provinces (code, name, country_id, fa_name) VALUES ('IR-27', 'Golestan', 1, 'گلستان');
INSERT INTO provinces (code, name, country_id, fa_name) VALUES ('IR-28', 'Qazvin', 1, 'قزوين');
INSERT INTO provinces (code, name, country_id, fa_name) VALUES ('IR-29', 'Khorasan-e Jonubi', 1, 'خراسان جنوبي');
INSERT INTO provinces (code, name, country_id, fa_name) VALUES ('IR-30', 'Khorasan-e Razavi', 1, 'خراسان رضوي');
INSERT INTO provinces (code, name, country_id, fa_name) VALUES ('IR-31', 'Khorasan-e Shomali', 1, 'خراسان شمالي');
INSERT INTO provinces (code, name, country_id, fa_name) VALUES ('IR-32', 'Alborz', 1, 'البرز');


INSERT INTO  cities (id,province,name ) VALUES
  (1, "Azarbayjane-e Sharqi", 'تبريز'),
  (2, "Azarbayjane-e Sharqi", 'كندوان'),
  (3, "Azarbayjane-e Sharqi", 'بندر شرفخانه'),
  (4, "Azarbayjane-e Sharqi", 'مراغه'),
  (5, "Azarbayjane-e Sharqi", 'ميانه'),
  (6, "Azarbayjane-e Sharqi", 'شبستر'),
  (7, "Azarbayjane-e Sharqi", 'مرند'),
  (8, "Azarbayjane-e Sharqi", 'جلفا'),
  (9, "Azarbayjane-e Sharqi", 'سراب'),
  (10, "Azarbayjane-e Sharqi", 'هاديشهر'),
  (11, "Azarbayjane-e Sharqi", 'بناب'),
  (12, "Azarbayjane-e Sharqi", 'كليبر'),
  (13, "Azarbayjane-e Sharqi", 'تسوج'),
  (14, "Azarbayjane-e Sharqi", 'اهر'),
  (15, "Azarbayjane-e Sharqi", 'هريس'),
  (16, "Azarbayjane-e Sharqi", 'عجبشير'),
  (17, "Azarbayjane-e Sharqi", 'هشترود'),
  (18, "Azarbayjane-e Sharqi", 'ملكان'),
  (19, "Azarbayjane-e Sharqi", 'بستان آباد'),
  (20, "Azarbayjane-e Sharqi", 'ورزقان'),
  (21, "Azarbayjane-e Sharqi", 'اسكو'),
  (22, "Azarbayjane-e Sharqi", 'آذر شهر'),
  (23, "Azarbayjane-e Sharqi", 'قره آغاج'),
  (24, "Azarbayjane-e Sharqi", 'ممقان'),
  (25, "Azarbayjane-e Sharqi", 'صوفیان'),
  (26, "Azarbayjane-e Sharqi", 'ایلخچی'),
  (27, "Azarbayjane-e Sharqi", 'خسروشهر'),
  (28, "Azarbayjane-e Sharqi", 'باسمنج'),
  (29, "Azarbayjane-e Sharqi", 'سهند'),
  (30, "Azarbayjane-e Gharbi", 'اروميه'),
  (31, "Azarbayjane-e Gharbi", 'نقده'),
  (32, "Azarbayjane-e Gharbi", 'ماكو'),
  (33, "Azarbayjane-e Gharbi", 'تكاب'),
  (34, "Azarbayjane-e Gharbi", 'خوي'),
  (35, "Azarbayjane-e Gharbi", 'مهاباد'),
  (36, "Azarbayjane-e Gharbi", 'سر دشت'),
  (37, "Azarbayjane-e Gharbi", 'چالدران'),
  (38, "Azarbayjane-e Gharbi", 'بوكان'),
  (39, "Azarbayjane-e Gharbi", 'مياندوآب'),
  (40, "Azarbayjane-e Gharbi", 'سلماس'),
  (41, "Azarbayjane-e Gharbi", 'شاهين دژ'),
  (42, "Azarbayjane-e Gharbi", 'پيرانشهر'),
  (43, "Azarbayjane-e Gharbi", 'سيه چشمه'),
  (44, "Azarbayjane-e Gharbi", 'اشنويه'),
  (45, "Azarbayjane-e Gharbi", 'چایپاره'),
  (46, "Azarbayjane-e Gharbi", 'پلدشت'),
  (47, "Azarbayjane-e Gharbi", 'شوط'),
  (48, "Ardabil", 'اردبيل'),
  (49, "Ardabil", 'سرعين'),
  (50, "Ardabil", 'بيله سوار'),
  (51, "Ardabil", 'پارس آباد'),
  (52, "Ardabil", 'خلخال'),
  (53, "Ardabil", 'مشگين شهر'),
  (54, "Ardabil", 'مغان'),
  (55, "Ardabil", 'نمين'),
  (56, "Ardabil", 'نير'),
  (57, "Ardabil", 'كوثر'),
  (58, "Ardabil", 'كيوي'),
  (59, "Ardabil", 'گرمي'),
  (60, "Esfahan", 'اصفهان'),
  (61, "Esfahan", 'فريدن'),
  (62, "Esfahan", 'فريدون شهر'),
  (63, "Esfahan", 'فلاورجان'),
  (64, "Esfahan", 'گلپايگان'),
  (65, "Esfahan", 'دهاقان'),
  (66, "Esfahan", 'نطنز'),
  (67, "Esfahan", 'نايين'),
  (68, "Esfahan", 'تيران'),
  (69, "Esfahan", 'كاشان'),
  (70, "Esfahan", 'فولاد شهر'),
  (71, "Esfahan", 'اردستان'),
  (72, "Esfahan", 'سميرم'),
  (73, "Esfahan", 'درچه'),
  (74, "Esfahan", 'کوهپایه'),
  (75, "Esfahan", 'مباركه'),
  (76, "Esfahan", 'شهرضا'),
  (77, "Esfahan", 'خميني شهر'),
  (78, "Esfahan", 'شاهين شهر'),
  (79, "Esfahan", 'نجف آباد'),
  (80, "Esfahan", 'دولت آباد'),
  (81, "Esfahan", 'زرين شهر'),
  (82, "Esfahan", 'آران و بيدگل'),
  (83, "Esfahan", 'باغ بهادران'),
  (84, "Esfahan", 'خوانسار'),
  (85, "Esfahan", 'مهردشت'),
  (86, "Esfahan", 'علويجه'),
  (87, "Esfahan", 'عسگران'),
  (88, "Esfahan", 'نهضت آباد'),
  (89, "Esfahan", 'حاجي آباد'),
  (90, "Esfahan", 'تودشک'),
  (91, "Esfahan", 'ورزنه'),
  (92, "Ilam", 'ايلام'),
  (93, "Ilam", 'مهران'),
  (94, "Ilam", 'دهلران'),
  (95, "Ilam", 'آبدانان'),
  (96, "Ilam", 'شيروان چرداول'),
  (97, "Ilam", 'دره شهر'),
  (98, "Ilam", 'ايوان'),
  (99, "Ilam", 'سرابله'),
  (100, "Bushehr", 'بوشهر'),
  (101, "Bushehr", 'تنگستان'),
  (102, "Bushehr", 'دشتستان'),
  (103, "Bushehr", 'دير'),
  (104, "Bushehr", 'ديلم'),
  (105, "Bushehr", 'كنگان'),
  (106, "Bushehr", 'گناوه'),
  (107, "Bushehr", 'ريشهر'),
  (108, "Bushehr", 'دشتي'),
  (109, "Bushehr", 'خورموج'),
  (110, "Bushehr", 'اهرم'),
  (111, "Bushehr", 'برازجان'),
  (112, "Bushehr", 'خارك'),
  (113, "Bushehr", 'جم'),
  (114, "Bushehr", 'کاکی'),
  (115, "Bushehr", 'عسلویه'),
  (116, "Bushehr", 'بردخون'),
  (117, "Tehran", 'تهران'),
  (118, "Tehran", 'ورامين'),
  (119, "Tehran", 'فيروزكوه'),
  (120, "Tehran", 'ري'),
  (121, "Tehran", 'دماوند'),
  (122, "Tehran", 'اسلامشهر'),
  (123, "Tehran", 'رودهن'),
  (124, "Tehran", 'لواسان'),
  (125, "Tehran", 'بومهن'),
  (126, "Tehran", 'تجريش'),
  (127, "Tehran", 'فشم'),
  (128, "Tehran", 'كهريزك'),
  (129, "Tehran", 'پاكدشت'),
  (130, "Tehran", 'چهاردانگه'),
  (131, "Tehran", 'شريف آباد'),
  (132, "Tehran", 'قرچك'),
  (133, "Tehran", 'باقرشهر'),
  (134, "Tehran", 'شهريار'),
  (135, "Tehran", 'رباط كريم'),
  (136, "Tehran", 'قدس'),
  (137, "Tehran", 'ملارد'),
  (138, "Chahar Mahal va Bakhtiari", 'شهركرد'),
  (139, "Chahar Mahal va Bakhtiari", 'فارسان'),
  (140, "Chahar Mahal va Bakhtiari", 'بروجن'),
  (141, "Chahar Mahal va Bakhtiari", 'چلگرد'),
  (142, "Chahar Mahal va Bakhtiari", 'اردل'),
  (143, "Chahar Mahal va Bakhtiari", 'لردگان'),
  (144, "Chahar Mahal va Bakhtiari", 'سامان'),
  (145, "Khorasan-e Jonubi", 'قائن'),
  (146, "Khorasan-e Jonubi", 'فردوس'),
  (147, "Khorasan-e Jonubi", 'بيرجند'),
  (148, "Khorasan-e Jonubi", 'نهبندان'),
  (149, "Khorasan-e Jonubi", 'سربيشه'),
  (150, "Khorasan-e Jonubi", 'طبس مسینا'),
  (151, "Khorasan-e Jonubi", 'قهستان'),
  (152, "Khorasan-e Jonubi", 'درمیان'),
  (153, "Khorasan-e Razavi", 'مشهد'),
  (154, "Khorasan-e Razavi", 'نيشابور'),
  (155, "Khorasan-e Razavi", 'سبزوار'),
  (156, "Khorasan-e Razavi", 'كاشمر'),
  (157, "Khorasan-e Razavi", 'گناباد'),
  (158, "Khorasan-e Razavi", 'طبس'),
  (159, "Khorasan-e Razavi", 'تربت حيدريه'),
  (160, "Khorasan-e Razavi", 'خواف'),
  (161, "Khorasan-e Razavi", 'تربت جام'),
  (162, "Khorasan-e Razavi", 'تايباد'),
  (163, "Khorasan-e Razavi", 'قوچان'),
  (164, "Khorasan-e Razavi", 'سرخس'),
  (165, "Khorasan-e Razavi", 'بردسكن'),
  (166, "Khorasan-e Razavi", 'فريمان'),
  (167, "Khorasan-e Razavi", 'چناران'),
  (168, "Khorasan-e Razavi", 'درگز'),
  (169, "Khorasan-e Razavi", 'كلات'),
  (170, "Khorasan-e Razavi", 'طرقبه'),
  (171, "Khorasan-e Razavi", 'سر ولایت'),
  (172, "Khorasan-e Shomali", 'بجنورد'),
  (173, "Khorasan-e Shomali", 'اسفراين'),
  (174, "Khorasan-e Shomali", 'جاجرم'),
  (175, "Khorasan-e Shomali", 'شيروان'),
  (176, "Khorasan-e Shomali", 'آشخانه'),
  (177, "Khorasan-e Shomali", 'گرمه'),
  (178, "Khorasan-e Shomali", 'ساروج'),
  (179, "Khuzestan", 'اهواز'),
  (180, "Khuzestan", 'ايرانشهر'),
  (181, "Khuzestan", 'شوش'),
  (182, "Khuzestan", 'آبادان'),
  (183, "Khuzestan", 'خرمشهر'),
  (184, "Khuzestan", 'مسجد سليمان'),
  (185, "Khuzestan", 'ايذه'),
  (186, "Khuzestan", 'شوشتر'),
  (187, "Khuzestan", 'انديمشك'),
  (188, "Khuzestan", 'سوسنگرد'),
  (189, "Khuzestan", 'هويزه'),
  (190, "Khuzestan", 'دزفول'),
  (191, "Khuzestan", 'شادگان'),
  (192, "Khuzestan", 'بندر ماهشهر'),
  (193, "Khuzestan", 'بندر امام خميني'),
  (194, "Khuzestan", 'اميديه'),
  (195, "Khuzestan", 'بهبهان'),
  (196, "Khuzestan", 'رامهرمز'),
  (197, "Khuzestan", 'باغ ملك'),
  (198, "Khuzestan", 'هنديجان'),
  (199, "Khuzestan", 'لالي'),
  (200, "Khuzestan", 'رامشیر'),
  (201, "Khuzestan", 'حمیدیه'),
  (202, "Khuzestan", 'دغاغله'),
  (203, "Khuzestan", 'ملاثانی'),
  (205, "Khuzestan", 'ویسی'),
  (206, "Zanjan", 'زنجان'),
  (207, "Zanjan", 'ابهر'),
  (208, "Zanjan", 'خدابنده'),
  (209, "Zanjan", 'طارم'),
  (210, "Zanjan", 'ماهنشان'),
  (211, "Zanjan", 'خرمدره'),
  (212, "Zanjan", 'ايجرود'),
  (213, "Zanjan", 'زرين آباد'),
  (214, "Zanjan", 'آب بر'),
  (215, "Zanjan", 'قيدار'),
  (216, "Semnan", 'سمنان'),
  (217, "Semnan", 'شاهرود'),
  (218, "Semnan", 'گرمسار'),
  (219, "Semnan", 'ايوانكي'),
  (220, "Semnan", 'دامغان'),
  (221, "Semnan", 'بسطام'),
  (222, "Sistan va Baluchestan", 'زاهدان'),
  (223, "Sistan va Baluchestan", 'چابهار'),
  (224, "Sistan va Baluchestan", 'خاش'),
  (225, "Sistan va Baluchestan", 'سراوان'),
  (226, "Sistan va Baluchestan", 'زابل'),
  (227, "Sistan va Baluchestan", 'سرباز'),
  (228, "Sistan va Baluchestan", 'نيكشهر'),
  (229, "Sistan va Baluchestan", 'ايرانشهر'),
  (230, "Sistan va Baluchestan", 'راسك'),
  (231, "Sistan va Baluchestan", 'ميرجاوه'),
  (232, "Fars", 'شيراز'),
  (233, "Fars", 'اقليد'),
  (234, "Fars", 'داراب'),
  (235, "Fars", 'فسا'),
  (236, "Fars", 'مرودشت'),
  (237, "Fars", 'خرم بيد'),
  (238, "Fars", 'آباده'),
  (239, "Fars", 'كازرون'),
  (240, "Fars", 'ممسني'),
  (241, "Fars", 'سپيدان'),
  (242, "Fars", 'لار'),
  (243, "Fars", 'فيروز آباد'),
  (244, "Fars", 'جهرم'),
  (245, "Fars", 'ني ريز'),
  (246, "Fars", 'استهبان'),
  (247, "Fars", 'لامرد'),
  (248, "Fars", 'مهر'),
  (249, "Fars", 'حاجي آباد'),
  (250, "Fars", 'نورآباد'),
  (251, "Fars", 'اردكان'),
  (252, "Fars", 'صفاشهر'),
  (253, "Fars", 'ارسنجان'),
  (254, "Fars", 'قيروكارزين'),
  (255, "Fars", 'سوريان'),
  (256, "Fars", 'فراشبند'),
  (257, "Fars", 'سروستان'),
  (258, "Fars", 'ارژن'),
  (259, "Fars", 'گويم'),
  (260, "Fars", 'داريون'),
  (261, "Fars", 'زرقان'),
  (262, "Fars", 'خان زنیان'),
  (263, "Fars", 'کوار'),
  (264, "Fars", 'ده بید'),
  (265, "Fars", 'باب انار/خفر'),
  (266, "Fars", 'بوانات'),
  (267, "Fars", 'خرامه'),
  (268, "Fars", 'خنج'),
  (269, "Fars", 'سیاخ دارنگون'),
  (270, "Qazvin", 'قزوين'),
  (271, "Qazvin", 'تاكستان'),
  (272, "Qazvin", 'آبيك'),
  (273, "Qazvin", 'بوئين زهرا'),
  (274, "Qom", 'قم'),
  (275, "Alborz", 'طالقان'),
  (276, "Alborz", 'نظرآباد'),
  (277, "Alborz", 'اشتهارد'),
  (278, "Alborz", 'هشتگرد'),
  (279, "Alborz", 'كن'),
  (280, "Alborz", 'آسارا'),
  (281, "Alborz", 'شهرک گلستان'),
  (282, "Alborz", 'اندیشه'),
  (283, "Alborz", 'كرج'),
  (284, "Alborz", 'نظر آباد'),
  (285, "Alborz", 'گوهردشت'),
  (286, "Alborz", 'ماهدشت'),
  (287, "Alborz", 'مشکین دشت'),
  (288, "Kordestan", 'سنندج'),
  (289, "Kordestan", 'ديواندره'),
  (290, "Kordestan", 'بانه'),
  (291, "Kordestan", 'بيجار'),
  (292, "Kordestan", 'سقز'),
  (293, "Kordestan", 'كامياران'),
  (294, "Kordestan", 'قروه'),
  (295, "Kordestan", 'مريوان'),
  (296, "Kordestan", 'صلوات آباد'),
  (297, "Kordestan", 'حسن آباد'),
  (298, "Kerman", 'كرمان'),
  (299, "Kerman", 'راور'),
  (300, "Kerman", 'بابك'),
  (301, "Kerman", 'انار'),
  (302, "Kerman", 'کوهبنان'),
  (303, "Kerman", 'رفسنجان'),
  (304, "Kerman", 'بافت'),
  (305, "Kerman", 'سيرجان'),
  (306, "Kerman", 'كهنوج'),
  (307, "Kerman", 'زرند'),
  (308, "Kerman", 'بم'),
  (309, "Kerman", 'جيرفت'),
  (310, "Kerman", 'بردسير'),
  (311, "Kermanshah", 'كرمانشاه'),
  (312, "Kermanshah", 'اسلام آباد غرب'),
  (313, "Kermanshah", 'سر پل ذهاب'),
  (314, "Kermanshah", 'كنگاور'),
  (315, "Kermanshah", 'سنقر'),
  (316, "Kermanshah", 'قصر شيرين'),
  (317, "Kermanshah", 'گيلان غرب'),
  (318, "Kermanshah", 'هرسين'),
  (319, "Kermanshah", 'صحنه'),
  (320, "Kermanshah", 'پاوه'),
  (321, "Kermanshah", 'جوانرود'),
  (322, "Kermanshah", 'شاهو'),
  (323, "Kohgiluye va Bowyer Ahmad", 'ياسوج'),
  (324, "Kohgiluye va Bowyer Ahmad", 'گچساران'),
  (325, "Kohgiluye va Bowyer Ahmad", 'دنا'),
  (326, "Kohgiluye va Bowyer Ahmad", 'دوگنبدان'),
  (327, "Kohgiluye va Bowyer Ahmad", 'سي سخت'),
  (328, "Kohgiluye va Bowyer Ahmad", 'دهدشت'),
  (329, "Kohgiluye va Bowyer Ahmad", 'ليكك'),
  (330, "Golestan", 'گرگان'),
  (331, "Golestan", 'آق قلا'),
  (332, "Golestan", 'گنبد كاووس'),
  (333, "Golestan", 'علي آباد كتول'),
  (334, "Golestan", 'مينو دشت'),
  (335, "Golestan", 'تركمن'),
  (336, "Golestan", 'كردكوي'),
  (337, "Golestan", 'بندر گز'),
  (338, "Golestan", 'كلاله'),
  (339, "Golestan", 'آزاد شهر'),
  (340, "Golestan", 'راميان'),
  (341, "Gilan", 'رشت'),
  (342, "Gilan", 'منجيل'),
  (343, "Gilan", 'لنگرود'),
  (344, "Gilan", 'رود سر'),
  (345, "Gilan", 'تالش'),
  (346, "Gilan", 'آستارا'),
  (347, "Gilan", 'ماسوله'),
  (348, "Gilan", 'آستانه اشرفيه'),
  (349, "Gilan", 'رودبار'),
  (350, "Gilan", 'فومن'),
  (351, "Gilan", 'صومعه سرا'),
  (352, "Gilan", 'بندرانزلي'),
  (353, "Gilan", 'كلاچاي'),
  (354, "Gilan", 'هشتپر'),
  (355, "Gilan", 'رضوان شهر'),
  (356, "Gilan", 'ماسال'),
  (357, "Gilan", 'شفت'),
  (358, "Gilan", 'سياهكل'),
  (359, "Gilan", 'املش'),
  (360, "Gilan", 'لاهیجان'),
  (361, "Gilan", 'خشک بيجار'),
  (362, "Gilan", 'خمام'),
  (363, "Gilan", 'لشت نشا'),
  (364, "Gilan", 'بندر کياشهر'),
  (365, "Lorestan", 'خرم آباد'),
  (366, "Lorestan", 'ماهشهر'),
  (367, "Lorestan", 'دزفول'),
  (368, "Lorestan", 'بروجرد'),
  (369, "Lorestan", 'دورود'),
  (370, "Lorestan", 'اليگودرز'),
  (371, "Lorestan", 'ازنا'),
  (372, "Lorestan", 'نور آباد'),
  (373, "Lorestan", 'كوهدشت'),
  (374, "Lorestan", 'الشتر'),
  (375, "Lorestan", 'پلدختر'),
  (376, "Mazandaran", 'ساري'),
  (377, "Mazandaran", 'آمل'),
  (378, "Mazandaran", 'بابل'),
  (379, "Mazandaran", 'بابلسر'),
  (380, "Mazandaran", 'بهشهر'),
  (381, "Mazandaran", 'تنكابن'),
  (382, "Mazandaran", 'جويبار'),
  (383, "Mazandaran", 'چالوس'),
  (384, "Mazandaran", 'رامسر'),
  (385, "Mazandaran", 'سواد كوه'),
  (386, "Mazandaran", 'قائم شهر'),
  (387, "Mazandaran", 'نكا'),
  (388, "Mazandaran", 'نور'),
  (389, "Mazandaran", 'بلده'),
  (390, "Mazandaran", 'نوشهر'),
  (391, "Mazandaran", 'پل سفيد'),
  (392, "Mazandaran", 'محمود آباد'),
  (393, "Mazandaran", 'فريدون كنار'),
  (394, "Markazi", 'اراك'),
  (395, "Markazi", 'آشتيان'),
  (396, "Markazi", 'تفرش'),
  (397, "Markazi", 'خمين'),
  (398, "Markazi", 'دليجان'),
  (399, "Markazi", 'ساوه'),
  (400, "Markazi", 'سربند'),
  (401, "Markazi", 'محلات'),
  (402, "Markazi", 'شازند'),
  (403, "Hormozgan", 'بندرعباس'),
  (404, "Hormozgan", 'قشم'),
  (405, "Hormozgan", 'كيش'),
  (406, "Hormozgan", 'بندر لنگه'),
  (407, "Hormozgan", 'بستك'),
  (408, "Hormozgan", 'حاجي آباد'),
  (409, "Hormozgan", 'دهبارز'),
  (410, "Hormozgan", 'انگهران'),
  (411, "Hormozgan", 'ميناب'),
  (412, "Hormozgan", 'ابوموسي'),
  (413, "Hormozgan", 'بندر جاسك'),
  (414, "Hormozgan", 'تنب بزرگ'),
  (415, "Hormozgan", 'بندر خمیر'),
  (416, "Hormozgan", 'پارسیان'),
  (418, "Hamadan", 'همدان'),
  (419, "Hamadan", 'ملاير'),
  (420, "Hamadan", 'تويسركان'),
  (421, "Hamadan", 'نهاوند'),
  (422, "Hamadan", 'كبودر اهنگ'),
  (423, "Hamadan", 'رزن'),
  (424, "Hamadan", 'اسدآباد'),
  (425, "Hamadan", 'بهار'),
  (426, "Yazd", 'يزد'),
  (427, "Yazd", 'تفت'),
  (428, "Yazd", 'اردكان'),
  (429, "Yazd", 'ابركوه'),
  (430, "Yazd", 'ميبد'),
  (431, "Yazd", 'طبس'),
  (432, "Yazd", 'بافق'),
  (433, "Yazd", 'مهريز'),
  (434, "Yazd", 'اشكذر'),
  (435, "Yazd", 'هرات'),
  (436, "Yazd", 'خضرآباد'),
  (437, "Yazd", 'شاهديه'),
  (438, "Yazd", 'حمیدیه شهر'),
  (439, "Yazd", 'سید میرزا'),
  (440, "Yazd", 'زارچ');

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back


