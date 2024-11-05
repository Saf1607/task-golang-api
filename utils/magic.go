package utils

func MagicSum(n int) int {
	return n + n
}

// menerima satu angka n dan menghitung n pangkat n.
func MagicPow(n int) int {
	result := 1
	for i := 0; i < n; i++ {
		result *= n
	}
	return result
}

func MagicOdd(n int) bool {
	return n%2 != 0
}

func MagicGrade(n int) string {
	switch n {
	case 0:
		return "Zero"
	case 1:
		return "Bad"
	case 2:
		return "Ok"
	case 3:
		return "Nice"
	case 4:
		return "Awesome"
	case 5:
		return "Exceptional"
	default:
		return "Unknown"
	}
}

func MagicName(n int) []string {

	myname := make([]string, n)
	for i := 0; i < n; i++ {
		myname[i] = "safira"
	}
	return myname
}

// menghitung jumlah dari 0 hingga n dengan menambahkan semua angka dari 0 sampai n ke dalam variabel sum.
func MagicTria(n int) int {

	sum := 0

	for i := 0; i <= n; i++ {
		sum += i
	}
	return sum
}

func MagicChange(n *int) {

	*n *= 2
}

type MagicNumber struct {
	Number int
}

func (m *MagicNumber) Multiply(n int) {
	m.Number *= n
}
