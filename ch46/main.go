package main

import (
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
	"strings"
	"xorm.io/xorm"
)

const JSONString = `{"logistics_companies_get_response":{"logistics_companies":[{"code":"STO","available":1,"logistics_company":"申通快递","id":1},{"code":"SHHT","available":0,"logistics_company":"上海汇通","id":2},{"code":"HT","available":1,"logistics_company":"百世快递","id":3},{"code":"SF","available":1,"logistics_company":"顺丰快递","id":44},{"code":"YTO","available":1,"logistics_company":"圆通快递","id":85},{"code":"NBYGZT","available":0,"logistics_company":"内部员工自提","id":86},{"code":"BBSD","available":0,"logistics_company":"奔奔速达","id":88},{"code":"SAD","available":0,"logistics_company":"赛澳递","id":89},{"code":"CHENGBANG","available":0,"logistics_company":"晟邦物流","id":90},{"code":"ZTO","available":1,"logistics_company":"中通快递","id":115},{"code":"QF","available":0,"logistics_company":"全峰快递","id":116},{"code":"YS","available":1,"logistics_company":"优速","id":117},{"code":"EMS","available":1,"logistics_company":"邮政EMS","id":118},{"code":"TT","available":1,"logistics_company":"天天快递","id":119},{"code":"JD","available":1,"logistics_company":"京东配送","id":120},{"code":"YUNDA","available":1,"logistics_company":"韵达快递","id":121},{"code":"KJ","available":0,"logistics_company":"快捷快递","id":122},{"code":"GTO","available":0,"logistics_company":"国通快递","id":124},{"code":"DDCBPS","available":0,"logistics_company":"当当出版配送","id":128},{"code":"ZJS","available":1,"logistics_company":"宅急送快递","id":129},{"code":"RFD","available":0,"logistics_company":"如风达","id":130},{"code":"DB","available":1,"logistics_company":"德邦快递","id":131},{"code":"YZXB","available":1,"logistics_company":"邮政快递包裹","id":132},{"code":"LBEX","available":0,"logistics_company":"龙邦快递","id":133},{"code":"FEDEX","available":1,"logistics_company":"联邦快递","id":135},{"code":"JIUYE","available":1,"logistics_company":"九曳供应链","id":136},{"code":"BCDRD","available":0,"logistics_company":"百城当日达快递","id":137},{"code":"DDKD","available":0,"logistics_company":"达达快递","id":138},{"code":"DDJWL","available":0,"logistics_company":"冻到家物流","id":139},{"code":"NJCHENGBANG","available":0,"logistics_company":"南京晟邦","id":140},{"code":"SXHONGMAJIA","available":0,"logistics_company":"山西红马甲","id":141},{"code":"WXWL","available":0,"logistics_company":"万象物流","id":142},{"code":"LIJISONG","available":0,"logistics_company":"立即送","id":143},{"code":"MENDUIMEN","available":0,"logistics_company":"门对门","id":144},{"code":"SAD2","available":0,"logistics_company":"赛澳递","id":145},{"code":"FENGCHENG","available":0,"logistics_company":"丰程","id":147},{"code":"ADX","available":0,"logistics_company":"安达信","id":148},{"code":"HWKD","available":0,"logistics_company":"海外快递","id":149},{"code":"GZLT","available":0,"logistics_company":"飞远物流","id":150},{"code":"NDKD","available":0,"logistics_company":"南都快递","id":151},{"code":"HUIWEN","available":0,"logistics_company":"汇文快递","id":152},{"code":"NDKD","available":0,"logistics_company":"南都快递","id":153},{"code":"HUANGMAJIA","available":0,"logistics_company":"黄马甲","id":154},{"code":"SURE","available":1,"logistics_company":"速尔快递","id":155},{"code":"YAMAXUNWULIU","available":1,"logistics_company":"亚马逊物流","id":156},{"code":"YCT","available":1,"logistics_company":"黑猫宅急便","id":157},{"code":"SFHY","available":0,"logistics_company":"顺丰航运","id":158},{"code":"YTHY","available":0,"logistics_company":"圆通航运","id":159},{"code":"PINHAOHUO","available":0,"logistics_company":"拼好货","id":160},{"code":"SHSAD","available":0,"logistics_company":"上海赛澳递","id":161},{"code":"BJCS","available":0,"logistics_company":"城市100","id":162},{"code":"ZMKM","available":0,"logistics_company":"芝麻开门","id":163},{"code":"SHUNJIEFENGDA","available":0,"logistics_company":"顺捷丰达","id":164},{"code":"HTXMJ","available":0,"logistics_company":"汇通小红马","id":165},{"code":"LNXHM","available":0,"logistics_company":"辽宁小红马","id":166},{"code":"LNHUANGMAJIA","available":0,"logistics_company":"辽宁黄马甲","id":167},{"code":"JSSAD","available":0,"logistics_company":"江苏赛澳递","id":168},{"code":"SANRENXING","available":0,"logistics_company":"三人行","id":169},{"code":"THJD","available":0,"logistics_company":"通和佳递","id":170},{"code":"SUJIEVIP","available":0,"logistics_company":"速捷","id":171},{"code":"XNXD","available":0,"logistics_company":"信诺迅达","id":172},{"code":"FXIANSHENG","available":0,"logistics_company":"风先生","id":173},{"code":"KUANRONG","available":0,"logistics_company":"宽容","id":174},{"code":"GZTK","available":0,"logistics_company":"广州途客","id":175},{"code":"XIAOHONGMAO","available":0,"logistics_company":"小红帽","id":176},{"code":"PENGDA","available":0,"logistics_company":"鹏达","id":177},{"code":"FJGZLT","available":0,"logistics_company":"福建飞远","id":178},{"code":"ETEKUAI","available":0,"logistics_company":"E特快","id":179},{"code":"SELF","available":1,"logistics_company":"其他","id":180},{"code":"YUNNIAO","available":0,"logistics_company":"云鸟","id":181},{"code":"BAODA","available":0,"logistics_company":"保达","id":182},{"code":"KYE","available":1,"logistics_company":"跨越速运","id":183},{"code":"JLHMJ","available":0,"logistics_company":"吉林黄马甲","id":184},{"code":"CHENGJI","available":0,"logistics_company":"城际速递","id":185},{"code":"USPS","available":1,"logistics_company":"usps","id":186},{"code":"ANJELEX","available":0,"logistics_company":"青岛安捷","id":187},{"code":"DHTY","available":0,"logistics_company":"大韩通运","id":188},{"code":"BANGBANGTANG","available":0,"logistics_company":"棒棒糖","id":189},{"code":"TUXIAN","available":0,"logistics_company":"途鲜","id":190},{"code":"CNKD","available":0,"logistics_company":"菜鸟快递","id":191},{"code":"EMSKD","available":0,"logistics_company":"EMS经济快递","id":192},{"code":"HZZX","available":0,"logistics_company":"汇站众享","id":193},{"code":"PAIKE","available":0,"logistics_company":"派客","id":194},{"code":"XLOBO","available":1,"logistics_company":"贝海国际速递","id":195},{"code":"TFGJKD","available":0,"logistics_company":"丰泰国际快递","id":196},{"code":"HUANQIU","available":1,"logistics_company":"环球速运","id":197},{"code":"SHUNFASUDI","available":0,"logistics_company":"168顺发速递","id":198},{"code":"QQSD","available":0,"logistics_company":"全球快递","id":199},{"code":"CG","available":1,"logistics_company":"程光物流","id":200},{"code":"UAPEX","available":0,"logistics_company":"全一快递","id":201},{"code":"HQSY","available":0,"logistics_company":"环球速运","id":202},{"code":"DJKJ","available":0,"logistics_company":"东骏快捷","id":203},{"code":"BSKD","available":0,"logistics_company":"百世快递","id":204},{"code":"YCGWL","available":1,"logistics_company":"远成快运","id":205},{"code":"FTGJSD","available":0,"logistics_company":"风腾国际速递","id":206},{"code":"BNZY","available":0,"logistics_company":"笨鸟转运","id":207},{"code":"ANNENG","available":0,"logistics_company":"安能快递","id":208},{"code":"EPS","available":0,"logistics_company":"联众国际快运","id":209},{"code":"HOAU","available":1,"logistics_company":"天地华宇","id":210},{"code":"ZHONGYOUWULIU","available":1,"logistics_company":"中邮速递","id":211},{"code":"HITAOYI","available":0,"logistics_company":"hi淘易","id":212},{"code":"INTEREMS","available":1,"logistics_company":"EMS-国际件","id":213},{"code":"ZTKY","available":1,"logistics_company":"中铁物流","id":214},{"code":"CYWL","available":0,"logistics_company":"楚源物流","id":215},{"code":"XBWL","available":1,"logistics_company":"新邦物流","id":216},{"code":"FLASH","available":0,"logistics_company":"Flash Express","id":217},{"code":"NSF","available":1,"logistics_company":"新顺丰NSF","id":218},{"code":"RLKD","available":0,"logistics_company":"锐朗快递","id":219},{"code":"WDGJWL","available":0,"logistics_company":"王道国际物流","id":220},{"code":"DCS","available":0,"logistics_company":"DCS GLOBAL","id":221},{"code":"XSKD","available":0,"logistics_company":"迅速快递","id":222},{"code":"FTD","available":1,"logistics_company":"富腾达","id":223},{"code":"QFWL","available":0,"logistics_company":"琦峰物流","id":224},{"code":"JYTWL","available":0,"logistics_company":"金运通物流","id":225},{"code":"EWE","available":1,"logistics_company":"EWE全球快递","id":226},{"code":"RRS","available":1,"logistics_company":"日日顺物流","id":227},{"code":"SNWL","available":1,"logistics_company":"苏宁快递","id":228},{"code":"BESTQJT","available":1,"logistics_company":"百世快运","id":229},{"code":"DEBANGWULIU","available":1,"logistics_company":"德邦物流","id":230},{"code":"WEITEPAI","available":0,"logistics_company":"微特派","id":231},{"code":"MYAAE","available":1,"logistics_company":"AAE全球专递","id":232},{"code":"ARAMEX","available":1,"logistics_company":"Aramex","id":233},{"code":"ASENDIA","available":0,"logistics_company":"Asendia USA","id":234},{"code":"CITYLINK","available":0,"logistics_company":"City-Link","id":235},{"code":"COE","available":1,"logistics_company":"COE东方快递","id":236},{"code":"DHLDE","available":0,"logistics_company":"DHL德国","id":237},{"code":"DHL","available":0,"logistics_company":"DHL全球","id":238},{"code":"DHLCN","available":1,"logistics_company":"DHL中国","id":239},{"code":"EMSGJ","available":0,"logistics_company":"EMS国际","id":240},{"code":"FEDEXUS","available":0,"logistics_company":"FedEx美国","id":241},{"code":"FEDEXCN","available":0,"logistics_company":"FedEx中国","id":242},{"code":"OCS","available":0,"logistics_company":"OCS","id":243},{"code":"ONTRAC","available":0,"logistics_company":"OnTrac","id":244},{"code":"TNT","available":0,"logistics_company":"TNT","id":245},{"code":"UPS","available":1,"logistics_company":"UPS","id":246},{"code":"POSTAL","available":0,"logistics_company":"阿尔巴尼亚邮政","id":247},{"code":"POSTAR","available":0,"logistics_company":"阿根廷邮政","id":248},{"code":"POSTAE","available":0,"logistics_company":"阿联酋邮政","id":249},{"code":"POSTEE","available":0,"logistics_company":"爱沙尼亚邮政","id":250},{"code":"POSTAT","available":0,"logistics_company":"奥地利邮政","id":252},{"code":"POSTAU","available":0,"logistics_company":"澳大利亚邮政","id":253},{"code":"POSTPK","available":0,"logistics_company":"巴基斯坦邮政","id":254},{"code":"POSTBR","available":0,"logistics_company":"巴西邮政","id":255},{"code":"POSTBY","available":0,"logistics_company":"白俄罗斯邮政","id":256},{"code":"EES","available":0,"logistics_company":"百福东方","id":257},{"code":"POSTB","available":0,"logistics_company":"包裹信件","id":258},{"code":"POSTBG","available":0,"logistics_company":"保加利亚邮政","id":259},{"code":"BLSYZ","available":0,"logistics_company":"比利时邮政","id":260},{"code":"BLYZ","available":0,"logistics_company":"波兰邮政","id":261},{"code":"CXCOD","available":1,"logistics_company":"传喜物流","id":262},{"code":"DTW","available":1,"logistics_company":"大田物流","id":263},{"code":"4PX","available":1,"logistics_company":"递四方","id":264},{"code":"RUSTON","available":0,"logistics_company":"俄速通","id":265},{"code":"FGYZ","available":0,"logistics_company":"法国邮政","id":266},{"code":"GZFY","available":0,"logistics_company":"凡宇快递","id":267},{"code":"ZTKY1","available":1,"logistics_company":"飞豹快递","id":268},{"code":"HZABC","available":0,"logistics_company":"飞远(爱彼西)配送","id":269},{"code":"POSTFI","available":0,"logistics_company":"芬兰邮政","id":270},{"code":"POSTCO","available":0,"logistics_company":"哥伦比亚邮政","id":271},{"code":"EPOST","available":0,"logistics_company":"韩国邮政","id":272},{"code":"HLWL","available":1,"logistics_company":"恒路物流","id":273},{"code":"HQKY","available":0,"logistics_company":"华企快运","id":274},{"code":"TMS56","available":1,"logistics_company":"加运美","id":275},{"code":"CNEX","available":1,"logistics_company":"佳吉快运","id":276},{"code":"JIAYI","available":1,"logistics_company":"佳怡物流","id":277},{"code":"KERRYEAS","available":0,"logistics_company":"嘉里大通","id":278},{"code":"JKYZ","available":0,"logistics_company":"捷克邮政","id":279},{"code":"JDYWL","available":0,"logistics_company":"筋斗云物流","id":280},{"code":"SZKKE","available":1,"logistics_company":"京广速递","id":281},{"code":"POSTHR","available":0,"logistics_company":"克罗地亚邮政","id":282},{"code":"POSTLV","available":0,"logistics_company":"拉脱维亚邮政","id":283},{"code":"POSTLB","available":0,"logistics_company":"黎巴嫩邮政","id":284},{"code":"LTS","available":1,"logistics_company":"联昊通","id":285},{"code":"POSTMT","available":0,"logistics_company":"马耳他邮政","id":286},{"code":"POSTMK","available":0,"logistics_company":"马其顿邮政","id":287},{"code":"POSTMU","available":0,"logistics_company":"毛里求斯邮政","id":288},{"code":"SERPOST","available":0,"logistics_company":"秘鲁邮政","id":289},{"code":"MBEX","available":0,"logistics_company":"民邦快递","id":290},{"code":"CAE","available":1,"logistics_company":"民航快递","id":291},{"code":"SZML56","available":0,"logistics_company":"明亮物流","id":292},{"code":"POSTMD","available":0,"logistics_company":"摩尔多瓦邮政","id":293},{"code":"POSTZA","available":0,"logistics_company":"南非邮政","id":294},{"code":"POSTNO","available":0,"logistics_company":"挪威邮政","id":295},{"code":"POSTPT","available":0,"logistics_company":"葡萄牙邮政","id":296},{"code":"QRT","available":0,"logistics_company":"全日通","id":297},{"code":"RBYZEMS","available":0,"logistics_company":"日本邮政","id":298},{"code":"POSTSE","available":0,"logistics_company":"瑞典邮政","id":299},{"code":"POSTCH","available":0,"logistics_company":"瑞士邮政","id":300},{"code":"POSTSRB","available":0,"logistics_company":"塞尔维亚邮政","id":301},{"code":"SANTAI","available":0,"logistics_company":"三态速递","id":302},{"code":"POSTSA","available":0,"logistics_company":"沙特邮政","id":303},{"code":"SZSA56","available":0,"logistics_company":"圣安物流","id":304},{"code":"FJSFWLJTYXGS","available":1,"logistics_company":"盛丰物流","id":305},{"code":"SHENGHUI","available":1,"logistics_company":"盛辉物流","id":306},{"code":"POSTSK","available":0,"logistics_company":"斯洛伐克邮政","id":307},{"code":"POSTSI","available":0,"logistics_company":"斯洛文尼亚邮政","id":308},{"code":"SUIJIAWL","available":0,"logistics_company":"穗佳物流","id":309},{"code":"POSTTH","available":0,"logistics_company":"泰国邮政","id":310},{"code":"POSTTR","available":0,"logistics_company":"土耳其邮政","id":311},{"code":"MANCOWL","available":1,"logistics_company":"万家物流","id":312},{"code":"POSTUA","available":0,"logistics_company":"乌克兰邮政","id":313},{"code":"POSTES","available":0,"logistics_company":"西班牙邮政","id":314},{"code":"XFWL","available":1,"logistics_company":"信丰物流","id":315},{"code":"POSTHU","available":0,"logistics_company":"匈牙利邮政","id":316},{"code":"AIR","available":1,"logistics_company":"亚风速递","id":317},{"code":"POSTAM","available":0,"logistics_company":"亚美尼亚邮政","id":318},{"code":"YWWL","available":1,"logistics_company":"燕文物流","id":319},{"code":"POSTIT","available":0,"logistics_company":"意大利邮政","id":320},{"code":"FEC","available":0,"logistics_company":"银捷速递","id":321},{"code":"POSTIN","available":0,"logistics_company":"印度邮政","id":322},{"code":"ROYALMAIL","available":0,"logistics_company":"英国皇家邮政","id":323},{"code":"POSTBBZ","available":1,"logistics_company":"邮政标准快递","id":324},{"code":"CNPOSTGJ","available":1,"logistics_company":"邮政国际包裹","id":325},{"code":"YFEXPRESS","available":0,"logistics_company":"越丰物流","id":326},{"code":"YTZG","available":0,"logistics_company":"运通中港快递","id":327},{"code":"ZENY","available":0,"logistics_company":"增益速递","id":328},{"code":"POSTCL","available":0,"logistics_company":"智利邮政","id":329},{"code":"SPSR","available":0,"logistics_company":"中俄快递","id":330},{"code":"CRE","available":1,"logistics_company":"中铁快运","id":332},{"code":"KFW","available":0,"logistics_company":"快服务快递","id":333},{"code":"KDN","available":0,"logistics_company":"快递鸟","id":334},{"code":"YOUBANG","available":1,"logistics_company":"优邦国际速运","id":335},{"code":"TJ","available":1,"logistics_company":"天际快递","id":336},{"code":"FY","available":1,"logistics_company":"飞洋快递","id":337},{"code":"BM","available":1,"logistics_company":"斑马物联网","id":338},{"code":"EKM","available":1,"logistics_company":"易客满","id":339},{"code":"JDKD","available":1,"logistics_company":"京东大件物流","id":340},{"code":"SUBIDA","available":1,"logistics_company":"速必达","id":341},{"code":"DJKJWL","available":0,"logistics_company":"东骏快捷","id":342},{"code":"ZTOKY","available":1,"logistics_company":"中通快运","id":343},{"code":"YDKY","available":1,"logistics_company":"韵达快运","id":344},{"code":"ANKY","available":1,"logistics_company":"安能快运","id":345},{"code":"ANDE","available":1,"logistics_company":"安得物流","id":346},{"code":"WM","available":1,"logistics_company":"中粮我买网","id":347},{"code":"YMDD","available":1,"logistics_company":"壹米滴答","id":348},{"code":"DD","available":1,"logistics_company":"当当网","id":349},{"code":"PJ","available":0,"logistics_company":"品骏","id":350},{"code":"OTP","available":0,"logistics_company":"承诺达特快","id":351},{"code":"AXWL","available":1,"logistics_company":"安迅物流","id":352},{"code":"YJ","available":0,"logistics_company":"友家速递","id":353},{"code":"SDSD","available":1,"logistics_company":"D速物流","id":354},{"code":"STOINTER","available":1,"logistics_company":"申通国际","id":355},{"code":"YZT","available":1,"logistics_company":"一智通","id":356},{"code":"JGSD","available":0,"logistics_company":"京广速递","id":357},{"code":"SXJD","available":1,"logistics_company":"顺心捷达","id":358},{"code":"QH","available":1,"logistics_company":"群航国际货运","id":359},{"code":"ZWYSD","available":1,"logistics_company":"中外运速递","id":360},{"code":"ZZSY","available":1,"logistics_company":"卓志速运","id":361},{"code":"JZMSD","available":1,"logistics_company":"加州猫速递","id":362},{"code":"GJ","available":1,"logistics_company":"高捷物流","id":363},{"code":"SQWL","available":1,"logistics_company":"商桥物流","id":364},{"code":"FR","available":1,"logistics_company":"复融供应链","id":365},{"code":"ZY","available":1,"logistics_company":"中远e环球","id":366},{"code":"YDGJ","available":1,"logistics_company":"韵达国际","id":367},{"code":"MKGJ","available":1,"logistics_company":"美快国际","id":368},{"code":"NFCM","available":0,"logistics_company":"南方传媒","id":369},{"code":"WSPY","available":1,"logistics_company":"威时沛运","id":370},{"code":"ZTOINTER","available":1,"logistics_company":"中通国际","id":371},{"code":"SFKY","available":1,"logistics_company":"顺丰快运","id":372},{"code":"MGWL","available":1,"logistics_company":"亚马逊综合物流","id":373},{"code":"HKE","available":1,"logistics_company":"HKE国际速递","id":374},{"code":"EFSPOST","available":1,"logistics_company":"新西兰平安物流","id":375},{"code":"HTINTER","available":1,"logistics_company":"百世国际","id":376},{"code":"BSE","available":0,"logistics_company":"蓝天国际快递","id":377},{"code":"YLJY","available":0,"logistics_company":"优联吉运","id":378},{"code":"ZYSFWL","available":1,"logistics_company":"转运四方物流","id":379},{"code":"WSKD","available":1,"logistics_company":"威盛快递","id":380},{"code":"YTGJ","available":1,"logistics_company":"圆通国际","id":381},{"code":"HXWL","available":1,"logistics_company":"海信物流","id":382},{"code":"HYWL","available":1,"logistics_company":"空港宏远电商物流","id":383},{"code":"JTSD","available":1,"logistics_company":"极兔速递","id":384},{"code":"UCS","available":1,"logistics_company":"合众速递","id":385},{"code":"SYNSHIP","available":1,"logistics_company":"SYNSHIP快递","id":386},{"code":"21CAKE","available":1,"logistics_company":"21cake物流","id":387},{"code":"WHHDJ","available":1,"logistics_company":"娃哈哈到家配送","id":388},{"code":"FZGJ","available":1,"logistics_company":"方舟国际速递","id":389},{"code":"STZNWL","available":1,"logistics_company":"圣塔智能物流","id":390},{"code":"SFGJ","available":1,"logistics_company":"顺丰国际","id":391},{"code":"TJWL","available":1,"logistics_company":"泰进物流","id":392},{"code":"QJW","available":1,"logistics_company":"千机网1小时达","id":393},{"code":"AuExpress","available":1,"logistics_company":"澳邮中国快运","id":394},{"code":"HSSY","available":1,"logistics_company":"汇森速运","id":395},{"code":"XFXBWL","available":1,"logistics_company":"幸福西饼物流","id":396},{"code":"DDKS","available":1,"logistics_company":"叮当快送","id":397},{"code":"QYT","available":1,"logistics_company":"泉源堂","id":398},{"code":"YFDYF","available":1,"logistics_company":"益丰大药房","id":399},{"code":"YTDKD","available":0,"logistics_company":"易达通快递","id":400},{"code":"ZJWL","available":1,"logistics_company":"中汲物流","id":401},{"code":"SHUANGHUI","available":0,"logistics_company":"双汇","id":402},{"code":"FENGWANG","available":1,"logistics_company":"丰网速运","id":403},{"code":"WUXINYAOFANG","available":0,"logistics_company":"五心药房","id":404},{"code":"LAOBAIXING","available":1,"logistics_company":"老百姓大药房","id":405},{"code":"YUELUWULIU","available":1,"logistics_company":"跃陆物流","id":406},{"code":"NFSQ","available":1,"logistics_company":"农夫山泉","id":407},{"code":"YIJIUYIJIU","available":1,"logistics_company":"1919酒类直供","id":408},{"code":"QUANYOUJIAJU","available":1,"logistics_company":"全友家居","id":409},{"code":"GUJIAJIAJU","available":0,"logistics_company":"顾家家居","id":410},{"code":"ZHIHUASHI","available":1,"logistics_company":"芝华仕","id":411},{"code":"TIANMAWULIU","available":0,"logistics_company":"天马物流","id":412},{"code":"HALUO","available":1,"logistics_company":"哈啰出行","id":413},{"code":"LINSHIMUYE","available":1,"logistics_company":"林氏物流","id":414},{"code":"SIJIAPP","available":1,"logistics_company":"商家自行配送","id":415},{"code":"SFTC","available":1,"logistics_company":"顺丰同城","id":416},{"code":"SHANSONG","available":1,"logistics_company":"闪送","id":417},{"code":"SFINTL","available":1,"logistics_company":"顺丰集运","id":418},{"code":"SZJY","available":1,"logistics_company":"神州集运","id":419},{"code":"DSDCD","available":1,"logistics_company":"代收点仓端","id":420},{"code":"PADTF","available":1,"logistics_company":"平安达腾飞","id":421},{"code":"SUTENG","available":1,"logistics_company":"速腾快递","id":422},{"code":"YXWL","available":1,"logistics_company":"宇鑫物流","id":423},{"code":"SUTWL","available":1,"logistics_company":"速通物流","id":424},{"code":"JUYUWL","available":1,"logistics_company":"具语平台物流","id":425},{"code":"AJJPWL","available":1,"logistics_company":"安居佳配物流","id":426},{"code":"SANZHIWL","available":1,"logistics_company":"三志物流","id":427},{"code":"FUHUWL","available":1,"logistics_company":"福虎物流","id":428},{"code":"YADSY","available":1,"logistics_company":"源安达速运","id":429},{"code":"LIANYUNHUI","available":1,"logistics_company":"联运汇","id":430},{"code":"YONGCHANGWL","available":1,"logistics_company":"永昌物流","id":431},{"code":"HMJT","available":1,"logistics_company":"澳全球","id":432},{"code":"NKHXWL","available":1,"logistics_company":"南康洪鑫物流","id":433},{"code":"JGWL","available":1,"logistics_company":"景光物流","id":434},{"code":"YUNJUWL","available":1,"logistics_company":"云聚物流","id":435},{"code":"YIXINSUYUN","available":1,"logistics_company":"益鑫速运","id":436},{"code":"KJKWL","available":1,"logistics_company":"快捷快物流","id":437},{"code":"LHTSD","available":0,"logistics_company":"联昊通速递","id":438},{"code":"GZLXWL","available":1,"logistics_company":"赣州龙鑫物流","id":439},{"code":"HNHTWL","available":1,"logistics_company":"河南鸿泰物流","id":440},{"code":"YANGBAOGUO","available":1,"logistics_company":"洋包裹","id":441},{"code":"SDXDWL","available":1,"logistics_company":"山东湘达物流","id":442},{"code":"GZCBWL","available":1,"logistics_company":"赣州楚邦物流","id":443},{"code":"YITUWULIU","available":1,"logistics_company":"易途物流","id":444},{"code":"YOUYUANWL","available":1,"logistics_company":"友源物流","id":445},{"code":"JITUJIYUN","available":1,"logistics_company":"极兔集运","id":446},{"code":"YLBDT","available":1,"logistics_company":"永利八达通","id":447},{"code":"ZTOCC","available":1,"logistics_company":"中通冷链","id":448},{"code":"CHINAUNICOM","available":1,"logistics_company":"联通自有物流","id":449},{"code":"YANWENJIYUN","available":1,"logistics_company":"燕文集运","id":450},{"code":"XLHJY","available":1,"logistics_company":"鑫隆华集运","id":451},{"code":"AMLJY","available":1,"logistics_company":"艾姆勒集运","id":452},{"code":"GUYIWULIU","available":1,"logistics_company":"古宜物流","id":453},{"code":"LUOTUOWULIU","available":1,"logistics_company":"骆驼物流","id":454},{"code":"XIAOMIWULIU","available":1,"logistics_company":"小米物流","id":455},{"code":"YSJE","available":1,"logistics_company":"139 express","id":456},{"code":"YBGJ","available":1,"logistics_company":"Jańa Post","id":457},{"code":"YZDSBK","available":1,"logistics_company":"邮政电商标快","id":458},{"code":"JIEZHOU","available":1,"logistics_company":"芥舟物流","id":459},{"code":"JE","available":1,"logistics_company":"Jingle Express","id":460},{"code":"WJXSD","available":1,"logistics_company":"微集新世代","id":461},{"code":"JYSJ","available":1,"logistics_company":"集运世家","id":462},{"code":"BUYUP","available":1,"logistics_company":"BUYUP","id":463},{"code":"AMICI","available":1,"logistics_company":"欧米奇速递","id":464},{"code":"ZYWL","available":1,"logistics_company":"中邮物流","id":465},{"code":"ECMS","available":1,"logistics_company":"易客满物流","id":466},{"code":"TST","available":1,"logistics_company":"TST速运通","id":467},{"code":"JPRSJY","available":1,"logistics_company":"日森集运","id":468},{"code":"NDJY","available":1,"logistics_company":"诺达集运","id":469},{"code":"DSJY","available":1,"logistics_company":"东澍集运","id":470},{"code":"CTJY","available":1,"logistics_company":"赤兔集运","id":471},{"code":"TDQQJY","available":1,"logistics_company":"通达集运","id":472},{"code":"DFJY","available":0,"logistics_company":"大发集运","id":473},{"code":"LTJY","available":1,"logistics_company":"龙通集运","id":474},{"code":"FRJY","available":1,"logistics_company":"FR集运","id":475},{"code":"TMJY","available":1,"logistics_company":"天马集运","id":476},{"code":"DHSHJY","available":1,"logistics_company":"缔惠盛合集运","id":477},{"code":"FSJY","available":1,"logistics_company":"flash集运","id":478},{"code":"MTJY","available":1,"logistics_company":"易淘集运","id":479},{"code":"XLJY","available":0,"logistics_company":"西里物流","id":480},{"code":"ZTGJJY","available":1,"logistics_company":"中通国际集运","id":481},{"code":"786EXPRESS","available":1,"logistics_company":"786 Express","id":482},{"code":"YMTRANS","available":1,"logistics_company":"YM TRANS","id":483},{"code":"BEEPOST","available":1,"logistics_company":"Bee Post","id":484},{"code":"HUOLALA","available":1,"logistics_company":"货拉拉","id":485},{"code":"WPJY","available":1,"logistics_company":"WePost集运","id":486},{"code":"JLJY","available":1,"logistics_company":"KEC-嘉里集运","id":487},{"code":"SHIBJY","available":1,"logistics_company":"十邦集运","id":488},{"code":"JIEJINGJY","available":1,"logistics_company":"捷竞集运","id":489},{"code":"SLHT","available":1,"logistics_company":"SMARTEX","id":490},{"code":"PSBYEX","available":1,"logistics_company":"顺捷速递","id":491},{"code":"AFL","available":1,"logistics_company":"AFL集运","id":492},{"code":"ZHONGYOUJY","available":1,"logistics_company":"中邮集运","id":493},{"code":"878EXPRESS","available":1,"logistics_company":"878 Express","id":494},{"code":"YUNDAJY","available":0,"logistics_company":"uda Express","id":495},{"code":"FBKY","available":1,"logistics_company":"飞豹快运","id":496},{"code":"EBUY","available":0,"logistics_company":"eBuy","id":497},{"code":"YIYUAN","available":0,"logistics_company":"壹圆国际","id":498},{"code":"GHT","available":0,"logistics_company":"GHT EXPRESS","id":499},{"code":"SLGJ","available":0,"logistics_company":"首领国际","id":500},{"code":"CCCP","available":0,"logistics_company":"CCCPexpress","id":501},{"code":"ONEX","available":0,"logistics_company":"ONEX","id":502},{"code":"GLOBBING","available":0,"logistics_company":"Globbing","id":503},{"code":"INEX","available":0,"logistics_company":"INEX","id":504}],"request_id":"17176535165381964"}}`

