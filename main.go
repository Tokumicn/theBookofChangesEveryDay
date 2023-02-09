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
			Wai:   BaGuaMap[Kun],
			Nei:   BaGuaMap[Kan],
			Name:  "",
			DuYin: "",
			Text:  "",
			Short: "",
			Desc:  "",
			Yao: [6]Yao{
				{
					Image:     YinYao,
					Text:      "",
					ImageText: "",
				},
				{
					Image:     YangYao,
					Text:      "",
					ImageText: "",
				},
				{
					Image:     YinYao,
					Text:      "",
					ImageText: "",
				},
				{
					Image:     YangYao,
					Text:      "",
					ImageText: "",
				},
				{
					Image:     YangYao,
					Text:      "",
					ImageText: "",
				},
				{
					Image:     YangYao,
					Text:      "",
					ImageText: "",
				},
			},
		},
		17: {
			Index: 17,
			Wai:   BaGuaMap[Kun],
			Nei:   BaGuaMap[Kan],
			Name:  "",
			DuYin: "",
			Text:  "",
			Short: "",
			Desc:  "",
			Yao: [6]Yao{
				{
					Image:     YinYao,
					Text:      "",
					ImageText: "",
				},
				{
					Image:     YangYao,
					Text:      "",
					ImageText: "",
				},
				{
					Image:     YinYao,
					Text:      "",
					ImageText: "",
				},
				{
					Image:     YangYao,
					Text:      "",
					ImageText: "",
				},
				{
					Image:     YangYao,
					Text:      "",
					ImageText: "",
				},
				{
					Image:     YangYao,
					Text:      "",
					ImageText: "",
				},
			},
		},
		18: {
			Index: 18,
			Wai:   BaGuaMap[Kun],
			Nei:   BaGuaMap[Kan],
			Name:  "",
			DuYin: "",
			Text:  "",
			Short: "",
			Desc:  "",
			Yao: [6]Yao{
				{
					Image:     YinYao,
					Text:      "",
					ImageText: "",
				},
				{
					Image:     YangYao,
					Text:      "",
					ImageText: "",
				},
				{
					Image:     YinYao,
					Text:      "",
					ImageText: "",
				},
				{
					Image:     YangYao,
					Text:      "",
					ImageText: "",
				},
				{
					Image:     YangYao,
					Text:      "",
					ImageText: "",
				},
				{
					Image:     YangYao,
					Text:      "",
					ImageText: "",
				},
			},
		},
		19: {
			Index: 19,
			Wai:   BaGuaMap[Kun],
			Nei:   BaGuaMap[Kan],
			Name:  "",
			DuYin: "",
			Text:  "",
			Short: "",
			Desc:  "",
			Yao: [6]Yao{
				{
					Image:     YinYao,
					Text:      "",
					ImageText: "",
				},
				{
					Image:     YangYao,
					Text:      "",
					ImageText: "",
				},
				{
					Image:     YinYao,
					Text:      "",
					ImageText: "",
				},
				{
					Image:     YangYao,
					Text:      "",
					ImageText: "",
				},
				{
					Image:     YangYao,
					Text:      "",
					ImageText: "",
				},
				{
					Image:     YangYao,
					Text:      "",
					ImageText: "",
				},
			},
		},
		20: {
			Index: 20,
			Wai:   BaGuaMap[Kun],
			Nei:   BaGuaMap[Kan],
			Name:  "",
			DuYin: "",
			Text:  "",
			Short: "",
			Desc:  "",
			Yao: [6]Yao{
				{
					Image:     YinYao,
					Text:      "",
					ImageText: "",
				},
				{
					Image:     YangYao,
					Text:      "",
					ImageText: "",
				},
				{
					Image:     YinYao,
					Text:      "",
					ImageText: "",
				},
				{
					Image:     YangYao,
					Text:      "",
					ImageText: "",
				},
				{
					Image:     YangYao,
					Text:      "",
					ImageText: "",
				},
				{
					Image:     YangYao,
					Text:      "",
					ImageText: "",
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
