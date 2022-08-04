package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	records, _ := readData("CH001_01.csv")
	for idx, rec := range records {
		fmt.Println(idx, ": ", rec)
		if idx == 20 {
			log.Fatalln("끝!")
		}
	}
}

func readData(fileName string) ([][]string, error) {

	f, err := os.Open(fileName)

	if err != nil {
		log.Fatalf("파일 불러오기 실패")
	}

	// 프로그램 종료하면 파일도 함께 종료하기 (메모리 낭비 방지)
	defer f.Close()

	r := csv.NewReader(f)

	// 첫줄 생략하기
	// Read()는 한 라인을 읽어 들이고, ReadAll()은 전체를 한꺼번에 읽어들인다.
	// r은 f라는 파일을 가진 변수고 이 r에서 Read()를 통해 첫 줄을 가져와서
	// 오류가 나지 않는다면 읽은 줄을 없앰
	// 단순히 r.Read()는 r1 = r[0]이 아니라 r1 = r.popleft() 방식을 동작
	if _, err := r.Read(); err != nil {
		log.Fatalf("파일의 첫 데이터 불러오기 실패")
	}
	records, err := r.ReadAll()

	if err != nil {
		log.Fatalf("파일 데이터 불로오기 실패")
	}

	return records, nil
}
