package plugins

import (
	"encoding/csv"
	//"io"
	"fmt"
	"os"
)

func WriteCSV(path string, startline []string, data [][]string) {
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM

	w := csv.NewWriter(f)
	w.Write(startline)

	for i := 0; i < len(data); i++ {
		w.Write(data[i])
	}

	w.Flush()
}

func ReadCSV(path string) (data [][]string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("ReadCSV1 Error:", err)
		//logger.Debugln("Error:", err)
		return
	}
	defer file.Close()
	reader := csv.NewReader(file)

	data, err2 := reader.ReadAll()

	if err2 != nil {
		fmt.Printf("ReadCSV2 Error:", err)
		//logger.Debugln("Error:", err2)
		return
	}
	return data
}
