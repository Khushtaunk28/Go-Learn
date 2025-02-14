package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)
type FileManager struct{
	InputFilePath string
	OutputFilePath string
}
func (fm FileManager) ReadLines() ([]string,error){
	file,err:= os.Open(fm.InputFilePath)
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

func (fm FileManager) WriteResult(data interface{}) error{
	file,err:=os.Create(fm.OutputFilePath)
	if(err!=nil){
		return errors.New("Error while writing into file")
	}
	encoder:= json.NewEncoder(file)
	err = encoder.Encode(data)
	if(err!=nil){
		return errors.New("Error while encoding into json")
	}
	file.Close()
	return nil
}

func NewFileManager(inputPath,outputPath string) FileManager {
	return FileManager{
		InputFilePath: inputPath,
		OutputFilePath: outputPath,
	}
}