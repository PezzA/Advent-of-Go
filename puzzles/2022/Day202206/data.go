package Day202206

var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string, int) {
	return 2022, 6, "Tuning Trouble", 2
}

func (td dayEntry) PuzzleInput() string {
	return `cdhccdbdggfjjgssjzjzggjnjpnpbbzbnzzflfjfnfrrpvrvbrvvrvggvlvnnbrnrcncsnndbndbnndbdndfdrdvrvvndvvbggnrrnbrnntffgttwzwnnmvmcvvhsstzzlnlwlttbzzpnpmnnjvjnntmnmfftwwrfwrwswmmfrrfrrgbrbffwvvshvhrhmhththbbmqbmqqlslhssrmmqdmmjtmtmjtmjtttnwnvwvqwqjjnbbbdbqbnbpnbnllglcglcgcdczdznnqhhfthtmtlldqlqrrmddrldlzdllvddjcddqfqbqsbqqnllwppqpqzzrbbdppzsppjdpdqpdqdfqfrrwbwrwwqcqcsqsvvpbvbbztzptzzpccdtdhdffvqvcvzzmwzwddjfdffplplqlvlmmmvggpmpvpddpbptplpvlplvpvvnrvnnbqqqjhhwfhwfhwhqhmmpphqpqvppfzpzjzddgzzwffjmjggwhwwnnmlmpmmhcmcpcrcddvzvpzzwnznfznffgdgvddvtvgvsvdsdbbjnjtntbnttgbbbvgvgrgrzrvzrzddlsddcndcnnfqnnmpmlppdlplzplzpzgzmzmddlvlnnbttbwwhbhdhfdfssjppmcpplpdddnpdnnljlwjljsjnjhnhvhvqqsffrbbdttjdjndjdwwsfspffnhfhhlvhvmmqjmmwzwszwwvdvpdvdbdtdsdtsshvvmtvmtmctclchccrllznzfffpjjvhhdmhhvphpghgsgmmhlhnlnmnlnslnlgngznnsqnqddllpwllmzmjmttptfpplglqlgglgqqptqqmvmtmjmddcchbblltslsvsmvmgghmmccnzcztczzmnmttrdrvvcvzvvzllbhllnldndbbqffbbgtgddbtdbbzttdptpccjnjppbllbzlblrlcllhrrhqhgqqbcqcvcdvvnnzfzvzttrptrrwmrmlrlddvttdbtdbdcdvccwlcwwhphmppwfppclpcllgqgnghhvfflfggrzrcchfhhrdhrdhdnhnmmhjjwqjjpmmwvmmdnmnzzqmqwwmtthtdhtthnnqhqdhqqndqqwffsbspbptpmmndnllsmmdhmhfhnhjhghshlslppbgpgngddlsljsjmmzqzhhswhssfzssfqqcmqcmqcmmqggjcjvvgssrccwddmpmwwdfdpdbdpdwdvvqfvfrvvvsbvbhvvmqqcjqqvzqzppncnhhqnnpgplpqqpjpbblpbbbshsthhvfhfmmqzmmznnvrvqrrwdrdlrlwlttzqttjvttqltqqnfqqqwjqwwqttfstftjffsqqnhhnsnqqhggbsgghfgglslmssqlqhlhthqhccdsspsnssshbbnmngnnhllwclcffqllsrszrssnqsqvqjvjcvcttqgqbqmmfqfsqfsqswwvcvffndnfdfvfcvvggsmsfmfwfpfwwzhznntgtlglmmlfmllwrlrwwhcchqchhznzjjcdjdbjjhcjcscwwlnnsgngqqtgqgngnwgnnhqnnhchmchhtchcnclcmccgffbmmzvvrnngwwvddzccnjntjtwjwwztwtmtddjddpsptpbpbvvbwwnlnmndmnmdnnclnnbsbddpfdfvvjtjqqtqqqzjzlzqllzzwwlppvfffpcffffprrncnnzsnznhhwvvqhqphpjjgqqvnnmdmqqglqlblgglrlsspscsjjpvpbpjjwccslsppdjpdjjwvjjmhjhtjjwqqbqjqzjqzqpqbbswwlssqzssbjjpjqjbbjcjpjspjssjjzhhhnjhnhbnhhwzhzwlcshqlqpzgggzmcwntcwmfgtrwwjdpnbdqqcgnzgbdrzdmpwgvtvqffqbpvjpjrcfswffllnvnwvhclpjcwqwgnwqwvwsfgflrgzzsswffwjdjgvdvlgmczcbthwbvhggwzwlzfmhvwvjpbpnhcczbgfhhgghsmjwnvnsvnvmqwstrgnncwbqgbqpgdngllcqnzgwswpgtwzgqzggnzsdgltrlqfctqfvlzdswccfpdtjbfnrbqsmpjclnplbmqbmvwbzzdflwbqrljvzjpcrmnqsmrpqlmfsgcmthqpwgwzvmrjnhqczljcpnzjbwzrhjrzmcqpmlbzhgmqrlzsjbjsvcmcngptzlrthwsrjrlmsrgjlzrvpzwmprwnpgvjtspsppfvwfwcvbnqcwwmzlbqthqmbnbmnsnzgsbbnqtrvhlzjhphclpjzrdblszrnftqgwwrhpznhjhgrncvsvrmtmmgssvzddjfrnrzhbrqrfffjvzrqdnrdbvjwgrvlcpbncfgczlwdggsjmwzhndcdbggjvwfljctjnsjwczwfdrfttbhnlswfdbpcnwpspdhnzwqbgdswwpccbpfpgmfmvvwpzbzqsbbjbfnhjpszcbnrdplmwtdjtpcsztdjcmczltnstzwlcdbtdhsdgsgtlvdfqggfmmrppjfrmtfwhpbjsppszjbhmthndqmvbmqcbtqsltwrcvlvblwspbgspjftwllzcmnsrvjpnstzrfmcflnhppsdfggwbzvnvlnjqlfvrlplnzvfrwvgcgqvnpfgtgchctvhcplclzmfpwgnfhqjgglfmsgpflqcpqmbbhwnvvdllcnhblpnndbdtmgvfbvvvlvzlrpfqmnvzbfrssjtlgcjtpfznshvdjrjnfshfcgvwcdbqlfsbhnzwmsgwhpbzttgfjsqgwvdmbdwjljhsndrbbzfrsqjhcbldzqpmtnfvnmzltjcrvrltwshnhqlnclmcnfpbzstsczlqmfmdftzfbcwqnhqppzfbzpbfjhmmtvtbmblmtshsbtjlvsqvmbmgstbbdmhprqmtpfdqqntmnlbmpsmwfgrvstjcllhwpcddnljdjvdrbwqmgrjnldpgnrhgqpzqrvwzsngrgnbpjnsffzjsbdptwnnfcqlscfhvggpfstsnqzcjbqqhgdpqsrlprcppgqmddpqpbnvgwtdqsbbgtvsqfrtqfsbdzhsztfmvwrrsjcbtcjgzrnhnpgldtwbwgmwbgmjjzsbbzlhgmlczrhjwtzrgwscmjvlstprldhglvftqzbtrmcwzgtjppbnjcdvjvcwvdbngnbrmjvvtnwdqfclbpgzcfnnnlnngtgmhsqsdmbjctjzjpbrwrhscqshmmwbtfnzjgsrjlnqqdsvdrjdzsdprphnfmwwcztqfcrjvnfhlvnqwbrfmcvhrbtgvcrqjjfcnzwmlfzzdcbbzvphhmsdltwjfdcgthpvszqzjdbfwrpvhbjqdhrscnvjhjvvcldnhgjclmzpbrrwnscgpcqrpdgsnjnwhctcdqgwqbrcszfzpmtdrhlftvwffdjrtznqrppqbdbwvzmtlpvsqqpcngjgfdrpngnspdwhhvlhqrtsphgqrlldggtrvqsprbfdmrpgcmqphdvjfmhlznpgtqlvtnllcdhzhhtjjlfvdlwhcrfmjmdjtmbllvsfgvmfqtqlmrlrjmqptszvjdpzhphppljnpjdjpwlrclssgdnstchhwhpflmlrtdqvqbbljrmnflrltzpqmgqfrczvfzrpfsrwsgpljvjfjdjdvjchcdmmtjgghqspwzdtwqgtvmnrrbfbgnhcrvnzznrdlqmgmdwmpwzlqdjtvpszwnjtjtmjqvfwvftlhgpvgzswpbvbllfcwpjnsmbhzrdpdzjsrpnhphdcqjmzvvhrjcwhgwjwcshqwzpbpmfnjjvqcjrqmvsrdrtdvfhwhrbpvrqrsfzflslqtdrtcsggtzmpvbszdgttlvpwwltvpcwqmnwqtpcfzgsvsmncvpqqdrljfwtncplmjlpfcnqmcctwzhrbmrfwvsrjsbnhjrjmrnbmmnnhsvlltwzzhsgwppnlmljgpcsmpchdjdzpgvrtwsfzffhnlbfmrldzbshvpqhnfzpwnvczgfvhbntcpztwqlfgtsmdhvcrgjhvqrhbpvbpzcpbgzrcfjztbnfjptbzfpztwprwf`
}

func (td dayEntry) PuzzleSpec() string {
	return ``
}
