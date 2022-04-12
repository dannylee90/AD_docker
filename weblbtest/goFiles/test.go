package main

import(
	"os"
	"fmt"
	"strconv"
	"math/rand"
	"time"
	"reflect"
)

var path="/usr/local/apache2/htdocs/index.html"
//var path="/Users/danny/docker/test/index.html"

func main(){
	writeFile()
}


func writeFile(){

	// 시간값을 시드로 활용해  100000 이하의 랜덤 숫자 생성
	timeSource := rand.NewSource(time.Now().UnixNano())
	random := rand.New(timeSource)
	no := random.Intn(100000)
	fmt.Println(no)
	fmt.Println(reflect.TypeOf(no))
	
	// 생성 된 랜덤 숫자를 문자로 변환
	int2str := strconv.Itoa(no)
	fmt.Println(int2str)
	fmt.Println(reflect.TypeOf(int2str))


	// 파일 생성 후 랜덤생성 된 문자열 숫자를 쓰고 파일닫음
	f, err := os.Create(path)
	
	if err != nil{
		return
	}
	defer f.Close()

	_, err= f.WriteString(int2str)
	if isError(err){
		return
	}
	
	err = f.Sync()
	if isError(err){
		return
	}
	
	fmt.Println("File created")

}

func isError(err error) bool{
	if err != nil{
		fmt.Println(err.Error())
	}
	return (err != nil)
}