type Response struct {
	LogisticsCompaniesGetResponse LogisticsCompaniesGetResponse `json:"logistics_companies_get_response"`
}
type LogisticsCompaniesGetResponse struct {
	LogisticsCompanies []LogisticsCompanies `json:"logistics_companies"`
	RequestID          string               `json:"request_id"`
}
type LogisticsCompanies struct {
	Code             string `json:"code"`
	Available        int    `json:"available"`
	LogisticsCompany string `json:"logistics_company"`
	Id               int    `json:"id"`
}
type SysShippingListChannels []SysShippingListChannel
type SysShippingListChannel struct {
	Id                int64  `json:"id"`
	SysShippingListId int64  `json:"sys_shipping_list_id"`
	ShippingCode      string `json:"shipping_code"`
	ChannelType       string `json:"channel_type"`
	ShippingName      string `json:"shipping_name"`
}
type SysShippingList struct {
	Id           int64  `json:"id"`
	Status       uint8  `json:"status"`
	ShippingCode string `json:"shipping_code"`
	ShippingName string `json:"shipping_name"`
}

const (
	SysShippingListTable        = `sys_shipping_list`
	SysShippingListChannelTable = `sys_shipping_list_channel`
)

func main() {
	var err error
	var engine *xorm.Engine
	engine, err = xorm.NewEngine("mysql", "root:ccqpolicP#ascfsa3!@tcp(172.16.9.113:3306)/oms_saas?charset=utf8&interpolateParams=true")
	if err != nil {
		log.Fatalln(err)
	}
	//engine.ShowSQL(true)
	//fmt.Printf("%+v\n", sysShippingLists)
	response := &Response{}
	err = json.Unmarshal([]byte(JSONString), response)
	if err != nil {
		return
	}
	var sysShippingListChannels = make(SysShippingListChannels, 0)
	for _, logisticsCompanies := range response.LogisticsCompaniesGetResponse.LogisticsCompanies {
		if logisticsCompanies.Available == 0 {
			continue
		}
		sysShippingList := SysShippingList{}
		b, err := engine.Table(SysShippingListTable).Where("shipping_code = ? or shipping_name like ?", logisticsCompanies.Code, strings.Join([]string{"%", logisticsCompanies.LogisticsCompany, "%"}, "")).Get(&sysShippingList)
		if err != nil {
			log.Fatalf("sysShippingList:%s", err)
		}

		if !b {
			//不存在
			sysShippingList.ShippingName = logisticsCompanies.LogisticsCompany
			sysShippingList.ShippingCode = logisticsCompanies.Code
			sysShippingList.Status = 1
			_, err = engine.Table(SysShippingListTable).Insert(&sysShippingList)
			if err != nil {
				log.Fatalf("sysShippingList insert fail:%s", err)
			}
		}
		sysShippingListChannel := SysShippingListChannel{}
		b, err = engine.Table(SysShippingListChannelTable).Where("sys_shipping_list_id = ? and channel_type = ?", sysShippingList.Id, "PDD").Get(&sysShippingListChannel)
		if err != nil {
			log.Fatalf("SysShippingListChange:%s", err)
		}
		if b {
			sysShippingListChannel.ShippingCode = strconv.Itoa(logisticsCompanies.Id)
			sysShippingListChannel.ShippingName = logisticsCompanies.LogisticsCompany
			_, err := engine.Table(SysShippingListChannelTable).ID(sysShippingListChannel.Id).Cols("shipping_code", "shipping_name").Update(sysShippingListChannel)
			if err != nil {
				return
			}
		} else {
			sysShippingListChannels = append(sysShippingListChannels, SysShippingListChannel{
				SysShippingListId: sysShippingList.Id,
				ShippingCode:      strconv.Itoa(logisticsCompanies.Id),
				ChannelType:       "PDD",
				ShippingName:      logisticsCompanies.LogisticsCompany,
			})
		}
		//fmt.Printf("%+v\n", sysShippingList)
	}
	if len(sysShippingListChannels) > 0 {
		engine.Table(SysShippingListChannelTable).Insert(sysShippingListChannels)
	}

}