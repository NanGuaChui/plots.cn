const data = [
  "地址：重庆市永川区汇龙东二路与人民西路交叉口西南60米",
  "姓名：陈晨",
  "电话：18086194631",
  "",
  "地址：重庆市永川区厦坤·天之韵花园1栋",
  "姓名：丁敏",
  "电话：18580166090",
  "",
  "地址：重庆市永川区人民南路66号竹映三清3栋4-1",
  "姓名：屈爱",
  "电话：15084373658",
  "",
  "地址：重庆市永川区长城大厦(人民南路)",
  "姓名：刘林",
  "电话：15025438021",
  "",
  "地址：九龙坡保利花半里玫瑰园1栋3-3",
  "姓名：陈昌",
  "电话：13012303868",
  "",
  "地址：九龙坡吉嘉碧秀苑南区",
  "姓名：黄生",
  "电话：18696579028",
  "",
  "地址：九龙坡区冶金路与前冶路交叉口西南40米 ",
  "姓名：黄龙",
  "电话：13637734713",
  "",
  "地址：九龙坡区西郊二村2号 ",
  "姓名：黄平平",
  "电话：18182267107",
  "",
  "地址：九龙坡区冶金路石坪桥三村 ",
  "姓名：丁玲玲",
  "电话：18883226951",
  "",
  "地址：九龙坡区石坪桥冶金村106号",
  "姓名：丁永林",
  "电话：15023944217",
  "",
  "地址：九龙坡区石坪桥横街冶金三村7号",
  "姓名：陈娟",
  "电话：13752836160",
  "",
  "地址：九龙坡区杨家坪冶金三村",
  "姓名：李宝亮",
  "电话：13637705752",
  "",
  "地址：九龙坡区石坪桥横街51号旁",
  "姓名：彭秀芬",
  "电话：17316793332",
  "",
  "地址：重庆九龙坡中冶·幸福岭2栋",
  "姓名：聂立先",
  "电话：13564644699",
  "",
  "地址：九龙坡区万民路与石坪桥后街交叉口东南100米 ",
  "姓名：黄宗玉",
  "电话：13110211406",
  "",
  "地址：九龙坡区冶金二村3号附3号",
  "姓名：李军",
  "电话：13002322240",
  "",
  "地址：九龙坡区石坪桥后街与后街支路交叉口西南40米",
  "姓名：陈树明",
  "电话：13896734416",
  "",
  "地址：九龙坡区石坪桥后街志龙石坪苑",
  "姓名：王泽轩",
  "电话：15683800752",
  "",
  "地址：九龙坡区石坪桥后街志龙石坪苑",
  "姓名：王泽轩",
  "电话：15223359523",
  "",
  "地址：九龙坡区石坪桥后街与石坪桥正街辅路交叉口南180米",
  "姓名：黄中梅",
  "电话：13320287979",
  "",
  "地址：九龙坡区石坪桥后街1号附8号",
  "姓名：周淑珍",
  "电话：18996019015",
  "",
  "地址：九龙坡区石坪桥后街1号",
  "姓名：曹淑芳",
  "电话：15909390399",
  "",
  "地址：九龙坡区石坪桥后街1号",
  "姓名：王泽轩",
  "电话：15909390399",
  "",
  "地址：九龙坡区万科西九2期 ",
  "姓名：魏显琴",
  "电话：19923253301",
  "",
  "地址：九龙坡区石坪桥横街附近",
  "姓名：胡淑荣",
  "电话：13110108518",
  "",
  "地址：九龙坡区石坪桥横街与彩虹路交叉口东北80米",
  "姓名：杨士林",
  "电话：13002303348",
  "",
  "地址：九龙坡区石坪桥横街18-14号",
  "姓名：曾梅",
  "电话：13452053791",
  "",
  "地址：九龙坡区锦尚一路与石坪桥横街辅路交叉口西南60米",
  "姓名：王用琼",
  "电话：13062394967",
  "",
  "地址：九龙坡区石坪桥横街辅路与石坪桥正街辅路交叉口北140米",
  "姓名：曹吉辉",
  "电话：13883643835",
  "",
  "地址：九龙坡区石坪桥横街2号",
  "姓名：廖琦",
  "电话：13608331112",
  "",
  "地址：九龙坡区石坪桥横街6-6号",
  "姓名：袁德华",
  "电话：17723580800",
  "",
  "地址：九龙坡区石坪桥嘉铠街坊邻居小区1幢1-10号 ",
  "姓名：李竹君",
  "电话：18908270917",
  "",
  "地址：九龙坡区彩虹路与云雁路交叉口西北200米  ",
  "姓名：陈润华",
  "电话：13983992088",
  "",
  "地址：九龙坡区石坪桥正街  ",
  "姓名：吴康吉",
  "电话：13983860710",
  "",
  "地址：九龙坡区石坪桥横街54号阳光心悦2楼附2号1门",
  "姓名：王瑞华",
  "电话：13022389688",
  "",
  "地址：九龙坡区石坪桥横街桃花溪畔小花市集  ",
  "姓名：余顺珍",
  "电话：18623388030",
  "",
  "地址：九龙坡区石坪桥正街119号保利爱尚里1期(火锅店旁) ",
  "姓名：田英",
  "电话：13883886885",
  "",
  "地址：九龙坡区石坪桥正街119号附14号 ",
  "姓名：马美森",
  "电话：18523308615",
  "",
  "地址：九龙坡区滨河美食街(保利爱尚里1期尚美店)",
  "姓名：陈树明",
  "电话：15608359560",
  "",
  "地址：九龙坡区石坪桥正街119号附75号",
  "姓名：刘文",
  "电话：15223394850",
  "",
  "地址：九龙坡区石坪桥正街119号附81号 ",
  "姓名：周景明",
  "电话：13650580609",
  "",
  "地址：九龙坡区石坪桥正街119号附59号 ",
  "姓名：罗华春",
  "电话：15923440552",
  "",
  "地址：九龙坡区石坪桥正街115号附62号   ",
  "姓名：张玲",
  "电话：13618284718",
  "",
  "地址：九龙坡区石坪桥正街115号保利爱尚里彩乐荟3幢2415室",
  "姓名：邓永清",
  "电话：15111991790",
  "",
  "",
];
