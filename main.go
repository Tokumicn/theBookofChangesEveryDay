package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	Yin    int = 0
	Yang   int = 1
	GuaLen     = 6
)

const (
	Qian = 1 // 乾卦
	Kun  = 2 // 坤卦
	Zhen = 3 // 震卦
	Gen  = 4 // 艮卦
	Li   = 5 //  离卦
	Kan  = 6 // 坎卦
	Dui  = 7 // 兑卦
	Xun  = 8 // 巽卦

	YangYao = "—"
	YinYao  = "--"
)

type Yao struct {
	Image     string `json:"image"`     // 爻象
	Text      string `json:"text"`      // 爻辞
	ImageText string `json:"imageText"` // 象辞
}

// 八卦
type BaGua struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Char       string `json:"char"`
	YinYangArr []int  `json:"yinYangArr"`
	Tips       string `json:"tips"`
}

// 卦象： 六十四卦
type GuaImage struct {
	Index   int         `json:"index"`   // 卦序
	Wai     BaGua       `json:"wai"`     // 上卦，外卦
	Nei     BaGua       `json:"nei"`     // 下卦，内卦
	Name    string      `json:"name"`    // 卦名
	DuYin   string      `json:"duYin"`   // 读音
	Text    string      `json:"text"`    // 卦辞
	Extra   string      `json:"extra"`   // 额外信息
	YongYao Yao         `json:"yongYao"` // 如用九、用六
	Short   string      `json:"short"`   // 卦简介
	Desc    string      `json:"desc"`    // 介绍
	Yao     [GuaLen]Yao `json:"yao"`     // 六爻
}

