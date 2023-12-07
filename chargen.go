package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

//func statdice() gera um valor randomico composto da soma de quatro valores entre 1 e 6 menos o menor valor

func statdice() int {
	//Esta parte utiliza o rand.Seed para gerar quatro valores aleatorios entre 1 e 6 (dicea, diceb, dicec, diced)
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 6
	var dicea = (rand.Intn(max-min+1) + min)
	var diceb = (rand.Intn(max-min+1) + min)
	var dicec = (rand.Intn(max-min+1) + min)
	var diced = (rand.Intn(max-min+1) + min)

	//Aqui juntamos os quatro valores em um array diceArray
	var diceArray = [4]int{dicea, diceb, dicec, diced}

	//Nesta parte é feito o sort dos valores em ordem decrescente
	sort.Slice(diceArray[:], func(i, j int) bool {
		return diceArray[:][i] > diceArray[:][j]
	})

	//Aqui iniciamos e damos valor para o newDiceArray igual ao diceArray menos o ultimo valor(numero menor) 
	var newDiceArray = diceArray[:len(diceArray)-1]

	//Inicializamos o valor final diceSum somente no final
	var diceSum int

	//Setamos o valor do diceSum igual a soma dos valores do array newDiceArray
	for i := 0; i < len(newDiceArray); i++ {
		diceSum = diceSum + newDiceArray[i]
	}

	//Retorna o diceSum
	return diceSum
}

func main() {

	//Aqui se gera um  sleepDuration randomico entre 1 a 5 milisegundos
	minSleep := 1
	maxSleep := 5
	randomSleep := rand.Intn(maxSleep-minSleep+1) + minSleep

	sleepDuration := time.Duration(randomSleep)

	//Aqui se inicializam e definem os 6 valores dos status de acordo com a função statdice()
	str := statdice()
	time.Sleep(sleepDuration)
	dex := statdice()
	time.Sleep(sleepDuration)
	con := statdice()
	time.Sleep(sleepDuration)
	wis := statdice()
	time.Sleep(sleepDuration)
	inl := statdice()
	time.Sleep(sleepDuration)
	cha := statdice()
		/*Na primeira versão deste programa, foi constatado que, devido a natureza da função rand.Seed
	utilizada para gerar os valores (a mesma gera um valor randomico baseado em uma seed, que por si é derivada do valor atual da Unix time)
	os 6 valores dos status ficavam iguais (os valores eram gerados simultaneamente e, por serem gerados com a mesma seed, retornavam iguais)
	por isso foi implementado uma função time.Sleep entre os status, para que cada valor seja gerado com uma seed diferente*/

	/*Em sua parte final definimos um limite máximo de pontos totais, caso a soma de todos os status seja maior que o valor estipulado,
	o console printa uma mensagem de erro, caso esteja dentro dos parametros estipulados ele printa os valores finais ordenados*/
	if str+dex+con+wis+inl+cha >= 80 {
		fmt.Print("Personagem muito OP, role novamente.")
	} else {
		fmt.Println("STR:", str)
		fmt.Println("DEX:", dex)
		fmt.Println("CON:", con)
		fmt.Println("WIS:", wis)
		fmt.Println("INT:", inl)
		fmt.Println("CHA:", cha)
	}
}
