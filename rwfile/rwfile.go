package rwfile

import (
	"os"
)

func Write(filename , history string){
	err := os.Mkdir("./users/"+filename, 0755)
	if err != nil{
		os.Mkdir("./users/", 0755)
		os.Mkdir("./users/"+filename, 0755)
	}
	os.WriteFile("./users/"+filename+"/"+filename+".json", []byte(history), 0644)
}
func Read(filename string)string{
	history , _ := os.ReadFile("./users/"+filename+"/"+filename+".json")

	return string(history)
}
func Check(filename string)bool{
	_, err := os.Stat("./users/"+filename+"/"+filename+".json")
	return err == nil
}


func main() {
	
}