package main

import (
	"io/ioutil"
	"log"

	"github.com/emman27/aoc2017/circus"
	"github.com/emman27/aoc2017/corruptionChecksum"
	"github.com/emman27/aoc2017/inverseCaptcha"
	"github.com/emman27/aoc2017/knothash"
	"github.com/emman27/aoc2017/memoryReallocation"
	"github.com/emman27/aoc2017/passphrase"
	"github.com/emman27/aoc2017/registers"
	"github.com/emman27/aoc2017/spiralMemory"
	"github.com/emman27/aoc2017/stream"
	"github.com/emman27/aoc2017/trampolines"
)

func main() {

	log.Println(inverseCaptcha.InverseCaptcha("5672987533353956199629683941564528646262567117433461547747793928322958646779832484689174151918261551689221756165598898428736782194511627829355718493723961323272136452517987471351381881946883528248611611258656199812998632682668749683588515362946994415852337196718476219162124978836537348924591957188827929753417884942133844664636969742547717228255739959316351852731598292529837885992781815131876183578461135791315287135243541659853734343376618419952776165544829717676988897684141328138348382882699672957866146524759879236555935723655326743713542931693477824289283542468639522271643257212833248165391957686226311246517978319253977276663825479144321155712866946255992634876158822855382331452649953283788863248192338245943966269197421474555779135168637263279579842885347152287275679811576594376535226167894981226866222987522415785244875882556414956724976341627123557214837873872723618395529735349273241686548287549763993653379539445435319698825465289817663294436458194867278623978745981799283789237555242728291337538498616929817268211698649236646127899982839523784837752863458819965485149812959121884771849954723259365778151788719941888128618552455879369919511319735525621198185634342538848462461833332917986297445388515717463168515123732455576143447454835849565757773325367469763383757677938748319968971312267871619951657267913817242485559771582167295794259441256284168356292785568858527184122231262465193612127961685513913835274823892596923786613299747347259254823531262185328274367529265868856512185135329652635938373266759964119863494798222245536758792389789818646655287856173534479551364115976811459677123592747375296313667253413698823655218254168196162883437389718167743871216373164865426458794239496224858971694877159591215772938396827435289734165853975267521291574436567193473814247981877735223376964125359992555885137816647382139596646856417424617847981855532914872251686719394341764324395254556782277426326331441981737557262581762412544849689472281645835957667217384334435391572985228286537574388834835693416821419655967456137395465649249256572866516984318344482684936625486311718525523265165"))

	log.Println(corruptionChecksum.Calculate([][]int{
		[]int{1919, 2959, 82, 507, 3219, 239, 3494, 1440, 3107, 259, 3544, 683, 207, 562, 276, 2963},
		[]int{587, 878, 229, 2465, 2575, 1367, 2017, 154, 152, 157, 2420, 2480, 138, 2512, 2605, 876},
		[]int{744, 6916, 1853, 1044, 2831, 4797, 213, 4874, 187, 6051, 6086, 7768, 5571, 6203, 247, 285},
		[]int{1210, 1207, 1130, 116, 1141, 563, 1056, 155, 227, 1085, 697, 735, 192, 1236, 1065, 156},
		[]int{682, 883, 187, 307, 269, 673, 290, 693, 199, 132, 505, 206, 231, 200, 760, 612},
		[]int{1520, 95, 1664, 1256, 685, 1446, 253, 88, 92, 313, 754, 1402, 734, 716, 342, 107},
		[]int{146, 1169, 159, 3045, 163, 3192, 1543, 312, 161, 3504, 3346, 3231, 771, 3430, 3355, 3537},
		[]int{177, 2129, 3507, 3635, 2588, 3735, 3130, 980, 324, 266, 1130, 3753, 175, 229, 517, 3893},
		[]int{4532, 164, 191, 5169, 4960, 3349, 3784, 3130, 5348, 5036, 2110, 151, 5356, 193, 1380, 3580},
		[]int{2544, 3199, 3284, 3009, 3400, 953, 3344, 3513, 102, 1532, 161, 143, 2172, 2845, 136, 2092},
		[]int{194, 5189, 3610, 4019, 210, 256, 5178, 4485, 5815, 5329, 5457, 248, 5204, 4863, 5880, 3754},
		[]int{3140, 4431, 4534, 4782, 3043, 209, 216, 5209, 174, 161, 3313, 5046, 1160, 160, 4036, 111},
		[]int{2533, 140, 4383, 1581, 139, 141, 2151, 2104, 2753, 4524, 4712, 866, 3338, 2189, 116, 4677},
		[]int{1240, 45, 254, 1008, 1186, 306, 633, 1232, 1457, 808, 248, 1166, 775, 1418, 1175, 287},
		[]int{851, 132, 939, 1563, 539, 1351, 1147, 117, 1484, 100, 123, 490, 152, 798, 1476, 543},
		[]int{1158, 2832, 697, 113, 121, 397, 1508, 118, 2181, 2122, 809, 2917, 134, 2824, 3154, 2791},
	}, corruptionChecksum.EvenDivisResults))
	log.Println(spiralMemory.PartA(325489))
	log.Println(spiralMemory.PartB(325489))
	log.Println(passphrase.PartA(passphrase.ActualData))
	log.Println(passphrase.PartB(passphrase.ActualData))
	// log.Println(trampolines.PartA(trampolines.TestData))
	log.Println(trampolines.PartB(trampolines.TestData))
	log.Println(memoryReallocation.PartA([]int{4, 1, 15, 12, 0, 9, 9, 5, 5, 8, 7, 3, 14, 5, 12, 3}))
	log.Println(memoryReallocation.PartB([]int{4, 1, 15, 12, 0, 9, 9, 5, 5, 8, 7, 3, 14, 5, 12, 3}))
	log.Println(circus.PartA("./circus/puzzle_input.txt"))
	log.Println(circus.PartB("./circus/puzzle_input.txt"))
	log.Println(registers.PartA("./registers/puzzle_input.txt"))
	log.Println(registers.PartB("./registers/puzzle_input.txt"))

	s, err := ioutil.ReadFile("./stream/puzzle_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(stream.PartA(string(s)))
	log.Println(stream.PartB(string(s)))
	log.Println(knothash.PartA([]int{46, 41, 212, 83, 1, 255, 157, 65, 139, 52, 39, 254, 2, 86, 0, 204}))
}
