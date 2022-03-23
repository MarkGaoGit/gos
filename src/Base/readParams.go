package Base

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func ReadParams() {

	//使用buffer读取文件
	//readFileByBuffer()

	//使用切片读取文件
	//readFileBySlice()

	//使用缓冲区读写文件

}

func catBySlice(f *os.File) {
	const NBUF = 521
	var buf [NBUF]byte
	for {
		switch nr, err := f.Read(buf[:]); true {
		case nr < 0:
			fmt.Fprintf(os.Stderr, "cat: error reading: %s\n", err.Error())
			os.Exit(1)
		case nr == 0:
			continue
		case nr > 0:
			if nw, ew := os.Stdout.Write(buf[0:nr]); nw != nr {
				fmt.Fprintf(os.Stderr, "cat: error writing: %s\n", ew.Error())
			}
		}
	}
}

func readFileBySlice() {
	flag.Parse()
	if flag.NArg() == 0 {
		catBySlice(os.Stdin)
	}

	for i := 0; i < flag.NArg(); i++ {
		f, _ := os.Open(flag.Arg(i))
		if f == nil {
			os.Exit(1)
		}
		catBySlice(f)
		f.Close()
	}
}

func readFileByBuffer() {
	flag.Parse()
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
	}

	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s:error reading from %s: %s\n", os.Args[0], flag.Arg(i), err.Error())
			continue
		}
		cat(bufio.NewReader(f))
		f.Close()
	}
}

func cat(r *bufio.Reader) {
	for {
		buf, err := r.ReadBytes('\n')
		fmt.Fprintf(os.Stdout, "%s", buf)
		if err == io.EOF {
			break
		}
	}
}

// 从命令行读取参数
func readS() {
	var newLineBool = flag.Bool("n", false, "print new line")

	space := " "
	newLine := "\n"

	flag.PrintDefaults()
	flag.Parse()
	fmt.Println(flag.NFlag())

	s := ""

	for i := 0; i < flag.NFlag(); i++ {
		if i > 0 {
			s += space
			if *newLineBool {
				s += newLine
			}
		}

		s += flag.Arg(i)
	}

	os.Stdout.WriteString(s)
}
