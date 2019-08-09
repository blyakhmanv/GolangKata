package main

import "fmt"
import "regexp"
import "strings"

func main() {
	var dir = "/+11-541-914-3010 Alphand_St. <J Steeve>\n 133, Green, Rd. <E Kustur> NY-56423 ;+1-541-914-3010\n" + "+1-541-984-3012 <P Reed> /PO Box 530; Pollocksville, NC-28573\n :+1-321-512-2222 <Paul Dive> Sequoia Alley PQ-67209\n" + "+1-741-984-3090 <Peter Reedgrave> _Chicago\n :+1-921-333-2222 <Anna Stevens> Haramburu_Street AA-67209\n" + "+1-111-544-8973 <Peter Pan> LA\n +1-921-512-2222 <Wilfrid Stevens> Wild Street AA-67209\n" + "<Peter Gone> LA ?+1-121-544-8974 \n <R Steell> Quora Street AB-47209 +1-481-512-2222\n" + "<Arthur Clarke> San Antonio $+1-121-504-8974 TT-45120\n <Ray Chandler> Teliman Pk. !+1-681-512-2222! AB-47209,\n" + "<Sophia Loren> +1-421-674-8974 Bern TP-46017\n <Peter O'Brien> High Street +1-908-512-2222; CC-47209\n" + "<Anastasia> +48-421-674-8974 Via Quirinal    Roma\n <P Salinger> Main Street, +1-098-512-2222, Denver\n" + "<C Powel> *+19-421-674-8974 Chateau des Fosses Strasbourg F-68000\n <Bernard Deltheil> +1-498-512-2222; Mount Av.  Eldorado\n" + "+1-099-500-8000 <Peter Crush> Labrador Bd.\n +1-931-512-4855 <William Saurin> Bison Street CQ-23071\n" + "<P Salinge> Main Street, +1-098-512-2222, Denve\n" + "<P Salinge> Main Street, +1-098-512-2222, Denve\n"
	fmt.Println(Phone(dir, "1-541-914-3010"))
	fmt.Println(Phone(dir, "1-098-512-2222"))
	fmt.Println(Phone(dir, "5-555-555-5555"))
	fmt.Println(Phone(dir, "1-908-512-2222"))
	fmt.Println(Phone(dir, "1-498-512-2222"))

}

//Phone - kata function
func Phone(dir, num string) string {
	//space regexp
	space := regexp.MustCompile(`\s+`)

	//find duplicates and missing numbers
	pnoneNumberReg := regexp.MustCompile("\\+" + num)
	i := len(pnoneNumberReg.FindAllStringIndex(dir, -1))
	switch {
	case i == 0:
		return "Error => Not found: " + num
	case i > 1:
		return "Error => Too many people: " + num
	}

	var nameRec string
	//split lines into slice
	lines := strings.Split(dir, "\n")
	for _, line := range lines {
		if len(pnoneNumberReg.FindAllStringIndex(line, -1)) == 1 {

			//find Name from the line
			nameRec = regexp.MustCompile(`<.*>`).FindString(line)

			//remove Name,phone number and extra spaces form the line
			line = strings.Replace(line, "+"+num, "", -1)
			line = strings.Replace(line, nameRec, "", -1)
			line = regexp.MustCompile(`\.`).ReplaceAllString(line, ". ")
			line = regexp.MustCompile(`[^0-9A-Za-z_\-\.]`).ReplaceAllString(line, " ")
			line = space.ReplaceAllString(line, " ")
			line = regexp.MustCompile(`_`).ReplaceAllString(line, " ")
			line = strings.TrimSpace(line)

			//remove <> fron nameRec
			nameRec = nameRec[1 : len(nameRec)-1]

			return "Phone => " + num + ", Name => " + nameRec + ", Address => " + line
		}
	}

	//pnoneNumber := regexp.MustCompile(`[\d\d]-\d\d\d-\d\d\d-\d\d\d\d`)
	//fmt.Println(pnoneNumber.FindAllString(dir, -1))
	return ""
}