var (
	BaGuaMap = map[int]BaGua{
		Qian: {
			Id:         Qian,
			Name:       "乾(天)",
			Char:       "☰",
			YinYangArr: []int{Yang, Yang, Yang},
			Tips:       "乾三连",
		},
		Kun: {
			Id:         Kun,
			Name:       "坤(地)",
			Char:       "☷",
			YinYangArr: []int{Yin, Yin, Yin},
			Tips:       "坤六断",
		},
		Zhen: {
			Id:         Zhen,
			Name:       "震(雷)",
			Char:       "☳",
			YinYangArr: []int{Yin, Yin, Yang},
			Tips:       "震仰盂", // 下实上虚，形似口朝上的钵盂
		},
		Gen: {
			Id:         Gen,
			Name:       "艮(山)",
			Char:       "☶",
			YinYangArr: []int{Yang, Yin, Yin},
			Tips:       "艮覆碗", // 上实，下虚，形似扣着的碗
		},
		Li: {
			Id:         Li,
			Name:       "離(火)",
			Char:       "☲",
			YinYangArr: []int{Yang, Yin, Yang},
			Tips:       "离中虚",
		},
		Kan: {
			Id:         Kan,
			Name:       "坎(水)",
			Char:       "☵",
			YinYangArr: []int{Yin, Yang, Yin},
			Tips:       "坎中满",
		},
		Dui: {
			Id:         Dui,
			Name:       "兌(澤)",
			Char:       "☱",
			YinYangArr: []int{Yin, Yang, Yang},
			Tips:       "兑上缺",
		},
		Xun: {
			Id:         Xun,
			Name:       "巽xùn(風)",
			Char:       "☴",
			YinYangArr: []int{Yang, Yang, Yin},
			Tips:       "巽下断",
		},
	}

	LiuShiSiGuaMap = map[int]GuaImage{
		1: {
			Index: 1,
			Wai:   BaGuaMap[Qian],
			Nei:   BaGuaMap[Qian],
			Name:  "乾为天",
			DuYin: "qián",
			Text:  "乾：元、亨、利、贞。",
			Extra: "",
			YongYao: Yao{
				Text:      "用九，见群龙无首，吉。",
				ImageText: "《象》曰：用九，天德不可为首也。",
			},
			Short: "象征天，含有“健”的意思，“健”也称为乾卦的卦德。《周易集解》：“言天之体以健为用，运行不息，应化无穷，故圣人则之。欲使人法天之用，不法天之体，故名‘乾’，不名天也”。",
			Desc:  "　　《彖》曰：大哉乾元！万物资始，乃统天。云行雨施，品物流形。大明终始，六位时成，时乘六龙以御天。乾道变化，各正性命，保合大和，乃利贞。首出庶物，万国咸寧。中国古籍全录\n　　《象》曰：天行健，君子以自强不息。\n　　《象》曰：潜龙勿用，阳在下也；见龙在田，德施普也；终日乾乾，反復道也；或跃在渊，进无咎也；飞龙在天，大人造也；亢龙有悔，盈不可久也；用九，天德不可为首也。\n　　《文言》曰：\n　　元者，善之长也；亨者，嘉之会也；利者，义之和也；贞者，事之干也。君子体仁足以长人，嘉会足以合礼，利物足以和义，贞固足以干事。君子行此四德者，故曰：乾，元、亨、利、贞。中国古籍全录\n　　初九曰「潜龙勿用」，何谓也？子曰：「龙德而隱者也。不易乎世，不成乎名，遯世无闷，不见是而无闷。乐则行之，忧则违之，確乎其不可拔，潜龙也。」\n　　九二曰「见龙在田，利见大人」何谓也？子曰：「龙德而正中者也。庸言之信，庸行之谨。闲邪存其诚，善世而不伐，德博而化。《易》曰：『见龙在田，利见大人。』君德也。」\n　　九三曰「君子终日乾乾，夕惕若厉，无咎」何谓也？子曰：「君子进德修业。忠信，所以进德也；修辞立其诚，所以居业也。知至至之，可与几也。知终终之，可与存义也。是故居上位而不骄，在下位而不忧。故乾乾因其时而惕，虽危无咎矣。」\n　　九四曰「或跃在渊，无咎。」何谓也？子曰：「上下无常，非为邪也；进退无恒，非离群也。君子进德修业，欲及时也，故无咎。」\n　　九五曰「飞龙在天，利见大人」，何谓也？子曰：「同声相应，同气相求；水流湿，火就燥，云从龙，风从虎。圣人作而万物覩。本乎天者亲上，本乎地者亲下，则各从其类也。」\n　　上九曰「亢龙有悔」何谓也？子曰：「贵而无位，高而无民，贤人在下位而无辅，是以动而有悔也。」\n　　「潜龙勿用」，下也。「见龙在田」，时舍也。终日乾乾，行事也；或跃在渊，自试也；飞龙在天，上治也；亢龙有悔，穷之灾也。乾元用九，天下治也。\n　　潜龙勿用，阳气潜藏；见龙在田，天下文明。终日乾乾，与时偕行；或跃在渊，乾道乃革；飞龙在天，乃位乎天德；亢龙有悔，与时偕极；乾元用九，乃见天则。\n　　乾元者，始而亨者也；利贞者，性情也。乾始，能以美利利天下，不言所利，大矣哉！大哉乾乎！刚健中正，纯粹精也；六爻发挥，旁通情也；时乘六龙，以御天也。云行雨施，天下平也。\n　　君子以成德为行，日可见之行也。潜之为言也，隱而未见，行而未成，是以君子弗用也。\n　　君子学以聚之，问以辩之，宽以居之，仁以行之。《易》曰：「见龙在田，利见大人。」君德也。\n　　九三重刚而不中，上不在天，下不在田。故乾乾因其时而惕，虽危无咎矣。\n　　九四重刚而不中，上不在天，下不在田，中不在人，故或之。或之者，疑之也，故无咎。\n　　夫大人者，与天地合其德，与日月合其明，与四时合其序，与鬼神合其吉凶。先天而天弗违，后天而奉天时。天且弗违，而况於人乎？况於鬼神乎？\n \n　　亢之为言也，知进而不知退，知存而不知亡，知得而不知丧。其唯圣人乎？知进退存亡而不失其正者，其唯圣人乎！",
			Yao: [6]Yao{
				{
					Image:     YangYao,
					Text:      "初九，潜龙勿用。",
					ImageText: "《象》曰：潜龙勿用，阳在下也。",
				},
				{
					Image:     YangYao,
					Text:      "九二，见龙在田，利见大人。",
					ImageText: "《象》曰：见龙在田，德施普也。",
				},
				{
					Image:     YangYao,
					Text:      "九三，君子终日乾乾，夕惕若厉，无咎。",
					ImageText: "《象》曰：终日乾乾，反復道也。",
				},
				{
					Image:     YangYao,
					Text:      "九四，或跃在渊，无咎。",
					ImageText: "《象》曰：或跃在渊，进无咎也。",
				},
				{
					Image:     YangYao,
					Text:      "九五，飞龙在天，利见大人。",
					ImageText: "《象》曰：飞龙在天，大人造也。",
				},
				{
					Image:     YangYao,
					Text:      "上九，亢龙有悔。",
					ImageText: "《象》曰：亢龙有悔，盈不可久也。",
				},
			},
		},
		2: {
			Index: 2,
			Wai:   BaGuaMap[Kun],
			Nei:   BaGuaMap[Kun],
			Name:  "坤为地",
			DuYin: "kūn",
			Text:  "《坤》：元亨。利牝马之贞。君子有攸往，先迷，後得主，利。西南得朋，东北丧朋。安贞吉。",
			Extra: "",
			YongYao: Yao{
				Text:      "用六，利永贞。",
				ImageText: "《象》曰：用六“永贞”，以大终也。",
			},
			Short: "地载万物，也可使万物归隐，所以坤有归与藏的意思。坤卦是唯一的纯阴卦，是“至柔”、“至静”之卦。充分体现了大地之美，女性之美，阴柔之美。坤为大地，承载万物，顺应天时，化育万物，大地具有宽厚、包容、正直、宏大、安静的胸怀，值得我们好好学习。\n　　坤卦，坤为地卦",
			Desc:  "《彖》曰：至哉坤元，万物资生，乃顺承天。坤厚载物，德合无疆。含弘光大，品物咸亨。牝马地类，行地无疆，柔顺利贞。君子。君子攸行，先迷失道，後顺得常。西南得朋，乃与类行。东北丧朋，乃终有庆。安贞之吉，应地无疆。\n　　《象》曰：地势坤。君子以厚德载物。\n　　初六：履霜，坚冰至。\n　　《象》曰：“履霜坚冰”，阴始凝也，驯致其道，至坚冰也。\n　　六二，直方大，不习，无不利。\n　　《象》曰：六二之动，直以方也。“不习无不利”，地道光也。\n　　六三，含章可贞，或从王事，无成有终。\n　　《象》曰“含章可贞”，以时发也。“或従王事”，知光大也。\n　　六四，括囊，无咎无誉。\n　　《象》曰：“括囊无咎”，慎不害也。\n　　六五，黄裳，元吉。\n　　《象》曰：“黄裳元吉”，文在中也。\n　　上六，龙战于野，其血玄黄。\n　　《象》曰：“龙战于野”，共道穷也。\n　　用六，利永贞。\n　　《象》曰：用六“永贞”，以大终也。\n　　《文言》曰：坤至柔而动也刚，至静而德方，后得主而有常，含万物而化光。坤道其顺乎，承天而时行。积善之家必有余庆，积不善之家必有余殃。臣弑其君，子弑其父，非一朝一夕之故，其所由来者渐矣，由辩之不早辩也。《易》曰：“履霜，坚冰至”，盖言顺也。\n　　“直”其正也，“方”其义也。君子敬以直内，义以方外，敬义立而德不孤。“直、方、大，不习无不利”，则不疑其所行也。\n　　阴虽有美，“含”之以従王事，弗敢成也。地道也，妻道也，臣道也，地道无成而代有终也。\n　　天地变化，草木蕃。天地闭，贤人隐。《易》曰：“括囊，无咎无誉”，盖言谨也。\n　　君子黄中通理，正位居体，美在其中而畅于四支，发于事业，美之至也。\n　　阴疑于阳必战，为其嫌于无阳也，故称“龙”焉。犹未离其类也，故称“血”焉。夫玄黄者，天地之杂也，天玄而地黄。",
			Yao: [6]Yao{
				{
					Image:     YinYao,
					Text:      "初六：履霜，坚冰至。",
					ImageText: "《象》曰：“履霜坚冰”，阴始凝也，驯致其道，至坚冰也。",
				},
				{
					Image:     YinYao,
					Text:      "六二，直方大，不习，无不利。",
					ImageText: "《象》曰：六二之动，直以方也。“不习无不利”，地道光也。",
				},
				{
					Image:     YinYao,
					Text:      "六三，含章可贞，或从王事，无成有终。",
					ImageText: "《象》曰“含章可贞”，以时发也。“或従王事”，知光大也。",
				},
				{
					Image:     YinYao,
					Text:      "六四，括囊，无咎无誉。",
					ImageText: "《象》曰：“括囊无咎”，慎不害也。",
				},
				{
					Image:     YinYao,
					Text:      "六五，黄裳，元吉。",
					ImageText: "《象》曰：“黄裳元吉”，文在中也。",
				},
				{
					Image:     YinYao,
					Text:      "上六，龙战于野，其血玄黄。",
					ImageText: "《象》曰：“龙战于野”，共道穷也。",
				},
			},
		},
		3: {
			Index: 3,
			Wai:   BaGuaMap[Kan],
			Nei:   BaGuaMap[Zhen],
			Name:  "水雷屯",
			DuYin: "tún",
			Text:  "《屯》：元亨，利贞。勿用有攸往。利建侯。",
			Extra: "",
			Short: "屯：卦名，象征事物的初生与萌芽。屯者，物之初生也。故屯象征初生。像种子萌芽，破土而出，萌生、破土多有艰难，所以有“难”义。初生之物应当强根固本，不可轻动。但此时也是王者建功立业的时候，所以应该坚定信念产，积极进取，不可安居无事。",
			Desc:  "《彖》曰：屯，刚柔始交而难生。动乎险中，大亨贞。雷雨之动满盈，天造草昧。宜寻建侯而不宁。\n　　《象》曰：云雷，屯。君子以经纶。\n　　初九，磐桓，利居贞。利建侯。\n　　《象》曰：虽磐桓，志行正也。以贵下贱，大得民也。\n　　六二，屯如邅如，乘马班如。匪寇，婚媾。女子贞不字，十年乃字。\n　　《象》曰：六二之难，乘刚也。十年乃字，反常也。\n　　六三，即鹿无虞，惟入于林中，君子几不如舍，往吝。\n　　《象》曰：“即鹿无虞”，以従禽也。君子舍之，往吝穷也。\n　　六四，乘马班如，求婚媾。往吉，无不利。\n　　《象》曰：求而往，明也。\n　　九五，屯其膏，小，贞吉；大，贞凶。\n　　《象》曰：“屯其膏”，施未光也。\n　　上六，乘马班如，泣血涟如。\n　　《象》曰：“泣血涟如”，何可长也。",
			Yao: [6]Yao{
				{
					Image:     YangYao,
					Text:      "初九，磐桓，利居贞。利建侯。",
					ImageText: "《象》曰：虽磐桓，志行正也。以贵下贱，大得民也。",
				},
				{
					Image:     YinYao,
					Text:      "六二，屯如邅如，乘马班如。匪寇，婚媾。女子贞不字，十年乃字。",
					ImageText: "《象》曰：六二之难，乘刚也。十年乃字，反常也。",
				},
				{
					Image:     YinYao,
					Text:      "六三，即鹿无虞，惟入于林中，君子几不如舍，往吝。",
					ImageText: "《象》曰：“即鹿无虞”，以従禽也。君子舍之，往吝穷也。",
				},
				{
					Image:     YinYao,
					Text:      "六四，乘马班如，求婚媾。往吉，无不利。",
					ImageText: "《象》曰：求而往，明也。",
				},
				{
					Image:     YangYao,
					Text:      "九五，屯其膏，小，贞吉；大，贞凶。",
					ImageText: "《象》曰：“屯其膏”，施未光也。",
				},
				{
					Image:     YinYao,
					Text:      "上六，乘马班如，泣血涟如。",
					ImageText: "《象》曰：“泣血涟如”，何可长也。",
				},
			},
		},
		4: {
			Index: 4,
			Wai:   BaGuaMap[Kan],
			Nei:   BaGuaMap[Zhen],
			Name:  "山水蒙",
			DuYin: "méng",
			Text:  "《蒙》：亨。匪我求童蒙，童蒙求我。初筮告，再三渎，渎则不告。利贞。",
			Short: "艮为山，坎为泉，山下出泉。泉水始流出山，则必将渐汇成江河，正如蒙稚渐启，又山下有险，因为有险停止不前，所以蒙昧不明。事物发展的初期阶段，必然蒙昧，所以教育是当务之急，培养学生纯正无邪的品质，是治蒙之道。",
			Desc:  "《彖》曰：蒙，山下有险，险而止，蒙。“蒙亨”，以亨行，时中也。“匪我求童蒙，童蒙求我”。志应也。“初筮告”，以刚中也。“再三渎，渎则不告”，渎蒙也。蒙以养正，圣功也。\n　　《象》曰：山下出泉，蒙。君子以果行育德。\n　　初六，发蒙，利用刑人，用说桎梏，以往吝。\n　　《象》曰：“利用刑人”，以正法也。\n　　九二，包蒙，吉。纳妇，吉。子克家。\n　　《象》曰：“子克家”，刚柔节也。\n　　六三，勿用取女，见金夫，不有躬。无攸利。\n　　《象》曰：“勿用取女”，行不顺也。\n　　六四，困蒙，吝。\n　　《象》曰：“困蒙之吝”，独远实也。\n　　六五，童蒙，吉。\n　　《象》曰：“童蒙”之“吉”，顺以巽也。\n　　上九，击蒙，不利为寇，利御寇。\n　　《象》曰：“利”用“御寇”，上下顺也。",
			Yao: [6]Yao{
				{
					Image:     YinYao,
					Text:      "初六，发蒙，利用刑人，用说桎梏，以往吝。",
					ImageText: "《象》曰：“利用刑人”，以正法也。",
				},
				{
					Image:     YangYao,
					Text:      "九二，包蒙，吉。纳妇，吉。子克家。",
					ImageText: "《象》曰：“子克家”，刚柔节也。",
				},
				{
					Image:     YinYao,
					Text:      "六三，勿用取女，见金夫，不有躬。无攸利。",
					ImageText: "《象》曰：“勿用取女”，行不顺也。",
				},
				{
					Image:     YinYao,
					Text:      "六四，困蒙，吝。",
					ImageText: "《象》曰：“困蒙之吝”，独远实也。",
				},
				{
					Image:     YinYao,
					Text:      "六五，童蒙，吉。",
					ImageText: "《象》曰：“童蒙”之“吉”，顺以巽也。",
				},
				{
					Image:     YangYao,
					Text:      "上九，击蒙，不利为寇，利御寇。",
					ImageText: "《象》曰：“利”用“御寇”，上下顺也。",
				},
			},
		},
		5: {
			Index: 5,
			Wai:   BaGuaMap[Kan],
			Nei:   BaGuaMap[Qian],
			Name:  "水天需",
			DuYin: "xū",
			Text:  "《需》：有孚，光亨。贞吉，利涉大川。",
			Short: "需卦，等待之意。乾为天，坎为云，云气上集于天，待时降雨，为需。需象征需待。物初蒙稚，得养而成，因此也含有需待饮食的意思。需卦给我们的启示最重要的是无论在哪里都要耐心等待，顺应天道，伺机而动，是人生的一种智慧。",
			Desc:  "《彖》曰：“需”，须也。险在前也，刚健而不陷，其义不困穷矣。“需，有孚，光亨，贞吉”，位乎天位，以正中也。“利涉大川”，往有功也。\n　　《象》曰：云上于天，需。君子以饮食宴乐。",
			Yao: [6]Yao{
				{
					Image:     YangYao,
					Text:      "初九，需于郊，利用恒，无咎。",
					ImageText: "《象》曰：“需于郊”，不犯难行也。“利用恒无咎”，未失常也。",
				},
				{
					Image:     YangYao,
					Text:      "九二，需于沙，小有言，终吉。",
					ImageText: "《象》曰：“需于沙”，衍在中也。虽小有言，以终吉也。",
				},
				{
					Image:     YangYao,
					Text:      "九三，需于泥，致寇至。",
					ImageText: "",
				},
				{
					Image:     YinYao,
					Text:      "六四，需于血，出自穴。",
					ImageText: "《象》曰：“需于血，”顺以听也。",
				},
				{
					Image:     YangYao,
					Text:      "九五，需于酒食，贞吉。",
					ImageText: "《象》曰：“酒食贞吉”，以中正也。",
				},
				{
					Image:     YinYao,
					Text:      "上六，入于穴，有不速之客三人来，敬之终吉。",
					ImageText: "《象》曰：“不速之客来，敬之终吉”，虽不当位，未大失也。",
				},
			},
		},
		6: {
			Index: 6,
			Wai:   BaGuaMap[Qian],
			Nei:   BaGuaMap[Kan],
			Name:  "天水讼",
			DuYin: "sòng",
			Text:  "《讼》：有孚窒惕，中吉，终凶。利见大人。不利涉大川。",
			Short: "讼卦，象征争论、诉讼。乾为天，坎为水，天西转与水东流背向而行，像人与人不和而争辩。讼象征争辩争论，含诉讼之义。当不易和解时，便会导致诉讼。应该找有大德大才的人进行决断，不要逞强冒险。",
			Desc:  "《彖》曰：讼，上刚下险，险而健，讼。“讼有孚窒惕，中吉”，刚来而得中也。“终凶”，讼不可成也。“利见大人”，尚中正也。“不利涉大川”，入于渊也。\n　　《象》曰：天与水违行，讼。君子以作事谋始。\n　　初六，不永所事，小有言，终吉。\n　　《象》曰：“不永所事”，讼不可长也。虽“小有言”，其辩明也。\n　　九二，不克讼，归而逋。其邑人三百户，无眚。\n　　《象》曰：“不克讼”，归逋窜也。自下讼上，患至掇也。\n　　六三，食旧德，贞厉，终吉。或従王事，无成。\n　　《象》曰：食旧德，従上吉也。\n　　九四，不克讼，复既命渝。安贞吉。\n　　《象》曰：复即命渝，安贞不失也。\n　　九五：讼，元吉。\n　　《象》曰：“讼，元吉”以中正也。\n　　上九：或锡之鞶带，终朝三褫之。\n　　《象》曰：以讼受服，亦不足敬也。",
			Yao: [6]Yao{
				{
					Image:     YinYao,
					Text:      "初六，不永所事，小有言，终吉。",
					ImageText: "《象》曰：“不永所事”，讼不可长也。虽“小有言”，其辩明也。",
				},
				{
					Image:     YangYao,
					Text:      "九二，不克讼，归而逋。其邑人三百户，无眚。",
					ImageText: "《象》曰：“不克讼”，归逋窜也。自下讼上，患至掇也。",
				},
				{
					Image:     YinYao,
					Text:      "六三，食旧德，贞厉，终吉。或従王事，无成。",
					ImageText: "《象》曰：食旧德，従上吉也。",
				},
				{
					Image:     YangYao,
					Text:      "九四，不克讼，复既命渝。安贞吉。",
					ImageText: "《象》曰：复即命渝，安贞不失也。",
				},
				{
					Image:     YangYao,
					Text:      "九五：讼，元吉。",
					ImageText: "《象》曰：“讼，元吉”以中正也。",
				},
				{
					Image:     YangYao,
					Text:      "上九：或锡之鞶带，终朝三褫之。",
					ImageText: "《象》曰：以讼受服，亦不足敬也。",
				},
			},
		},
		7: {
			Index: 7,
			Wai:   BaGuaMap[Kun],
			Nei:   BaGuaMap[Kan],
			Name:  "地水师",
			DuYin: "shī",
			Text:  "《师》：贞丈人吉，无咎。",
			Short: "坤为地，坎为水，地中有水。地中众者，莫过于水。师为众，部属兵士众多的意思。持正的“仁义之师”，才可攻伐天下使百姓服从，用兵胜负在于择将选帅，持重老成的人统兵可获吉祥，这样才没有灾祸。",
			Desc:  "《彖》曰：师，众也。贞，正也。能以众正，可以王矣。刚中而应，行险而顺，以此毒天下，而民従之，吉又何咎矣。\n　　《象》曰：地中有水，师。君子以容民畜众。",
			Yao: [6]Yao{
				{
					Image:     YinYao,
					Text:      "初六，师出以律，否臧凶。",
					ImageText: "《象》曰：“师出以律，”失律凶也。",
				},
				{
					Image:     YangYao,
					Text:      "九二，在师中吉，无咎，王三锡命。",
					ImageText: "《象》曰：“在师中吉”，承天宠也。“王三锡命”，怀万邦也。",
				},
				{
					Image:     YinYao,
					Text:      "六三，师或舆尸，凶。",
					ImageText: "《象》曰：“师或舆尸”，大无功也。",
				},
				{
					Image:     YinYao,
					Text:      "六四，师左次，无咎。",
					ImageText: "《象》曰：“左次无咎”，未失常也。",
				},
				{
					Image:     YinYao,
					Text:      "六五，田有禽。利执言，无咎。长子帅师，弟子舆尸，贞凶。",
					ImageText: "《象》曰：“长子帅师”，以中行也。“弟子舆尸”，使不当也。",
				},
				{
					Image:     YinYao,
					Text:      "上六，大君有命，开国承家，小人勿用。",
					ImageText: "《象》曰：“大君有命”，以正功也。“小人勿用”，必乱邦也。",
				},
			},
		},
		8: {
			Index: 8,
			Wai:   BaGuaMap[Kan],
			Nei:   BaGuaMap[Kun],
			Name:  "水地比",
			DuYin: "bǐ",
			Text:  "《比》：吉。原筮，元，永贞，无咎。不宁方来，后夫凶。",
			Short: "比卦，亲比，亲密的辅佐。冲为地，坎为水，地上有水。水得地而蓄而流，地得水而柔而润，水与地亲密无间。比者，辅也，密也。故比象征亲密比辅。彼此能亲密比辅自然吉祥，但应比辅于守持正固而有德的长者，择善而从。",
			Desc:  "《彖》曰：比，吉也；比，辅也，下顺従也。“原筮，元永贞，无咎”，以刚中也。“不宁方来”，上下应也。“后夫凶”，其道穷也。\n　　《象》曰：地上有水，比。先王以建万国，亲诸侯。",
			Yao: [6]Yao{
				{
					Image:     YinYao,
					Text:      "初六，有孚比之，无咎。有孚盈缶，终来有它，吉。",
					ImageText: "《象》曰：比之初六，有它吉也。",
				},
				{
					Image:     YangYao,
					Text:      "六二，比之自内，贞吉。",
					ImageText: "《象》曰：“比之自内”，不自失也。",
				},
				{
					Image:     YinYao,
					Text:      "六三，比之匪人。",
					ImageText: "《象》曰：比之匪人”，不亦伤乎？斋",
				},
				{
					Image:     YinYao,
					Text:      "六四，外比之，贞吉。",
					ImageText: "《象》曰：外比于贤，以従上也。",
				},
				{
					Image:     YangYao,
					Text:      "九五，显比，王用三驱，失前禽，邑人不诫，吉。",
					ImageText: "《象》曰：“显比”之吉，位正中也。舍逆取顺，失前禽也。邑人不诫，上使中也。",
				},
				{
					Image:     YinYao,
					Text:      "上六，比之无首，凶。",
					ImageText: "《象》曰：“比之无首”，无所终也。",
				},
			},
		},
		9: {
			Index: 9,
			Wai:   BaGuaMap[Xun],
			Nei:   BaGuaMap[Qian],
			Name:  "风天小畜",
			DuYin: "xù",
			Text:  "《小畜》：亨。密云不雨。自我西郊。",
			Short: "乾为天，巽为风，风飘行天上，微畜而未下行。畜有畜聚、畜养、畜止之义。小畜象征小有畜聚，所畜甚微之象。以小畜大，以下济上，有利于刚大者之行。但阴气从西方升起聚阳甚微，不足以成雨。",
			Desc:  "《彖》曰：“小畜”，柔得位而上下应之，曰小畜。健而巽，刚中而志行，乃亨。“密云不雨”，尚往也。“自我西郊”，施未行也。\n　　《象》曰：风行天上，“小畜”。君子以懿文德。",
			Yao: [6]Yao{
				{
					Image:     YangYao,
					Text:      "初九，“复自道，何其咎？吉。",
					ImageText: "《象》曰：“复自道”，其义“吉”也。",
				},
				{
					Image:     YangYao,
					Text:      "九二，牵复，吉。",
					ImageText: "《象》曰：牵复在中，亦不自失也。",
				},
				{
					Image:     YangYao,
					Text:      "九三，舆说辐。夫妻反目。",
					ImageText: "《象》曰：“夫妻反目”，不能正室也。",
				},
				{
					Image:     YinYao,
					Text:      "六四，有孚，血去，惕出无咎。",
					ImageText: "《象》曰：“有孚惕出”，上合志也。",
				},
				{
					Image:     YangYao,
					Text:      "九五，有孚挛如，富以其邻。",
					ImageText: "《象》曰：“有孚挛如”，不独富也。",
				},
				{
					Image:     YangYao,
					Text:      "上九，既雨既处，尚德载。妇贞厉。月几望，君子征凶。",
					ImageText: "《象》曰：“既雨既处”，德积载也。“君子征凶”，有所疑也。",
				},
			},
		},
		10: {
			Index: 10,
			Wai:   BaGuaMap[Qian],
			Nei:   BaGuaMap[Dui],
			Name:  "天泽履",
			DuYin: "lǚ",
			Text:  "《履》：履虎尾，不咥人。亨。",
			Short: "履卦，象征履行、实践。乾为天，兑为泽，天在上，泽在下，为土下之正理。又乾为刚健，兑为和悦，有和悦应合刚健之象。履象征慎行，循礼而行的意思。遇事循礼慎行，即使有危也无害，所以诸事顺利。",
			Desc:  "《彖》曰：“履”，柔履刚也。说而应乎乾，是以“履虎尾，不咥人”。亨，刚中正，履帝位而不疚，光明也。\n　　《象》曰：上天下泽，“履”。君子以辨上下，定民志。",
			Yao: [6]Yao{
				{
					Image:     YangYao,
					Text:      "初九，素履往，无咎。",
					ImageText: "《象》曰：“素履之往”，独行愿也。",
				},
				{
					Image:     YangYao,
					Text:      "九二，履道坦坦，幽人贞吉。",
					ImageText: "《象》曰：“幽人贞吉”，中不自乱也。",
				},
				{
					Image:     YinYao,
					Text:      "六三，眇能视，跛能履，履虎尾，咥人，凶。武人为于大君。",
					ImageText: "《象》曰：“眇能视”，不足以有明也。“跛能履”，不足以与行也。“咥人之凶”，位不当也。“武人为于大君”，志刚也。",
				},
				{
					Image:     YangYao,
					Text:      "九四，履虎尾，愬愬，终吉。",
					ImageText: "《象》曰：“愬愬终吉”。志行也。",
				},
				{
					Image:     YangYao,
					Text:      "九五，夬履，贞厉。",
					ImageText: "《象》曰：“夬履贞厉”，位正当也。",
				},
				{
					Image:     YangYao,
					Text:      "上九，视履考祥，其旋元吉。",
					ImageText: "《象》曰：元吉在上，大有庆也。",
				},
			},
		},
		11: {
			Index: 11,
			Wai:   BaGuaMap[Kan],
			Nei:   BaGuaMap[Qian],
			Name:  "地天泰",
			DuYin: "tài",
			Text:  "《泰》：小往大来，吉，亨。",
			Short: "泰卦，象征通泰、平安。乾为天，坤为地，天气下降，地气上升，天地阴阳交合，万物的生养之道畅通。泰为通，泰象征通泰。即安泰亨通。通泰之时，阴者哀而往，阳者盛而来，所以既吉祥又顺利。",
			Desc:  "《彖》曰：“泰，小往大来。吉，亨。”则是天地交而万物通也，上下交而其志同也。内阳而外阴，内健而外顺，内君子而外小人，君子道长，小人道消也。\n　　《象》曰：天地交，泰。后以财成天地之道，辅相天地之宜，以左右民。",
			Yao: [6]Yao{
				{
					Image:     YangYao,
					Text:      "初九，拔茅茹以其汇。征吉。",
					ImageText: "《象》曰：“拔茅征吉”，志在外也。",
				},
				{
					Image:     YangYao,
					Text:      "九二，包荒，用冯河，不遐遗。朋亡，得尚于中行。",
					ImageText: "《象》曰：“包荒，得尚于中行”，以光大也。",
				},
				{
					Image:     YangYao,
					Text:      "九三，无平不陂，无往不复。艰贞无咎。勿恤其孚，于食有福。",
					ImageText: "《象》曰：“无往不复”，天地际也。",
				},
				{
					Image:     YinYao,
					Text:      "六四，翩翩，不富以其邻，不戒以孚。",
					ImageText: "《象》曰：“翩翩，不富”，皆失实也。“不戒以孚”，中心愿也。",
				},
				{
					Image:     YinYao,
					Text:      "六五，帝乙归妹，以祉元吉。",
					ImageText: "《象》曰：“以祉元吉”，中以行愿也。",
				},
				{
					Image:     YinYao,
					Text:      "上六，城复于隍，勿用师，自邑告命。贞吝。",
					ImageText: "《象》曰：“城复于隍”，其命乱也。",
				},
			},
		},
		12: {
			Index: 12,
			Wai:   BaGuaMap[Qian],
			Nei:   BaGuaMap[Kun],
			Name:  "天地否",
			DuYin: "pǐ",
			Text:  "《否》：否之匪人，不利君子贞，大往小来。",
			Short: "否卦，象征闭塞不通。坤下乾上，天气上升，地气下沉，天地阴阳二气互不交合，万物生养不得畅通，为否。否者，闭也。所以否象征否闭、闭塞。否闭之世，人道不通，天下无利。是小人得势，君子被排斥的形象。",
			Desc:  "《彖》曰：“否之匪人，不利君子贞，大往小来。”则是天地不交而万物不通也，上下不交而天下无邦也；内阴而外阳，内柔而外刚，内小人而外君子，小人道长，君子道消也。\n　　《象》曰：天地不交，“否”。君子以俭德辟难，不可荣以禄。",
			Yao: [6]Yao{
				{
					Image:     YinYao,
					Text:      "初六，拔茅茹以其汇。贞吉，亨。",
					ImageText: "《象》曰：“拔茅贞吉”，志在君也。",
				},
				{
					Image:     YinYao,
					Text:      "六二，包承，小人吉，大人否。亨。",
					ImageText: "《象》曰：“大人否亨”，不乱群也。",
				},
				{
					Image:     YinYao,
					Text:      "六三，包羞。",
					ImageText: "《象》曰：“包羞”，位不当也。",
				},
				{
					Image:     YangYao,
					Text:      "九四，有命，无咎，畴离祉。",
					ImageText: "《象》曰：“有命无咎”，志行也。",
				},
				{
					Image:     YangYao,
					Text:      "九五，休否，大人吉。其亡其亡，系于苞桑。",
					ImageText: "《象》曰：大人之吉，位正当也。",
				},
				{
					Image:     YangYao,
					Text:      "上九，倾否，先否后喜。",
					ImageText: "《象》曰：否终则倾，何可长也。",
				},
			},
		},
		13: {
			Index: 13,
			Wai:   BaGuaMap[Qian],
			Nei:   BaGuaMap[Li],
			Name:  "天火同人",
			DuYin: "",
			Text:  "《同人》：同人于野，亨。利涉大川。利君子贞。",
			Short: "同人卦，象征大家同心同德之意。离为火，乾为天，火光上升，即天、火相互亲和，为同人。象征和同于人。天下为公，有和睦、和平之义。促成世界大同，必须有广阔无私、光明磊落的境界，方顺利畅通，而这也是君子的正道。",
			Desc:  "《彖》曰：“同人”，柔得位得中，而应乎乾，曰同人。同人曰：“同人于野，亨。利涉大川”，乾行也。文明以健，中正而应，君子正也。唯君子为能通天下之志。\n　　《象》曰：天与火，同人。君子以类族辨物。",
			Yao: [6]Yao{
				{
					Image:     YangYao,
					Text:      "初九，同人于门，无咎。",
					ImageText: "《象》曰：“出门同人”，又谁咎也。",
				},
				{
					Image:     YinYao,
					Text:      "　六二，同人于宗，吝。",
					ImageText: "《象》曰：“同人于宗”，吝道也。",
				},
				{
					Image:     YangYao,
					Text:      "九三，伏戎于莽，升其高陵，三岁不兴。",
					ImageText: "《象》曰：“伏戎于莽”，敌刚也。“三岁不兴”，安行也。",
				},
				{
					Image:     YangYao,
					Text:      "九四，乘其墉，弗克攻，吉主。",
					ImageText: "《象》曰：“乘其墉”，义弗克也。其“吉”，则困而反则也。",
				},
				{
					Image:     YangYao,
					Text:      "九五，同人先号咷而后笑，大师克，相遇。",
					ImageText: "《象》曰：同人之先，以中直也。大师相遇，言相克也。",
				},
				{
					Image:     YangYao,
					Text:      "上九，同人于郊，无悔。",
					ImageText: "《象》曰：“同人于郊”，志未得也。",
				},
			},
		},
		14: {
			Index: 14,
			Wai:   BaGuaMap[Li],
			Nei:   BaGuaMap[Qian],
			Name:  "火天大有",
			DuYin: "",
			Text:  "《大有》：元亨。",
			Short: "大有卦，象征大有收获。离为火，乾为天，火焰高悬天上。即太阳当空照耀，大地五谷丰登，大获所有。故大有有收获之义，象征大获所有。又卦中一阴居尊位，获五阳之应，故为“大有”。",
			Desc:  "　《彖》曰：“大有”，柔得尊位大中，而上下应之，曰“大有”。其德刚健而文明，应乎天而时行，是以元亨。\n　　《象》曰：火在天上，“大有”。君子以遏恶扬善，顺天休命。",
			Yao: [6]Yao{
				{
					Image:     YangYao,
					Text:      "初九，无交害匪咎。艰则无咎。",
					ImageText: "《象》曰：大有初九，无交害也。",
				},
				{
					Image:     YangYao,
					Text:      "九二，大车以载，有攸往，无咎。",
					ImageText: "《象》曰：“大车以载”，积中不败也。",
				},
				{
					Image:     YangYao,
					Text:      "九三，公用亨于天子，小人弗克。",
					ImageText: "《象》曰：公用亨于天子，小人害也。",
				},
				{
					Image:     YangYao,
					Text:      "九四，匪其彭，无咎。",
					ImageText: "《象》曰：“匪其彭，无咎。”明辨晰也。",
				},
				{
					Image:     YinYao,
					Text:      "六五，厥孚交如威如，吉。",
					ImageText: "《象》曰：“厥孚交如”，信以发志也。“威如之吉”，易而无备也。",
				},
				{
					Image:     YangYao,
					Text:      "上九，自天祐之，吉，无不利。",
					ImageText: "《象》曰：大有上吉，自天祐也。",
				},
			},
		},
		15: {
			Index: 15,
			Wai:   BaGuaMap[Kun],
			Nei:   BaGuaMap[Kan],
			Name:  "地山谦",
			DuYin: "qiān",
			Text:  "《谦》：亨。君子有终。",
			Short: "谦卦，象征谦虚、谦逊。艮象征山、止，坤象征地、顺，地中有山。山体高大，但在地下，高能下，下谦之象。卑下之中，蕴其崇高，屈躬下物，先人后己，所以谦象征谦虚。如此谦虚地待物、待事，所以诸事顺利。但是只有君子才能始终保持谦虚的美德。",
			Desc:  "《彖》曰：谦，亨。天道下济而光明，地道卑而上行。天道亏盈而益谦，地道变盈而流谦，鬼神害盈而福谦，人道恶盈而好谦。谦，尊而光，卑而不可逾，君子之终也。\n　　《象》曰：地中有山，谦。君子以裒多益寡，称物平施。",
			Yao: [6]Yao{
				{
					Image:     YinYao,
					Text:      "初六，谦谦君子，用涉大川，吉。",
					ImageText: "《象》曰：“谦谦君子”，卑以自牧也。",
				},
				{
					Image:     YangYao,
					Text:      "六二，鸣谦，贞吉。",
					ImageText: "《象》曰：“鸣谦贞吉”，中心得也。",
				},
				{
					Image:     YangYao,
					Text:      "九三，劳谦君子，有终，吉。",
					ImageText: "《象》曰：“劳谦君子”，万民服也。",
				},
				{
					Image:     YinYao,
					Text:      "六四，无不利，捴谦。",
					ImageText: "《象》曰：“无不利，捴谦”，不违则也。",
				},
				{
					Image:     YinYao,
					Text:      "六五，不富以其邻，利用侵伐，无不利。",
					ImageText: "《象》曰：“利用侵伐”，征不服也。",
				},
				{
					Image:     YinYao,
					Text:      "上六，鸣谦，利用行师征邑国。",
					ImageText: "《象》曰：“鸣谦”，志未得也。“可用行师”，征邑国也。",
				},
			},
		},
		16: {
			Index: 16,
			Wai:   BaGuaMap[Zhen],
			Nei:   BaGuaMap[Kun],
			Name:  "雷地豫",
			DuYin: "yù",
			Text:  "利建侯行师。",
			Short: "",
			Desc:  "《彖》曰：豫，刚应而志行，顺以动，豫。豫顺以动，故天地如之，而况建侯行师乎？天地以顺动，故日月不过，而四时不忒。圣人以顺动，则刑罚清而民服，豫之时义大矣哉！\n\n《象》曰：雷出地奋，豫。先王以作乐崇德，殷荐之上帝，以配祖考。",
			Yao: [6]Yao{
				{
					Image:     YinYao,
					Text:      "初六：鸣豫，凶。",
					ImageText: "《象》曰：初六鸣豫，志穷凶也。",
				},
				{
					Image:     YinYao,
					Text:      "六二：介于石，不终日，贞吉。",
					ImageText: "《象》曰：不终日贞吉，以中正也。",
				},
				{
					Image:     YangYao,
					Text:      "六三：盱豫，悔，迟有悔。 ",
					ImageText: "《象》曰：盱豫不悔，位不当也。",
				},
				{
					Image:     YinYao,
					Text:      "九四：由豫，大有得，勿疑。朋盍簪。",
					ImageText: "《象》曰：由豫大有得，志大行也。",
				},
				{
					Image:     YinYao,
					Text:      "六五：贞疾，恒不死。",
					ImageText: "《象》曰：六五贞疾，乘刚也。恒不死，中未亡也。",
				},
				{
					Image:     YinYao,
					Text:      "上六：冥豫，成有渝。无咎。",
					ImageText: "《象》曰：“冥豫”在上，何可长也？",
				},
			},
		},
		17: {
			Index: 17,
			Wai:   BaGuaMap[Dui],  // 兑为泽
			Nei:   BaGuaMap[Zhen], // 震为雷
			Name:  "泽雷随",
			DuYin: "suí",
			Text:  "元亨，利贞，无咎。",
			Short: "",
			Desc:  "《彖》曰：随，刚来而下柔，动而说，随。大亨贞无咎，而天下随时，随时之义大矣哉！\n\n《象》曰：泽中有雷，随。君子以向晦入宴息。",
			Yao: [6]Yao{
				{
					Image:     YangYao,
					Text:      "初九：官有渝，贞吉，出门交有功。",
					ImageText: "《象》曰：官有渝，从正吉也。出门交有功，不失也。",
				},
				{
					Image:     YinYao,
					Text:      "六二：系小子，失丈夫。",
					ImageText: "《象》曰：系小子，弗兼与也。",
				},
				{
					Image:     YinYao,
					Text:      "六三：系丈夫，失小子，随有求，得。利居贞。",
					ImageText: "《象》曰：系丈夫，志舍下也。",
				},
				{
					Image:     YangYao,
					Text:      "九四：随有获，贞凶。有孚在道，以明，何咎？",
					ImageText: "《象》曰：随有获，其义凶也。有孚在道，明功也。",
				},
				{
					Image:     YangYao,
					Text:      "九五：孚于嘉，吉。",
					ImageText: "《象》曰：孚于嘉吉，位正中也。",
				},
				{
					Image:     YinYao,
					Text:      "上六：拘系之，乃从维之，王用亨于西山。",
					ImageText: "《象》曰：拘系之，上穷也。",
				},
			},
		},
		18: {
			Index: 18,
			Wai:   BaGuaMap[Gen], // 艮为山
			Nei:   BaGuaMap[Xun], // 巽为风
			Name:  "山风蛊",
			DuYin: "gǔ",
			Text:  "元亨。利涉大川，先甲三日，后甲三日。",
			Short: "",
			Desc:  "《彖》曰：蛊，刚上而柔下，巽而止，蛊。蛊，元亨而天下治也。利涉大川，往有事也。先甲三日，后甲三日，终则有始，天行也。\n\n《象》曰：山下有风，蛊。君子以振民育德。",
			Yao: [6]Yao{
				{
					Image:     YinYao,
					Text:      "初六：干父之蛊，有子，考无咎。厉，终吉。",
					ImageText: "《象》曰：干父之蛊，意承考也。",
				},
				{
					Image:     YangYao,
					Text:      "九二：干母之蛊，不可贞。",
					ImageText: "《象》曰：干母之蛊，得中道也。",
				},
				{
					Image:     YangYao,
					Text:      "九三：干父之蛊，小有悔，无大咎。",
					ImageText: "《象》曰：干父之蛊，终无咎也。",
				},
				{
					Image:     YinYao,
					Text:      "六四：裕父之蛊，往见吝。",
					ImageText: "《象》曰：裕父之蛊，往未得也。",
				},
				{
					Image:     YinYao,
					Text:      "六五：干父之蛊，用誉。",
					ImageText: "《象》曰：干父用誉，承以德也。",
				},
				{
					Image:     YangYao,
					Text:      "上九：不事王侯，高尚其事。 ",
					ImageText: "《象》曰：不事王侯，志可则也。",
				},
			},
		},
		19: {
			Index: 19,
			Wai:   BaGuaMap[Kun], // 坤为地
			Nei:   BaGuaMap[Dui], // 兑为泽
			Name:  "地泽临",
			DuYin: "",
			Text:  "元亨，利贞。至于八月有凶。",
			Short: "",
			Desc:  "《彖》曰：临，刚浸而长，说而顺，刚中而应。大亨以正，天之道也。至于八月有凶，消不久也。\n\n《象》曰：泽上有地，临。君子以教思无穷，容保民无疆。",
			Yao: [6]Yao{
				{
					Image:     YangYao,
					Text:      "初九：咸临，贞吉。",
					ImageText: "《象》曰：咸临贞吉，志行正也。",
				},
				{
					Image:     YangYao,
					Text:      "九二：咸临，吉，无不利。",
					ImageText: "《象》曰：咸临吉无不利，未顺命也。",
				},
				{
					Image:     YinYao,
					Text:      "六三：甘临，无攸利；既忧之，无咎。",
					ImageText: "《象》曰：甘临，位不当也。既忧之。咎不长也。",
				},
				{
					Image:     YinYao,
					Text:      "六四：至临，无咎。",
					ImageText: "《象》曰：至临无咎，位当也。",
				},
				{
					Image:     YinYao,
					Text:      "六五：知临，大君之宜，吉。",
					ImageText: "《象》曰：大君之宜，行中之谓也。",
				},
				{
					Image:     YinYao,
					Text:      "上六：敦临，吉，无咎。",
					ImageText: "《象》曰：敦临之吉，志在内也。",
				},
			},
		},
		20: {
			Index: 20,
			Wai:   BaGuaMap[Xun], // 巽为风
			Nei:   BaGuaMap[Kun], // 坤为地
			Name:  "风地观",
			DuYin: "",
			Text:  "盥而不荐。有孚颙若。",
			Short: "",
			Desc:  "《彖》曰：大观在上，顺而巽，中正以观天下，观。“盥而不荐，有孚颙若”，下观而化也。观天之神道，而四时不忒，圣人以神道设教，而天下服矣。\n\n《象》曰：风行地上，观。先王以省方观民设教。",
			Yao: [6]Yao{
				{
					Image:     YinYao,
					Text:      "初六：童观，小人无咎，君子吝。",
					ImageText: "《象》曰：初六童观，小人道也。",
				},
				{
					Image:     YinYao,
					Text:      "六二：窥观，利女贞。",
					ImageText: "《象》曰：窥观女贞，亦可丑也。",
				},
				{
					Image:     YinYao,
					Text:      "六三：观我生，进退。",
					ImageText: "《象》曰：观我生进退，未失道也。",
				},
				{
					Image:     YinYao,
					Text:      "六四：观国之光，利用宾于王。",
					ImageText: "《象》曰：观国之光，尚宾也。",
				},
				{
					Image:     YangYao,
					Text:      "九五：观我生，君子无咎。",
					ImageText: "《象》曰：观我生，观民也。",
				},
				{
					Image:     YangYao,
					Text:      "上九，观其生，君子无咎。",
					ImageText: "《象》曰：观其生，志未平也。",
				},
			},
		},
		21: {
			Index: 21,
			Wai:   BaGuaMap[Li],   // 离为火
			Nei:   BaGuaMap[Zhen], // 震为雷
			Name:  "火雷噬嗑",
			DuYin: "",
			Text:  "亨。利用狱。",
			Short: "",
			Desc:  "《彖》曰：颐中有物曰噬嗑。噬嗑而亨，刚柔分，动而明，雷电合而章。柔得中而上行，虽不当位，利用狱也。\n\n《象》曰：雷电，噬嗑。先王以明罚敕法。",
			Yao: [6]Yao{
				{
					Image:     YangYao,
					Text:      "初九：屦校灭趾，无咎。",
					ImageText: "《象》曰：屦校灭趾，不行也。",
				},
				{
					Image:     YinYao,
					Text:      "六二：噬肤灭鼻，无咎。",
					ImageText: "《象》曰：噬肤灭鼻，乘刚也。",
				},
				{
					Image:     YinYao,
					Text:      "六三：噬腊肉遇毒，小吝，无咎。",
					ImageText: "《象》曰：遇毒，位不当也。",
				},
				{
					Image:     YangYao,
					Text:      "九四：噬干胏，得金矢。利艰贞，吉。",
					ImageText: "《象》曰：利艰贞吉，未光也。",
				},
				{
					Image:     YinYao,
					Text:      "六五：噬干肉得黄金。贞厉，无咎。",
					ImageText: "《象》曰：“贞厉无咎”，得当也。",
				},
				{
					Image:     YangYao,
					Text:      "上九：何校灭耳，凶。 ",
					ImageText: "《象》曰：何校灭耳，聪不明也。",
				},
			},
		},
		22: {
			Index: 22,
			Wai:   BaGuaMap[Gen], // 艮为山
			Nei:   BaGuaMap[Li],  // 离为火
			Name:  "山火贲",
			DuYin: "",
			Text:  "亨。小利有攸往。",
			Short: "",
			Desc:  "《彖》曰：贲亨，柔来而文刚，故亨。分，刚上而文柔，故小利有攸往。刚柔交错，天文也。文明以止，人文也。观乎天文，以察时变；观乎人文，以化成天下。\n\n《象》曰：山下有火，贲。君子以明庶政，无敢折狱。",
			Yao: [6]Yao{
				{
					Image:     YangYao,
					Text:      "初九：贲其趾，舍车而徒。",
					ImageText: "《象》曰：舍车而徒，义弗乘也。",
				},
				{
					Image:     YinYao,
					Text:      "六二：贲其须。",
					ImageText: "《象》曰：贲其须，与上兴也。",
				},
				{
					Image:     YangYao,
					Text:      "九三：贲如，濡如，永贞吉。",
					ImageText: "《象》曰：永贞之吉，终莫之陵也。",
				},
				{
					Image:     YinYao,
					Text:      "六四：贲如皤如，白马翰如。匪寇，婚媾。",
					ImageText: "《象》曰：六四，当位疑也。匪寇婚媾，终无尤也。",
				},
				{
					Image:     YinYao,
					Text:      "六五：贲于丘园，束帛戋戋，吝，终吉。",
					ImageText: "《象》曰：六五之吉，有喜也。",
				},
				{
					Image:     YangYao,
					Text:      "上九：白贲，无咎。",
					ImageText: "《象》曰：白贲无咎，上得志也。",
				},
			},
		},
		23: {
			Index: 23,
			Wai:   BaGuaMap[Gen], // 艮为山
			Nei:   BaGuaMap[Kun], // 坤为地
			Name:  "山地剥",
			DuYin: "",
			Text:  "不利有攸往。",
			Short: "",
			Desc:  "《彖》曰：剥，剥也。柔变刚也。不利有攸往，小人长也。顺而止之，观象也。君子尚消息盈虚，天行也。\n\n《象》曰：出附于地，剥。上以厚下安宅。",
			Yao: [6]Yao{
				{
					Image:     YinYao,
					Text:      "初六：剥床以足，蔑贞凶。",
					ImageText: "《象》曰：剥床以足，以灭下也。",
				},
				{
					Image:     YinYao,
					Text:      "六二：剥床以辨，蔑贞凶。",
					ImageText: "《象》曰：剥床以辨，未有与也。",
				},
				{
					Image:     YinYao,
					Text:      "六三：剥之，无咎。",
					ImageText: "《象》曰：剥之无咎，失上下也。",
				},
				{
					Image:     YinYao,
					Text:      "六四：剥床以肤，凶。",
					ImageText: "《象》曰：“剥床以肤”，切近灾也。",
				},
				{
					Image:     YinYao,
					Text:      "六五：贯鱼以宫人宠，无不利。",
					ImageText: "《象》曰：“以宫人宠”，终无尤也。",
				},
				{
					Image:     YangYao,
					Text:      "上九：硕果不食，君子得舆，小人剥庐。",
					ImageText: "《象》曰：君子得舆，民所载也。小人剥庐，终不可用也",
				},
			},
		},
		24: {
			Index: 24,
			Wai:   BaGuaMap[Kun],  // 坤为地
			Nei:   BaGuaMap[Zhen], // 震为雷
			Name:  "地雷复",
			DuYin: "",
			Text:  "亨。出入无疾。朋来无咎。反复其道，七日来复，利有攸往。",
			Short: "",
			Desc:  "《彖》曰：复，亨。刚反，动而以顺行。是以出入无疾，朋来无咎。反复其道，七日来复，天行也。利有攸往，刚长也。复，其见天地之心乎。\n\n《象》曰：雷在地中，复。先王以至日闭关，商旅不行，后不省方。",
			Yao: [6]Yao{
				{
					Image:     YangYao,
					Text:      "初九：不远复，无祗悔，元吉。",
					ImageText: "《象》曰：不远之复，以修身也。",
				},
				{
					Image:     YinYao,
					Text:      "六二：休复，吉。",
					ImageText: "《象》曰：休复之吉，以下仁也。",
				},
				{
					Image:     YinYao,
					Text:      "六三：频复，厉，无咎。",
					ImageText: "《象》曰：频复之厉，义无咎也。",
				},
				{
					Image:     YinYao,
					Text:      "六四：中行独复。",
					ImageText: "《象》曰：中行独复，以从道也。",
				},
				{
					Image:     YinYao,
					Text:      "六五：敦复，无悔。",
					ImageText: "《象》曰：敦复无悔，中以自考也。",
				},
				{
					Image:     YinYao,
					Text:      "上六：迷复，凶，有灾眚。用行师，终有大败，以其国君凶，至于十年不克征。",
					ImageText: "《象》曰：迷复之凶，反君道也。",
				},
			},
		},
		25: {
			Index: 25,
			Wai:   BaGuaMap[Qian], // 乾为天
			Nei:   BaGuaMap[Zhen], // 震为雷
			Name:  "天雷无妄",
			DuYin: "",
			Text:  "元亨，利贞。其匪正有眚，不利有攸往。",
			Short: "",
			Desc:  "《彖》曰：无妄，刚自外来而为主于内，动而健，刚中而应。大亨以正，天之命也。“其匪正有眚，不利有攸往”，无妄之往何之矣？天命不祐，行矣哉！\n\n《象》曰：天下雷行，物与无妄。先王以茂对时育万物。",
			Yao: [6]Yao{
				{
					Image:     YangYao,
					Text:      "初九：无妄往，吉。",
					ImageText: "《象》曰：无妄之往，得志也。",
				},
				{
					Image:     YinYao,
					Text:      "六二：不耕获，不菑畬，则利用攸往。",
					ImageText: "《象》曰：不耕获，未富也。",
				},
				{
					Image:     YinYao,
					Text:      "六三：无妄之灾，或系之牛，行人之得，邑人之灾。",
					ImageText: "《象》曰：行人得牛，邑人灾也。",
				},
				{
					Image:     YangYao,
					Text:      "九四：可贞。无咎。",
					ImageText: "《象》曰：可贞无咎，固有之也。",
				},
				{
					Image:     YangYao,
					Text:      "九五：无妄之疾，勿药有喜。",
					ImageText: "《象》曰：无妄之药，不可试也。",
				},
				{
					Image:     YangYao,
					Text:      "上九：无妄行，有眚，无攸利。",
					ImageText: "《象》曰：无妄之行，穷之灾也。",
				},
			},
		},
		26: {
			Index: 26,
			Wai:   BaGuaMap[Gen],  // 艮为山
			Nei:   BaGuaMap[Qian], // 乾为天
			Name:  "山天大畜",
			DuYin: "",
			Text:  "利贞。不家食吉。利涉大川。",
			Short: "",
			Desc:  "《彖》曰：大畜，刚健笃实，辉光日新。其德刚上而尚贤，能止健，大正也。不家食吉，养贤也。利涉大川，应乎天也。\n\n《象》曰：天在山中，大畜。君子以多识前言往行，以畜其德。",
			Yao: [6]Yao{
				{
					Image:     YangYao,
					Text:      "初九：有厉，利已。",
					ImageText: "《象》曰：有厉利已，不犯灾也。",
				},
				{
					Image:     YangYao,
					Text:      "九二：舆说輹。",
					ImageText: "《象》曰：舆说輹，中无尤也。",
				},
				{
					Image:     YangYao,
					Text:      "九三：良马逐，利艰贞，曰闲舆卫，利有攸往。",
					ImageText: "《象》曰：利有攸往，上合志也。",
				},
				{
					Image:     YinYao,
					Text:      "六四：童牛之牿，元吉。",
					ImageText: "《象》曰：六四元吉，有喜也。",
				},
				{
					Image:     YinYao,
					Text:      "六五：豮豕之牙，吉。",
					ImageText: "《象》曰：六五之吉，有庆也。",
				},
				{
					Image:     YangYao,
					Text:      "上九：何天之衢，亨。",
					ImageText: "《象》曰：何天之衢，道大行也。",
				},
			},
		},
		27: {
			Index: 27,
			Wai:   BaGuaMap[Gen],  // 艮为山
			Nei:   BaGuaMap[Zhen], // 震为雷
			Name:  "山雷颐",
			DuYin: "",
			Text:  "贞吉。观颐，自求口实。",
			Short: "",
			Desc:  "《彖》曰：颐，贞吉，养正则吉也。观颐，观其所养也。自求口实，观其自养也。天地养万物，圣人养贤以及万民，颐之时大矣哉！\n\n《象》曰：山下有雷，颐。君子以慎言语，节饮食。",
			Yao: [6]Yao{
				{
					Image:     YangYao,
					Text:      "初九：舍尔灵龟，观我朵颐，凶。",
					ImageText: "《象》曰：观我朵颐，亦不足贵也。",
				},
				{
					Image:     YinYao,
					Text:      "六二：颠颐拂经于丘颐，征凶。",
					ImageText: "《象》曰：六二征凶，行失类也。",
				},
				{
					Image:     YinYao,
					Text:      "六三，拂颐，贞凶，十年勿用，无攸利。",
					ImageText: "《象》曰：十年勿用，道大悖也。",
				},
				{
					Image:     YinYao,
					Text:      "六四：颠颐，吉。虎视眈眈，其欲逐逐，无咎。",
					ImageText: "《象》曰：颠颐之吉，上施光也。",
				},
				{
					Image:     YinYao,
					Text:      "六五：拂经，居贞吉，不可涉大川。",
					ImageText: "《象》曰：居贞之吉，顺以从上也。",
				},
				{
					Image:     YangYao,
					Text:      "上九：由颐，厉，吉。利涉大川。",
					ImageText: "《象》曰：由颐厉吉，大有庆也。",
				},
			},
		},
		28: {
			Index: 28,
			Wai:   BaGuaMap[Li],   // 兑为泽
			Nei:   BaGuaMap[Zhen], // 巽为风
			Name:  "泽风大过",
			DuYin: "",
			Text:  "栋挠，利有攸往，亨。",
			Short: "",
			Desc:  "《彖》曰：大过，大者过也。栋挠，本末弱也。刚过而中，巽而说，行。利有攸往，乃亨。大过之时大矣哉！\n\n《象》曰：泽灭木，大过。君子以独立不惧，遯世无闷。",
			Yao: [6]Yao{
				{
					Image:     YinYao,
					Text:      "初六：藉用白茅，无咎。",
					ImageText: "《象》曰：藉用白茅，柔在下也。",
				},
				{
					Image:     YangYao,
					Text:      "九二：枯杨生稊，老夫得其女妻，无不利。",
					ImageText: "《象》曰：老夫女妻，过以相与也。",
				},
				{
					Image:     YangYao,
					Text:      "九三：栋桡，凶。",
					ImageText: "《象》曰：栋桡之凶，不可以有辅也。",
				},
				{
					Image:     YangYao,
					Text:      "九四：栋隆，吉。有它，吝。",
					ImageText: "《象》曰：栋隆之吉，不桡乎下也。",
				},
				{
					Image:     YangYao,
					Text:      "九五：枯杨生华，老妇得其士夫，无咎无誉。",
					ImageText: "《象》曰：枯杨生华，何可久也。老妇士夫，亦可丑也。",
				},
				{
					Image:     YinYao,
					Text:      "上六：过涉灭顶，凶。无咎。",
					ImageText: "《象》曰：过涉之凶，不可咎也。",
				},
			},
		},
		29: {
			Index: 29,
			Wai:   BaGuaMap[Kan], // 坎为水
			Nei:   BaGuaMap[Kan], // 坎为水
			Name:  "坎为水",
			DuYin: "",
			Text:  "有孚维心，亨。行有尚。",
			Short: "",
			Desc:  "《彖》曰：习坎，重险也。水流而不盈。行险而不失其信。维心亨，乃以刚中也。行有尚，往有功也。天险，不可升也。地险，山川丘陵也。王公设险以守其国。险之时用大矣哉！\n\n《象》曰：水洊至，习坎。君子以常德行，习教事。",
			Yao: [6]Yao{
				{
					Image:     YinYao,
					Text:      "初六：习坎，入于坎，窞，凶。",
					ImageText: "《象》曰：习坎入坎，失道，凶也。",
				},
				{
					Image:     YangYao,
					Text:      "九二：坎有险，求小得。",
					ImageText: "《象》曰：求小得，未出中也。",
				},
				{
					Image:     YinYao,
					Text:      "六三：来之坎，坎险且枕，入于坎，窞，勿用。",
					ImageText: "《象》曰：来之坎坎，终无功也。",
				},
				{
					Image:     YinYao,
					Text:      "六四，樽酒簋贰用缶，纳约自牖，终无咎。",
					ImageText: "《象》曰：樽酒簋贰，刚柔际也。",
				},
				{
					Image:     YangYao,
					Text:      "九五：坎不盈，祗既平，无咎。",
					ImageText: "《象》曰：坎不盈，中未大也。",
				},
				{
					Image:     YinYao,
					Text:      "上六：系用徽纆，窴于丛棘，三岁不得，凶。",
					ImageText: "《象》曰：上六失道，凶三岁也。",
				},
			},
		},
		30: {
			Index: 30,
			Wai:   BaGuaMap[Li], // 离为火
			Nei:   BaGuaMap[Li], // 离为火
			Name:  "离为火",
			DuYin: "",
			Text:  "利贞。亨。畜牝牛吉。",
			Short: "",
			Desc:  "《彖》曰：离，丽也。日月丽乎天，百谷草木丽乎土。重明以丽乎正，乃化成天下。柔丽乎中正，故亨，是以畜牝牛吉也。\n\n《象》曰：明两作，离。大人以继明照于四方。",
			Yao: [6]Yao{
				{
					Image:     YangYao,
					Text:      "初九：履错然，敬之无咎。",
					ImageText: "《象》曰：履错之敬，以辟咎也。",
				},
				{
					Image:     YinYao,
					Text:      "六二：黄离，元吉。",
					ImageText: "《象》曰：黄离元吉，得中道也。",
				},
				{
					Image:     YangYao,
					Text:      "九三：日昃之离，不鼓缶而歌，则大耋之嗟，凶。",
					ImageText: "《象》曰：日昃之离，何可久也？",
				},
				{
					Image:     YangYao,
					Text:      "九四：突如，其来如，焚如，死如，弃如。 ",
					ImageText: "《象》曰：突如其来如，无所容也。",
				},
				{
					Image:     YinYao,
					Text:      "六五：出涕沱若，戚嗟若，吉。",
					ImageText: "《象》曰：六五之吉，离王公也。",
				},
				{
					Image:     YangYao,
					Text:      "上九：王用出征，有嘉折首，获匪其丑，无咎。",
					ImageText: "《象》曰：王用出征，以正邦也。",
				},
			},
		},
		31: {
			Index: 31,
			Wai:   BaGuaMap[Zhen], // 兑为泽
			Nei:   BaGuaMap[Xun],  // 艮为山
			Name:  "泽山咸",
			DuYin: "",
			Text:  "亨。利贞。取女吉。",
			Short: "",
			Desc:  "《彖》曰：咸，感也。柔上而刚下，二气感应以相与。止而说，男下女，是以亨利贞、取女吉也。天地感而万物化生，圣人感人心而天下和平。观其所感，而天地万物之情可见矣。\n\n《象》曰：山上有泽，咸。君子以虚受人。",
			Yao: [6]Yao{
				{
					Image:     YinYao,
					Text:      "初六：咸其拇。",
					ImageText: "《象》曰：咸其拇，志在外也。",
				},
				{
					Image:     YinYao,
					Text:      "六二：咸其腓，凶。居吉。",
					ImageText: "《象》曰：虽凶居吉，顺不害也。",
				},
				{
					Image:     YangYao,
					Text:      "九三：咸其股，执其随，往吝。",
					ImageText: "《象》曰：咸其股，亦不处也。志在随人，所执下也。",
				},
				{
					Image:     YangYao,
					Text:      "九四：贞吉。悔亡。憧憧往来，朋从尔思。",
					ImageText: "《象》曰：贞吉悔亡，未感害也。憧憧往来，未光大也。",
				},
				{
					Image:     YangYao,
					Text:      "九五：咸其脢，无悔。",
					ImageText: "《象》曰：咸其脢，志末也。",
				},
				{
					Image:     YinYao,
					Text:      "上六：咸其辅颊舌。",
					ImageText: "《象》曰：咸其辅颊舌，滕口说也。",
				},
			},
		},
		32: {
			Index: 32,
			Wai:   BaGuaMap[Zhen], // 震为雷
			Nei:   BaGuaMap[Xun],  // 巽为风
			Name:  "雷风恒",
			DuYin: "",
			Text:  "亨。无咎。利贞。利有攸往。",
			Short: "",
			Desc:  "《彖》曰：恒，久也。刚上而柔下。雷风相与，巽而动，刚柔皆应，恒。恒亨无咎利贞，久于其道也。天地之道恒久而不已也。利有攸往，终则有始也。日月得天而能久照，四时变化而能久成。圣人久于其道而天下化成。观其所恒，而天地万物之情可见矣。\n\n《象》曰：雷风，恒。君子以立不易方。",
			Yao: [6]Yao{
				{
					Image:     YinYao,
					Text:      "初六：浚恒，贞凶，无攸利。",
					ImageText: "《象》曰：浚恒之凶，始求深也。",
				},
				{
					Image:     YangYao,
					Text:      "九二：悔亡。",
					ImageText: "《象》曰：九二悔亡，能久中也。",
				},
				{
					Image:     YangYao,
					Text:      "九三：不恒其德，或承之羞，贞吝。",
					ImageText: "《象》曰：不恒其德，无所容也。",
				},
				{
					Image:     YangYao,
					Text:      "九四：田无禽。",
					ImageText: "《象》曰：久非其位，安得禽也。",
				},
				{
					Image:     YinYao,
					Text:      "六五：恒其德，贞，妇人吉，夫子凶。",
					ImageText: "《象》曰：妇人贞吉，从一而终也。夫子制义，从妇凶也。",
				},
				{
					Image:     YangYao,
					Text:      "上六：振恒，凶。",
					ImageText: "《象》曰：振恒在上，大无功也。",
				},
			},
		},
		33: {
			Index: 33,
			Wai:   BaGuaMap[Qian], // 乾为天
			Nei:   BaGuaMap[Gen],  // 艮为山
			Name:  "天山遯",
			DuYin: "dùn",
			Text:  "亨。小利贞。",
			Short: "",
			Desc:  "《彖》曰：遯亨，遯而亨也。刚当位而应，与时行也。小利贞，浸而长也。遯之时义大矣哉！\n\n《象》曰：天下有山，遯。君子以远小人，不恶而严。",
			Yao: [6]Yao{
				{
					Image:     YinYao,
					Text:      "初六：遯尾，厉，勿用有攸往。",
					ImageText: "《象》曰：遯尾之厉，不往何灾也？",
				},
				{
					Image:     YinYao,
					Text:      "六二：执之用黄牛之革，莫之胜说。",
					ImageText: "《象》曰：执用黄牛，固志也。",
				},
				{
					Image:     YangYao,
					Text:      "九三：系遯，有疾厉，畜臣妾吉。",
					ImageText: "《象》曰：系遯之厉，有疾惫也。畜臣妾吉，不可大事也。",
				},
				{
					Image:     YangYao,
					Text:      "九四：好遯，君子吉，小人否。",
					ImageText: "《象》曰：君子好遯，小人否也。",
				},
				{
					Image:     YangYao,
					Text:      "九五：嘉遯，贞吉。",
					ImageText: "《象》曰：嘉遯贞吉，以正志也。",
				},
				{
					Image:     YangYao,
					Text:      "上九：肥遯，无不利。",
					ImageText: "《象》曰：肥遯无不利，无所疑也。",
				},
			},
		},
		34: {
			Index: 34,
			Wai:   BaGuaMap[Zhen], // 震为雷
			Nei:   BaGuaMap[Qian], // 乾为天
			Name:  "雷天大壮",
			DuYin: "",
			Text:  "利贞。",
			Short: "",
			Desc:  "《彖》曰：大壮，大者壮也。刚以动，故壮。大壮利贞，大者正也。正大，而天地之情可见矣。\n\n《象》曰：雷在天上，大壮。君子以非礼弗履。",
			Yao: [6]Yao{
				{
					Image:     YangYao,
					Text:      "初九：壮于趾，征凶，有孚。",
					ImageText: "《象》曰：壮于趾，其孚穷也。",
				},
				{
					Image:     YangYao,
					Text:      "九二：贞吉。",
					ImageText: "《象》曰：九二“贞吉”，以中也。",
				},
				{
					Image:     YangYao,
					Text:      "九三：小人用壮，君子用罔，贞厉。羝羊触藩，羸其角。",
					ImageText: "《象》曰：小人用壮，君子以罔也。",
				},
				{
					Image:     YangYao,
					Text:      "九四：贞吉，悔亡。藩决不羸，壮于大舆之輹。",
					ImageText: "《象》曰：藩决不羸，尚往也。",
				},
				{
					Image:     YinYao,
					Text:      "六五：丧羊于易，无悔。",
					ImageText: "《象》曰：丧羊于易，位不当也。",
				},
				{
					Image:     YinYao,
					Text:      "上六：羝羊触藩，不能退，不能遂，无攸利，艰则吉。",
					ImageText: "《象》曰：不能退，不能遂，不详也。艰则吉，咎不长也。",
				},
			},
		},
		35: {
			Index: 35,
			Wai:   BaGuaMap[Li],   // 离为火
			Nei:   BaGuaMap[Zhen], // 坤为地
			Name:  "火地晋",
			DuYin: "",
			Text:  "康侯用锡马蕃庶，昼日三接。",
			Short: "",
			Desc:  "《彖》曰：晋，进也，明出地上。顺而丽乎大明，柔进而上行，是以康侯用锡马蕃庶，昼日三接也。\n\n《象》曰：明出地上，晋。君子以自昭明德。",
			Yao: [6]Yao{
				{
					Image:     YinYao,
					Text:      "初六：晋如摧如，贞吉。罔孚，裕无咎。",
					ImageText: "《象》曰：晋如摧如，独行正也。裕无咎。未受命也。",
				},
				{
					Image:     YinYao,
					Text:      "六二：晋如，愁如，贞吉。受兹介福于，其王母。",
					ImageText: "《象》曰：受兹介福，以中正也。",
				},
				{
					Image:     YinYao,
					Text:      "六三：众允，悔亡。",
					ImageText: "《象》曰：众允之志，上行也。",
				},
				{
					Image:     YangYao,
					Text:      "九四：晋如鼫鼠，贞厉。",
					ImageText: "《象》曰：鼫鼠贞厉，位不当也。",
				},
				{
					Image:     YinYao,
					Text:      "六五：悔亡，失得，勿恤。往吉，无不利。",
					ImageText: "《象》》曰：失得勿恤，往有庆也。",
				},
				{
					Image:     YangYao,
					Text:      "上九：晋其角，维用伐邑，厉吉，无咎，贞吝。",
					ImageText: "《象》曰：维用伐邑，道未光也。",
				},
			},
		},
		36: {
			Index: 36,
			Wai:   BaGuaMap[Kun], // 坤为地
			Nei:   BaGuaMap[Li],  // 离为火
			Name:  "地火明夷",
			DuYin: "",
			Text:  "利艰贞。",
			Short: "",
			Desc:  "《彖》曰：明入地中，“明夷”。内文明而外柔顺，以蒙大难，文王以之。利艰贞，晦其明也，内难而能正其志，箕子以之。\n\n《象》曰：明入地中，明夷。君子以莅众用晦而明。",
			Yao: [6]Yao{
				{
					Image:     YangYao,
					Text:      "初九：明夷，于飞垂其翼。君子于行，三日不食。有攸往，主人有言。",
					ImageText: "《象》曰：君子于行，义不食也。",
				},
				{
					Image:     YinYao,
					Text:      "六二：明夷夷于左股，用拯马壮，吉。",
					ImageText: "《象》曰：六二之吉，顺以则也。",
				},
				{
					Image:     YangYao,
					Text:      "九三：明夷于南狩，得其大首，不可疾贞。",
					ImageText: "《象》曰：南狩之志，乃得大也。",
				},
				{
					Image:     YinYao,
					Text:      "六四：入于左腹，获明夷之心，于出门庭。",
					ImageText: "《象》曰：入于左腹，获心意也。",
				},
				{
					Image:     YinYao,
					Text:      "六五：箕子之明夷，利贞。",
					ImageText: "《象》曰：箕子之贞，明不可息也。",
				},
				{
					Image:     YangYao,
					Text:      "上六：不明，晦，初登于天，后入于地。",
					ImageText: "《象》曰：初登于天，照四国也。后入天地，失则也。",
				},
			},
		},
		37: {
			Index: 37,
			Wai:   BaGuaMap[Li],   //
			Nei:   BaGuaMap[Zhen], //
			Name:  "风火家人",
			DuYin: "",
			Text:  "利女贞。",
			Short: "",
			Desc:  "《彖》曰：家人，女正位乎内，男正位乎外。男女正，天地之大义也。家人有严君焉，父母之谓也。父父，子子，兄兄，弟弟，夫夫，妇妇，而家道正。正家而天下定矣。\n\n《象》曰：风自火出，家人。君子以言有物而行有恒。",
			Yao: [6]Yao{
				{
					Image:     YangYao,
					Text:      "初九：闲有家，悔亡。",
					ImageText: "《象》曰：闲有家，志未变也。",
				},
				{
					Image:     YinYao,
					Text:      "六二：无攸遂，在中馈，贞吉。",
					ImageText: "《象》曰：六二之吉，顺以巽也。",
				},
				{
					Image:     YinYao,
					Text:      "九三：家人嗃嗃，悔厉吉；妇子嘻嘻，终吝。",
					ImageText: "《象》曰：家人嗃嗃，未失也。妇子嘻嘻，失家节也。",
				},
				{
					Image:     YangYao,
					Text:      "九四：富家，大吉。",
					ImageText: "《象》曰：富家大吉，顺在位也。",
				},
				{
					Image:     YangYao,
					Text:      "九五：王假有家，勿恤，吉。",
					ImageText: "《象》曰：王假有家，交相爱也。",
				},
				{
					Image:     YangYao,
					Text:      "上九，有孚威如，终吉。",
					ImageText: "《象》曰：威如之吉，反身之谓也。",
				},
			},
		},
		38: {
			Index: 38,
			Wai:   BaGuaMap[Li],  // 离为火
			Nei:   BaGuaMap[Dui], // 兑为泽
			Name:  "火泽睽",
			DuYin: "",
			Text:  "小事吉。",
			Short: "",
			Desc:  "《彖》曰：睽，火动而上，泽动而下。二女同居，其志不同行。说而丽乎明，柔进而上行，得中而应乎刚，是以小事吉。天地睽而其事同也。男女睽而其志通也。万物睽而其事类也，睽之时用大矣哉！\n\n《象》曰：上火下泽，睽。君子以同而异。",
			Yao: [6]Yao{
				{
					Image:     YangYao,
					Text:      "初九：悔亡。丧马勿逐自复。见恶人无咎。",
					ImageText: "《象》曰：见恶人，以辟咎也。",
				},
				{
					Image:     YangYao,
					Text:      "九二：遇主于巷，无咎。",
					ImageText: "《象》曰：遇主于巷，未失道也。",
				},
				{
					Image:     YinYao,
					Text:      "六三：见舆曳，其牛掣，其人天且劓，无初有终。",
					ImageText: "《象》曰：见舆曳，位不当也。无初有终，遇刚也。",
				},
				{
					Image:     YangYao,
					Text:      "九四：睽孤遇元夫，交孚，厉，无咎。",
					ImageText: "《象》曰：交孚无咎，志行也。",
				},
				{
					Image:     YinYao,
					Text:      "六五：悔亡。厥宗噬肤，往何咎。",
					ImageText: "《象》曰：厥宗噬肤，往有庆也。",
				},
				{
					Image:     YangYao,
					Text:      "上九：睽孤见豕负途，载鬼一车，先张之弧，后说之弧，匪寇，婚媾。往遇雨则吉。",
					ImageText: "《象》曰：遇雨之吉，群疑亡也。",
				},
			},
		},
		39: {
			Index: 39,
			Wai:   BaGuaMap[Li],   //
			Nei:   BaGuaMap[Zhen], //
			Name:  "水山蹇",
			DuYin: "jiǎn",
			Text:  "利西南，不利东北。利见大人。贞吉。",
			Short: "",
			Desc:  "《彖》曰：蹇，难也，险在前也。见险而能止，知矣哉！蹇，利西南”，往得中也。不利东北，其道穷也。利见大人，往有功也。当位贞吉”，以正邦也。蹇之时用大矣哉！\n\n《象》曰：山上有水，蹇。君子以反身修德。",
			Yao: [6]Yao{
				{
					Image:     YinYao,
					Text:      "初六：往蹇来誉。",
					ImageText: "《象》曰：往蹇来誉，宜待也。",
				},
				{
					Image:     YinYao,
					Text:      "六二：王臣蹇蹇，匪躬之故。",
					ImageText: "《象》曰：王臣蹇蹇，终无尤也。",
				},
				{
					Image:     YangYao,
					Text:      "九三：往蹇来反。 ",
					ImageText: "《象》曰：往蹇来反，内喜之也。",
				},
				{
					Image:     YinYao,
					Text:      "六四：往蹇来连。",
					ImageText: "《象》曰：往蹇来连，当位实也。",
				},
				{
					Image:     YangYao,
					Text:      "九五：大蹇朋来。 ",
					ImageText: "《象》曰：大蹇朋来，以中节也。",
				},
				{
					Image:     YinYao,
					Text:      "上六：往蹇来硕，吉，利见大人。",
					ImageText: "《象》曰：往蹇来硕，志在内也。利见大人，以从贵也。",
				},
			},
		},
		40: {
			Index: 40,
			Wai:   BaGuaMap[Zhen], // 震为雷
			Nei:   BaGuaMap[Kan],  // 坎为水
			Name:  "雷水解",
			DuYin: "",
			Text:  "利西南。无所往，其来复吉。有攸往，夙吉。",
			Short: "",
			Desc:  "《彖》曰：解，险以动，动而免乎险，解。解，利西南，往得众也。其来复吉，乃得中也。有攸往夙吉，往有功也。天地解而雷雨作，雷雨作而百果草木皆甲坼。解之时大矣哉！\n\n《象》曰：雷雨作，解。君子以赦过宥罪。",
			Yao: [6]Yao{
				{
					Image:     YinYao,
					Text:      "初六：无咎。",
					ImageText: "《象》曰：刚柔之际，义无咎也。",
				},
				{
					Image:     YangYao,
					Text:      "九二：田获三狐，得黄矢，贞吉。",
					ImageText: "《象》曰：九二贞吉，得中道也。",
				},
				{
					Image:     YinYao,
					Text:      "六三：负且乘，致寇至，贞吝。",
					ImageText: "《象》曰：负且乘，亦可丑也。自我致戎，又谁咎也？",
				},
				{
					Image:     YangYao,
					Text:      "九四：解而拇，朋至斯孚。",
					ImageText: "《象》曰：解而拇，未当位也。",
				},
				{
					Image:     YinYao,
					Text:      "六五：君子维有解，吉，有孚于小人。",
					ImageText: "《象》曰：君子有解，小人退也。",
				},
				{
					Image:     YinYao,
					Text:      "上六：公用射隼于高墉之上，获之，无不利。",
					ImageText: "《象》曰：公用射隼，以解悖也。",
				},
			},
		},
		41: {
			Index: 41,
			Wai:   BaGuaMap[Gen], // 艮为山
			Nei:   BaGuaMap[Dui], // 兑为泽
			Name:  "山泽损",
			DuYin: "",
			Text:  "有孚，元吉，无咎。可贞，利有攸往。曷之用，二簋可用享。",
			Short: "",
			Desc:  "《彖》曰：损，损下益上，其道上行。损而有孚，元吉，无咎，可贞，利有攸往，曷之用？二簋可用享。二簋应有时。损刚益柔有时，损益盈虚，与时偕行。\n\n《象》曰：山下有泽，损。君子以惩忿窒欲。",
			Yao: [6]Yao{
				{
					Image:     YangYao,
					Text:      "初九：已事遄往，无咎。酌损之。",
					ImageText: "《象》曰：已事遄往，尚合志也。",
				},
				{
					Image:     YangYao,
					Text:      "九二：利贞。征凶，弗损，益之。 ",
					ImageText: "《象》曰：九二利贞，中以为志也。",
				},
				{
					Image:     YinYao,
					Text:      "六三：三人行则损一人，一人行则得其友。",
					ImageText: "《象》曰：一人行，三则疑也。",
				},
				{
					Image:     YinYao,
					Text:      "六四：损其疾，使遄有喜，无咎。",
					ImageText: "《象》曰：损其疾，亦可喜也。",
				},
				{
					Image:     YinYao,
					Text:      "六五：或益之十朋之龟，弗克违，元吉。",
					ImageText: "《象》曰：六五元吉，自上祐也。",
				},
				{
					Image:     YangYao,
					Text:      "上九：弗损，益之，无咎，贞吉，利有攸往，得臣无家。",
					ImageText: "《象》曰：弗损，益之，大得志也。",
				},
			},
		},
		42: {
			Index: 42,
			Wai:   BaGuaMap[Xun],  // 巽为风
			Nei:   BaGuaMap[Zhen], // 震为雷
			Name:  "风雷益",
			DuYin: "",
			Text:  "利有攸往。利涉大川。",
			Short: "",
			Desc:  "《彖》曰：益，损上益下，民说无疆。自上下下，其道大光。利有攸往，中正有庆。利涉大川，木道乃行。益动而巽，日进无疆。天施地生，其益无方。凡益之道，与时偕行。\n\n《象》曰：风雷，益。君子以见善则迁，有过则改。",
			Yao: [6]Yao{
				{
					Image:     YangYao,
					Text:      "初九：利用为大作，元吉，无咎。",
					ImageText: "《象》曰：元吉无咎，下不厚事也。",
				},
				{
					Image:     YinYao,
					Text:      "六二：或益之十朋之龟，弗克违。永贞吉。王用享于帝，吉。",
					ImageText: "《象》曰：或益之，自外来也。",
				},
				{
					Image:     YinYao,
					Text:      "六三：益之用凶事，无咎。有孚。中行告公用圭。",
					ImageText: "《象》曰：益用凶事，固有之也。",
				},
				{
					Image:     YinYao,
					Text:      "六四：中行告公，从，利用为依迁国。",
					ImageText: "《象》曰：告公从，以益志也。",
				},
				{
					Image:     YangYao,
					Text:      "九五：有孚惠心，勿问，元吉。有孚，惠我德。",
					ImageText: "《象》曰：有孚惠心，勿问之矣。惠我德，大得志也。",
				},
				{
					Image:     YangYao,
					Text:      "上九：莫益之，或击之，立心勿恒，凶。",
					ImageText: "《象》曰：莫益之，偏辞也。或击之，自外来也。",
				},
			},
		},
		43: {
			Index: 43,
			Wai:   BaGuaMap[Dui],  // 兑为泽
			Nei:   BaGuaMap[Qian], // 乾为天
			Name:  "泽天夬",
			DuYin: "guài",
			Text:  "扬于王庭，孚号。有厉，告自邑。不利即戎，利有攸往。",
			Short: "",
			Desc:  "《彖》曰：夬，决也，刚决柔也。健而说，决而和。扬于王庭，柔乘五刚也。孚号有厉，其危乃光也。告自邑，不利即戎，所尚乃穷也。利有攸往，刚长乃终也。\n\n《象》曰：泽上于天，夬。君子以施禄及下，居德则忌。",
			Yao: [6]Yao{
				{
					Image:     YangYao,
					Text:      "初九：壮于前趾，往不胜，为咎。",
					ImageText: "《象》曰：不胜而往，咎也。",
				},
				{
					Image:     YangYao,
					Text:      "九二：惕号，莫夜有戎，勿恤。",
					ImageText: "《象》曰：有戎勿恤，得中道也。",
				},
				{
					Image:     YangYao,
					Text:      "九三：壮于頄，有凶。君子夬夬独行，遇雨若濡，有愠无咎。",
					ImageText: "《象》曰：君子夬夬，终无咎也。",
				},
				{
					Image:     YangYao,
					Text:      "九四：臀无肤，其行次且。牵羊悔亡，闻言不信。",
					ImageText: "《象》曰：其行次且，位不当也。闻言不信，聪不明也。",
				},
				{
					Image:     YangYao,
					Text:      "九五：苋陆夬夬中行，无咎。",
					ImageText: "《象》曰：中行无咎，中未光也。",
				},
				{
					Image:     YinYao,
					Text:      "上六：无号，终有凶。",
					ImageText: "《象》曰：无号之凶，终不可长也。",
				},
			},
		},
		44: {
			Index: 44,
			Wai:   BaGuaMap[Qian], // 乾为天
			Nei:   BaGuaMap[Xun],  // 巽为风
			Name:  "天风姤",
			DuYin: "gòu、dù", // 读作姤（gòu或dù），读作姤（gòu）时同“遘”，也指善，美好 [5]。读作姤（dù）时古同“妒”，忌妒；忌恨。
			Text:  "女壮，勿用取女。",
			Short: "",
			Desc:  "《彖》曰：姤，遇也，柔遇刚也。勿用取女，不可与长也。天地相遇，品物咸章也。刚遇中正，天下大行也。姤之时义大矣哉！\n\n《象》曰：天下有风，姤。后以施命诰四方。",
			Yao: [6]Yao{
				{
					Image:     YinYao,
					Text:      "初六：系于金柅，贞吉。有攸往，见凶，羸豕孚蹢躅。",
					ImageText: "《象》曰：系于金柅，柔道牵也。",
				},
				{
					Image:     YangYao,
					Text:      "九二：包有鱼，无咎，不利宾。",
					ImageText: "《象》曰：包有鱼，义不及宾也。",
				},
				{
					Image:     YangYao,
					Text:      "九三：臀无肤，其行次且，厉，无大咎。",
					ImageText: "《象》曰：其行次且，行未牵也。",
				},
				{
					Image:     YangYao,
					Text:      "九四：包无鱼，起凶。",
					ImageText: "《象》曰：无鱼之凶，远民也。",
				},
				{
					Image:     YangYao,
					Text:      "九五：以杞包瓜，含章，有陨自天。",
					ImageText: "《象》曰：九五含章，中正也。有陨自天，志不舍命也。",
				},
				{
					Image:     YangYao,
					Text:      "上九：姤其角，吝，无咎。",
					ImageText: "《象》曰：姤其角，上穷吝也。",
				},
			},
		},
		45: {
			Index: 45,
			Wai:   BaGuaMap[Dui], // 兑为泽
			Nei:   BaGuaMap[Kun], // 坤为地
			Name:  "泽地萃",
			DuYin: "cuì",
			Text:  "亨，王假有庙。利见大人。亨，利贞，用大牲吉。利有攸往。",
			Short: "",
			Desc:  "《彖》曰：萃，聚也。顺以说，刚中而应，故聚也。王假有庙，致孝享也。利见大人亨，聚以正也。用大牲吉，利有攸往，顺天命也。观其所聚，而天地万物之情可见矣。\n\n《象》曰：泽上于地，萃。君子以除戎器，戒不虞。",
			Yao: [6]Yao{
				{
					Image:     YinYao,
					Text:      "初六：有孚不终，乃乱乃萃，若号，一握为笑，勿恤，往无咎。",
					ImageText: "《象》曰：乃乱乃萃，其志乱也。",
				},
				{
					Image:     YinYao,
					Text:      "六二：引吉，无咎，孚乃利用禴。",
					ImageText: "《象》曰：引吉无咎，中未变也。",
				},
				{
					Image:     YinYao,
					Text:      "六三：萃如嗟如，无攸利，往无咎，小吝。",
					ImageText: "《象》曰：往无咎，上巽也。",
				},
				{
					Image:     YangYao,
					Text:      "九四：大吉无咎。",
					ImageText: "《象》曰：大吉无咎，位不当也。",
				},
				{
					Image:     YangYao,
					Text:      "九五：萃有位，无咎。匪孚，元永贞，悔亡。",
					ImageText: "《象》曰：萃有位，志未光也。",
				},
				{
					Image:     YinYao,
					Text:      "上六：赍咨涕洟，无咎。",
					ImageText: "《象》曰：赍咨涕洟，未安上也。",
				},
			},
		},
		46: {
			Index: 46,
			Wai:   BaGuaMap[Kun], // 坤为地
			Nei:   BaGuaMap[Xun], // 巽为风
			Name:  "地风升",
			DuYin: "",
			Text:  "元亨。用见大人，勿恤。南征吉。",
			Short: "",
			Desc:  "《彖》曰：柔以时升，巽而顺，刚中而应，是以大亨，用见大人勿恤，有庆也。南征吉，志行也。\n\n《象》曰：地中生木，升。君子以顺德，积小以高大。",
			Yao: [6]Yao{
				{
					Image:     YinYao,
					Text:      "初六：允升，大吉。",
					ImageText: "《象》曰：允升大吉，上合志也。",
				},
				{
					Image:     YangYao,
					Text:      "九二：孚乃利用禴，无咎。",
					ImageText: "《象》曰：九二之孚，有喜也。",
				},
				{
					Image:     YangYao,
					Text:      "九三：升虚邑。",
					ImageText: "《象》曰：升虚邑，无所疑也。",
				},
				{
					Image:     YinYao,
					Text:      "六四：王用亨于岐山，吉，无咎。",
					ImageText: "《象》曰：王用亨于岐山，顺事也。",
				},
				{
					Image:     YinYao,
					Text:      "六五：贞吉，升阶。",
					ImageText: "《象》曰：贞吉升阶，大得志也。",
				},
				{
					Image:     YangYao,
					Text:      "上六：冥升，利于不息之贞。",
					ImageText: "《象》曰：冥升在上，消不富也。",
				},
			},
		},
		47: {
			Index: 47,
			Wai:   BaGuaMap[Dui], // 兑为泽
			Nei:   BaGuaMap[Kan], // 坎为水
			Name:  "泽水困",
			DuYin: "",
			Text:  "亨。贞大人吉，无咎。有言不信。",
			Short: "",
			Desc:  "《彖》曰：困，刚揜也。险以说，因而不失其所，亨，其唯君子乎。贞大人吉，以刚中也。有言不信，尚口乃穷也。\n\n《象》曰：泽无水，困。君子以致命遂志。",
			Yao: [6]Yao{
				{
					Image:     YinYao,
					Text:      "初六：臀困于株木，入于幽谷，三岁不觌。",
					ImageText: "《象》曰：入于幽谷，幽不明也。",
				},
				{
					Image:     YangYao,
					Text:      "九二：困于酒食，朱绂方来。利用享祀。征凶，无咎。",
					ImageText: "《象》曰：困于酒食，中有庆也。",
				},
				{
					Image:     YinYao,
					Text:      "六三：困于石，据于蒺藜，入于其宫，不见其妻，凶。",
					ImageText: "《象》曰：据于蒺藜，乘刚也。入于其宫，不见其妻，不祥也。",
				},
				{
					Image:     YangYao,
					Text:      "九四：来徐徐，困于金车，吝，有终。",
					ImageText: "《象》曰：来徐徐，志在下也。虽不当位，有与也。",
				},
				{
					Image:     YangYao,
					Text:      "九五：劓刖，困于赤绂，乃徐有说，利用祭祀。",
					ImageText: "《象》曰：劓刖，志未得也。乃徐有说，以中直也。利用祭祀，受福也。",
				},
				{
					Image:     YinYao,
					Text:      "上六：困于葛藟，于臲(niè)，曰动悔有悔，征吉。", // 臲，读作niè，同“絅”，禅衣。如：衣锦臲衣
					ImageText: "《象》曰：困于葛藟，未当也。动悔有悔，吉行也。",
				},
			},
		},
		48: {
			Index: 48,
			Wai:   BaGuaMap[Kan], // 坎为水
			Nei:   BaGuaMap[Xun], // 巽为风
			Name:  "水风井",
			DuYin: "",
			Text:  "改邑不改井，无丧无得。往来井井。汔至，亦未繘井，羸其瓶，凶。",
			Short: "",
			Desc:  "《彖》曰：巽乎水而上水，井。井养而不穷也。改邑不改井，乃以刚中也。汔至亦未繘，井，未有功也。羸其瓶，是以凶也。\n\n《象》曰：木上有水，井。君子以劳民劝相。",
			Yao: [6]Yao{
				{
					Image:     YinYao,
					Text:      "初六：井泥不食。旧井无禽。",
					ImageText: "《象》曰：井泥不食，下也。旧井无禽，时舍也。",
				},
				{
					Image:     YangYao,
					Text:      "九二：井谷射鲋，瓮敝漏。",
					ImageText: "《象》曰：井谷射鲋，无与也。",
				},
				{
					Image:     YangYao,
					Text:      "九三：井渫不食，为我心恻。可用汲，王明并受其福。",
					ImageText: "《象》曰：井渫不食，行恻也。求王明，受福也。",
				},
				{
					Image:     YinYao,
					Text:      "六四：井甃，无咎。",
					ImageText: "《象》曰：井甃无咎，修井也。",
				},
				{
					Image:     YangYao,
					Text:      "九五：井洌，寒泉食。",
					ImageText: "《象》曰：寒泉之食，中正也。",
				},
				{
					Image:     YinYao,
					Text:      "上六：井收勿幕，有孚元吉。",
					ImageText: "《象》曰：元吉在上，大成也。",
				},
			},
		},
		49: {
			Index: 49,
			Wai:   BaGuaMap[Dui], // 兑为泽
			Nei:   BaGuaMap[Li],  // 离为火
			Name:  "泽火革",
			DuYin: "",
			Text:  "已日乃孚。元亨。利贞，悔亡。",
			Short: "",
			Desc:  "《彖》曰：革，水火相息，二女同居，其志不相得曰革。已日乃孚，革而信之。文明以说，大亨以正。革而当，其悔乃亡。天地革而四时成，汤武革命，顺乎天而应乎人。革之时大矣哉！\n\n《象》曰：泽中有火，革。君子以治历明时。",
			Yao: [6]Yao{
				{
					Image:     YangYao,
					Text:      "初九：巩用黄牛之革。",
					ImageText: "《象》曰：巩用黄牛，不可以有为也。",
				},
				{
					Image:     YinYao,
					Text:      "六二：巳日乃革之，征吉，无咎。",
					ImageText: "《象》曰：巳日革之，行有嘉也。",
				},
				{
					Image:     YangYao,
					Text:      "九三：征凶。贞厉。革言三就，有孚。",
					ImageText: "《象》曰：革言三就，又何之矣。",
				},
				{
					Image:     YangYao,
					Text:      "九四：悔亡。有孚改命，吉。",
					ImageText: "《象》曰：改命之吉，信志也。",
				},
				{
					Image:     YangYao,
					Text:      "九五：大人虎变，未占有孚。",
					ImageText: "《象》曰：大人虎变，其文炳也。",
				},
				{
					Image:     YangYao,
					Text:      "上六：君子豹变，小人革面，征凶，居贞吉。",
					ImageText: "《象》曰：君子豹变，其文蔚也。小人革面，顺以从君也。",
				},
			},
		},
		50: {
			Index: 50,
			Wai:   BaGuaMap[Li],   // 离为火
			Nei:   BaGuaMap[Zhen], // 巽为风
			Name:  "火风鼎",
			DuYin: "",
			Text:  "元吉，亨。",
			Short: "",
			Desc:  "《彖》曰：鼎，象也。以木巽火，亨饪也。圣人亨以享上帝，而大亨以养圣贤。巽而耳目聪明，柔进而上行，得中而应乎刚，是以元亨。\n\n《象》曰：木上有火，鼎。君子以正位凝命。\n\n",
			Yao: [6]Yao{
				{
					Image:     YinYao,
					Text:      "初六：鼎颠趾，利出否。得妾以其子，无咎。",
					ImageText: "《象》曰：鼎颠趾，未悖也。利出否，以从贵也。",
				},
				{
					Image:     YangYao,
					Text:      "九二：鼎有实，我仇有疾，不我能即，吉。",
					ImageText: "《象》曰：鼎有实，慎所之也。我仇有疾，终无尤也。",
				},
				{
					Image:     YangYao,
					Text:      "九三：鼎耳革，其行塞，雉膏不食，方雨，亏悔，终吉。",
					ImageText: "《象》曰：鼎耳革，失其义也。",
				},
				{
					Image:     YangYao,
					Text:      "九四：鼎折足，覆公餗，其形渥，凶。",
					ImageText: "《象》曰：覆公餗，信如何也。",
				},
				{
					Image:     YinYao,
					Text:      "六五：鼎黄耳金铉，利贞。",
					ImageText: "《象》曰：鼎黄耳，中以为实也。",
				},
				{
					Image:     YangYao,
					Text:      "上九：鼎玉铉，大吉，无不利。",
					ImageText: "《象》曰：玉铉在上，刚柔节也。",
				},
			},
		},
		51: {
			Index: 51,
			Wai:   BaGuaMap[Zhen], // 震为雷
			Nei:   BaGuaMap[Zhen], // 震为雷
			Name:  "震为雷",
			DuYin: "",
			Text:  "亨。震来虩虩(xì)，笑言哑哑，震惊百里，不丧匕鬯。",
			Short: "",
			Desc:  "《彖》曰：震，亨。“震来虩虩”，恐致福也。“笑言哑哑”，后有则也。震惊百里，惊远而惧迩也。不丧匕鬯，出可以守宗庙社稷，以为祭主也。\n\n《象》曰：洊雷，震。君子以恐惧修省。",
			Yao: [6]Yao{
				{
					Image:     YangYao,
					Text:      "初九：震来虩虩，后笑言哑哑，吉。", // 读作虩（xì），〔虩虩〕恐惧的样子，如“震来虩虩，笑言哑哑。
					ImageText: "《象》曰：震来虩虩，恐致福也。笑言哑哑，后有则也。",
				},
				{
					Image:     YinYao,
					Text:      "六二：震来厉，亿丧贝，跻于九陵，勿逐，七日得。",
					ImageText: "《象》曰：震来厉，乘刚也。",
				},
				{
					Image:     YinYao,
					Text:      "六三：震苏苏，震行无眚。",
					ImageText: "《象》曰：震苏苏，位不当也。",
				},
				{
					Image:     YangYao,
					Text:      "九四：震遂泥。",
					ImageText: "《象》曰：震遂泥，未光也。",
				},
				{
					Image:     YinYao,
					Text:      "六五：震往来，厉，意无丧，有事。",
					ImageText: "《象》曰：震往来厉，危行也。其事在中，大无丧也。",
				},
				{
					Image:     YinYao,
					Text:      "上六：震索索，视矍矍(jué)，征凶。震不于其躬，于其邻，无咎。婚媾有言。", // 矍是一个汉字，字音：jué，基本意思是形容老人目光炯炯、精神健旺。
					ImageText: "《象》曰：震索索，中未得也。虽凶无咎，畏邻戒也。",
				},
			},
		},
		52: {
			Index: 52,
			Wai:   BaGuaMap[Gen], // 艮为山
			Nei:   BaGuaMap[Gen], // 艮为山
			Name:  "艮为山",
			DuYin: "",
			Text:  "艮其背，不获其身，行其庭，不见其人，无咎。",
			Short: "",
			Desc:  "《彖》曰：艮，止也。时止则止，时行则行，动静不失其时，其道光明。艮其止，止其所也。上下敌应，不相与也。是以不获其身、行其庭不见其人、无咎也。\n\n《象》曰：兼山，艮。君子以思不出其位。",
			Yao: [6]Yao{
				{
					Image:     YinYao,
					Text:      "初六：艮其趾，无咎。利永贞。",
					ImageText: "《象》曰：艮其趾，未失正也。",
				},
				{
					Image:     YinYao,
					Text:      "六二：艮其腓，不拯其随，其心不快。",
					ImageText: "《象》曰：不拯其随，未退听也。",
				},
				{
					Image:     YangYao,
					Text:      "九三，艮其限，列其夤，厉，熏心。",
					ImageText: "《象》曰：艮其限，危熏心也。",
				},
				{
					Image:     YinYao,
					Text:      "六四：艮其身，无咎。",
					ImageText: "《象》曰：艮其身，止诸躬也。",
				},
				{
					Image:     YinYao,
					Text:      "六五：艮其辅，言有序，悔亡。",
					ImageText: "《象》曰：艮其辅，以中正也。",
				},
				{
					Image:     YangYao,
					Text:      "上九，敦艮，吉。",
					ImageText: "《象》曰：敦艮之吉，以厚终也。",
				},
			},
		},
		53: {
			Index: 53,
			Wai:   BaGuaMap[Xun], // 巽为风
			Nei:   BaGuaMap[Gen], // 艮为山
			Name:  "风山渐",
			DuYin: "",
			Text:  "女归吉，利贞。",
			Short: "",
			Desc:  "《彖》曰：渐之进也，女归吉也。进得位，往有功也。进以正，可以正邦也。其位刚得中也。止而巽，动不穷也。\n\n《象》曰：山上有木，渐。君子以居贤德善俗。",
			Yao: [6]Yao{
				{
					Image:     YinYao,
					Text:      "初六：鸿渐于干。小子厉，有言，无咎。",
					ImageText: "《象》曰：小子之厉，义无咎也。",
				},
				{
					Image:     YinYao,
					Text:      "六二：鸿渐于磐，饮食衎衎，吉。",
					ImageText: "《象》曰：饮食衎衎，不素饱也。",
				},
				{
					Image:     YangYao,
					Text:      "九三：鸿渐于陆。夫征不复，妇孕不育，凶。利御寇。",
					ImageText: "《象》曰：夫征不复，离群丑也。妇孕不育，失其道也。利用御寇，顺相保也。",
				},
				{
					Image:     YinYao,
					Text:      "六四：鸿渐于木，或得其桷，无咎。",
					ImageText: "《象》曰：或得其桷，顺以巽也。",
				},
				{
					Image:     YangYao,
					Text:      "九五：鸿渐于陵，妇三岁不孕，终莫之胜，吉。",
					ImageText: "《象》曰：终莫之胜吉，得所愿也。",
				},
				{
					Image:     YangYao,
					Text:      "上九：鸿渐于陆，其羽可用为仪，吉。",
					ImageText: "《象》曰：其羽可用为仪吉，不可乱也。",
				},
			},
		},
		54: {
			Index: 54,
			Wai:   BaGuaMap[Zhen], // 震为雷
			Nei:   BaGuaMap[Dui],  // 兑为泽
			Name:  "雷泽归妹",
			DuYin: "",
			Text:  "征凶，无攸利。",
			Short: "",
			Desc:  "《彖》曰：归妹，天地之大义也。天地不交而万物不兴。归妹，人之终始也。说以动，所归妹也。征凶”，位不当也。无攸利，柔乘刚也。\n\n《象》曰：泽上有雷，归妹。君子以永终知敝。",
			Yao: [6]Yao{
				{
					Image:     YangYao,
					Text:      "初九：归妹以娣。跛能履，征吉。",
					ImageText: "《象》曰：归妹以娣，以恒也。跛能履吉，相承也。",
				},
				{
					Image:     YangYao,
					Text:      "九二：眇能视，利幽人之贞。",
					ImageText: "《象》曰：“利幽人之贞”，未变常也。",
				},
				{
					Image:     YinYao,
					Text:      "六三，归妹以须，反归以娣。",
					ImageText: "《象》曰：归妹以须，未当也。",
				},
				{
					Image:     YangYao,
					Text:      "九四：归妹愆期，迟归有时。",
					ImageText: "《象》曰：愆期之志，有待而行也。",
				},
				{
					Image:     YinYao,
					Text:      "六五：帝乙归妹，其君之袂不如其娣之袂良。月几望，吉。",
					ImageText: "《象》曰：帝乙归妹，不如其娣之袂良也。其位在中，以贵行也。",
				},
				{
					Image:     YinYao,
					Text:      "上六：女承筐无实，士刲羊无血，无攸利。",
					ImageText: "《象》曰：上六无实，承虚筐也。",
				},
			},
		},
		55: {
			Index: 55,
			Wai:   BaGuaMap[Zhen], // 震为雷
			Nei:   BaGuaMap[Li],   // 离为火
			Name:  "雷火丰",
			DuYin: "",
			Text:  "亨，王假之。勿忧，宜日中。",
			Short: "",
			Desc:  "《彖》曰：丰，大也。明以动，故丰。王假之，尚大也。勿忧宜日中，宜照天下也。日中则昃，月盈则食，天地盈虚，与时消息，而况于人乎，况于鬼神乎。\n\n《象》曰：雷电皆至，丰。君子以折狱致刑。",
			Yao: [6]Yao{
				{
					Image:     YangYao,
					Text:      "初九：遇其配主，虽旬无咎，往有尚。",
					ImageText: "《象》曰：虽旬无咎，过旬灾也。",
				},
				{
					Image:     YinYao,
					Text:      "六二：丰其蔀，日中见斗。往得疑疾，有孚发若，吉。",
					ImageText: "《象》曰：有孚发若，信以发志也。",
				},
				{
					Image:     YangYao,
					Text:      "九三：丰其沛，日中见沫，折其右肱，无咎。",
					ImageText: "《象》曰：丰其沛，不可大事也。折其右肱，终不可用也。",
				},
				{
					Image:     YangYao,
					Text:      "九四：丰其蔀，日中见斗，遇其夷主，吉。",
					ImageText: "《象》曰：丰其蔀，位不当也。日中见斗，幽不明也。遇其夷主，吉行也。",
				},
				{
					Image:     YinYao,
					Text:      "六五：来章有庆誉，吉。",
					ImageText: "《象》曰：六五之吉，有庆也。",
				},
				{
					Image:     YinYao,
					Text:      "上六：丰其屋，蔀其家，窥其户，阒其无人，三岁不觌，凶。",
					ImageText: "《象》曰：丰其屋，天际翔也。窥其户，阒其无人，自藏也。",
				},
			},
		},
		56: {
			Index: 56,
			Wai:   BaGuaMap[Li],  // 离为火
			Nei:   BaGuaMap[Gen], // 艮为山
			Name:  "火山旅",
			DuYin: "",
			Text:  "小亨。旅贞吉。",
			Short: "",
			Desc:  "《彖》曰：旅小亨，柔得中乎外，而顺乎刚，止而丽乎明，是以小亨旅贞吉也。旅之时义大矣哉！\n\n《象》曰：山上有火，旅。君子以明慎用刑而不留狱。",
			Yao: [6]Yao{
				{
					Image:     YinYao,
					Text:      "初六：旅琐琐，斯其所取灾。",
					ImageText: "《象》曰：旅琐琐，志穷灾也。",
				},
				{
					Image:     YinYao,
					Text:      "六二：旅即次，怀其资，得童仆，贞。",
					ImageText: "《象》曰：得童仆贞，终无尤也。",
				},
				{
					Image:     YangYao,
					Text:      "九三：旅焚其次，丧其童仆，贞厉。",
					ImageText: "《象》曰：旅焚其次，亦以伤矣。以旅与下，其义丧也。",
				},
				{
					Image:     YangYao,
					Text:      "九四：旅于处，得其资斧，我心不快。",
					ImageText: "《象》曰：旅于处，未得位也。得其资斧，心未快也。",
				},
				{
					Image:     YinYao,
					Text:      "六五：射雉，一矢亡，终以誉命。",
					ImageText: "《象》曰：终以誉命，上逮也。",
				},
				{
					Image:     YangYao,
					Text:      "上九：鸟焚其巢，旅人先笑后号咷。丧牛于易，凶。",
					ImageText: "《象》曰：以旅在上，其义焚也。丧牛于易，终莫之闻也。",
				},
			},
		},
		57: {
			Index: 57,
			Wai:   BaGuaMap[Xun], // 巽为风
			Nei:   BaGuaMap[Xun], // 巽为风
			Name:  "巽为风",
			DuYin: "",
			Text:  "小亨。利有攸往。利见大人。",
			Short: "",
			Desc:  "《彖》曰：重巽以申命。刚巽乎中正而志行。柔皆顺乎刚，是以小亨、利有攸往、利见大人。\n\n《象》曰：随风，巽。君子以申命行事。",
			Yao: [6]Yao{
				{
					Image:     YinYao,
					Text:      "初六：进退，利武人之贞。",
					ImageText: "《象》曰：进退，志疑也。利武人之贞，志治也。",
				},
				{
					Image:     YangYao,
					Text:      "九二：巽在床下，用史巫纷若，吉，无咎。",
					ImageText: "《象》曰：纷若之吉，得中也。",
				},
				{
					Image:     YangYao,
					Text:      "九三：频巽，吝。",
					ImageText: "《象》曰：频巽之吝，志穷也。",
				},
				{
					Image:     YinYao,
					Text:      "六四：悔亡，田获三品。",
					ImageText: "《象》曰：田获三品，有功也。",
				},
				{
					Image:     YangYao,
					Text:      "九五：贞吉，悔亡，无不利，无初有终。先庚三日，后庚三日，吉。",
					ImageText: "《象》曰：九五之吉，位正中也。",
				},
				{
					Image:     YangYao,
					Text:      "上九：巽在床下，丧其资斧，贞凶。",
					ImageText: "《象》曰：巽在床下，上穷也。丧其资斧，正乎凶也。",
				},
			},
		},
		58: {
			Index: 58,
			Wai:   BaGuaMap[Dui], // 兑为泽
			Nei:   BaGuaMap[Dui], // 兑为泽
			Name:  "兑为泽",
			DuYin: "",
			Text:  "亨。利贞。",
			Short: "",
			Desc:  "《彖》曰：兑，说也。刚中而柔外，说以利贞，是以顺乎天而应乎人。说以先民，民忘其劳。说以犯难，民忘其死。说之大，民劝矣哉！\n\n《象》曰：丽泽，兑。君子以朋友讲习。",
			Yao: [6]Yao{
				{
					Image:     YangYao,
					Text:      "初九：和兑，吉。x",
					ImageText: "《象》曰：和兑之吉，行未疑也。",
				},
				{
					Image:     YangYao,
					Text:      "九二：孚兑，吉，悔亡。",
					ImageText: "《象》曰：孚兑之吉，信志也。",
				},
				{
					Image:     YinYao,
					Text:      "六三：来兑，凶。",
					ImageText: "《象》曰：来兑之凶，位不当也。",
				},
				{
					Image:     YangYao,
					Text:      "九四：商兑未宁，介疾有喜。",
					ImageText: "《象》曰：九四之喜，有庆也。",
				},
				{
					Image:     YangYao,
					Text:      "九五：孚于剥，有厉。",
					ImageText: "《象》曰：孚于剥，位正当也。",
				},
				{
					Image:     YinYao,
					Text:      "上六：引兑。",
					ImageText: "《象》曰：上六引兑，未光也。",
				},
			},
		},
		59: {
			Index: 59,
			Wai:   BaGuaMap[Li],   // 巽为风
			Nei:   BaGuaMap[Zhen], // 坎为水
			Name:  "风水涣",
			DuYin: "",
			Text:  "亨。王假有庙。利涉大川，利贞。",
			Short: "",
			Desc:  "《彖》曰：涣，亨”，刚来而不穷，柔得位乎外而上同。“王假有庙”，王乃在中也。“利涉大川”，乘木有功也。\n\n《象》曰：风行水上，涣。先王以享于帝，立庙。",
			Yao: [6]Yao{
				{
					Image:     YinYao,
					Text:      "初六：用拯马壮，吉。",
					ImageText: "《象》曰：初六之吉顺也。",
				},
				{
					Image:     YangYao,
					Text:      "九二：涣奔其机，悔亡。",
					ImageText: "《象》曰：涣奔其机，得愿也。",
				},
				{
					Image:     YinYao,
					Text:      "六三：涣其躬，无悔。",
					ImageText: "《象》曰：涣其躬，志在外也。",
				},
				{
					Image:     YinYao,
					Text:      "六四：涣其群，元吉。涣有丘，匪夷所思。",
					ImageText: "《象》曰：涣其群元吉，光大也。",
				},
				{
					Image:     YangYao,
					Text:      "九五：涣汗其大号，涣王居，无咎。",
					ImageText: "《象》曰：王居无咎，正位也。",
				},
				{
					Image:     YangYao,
					Text:      "上九：涣其血，去逖出，无咎。",
					ImageText: "《象》曰：涣其血，远害也。",
				},
			},
		},
		60: {
			Index: 60,
			Wai:   BaGuaMap[Kan], // 坎为水
			Nei:   BaGuaMap[Dui], // 兑为泽
			Name:  "水泽节",
			DuYin: "",
			Text:  "亨。苦节，不可贞。",
			Short: "",
			Desc:  "《彖》曰：节亨。刚柔分而刚得中。苦节不可贞，其道穷也。说以行险，当位以节，中正以通。天地节而四时成。节以制度，不伤财，不害民。\n\n《象》曰：泽上有水，节。君子以制数度，议德行。",
			Yao: [6]Yao{
				{
					Image:     YangYao,
					Text:      "初九：不出户庭，无咎。",
					ImageText: "《象》曰：不出户庭，知通塞也。",
				},
				{
					Image:     YangYao,
					Text:      "九二：不出门庭，凶。",
					ImageText: "《象》曰：不出门庭凶，失时极也。",
				},
				{
					Image:     YinYao,
					Text:      "六三：不节若，则嗟若，无咎。",
					ImageText: "《象》曰：不节之嗟，又谁咎也。",
				},
				{
					Image:     YinYao,
					Text:      "六四：安节。亨。",
					ImageText: "《象》曰：安节之亨，承上道也。",
				},
				{
					Image:     YangYao,
					Text:      "九五：甘节，吉，往有尚。",
					ImageText: "《象》曰：甘节之吉，居位中也。",
				},
				{
					Image:     YinYao,
					Text:      "上六：苦节，贞凶，悔亡。",
					ImageText: "《象》曰：苦节贞凶，其道穷也。",
				},
			},
		},
		61: {
			Index: 61,
			Wai:   BaGuaMap[Xun], // 巽为风
			Nei:   BaGuaMap[Dui], // 兑为泽
			Name:  "风泽中孚",
			DuYin: "",
			Text:  "豚鱼，吉。利涉大川，利贞。",
			Short: "",
			Desc:  "《彖》曰：中孚，柔在内而刚得中，说而巽，孚乃化邦也。豚鱼吉，信及豚鱼也。利涉大川，乘木舟虚也。中孚以利贞，乃应乎天也。\n\n《象》曰：泽上有风，中孚。君子以议狱缓死。",
			Yao: [6]Yao{
				{
					Image:     YangYao,
					Text:      "初九：虞吉，有它不燕。",
					ImageText: "《象》曰：初九“虞吉”，志未变也。",
				},
				{
					Image:     YangYao,
					Text:      "九二：鸣鹤在阴，其子和之。我有好爵，吾与尔靡之。",
					ImageText: "《象》曰：其子和之，中心愿也。",
				},
				{
					Image:     YinYao,
					Text:      "六三：得敌，或鼓或罢，或泣或歌。",
					ImageText: "《象》曰：或鼓或罢，位不当也。",
				},
				{
					Image:     YinYao,
					Text:      "六四：月几望，马匹亡，无咎。",
					ImageText: "《象》曰：马匹亡，绝类上也。",
				},
				{
					Image:     YangYao,
					Text:      "九五：有孚挛如，无咎。",
					ImageText: "《象》曰：有孚挛如，位正当也。",
				},
				{
					Image:     YangYao,
					Text:      "上九：翰音登于天，贞凶。",
					ImageText: "《象》曰：翰音登于天，何可长也。",
				},
			},
		},
		62: {
			Index: 62,
			Wai:   BaGuaMap[Zhen], // 震为雷
			Nei:   BaGuaMap[Gen],  // 艮为山
			Name:  "雷山小过",
			DuYin: "",
			Text:  "亨。利贞。可小事，不可大事。飞鸟遗之音，不宜上，宜下，大吉。",
			Short: "",
			Desc:  "《彖》曰：小过，小者过而亨也。过以利贞，与时行也。柔得中，是以小事吉也。刚失位而不中，是以不可大事也。有飞鸟之象焉，飞鸟遗之音、不宜上宜下、大吉，上逆而下顺也。\n\n《象》曰：山上有雷，小过。君子以行过乎恭，丧过乎哀，用过乎俭。",
			Yao: [6]Yao{
				{
					Image:     YinYao,
					Text:      "初六：飞鸟以凶。",
					ImageText: "《象》曰：飞鸟以凶，不可如何也。",
				},
				{
					Image:     YinYao,
					Text:      "六二：过其祖，遇其妣。不及其君，遇其臣。无咎。",
					ImageText: "《象》曰：不及其君，臣不可过也。",
				},
				{
					Image:     YangYao,
					Text:      "九三：弗过防之，从或戕之，凶。",
					ImageText: "《象》曰：从或戕之，凶如何也？",
				},
				{
					Image:     YangYao,
					Text:      "九四：无咎。弗过遇之，往厉必戒，勿用永贞。",
					ImageText: "《象》曰：弗过遇之，位不当也。往厉必戒，终不可长也。",
				},
				{
					Image:     YinYao,
					Text:      "六五：密云不雨，自我西郊。公弋取彼在穴。",
					ImageText: "《象》曰：密云不雨，已上也。",
				},
				{
					Image:     YinYao,
					Text:      "上六：弗遇过之，飞鸟离之，凶，是谓灾眚。",
					ImageText: "《象》曰：弗遇过之，已亢也。",
				},
			},
		},
		63: {
			Index: 63,
			Wai:   BaGuaMap[Kan], // 坎为水
			Nei:   BaGuaMap[Li],  // 离为火
			Name:  "水火既济",
			DuYin: "",
			Text:  "亨小，利贞。初吉终乱。",
			Short: "",
			Desc:  "《彖》曰：既济亨，小者亨也。利贞。刚柔正而位当也。初吉，柔得中也。终止则乱，其道穷也。\n\n《象》曰：水在火上，既济。君子以思患而豫防之。",
			Yao: [6]Yao{
				{
					Image:     YangYao,
					Text:      "初九：曳其轮，濡其尾，无咎。",
					ImageText: "《象》曰：曳其轮，义无咎也。",
				},
				{
					Image:     YinYao,
					Text:      "六二：妇丧其茀，勿逐，七日得。",
					ImageText: "《象》曰：七日得，以中道也。",
				},
				{
					Image:     YangYao,
					Text:      "九三：高宗伐鬼方，三年克之，小人勿用。",
					ImageText: "《象》曰：三年克之，惫也。",
				},
				{
					Image:     YinYao,
					Text:      "六四：繻有衣袽，终日戒。",
					ImageText: "《象》曰：终日戒，有所疑也。",
				},
				{
					Image:     YangYao,
					Text:      "九五：东邻杀牛，不如西邻之禴祭，实受其福。",
					ImageText: "《象》曰：东邻杀牛，不如西邻之时也。实受其福，吉大来也。",
				},
				{
					Image:     YinYao,
					Text:      "上六：濡其首，厉。",
					ImageText: "《象》曰：濡其首厉，何可久也。",
				},
			},
		},
		64: {
			Index: 64,
			Wai:   BaGuaMap[Li],  // 离为火
			Nei:   BaGuaMap[Kan], // 坎为水
			Name:  "火水未济",
			DuYin: "",
			Text:  "亨。小狐汔济，濡其尾，无攸利。",
			Short: "",
			Desc:  "《彖》曰：未济亨，柔得中也。小狐汔济，未出中也。濡其尾，无攸利，不续终也。虽不当位，刚柔应也。\n\n《象》曰：火在水上，未济。君子以慎辨物居方。",
			Yao: [6]Yao{
				{
					Image:     YinYao,
					Text:      "初六：濡其尾，吝。",
					ImageText: "《象》曰：濡其尾，亦不知极也。",
				},
				{
					Image:     YangYao,
					Text:      "九二：曳其轮，贞吉。",
					ImageText: "《象》曰：九二贞吉，中以行正也。",
				},
				{
					Image:     YinYao,
					Text:      "六三：未济，征凶。利涉大川。",
					ImageText: "《象》曰：未济征凶，位不当也。",
				},
				{
					Image:     YangYao,
					Text:      "九四：贞吉，悔亡。震用伐鬼方，三年，有赏于大国。",
					ImageText: "《象》曰：贞吉悔亡，志行也。",
				},
				{
					Image:     YinYao,
					Text:      "六五：贞吉，无悔。君子之光，有孚吉。",
					ImageText: "《象》曰：君子之光，其辉吉也。",
				},
				{
					Image:     YangYao,
					Text:      "上九：有孚于饮酒，无咎。濡其首，有孚失是。",
					ImageText: "《象》曰：饮酒濡首，亦不知节也。",
				},
			},
		},
	}
)

