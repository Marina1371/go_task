package main
   
import (
	"crypto/md5"
	"encoding/hex"
   )
func mdHashing(input string) string {
byteInput := []byte(input)
md5Hash := md5.Sum(byteInput)
return hex.EncodeToString(md5Hash[:]) // by referring to it as a string
   }
func main ()  { 
fmt.Println(mdHashing( input:"abcdefghijklmnopqrstuvwxyz"))
fmt.Println(mdHashing( input:"ABCDEFGHIJKLMNOPQRSTUVWXYZ"))
}
   



