package misollar

import (
	"fmt"
	"math"
)

func Juft() {
	var son1, son2, son3 int =12,24,65
	var juft,toq []int
	if son1%2== 0{
		juft = append(juft, son1)
	}else {
		toq = append(toq, son1)
	}
	if son2%2== 0{
		juft = append(juft, son2)
	}else {
		toq = append(toq, son2)
	}
	if son3%2== 0{
		juft = append(juft, son3)
	}else {
		toq = append(toq, son3)
	}
	a:= len(toq)
	b:=len(juft)
	fmt.Println("juft", b)
	fmt.Println("toq", a)
}
func Tub() {
	var sonlar []int 
    var son int = 992
for a:=1 ; a<=son/2 ; a++ {
        if son%a == 0 {
			sonlar = append(sonlar, a)
		}
}
if len(sonlar)<=2 {
 fmt.Println("tub son: ", son)
}else {
	fmt.Println("tub son emas: ", son)
}
}
func Gacha() {
	var son1 int =1000
	var tubl,sonlar int 
   for son:=2 ; son1>son ; son++{
	 for a:=1 ; a<=son/2 ; a++ {
        if son%a == 0 {
			sonlar +=1	
		}
     }
        if sonlar<=1 {	
         tubl+=1
        }
         sonlar=0
    }
fmt.Println(tubl)
}
func Orasi() {
	var tubl,sonlar int
	var num1, num2 int = 2, 100
	for son:=num1 ; num2>son ; son++{
		for a:=1 ; a<=son/2 ; a++ {
		   if son%a == 0 {
			   sonlar +=1	
		   }
		}
		 if sonlar<=1 {	
			tubl+=1
		   }
			sonlar=0
	   }
		   fmt.Println(tubl)
}
func Teskari() {
 var qoldiq,teskari int 
 var son int = 64665466548
 for teskari= 0; son!=0 ;teskari *=10{
	qoldiq = son%10
	son/=10
	teskari=teskari+qoldiq
	 
 }
 teskari/=10
 fmt.Println(teskari)
}
func Kabisa() {
	var kabisa []int
	var boshi, oxiri,yil int = 1945, 2024,0
	for yil = boshi; yil<=oxiri; yil++ {
	if yil%4==0 {
		if yil%100!=0 || yil%400==0 {
	      kabisa = append(kabisa, yil)
		}
    }
 }
 fmt.Println("kabisa yillar: ", kabisa)
}
func BinaryToDesimal() {
	var binary string = "11111111"
	var desimal,son float32
	var a float64
	b:= len(binary)
	for a=0 ; 1<b; a++ {
		b-=1
		fmt.Println(a)
     }
	for _,v := range binary {
		
		if v == '1' {
			son = float32(math.Pow(2, a))
          desimal+=son
		  
		} 
		a-=1 
	}
	fmt.Println(desimal)
}
func DesimalToBinary() {
	var desimal,binary int
	var teskari []int 
	desimal= 1555
	for desimal>0 {
		qoldiq:= desimal%2
		desimal/=2
		teskari = append(teskari, qoldiq)
	}
	uz:=len(teskari)-1
   for uz>0 {
	binary*=10
	binary = binary+teskari[uz]
	uz-=1	
   }
   fmt.Println(binary)	
}
func MukammalSon() {
var son, sum int= 28,0
	var sonlar []int
	for a := 1; a <= (son / 2); a++ {
        if son%a == 0 {
			sonlar = append(sonlar, a)
			sum = sum + a
		}
	}
	if sum == son {
		fmt.Println("mukammal son", sonlar, "=", son)
	} else {
		fmt.Println("mukammal son emas: ", son, "teng emas", sum)
	}
}