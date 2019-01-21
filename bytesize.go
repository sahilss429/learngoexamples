package main
import "fmt"

const(
	Bits = 1
	B = Bits * 8
	KB = B * 1024
	MB = KB * 1024
	GB = MB * 1024
	TB = GB * 1024
	PB = TB * 1024
)

type ByteSize float64

func (b ByteSize) String() string {
	switch {
	case b >= PB:
		return "Very Big"
	case b >= TB:
		return fmt.Sprintf("%.2fTB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%.2fGB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2fMB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.2fKB", b/KB)
	case b >= B:
		return fmt.Sprintf("%.2fB", b/B)
	}
	return fmt.Sprintf("%.fBits", b)
}
func main(){
fmt.Println(ByteSize(2048))
fmt.Println(ByteSize(3492103452))
fmt.Println(ByteSize(1021))
}