func main() {
	rand.Seed(time.Now().UnixNano())

	arr := []int{}
	for i := 0; i < 6; i++ {
		tempYinYang := rand.Intn(1)
		arr = append(arr, tempYinYang)
	}

	fmt.Println("《易经》第六卦 讼 天水讼 乾上坎下\n　　讼卦，象征争论、诉讼。乾为天，坎为水，天西转与水东流背向而行，像人与人不和而争辩。讼象征争辩争论，含诉讼之义。当不易和解时，便会导致诉讼。应该找有大德大才的人进行决断，不要逞强冒险。\n　　《讼》：有孚窒惕，中吉，终凶。利见大人。不利涉大川。\n　　《彖》曰：讼，上刚下险，险而健，讼。“讼有孚窒惕，中吉”，刚来而得中也。“终凶”，讼不可成也。“利见大人”，尚中正也。“不利涉大川”，入于渊也。\n　　《象》曰：天与水违行，讼。君子以作事谋始。\n　　初六，不永所事，小有言，终吉。\n　　《象》曰：“不永所事”，讼不可长也。虽“小有言”，其辩明也。\n　　九二，不克讼，归而逋。其邑人三百户，无眚。\n　　《象》曰：“不克讼”，归逋窜也。自下讼上，患至掇也。\n　　六三，食旧德，贞厉，终吉。或従王事，无成。\n　　《象》曰：食旧德，従上吉也。\n　　九四，不克讼，复既命渝。安贞吉。\n　　《象》曰：复即命渝，安贞不失也。\n　　九五：讼，元吉。\n　　《象》曰：“讼，元吉”以中正也。\n　　上九：或锡之鞶带，终朝三褫之。\n　　《象》曰：以讼受服，亦不足敬也。")
}
