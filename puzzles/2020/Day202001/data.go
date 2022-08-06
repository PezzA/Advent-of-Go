package Day202001

var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string, int) {
	return 2020, 1, "Report Repair", 2
}

func (td dayEntry) PuzzleSpec() string {

	return `--- Day 1: Report Repair ---
After saving Christmas five years in a row, you've decided to take a vacation at a nice resort on a tropical island. Surely, Christmas will go on without you.

The tropical island has its own currency and is entirely cash-only. The gold coins used there have a little picture of a starfish; the locals just call them stars. None of the currency exchanges seem to have heard of them, but somehow, you'll need to find fifty of these coins by the time you arrive so you can pay the deposit on your room.

To save your vacation, you need to get all fifty stars by December 25th.

Collect stars by solving puzzles. Two puzzles will be made available on each day in the Advent calendar; the second puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!

Before you leave, the Elves in accounting just need you to fix your expense report (your puzzle input); apparently, something isn't quite adding up.

Specifically, they need you to find the two entries that sum to 2020 and then multiply those two numbers together.

For example, suppose your expense report contained the following:

1721
979
366
299
675
1456
In this list, the two entries that sum to 2020 are 1721 and 299. Multiplying them together produces 1721 * 299 = 514579, so the correct answer is 514579.

Of course, your expense report is much larger. Find the two entries that sum to 2020; what do you get if you multiply them together?


--- Part Two ---
The Elves in accounting are thankful for your help; one of them even offers you a starfish coin they had left over from a past vacation. They offer you a second one if you can find three numbers in your expense report that meet the same criteria.

Using the above example again, the three entries that sum to 2020 are 979, 366, and 675. Multiplying them together produces the answer, 241861950.

In your expense report, what is the product of the three entries that sum to 2020?
`
}

func (td dayEntry) PuzzleInput() string {
	return `1833
1949
1745
1863
1422
1980
1695
1233
1407
1971
1486
1922
1802
1689
214
1864
1091
1395
1772
1901
1542
1730
1782
1815
1818
1236
1604
1219
1834
1813
1532
1963
2003
1149
1577
1408
1871
1417
1727
1155
1953
1287
1259
1548
1126
1927
1225
1172
11
1301
1869
1808
1238
1989
1027
321
1564
636
1847
1877
1716
1275
1738
1054
1966
1019
1256
1940
1821
1914
1556
1389
1020
1293
1935
1804
1945
508
1856
1674
1721
1541
1435
1852
1394
2006
1366
1473
1274
623
1701
1842
1954
1999
1195
1246
1967
1153
1851
1294
1152
1812
1732
1030
1956
1132
1948
1865
1835
1231
1975
1759
1843
1379
1657
1267
1062
1850
1817
1543
1805
1252
1974
1161
876
1647
1796
1634
1177
1519
1527
1249
1158
2007
1702
1714
1040
1826
1837
1361
1814
1858
1828
1951
1140
1845
1476
1337
1262
1923
1397
1025
1412
1930
1164
1300
1369
1777
1591
1919
1874
1482
2010
1957
1897
1854
1992
1735
1786
1661
1404
1254
1803
1794
1614
1711
1007
1979
1928
1505
2001
1094
2005
1297
1933
1976
1104
1279
1012
1943
1679
1958
1841
1623
1809
1800
919
1620
1936
1209`
}
