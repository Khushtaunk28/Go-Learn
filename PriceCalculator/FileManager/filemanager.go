package filemanager

import (
	"bufio"
	"errors"
	"os"
)
func ReadLines(path string) ([]string,error){
	file,err:= os.Open(path)
	if(err!=nil){
		return nil,errors.New("Error while opening file")
	}
	scanner:=bufio.NewScanner(file)
	var line []string
	for scanner.Scan(){
		line =append(line,scanner.Text())

	}
	err=scanner.Err()
	if err!=nil{
		file.Close()
		return nil,errors.New("Error while reading file")
	}
	file.Close()
	return line,nil

}