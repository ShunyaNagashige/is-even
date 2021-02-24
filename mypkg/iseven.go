package mypkg

type MyError string

func IsEven(n int)(bool,error){
	if(n%2==0){
		return true,nil
	}
	return false,nil
}