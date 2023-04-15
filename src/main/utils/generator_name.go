package utils

import (
	"math/rand"
	"time"
)

var (
	surname  = []string{"赵", "钱", "孙", "李", "周", "吴", "郑", "王", "冯", "陈", "楮", "卫", "蒋", "沈", "韩", "杨", "朱", "秦", "尤", "许", "何", "吕", "施", "张", "孔", "曹", "严", "华", "金", "魏", "陶", "姜", "戚", "谢", "邹", "喻", "柏", "水", "窦", "章", "云", "苏", "潘", "葛", "奚", "范", "彭", "郎", "鲁", "韦", "昌", "马", "苗", "凤", "花", "方", "俞", "任", "袁", "柳", "酆", "鲍", "史", "唐", "费", "廉", "岑", "薛", "雷", "贺", "倪", "汤", "藤", "殷", "罗", "毕", "郝", "邬", "安", "常", "乐", "于", "时", "傅", "皮", "卞", "齐", "康", "伍", "余", "元", "卜", "顾", "孟", "平", "黄", "和", "穆", "萧", "尹", "欧阳", "慕容"}
	maleName = []string{"建国", "继伟", "国强", "国泰", "国宝", "国庆", "宏伟", "宏达", "强", "刚", "健", "俊", "明", "鹏", "飞", "海", "亮", "军", "伟", "勇", "祥", "华", "坚", "宝", "磊", "利", "龙", "成", "康", "杰", "建军", "建华", "建设", "志强", "志明", "志勇", "学文", "学友", "学庆", "秀英", "秀兰", "宜华", "宜春", "宜民", "宜昌", "运涛", "运发", "运城", "运海", "元龙", "元福", "元忠", "永康", "永生", "永安", "永和", "永瑞", "友华", "友明", "友善", "有信", "玉树"}
	// femaleName 为女性名字
	femaleName = []string{"秀英", "秀兰", "宜华", "宜春", "宜民", "宜昌", "运涛", "运发", "运城", "运海", "婷婷", "丽丽", "敏敏", "雅雅", "慧慧", "红红", "佳佳", "倩倩", "娜娜", "静静", "梅梅", "燕燕", "莉莉", "艳艳", "雪雪", "琳琳", "晨晨", "辉辉", "欣欣", "阳阳", "峰峰", "文文", "亮亮", "明明", "莹莹", "慧慧", "丹丹", "婷婷", "洁洁", "蓉蓉", "瑞瑞", "艳艳", "翠翠", "娟娟", "秀秀", "霞霞", "梦梦", "飞飞", "红梅", "晓燕", "婷婷", "玲玲", "秀珍", "凤凰", "雯雯", "倩倩", "媛媛", "月月", "瑜瑜", "萍萍", "云云", "莺莺", "燕子"}
)

var (
	surnames = []string{"Smith", "Johnson", "Williams", "Jones", "Brown", "Garcia", "Miller", "Davis", "Rodriguez", "Martinez", "Hernandez", "Lopez", "Gonzalez", "Perez", "Taylor", "Anderson", "Wilson", "Jackson", "White", "Harris", "Martin", "Thompson", "Moore", "Young", "Allen", "King", "Wright", "Scott", "Green", "Baker", "Adams", "Nelson", "Carter", "Mitchell", "Perez", "Roberts", "Turner", "Phillips", "Campbell", "Parker", "Evans", "Edwards", "Collins", "Stewart", "Sanchez", "Morris", "Rogers", "Reed", "Cook"}

	// 名字数据集合
	maleNames = []string{"John", "William", "David", "Richard", "Thomas", "Charles", "Christopher", "Daniel", "Matthew", "Anthony", "Joseph", "Donald", "Mark", "Paul", "George", "Steven", "Andrew", "Edward", "Brian", "Kevin", "Robert", "Joshua", "Justin", "Eric", "Jason", "Jacob", "Alexander", "Adam", "Benjamin", "Jonathan", "Tyler", "Nathan", "Brandon", "Aaron", "Zachary", "Stephen", "Timothy", "Jacob", "Jesse", "Logan", "Nicholas", "Michael", "Dylan", "Ethan", "Cameron", "Samuel", "Ryan", "Ian", "Travis", "Connor"}

	femaleNames = []string{"Mary", "Patricia", "Linda", "Barbara", "Elizabeth", "Jennifer", "Maria", "Susan", "Margaret", "Dorothy", "Lisa", "Nancy", "Karen", "Betty", "Helen", "Sandra", "Donna", "Carol", "Ruth", "Sharon", "Michelle", "Laura", "Sarah", "Kimberly", "Deborah", "Jessica", "Amy", "Angela", "Melissa", "Rebecca", "Stephanie", "Lori", "Julie", "Tiffany", "Emily", "Christine", "Catherine", "Virginia", "Samantha", "Janet", "Katherine", "Jacqueline", "Frances", "Ann", "Alice", "Jean", "Teresa"}
)

func GenerateName() string {
	rand.Seed(time.Now().UnixNano())

	// 随机选择一个姓氏
	name := surnames[rand.Intn(len(surnames))]

	// 50% 的概率选择男性名字，50% 的概率选择女性名字
	if rand.Intn(2) == 0 {
		name += maleNames[rand.Intn(len(maleNames))]
	} else {
		name += femaleNames[rand.Intn(len(femaleNames))]
	}
	return name
}
