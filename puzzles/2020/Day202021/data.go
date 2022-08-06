package Day202021

var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string, int) {
	return 2020, 21, "Allergen Assessment"
}

func (td dayEntry) PuzzleInput() string {
	return `jtgktc gnvtb sfb vdlsx sdh przbq cvslrn ccntj hvbfjv rkvs jctgkq cfkv kjz rkkrx fvxql pgfzp vrqkb bqkndvb bmrmhm frdf xgms pvhf cccmn glbktsc phxqzn hqqt bqtvr qxgsq gpddp dtsjv jcpb txjkrbh cnj rlvkp ldxh qlcsjhg fcfvhz vflms zcx jrvgpbht bxtztd bvpq rgf qzkjrtl xfsrjp mzvqjf rtjzgg sbkztg lrc bbkzrl hfrhh hnbvdq snhrpv qknlg khjjdk (contains shellfish, fish)
qgbft rkkrx ljrblp htjggs bqtvr frlgls sbkztg fcgbkq zhqlcr qlcsjhg bxtztd mtdbxq hzdk gnvtb fjvdbs nljs mmhv lffcp lrc hvtzz glbktsc phqjq czvjt fglxll hnknpb hqqt sdh bmrmhm ltftm cfkv bqkndvb bdqbc jmctq jkrsn qzkjrtl gbpqhk vflms tplzk nsdrd qtdhj snhrpv fcfvhz rtjzgg dgmb nfrln trl bbkzrl frdf vknkn hnbvdq rlvkp jvrhm pzzkssr sbvcdk tzdz qgqbpmp lrzfcf fvxql mcsbdnpt zfrp phxqzn tqsjd shmzdgxm sqp zlxtk mmzn gpddp drgqqp ldxls (contains peanuts)
bmrmhm kjz fjvdbs dtsjv khjjdk vrqkb zhqlcr sqp jrvgpbht ttvvt vxkzt nrrh przbq phqjq gpddp htcblx htsrn tplzk frdf mtdbxq bdqbc fcfkxs rh ltftm frlgls zfrp sdh drgqqp nfrln qlkdr bxtztd pzzkssr htrvf qzkjrtl jctgkq dgmb nhxjbx hfrhh rkkrx gnrxbp trpllnr lflxjnh fsskh hrbrc gbpqhk rkvs zmb gbqfd szj snhrpv ggkz pgfzp jkrsn ldxls fqxx lzbjg lxfd phxqzn qlcsjhg qgbft vflms bqtvr nbzxn (contains peanuts)
drgqqp sqp jlvmh shmzdgxm pgkttq lrc vflms nfzn jvpslhn dcqz nhxjbx bmppb czvjt bmrmhm snhrpv qmlndrj qvmtms lgtnxs jrvgpbht nrhftzl htcblx fvxql gpddp svqtgll lrzfcf fjvdbs vfgqpk fglxll vrqkb ztmkvt lxfm tplzk bqkndvb xfsrjp zlxtk mtdbxq htsrn ttvvt hfrhh rkkrx sbvcdk mcsbdnpt sdh msrg zmb tcbbkg trl bqtvr gbpqhk jkrsn hskzxjz (contains peanuts, nuts)
qbnzjb hrbrc bqtvr qxgsq ldxh zmb pvhf jkrsn htrvf dkgs vvpxz lxfm tcbbkg tktj jrjfv gjnnlz kxbvz jxxbr bxtztd rgf vfgqpk jvpslhn vknkn nrhftzl zhnlt jrvgpbht rxv pgfzp ljrblp jcgfs clcr fjvdbs fcfkxs hnbvdq tplzk vtxnv frlgls rxbx qgqbpmp jctgkq szj mzvqjf lxfd khjjdk qlcsjhg zcx jmctq lnrd bmrmhm jcpb bdqbc bqkndvb nsdrd shmzdgxm qzkjrtl snhrpv tjxbjsj zkhdzrb rkkrx ccntj ghljzv svqtgll trpllnr ztxf lkfzc gpddp mcsbdnpt (contains fish, eggs)
nrrh clcr zhqlcr rkvs qzkjrtl qlr glbktsc vnm gnvtb bmrmhm nbzxn pbgff qtdhj gbqfd nsdrd kjz bmppb kqsnfjf bvv srzvnxlx jcgfs snhrpv lkfzc hplfr htcblx bxtztd gbqtp lzbjg vflms nfrln jlvmh jvpslhn rkkrx qlcsjhg rlvkp qbnzjb fjphsbb xgms ghljzv jvrhm qgbft rvrzn vxkzt tktj kdm msrg hskzxjz jrvgpbht csxvh fjvdbs hzdk gjnnlz mtdbxq lgtnxs ljrblp zlxtk jctgkq vprjz gbpqhk hmzvmm qknlg htjggs phxqzn ltftm qlkdr bqkndvb fglxll krzchjs hxs dsjvgg shmzdgxm hmjs bjbpq nljs gpddp mmhv fqxx zmb bvpq (contains dairy, peanuts)
hcgsh hfrhh gcgqb pfszh bvv lzbjg htcblx pzzkssr bbkzrl gbpqhk ldxls gnrxbp bqtvr bmrmhm jkrsn zmb dnbzjcg ztmkvt jjjjkmnt bxtztd gpddp cjgxf bmppb jth rvrzn jcpb tzdz zhqlcr dtsjv cczcxsl lrc vtxnv fsskh ljrblp lnrd jctgkq ddt snhrpv lxfd mmzn dgmb gbqtp vvrdp qlr vflms hxs jrjfv tjxbjsj jvpslhn sbvcdk rh szj nhxjbx htjggs ztxf przbq pvhf hnbvdq kqsnfjf qvmtms qzkjrtl kdm bvpq clcr qlcsjhg rkkrx sbkztg nbzxn gzlj vdlsx kjz xfsrjp jvrhm lrzfcf dsjvgg (contains fish, wheat, eggs)
nhxjbx lffcp sfb hvbfjv trpllnr qtdhj cfkv szdnqf tbb snhrpv jxxbr lflxjnh gpddp dcqz khjjdk hrbrc ljrblp lgtnxs ggkz fjphsbb lxfm rkkrx cczcxsl zhqlcr jth hfrhh gjnnlz jctgkq ccntj mn bjbpq jjjjkmnt lrzfcf rkvs htrvf lxfd jcgfs pvhf bdqbc bxtztd vvpxz rxp mzvqjf rh nsdrd bqkndvb mmzn hmjs bmrmhm vflms fjvdbs hxs pgkttq ttvvt nbzxn mtdbxq qknlg fcfkxs nljs rvrzn jlvmh jslg vtxnv mmhv qgbft cgplg zmb hnknpb bqtvr fcgbkq pbgff nrrh (contains nuts, wheat)
przbq pbgff vflms kqsnfjf qzkjrtl pgkttq pgfzp txjkrbh zmb bqtvr sdh ccntj ldxh jkrsn zfrp cvslrn vnm lflxjnh mmzn khjjdk vrqkb nfrln cnj vfgqpk nbzxn tbb jrvgpbht lxfd qknlg rtjzgg lrzfcf bvv czvjt vxkzt dgmb ztmkvt fcfkxs qlr rgf rxp frlgls mmhv htcblx phxqzn lnrd snhrpv tzdz vdlsx bbkzrl mcsbdnpt gzlj kdm vvrdp fglxll rkkrx bjbpq bqkndvb jcgfs fcfvhz gcgqb bvpq jvpslhn ltftm xgms rlvkp tplzk mtdbxq tcbbkg ljrblp mzvqjf phqjq (contains wheat, eggs)
ghljzv snhrpv ttvvt przbq bqtvr jslg sqp trpllnr gzlj qzkjrtl gbqfd glbktsc rxv vflms kqsnfjf jsfl sbvcdk zmb cfkv gcgqb hvtzz qmlndrj bvpq tktj fdl krzchjs ggkz vknkn fqxx bxtztd jvrhm bmrmhm bdqbc lgtnxs bqkndvb nbzxn qknlg (contains eggs)
nrrh frdf vrqkb bdqbc dcqz jjjjkmnt ldxls bmrmhm vknkn ztxf czvjt dkgs msrg trl phqjq nljs drgqqp dnbzjcg rkvs kjz zlxtk xgms hqqt ggkz ghljzv sfb bbkzrl bqtvr rgf zpqnhq lnrd vtxnv rkkrx qgqbpmp jvrhm zmb rxv lffcp sbkztg fcfkxs jslg kdm snhrpv mzvqjf vfgqpk mtdbxq bmppb hskzxjz lflxjnh qbnzjb tplzk zhqlcr jrjfv bjbpq qlcsjhg kxbvz sdh pgfzp ljrblp jth csxvh tbb qzkjrtl fsskh jrvgpbht nbzxn bqkndvb jctgkq (contains fish, wheat, soy)
ldxls nrhftzl qknlg qvmtms vrqkb ddt csxvh phxqzn jvrhm bmrmhm mcsbdnpt rlvkp hzdk clcr gnvtb frlgls rtjzgg lflxjnh lgtnxs sfb jkrsn qlr vtxnv vflms fjphsbb lxfd snhrpv jtgktc dnbzjcg pvhf bvv qzkjrtl frdf phqjq glbktsc ttvvt htsrn cczcxsl tcbbkg qgbft vvpxz fsskh bqkndvb mtdbxq htcblx hmzvmm tplzk hnbvdq bvpq bqtvr mmzn hxs jrjfv ldxh zmb nxnbpj srzvnxlx bbkzrl nfrln bmppb jmctq zlxtk czvjt rkvs (contains nuts, shellfish, eggs)
vxkzt txjkrbh nfrln frdf ljrblp zfrp srzvnxlx qvmtms dgmb nbzxn htjggs bqkndvb hplfr gnvtb rkkrx vfgqpk rkvs vnm trl kjz rxbx lffcp vflms bmrmhm zlxtk hcgsh vprjz jkrsn cgplg qknlg qzkjrtl fdl gnrxbp msrg rh tbb mzvqjf ldxh fjvdbs lnrd bqtvr bvv phxqzn jcpb lrzfcf sqp mmzn zhqlcr sdh hmzvmm jrjfv zpqnhq cjgxf tzdz zmb cvslrn fsskh (contains wheat, fish)
bmrmhm qlkdr bqkndvb bqtvr qgbft bvpq txjkrbh hmzvmm vtxnv lxfm mmhv fglxll zfrp hplfr zlxtk kxbvz rgf srzvnxlx pbgff ldxh tplzk fcfvhz frlgls xfsrjp fcgbkq kjz pzzkssr rkkrx gbqtp fqxx phqjq vvrdp snhrpv hqqt kdm dkgs bvv tbb hcgsh szj czvjt zmb xgms mmzn vflms qgqbpmp jxxbr zkhdzrb nbzxn ljrblp pgkttq zhnlt lflxjnh mn clcr htjggs (contains dairy)
qgbft rxv nfrln pbgff lffcp pvhf lxfm jsfl fcgbkq jcgfs jvpslhn gjnnlz jkrsn rkkrx gnvtb gcgqb sdh gnrxbp qzkjrtl snhrpv frlgls frdf glbktsc bqkndvb bbkzrl tktj vflms jslg zlxtk fglxll vvrdp vfgqpk zmb lflxjnh zpqnhq tplzk gbqtp ghljzv jrvgpbht rlvkp qgqbpmp pfszh przbq bqtvr tcbbkg vxkzt ttvvt clcr mzvqjf zhqlcr ggkz (contains dairy)
nsdrd pbgff ccntj jlvmh qgbft dnbzjcg ztxf bqkndvb jrvgpbht fqxx bqtvr nxnbpj gpddp zmb ghljzv ltftm jsfl gnrxbp hqqt qlkdr svqtgll mmhv fcfvhz zhnlt sdh khjjdk jcgfs cgplg dkgs hxs hfrhh bmrmhm hzdk qvmtms szdnqf lflxjnh lkfzc bvpq hplfr mcsbdnpt gzlj glbktsc ldxls nbzxn jkrsn sbvcdk trpllnr cczcxsl vflms rtjzgg tcbbkg zfrp jvrhm nfrln htjggs jctgkq nljs przbq ttvvt tplzk krzchjs hvbfjv mmzn rkkrx snhrpv frdf dcqz frlgls htrvf pzzkssr rvrzn hcgsh vnm (contains fish, eggs, peanuts)
vnm trpllnr bvv lrzfcf qgqbpmp fdl tcbbkg mcsbdnpt nxnbpj rkkrx ggkz qbcg bmrmhm dsjvgg pbgff jlvmh fqxx tzdz bqtvr fcfkxs mzvqjf nfzn rtjzgg czvjt jvrhm lgtnxs jmctq pgkttq hcgsh kjz dtsjv lxfm tjxbjsj bdqbc qzkjrtl cgplg hnbvdq rkvs jctgkq jtgktc gnrxbp dgmb jcgfs tqsjd gcgqb rxbx jcpb cfkv zkhdzrb vdlsx qlkdr gbpqhk zcx vknkn hskzxjz vflms zmb ltftm pfszh bqkndvb jxxbr sbvcdk dnbzjcg nljs szdnqf krzchjs phqjq sbkztg nhxjbx mtdbxq (contains wheat, fish)
bjbpq xfsrjp hrbrc dgmb gjnnlz lrzfcf szj jrvgpbht vdlsx ldxls fsskh cccmn vxkzt clcr cczcxsl krzchjs hmjs dtsjv tktj lxfd htcblx vflms bvv lrc srzvnxlx qknlg phxqzn rkvs ztxf fqxx shmzdgxm cvslrn cgplg hcgsh nfzn trpllnr phqjq qgbft qlcsjhg jmctq jslg pgfzp trl gnrxbp cjgxf bqkndvb zmb mmhv ddt rxp lnrd qzkjrtl lzbjg tcbbkg zhnlt dcqz frdf sfb rkkrx jlvmh gbpqhk bmrmhm rh przbq qbnzjb mn rvrzn hskzxjz snhrpv jkrsn pzzkssr dkgs tzdz qtdhj glbktsc cnj kjz mcsbdnpt hnbvdq fvxql nrhftzl nfrln hvbfjv (contains eggs)
ttvvt ztmkvt hvbfjv nxnbpj hplfr lkfzc rxbx srzvnxlx nhxjbx nrhftzl ltftm qlr txjkrbh czvjt hcgsh gnrxbp bjbpq jsfl zhnlt jslg sbvcdk bbkzrl sqp lnrd lflxjnh qgqbpmp gjnnlz vvpxz bmrmhm jmctq gcgqb mn ghljzv shmzdgxm jth lxfm dsjvgg qknlg pgkttq vflms szdnqf nbzxn dkgs khjjdk rkkrx ddt vknkn tqsjd htrvf tjxbjsj qmlndrj sfb qzkjrtl fjphsbb bvv rxv rxp bqkndvb bdqbc fsskh rkvs zmb phqjq qbnzjb cjgxf kjz pfszh jcgfs jcpb jjjjkmnt fcfkxs cgplg fvxql snhrpv tzdz krzchjs frdf bxtztd (contains nuts)
nrhftzl lrzfcf fdl fvxql nsdrd zkhdzrb zhqlcr zlxtk jkrsn ljrblp dgmb glbktsc hvbfjv pfszh tplzk pgfzp gpddp hxs sfb dkgs tjxbjsj qgbft bqkndvb snhrpv vdlsx hplfr jrvgpbht bbkzrl hvtzz vfgqpk sqp nfrln vflms cvslrn htcblx gbqtp bmrmhm qxgsq jcpb nxnbpj zmb qzkjrtl dsjvgg cczcxsl htrvf mzvqjf xfsrjp gbqfd cnj gcgqb ghljzv nfzn lffcp kjz rkkrx czvjt pzzkssr (contains dairy, nuts)
bjbpq fcfkxs sbkztg qzkjrtl dgmb bxtztd vvrdp bmppb trpllnr jkrsn ldxls dcqz lxfd dnbzjcg hxs cvslrn przbq clcr hskzxjz nhxjbx bvpq cccmn glbktsc cczcxsl hvtzz rlvkp ggkz tplzk lzbjg ddt rkkrx vknkn czvjt fqxx hzdk bqtvr rxbx lgtnxs fdl gcgqb jsfl zhnlt jrvgpbht qlcsjhg hnknpb jslg nfzn gpddp jmctq vflms fglxll jrjfv zlxtk vvpxz nsdrd tktj gbqfd bqkndvb pvhf bmrmhm cfkv jxxbr qxgsq sfb lnrd rkvs zmb vnm mtdbxq qknlg pgkttq cgplg (contains dairy, soy, peanuts)
bdqbc gcgqb dsjvgg hvtzz lffcp lnrd qlr fcfkxs csxvh gbqtp czvjt ddt qzkjrtl tbb htcblx kdm jth jkrsn dkgs sbvcdk jjjjkmnt zhnlt qknlg tcbbkg phxqzn snhrpv trl szj nsdrd zlxtk bvpq kqsnfjf hnknpb bqtvr pbgff zkhdzrb rkkrx vflms xfsrjp glbktsc vknkn ghljzv jxxbr zmb kxbvz jcpb vfgqpk jrjfv pfszh bmrmhm nljs fvxql cgplg cczcxsl cnj bbkzrl gnvtb (contains wheat, peanuts, soy)
htrvf pgfzp sqp rgf kdm pbgff pgkttq htjggs hmjs sbvcdk jmctq kxbvz ljrblp zfrp dkgs cczcxsl nhxjbx bmrmhm przbq ldxls snhrpv phqjq trpllnr rlvkp nrhftzl dgmb dnbzjcg vxkzt vknkn dtsjv krzchjs vvrdp qtdhj zlxtk txjkrbh qmlndrj rvrzn gpddp jrjfv cjgxf qlr lffcp qzkjrtl msrg rkvs qknlg tcbbkg xfsrjp nxnbpj lnrd mzvqjf ztxf gzlj hnbvdq pvhf zkhdzrb fdl bqkndvb tjxbjsj rkkrx vflms hplfr hxs xgms shmzdgxm ddt qgbft tqsjd jvrhm ztmkvt zmb qvmtms jcpb jslg hfrhh lrzfcf rxv (contains wheat, eggs)
qlcsjhg hqqt rlvkp nrrh htcblx lffcp fsskh krzchjs jtgktc bqtvr rkkrx tqsjd vnm lrzfcf dcqz nljs jlvmh kqsnfjf zhqlcr vxkzt zmb vflms tplzk bmrmhm ztmkvt cnj qvmtms jvpslhn vdlsx bxtztd jcpb qlkdr xfsrjp snhrpv gnrxbp rxbx zcx ghljzv tjxbjsj fqxx bqkndvb kxbvz msrg svqtgll trl rh ljrblp nxnbpj hvbfjv lrc nfzn nbzxn dtsjv dsjvgg fjvdbs frdf drgqqp dkgs ltftm ggkz htrvf jmctq nsdrd hmzvmm rvrzn fglxll sqp lxfd cgplg clcr sdh lzbjg vtxnv lnrd lgtnxs htjggs bvv (contains wheat)
hcgsh lzbjg przbq kjz vflms hxs fcfvhz ljrblp jsfl mzvqjf mtdbxq rkvs xgms tzdz rkkrx bqtvr txjkrbh sbvcdk vrqkb qtdhj rxbx bbkzrl jth vvrdp fvxql cjgxf vnm rvrzn htrvf khjjdk jvpslhn tqsjd kdm czvjt zmb qknlg pgfzp szdnqf csxvh trpllnr gzlj zlxtk bqkndvb hmzvmm bmrmhm lxfm qzkjrtl fjvdbs gnvtb kxbvz vprjz hskzxjz ddt vvpxz htjggs qmlndrj htsrn xfsrjp szj rxv (contains shellfish, fish, soy)
vflms frdf dnbzjcg bxtztd rvrzn mn hvtzz qgbft lrzfcf fglxll rxv rh ggkz trl zkhdzrb gzlj bjbpq rkkrx phqjq zfrp qbnzjb snhrpv pgfzp qgqbpmp szdnqf mmhv hxs rxbx sfb qbcg kqsnfjf csxvh htrvf shmzdgxm sqp vrqkb xgms qlcsjhg bqkndvb msrg bqtvr vvpxz hrbrc cjgxf dgmb bvv cczcxsl gcgqb lxfm jcgfs gbpqhk cnj lzbjg szj bdqbc nsdrd jrjfv qzkjrtl pgkttq vnm ljrblp hfrhh drgqqp nfzn kdm jvpslhn vfgqpk zhnlt nbzxn zmb jctgkq tktj mzvqjf tcbbkg pvhf trpllnr qlr (contains shellfish, dairy)
fglxll qlcsjhg tbb hskzxjz lgtnxs zmb fdl vrqkb trpllnr lxfd bmrmhm hqqt cgplg sbvcdk zlxtk tqsjd ldxls zcx rkkrx txjkrbh vdlsx gbqtp lflxjnh qbcg kqsnfjf vfgqpk jvpslhn pbgff rkvs jlvmh dsjvgg qmlndrj bqkndvb htjggs cccmn qzkjrtl kjz bqtvr svqtgll gnvtb gcgqb snhrpv szdnqf zhnlt jxxbr jth qgbft krzchjs lkfzc dtsjv nfzn jvrhm fvxql qgqbpmp dkgs jjjjkmnt bvv phqjq hxs vxkzt hzdk ltftm mcsbdnpt csxvh cnj lrc hnknpb drgqqp jcgfs (contains fish, nuts, soy)
ghljzv zfrp bqkndvb zmb fsskh mcsbdnpt trl bqtvr rkkrx qzkjrtl xfsrjp bvv rtjzgg jctgkq qlcsjhg snhrpv ztxf pvhf mtdbxq htrvf ldxh nhxjbx tktj bdqbc kxbvz szdnqf vvpxz gcgqb htcblx qtdhj bvpq hvtzz sdh rlvkp przbq nfrln lzbjg bmrmhm hmjs khjjdk qxgsq mmzn fcfvhz fdl lgtnxs ddt jcgfs rkvs qbnzjb phqjq bjbpq gbqfd fjphsbb cnj gpddp zcx hxs tjxbjsj ccntj szj vdlsx ggkz fjvdbs nbzxn tqsjd jrvgpbht txjkrbh nrrh hfrhh bmppb hskzxjz vxkzt zhnlt bxtztd clcr sfb jtgktc cvslrn gjnnlz (contains eggs, fish)
glbktsc kdm zfrp fglxll jslg hzdk bqkndvb qlkdr ldxh szdnqf hxs nbzxn sdh mcsbdnpt qmlndrj fvxql bqtvr snhrpv nhxjbx hmzvmm fjvdbs qvmtms vnm trl frdf gzlj tcbbkg qzkjrtl gbqtp bmppb csxvh zmb qgqbpmp rkkrx lzbjg rxp phxqzn rtjzgg lffcp bmrmhm gbqfd phqjq clcr rvrzn hrbrc bbkzrl (contains soy, fish, nuts)
phqjq qbcg ztxf kjz bmrmhm hvtzz bvv qlr xgms jkrsn vfgqpk jslg fcgbkq mzvqjf htcblx jvpslhn vflms lrzfcf nsdrd snhrpv jlvmh krzchjs rkvs zpqnhq txjkrbh sdh lrc phxqzn trpllnr cccmn sqp glbktsc vnm cczcxsl mn mmzn tqsjd qmlndrj hfrhh zmb clcr hnbvdq vvrdp fjvdbs nfzn vknkn hxs dtsjv cvslrn jth bqkndvb rh tcbbkg rtjzgg tbb pfszh szj jjjjkmnt qzkjrtl frdf bqtvr rlvkp hnknpb pzzkssr dcqz (contains shellfish, dairy, nuts)
glbktsc mmzn tbb szj vflms qlkdr vvpxz rvrzn cnj htcblx vnm cgplg clcr jcgfs vfgqpk qtdhj ggkz xgms gbqtp snhrpv zpqnhq tktj gcgqb bxtztd hzdk zlxtk hvtzz przbq drgqqp bmrmhm jslg sdh frlgls lrzfcf ttvvt fcfkxs ltftm rgf fjvdbs qzkjrtl ztxf zfrp jjjjkmnt tplzk rkkrx nxnbpj bqkndvb qbnzjb hfrhh jvpslhn gnvtb ghljzv bqtvr nsdrd jmctq qknlg trpllnr ddt htsrn jth rxbx (contains soy, wheat, peanuts)
vfgqpk hcgsh bbkzrl jxxbr bmrmhm nrrh bqkndvb tktj ttvvt rtjzgg rvrzn sbkztg nbzxn hplfr kdm dsjvgg msrg cgplg pbgff pgkttq vvpxz hrbrc jsfl lxfm zhnlt qzkjrtl cjgxf vrqkb hskzxjz hmjs fjphsbb lzbjg qvmtms sfb cnj jth sqp rgf zlxtk vflms jrjfv tjxbjsj mcsbdnpt zfrp lrc bqtvr bmppb csxvh rkkrx snhrpv jjjjkmnt vvrdp lflxjnh jlvmh fcfvhz drgqqp vnm nfzn lffcp gcgqb zkhdzrb dgmb hqqt hvtzz mtdbxq qlkdr (contains peanuts)
frlgls rtjzgg tcbbkg cfkv dsjvgg snhrpv zpqnhq jvrhm pgkttq hfrhh vflms ljrblp hcgsh zcx nljs qgbft bqtvr bjbpq ccntj nsdrd jrvgpbht qlcsjhg qzkjrtl bmrmhm glbktsc rkvs zlxtk tbb jrjfv vxkzt bvpq cgplg pzzkssr lnrd zfrp phqjq bqkndvb mn hnbvdq tjxbjsj rxv txjkrbh ltftm vprjz vvpxz fjphsbb jth khjjdk gjnnlz mmhv jslg drgqqp ztmkvt tqsjd zmb bdqbc ghljzv hnknpb qknlg gnrxbp csxvh bvv lkfzc hqqt krzchjs bxtztd szj sbvcdk hskzxjz cnj nfzn pgfzp cjgxf msrg jjjjkmnt gbpqhk lflxjnh xgms qvmtms (contains dairy, shellfish)
nrrh bqkndvb kdm glbktsc qxgsq hnbvdq svqtgll pgfzp gnrxbp jvrhm lxfm nxnbpj snhrpv fjphsbb zmb vnm kqsnfjf pgkttq hrbrc phqjq jxxbr bqtvr hzdk dtsjv cccmn lnrd fdl shmzdgxm qzkjrtl qtdhj hvbfjv csxvh htsrn gbqtp tjxbjsj cczcxsl hmjs zhnlt fvxql lxfd pzzkssr msrg ztmkvt bmrmhm xgms jslg nljs rkkrx (contains nuts, eggs, wheat)
snhrpv lzbjg ztmkvt tktj czvjt cczcxsl hzdk vflms fdl hxs mmhv gpddp nfrln zmb htsrn kdm bmrmhm jth gnrxbp jrvgpbht mn rkkrx zhqlcr htrvf drgqqp vnm qlr fglxll ztxf pgkttq bqkndvb pgfzp msrg gjnnlz trl lffcp gcgqb htcblx qlcsjhg qzkjrtl csxvh ljrblp jmctq cgplg sfb przbq glbktsc hvtzz ggkz (contains wheat, dairy, soy)
gpddp jjjjkmnt trl jvrhm vflms tplzk cnj qbnzjb jlvmh gnrxbp frlgls mtdbxq fglxll ztmkvt mzvqjf htsrn hxs fvxql rkvs hcgsh vtxnv jctgkq cvslrn qlr qzkjrtl rkkrx bqkndvb nsdrd pbgff qxgsq bqtvr zpqnhq vvrdp qknlg drgqqp szj bmrmhm ldxls gjnnlz snhrpv cczcxsl jsfl fqxx trpllnr zlxtk ddt (contains wheat, dairy, nuts)
qzkjrtl srzvnxlx rkkrx cgplg fjphsbb jkrsn bqkndvb lnrd fdl gpddp lxfm hvbfjv fcfvhz gjnnlz pgfzp vrqkb vflms nsdrd rlvkp bmrmhm lflxjnh vvpxz jth phxqzn qknlg rgf nhxjbx snhrpv lffcp gbqtp nrrh lxfd lrc hmzvmm bjbpq frlgls kdm bvpq cczcxsl vvrdp kxbvz fcgbkq nbzxn zcx krzchjs fglxll rkvs sqp qtdhj tbb zkhdzrb pzzkssr gnvtb tcbbkg jxxbr mtdbxq tzdz qgbft zfrp lkfzc hfrhh dsjvgg rxp qxgsq bmppb jcpb vdlsx hqqt gbpqhk cnj zmb (contains peanuts, eggs)
phxqzn zkhdzrb fcfvhz qtdhj fcgbkq hqqt jth gbqfd rxv czvjt rkkrx jlvmh tktj htjggs pgfzp sdh nsdrd qgqbpmp qlr lkfzc vflms mzvqjf vvrdp vfgqpk zhqlcr bqtvr bmrmhm qknlg hrbrc rh shmzdgxm bqkndvb fjvdbs fsskh lrc qbnzjb ztmkvt ccntj snhrpv vvpxz lgtnxs hmjs szj hvbfjv hskzxjz trl qlcsjhg jjjjkmnt zcx hnbvdq hvtzz hmzvmm khjjdk rkvs qzkjrtl mn qvmtms (contains soy)
rgf qknlg rlvkp lflxjnh zhqlcr lrc clcr gjnnlz htjggs htrvf gcgqb srzvnxlx vrqkb xfsrjp bmppb nsdrd kjz jkrsn jslg qgqbpmp fcfkxs kdm kqsnfjf przbq dsjvgg gbpqhk trl zcx hplfr nrhftzl qlr jrvgpbht tqsjd ztxf nbzxn hvbfjv lxfd hvtzz dnbzjcg qgbft lzbjg cccmn mtdbxq jxxbr vnm shmzdgxm hrbrc jctgkq qvmtms glbktsc qtdhj vvrdp qzkjrtl cvslrn htcblx jvrhm pgfzp pgkttq tjxbjsj khjjdk jmctq rkvs gbqfd bmrmhm jth zmb jcpb hzdk vvpxz zkhdzrb rkkrx zlxtk ghljzv vflms hxs dgmb fqxx pbgff cfkv txjkrbh lkfzc pzzkssr qmlndrj snhrpv dkgs csxvh bvpq vknkn bqkndvb vprjz msrg ljrblp fcgbkq cjgxf czvjt ldxh mmzn (contains fish, dairy, eggs)
rlvkp hskzxjz glbktsc vrqkb gcgqb kjz cnj ddt trpllnr ldxh bmrmhm gpddp nbzxn lffcp cccmn zhnlt pgfzp hfrhh vknkn vdlsx jkrsn bqkndvb vflms vprjz phxqzn jvrhm qgbft bvv vtxnv rkkrx tzdz ttvvt phqjq tqsjd hmjs jrvgpbht htcblx zfrp bbkzrl khjjdk hnknpb bqtvr zmb cczcxsl snhrpv jmctq czvjt lrzfcf (contains nuts, fish)
rxbx dtsjv snhrpv phqjq mn pgkttq krzchjs hzdk vflms fsskh ltftm qgqbpmp zmb hmjs tbb cgplg sbvcdk cjgxf bmrmhm bqtvr qzkjrtl zfrp ldxh fjphsbb nhxjbx ccntj zhqlcr lzbjg szj vfgqpk qlcsjhg qxgsq hqqt bvv dcqz vdlsx jvpslhn lxfm pzzkssr hmzvmm rxv qvmtms qbnzjb fdl jmctq przbq jslg zhnlt pvhf szdnqf gbqfd lffcp qlkdr bqkndvb gpddp gnvtb (contains eggs, dairy, peanuts)
bdqbc ggkz szdnqf fsskh nbzxn dtsjv clcr qxgsq jlvmh mzvqjf fcgbkq hqqt rkkrx mmhv vtxnv gzlj zkhdzrb jrjfv vfgqpk qgbft fdl lgtnxs snhrpv fglxll jvpslhn pfszh bqtvr bqkndvb zcx pvhf jcpb hxs przbq vvrdp nrrh tqsjd hzdk ldxh qvmtms hcgsh vflms mcsbdnpt gbqtp zlxtk qzkjrtl kjz gcgqb bxtztd szj nrhftzl srzvnxlx tplzk bmrmhm ghljzv lzbjg jvrhm dgmb nhxjbx lxfd (contains peanuts, shellfish)
qzkjrtl zkhdzrb tbb pgkttq gbqfd jctgkq vprjz rkkrx zhnlt frlgls qbnzjb ggkz msrg bmppb shmzdgxm jcgfs nhxjbx lnrd zfrp jrjfv cccmn vvpxz pvhf nljs tqsjd qlcsjhg tktj fglxll qvmtms nxnbpj mmhv kqsnfjf hnbvdq fqxx nrhftzl nfzn lflxjnh jjjjkmnt svqtgll dsjvgg gpddp kxbvz hfrhh dcqz ddt bqkndvb vtxnv lkfzc jslg zmb krzchjs cfkv hmjs lrzfcf mmzn nrrh bqtvr qbcg vflms hvtzz jkrsn ccntj vfgqpk lzbjg htrvf dkgs rxp gzlj hxs lxfm jlvmh gbqtp khjjdk bmrmhm (contains peanuts)`
}

func (td dayEntry) PuzzleSpec() string {
	return ``
}
